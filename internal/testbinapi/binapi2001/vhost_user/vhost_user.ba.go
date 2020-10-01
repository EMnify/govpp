// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.01
// source: .vppapi/core/vhost_user.api.json

// Package vhost_user contains generated bindings for API file vhost_user.api.
//
// Contents:
//   2 aliases
//   8 enums
//   8 messages
//
package vhost_user

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	"net"
	"strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "vhost_user"
	APIVersion = "4.0.0"
	VersionCrc = 0xb547a33d
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

// VirtioNetFeaturesFirst32 defines enum 'virtio_net_features_first_32'.
type VirtioNetFeaturesFirst32 uint32

const (
	VIRTIO_NET_F_API_CSUM              VirtioNetFeaturesFirst32 = 1
	VIRTIO_NET_F_API_GUEST_CSUM        VirtioNetFeaturesFirst32 = 2
	VIRTIO_NET_F_API_GUEST_TSO4        VirtioNetFeaturesFirst32 = 128
	VIRTIO_NET_F_API_GUEST_TSO6        VirtioNetFeaturesFirst32 = 256
	VIRTIO_NET_F_API_GUEST_UFO         VirtioNetFeaturesFirst32 = 1024
	VIRTIO_NET_F_API_HOST_TSO4         VirtioNetFeaturesFirst32 = 2048
	VIRTIO_NET_F_API_HOST_TSO6         VirtioNetFeaturesFirst32 = 4096
	VIRTIO_NET_F_API_HOST_UFO          VirtioNetFeaturesFirst32 = 16384
	VIRTIO_NET_F_API_MRG_RXBUF         VirtioNetFeaturesFirst32 = 32768
	VIRTIO_NET_F_API_CTRL_VQ           VirtioNetFeaturesFirst32 = 131072
	VIRTIO_NET_F_API_GUEST_ANNOUNCE    VirtioNetFeaturesFirst32 = 2097152
	VIRTIO_NET_F_API_MQ                VirtioNetFeaturesFirst32 = 4194304
	VHOST_F_API_LOG_ALL                VirtioNetFeaturesFirst32 = 67108864
	VIRTIO_F_API_ANY_LAYOUT            VirtioNetFeaturesFirst32 = 134217728
	VIRTIO_F_API_INDIRECT_DESC         VirtioNetFeaturesFirst32 = 268435456
	VHOST_USER_F_API_PROTOCOL_FEATURES VirtioNetFeaturesFirst32 = 1073741824
)

var (
	VirtioNetFeaturesFirst32_name = map[uint32]string{
		1:          "VIRTIO_NET_F_API_CSUM",
		2:          "VIRTIO_NET_F_API_GUEST_CSUM",
		128:        "VIRTIO_NET_F_API_GUEST_TSO4",
		256:        "VIRTIO_NET_F_API_GUEST_TSO6",
		1024:       "VIRTIO_NET_F_API_GUEST_UFO",
		2048:       "VIRTIO_NET_F_API_HOST_TSO4",
		4096:       "VIRTIO_NET_F_API_HOST_TSO6",
		16384:      "VIRTIO_NET_F_API_HOST_UFO",
		32768:      "VIRTIO_NET_F_API_MRG_RXBUF",
		131072:     "VIRTIO_NET_F_API_CTRL_VQ",
		2097152:    "VIRTIO_NET_F_API_GUEST_ANNOUNCE",
		4194304:    "VIRTIO_NET_F_API_MQ",
		67108864:   "VHOST_F_API_LOG_ALL",
		134217728:  "VIRTIO_F_API_ANY_LAYOUT",
		268435456:  "VIRTIO_F_API_INDIRECT_DESC",
		1073741824: "VHOST_USER_F_API_PROTOCOL_FEATURES",
	}
	VirtioNetFeaturesFirst32_value = map[string]uint32{
		"VIRTIO_NET_F_API_CSUM":              1,
		"VIRTIO_NET_F_API_GUEST_CSUM":        2,
		"VIRTIO_NET_F_API_GUEST_TSO4":        128,
		"VIRTIO_NET_F_API_GUEST_TSO6":        256,
		"VIRTIO_NET_F_API_GUEST_UFO":         1024,
		"VIRTIO_NET_F_API_HOST_TSO4":         2048,
		"VIRTIO_NET_F_API_HOST_TSO6":         4096,
		"VIRTIO_NET_F_API_HOST_UFO":          16384,
		"VIRTIO_NET_F_API_MRG_RXBUF":         32768,
		"VIRTIO_NET_F_API_CTRL_VQ":           131072,
		"VIRTIO_NET_F_API_GUEST_ANNOUNCE":    2097152,
		"VIRTIO_NET_F_API_MQ":                4194304,
		"VHOST_F_API_LOG_ALL":                67108864,
		"VIRTIO_F_API_ANY_LAYOUT":            134217728,
		"VIRTIO_F_API_INDIRECT_DESC":         268435456,
		"VHOST_USER_F_API_PROTOCOL_FEATURES": 1073741824,
	}
)

