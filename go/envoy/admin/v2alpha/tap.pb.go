// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/tap.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	v2alpha "github.com/cilium/proxy/go/envoy/service/tap/v2alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The /tap admin request body that is used to configure an active tap session.
type TapRequest struct {
	// The opaque configuration ID used to match the configuration to a loaded extension.
	// A tap extension configures a similar opaque ID that is used to match.
	ConfigId string `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	// The tap configuration to load.
	TapConfig            *v2alpha.TapConfig `protobuf:"bytes,2,opt,name=tap_config,json=tapConfig,proto3" json:"tap_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TapRequest) Reset()         { *m = TapRequest{} }
func (m *TapRequest) String() string { return proto.CompactTextString(m) }
func (*TapRequest) ProtoMessage()    {}
func (*TapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_69ef50f493e7b8aa, []int{0}
}

func (m *TapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TapRequest.Unmarshal(m, b)
}
func (m *TapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TapRequest.Marshal(b, m, deterministic)
}
func (m *TapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapRequest.Merge(m, src)
}
func (m *TapRequest) XXX_Size() int {
	return xxx_messageInfo_TapRequest.Size(m)
}
func (m *TapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TapRequest proto.InternalMessageInfo

func (m *TapRequest) GetConfigId() string {
	if m != nil {
		return m.ConfigId
	}
	return ""
}

func (m *TapRequest) GetTapConfig() *v2alpha.TapConfig {
	if m != nil {
		return m.TapConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*TapRequest)(nil), "envoy.admin.v2alpha.TapRequest")
}

func init() { proto.RegisterFile("envoy/admin/v2alpha/tap.proto", fileDescriptor_69ef50f493e7b8aa) }

var fileDescriptor_69ef50f493e7b8aa = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0x87, 0x79, 0x8b, 0x7f, 0xda, 0x78, 0x91, 0x7a, 0x70, 0x0c, 0x84, 0x2a, 0x43, 0x76, 0x4a,
	0x60, 0x7e, 0x83, 0x78, 0xf2, 0x20, 0x8c, 0xd2, 0xfb, 0x78, 0x6d, 0xa3, 0x06, 0xd6, 0xbe, 0xaf,
	0x59, 0x0c, 0xee, 0xe8, 0xd5, 0x8f, 0xbc, 0x93, 0x34, 0xa9, 0x3b, 0x79, 0x4b, 0xf2, 0xfc, 0x1e,
	0x78, 0x22, 0x6e, 0xcc, 0x10, 0x68, 0xaf, 0xb0, 0xeb, 0xed, 0xa0, 0xc2, 0x0a, 0xb7, 0xfc, 0x8e,
	0xca, 0x23, 0x4b, 0x76, 0xe4, 0xa9, 0xbc, 0x8a, 0x58, 0x46, 0x2c, 0x27, 0x3c, 0xbf, 0x4f, 0xce,
	0xce, 0xb8, 0x60, 0x5b, 0x33, 0xae, 0x8f, 0x66, 0x4b, 0x7d, 0x4f, 0x43, 0x92, 0xe7, 0xd7, 0x01,
	0xb7, 0xb6, 0x43, 0x6f, 0xd4, 0xdf, 0x21, 0x81, 0xbb, 0x6f, 0x10, 0xa2, 0x41, 0xae, 0xcd, 0xc7,
	0xa7, 0xd9, 0xf9, 0x72, 0x21, 0x8a, 0x96, 0x86, 0x57, 0xfb, 0xb6, 0xb1, 0xdd, 0x0c, 0x2a, 0x58,
	0x16, 0xfa, 0xfc, 0xa0, 0x4f, 0x5c, 0x56, 0x41, 0x9d, 0x27, 0xf2, 0xd4, 0x95, 0xcf, 0x42, 0x78,
	0xe4, 0x4d, 0xba, 0xcf, 0xb2, 0x0a, 0x96, 0x17, 0xab, 0x85, 0x4c, 0x7d, 0x53, 0x8a, 0x1c, 0xc3,
	0xa7, 0x14, 0xd9, 0x20, 0x3f, 0xc6, 0xad, 0xce, 0x0f, 0xfa, 0xf4, 0x07, 0xb2, 0x4b, 0xa8, 0x0b,
	0x7f, 0x7c, 0x54, 0xe2, 0xd6, 0x52, 0xd2, 0xd9, 0xd1, 0xd7, 0x5e, 0xfe, 0xf3, 0x53, 0x9d, 0x37,
	0xc8, 0xeb, 0x31, 0x79, 0x0d, 0x2f, 0x67, 0xb1, 0xfd, 0xe1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x9a,
	0x67, 0x33, 0x15, 0x32, 0x01, 0x00, 0x00,
}
