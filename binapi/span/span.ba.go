// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0
//  VPP:              20.05.1-release
// source: /usr/share/vpp/api/core/span.api.json

// Package span contains generated bindings for API file span.api.
//
// Contents:
//   1 enum
//   4 messages
//
package span

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	interface_types "git.fd.io/govpp.git/binapi/interface_types"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "span"
	APIVersion = "2.0.0"
	VersionCrc = 0x9f4dec7c
)

// SpanState defines enum 'span_state'.
type SpanState uint32

const (
	SPAN_STATE_API_DISABLED SpanState = 0
	SPAN_STATE_API_RX       SpanState = 1
	SPAN_STATE_API_TX       SpanState = 2
	SPAN_STATE_API_RX_TX    SpanState = 3
)

var (
	SpanState_name = map[uint32]string{
		0: "SPAN_STATE_API_DISABLED",
		1: "SPAN_STATE_API_RX",
		2: "SPAN_STATE_API_TX",
		3: "SPAN_STATE_API_RX_TX",
	}
	SpanState_value = map[string]uint32{
		"SPAN_STATE_API_DISABLED": 0,
		"SPAN_STATE_API_RX":       1,
		"SPAN_STATE_API_TX":       2,
		"SPAN_STATE_API_RX_TX":    3,
	}
)

func (x SpanState) String() string {
	s, ok := SpanState_name[uint32(x)]
	if ok {
		return s
	}
	return "SpanState(" + strconv.Itoa(int(x)) + ")"
}

// SwInterfaceSpanDetails defines message 'sw_interface_span_details'.
type SwInterfaceSpanDetails struct {
	SwIfIndexFrom interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index_from" json:"sw_if_index_from,omitempty"`
	SwIfIndexTo   interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index_to" json:"sw_if_index_to,omitempty"`
	State         SpanState                      `binapi:"span_state,name=state" json:"state,omitempty"`
	IsL2          bool                           `binapi:"bool,name=is_l2" json:"is_l2,omitempty"`
}

func (m *SwInterfaceSpanDetails) Reset()               { *m = SwInterfaceSpanDetails{} }
func (*SwInterfaceSpanDetails) GetMessageName() string { return "sw_interface_span_details" }
func (*SwInterfaceSpanDetails) GetCrcString() string   { return "055643fc" }
func (*SwInterfaceSpanDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceSpanDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndexFrom
	size += 4 // m.SwIfIndexTo
	size += 4 // m.State
	size += 1 // m.IsL2
	return size
}
func (m *SwInterfaceSpanDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndexFrom))
	buf.EncodeUint32(uint32(m.SwIfIndexTo))
	buf.EncodeUint32(uint32(m.State))
	buf.EncodeBool(m.IsL2)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSpanDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndexFrom = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.SwIfIndexTo = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.State = SpanState(buf.DecodeUint32())
	m.IsL2 = buf.DecodeBool()
	return nil
}

// SwInterfaceSpanDump defines message 'sw_interface_span_dump'.
type SwInterfaceSpanDump struct {
	IsL2 bool `binapi:"bool,name=is_l2" json:"is_l2,omitempty"`
}

func (m *SwInterfaceSpanDump) Reset()               { *m = SwInterfaceSpanDump{} }
func (*SwInterfaceSpanDump) GetMessageName() string { return "sw_interface_span_dump" }
func (*SwInterfaceSpanDump) GetCrcString() string   { return "d6cf0c3d" }
func (*SwInterfaceSpanDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceSpanDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsL2
	return size
}
func (m *SwInterfaceSpanDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsL2)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSpanDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsL2 = buf.DecodeBool()
	return nil
}

// SwInterfaceSpanEnableDisable defines message 'sw_interface_span_enable_disable'.
type SwInterfaceSpanEnableDisable struct {
	SwIfIndexFrom interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index_from" json:"sw_if_index_from,omitempty"`
	SwIfIndexTo   interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index_to" json:"sw_if_index_to,omitempty"`
	State         SpanState                      `binapi:"span_state,name=state" json:"state,omitempty"`
	IsL2          bool                           `binapi:"bool,name=is_l2" json:"is_l2,omitempty"`
}

func (m *SwInterfaceSpanEnableDisable) Reset() { *m = SwInterfaceSpanEnableDisable{} }
func (*SwInterfaceSpanEnableDisable) GetMessageName() string {
	return "sw_interface_span_enable_disable"
}
func (*SwInterfaceSpanEnableDisable) GetCrcString() string { return "acc8fea1" }
func (*SwInterfaceSpanEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceSpanEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndexFrom
	size += 4 // m.SwIfIndexTo
	size += 4 // m.State
	size += 1 // m.IsL2
	return size
}
func (m *SwInterfaceSpanEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndexFrom))
	buf.EncodeUint32(uint32(m.SwIfIndexTo))
	buf.EncodeUint32(uint32(m.State))
	buf.EncodeBool(m.IsL2)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSpanEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndexFrom = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.SwIfIndexTo = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.State = SpanState(buf.DecodeUint32())
	m.IsL2 = buf.DecodeBool()
	return nil
}

// SwInterfaceSpanEnableDisableReply defines message 'sw_interface_span_enable_disable_reply'.
type SwInterfaceSpanEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SwInterfaceSpanEnableDisableReply) Reset() { *m = SwInterfaceSpanEnableDisableReply{} }
func (*SwInterfaceSpanEnableDisableReply) GetMessageName() string {
	return "sw_interface_span_enable_disable_reply"
}
func (*SwInterfaceSpanEnableDisableReply) GetCrcString() string { return "e8d4e804" }
func (*SwInterfaceSpanEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceSpanEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SwInterfaceSpanEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSpanEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_span_binapi_init() }
func file_span_binapi_init() {
	api.RegisterMessage((*SwInterfaceSpanDetails)(nil), "sw_interface_span_details_055643fc")
	api.RegisterMessage((*SwInterfaceSpanDump)(nil), "sw_interface_span_dump_d6cf0c3d")
	api.RegisterMessage((*SwInterfaceSpanEnableDisable)(nil), "sw_interface_span_enable_disable_acc8fea1")
	api.RegisterMessage((*SwInterfaceSpanEnableDisableReply)(nil), "sw_interface_span_enable_disable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SwInterfaceSpanDetails)(nil),
		(*SwInterfaceSpanDump)(nil),
		(*SwInterfaceSpanEnableDisable)(nil),
		(*SwInterfaceSpanEnableDisableReply)(nil),
	}
}
