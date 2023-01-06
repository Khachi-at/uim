package mock

import (
	"context"
	"net"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/uim"
	"github.com/uim/logger"
	"github.com/uim/tcp"
	"github.com/uim/websocket"
)

type ClientDemo struct{}

func (c *ClientDemo) Start(userID, protocol, addr string) {
	var cli uim.Client

	// 1. Init client.
	if protocol == "ws" {
		cli = websocket.NewClient(userID, "client", websocket.ClientOptions{})
		// set dialer
		cli.SetDialer(&WebSocketDialer{})
	} else if protocol == "tcp" {
		cli = tcp.NewClient("test1", "client", tcp.ClientOptions{})
		cli.SetDialer(&TCPDialer{})
	}

	// 2. Create connection.
	err := cli.Connect(addr)
	if err != nil {
		logger.Error(err)
	}
	count := 10
	go func() {
		// 3. Send message and return.
		for i := 0; i < count; i++ {
			err := cli.Send([]byte("hello"))
			if err != nil {
				logger.Error(err)
				return
			}
			time.Sleep(time.Second)
		}
	}()

	// 4. Receive message.
	recv := 0
	for {
		frame, err := cli.Read()
		if err != nil {
			logger.Info(err)
			break
		}
		if frame.GetOpCode() != uim.OpBinary {
			continue
		}
		recv++
		logger.Warnf("%s receive message [%s]", cli.ServiceID(), frame.GetPayload())
		if recv == count { // Complete message Receive.
			break
		}
	}
	// return
	cli.Close()
}

type ClientHandler struct{}

// Receive default listener.
func (h *ClientHandler) Receive(ag uim.Agent, payload []byte) {
	logger.Warnf("%s receive message [%s]", ag.ID(), string(payload))
}

func (h *ClientHandler) Disconnect(id string) error {
	logger.Warnf("disconnect %s", id)
	return nil
}

// WebsocketDialer.
type WebSocketDialer struct{}

// DialAndHandshake.
func (d *WebSocketDialer) DialAndHandshake(ctx uim.DialerContext) (net.Conn, error) {
	// 1. Call ws.Dial().
	conn, _, _, err := ws.Dial(context.TODO(), ctx.Address)
	if err != nil {
		return nil, err
	}
	// 2. Send auth info.
	err = wsutil.WriteClientBinary(conn, []byte(ctx.Id))
	if err != nil {
		return nil, err
	}

	// 3. Return connection.
	return conn, nil
}

type TCPDialer struct{}

func (d *TCPDialer) DialAndHandshake(ctx uim.DialerContext) (net.Conn, error) {
	logger.Info("start dial: ", ctx.Address)
	conn, err := net.DialTimeout("tcp", ctx.Address, ctx.Timeout)
	if err != nil {
		return nil, err
	}
	err = tcp.WriteFrame(conn, uim.OpBinary, []byte(ctx.Id))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
