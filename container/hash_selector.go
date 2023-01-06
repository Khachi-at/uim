package container

import (
	"github.com/uim"
	"github.com/uim/wire/pkt"
)

type HashSelector struct{}

func (s *HashSelector) Lookup(header *pkt.Header, srvs []uim.Service) string {
	ll := len(srvs)
	code := HashCode(header.ChannelId)
	return srvs[code%ll].ServiceID()
}
