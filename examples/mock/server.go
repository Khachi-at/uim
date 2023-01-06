package mock

import (
	"errors"
	"time"

	"github.com/uim"
	"github.com/uim/logger"
	"github.com/uim/naming"
	"github.com/uim/tcp"
	"github.com/uim/websocket"
)

type ServerDemo struct{}

func (s *ServerDemo) Start(id, protocol, addr string) {
	var srv uim.Server
	service := &naming.DefaultService{
		Id:       id,
		Protocol: protocol,
	}
	// Ignore server's internal logic.
	if protocol == "ws" {
		srv = websocket.NewServer(addr, service)
	} else if protocol == "tcp" {
		srv = tcp.NewServer(addr, service)
	}

	handler := &ServerHandler{}

	srv.SetReadWait(time.Minute)
	srv.SetAcceptor(handler)
	srv.SetMessageListener(handler)
	srv.SetStateListener(handler)

	err := srv.Start()
	if err != nil {
		panic(err)
	}
}

// ServerHandler.
type ServerHandler struct{}

// Accept this connection.
func (h *ServerHandler) Accept(conn uim.Conn, timeout time.Duration) (string, error) {
	// 1. Read auth frame.
	frame, err := conn.ReadFrame()
	if err != nil {
		return "", nil
	}
	logger.Info("recv ", frame.GetOpCode())
	// 2. Parse userId.
	userID := string(frame.GetPayload())
	// 3. Auth.
	if userID == "" {
		return "", errors.New("user id is invalid")
	}
	return userID, nil
}

// Receive default listener.
func (h *ServerHandler) Receive(ag uim.Agent, payload []byte) {
	logger.Info("Receive")
	ack := string(payload) + " from server "
	_ = ag.Push([]byte(ack))
}

// Disconnect default listener.
func (h *ServerHandler) Disconnect(id string) error {
	logger.Warnf("disconnect %s", id)
	return nil
}
