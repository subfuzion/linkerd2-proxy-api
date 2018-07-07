// Code generated by protoc-gen-go. DO NOT EDIT.
// source: net.proto

package net // import "github.com/linkerd/linkerd2-proxy-api/go/net"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IPAddress struct {
	// Types that are valid to be assigned to Ip:
	//	*IPAddress_Ipv4
	//	*IPAddress_Ipv6
	Ip                   isIPAddress_Ip `protobuf_oneof:"ip"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *IPAddress) Reset()         { *m = IPAddress{} }
func (m *IPAddress) String() string { return proto.CompactTextString(m) }
func (*IPAddress) ProtoMessage()    {}
func (*IPAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_net_0e160d3f8ec67e7e, []int{0}
}
func (m *IPAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPAddress.Unmarshal(m, b)
}
func (m *IPAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPAddress.Marshal(b, m, deterministic)
}
func (dst *IPAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPAddress.Merge(dst, src)
}
func (m *IPAddress) XXX_Size() int {
	return xxx_messageInfo_IPAddress.Size(m)
}
func (m *IPAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_IPAddress.DiscardUnknown(m)
}

var xxx_messageInfo_IPAddress proto.InternalMessageInfo

type isIPAddress_Ip interface {
	isIPAddress_Ip()
}

type IPAddress_Ipv4 struct {
	Ipv4 uint32 `protobuf:"fixed32,1,opt,name=ipv4,proto3,oneof"`
}
type IPAddress_Ipv6 struct {
	Ipv6 *IPv6 `protobuf:"bytes,2,opt,name=ipv6,proto3,oneof"`
}

func (*IPAddress_Ipv4) isIPAddress_Ip() {}
func (*IPAddress_Ipv6) isIPAddress_Ip() {}

func (m *IPAddress) GetIp() isIPAddress_Ip {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *IPAddress) GetIpv4() uint32 {
	if x, ok := m.GetIp().(*IPAddress_Ipv4); ok {
		return x.Ipv4
	}
	return 0
}

func (m *IPAddress) GetIpv6() *IPv6 {
	if x, ok := m.GetIp().(*IPAddress_Ipv6); ok {
		return x.Ipv6
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*IPAddress) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _IPAddress_OneofMarshaler, _IPAddress_OneofUnmarshaler, _IPAddress_OneofSizer, []interface{}{
		(*IPAddress_Ipv4)(nil),
		(*IPAddress_Ipv6)(nil),
	}
}

func _IPAddress_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*IPAddress)
	// ip
	switch x := m.Ip.(type) {
	case *IPAddress_Ipv4:
		b.EncodeVarint(1<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(x.Ipv4))
	case *IPAddress_Ipv6:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ipv6); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("IPAddress.Ip has unexpected type %T", x)
	}
	return nil
}

func _IPAddress_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*IPAddress)
	switch tag {
	case 1: // ip.ipv4
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Ip = &IPAddress_Ipv4{uint32(x)}
		return true, err
	case 2: // ip.ipv6
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IPv6)
		err := b.DecodeMessage(msg)
		m.Ip = &IPAddress_Ipv6{msg}
		return true, err
	default:
		return false, nil
	}
}

func _IPAddress_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*IPAddress)
	// ip
	switch x := m.Ip.(type) {
	case *IPAddress_Ipv4:
		n += 1 // tag and wire
		n += 4
	case *IPAddress_Ipv6:
		s := proto.Size(x.Ipv6)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type IPv6 struct {
	First                uint64   `protobuf:"fixed64,1,opt,name=first,proto3" json:"first,omitempty"`
	Last                 uint64   `protobuf:"fixed64,2,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPv6) Reset()         { *m = IPv6{} }
func (m *IPv6) String() string { return proto.CompactTextString(m) }
func (*IPv6) ProtoMessage()    {}
func (*IPv6) Descriptor() ([]byte, []int) {
	return fileDescriptor_net_0e160d3f8ec67e7e, []int{1}
}
func (m *IPv6) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPv6.Unmarshal(m, b)
}
func (m *IPv6) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPv6.Marshal(b, m, deterministic)
}
func (dst *IPv6) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPv6.Merge(dst, src)
}
func (m *IPv6) XXX_Size() int {
	return xxx_messageInfo_IPv6.Size(m)
}
func (m *IPv6) XXX_DiscardUnknown() {
	xxx_messageInfo_IPv6.DiscardUnknown(m)
}

var xxx_messageInfo_IPv6 proto.InternalMessageInfo