func (x VirtioNetFeaturesFirst32) String() string {
	s, ok := VirtioNetFeaturesFirst32_name[uint32(x)]
	if ok {
		return s
	}
	return "VirtioNetFeaturesFirst32(" + strconv.Itoa(int(x)) + ")"
}

// VirtioNetFeaturesLast32 defines enum 'virtio_net_features_last_32'.
type VirtioNetFeaturesLast32 uint32

const (
	VIRTIO_F_API_VERSION_1 VirtioNetFeaturesLast32 = 1
)

var (
	VirtioNetFeaturesLast32_name = map[uint32]string{
		1: "VIRTIO_F_API_VERSION_1",
	}
	VirtioNetFeaturesLast32_value = map[string]uint32{
		"VIRTIO_F_API_VERSION_1": 1,
	}
)

func (x VirtioNetFeaturesLast32) String() string {
	s, ok := VirtioNetFeaturesLast32_name[uint32(x)]
	if ok {
		return s
	}
	return "VirtioNetFeaturesLast32(" + strconv.Itoa(int(x)) + ")"
}

// InterfaceIndex defines alias 'interface_index'.
type InterfaceIndex uint32

// MacAddress defines alias 'mac_address'.
type MacAddress [6]uint8

func ParseMacAddress(s string) (MacAddress, error) {
	var macaddr MacAddress
	mac, err := net.ParseMAC(s)
	if err != nil {
		return macaddr, err
	}
	copy(macaddr[:], mac[:])
	return macaddr, nil
}
func (x MacAddress) ToMAC() net.HardwareAddr {
	return net.HardwareAddr(x[:])
}
func (x MacAddress) String() string {
	return x.ToMAC().String()
}
func (x *MacAddress) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}
func (x *MacAddress) UnmarshalText(text []byte) error {
	mac, err := ParseMacAddress(string(text))
	if err != nil {
		return err
	}
	*x = mac
	return nil
}

// CreateVhostUserIf defines message 'create_vhost_user_if'.
type CreateVhostUserIf struct {
	IsServer            bool       `binapi:"bool,name=is_server" json:"is_server,omitempty"`
	SockFilename        string     `binapi:"string[256],name=sock_filename" json:"sock_filename,omitempty"`
	Renumber            bool       `binapi:"bool,name=renumber" json:"renumber,omitempty"`
	DisableMrgRxbuf     bool       `binapi:"bool,name=disable_mrg_rxbuf" json:"disable_mrg_rxbuf,omitempty"`
	DisableIndirectDesc bool       `binapi:"bool,name=disable_indirect_desc" json:"disable_indirect_desc,omitempty"`
	EnableGso           bool       `binapi:"bool,name=enable_gso" json:"enable_gso,omitempty"`
	CustomDevInstance   uint32     `binapi:"u32,name=custom_dev_instance" json:"custom_dev_instance,omitempty"`
	UseCustomMac        bool       `binapi:"bool,name=use_custom_mac" json:"use_custom_mac,omitempty"`
	MacAddress          MacAddress `binapi:"mac_address,name=mac_address" json:"mac_address,omitempty"`
	Tag                 string     `binapi:"string[64],name=tag" json:"tag,omitempty"`
}

