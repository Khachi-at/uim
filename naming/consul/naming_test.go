package consul

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/uim"
	"github.com/uim/naming"
)

func Test_Naming(t *testing.T) {
	ns, err := NewNaming("localhost:8500")
	assert.Nil(t, err)

	_ = ns.Deregister("test_1")
	_ = ns.Deregister("test_2")

	serviceName := "for_test"
	// 1. Register test_1.
	err = ns.Register(&naming.DefaultService{
		Id:        "test_1",
		Name:      serviceName,
		Namespace: "",
		Address:   "localhost",
		Port:      8000,
		Protocol:  "ws",
		Tags:      []string{"tab1", "gate"},
	})
	assert.Nil(t, err)

	// 2. Service discover.
	servs, err := ns.Find(serviceName)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(servs))
	t.Log(servs)

	wg := sync.WaitGroup{}
	wg.Add(1)

	// 3. Subscribe service change.
	_ = ns.Subscribe(serviceName, func(services []uim.ServiceRegistration) {
		t.Log(len(services))

		assert.Equal(t, 2, len(services))
		assert.Equal(t, "test_2", services[1].ServiceID())
		wg.Done()
	})
	time.Sleep(time.Second)

	// 4. register test_2
	err = ns.Register(&naming.DefaultService{
		Id:        "test_2",
		Name:      serviceName,
		Namespace: "",
		Address:   "localhost",
		Port:      8001,
		Protocol:  "ws",
		Tags:      []string{"tab2", "gate"},
	})
	assert.Nil(t, err)

	wg.Wait()

	// 5. Service discover.
	servs, _ = ns.Find(serviceName, "gate")
	assert.Equal(t, 2, len(servs))

	// 6. Verify Tag query.
	servs, _ = ns.Find(serviceName, "tab2")
	assert.Equal(t, 1, len(servs))
	assert.Equal(t, "test_2", servs[0].ServiceID())

	// 7. Deregister test_2
	err = ns.Deregister("test_2")
	assert.Nil(t, err)

	// 8. Service discovery
	servs, err = ns.Find(serviceName)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(servs))
	assert.Equal(t, "test_1", servs[0].ServiceID())

	// 9. Register test_1.
	err = ns.Deregister("test1")
	assert.Nil(t, err)
}
