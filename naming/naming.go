package naming

import (
	"errors"

	"github.com/uim"
)

// errors
var (
	ErrNotFound = errors.New("service not found")
)

// Naming define methods of the naming service.
type Naming interface {
	Find(serviceName string, tags ...string) ([]uim.ServiceRegistration, error)
	Subscribe(serviceName string, callback func(services []uim.ServiceRegistration)) error
	Unsubscribe(serviceName string) error
	Register(service uim.ServiceRegistration) error
	Deregister(serviceID string) error
}
