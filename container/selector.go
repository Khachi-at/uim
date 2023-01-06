package container

import (
	"hash/crc32"

	"github.com/uim"
	"github.com/uim/wire/pkt"
)

func HashCode(key string) int {
	hash32 := crc32.NewIEEE()
	hash32.Write([]byte(key))
	return int(hash32.Sum32())
}

type Selector interface {
	Lookup(*pkt.Header, []uim.Service) string
}
