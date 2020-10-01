// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.05-release
// source: /usr/share/vpp/api/plugins/dhcp6_pd_client_cp.api.json

// Package dhcp6_pd_client_cp contains generated bindings for API file dhcp6_pd_client_cp.api.
//
// Contents:
//   4 messages
//
package dhcp6_pd_client_cp

import (
	api "git.fd.io/govpp.git/api"
	interface_types "git.fd.io/govpp.git/binapi/interface_types"
	ip_types "git.fd.io/govpp.git/binapi/ip_types"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "dhcp6_pd_client_cp"
	APIVersion = "2.0.0"
	VersionCrc = 0x96f41948
)

// DHCP6PdClientEnableDisable defines message 'dhcp6_pd_client_enable_disable'.
type DHCP6PdClientEnableDisable struct {
	SwIfIndex   interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	PrefixGroup string                         `binapi:"string[64],name=prefix_group" json:"prefix_group,omitempty"`
	Enable      bool                           `binapi:"bool,name=enable" json:"enable,omitempty"`
}

func (m *DHCP6PdClientEnableDisable) Reset()               { *m = DHCP6PdClientEnableDisable{} }
func (*DHCP6PdClientEnableDisable) GetMessageName() string { return "dhcp6_pd_client_enable_disable" }
func (*DHCP6PdClientEnableDisable) GetCrcString() string   { return "a75a0772" }
func (*DHCP6PdClientEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DHCP6PdClientEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4  // m.SwIfIndex
	size += 64 // m.PrefixGroup
	size += 1  // m.Enable
	return size
}
func (m *DHCP6PdClientEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeString(m.PrefixGroup, 64)
	buf.EncodeBool(m.Enable)
	return buf.Bytes(), nil
}
func (m *DHCP6PdClientEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.PrefixGroup = buf.DecodeString(64)
	m.Enable = buf.DecodeBool()
	return nil
}

// DHCP6PdClientEnableDisableReply defines message 'dhcp6_pd_client_enable_disable_reply'.
type DHCP6PdClientEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *DHCP6PdClientEnableDisableReply) Reset() { *m = DHCP6PdClientEnableDisableReply{} }
func (*DHCP6PdClientEnableDisableReply) GetMessageName() string {
	return "dhcp6_pd_client_enable_disable_reply"
}
func (*DHCP6PdClientEnableDisableReply) GetCrcString() string { return "e8d4e804" }
func (*DHCP6PdClientEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DHCP6PdClientEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *DHCP6PdClientEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *DHCP6PdClientEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// IP6AddDelAddressUsingPrefix defines message 'ip6_add_del_address_using_prefix'.
type IP6AddDelAddressUsingPrefix struct {
	SwIfIndex         interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	PrefixGroup       string                         `binapi:"string[64],name=prefix_group" json:"prefix_group,omitempty"`
	AddressWithPrefix ip_types.IP6AddressWithPrefix  `binapi:"ip6_address_with_prefix,name=address_with_prefix" json:"address_with_prefix,omitempty"`
	IsAdd             bool                           `binapi:"bool,name=is_add" json:"is_add,omitempty"`
}

func (m *IP6AddDelAddressUsingPrefix) Reset() { *m = IP6AddDelAddressUsingPrefix{} }
func (*IP6AddDelAddressUsingPrefix) GetMessageName() string {
	return "ip6_add_del_address_using_prefix"
}
func (*IP6AddDelAddressUsingPrefix) GetCrcString() string { return "9b3d11e0" }
func (*IP6AddDelAddressUsingPrefix) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IP6AddDelAddressUsingPrefix) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.SwIfIndex
	size += 64     // m.PrefixGroup
	size += 1 * 16 // m.AddressWithPrefix.Address
	size += 1      // m.AddressWithPrefix.Len
	size += 1      // m.IsAdd
	return size
}
func (m *IP6AddDelAddressUsingPrefix) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeString(m.PrefixGroup, 64)
	buf.EncodeBytes(m.AddressWithPrefix.Address[:], 16)
	buf.EncodeUint8(m.AddressWithPrefix.Len)
	buf.EncodeBool(m.IsAdd)
	return buf.Bytes(), nil
}
func (m *IP6AddDelAddressUsingPrefix) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.PrefixGroup = buf.DecodeString(64)
	copy(m.AddressWithPrefix.Address[:], buf.DecodeBytes(16))
	m.AddressWithPrefix.Len = buf.DecodeUint8()
	m.IsAdd = buf.DecodeBool()
	return nil
}

// IP6AddDelAddressUsingPrefixReply defines message 'ip6_add_del_address_using_prefix_reply'.
type IP6AddDelAddressUsingPrefixReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IP6AddDelAddressUsingPrefixReply) Reset() { *m = IP6AddDelAddressUsingPrefixReply{} }
func (*IP6AddDelAddressUsingPrefixReply) GetMessageName() string {
	return "ip6_add_del_address_using_prefix_reply"
}
func (*IP6AddDelAddressUsingPrefixReply) GetCrcString() string { return "e8d4e804" }
func (*IP6AddDelAddressUsingPrefixReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IP6AddDelAddressUsingPrefixReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IP6AddDelAddressUsingPrefixReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IP6AddDelAddressUsingPrefixReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_dhcp6_pd_client_cp_binapi_init() }
func file_dhcp6_pd_client_cp_binapi_init() {
	api.RegisterMessage((*DHCP6PdClientEnableDisable)(nil), "dhcp6_pd_client_enable_disable_a75a0772")
	api.RegisterMessage((*DHCP6PdClientEnableDisableReply)(nil), "dhcp6_pd_client_enable_disable_reply_e8d4e804")
	api.RegisterMessage((*IP6AddDelAddressUsingPrefix)(nil), "ip6_add_del_address_using_prefix_9b3d11e0")
	api.RegisterMessage((*IP6AddDelAddressUsingPrefixReply)(nil), "ip6_add_del_address_using_prefix_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*DHCP6PdClientEnableDisable)(nil),
		(*DHCP6PdClientEnableDisableReply)(nil),
		(*IP6AddDelAddressUsingPrefix)(nil),
		(*IP6AddDelAddressUsingPrefixReply)(nil),
	}
}
