// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/riteshRcH/go-edge-device-lib-core/peerstore (interfaces: Peerstore)

// Package pstoremanager_test is a generated GoMock package.
package pstoremanager_test

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	peer "github.com/libp2p/go-libp2p-core/peer"
	multiaddr "github.com/riteshRcH/go-multiaddr"
)

// MockPeerstore is a mock of Peerstore interface.
type MockPeerstore struct {
	ctrl     *gomock.Controller
	recorder *MockPeerstoreMockRecorder
}

// MockPeerstoreMockRecorder is the mock recorder for MockPeerstore.
type MockPeerstoreMockRecorder struct {
	mock *MockPeerstore
}

// NewMockPeerstore creates a new mock instance.
func NewMockPeerstore(ctrl *gomock.Controller) *MockPeerstore {
	mock := &MockPeerstore{ctrl: ctrl}
	mock.recorder = &MockPeerstoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeerstore) EXPECT() *MockPeerstoreMockRecorder {
	return m.recorder
}

// AddAddr mocks base method.
func (m *MockPeerstore) AddAddr(arg0 peer.ID, arg1 multiaddr.Multiaddr, arg2 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddAddr", arg0, arg1, arg2)
}

// AddAddr indicates an expected call of AddAddr.
func (mr *MockPeerstoreMockRecorder) AddAddr(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddr", reflect.TypeOf((*MockPeerstore)(nil).AddAddr), arg0, arg1, arg2)
}

// AddAddrs mocks base method.
func (m *MockPeerstore) AddAddrs(arg0 peer.ID, arg1 []multiaddr.Multiaddr, arg2 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddAddrs", arg0, arg1, arg2)
}

// AddAddrs indicates an expected call of AddAddrs.
func (mr *MockPeerstoreMockRecorder) AddAddrs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddrs", reflect.TypeOf((*MockPeerstore)(nil).AddAddrs), arg0, arg1, arg2)
}

// AddPrivKey mocks base method.
func (m *MockPeerstore) AddPrivKey(arg0 peer.ID, arg1 crypto.PrivKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPrivKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPrivKey indicates an expected call of AddPrivKey.
func (mr *MockPeerstoreMockRecorder) AddPrivKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPrivKey", reflect.TypeOf((*MockPeerstore)(nil).AddPrivKey), arg0, arg1)
}

// AddProtocols mocks base method.
func (m *MockPeerstore) AddProtocols(arg0 peer.ID, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddProtocols", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProtocols indicates an expected call of AddProtocols.
func (mr *MockPeerstoreMockRecorder) AddProtocols(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProtocols", reflect.TypeOf((*MockPeerstore)(nil).AddProtocols), varargs...)
}

// AddPubKey mocks base method.
func (m *MockPeerstore) AddPubKey(arg0 peer.ID, arg1 crypto.PubKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPubKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPubKey indicates an expected call of AddPubKey.
func (mr *MockPeerstoreMockRecorder) AddPubKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPubKey", reflect.TypeOf((*MockPeerstore)(nil).AddPubKey), arg0, arg1)
}

// AddrStream mocks base method.
func (m *MockPeerstore) AddrStream(arg0 context.Context, arg1 peer.ID) <-chan multiaddr.Multiaddr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrStream", arg0, arg1)
	ret0, _ := ret[0].(<-chan multiaddr.Multiaddr)
	return ret0
}

// AddrStream indicates an expected call of AddrStream.
func (mr *MockPeerstoreMockRecorder) AddrStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrStream", reflect.TypeOf((*MockPeerstore)(nil).AddrStream), arg0, arg1)
}

// Addrs mocks base method.
func (m *MockPeerstore) Addrs(arg0 peer.ID) []multiaddr.Multiaddr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addrs", arg0)
	ret0, _ := ret[0].([]multiaddr.Multiaddr)
	return ret0
}

// Addrs indicates an expected call of Addrs.
func (mr *MockPeerstoreMockRecorder) Addrs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addrs", reflect.TypeOf((*MockPeerstore)(nil).Addrs), arg0)
}

