package container

import (
	"sync"

	"github.com/uim"
	"github.com/uim/logger"
)

type ClientMap interface {
	Add(client uim.Client)
	Remove(id string)
	Get(id string) (client uim.Client, ok bool)
	Services(kvs ...string) []uim.Service
}

type ClientsImpl struct {
	clients *sync.Map
}

func NewClients(num int) ClientMap {
	return &ClientsImpl{
		clients: new(sync.Map),
	}
}

// Add channel.
func (ch *ClientsImpl) Add(client uim.Client) {
	if client.ServiceID() == "" {
		logger.WithFields(logger.Fields{
			"module": "ClientsImpl",
		}).Error("client id is required")
	}
	ch.clients.Store(client.ServiceID(), client)
}

// Remove Channel.
func (ch *ClientsImpl) Remove(id string) {
	ch.clients.Delete(id)
}

func (ch *ClientsImpl) Get(id string) (uim.Client, bool) {
	if id == "" {
		logger.WithFields(logger.Fields{
			"module": "ClientsImpl",
		}).Error("client id is required")
	}

	val, ok := ch.clients.Load(id)
	if !ok {
		return nil, false
	}
	return val.(uim.Client), true
}

func (ch *ClientsImpl) Services(kvs ...string) []uim.Service {
	kvLen := len(kvs)
	if kvLen != 0 && kvLen != 2 {
		return nil
	}
	arr := make([]uim.Service, 0)
	ch.clients.Range(func(key, val any) bool {
		ser := val.(uim.Service)
		if kvLen > 0 && ser.GetMeta()[kvs[0]] != kvs[1] {
			return true
		}
		arr = append(arr, ser)
		return true
	})
	return arr
}
