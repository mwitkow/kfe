// Code generated by protoc-gen-go.
// source: kedge/config/http/adhoc/matcher.proto
// DO NOT EDIT!

/*
Package kedge_config_http_routes is a generated protocol buffer package.

It is generated from these files:
	kedge/config/http/adhoc/matcher.proto

It has these top-level messages:
	Route
*/
package kedge_config_http_routes

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

type ProxyMode int32

const (
	ProxyMode_ANY ProxyMode = 0
	// / Reverse Proxy is when the FE serves an authority (Host) publicly and clients connect to that authority
	// / directly. This is used to expose publicly DNS-resolvable names.
	ProxyMode_REVERSE_PROXY ProxyMode = 1
	// / Forward Proxy is when the FE serves as an HTTP_PROXY for a browser or an application. The resolution of the
	// / backend is done by the FE itself, so non-public names can be addressed.
	// / This may be from the 90s, but it still is very useful.
	// /
	// / IMPORTANT: If you have a PAC file configured in Firefox, the HTTPS rule behaves differently than in Chrome. The
	// / proxied requests are not FORWARD_PROXY requests but REVERSE_PROXY_REQUESTS.
	ProxyMode_FORWARD_PROXY ProxyMode = 2
)

var ProxyMode_name = map[int32]string{
	0: "ANY",
	1: "REVERSE_PROXY",
	2: "FORWARD_PROXY",
}
var ProxyMode_value = map[string]int32{
	"ANY":           0,
	"REVERSE_PROXY": 1,
	"FORWARD_PROXY": 2,
}