func (m *IPv6) GetFirst() uint64 {
	if m != nil {
		return m.First
	}
	return 0
}

func (m *IPv6) GetLast() uint64 {
	if m != nil {
		return m.Last
	}
	return 0
}

type TcpAddress struct {
	Ip                   *IPAddress `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 uint32     `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TcpAddress) Reset()         { *m = TcpAddress{} }
func (m *TcpAddress) String() string { return proto.CompactTextString(m) }
func (*TcpAddress) ProtoMessage()    {}
func (*TcpAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_net_0e160d3f8ec67e7e, []int{2}
}
func (m *TcpAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpAddress.Unmarshal(m, b)
}
func (m *TcpAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpAddress.Marshal(b, m, deterministic)
}
func (dst *TcpAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpAddress.Merge(dst, src)
}
func (m *TcpAddress) XXX_Size() int {
	return xxx_messageInfo_TcpAddress.Size(m)
}
func (m *TcpAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpAddress.DiscardUnknown(m)
}

var xxx_messageInfo_TcpAddress proto.InternalMessageInfo

func (m *TcpAddress) GetIp() *IPAddress {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *TcpAddress) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*IPAddress)(nil), "io.linkerd.proxy.net.IPAddress")
	proto.RegisterType((*IPv6)(nil), "io.linkerd.proxy.net.IPv6")
	proto.RegisterType((*TcpAddress)(nil), "io.linkerd.proxy.net.TcpAddress")
}

func init() { proto.RegisterFile("net.proto", fileDescriptor_net_0e160d3f8ec67e7e) }

var fileDescriptor_net_0e160d3f8ec67e7e = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd0, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x06, 0x60, 0x12, 0x99, 0xa0, 0x5e, 0xc5, 0x62, 0x65, 0xa8, 0x58, 0xa8, 0x32, 0x75, 0xa0,
	0x76, 0x15, 0x50, 0x76, 0x3a, 0xd1, 0xad, 0x58, 0x2c, 0xb0, 0xb5, 0x8d, 0x29, 0x27, 0x8a, 0x6d,
	0xd9, 0x26, 0x82, 0xb7, 0x47, 0xbe, 0x24, 0x1b, 0x9d, 0x7c, 0xf7, 0xcb, 0xf7, 0xd9, 0x3a, 0x98,
	0x18, 0x1d, 0x85, 0xf3, 0x36, 0x5a, 0x5e, 0xa2, 0x15, 0x27, 0x34, 0x9f, 0xda, 0xb7, 0x29, 0xf9,
	0xf9, 0x15, 0x46, 0xc7, 0xea, 0x15, 0x26, 0x9b, 0xed, 0x63, 0xdb, 0x7a, 0x1d, 0x02, 0x2f, 0x81,
	0xa1, 0xeb, 0x1e, 0x66, 0xd9, 0x3c, 0x5b, 0x5c, 0x3d, 0x5d, 0x28, 0xea, 0xf8, 0x8a, 0xd2, 0x66,
	0x96, 0xcf, 0xb3, 0xc5, 0xb4, 0xbe, 0x11, 0xff, 0x39, 0x62, 0xb3, 0xed, 0x9a, 0x61, 0xa2, 0x59,
	0x33, 0xc8, 0xd1, 0x55, 0x2b, 0x60, 0x29, 0xe5, 0x25, 0x5c, 0xbe, 0xa3, 0x0f, 0x91, 0xd8, 0x42,
	0xf5, 0x0d, 0xe7, 0xc0, 0x4e, 0xbb, 0x10, 0x49, 0x2d, 0x14, 0xd5, 0xd5, 0x33, 0xc0, 0xcb, 0xc1,
	0x8d, 0xbf, 0x91, 0x49, 0xa1, 0xa1, 0x69, 0x7d, 0x7b, 0xee, 0xd5, 0xe1, 0xb2, 0xca, 0xd1, 0x25,
	0xd2, 0x59, 0xdf, 0x93, 0xd7, 0x8a, 0xea, 0xb5, 0x78, 0xbb, 0x3b, 0x62, 0xfc, 0xf8, 0xde, 0x8b,
	0x83, 0xfd, 0x92, 0x83, 0x30, 0x9e, 0xf5, 0x92, 0xa8, 0xe5, 0xce, 0xa1, 0x3c, 0x5a, 0x69, 0x74,
	0xdc, 0x17, 0xb4, 0xac, 0xfb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x93, 0x2c, 0x66, 0x39,
	0x01, 0x00, 0x00,
}
