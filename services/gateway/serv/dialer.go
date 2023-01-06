package serv

import (
	"net"

	"github.com/uim"
	"github.com/uim/logger"
	"github.com/uim/tcp"
	"github.com/uim/wire/pkt"
	"google.golang.org/protobuf/proto"
)

type TcpDialer struct {
	ServiceId string
}

func NewDialer(serviceId string) uim.Dialer {
	return &TcpDialer{
		ServiceId: serviceId,
	}
}

func (d *TcpDialer) DialAndHandshake(ctx uim.DialerContext) (net.Conn, error) {
	// 1. Dialer and create connection.
	conn, err := net.DialTimeout("tcp", ctx.Address, ctx.Timeout)
	if err != nil {
		return nil, err
	}
	req := &pkt.InnerHandshakeReq{
		ServiceId: d.ServiceId,
	}
	logger.Infof("send req %v", req)
	// 2. Send my serviceId to peer.
	bts, _ := proto.Marshal(req)
	err = tcp.WriteFrame(conn, uim.OpBinary, bts)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
