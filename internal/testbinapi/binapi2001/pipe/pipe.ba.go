// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.01
// source: .vppapi/core/pipe.api.json

// Package pipe contains generated bindings for API file pipe.api.
//
// Contents:
//   1 alias
//   6 enums
//   6 messages
//
package pipe

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	"strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "pipe"
	APIVersion = "1.0.1"
	VersionCrc = 0x1d68f11
)

// IfStatusFlags defines enum 'if_status_flags'.
type IfStatusFlags uint32

const (
	IF_STATUS_API_FLAG_ADMIN_UP IfStatusFlags = 1
	IF_STATUS_API_FLAG_LINK_UP  IfStatusFlags = 2
)

var (
	IfStatusFlags_name = map[uint32]string{
		1: "IF_STATUS_API_FLAG_ADMIN_UP",
		2: "IF_STATUS_API_FLAG_LINK_UP",
	}
	IfStatusFlags_value = map[string]uint32{
		"IF_STATUS_API_FLAG_ADMIN_UP": 1,
		"IF_STATUS_API_FLAG_LINK_UP":  2,
	}
)

func (x IfStatusFlags) String() string {
	s, ok := IfStatusFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := IfStatusFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "IfStatusFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
}

// IfType defines enum 'if_type'.
type IfType uint32

const (
	IF_API_TYPE_HARDWARE IfType = 1
	IF_API_TYPE_SUB      IfType = 2
	IF_API_TYPE_P2P      IfType = 3
	IF_API_TYPE_PIPE     IfType = 4
)

var (
	IfType_name = map[uint32]string{
		1: "IF_API_TYPE_HARDWARE",
		2: "IF_API_TYPE_SUB",
		3: "IF_API_TYPE_P2P",
		4: "IF_API_TYPE_PIPE",
	}
	IfType_value = map[string]uint32{
		"IF_API_TYPE_HARDWARE": 1,
		"IF_API_TYPE_SUB":      2,
		"IF_API_TYPE_P2P":      3,
		"IF_API_TYPE_PIPE":     4,
	}
)

func (x IfType) String() string {
	s, ok := IfType_name[uint32(x)]
	if ok {
		return s
	}
	return "IfType(" + strconv.Itoa(int(x)) + ")"
}

// LinkDuplex defines enum 'link_duplex'.
type LinkDuplex uint32

const (
	LINK_DUPLEX_API_UNKNOWN LinkDuplex = 0
	LINK_DUPLEX_API_HALF    LinkDuplex = 1
	LINK_DUPLEX_API_FULL    LinkDuplex = 2
)

var (
	LinkDuplex_name = map[uint32]string{
		0: "LINK_DUPLEX_API_UNKNOWN",
		1: "LINK_DUPLEX_API_HALF",
		2: "LINK_DUPLEX_API_FULL",
	}
	LinkDuplex_value = map[string]uint32{
		"LINK_DUPLEX_API_UNKNOWN": 0,
		"LINK_DUPLEX_API_HALF":    1,
		"LINK_DUPLEX_API_FULL":    2,
	}
)

func (x LinkDuplex) String() string {
	s, ok := LinkDuplex_name[uint32(x)]
	if ok {
		return s
	}
	return "LinkDuplex(" + strconv.Itoa(int(x)) + ")"
}

// MtuProto defines enum 'mtu_proto'.
type MtuProto uint32

const (
	MTU_PROTO_API_L3   MtuProto = 1
	MTU_PROTO_API_IP4  MtuProto = 2
	MTU_PROTO_API_IP6  MtuProto = 3
	MTU_PROTO_API_MPLS MtuProto = 4
	MTU_PROTO_API_N    MtuProto = 5
)

var (
	MtuProto_name = map[uint32]string{
		1: "MTU_PROTO_API_L3",
		2: "MTU_PROTO_API_IP4",
		3: "MTU_PROTO_API_IP6",
		4: "MTU_PROTO_API_MPLS",
		5: "MTU_PROTO_API_N",
	}
	MtuProto_value = map[string]uint32{
		"MTU_PROTO_API_L3":   1,
		"MTU_PROTO_API_IP4":  2,
		"MTU_PROTO_API_IP6":  3,
		"MTU_PROTO_API_MPLS": 4,
		"MTU_PROTO_API_N":    5,
	}
)

func (x MtuProto) String() string {
	s, ok := MtuProto_name[uint32(x)]
	if ok {
		return s
	}
	return "MtuProto(" + strconv.Itoa(int(x)) + ")"
}

// RxMode defines enum 'rx_mode'.
type RxMode uint32

