package unittest

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/uim"
	"github.com/uim/examples/dialer"
	"github.com/uim/websocket"
)

func login(account string) (uim.Client, error) {
	cli := websocket.NewClient(account, "unittest", websocket.ClientOptions{})

	cli.SetDialer(&dialer.ClientDialer{})
	if err := cli.Connect("ws://localhost:8005"); err != nil {
		return nil, err
	}
	return cli, nil
}

func Test_login(t *testing.T) {
	cli, err := login("test1")
	assert.Nil(t, err)
	time.Sleep(time.Second * 3)
	cli.Close()
}