// ClearAddrs mocks base method.
func (m *MockPeerstore) ClearAddrs(arg0 peer.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearAddrs", arg0)
}

// ClearAddrs indicates an expected call of ClearAddrs.
func (mr *MockPeerstoreMockRecorder) ClearAddrs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearAddrs", reflect.TypeOf((*MockPeerstore)(nil).ClearAddrs), arg0)
}

// Close mocks base method.
func (m *MockPeerstore) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockPeerstoreMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPeerstore)(nil).Close))
}

// FirstSupportedProtocol mocks base method.
func (m *MockPeerstore) FirstSupportedProtocol(arg0 peer.ID, arg1 ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FirstSupportedProtocol", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FirstSupportedProtocol indicates an expected call of FirstSupportedProtocol.
func (mr *MockPeerstoreMockRecorder) FirstSupportedProtocol(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstSupportedProtocol", reflect.TypeOf((*MockPeerstore)(nil).FirstSupportedProtocol), varargs...)
}

// Get mocks base method.
func (m *MockPeerstore) Get(arg0 peer.ID, arg1 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPeerstoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPeerstore)(nil).Get), arg0, arg1)
}

// GetProtocols mocks base method.
func (m *MockPeerstore) GetProtocols(arg0 peer.ID) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProtocols", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProtocols indicates an expected call of GetProtocols.
func (mr *MockPeerstoreMockRecorder) GetProtocols(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProtocols", reflect.TypeOf((*MockPeerstore)(nil).GetProtocols), arg0)
}

// LatencyEWMA mocks base method.
func (m *MockPeerstore) LatencyEWMA(arg0 peer.ID) time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatencyEWMA", arg0)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// LatencyEWMA indicates an expected call of LatencyEWMA.
func (mr *MockPeerstoreMockRecorder) LatencyEWMA(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatencyEWMA", reflect.TypeOf((*MockPeerstore)(nil).LatencyEWMA), arg0)
}

// PeerInfo mocks base method.
func (m *MockPeerstore) PeerInfo(arg0 peer.ID) peer.AddrInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeerInfo", arg0)
	ret0, _ := ret[0].(peer.AddrInfo)
	return ret0
}

// PeerInfo indicates an expected call of PeerInfo.
func (mr *MockPeerstoreMockRecorder) PeerInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerInfo", reflect.TypeOf((*MockPeerstore)(nil).PeerInfo), arg0)
}

// Peers mocks base method.
func (m *MockPeerstore) Peers() peer.IDSlice {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peers")
	ret0, _ := ret[0].(peer.IDSlice)
	return ret0
}

// Peers indicates an expected call of Peers.
func (mr *MockPeerstoreMockRecorder) Peers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peers", reflect.TypeOf((*MockPeerstore)(nil).Peers))
}

// PeersWithAddrs mocks base method.
func (m *MockPeerstore) PeersWithAddrs() peer.IDSlice {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeersWithAddrs")
	ret0, _ := ret[0].(peer.IDSlice)
	return ret0
}

// PeersWithAddrs indicates an expected call of PeersWithAddrs.
func (mr *MockPeerstoreMockRecorder) PeersWithAddrs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeersWithAddrs", reflect.TypeOf((*MockPeerstore)(nil).PeersWithAddrs))
}

// PeersWithKeys mocks base method.
func (m *MockPeerstore) PeersWithKeys() peer.IDSlice {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeersWithKeys")
	ret0, _ := ret[0].(peer.IDSlice)
	return ret0
}

// PeersWithKeys indicates an expected call of PeersWithKeys.
func (mr *MockPeerstoreMockRecorder) PeersWithKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeersWithKeys", reflect.TypeOf((*MockPeerstore)(nil).PeersWithKeys))
}

// PrivKey mocks base method.
func (m *MockPeerstore) PrivKey(arg0 peer.ID) crypto.PrivKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivKey", arg0)
	ret0, _ := ret[0].(crypto.PrivKey)
	return ret0
}

// PrivKey indicates an expected call of PrivKey.
func (mr *MockPeerstoreMockRecorder) PrivKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivKey", reflect.TypeOf((*MockPeerstore)(nil).PrivKey), arg0)
}

