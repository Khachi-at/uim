package handler

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/uim"
	"github.com/uim/wire"
	"github.com/uim/wire/pkt"
)

func TestLoginHandler_DoSysLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dispatcher := uim.NewMockDispatcher(ctrl)
	cache := uim.NewMockSessionStorage(ctrl)
	session := &pkt.Session{
		ChannelId: "channel1",
		Account:   "test1",
		GateId:    "gateway1",
	}
	// resp.
	dispatcher.EXPECT().Push(session.GateId, []string{"channel1"}, gomock.Any()).Times(1)
	// kickout notify
	dispatcher.EXPECT().Push(session.GateId, []string{"channel2"}, gomock.Any()).Times(1)

	cache.EXPECT().GetLocation(session.Account, "").DoAndReturn(func(account, device string) (*uim.Location, error) {
		return &uim.Location{
			ChannelId: "channel2",
			GateId:    "gateway1",
		}, nil
	})
	cache.EXPECT().Add(gomock.Any()).Times(1).DoAndReturn(func(add *pkt.Session) error {
		assert.Equal(t, session.ChannelId, add.ChannelId)
		assert.Equal(t, session.Account, add.Account)
		return nil
	})

	loginreq := pkt.New(wire.CommandLoginSignIn).WriteBody(session)

	r := uim.NewRouter()
	// login.
	loginHandler := NewLoginHandler()
	r.Handle(wire.CommandLoginSignIn, loginHandler.DoSysLogin)
	err := r.Serve(loginreq, dispatcher, cache, session)
	assert.Nil(t, err)
}