func (x ProxyMode) String() string {
	return proto.EnumName(ProxyMode_name, int32(x))
}
func (ProxyMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// / Backend is a gRPC ClientConn pool maintained to a single serivce.
type Route struct {
	// / backend_name is the string identifying the HTTP backend pool to send data to.
	BackendName string `protobuf:"bytes,1,opt,name=backend_name,json=backendName" json:"backend_name,omitempty"`
	// / path_rules is a globbing expression that matches a URL path of the request.
	// / See: https://cloud.google.com/compute/docs/load-balancing/http/url-map
	// / If not present, '/*' is default.
	PathRules []string `protobuf:"bytes,2,rep,name=path_rules,json=pathRules" json:"path_rules,omitempty"`
	// / host_matcher matches on the ':authority' header (a.k.a. Host header) enabling Virtual Host-like proxying.
	// / The matching is done through lower-case string-equality.
	// / If none are present, the route skips ':authority' checks.
	HostMatcher string `protobuf:"bytes,3,opt,name=host_matcher,json=hostMatcher" json:"host_matcher,omitempty"`
	// / metadata_matcher matches any HTTP inbound request Headers.
	// / Eeach key provided must find a match for the route to match.
	// / The matching is done through lower-case key match and explicit string-equality of values.
	// / If none are present, the route skips metadata checks.
	HeaderMatcher map[string]string `protobuf:"bytes,4,rep,name=header_matcher,json=headerMatcher" json:"header_matcher,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// / proxy_mode controlls what kind of inbound requests this route matches. See
	ProxyMode ProxyMode `protobuf:"varint,5,opt,name=proxy_mode,json=proxyMode,enum=kedge.config.http.routes.ProxyMode" json:"proxy_mode,omitempty"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Route) GetBackendName() string {
	if m != nil {
		return m.BackendName
	}
	return ""
}

func (m *Route) GetPathRules() []string {
	if m != nil {
		return m.PathRules
	}
	return nil
}

func (m *Route) GetHostMatcher() string {
	if m != nil {
		return m.HostMatcher
	}
	return ""
}

func (m *Route) GetHeaderMatcher() map[string]string {
	if m != nil {
		return m.HeaderMatcher
	}
	return nil
}

func (m *Route) GetProxyMode() ProxyMode {
	if m != nil {
		return m.ProxyMode
	}
	return ProxyMode_ANY
}

func init() {
	proto.RegisterType((*Route)(nil), "kedge.config.http.routes.Route")
	proto.RegisterEnum("kedge.config.http.routes.ProxyMode", ProxyMode_name, ProxyMode_value)
}

func init() { proto.RegisterFile("kedge/config/http/adhoc/routes.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0x4d, 0x4b, 0xc3, 0x30,
	0x18, 0xc7, 0x6d, 0xeb, 0x94, 0x3e, 0x73, 0x63, 0x06, 0x0f, 0x45, 0x10, 0xea, 0xcb, 0xa1, 0x78,
	0x48, 0x61, 0x5e, 0x64, 0x27, 0x27, 0x56, 0xbc, 0xec, 0x85, 0x08, 0xea, 0x4e, 0x25, 0x6b, 0x1f,
	0x17, 0xd9, 0xda, 0x94, 0x2c, 0x13, 0xfb, 0x21, 0xfd, 0x4e, 0x92, 0xb6, 0x0e, 0x41, 0x76, 0xfb,
	0xe7, 0x97, 0x7f, 0x7e, 0x79, 0x48, 0xe0, 0x6a, 0x89, 0xe9, 0x02, 0xc3, 0x44, 0xe6, 0xef, 0x1f,
	0x8b, 0x50, 0x68, 0x5d, 0x84, 0x3c, 0x15, 0x32, 0x09, 0x95, 0xdc, 0x68, 0x5c, 0xd3, 0x42, 0x49,
	0x2d, 0x89, 0x57, 0xb5, 0x68, 0xdd, 0xa2, 0xa6, 0x45, 0xeb, 0xfd, 0x8b, 0x6f, 0x1b, 0x5a, 0xcc,
	0x44, 0x72, 0x0e, 0x47, 0x73, 0x9e, 0x2c, 0x31, 0x4f, 0xe3, 0x9c, 0x67, 0xe8, 0x59, 0xbe, 0x15,
	0xb8, 0xac, 0xdd, 0xb0, 0x31, 0xcf, 0x90, 0x9c, 0x01, 0x14, 0x5c, 0x8b, 0x58, 0x6d, 0x56, 0xb8,
	0xf6, 0x6c, 0xdf, 0x09, 0x5c, 0xe6, 0x1a, 0xc2, 0x0c, 0x30, 0x06, 0x21, 0xd7, 0x3a, 0xce, 0xb8,
	0x4e, 0x04, 0x2a, 0xcf, 0xa9, 0x0d, 0x86, 0x8d, 0x6a, 0x44, 0x66, 0xd0, 0x15, 0xc8, 0x53, 0x54,
	0xdb, 0xd2, 0xbe, 0xef, 0x04, 0xed, 0x7e, 0x9f, 0xee, 0x9a, 0x90, 0x56, 0xd3, 0xd1, 0xa7, 0xea,
	0x54, 0xa3, 0x89, 0x72, 0xad, 0x4a, 0xd6, 0x11, 0x7f, 0x19, 0xb9, 0x07, 0x28, 0x94, 0xfc, 0x2a,
	0xe3, 0x4c, 0xa6, 0xe8, 0xb5, 0x7c, 0x2b, 0xe8, 0xf6, 0x2f, 0x77, 0x6b, 0xa7, 0xa6, 0x3b, 0x92,
	0x29, 0x32, 0xb7, 0xf8, 0x8d, 0xa7, 0x77, 0x40, 0xfe, 0x5f, 0x44, 0x7a, 0xe0, 0x2c, 0xb1, 0x6c,
	0x1e, 0xc4, 0x44, 0x72, 0x02, 0xad, 0x4f, 0xbe, 0xda, 0xa0, 0x67, 0x57, 0xac, 0x5e, 0x0c, 0xec,
	0x5b, 0xeb, 0x7a, 0x00, 0xee, 0xd6, 0x4c, 0x0e, 0xc1, 0x19, 0x8e, 0x67, 0xbd, 0x3d, 0x72, 0x0c,
	0x1d, 0x16, 0xbd, 0x44, 0xec, 0x39, 0x8a, 0xa7, 0x6c, 0xf2, 0x36, 0xeb, 0x59, 0x06, 0x3d, 0x4e,
	0xd8, 0xeb, 0x90, 0x3d, 0x34, 0xc8, 0x9e, 0x1f, 0x54, 0x9f, 0x75, 0xf3, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x4e, 0x74, 0x13, 0x0f, 0xd4, 0x01, 0x00, 0x00,
}
