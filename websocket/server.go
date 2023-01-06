package websocket

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/segmentio/ksuid"
	"github.com/uim"
	"github.com/uim/logger"
	"github.com/uim/naming"
)

// ServerOptions
type ServerOptions struct {
	loginwait time.Duration
	readwait  time.Duration
	writewait time.Duration
}

// Server is a websocket implement of the Server.
type Server struct {
	listen string
	naming.ServiceRegistration
	uim.ChannelMap
	uim.Acceptor
	uim.MessageListener
	uim.StateListener
	once    sync.Once
	options ServerOptions
}

var _ uim.Server = (*Server)(nil)

// NewServer
func NewServer(listen string, service naming.ServiceRegistration) uim.Server {
	return &Server{
		listen:              listen,
		ServiceRegistration: service,
		options: ServerOptions{
			loginwait: uim.DefaultLoginWait,
			readwait:  uim.DefaultReadWait,
			writewait: uim.DefaultWriteWait,
		},
	}
}

// Start server.
func (s *Server) Start() error {
	mux := http.NewServeMux()
	log := logger.WithFields(logger.Fields{
		"module": "ws.server",
		"listen": s.listen,
		"id":     s.ServiceID(),
	})

	if s.Acceptor == nil {
		s.Acceptor = new(defaultAcceptor)
	}
	if s.StateListener == nil {
		return fmt.Errorf("StateListener is nil")
	}
	if s.ChannelMap == nil {
		s.ChannelMap = uim.NewChannels(100)
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// step 1.
		rawconn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			resp(w, http.StatusBadRequest, err.Error())
			return
		}
		// step 2.
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
		channel.SetWriteWait(s.options.writewait)
		channel.SetReadWait(s.options.readwait)
		s.Add(channel)
		go func(ch uim.Channel) {
			// step 5.
			err := ch.Readloop(s.MessageListener)
			if err != nil {
				log.Info(err)
			}
			// step 6.
			s.Remove(ch.ID())
			err = s.Disconnect(ch.ID())
			if err != nil {
				log.Warn(err)
			}
			ch.Close()
		}(channel)
	})
	log.Infoln("started")
	return http.ListenAndServe(s.listen, mux)
}

func (s *Server) Shutdown(ctx context.Context) error {
	log := logger.WithFields(logger.Fields{
		"module": "ws.server",
		"id":     s.ServiceID(),
	})
	s.once.Do(func() {
		defer func() {
			log.Infoln("shutdown")
		}()
		// Close channels.
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
		return errors.New("channel not found")
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

func (s *Server) SetChannelMap(channels uim.ChannelMap) {
	s.ChannelMap = channels
}

// Set read wait duration.
func (s *Server) SetReadWait(readwait time.Duration) {
	s.options.readwait = readwait
}

func resp(w http.ResponseWriter, code int, body string) {
	w.WriteHeader(code)
	if body != "" {
		_, _ = w.Write([]byte(body))
	}
	logger.Warnf("response with code: %d %s", code, body)
}

type defaultAcceptor struct{}

func (a *defaultAcceptor) Accept(conn uim.Conn, timeout time.Duration) (string, error) {
	return ksuid.New().String(), nil
}
