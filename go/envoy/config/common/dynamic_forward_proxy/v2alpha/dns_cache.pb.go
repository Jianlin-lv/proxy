// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/common/dynamic_forward_proxy/v2alpha/dns_cache.proto

package envoy_config_common_dynamic_forward_proxy_v2alpha

import (
	fmt "fmt"
	v2 "github.com/cilium/proxy/go/envoy/api/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Configuration for the dynamic forward proxy DNS cache. See the :ref:`architecture overview
// <arch_overview_http_dynamic_forward_proxy>` for more information.
// [#next-free-field: 6]
type DnsCacheConfig struct {
	// The name of the cache. Multiple named caches allow independent dynamic forward proxy
	// configurations to operate within a single Envoy process using different configurations. All
	// configurations with the same name *must* otherwise have the same settings when referenced
	// from different configuration components. Configuration will fail to load if this is not
	// the case.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The DNS lookup family to use during resolution.
	//
	// [#comment:TODO(mattklein123): Figure out how to support IPv4/IPv6 "happy eyeballs" mode. The
	// way this might work is a new lookup family which returns both IPv4 and IPv6 addresses, and
	// then configures a host to have a primary and fall back address. With this, we could very
	// likely build a "happy eyeballs" connection pool which would race the primary / fall back
	// address and return the one that wins. This same method could potentially also be used for
	// QUIC to TCP fall back.]
	DnsLookupFamily v2.Cluster_DnsLookupFamily `protobuf:"varint,2,opt,name=dns_lookup_family,json=dnsLookupFamily,proto3,enum=envoy.api.v2.Cluster_DnsLookupFamily" json:"dns_lookup_family,omitempty"`
	// The DNS refresh rate for currently cached DNS hosts. If not specified defaults to 60s.
	//
	// .. note:
	//
	//  The returned DNS TTL is not currently used to alter the refresh rate. This feature will be
	//  added in a future change.
	DnsRefreshRate *duration.Duration `protobuf:"bytes,3,opt,name=dns_refresh_rate,json=dnsRefreshRate,proto3" json:"dns_refresh_rate,omitempty"`
	// The TTL for hosts that are unused. Hosts that have not been used in the configured time
	// interval will be purged. If not specified defaults to 5m.
	//
	// .. note:
	//
	//   The TTL is only checked at the time of DNS refresh, as specified by *dns_refresh_rate*. This
	//   means that if the configured TTL is shorter than the refresh rate the host may not be removed
	//   immediately.
	//
	//  .. note:
	//
	//   The TTL has no relation to DNS TTL and is only used to control Envoy's resource usage.
	HostTtl *duration.Duration `protobuf:"bytes,4,opt,name=host_ttl,json=hostTtl,proto3" json:"host_ttl,omitempty"`
	// The maximum number of hosts that the cache will hold. If not specified defaults to 1024.
	//
	// .. note:
	//
	//   The implementation is approximate and enforced independently on each worker thread, thus
	//   it is possible for the maximum hosts in the cache to go slightly above the configured
	//   value depending on timing. This is similar to how other circuit breakers work.
	MaxHosts             *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=max_hosts,json=maxHosts,proto3" json:"max_hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DnsCacheConfig) Reset()         { *m = DnsCacheConfig{} }
func (m *DnsCacheConfig) String() string { return proto.CompactTextString(m) }
func (*DnsCacheConfig) ProtoMessage()    {}
func (*DnsCacheConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2d9297e0c94cb56, []int{0}
}

func (m *DnsCacheConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsCacheConfig.Unmarshal(m, b)
}
func (m *DnsCacheConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsCacheConfig.Marshal(b, m, deterministic)
}
func (m *DnsCacheConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsCacheConfig.Merge(m, src)
}
func (m *DnsCacheConfig) XXX_Size() int {
	return xxx_messageInfo_DnsCacheConfig.Size(m)
}
func (m *DnsCacheConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsCacheConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DnsCacheConfig proto.InternalMessageInfo

func (m *DnsCacheConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DnsCacheConfig) GetDnsLookupFamily() v2.Cluster_DnsLookupFamily {
	if m != nil {
		return m.DnsLookupFamily
	}
	return v2.Cluster_AUTO
}

func (m *DnsCacheConfig) GetDnsRefreshRate() *duration.Duration {
	if m != nil {
		return m.DnsRefreshRate
	}
	return nil
}

func (m *DnsCacheConfig) GetHostTtl() *duration.Duration {
	if m != nil {
		return m.HostTtl
	}
	return nil
}

func (m *DnsCacheConfig) GetMaxHosts() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxHosts
	}
	return nil
}

func init() {
	proto.RegisterType((*DnsCacheConfig)(nil), "envoy.config.common.dynamic_forward_proxy.v2alpha.DnsCacheConfig")
}

func init() {
	proto.RegisterFile("envoy/config/common/dynamic_forward_proxy/v2alpha/dns_cache.proto", fileDescriptor_d2d9297e0c94cb56)
}

