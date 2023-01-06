package pkt

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uim/wire"
)

func TestReadPkt(t *testing.T) {
	seq := wire.Seq.Next()
	packet := New("auth.login.aa", WithSeq(seq), WithStatus(Status_Success))
	assert.Equal(t, "auth", packet.ServiceName())

	packet = New("auth.login", WithSeq(seq), WithStatus(Status_Success))
	packet.WriteBody(&LoginReq{
		Token: "test token",
	})
	packet.AddMeta(&Meta{
		Key:   "test",
		Value: "test",
	}, &Meta{
		Key:   wire.MetaDestServer,
		Value: "test",
	}, &Meta{
		Key:   wire.MetaDestChannels,
		Value: "test1,test2",
	})
	buf := new(bytes.Buffer)
	_ = packet.Encode(buf)
	t.Log(buf.Bytes())

	got, err := Read(buf)
	p := got.(*LogicPkt)
	assert.Nil(t, err)
	assert.Equal(t, "auth.login", p.Command)
	assert.Equal(t, seq, p.Sequence)
	assert.Equal(t, Status_Success, p.Status)
	assert.Equal(t, 3, len(packet.Meta))
	packet.DelMeta(wire.MetaDestServer)
	assert.Equal(t, 2, len(packet.Meta))
	assert.Equal(t, wire.MetaDestChannels, packet.Meta[1].Key)
	packet.DelMeta(wire.MetaDestChannels)
	assert.Equal(t, 1, len(packet.Meta))
}

func Test_Encode(t *testing.T) {
	var pkt = struct {
		Source   uint32
		Sequence uint64
		Data     []byte
	}{
		Source:   0x010201,
		Sequence: 2<<60 + 3,
		Data:     []byte("hello world"),
	}
	endian := binary.BigEndian

	buf := make([]byte, 1024)
	i := 0
	endian.PutUint32(buf[i:i+4], pkt.Source)
	i += 4
	endian.PutUint64(buf[i:i+8], pkt.Sequence)
	i += 8
	dataLen := len(pkt.Data)
	endian.PutUint32(buf[i:i+4], uint32(dataLen))
	i += 4
	copy(buf[i:i+dataLen], pkt.Data)
	i += dataLen
	t.Log(buf[0:i])
	t.Log("length", i)
}

func Test_Decode(t *testing.T) {
	var pkt struct {
		Source   uint32
		Sequence uint64
		Data     []byte
	}
	recv := []byte{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 11, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	endian := binary.BigEndian
	i := 0
	pkt.Source = endian.Uint32(recv[i : i+4])
	i += 4
	pkt.Sequence = endian.Uint64(recv[i : i+8])
	i += 8
	dataLen := endian.Uint32(recv[i : i+4])
	i += 4
	pkt.Data = make([]byte, dataLen)
	copy(pkt.Data, recv[i:i+int(dataLen)])
	t.Logf("Src:%d Seq:%d Data:%s", pkt.Source, pkt.Sequence, pkt.Data)

}

func Benchmark_Encode(t *testing.B) {
	var pkt = struct {
		Source   uint32
		Sequence uint64
		Data     []byte
	}{
		Source:   10000000,
		Sequence: 2<<60 + 3,
		Data:     []byte("hello world"),
	}
	for i := 0; i < t.N; i++ {
		endian := binary.BigEndian
		buf := make([]byte, 30)
		i := 0
		endian.PutUint32(buf[i:i+4], pkt.Source)
		i += 4
		endian.PutUint64(buf[i:i+8], pkt.Sequence)
		i += 8
		dataLen := len(pkt.Data)
		endian.PutUint32(buf[i:i+4], uint32(dataLen))
		i += 4
		copy(buf[i:i+dataLen], pkt.Data)
		i += dataLen
	}
}