func (m *CreateVhostUserIf) Reset()               { *m = CreateVhostUserIf{} }
func (*CreateVhostUserIf) GetMessageName() string { return "create_vhost_user_if" }
func (*CreateVhostUserIf) GetCrcString() string   { return "591ee951" }
func (*CreateVhostUserIf) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CreateVhostUserIf) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1     // m.IsServer
	size += 256   // m.SockFilename
	size += 1     // m.Renumber
	size += 1     // m.DisableMrgRxbuf
	size += 1     // m.DisableIndirectDesc
	size += 1     // m.EnableGso
	size += 4     // m.CustomDevInstance
	size += 1     // m.UseCustomMac
	size += 1 * 6 // m.MacAddress
	size += 64    // m.Tag
	return size
}
func (m *CreateVhostUserIf) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsServer)
	buf.EncodeString(m.SockFilename, 256)
	buf.EncodeBool(m.Renumber)
	buf.EncodeBool(m.DisableMrgRxbuf)
	buf.EncodeBool(m.DisableIndirectDesc)
	buf.EncodeBool(m.EnableGso)
	buf.EncodeUint32(m.CustomDevInstance)
	buf.EncodeBool(m.UseCustomMac)
	buf.EncodeBytes(m.MacAddress[:], 6)
	buf.EncodeString(m.Tag, 64)
	return buf.Bytes(), nil
}
func (m *CreateVhostUserIf) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsServer = buf.DecodeBool()
	m.SockFilename = buf.DecodeString(256)
	m.Renumber = buf.DecodeBool()
	m.DisableMrgRxbuf = buf.DecodeBool()
	m.DisableIndirectDesc = buf.DecodeBool()
	m.EnableGso = buf.DecodeBool()
	m.CustomDevInstance = buf.DecodeUint32()
	m.UseCustomMac = buf.DecodeBool()
	copy(m.MacAddress[:], buf.DecodeBytes(6))
	m.Tag = buf.DecodeString(64)
	return nil
}

// CreateVhostUserIfReply defines message 'create_vhost_user_if_reply'.
type CreateVhostUserIfReply struct {
	Retval    int32          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *CreateVhostUserIfReply) Reset()               { *m = CreateVhostUserIfReply{} }
func (*CreateVhostUserIfReply) GetMessageName() string { return "create_vhost_user_if_reply" }
func (*CreateVhostUserIfReply) GetCrcString() string   { return "5383d31f" }
func (*CreateVhostUserIfReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CreateVhostUserIfReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *CreateVhostUserIfReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *CreateVhostUserIfReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	return nil
}

// DeleteVhostUserIf defines message 'delete_vhost_user_if'.
type DeleteVhostUserIf struct {
	SwIfIndex InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *DeleteVhostUserIf) Reset()               { *m = DeleteVhostUserIf{} }
func (*DeleteVhostUserIf) GetMessageName() string { return "delete_vhost_user_if" }
func (*DeleteVhostUserIf) GetCrcString() string   { return "f9e6675e" }
func (*DeleteVhostUserIf) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DeleteVhostUserIf) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *DeleteVhostUserIf) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *DeleteVhostUserIf) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	return nil
}

