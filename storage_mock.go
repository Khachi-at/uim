package uim

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/uim/wire/pkt"
)

// MockSessionStorage is a mock of SessionStorage interface.
type MockSessionStorage struct {
	ctrl     *gomock.Controller
	recorder *MockSessionStorageMockRecorder
}

type MockSessionStorageMockRecorder struct {
	mock *MockSessionStorage
}

func NewMockSessionStorage(ctrl *gomock.Controller) *MockSessionStorage {
	mock := &MockSessionStorage{ctrl: ctrl}
	mock.recorder = &MockSessionStorageMockRecorder{mock: mock}
	return mock
}

func (m *MockSessionStorage) EXPECT() *MockSessionStorageMockRecorder {
	return m.recorder
}

func (m *MockSessionStorage) Add(session *pkt.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", session)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockSessionStorageMockRecorder) Add(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockSessionStorage)(nil).Add), session)
}

func (m *MockSessionStorage) Delete(account, channelId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", account, channelId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockSessionStorageMockRecorder) Delete(account, channelId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSessionStorage)(nil).Delete), account, channelId)
}

// Get mocks base method.
func (m *MockSessionStorage) Get(channelId string) (*pkt.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", channelId)
	ret0, _ := ret[0].(*pkt.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSessionStorageMockRecorder) Get(channelId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSessionStorage)(nil).Get), channelId)
}

// GetLocation mocks base method.
func (m *MockSessionStorage) GetLocation(account, device string) (*Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocation", account, device)
	ret0, _ := ret[0].(*Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocation indicates an expected call of GetLocation.
func (mr *MockSessionStorageMockRecorder) GetLocation(account, device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocation", reflect.TypeOf((*MockSessionStorage)(nil).GetLocation), account, device)
}

// GetLocations mocks base method.
func (m *MockSessionStorage) GetLocations(account ...string) ([]*Location, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range account {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLocations", varargs...)
	ret0, _ := ret[0].([]*Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocations indicates an expected call of GetLocations.
func (mr *MockSessionStorageMockRecorder) GetLocations(account ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocations", reflect.TypeOf((*MockSessionStorage)(nil).GetLocations), account...)
}