const (
	RX_MODE_API_UNKNOWN   RxMode = 0
	RX_MODE_API_POLLING   RxMode = 1
	RX_MODE_API_INTERRUPT RxMode = 2
	RX_MODE_API_ADAPTIVE  RxMode = 3
	RX_MODE_API_DEFAULT   RxMode = 4
)

var (
	RxMode_name = map[uint32]string{
		0: "RX_MODE_API_UNKNOWN",
		1: "RX_MODE_API_POLLING",
		2: "RX_MODE_API_INTERRUPT",
		3: "RX_MODE_API_ADAPTIVE",
		4: "RX_MODE_API_DEFAULT",
	}
	RxMode_value = map[string]uint32{
		"RX_MODE_API_UNKNOWN":   0,
		"RX_MODE_API_POLLING":   1,
		"RX_MODE_API_INTERRUPT": 2,
		"RX_MODE_API_ADAPTIVE":  3,
		"RX_MODE_API_DEFAULT":   4,
	}
)

func (x RxMode) String() string {
	s, ok := RxMode_name[uint32(x)]
	if ok {
		return s
	}
	return "RxMode(" + strconv.Itoa(int(x)) + ")"
}

// SubIfFlags defines enum 'sub_if_flags'.
type SubIfFlags uint32

const (
	SUB_IF_API_FLAG_NO_TAGS           SubIfFlags = 1
	SUB_IF_API_FLAG_ONE_TAG           SubIfFlags = 2
	SUB_IF_API_FLAG_TWO_TAGS          SubIfFlags = 4
	SUB_IF_API_FLAG_DOT1AD            SubIfFlags = 8
	SUB_IF_API_FLAG_EXACT_MATCH       SubIfFlags = 16
	SUB_IF_API_FLAG_DEFAULT           SubIfFlags = 32
	SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY SubIfFlags = 64
	SUB_IF_API_FLAG_INNER_VLAN_ID_ANY SubIfFlags = 128
	SUB_IF_API_FLAG_MASK_VNET         SubIfFlags = 254
	SUB_IF_API_FLAG_DOT1AH            SubIfFlags = 256
)

var (
	SubIfFlags_name = map[uint32]string{
		1:   "SUB_IF_API_FLAG_NO_TAGS",
		2:   "SUB_IF_API_FLAG_ONE_TAG",
		4:   "SUB_IF_API_FLAG_TWO_TAGS",
		8:   "SUB_IF_API_FLAG_DOT1AD",
		16:  "SUB_IF_API_FLAG_EXACT_MATCH",
		32:  "SUB_IF_API_FLAG_DEFAULT",
		64:  "SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY",
		128: "SUB_IF_API_FLAG_INNER_VLAN_ID_ANY",
		254: "SUB_IF_API_FLAG_MASK_VNET",
		256: "SUB_IF_API_FLAG_DOT1AH",
	}
	SubIfFlags_value = map[string]uint32{
		"SUB_IF_API_FLAG_NO_TAGS":           1,
		"SUB_IF_API_FLAG_ONE_TAG":           2,
		"SUB_IF_API_FLAG_TWO_TAGS":          4,
		"SUB_IF_API_FLAG_DOT1AD":            8,
		"SUB_IF_API_FLAG_EXACT_MATCH":       16,
		"SUB_IF_API_FLAG_DEFAULT":           32,
		"SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY": 64,
		"SUB_IF_API_FLAG_INNER_VLAN_ID_ANY": 128,
		"SUB_IF_API_FLAG_MASK_VNET":         254,
		"SUB_IF_API_FLAG_DOT1AH":            256,
	}
)

func (x SubIfFlags) String() string {
	s, ok := SubIfFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := SubIfFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "SubIfFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
}

// InterfaceIndex defines alias 'interface_index'.
type InterfaceIndex uint32

// PipeCreate defines message 'pipe_create'.
type PipeCreate struct {
	IsSpecified  bool   `binapi:"bool,name=is_specified" json:"is_specified,omitempty"`
	UserInstance uint32 `binapi:"u32,name=user_instance" json:"user_instance,omitempty"`
}

func (m *PipeCreate) Reset()               { *m = PipeCreate{} }
func (*PipeCreate) GetMessageName() string { return "pipe_create" }
func (*PipeCreate) GetCrcString() string   { return "bb263bd3" }
func (*PipeCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PipeCreate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsSpecified
	size += 4 // m.UserInstance
	return size
}
func (m *PipeCreate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsSpecified)
	buf.EncodeUint32(m.UserInstance)
	return buf.Bytes(), nil
}
func (m *PipeCreate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsSpecified = buf.DecodeBool()
	m.UserInstance = buf.DecodeUint32()
	return nil
}