// DeleteVhostUserIfReply defines message 'delete_vhost_user_if_reply'.
type DeleteVhostUserIfReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *DeleteVhostUserIfReply) Reset()               { *m = DeleteVhostUserIfReply{} }
func (*DeleteVhostUserIfReply) GetMessageName() string { return "delete_vhost_user_if_reply" }
func (*DeleteVhostUserIfReply) GetCrcString() string   { return "e8d4e804" }
func (*DeleteVhostUserIfReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DeleteVhostUserIfReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *DeleteVhostUserIfReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *DeleteVhostUserIfReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// ModifyVhostUserIf defines message 'modify_vhost_user_if'.
type ModifyVhostUserIf struct {
	SwIfIndex         InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IsServer          bool           `binapi:"bool,name=is_server" json:"is_server,omitempty"`
	SockFilename      string         `binapi:"string[256],name=sock_filename" json:"sock_filename,omitempty"`
	Renumber          bool           `binapi:"bool,name=renumber" json:"renumber,omitempty"`
	EnableGso         bool           `binapi:"bool,name=enable_gso" json:"enable_gso,omitempty"`
	CustomDevInstance uint32         `binapi:"u32,name=custom_dev_instance" json:"custom_dev_instance,omitempty"`
}

func (m *ModifyVhostUserIf) Reset()               { *m = ModifyVhostUserIf{} }
func (*ModifyVhostUserIf) GetMessageName() string { return "modify_vhost_user_if" }
func (*ModifyVhostUserIf) GetCrcString() string   { return "fcfeaf16" }
func (*ModifyVhostUserIf) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ModifyVhostUserIf) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4   // m.SwIfIndex
	size += 1   // m.IsServer
	size += 256 // m.SockFilename
	size += 1   // m.Renumber
	size += 1   // m.EnableGso
	size += 4   // m.CustomDevInstance
	return size
}
func (m *ModifyVhostUserIf) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.IsServer)
	buf.EncodeString(m.SockFilename, 256)
	buf.EncodeBool(m.Renumber)
	buf.EncodeBool(m.EnableGso)
	buf.EncodeUint32(m.CustomDevInstance)
	return buf.Bytes(), nil
}
func (m *ModifyVhostUserIf) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	m.IsServer = buf.DecodeBool()
	m.SockFilename = buf.DecodeString(256)
	m.Renumber = buf.DecodeBool()
	m.EnableGso = buf.DecodeBool()
	m.CustomDevInstance = buf.DecodeUint32()
	return nil
}

