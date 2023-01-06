package serv

import (
	"bytes"
	"strings"
	"time"

	"github.com/uim"
	"github.com/uim/container"
	"github.com/uim/logger"
	"github.com/uim/wire"
	"github.com/uim/wire/pkt"
	"google.golang.org/protobuf/proto"
)

var log = logger.WithFields(logger.Fields{
	"service": wire.SNChat,
	"pkg":     "serv",
})

type ServHandler struct {
	r          *uim.Router
	cache      uim.SessionStorage
	dispatcher *ServerDispatcher
}

func NewServHandler(r *uim.Router, cache uim.SessionStorage) *ServHandler {
	return &ServHandler{
		r:          r,
		dispatcher: &ServerDispatcher{},
		cache:      cache,
	}
}

// Accept this connection.
func (h *ServHandler) Accept(conn uim.Conn, timeout time.Duration) (string, error) {
	log.Infoln("enter")

	_ = conn.SetReadDeadline(time.Now().Add(timeout))
	frame, err := conn.ReadFrame()
	if err != nil {
		return "", err
	}

	var req pkt.InnerHandshakeReq
	_ = proto.Unmarshal(frame.GetPayload(), &req)
	log.Info("Accept -- ", req.ServiceId)

	return req.ServiceId, nil
}

func (h *ServHandler) Receive(ag uim.Agent, payload []byte) {
	buf := bytes.NewBuffer(payload)
	packet, err := pkt.MustReadLogicPkt(buf)
	if err != nil {
		log.Error(err)
		return
	}
	var session *pkt.Session
	if packet.Command == wire.CommandLoginSignIn {
		server, _ := packet.GetMeta(wire.MetaDestServer)
		session = &pkt.Session{
			ChannelId: packet.ChannelId,
			GateId:    server.(string),
			Tags:      []string{"AutoGenerated"},
		}
	} else {
		// TODO
		session, err = h.cache.Get(packet.ChannelId)
		if err == uim.ErrSessionNil {
			_ = RespErr(ag, packet, pkt.Status_SessionNotFound)
			return
		} else if err != nil {
			_ = RespErr(ag, packet, pkt.Status_SystemException)
			return
		}
	}
	log.Debug("recv a message from %s %s", session, &packet.Header)
	err = h.r.Serve(packet, h.dispatcher, h.cache, session)
	if err != nil {
		log.Warn(err)
	}
}

// Disconnect default listener.
func (h *ServHandler) Disconnect(id string) error {
	logger.Warnf("close event of %s", id)
	return nil
}

func RespErr(ag uim.Agent, p *pkt.LogicPkt, status pkt.Status) error {
	packet := pkt.NewFrom(&p.Header)
	packet.Status = status
	packet.Flag = pkt.Flag_Response
	return ag.Push(pkt.Marshal(packet))
}

type ServerDispatcher struct{}

func (d *ServerDispatcher) Push(gateway string, channels []string, p *pkt.LogicPkt) error {
	p.AddStringMeta(wire.MetaDestChannels, strings.Join(channels, ","))
	return container.Push(gateway, p)
}