var fileDescriptor_d2d9297e0c94cb56 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0xeb, 0xb0, 0xa5, 0x5b, 0x23, 0x96, 0x92, 0x03, 0x84, 0x82, 0x50, 0x84, 0x84, 0xb4,
	0xea, 0xc1, 0x16, 0xe9, 0x1d, 0x44, 0x76, 0x85, 0x40, 0x02, 0xa9, 0x8a, 0x80, 0x0b, 0x87, 0x68,
	0x1a, 0x3b, 0xbb, 0x16, 0x8e, 0x6d, 0xd9, 0x4e, 0xba, 0x7b, 0xe5, 0x71, 0x78, 0x27, 0x5e, 0xa4,
	0x27, 0x14, 0xbb, 0x39, 0x20, 0x40, 0x82, 0x5b, 0x26, 0x33, 0xff, 0x37, 0xf3, 0xcf, 0x18, 0xbf,
	0xe6, 0x6a, 0xd0, 0x7b, 0xda, 0x68, 0xd5, 0x8a, 0x0d, 0x6d, 0x74, 0xd7, 0x69, 0x45, 0xd9, 0x5e,
	0x41, 0x27, 0x9a, 0xba, 0xd5, 0xf6, 0x0a, 0x2c, 0xab, 0x8d, 0xd5, 0xbb, 0x3d, 0x1d, 0x0a, 0x90,
	0x66, 0x0b, 0x94, 0x29, 0x57, 0x37, 0xd0, 0x6c, 0x39, 0x31, 0x56, 0x7b, 0x9d, 0xbe, 0x08, 0x08,
	0x12, 0x11, 0x24, 0x22, 0xc8, 0x1f, 0x11, 0xe4, 0x06, 0x71, 0xfa, 0x20, 0x76, 0x05, 0x23, 0xe8,
	0x50, 0xd0, 0x86, 0xb9, 0x88, 0x3a, 0x7d, 0xba, 0xd1, 0x7a, 0x23, 0x39, 0x0d, 0xd1, 0x65, 0xdf,
	0x52, 0xd6, 0x5b, 0xf0, 0x42, 0xab, 0xbf, 0xe5, 0xaf, 0x2c, 0x18, 0xc3, 0xed, 0xa4, 0x7f, 0x38,
	0x80, 0x14, 0x0c, 0x3c, 0xa7, 0xd3, 0x47, 0x4c, 0x3c, 0xfb, 0x91, 0xe0, 0xc5, 0x5a, 0xb9, 0xd5,
	0x38, 0xf6, 0x2a, 0x0c, 0x9a, 0x3e, 0xc6, 0x33, 0x05, 0x1d, 0xcf, 0x50, 0x8e, 0x96, 0xc7, 0xe5,
	0xd1, 0x75, 0x39, 0xb3, 0x49, 0x8e, 0xaa, 0xf0, 0x33, 0xfd, 0x82, 0xef, 0x8f, 0x36, 0xa5, 0xd6,
	0x5f, 0x7b, 0x53, 0xb7, 0xd0, 0x09, 0xb9, 0xcf, 0x92, 0x1c, 0x2d, 0x17, 0xc5, 0x73, 0x12, 0xfd,
	0x82, 0x11, 0x64, 0x28, 0xc8, 0x4a, 0xf6, 0xce, 0x73, 0x4b, 0xd6, 0xca, 0xbd, 0x0f, 0xd5, 0x6f,
	0x42, 0x71, 0x39, 0xbf, 0x2e, 0x0f, 0xbf, 0xa1, 0xe4, 0x04, 0x55, 0xf7, 0xd8, 0xaf, 0xa9, 0xf4,
	0x03, 0x3e, 0x19, 0xe1, 0x96, 0xb7, 0x96, 0xbb, 0x6d, 0x6d, 0xc1, 0xf3, 0xec, 0x56, 0x8e, 0x96,
	0x77, 0x8a, 0x47, 0x24, 0x1a, 0x24, 0x93, 0x41, 0xb2, 0xbe, 0x59, 0x40, 0xe0, 0x7d, 0x47, 0xc9,
	0xd9, 0x41, 0xb5, 0x60, 0xca, 0x55, 0x51, 0x5b, 0x81, 0xe7, 0xe9, 0x4b, 0x3c, 0xdf, 0x6a, 0xe7,
	0x6b, 0xef, 0x65, 0x36, 0xfb, 0x77, 0xcc, 0xd1, 0x28, 0xfa, 0xe8, 0x65, 0x5a, 0xe2, 0xe3, 0x0e,
	0x76, 0xf5, 0x18, 0xba, 0xec, 0x30, 0x00, 0x9e, 0xfc, 0x06, 0xf8, 0xf4, 0x4e, 0xf9, 0xf3, 0xe2,
	0x33, 0xc8, 0x9e, 0x87, 0x5d, 0x9d, 0x25, 0xf9, 0x41, 0x35, 0xef, 0x60, 0xf7, 0x76, 0x94, 0x95,
	0x15, 0x7e, 0x25, 0x74, 0x5c, 0x4c, 0xbc, 0xf4, 0x7f, 0xbf, 0x89, 0xf2, 0xee, 0x74, 0x9f, 0x8b,
	0xb1, 0xe7, 0x05, 0xba, 0xbc, 0x1d, 0x9a, 0x9f, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xfa,
	0xa5, 0x1e, 0xa3, 0x02, 0x00, 0x00,
}
