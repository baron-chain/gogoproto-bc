// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue444/issue444.proto

package issue444

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SizeMe struct {
	Foo                  string   `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SizeMe) Reset()         { *m = SizeMe{} }
func (m *SizeMe) String() string { return proto.CompactTextString(m) }
func (*SizeMe) ProtoMessage()    {}
func (*SizeMe) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a9f5042d428c4e2, []int{0}
}
func (m *SizeMe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SizeMe.Unmarshal(m, b)
}
func (m *SizeMe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SizeMe.Marshal(b, m, deterministic)
}
func (m *SizeMe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SizeMe.Merge(m, src)
}
func (m *SizeMe) XXX_Size() int {
	return xxx_messageInfo_SizeMe.Size(m)
}
func (m *SizeMe) XXX_DiscardUnknown() {
	xxx_messageInfo_SizeMe.DiscardUnknown(m)
}

var xxx_messageInfo_SizeMe proto.InternalMessageInfo

func (m *SizeMe) GetFoo() string {
	if m != nil {
		return m.Foo
	}
	return ""
}

func init() {
	proto.RegisterType((*SizeMe)(nil), "issue444.SizeMe")
}

func init() { proto.RegisterFile("issue444/issue444.proto", fileDescriptor_0a9f5042d428c4e2) }

var fileDescriptor_0a9f5042d428c4e2 = []byte{
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x31, 0x31, 0xd1, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60,
	0x7c, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0xa0, 0x3e, 0x88, 0x05, 0x91, 0x57, 0x92, 0xe2,
	0x62, 0x0b, 0xce, 0xac, 0x4a, 0xf5, 0x4d, 0x15, 0x12, 0xe0, 0x62, 0x4e, 0xcb, 0xcf, 0x97, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x9d, 0x58, 0x1e, 0x3c, 0x92, 0x63, 0x4c, 0x62, 0x03,
	0x2b, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xac, 0x21, 0xf5, 0x63, 0x00, 0x00, 0x00,
}

func (m *SizeMe) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Foo)
	if l > 0 {
		n += 1 + l + sovIssue444(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIssue444(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIssue444(x uint64) (n int) {
	return sovIssue444(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}