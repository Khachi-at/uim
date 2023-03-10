package pkt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uim/wire"
)

func TestMarshal(t *testing.T) {
	bp := &BasicPkt{
		Code: CodePing,
	}

	bts := Marshal(bp)
	t.Log(bts)

	assert.Equal(t, wire.MagicBasicPkt[1], bts[1])
	assert.Equal(t, wire.MagicBasicPkt[2], bts[2])

	lp := New("login.signin")
	bts2 := Marshal(lp)
	t.Log(bts2)

	assert.Equal(t, wire.MagicLogicPkt[1], bts2[1])
	assert.Equal(t, wire.MagicLogicPkt[2], bts2[2])
}
