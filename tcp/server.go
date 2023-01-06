package tcp

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/uim"
	"github.com/uim/logger"
	"github.com/uim/naming"
)

type ServerOptions struct {
	loginwait time.Duration
	readwait  time.Duration
	writewait time.Duration
}

type Server struct {
	listen string
	naming.ServiceRegistration
	uim.ChannelMap
	uim.Acceptor
	uim.MessageListener
	uim.StateListener
	once    sync.Once
	options ServerOptions
	quit    *uim.Event
}

var _ uim.Server = (*Server)(nil)

func NewServer(listen string, service naming.ServiceRegistration) uim.Server {
	return &Server{
		listen:              listen,
		ServiceRegistration: service,
		ChannelMap:          uim.NewChannels(100),
		quit:                uim.NewEvent(),
		options: ServerOptions{
			loginwait: uim.DefaultLoginWait,
			readwait:  uim.DefaultReadWait,
			writewait: uim.DefaultWriteWait,
		},
	}
}

func (s *Server) Start() error {
	log := logger.WithFields(logger.Fields{
		"module": "tcp.server",
		"listen": s.listen,
		"id":     s.ServiceID(),
	})

	if s.StateListener == nil {
		return fmt.Errorf("StateListener is nil")
	}
	if s.Acceptor == nil {
		s.Acceptor = new(defaultAcceptor)
	}

	// step 1.
	lst, err := net.Listen("tcp", s.listen)
	if err != nil {
		return err
	}
	log.Info("started")
	for {
		// step 2.
		rawconn, err := lst.Accept()
		if err != nil {
			rawconn.Close()
			log.Warn(err)
			continue
		}

		go func(rawconn net.Conn) {
			conn := NewConn(rawconn)
			// step 3.
			id, err := s.Accept(conn, s.options.loginwait)
			if err != nil {
				_ = conn.WriteFrame(uim.OpClose, []byte(err.Error()))
				conn.Close()
				return
			}
			if _, ok := s.Get(id); ok {
				log.Warnf("channel %s existed", id)
				_ = conn.WriteFrame(uim.OpClose, []byte("channelId is repeated"))
				conn.Close()
				return
			}
			// step 4.
			channel := uim.NewChannel(id, conn)
			channel.SetReadWait(s.options.readwait)
			channel.SetWriteWait(s.options.writewait)
			s.Add(channel)

			log.Info("accept", channel)
			// step 5.
			err = channel.Readloop(s.MessageListener)
			if err != nil {
				log.Info(err)
			}
			// step 6.
			s.Remove(channel.ID())
			_ = s.Disconnect(channel.ID())
			channel.Close()
		}(rawconn)
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	log := logger.WithFields(logger.Fields{
		"module": "tcp.server",
		"id":     s.ServiceID(),
	})
	s.once.Do(func() {
		defer func() {
			log.Infoln("shutdown")
		}()
		// close channels
		channels := s.ChannelMap.All()
		for _, ch := range channels {
			ch.Close()

			select {
			case <-ctx.Done():
				return
			default:
				continue
			}
		}
	})
	return nil
}

func (s *Server) Push(id string, data []byte) error {
	ch, ok := s.ChannelMap.Get(id)
	if !ok {
		return errors.New("channel no found")
	}
	return ch.Push(data)
}

func (s *Server) SetAcceptor(acceptor uim.Acceptor) {
	s.Acceptor = acceptor
}

func (s *Server) SetMessageListener(listener uim.MessageListener) {
	s.MessageListener = listener
}

func (s *Server) SetStateListener(listener uim.StateListener) {
	s.StateListener = listener
}

func (s *Server) SetReadWait(readwait time.Duration) {
	s.options.readwait = readwait
}

func (s *Server) SetChannelMap(channels uim.ChannelMap) {
	s.ChannelMap = channels
}

type defaultAcceptor struct{}

func (a *defaultAcceptor) Accept(conn uim.Conn, timeout time.Duration) (string, error) {
	return ksuid.New().String(), nil
}