// ModifyVhostUserIfReply defines message 'modify_vhost_user_if_reply'.
type ModifyVhostUserIfReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *ModifyVhostUserIfReply) Reset()               { *m = ModifyVhostUserIfReply{} }
func (*ModifyVhostUserIfReply) GetMessageName() string { return "modify_vhost_user_if_reply" }
func (*ModifyVhostUserIfReply) GetCrcString() string   { return "e8d4e804" }
func (*ModifyVhostUserIfReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ModifyVhostUserIfReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *ModifyVhostUserIfReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *ModifyVhostUserIfReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SwInterfaceVhostUserDetails defines message 'sw_interface_vhost_user_details'.
type SwInterfaceVhostUserDetails struct {
	SwIfIndex       InterfaceIndex           `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	InterfaceName   string                   `binapi:"string[64],name=interface_name" json:"interface_name,omitempty"`
	VirtioNetHdrSz  uint32                   `binapi:"u32,name=virtio_net_hdr_sz" json:"virtio_net_hdr_sz,omitempty"`
	FeaturesFirst32 VirtioNetFeaturesFirst32 `binapi:"virtio_net_features_first_32,name=features_first_32" json:"features_first_32,omitempty"`
	FeaturesLast32  VirtioNetFeaturesLast32  `binapi:"virtio_net_features_last_32,name=features_last_32" json:"features_last_32,omitempty"`
	IsServer        bool                     `binapi:"bool,name=is_server" json:"is_server,omitempty"`
	SockFilename    string                   `binapi:"string[256],name=sock_filename" json:"sock_filename,omitempty"`
	NumRegions      uint32                   `binapi:"u32,name=num_regions" json:"num_regions,omitempty"`
	SockErrno       int32                    `binapi:"i32,name=sock_errno" json:"sock_errno,omitempty"`
}

func (m *SwInterfaceVhostUserDetails) Reset()               { *m = SwInterfaceVhostUserDetails{} }
func (*SwInterfaceVhostUserDetails) GetMessageName() string { return "sw_interface_vhost_user_details" }
func (*SwInterfaceVhostUserDetails) GetCrcString() string   { return "98530df1" }
func (*SwInterfaceVhostUserDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceVhostUserDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4   // m.SwIfIndex
	size += 64  // m.InterfaceName
	size += 4   // m.VirtioNetHdrSz
	size += 4   // m.FeaturesFirst32
	size += 4   // m.FeaturesLast32
	size += 1   // m.IsServer
	size += 256 // m.SockFilename
	size += 4   // m.NumRegions
	size += 4   // m.SockErrno
	return size
}
func (m *SwInterfaceVhostUserDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeString(m.InterfaceName, 64)
	buf.EncodeUint32(m.VirtioNetHdrSz)
	buf.EncodeUint32(uint32(m.FeaturesFirst32))
	buf.EncodeUint32(uint32(m.FeaturesLast32))
	buf.EncodeBool(m.IsServer)
	buf.EncodeString(m.SockFilename, 256)
	buf.EncodeUint32(m.NumRegions)
	buf.EncodeInt32(m.SockErrno)
	return buf.Bytes(), nil
}
func (m *SwInterfaceVhostUserDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	m.InterfaceName = buf.DecodeString(64)
	m.VirtioNetHdrSz = buf.DecodeUint32()
	m.FeaturesFirst32 = VirtioNetFeaturesFirst32(buf.DecodeUint32())
	m.FeaturesLast32 = VirtioNetFeaturesLast32(buf.DecodeUint32())
	m.IsServer = buf.DecodeBool()
	m.SockFilename = buf.DecodeString(256)
	m.NumRegions = buf.DecodeUint32()
	m.SockErrno = buf.DecodeInt32()
	return nil
}

// SwInterfaceVhostUserDump defines message 'sw_interface_vhost_user_dump'.
type SwInterfaceVhostUserDump struct {
	SwIfIndex InterfaceIndex `binapi:"interface_index,name=sw_if_index,default=4294967295" json:"sw_if_index,omitempty"`
}

func (m *SwInterfaceVhostUserDump) Reset()               { *m = SwInterfaceVhostUserDump{} }
func (*SwInterfaceVhostUserDump) GetMessageName() string { return "sw_interface_vhost_user_dump" }
func (*SwInterfaceVhostUserDump) GetCrcString() string   { return "f9e6675e" }
func (*SwInterfaceVhostUserDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceVhostUserDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *SwInterfaceVhostUserDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *SwInterfaceVhostUserDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	return nil
}

func init() { file_vhost_user_binapi_init() }
func file_vhost_user_binapi_init() {
	api.RegisterMessage((*CreateVhostUserIf)(nil), "create_vhost_user_if_591ee951")
	api.RegisterMessage((*CreateVhostUserIfReply)(nil), "create_vhost_user_if_reply_5383d31f")
	api.RegisterMessage((*DeleteVhostUserIf)(nil), "delete_vhost_user_if_f9e6675e")
	api.RegisterMessage((*DeleteVhostUserIfReply)(nil), "delete_vhost_user_if_reply_e8d4e804")
	api.RegisterMessage((*ModifyVhostUserIf)(nil), "modify_vhost_user_if_fcfeaf16")
	api.RegisterMessage((*ModifyVhostUserIfReply)(nil), "modify_vhost_user_if_reply_e8d4e804")
	api.RegisterMessage((*SwInterfaceVhostUserDetails)(nil), "sw_interface_vhost_user_details_98530df1")
	api.RegisterMessage((*SwInterfaceVhostUserDump)(nil), "sw_interface_vhost_user_dump_f9e6675e")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*CreateVhostUserIf)(nil),
		(*CreateVhostUserIfReply)(nil),
		(*DeleteVhostUserIf)(nil),
		(*DeleteVhostUserIfReply)(nil),
		(*ModifyVhostUserIf)(nil),
		(*ModifyVhostUserIfReply)(nil),
		(*SwInterfaceVhostUserDetails)(nil),
		(*SwInterfaceVhostUserDump)(nil),
	}
}