// PipeCreateReply defines message 'pipe_create_reply'.
type PipeCreateReply struct {
	Retval        int32             `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex     InterfaceIndex    `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	PipeSwIfIndex [2]InterfaceIndex `binapi:"interface_index[2],name=pipe_sw_if_index" json:"pipe_sw_if_index,omitempty"`
}

func (m *PipeCreateReply) Reset()               { *m = PipeCreateReply{} }
func (*PipeCreateReply) GetMessageName() string { return "pipe_create_reply" }
func (*PipeCreateReply) GetCrcString() string   { return "d4c2c2b3" }
func (*PipeCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PipeCreateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	for j1 := 0; j1 < 2; j1++ {
		size += 4 // m.PipeSwIfIndex[j1]
	}
	return size
}
func (m *PipeCreateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	for j0 := 0; j0 < 2; j0++ {
		buf.EncodeUint32(uint32(m.PipeSwIfIndex[j0]))
	}
	return buf.Bytes(), nil
}
func (m *PipeCreateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	for j0 := 0; j0 < 2; j0++ {
		m.PipeSwIfIndex[j0] = InterfaceIndex(buf.DecodeUint32())
	}
	return nil
}

// PipeDelete defines message 'pipe_delete'.
type PipeDelete struct {
	SwIfIndex InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *PipeDelete) Reset()               { *m = PipeDelete{} }
func (*PipeDelete) GetMessageName() string { return "pipe_delete" }
func (*PipeDelete) GetCrcString() string   { return "f9e6675e" }
func (*PipeDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PipeDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *PipeDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *PipeDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	return nil
}

// PipeDeleteReply defines message 'pipe_delete_reply'.
type PipeDeleteReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *PipeDeleteReply) Reset()               { *m = PipeDeleteReply{} }
func (*PipeDeleteReply) GetMessageName() string { return "pipe_delete_reply" }
func (*PipeDeleteReply) GetCrcString() string   { return "e8d4e804" }
func (*PipeDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PipeDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *PipeDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *PipeDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// PipeDetails defines message 'pipe_details'.
type PipeDetails struct {
	SwIfIndex     InterfaceIndex    `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	PipeSwIfIndex [2]InterfaceIndex `binapi:"interface_index[2],name=pipe_sw_if_index" json:"pipe_sw_if_index,omitempty"`
	Instance      uint32            `binapi:"u32,name=instance" json:"instance,omitempty"`
}

func (m *PipeDetails) Reset()               { *m = PipeDetails{} }
func (*PipeDetails) GetMessageName() string { return "pipe_details" }
func (*PipeDetails) GetCrcString() string   { return "43ac107a" }
func (*PipeDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PipeDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	for j1 := 0; j1 < 2; j1++ {
		size += 4 // m.PipeSwIfIndex[j1]
	}
	size += 4 // m.Instance
	return size
}
func (m *PipeDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	for j0 := 0; j0 < 2; j0++ {
		buf.EncodeUint32(uint32(m.PipeSwIfIndex[j0]))
	}
	buf.EncodeUint32(m.Instance)
	return buf.Bytes(), nil
}
func (m *PipeDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	for j0 := 0; j0 < 2; j0++ {
		m.PipeSwIfIndex[j0] = InterfaceIndex(buf.DecodeUint32())
	}
	m.Instance = buf.DecodeUint32()
	return nil
}

// PipeDump defines message 'pipe_dump'.
type PipeDump struct{}

func (m *PipeDump) Reset()               { *m = PipeDump{} }
func (*PipeDump) GetMessageName() string { return "pipe_dump" }
func (*PipeDump) GetCrcString() string   { return "51077d14" }
func (*PipeDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PipeDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *PipeDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *PipeDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_pipe_binapi_init() }
func file_pipe_binapi_init() {
	api.RegisterMessage((*PipeCreate)(nil), "pipe_create_bb263bd3")
	api.RegisterMessage((*PipeCreateReply)(nil), "pipe_create_reply_d4c2c2b3")
	api.RegisterMessage((*PipeDelete)(nil), "pipe_delete_f9e6675e")
	api.RegisterMessage((*PipeDeleteReply)(nil), "pipe_delete_reply_e8d4e804")
	api.RegisterMessage((*PipeDetails)(nil), "pipe_details_43ac107a")
	api.RegisterMessage((*PipeDump)(nil), "pipe_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*PipeCreate)(nil),
		(*PipeCreateReply)(nil),
		(*PipeDelete)(nil),
		(*PipeDeleteReply)(nil),
		(*PipeDetails)(nil),
		(*PipeDump)(nil),
	}
}