// PubKey mocks base method.
func (m *MockPeerstore) PubKey(arg0 peer.ID) crypto.PubKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PubKey", arg0)
	ret0, _ := ret[0].(crypto.PubKey)
	return ret0
}

// PubKey indicates an expected call of PubKey.
func (mr *MockPeerstoreMockRecorder) PubKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubKey", reflect.TypeOf((*MockPeerstore)(nil).PubKey), arg0)
}

// Put mocks base method.
func (m *MockPeerstore) Put(arg0 peer.ID, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockPeerstoreMockRecorder) Put(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockPeerstore)(nil).Put), arg0, arg1, arg2)
}

// RecordLatency mocks base method.
func (m *MockPeerstore) RecordLatency(arg0 peer.ID, arg1 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RecordLatency", arg0, arg1)
}

// RecordLatency indicates an expected call of RecordLatency.
func (mr *MockPeerstoreMockRecorder) RecordLatency(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordLatency", reflect.TypeOf((*MockPeerstore)(nil).RecordLatency), arg0, arg1)
}

// RemovePeer mocks base method.
func (m *MockPeerstore) RemovePeer(arg0 peer.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemovePeer", arg0)
}

// RemovePeer indicates an expected call of RemovePeer.
func (mr *MockPeerstoreMockRecorder) RemovePeer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePeer", reflect.TypeOf((*MockPeerstore)(nil).RemovePeer), arg0)
}

// RemoveProtocols mocks base method.
func (m *MockPeerstore) RemoveProtocols(arg0 peer.ID, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveProtocols", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProtocols indicates an expected call of RemoveProtocols.
func (mr *MockPeerstoreMockRecorder) RemoveProtocols(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProtocols", reflect.TypeOf((*MockPeerstore)(nil).RemoveProtocols), varargs...)
}

// SetAddr mocks base method.
func (m *MockPeerstore) SetAddr(arg0 peer.ID, arg1 multiaddr.Multiaddr, arg2 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAddr", arg0, arg1, arg2)
}

// SetAddr indicates an expected call of SetAddr.
func (mr *MockPeerstoreMockRecorder) SetAddr(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAddr", reflect.TypeOf((*MockPeerstore)(nil).SetAddr), arg0, arg1, arg2)
}

// SetAddrs mocks base method.
func (m *MockPeerstore) SetAddrs(arg0 peer.ID, arg1 []multiaddr.Multiaddr, arg2 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAddrs", arg0, arg1, arg2)
}

// SetAddrs indicates an expected call of SetAddrs.
func (mr *MockPeerstoreMockRecorder) SetAddrs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAddrs", reflect.TypeOf((*MockPeerstore)(nil).SetAddrs), arg0, arg1, arg2)
}

// SetProtocols mocks base method.
func (m *MockPeerstore) SetProtocols(arg0 peer.ID, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetProtocols", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProtocols indicates an expected call of SetProtocols.
func (mr *MockPeerstoreMockRecorder) SetProtocols(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProtocols", reflect.TypeOf((*MockPeerstore)(nil).SetProtocols), varargs...)
}

// Start mocks base method.
func (m *MockPeerstore) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockPeerstoreMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockPeerstore)(nil).Start))
}

// SupportsProtocols mocks base method.
func (m *MockPeerstore) SupportsProtocols(arg0 peer.ID, arg1 ...string) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SupportsProtocols", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SupportsProtocols indicates an expected call of SupportsProtocols.
func (mr *MockPeerstoreMockRecorder) SupportsProtocols(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportsProtocols", reflect.TypeOf((*MockPeerstore)(nil).SupportsProtocols), varargs...)
}

// UpdateAddrs mocks base method.
func (m *MockPeerstore) UpdateAddrs(arg0 peer.ID, arg1, arg2 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateAddrs", arg0, arg1, arg2)
}

// UpdateAddrs indicates an expected call of UpdateAddrs.
func (mr *MockPeerstoreMockRecorder) UpdateAddrs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddrs", reflect.TypeOf((*MockPeerstore)(nil).UpdateAddrs), arg0, arg1, arg2)
}
