package dialer

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/uim"
	"github.com/uim/logger"
	"github.com/uim/wire"
	"github.com/uim/wire/pkt"
	"github.com/uim/wire/token"
)

type ClientDialer struct{}

func (d *ClientDialer) DialAndHandshake(ctx uim.DialerContext) (net.Conn, error) {
	logger.Info("DialAndHandshake called")
	// 1. Dialer.
	conn, _, _, err := ws.Dial(context.TODO(), ctx.Address)
	if err != nil {
		return nil, err
	}
	// 2. Generate token.
	tk, err := token.Generate(token.DefaultSecret, &token.Token{
		Account: ctx.Id,
		App:     "uim",
		Exp:     time.Now().AddDate(0, 0, 1).Unix(),
	})
	if err != nil {
		return nil, err
	}
	// 3. Send CommandLoginSignIn message.
	loginreq := pkt.New(wire.CommandLoginSignIn).WriteBody(&pkt.LoginReq{
		Token: tk,
	})
	err = wsutil.WriteClientBinary(conn, pkt.Marshal(loginreq))
	if err != nil {
		return nil, err
	}

	// wait resp.
	logger.Info("waiting for login response")
	_ = conn.SetReadDeadline(time.Now().Add(ctx.Timeout))
	frame, err := ws.ReadFrame(conn)
	if err != nil {
		return nil, err
	}
	ack, err := pkt.MustReadLogicPkt(bytes.NewBuffer(frame.Payload))
	if err != nil {
		return nil, err
	}
	// 4. Check login success.
	if ack.Status != pkt.Status_Success {
		return nil, fmt.Errorf("login failed: %v", &ack.Header)
	}
	var resp = new(pkt.LoginResp)
	_ = ack.ReadBody(resp)

	logger.Info("logined", resp.GetChannelId())
	return conn, nil
}
