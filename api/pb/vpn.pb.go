// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vpn.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VPNProto int32

const (
	VPNProto_NOPREF VPNProto = 0
	VPNProto_UDP    VPNProto = 1
	VPNProto_TCP    VPNProto = 2
)

var VPNProto_name = map[int32]string{
	0: "NOPREF",
	1: "UDP",
	2: "TCP",
}
var VPNProto_value = map[string]int32{
	"NOPREF": 0,
	"UDP":    1,
	"TCP":    2,
}

func (x VPNProto) String() string {
	return proto.EnumName(VPNProto_name, int32(x))
}
func (VPNProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_vpn_a36938fb10df2565, []int{0}
}

type VPNStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VPNStatusRequest) Reset()         { *m = VPNStatusRequest{} }
func (m *VPNStatusRequest) String() string { return proto.CompactTextString(m) }
func (*VPNStatusRequest) ProtoMessage()    {}
func (*VPNStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_vpn_a36938fb10df2565, []int{0}
}
func (m *VPNStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNStatusRequest.Unmarshal(m, b)
}
func (m *VPNStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNStatusRequest.Marshal(b, m, deterministic)
}
func (dst *VPNStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNStatusRequest.Merge(dst, src)
}
func (m *VPNStatusRequest) XXX_Size() int {
	return xxx_messageInfo_VPNStatusRequest.Size(m)
}
func (m *VPNStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VPNStatusRequest proto.InternalMessageInfo

type VPNInitRequest struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Port                 string   `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
	ProtoPref            VPNProto `protobuf:"varint,3,opt,name=proto_pref,json=protoPref,enum=pb.VPNProto" json:"proto_pref,omitempty"`
	IpBlock              string   `protobuf:"bytes,4,opt,name=ip_block,json=ipBlock" json:"ip_block,omitempty"`
	Dns                  string   `protobuf:"bytes,5,opt,name=dns" json:"dns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VPNInitRequest) Reset()         { *m = VPNInitRequest{} }
func (m *VPNInitRequest) String() string { return proto.CompactTextString(m) }
func (*VPNInitRequest) ProtoMessage()    {}
func (*VPNInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_vpn_a36938fb10df2565, []int{1}
}
func (m *VPNInitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNInitRequest.Unmarshal(m, b)
}
func (m *VPNInitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNInitRequest.Marshal(b, m, deterministic)
}
func (dst *VPNInitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNInitRequest.Merge(dst, src)
}
func (m *VPNInitRequest) XXX_Size() int {
	return xxx_messageInfo_VPNInitRequest.Size(m)
}
func (m *VPNInitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNInitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VPNInitRequest proto.InternalMessageInfo

func (m *VPNInitRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *VPNInitRequest) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *VPNInitRequest) GetProtoPref() VPNProto {
	if m != nil {
		return m.ProtoPref
	}
	return VPNProto_NOPREF
}

func (m *VPNInitRequest) GetIpBlock() string {
	if m != nil {
		return m.IpBlock
	}
	return ""
}

func (m *VPNInitRequest) GetDns() string {
	if m != nil {
		return m.Dns
	}
	return ""
}

type VPNUpdateRequest struct {
	IpBlock              string   `protobuf:"bytes,1,opt,name=ip_block,json=ipBlock" json:"ip_block,omitempty"`
	Dns                  string   `protobuf:"bytes,2,opt,name=dns" json:"dns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VPNUpdateRequest) Reset()         { *m = VPNUpdateRequest{} }
func (m *VPNUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*VPNUpdateRequest) ProtoMessage()    {}
func (*VPNUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_vpn_a36938fb10df2565, []int{2}
}
func (m *VPNUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNUpdateRequest.Unmarshal(m, b)
}
func (m *VPNUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNUpdateRequest.Marshal(b, m, deterministic)
}
func (dst *VPNUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNUpdateRequest.Merge(dst, src)
}
func (m *VPNUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_VPNUpdateRequest.Size(m)
}
func (m *VPNUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VPNUpdateRequest proto.InternalMessageInfo

func (m *VPNUpdateRequest) GetIpBlock() string {
	if m != nil {
		return m.IpBlock
	}
	return ""
}

func (m *VPNUpdateRequest) GetDns() string {
	if m != nil {
		return m.Dns
	}
	return ""
}

type VPNRestartRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VPNRestartRequest) Reset()         { *m = VPNRestartRequest{} }
func (m *VPNRestartRequest) String() string { return proto.CompactTextString(m) }
func (*VPNRestartRequest) ProtoMessage()    {}
func (*VPNRestartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_vpn_a36938fb10df2565, []int{3}
}
func (m *VPNRestartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNRestartRequest.Unmarshal(m, b)
}
func (m *VPNRestartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNRestartRequest.Marshal(b, m, deterministic)
}
func (dst *VPNRestartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNRestartRequest.Merge(dst, src)
}
func (m *VPNRestartRequest) XXX_Size() int {
	return xxx_messageInfo_VPNRestartRequest.Size(m)
}
func (m *VPNRestartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNRestartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VPNRestartRequest proto.InternalMessageInfo

type VPNStatusResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	SerialNumber         string   `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
	Hostname             string   `protobuf:"bytes,3,opt,name=hostname" json:"hostname,omitempty"`
	Port                 string   `protobuf:"bytes,4,opt,name=port" json:"port,omitempty"`
	Cert                 string   `protobuf:"bytes,5,opt,name=cert" json:"cert,omitempty"`
	CaCert               string   `protobuf:"bytes,6,opt,name=ca_cert,json=caCert" json:"ca_cert,omitempty"`
	Net                  string   `protobuf:"bytes,7,opt,name=net" json:"net,omitempty"`
	Mask                 string   `protobuf:"bytes,8,opt,name=mask" json:"mask,omitempty"`
	CreatedAt            string   `protobuf:"bytes,9,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Proto                string   `protobuf:"bytes,10,opt,name=proto" json:"proto,omitempty"`
	Dns                  string   `protobuf:"bytes,11,opt,name=dns" json:"dns,omitempty"`
	ExpiresAt            string   `protobuf:"bytes,12,opt,name=expires_at,json=expiresAt" json:"expires_at,omitempty"`
	CaExpiresAt          string   `protobuf:"bytes,13,opt,name=ca_expires_at,json=caExpiresAt" json:"ca_expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VPNStatusResponse) Reset()         { *m = VPNStatusResponse{} }
func (m *VPNStatusResponse) String() string { return proto.CompactTextString(m) }
func (*VPNStatusResponse) ProtoMessage()    {}
func (*VPNStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_vpn_a36938fb10df2565, []int{4}
}
func (m *VPNStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNStatusResponse.Unmarshal(m, b)
}
func (m *VPNStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNStatusResponse.Marshal(b, m, deterministic)
}
func (dst *VPNStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNStatusResponse.Merge(dst, src)
}
func (m *VPNStatusResponse) XXX_Size() int {
	return xxx_messageInfo_VPNStatusResponse.Size(m)
}
func (m *VPNStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VPNStatusResponse proto.InternalMessageInfo

func (m *VPNStatusResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VPNStatusResponse) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *VPNStatusResponse) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *VPNStatusResponse) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *VPNStatusResponse) GetCert() string {
	if m != nil {
		return m.Cert
	}
	return ""
}

func (m *VPNStatusResponse) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

func (m *VPNStatusResponse) GetNet() string {
	if m != nil {
		return m.Net
	}
	return ""
}

func (m *VPNStatusResponse) GetMask() string {
	if m != nil {
		return m.Mask
	}
	return ""
}

func (m *VPNStatusResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *VPNStatusResponse) GetProto() string {
	if m != nil {
		return m.Proto
	}
	return ""
}

func (m *VPNStatusResponse) GetDns() string {
	if m != nil {
		return m.Dns
	}
	return ""
}

func (m *VPNStatusResponse) GetExpiresAt() string {
	if m != nil {
		return m.ExpiresAt
	}
	return ""
}

func (m *VPNStatusResponse) GetCaExpiresAt() string {
	if m != nil {
		return m.CaExpiresAt
	}
	return ""
}

type VPNInitResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VPNInitResponse) Reset()         { *m = VPNInitResponse{} }
func (m *VPNInitResponse) String() string { return proto.CompactTextString(m) }
func (*VPNInitResponse) ProtoMessage()    {}
func (*VPNInitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_vpn_a36938fb10df2565, []int{5}
}
func (m *VPNInitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNInitResponse.Unmarshal(m, b)
}
func (m *VPNInitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNInitResponse.Marshal(b, m, deterministic)
}
func (dst *VPNInitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNInitResponse.Merge(dst, src)
}
func (m *VPNInitResponse) XXX_Size() int {
	return xxx_messageInfo_VPNInitResponse.Size(m)
}
func (m *VPNInitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNInitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VPNInitResponse proto.InternalMessageInfo

type VPNUpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VPNUpdateResponse) Reset()         { *m = VPNUpdateResponse{} }
func (m *VPNUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*VPNUpdateResponse) ProtoMessage()    {}
func (*VPNUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_vpn_a36938fb10df2565, []int{6}
}
func (m *VPNUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNUpdateResponse.Unmarshal(m, b)
}
func (m *VPNUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNUpdateResponse.Marshal(b, m, deterministic)
}
func (dst *VPNUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNUpdateResponse.Merge(dst, src)
}
func (m *VPNUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_VPNUpdateResponse.Size(m)
}
func (m *VPNUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VPNUpdateResponse proto.InternalMessageInfo

type VPNRestartResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VPNRestartResponse) Reset()         { *m = VPNRestartResponse{} }
func (m *VPNRestartResponse) String() string { return proto.CompactTextString(m) }
func (*VPNRestartResponse) ProtoMessage()    {}
func (*VPNRestartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_vpn_a36938fb10df2565, []int{7}
}
func (m *VPNRestartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VPNRestartResponse.Unmarshal(m, b)
}
func (m *VPNRestartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VPNRestartResponse.Marshal(b, m, deterministic)
}
func (dst *VPNRestartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VPNRestartResponse.Merge(dst, src)
}
func (m *VPNRestartResponse) XXX_Size() int {
	return xxx_messageInfo_VPNRestartResponse.Size(m)
}
func (m *VPNRestartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VPNRestartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VPNRestartResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VPNStatusRequest)(nil), "pb.VPNStatusRequest")
	proto.RegisterType((*VPNInitRequest)(nil), "pb.VPNInitRequest")
	proto.RegisterType((*VPNUpdateRequest)(nil), "pb.VPNUpdateRequest")
	proto.RegisterType((*VPNRestartRequest)(nil), "pb.VPNRestartRequest")
	proto.RegisterType((*VPNStatusResponse)(nil), "pb.VPNStatusResponse")
	proto.RegisterType((*VPNInitResponse)(nil), "pb.VPNInitResponse")
	proto.RegisterType((*VPNUpdateResponse)(nil), "pb.VPNUpdateResponse")
	proto.RegisterType((*VPNRestartResponse)(nil), "pb.VPNRestartResponse")
	proto.RegisterEnum("pb.VPNProto", VPNProto_name, VPNProto_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VPNService service

type VPNServiceClient interface {
	Status(ctx context.Context, in *VPNStatusRequest, opts ...grpc.CallOption) (*VPNStatusResponse, error)
	Init(ctx context.Context, in *VPNInitRequest, opts ...grpc.CallOption) (*VPNInitResponse, error)
	Update(ctx context.Context, in *VPNUpdateRequest, opts ...grpc.CallOption) (*VPNUpdateResponse, error)
	Restart(ctx context.Context, in *VPNRestartRequest, opts ...grpc.CallOption) (*VPNRestartResponse, error)
}

type vPNServiceClient struct {
	cc *grpc.ClientConn
}

func NewVPNServiceClient(cc *grpc.ClientConn) VPNServiceClient {
	return &vPNServiceClient{cc}
}

func (c *vPNServiceClient) Status(ctx context.Context, in *VPNStatusRequest, opts ...grpc.CallOption) (*VPNStatusResponse, error) {
	out := new(VPNStatusResponse)
	err := grpc.Invoke(ctx, "/pb.VPNService/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vPNServiceClient) Init(ctx context.Context, in *VPNInitRequest, opts ...grpc.CallOption) (*VPNInitResponse, error) {
	out := new(VPNInitResponse)
	err := grpc.Invoke(ctx, "/pb.VPNService/Init", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vPNServiceClient) Update(ctx context.Context, in *VPNUpdateRequest, opts ...grpc.CallOption) (*VPNUpdateResponse, error) {
	out := new(VPNUpdateResponse)
	err := grpc.Invoke(ctx, "/pb.VPNService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vPNServiceClient) Restart(ctx context.Context, in *VPNRestartRequest, opts ...grpc.CallOption) (*VPNRestartResponse, error) {
	out := new(VPNRestartResponse)
	err := grpc.Invoke(ctx, "/pb.VPNService/Restart", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VPNService service

type VPNServiceServer interface {
	Status(context.Context, *VPNStatusRequest) (*VPNStatusResponse, error)
	Init(context.Context, *VPNInitRequest) (*VPNInitResponse, error)
	Update(context.Context, *VPNUpdateRequest) (*VPNUpdateResponse, error)
	Restart(context.Context, *VPNRestartRequest) (*VPNRestartResponse, error)
}

func RegisterVPNServiceServer(s *grpc.Server, srv VPNServiceServer) {
	s.RegisterService(&_VPNService_serviceDesc, srv)
}

func _VPNService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VPNStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPNServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VPNService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPNServiceServer).Status(ctx, req.(*VPNStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VPNService_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VPNInitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPNServiceServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VPNService/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPNServiceServer).Init(ctx, req.(*VPNInitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VPNService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VPNUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPNServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VPNService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPNServiceServer).Update(ctx, req.(*VPNUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VPNService_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VPNRestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPNServiceServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VPNService/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPNServiceServer).Restart(ctx, req.(*VPNRestartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VPNService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.VPNService",
	HandlerType: (*VPNServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _VPNService_Status_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _VPNService_Init_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _VPNService_Update_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _VPNService_Restart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vpn.proto",
}

func init() { proto.RegisterFile("vpn.proto", fileDescriptor_vpn_a36938fb10df2565) }

var fileDescriptor_vpn_a36938fb10df2565 = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xb1, 0x93, 0xe6, 0x63, 0xfa, 0x81, 0x3b, 0x4d, 0xa9, 0x09, 0x54, 0xaa, 0xcc, 0xa5,
	0x2a, 0x52, 0x23, 0xca, 0x8d, 0x0b, 0x2a, 0xa5, 0x48, 0x48, 0xc8, 0x98, 0xd0, 0xe6, 0x6a, 0x6d,
	0xdc, 0x6d, 0xb1, 0xda, 0xac, 0x97, 0xdd, 0x4d, 0xc4, 0x99, 0x57, 0xe0, 0xc6, 0x6b, 0x71, 0x82,
	0x33, 0x0f, 0x82, 0x3c, 0xbb, 0x0e, 0x76, 0x04, 0xb7, 0xf1, 0x7f, 0x66, 0x7f, 0x9a, 0xf9, 0xcf,
	0x18, 0xfa, 0x0b, 0x29, 0x8e, 0xa5, 0x2a, 0x4c, 0x81, 0xbe, 0x9c, 0x0e, 0x1f, 0xdf, 0x14, 0xc5,
	0xcd, 0x1d, 0x1f, 0x31, 0x99, 0x8f, 0x98, 0x10, 0x85, 0x61, 0x26, 0x2f, 0x84, 0xb6, 0x15, 0x11,
	0x42, 0x30, 0x49, 0xe2, 0x8f, 0x86, 0x99, 0xb9, 0x1e, 0xf3, 0xcf, 0x73, 0xae, 0x4d, 0xf4, 0xdd,
	0x83, 0xad, 0x49, 0x12, 0xbf, 0x15, 0xb9, 0x71, 0x12, 0x0e, 0xa1, 0xf7, 0xa9, 0xd0, 0x46, 0xb0,
	0x19, 0x0f, 0xbd, 0x03, 0xef, 0xb0, 0x3f, 0x5e, 0x7e, 0x23, 0x42, 0x5b, 0x16, 0xca, 0x84, 0x3e,
	0xe9, 0x14, 0xe3, 0x53, 0x00, 0xe2, 0xa7, 0x52, 0xf1, 0xeb, 0xb0, 0x75, 0xe0, 0x1d, 0x6e, 0x9d,
	0x6c, 0x1c, 0xcb, 0xe9, 0xf1, 0x24, 0x89, 0x93, 0x32, 0x31, 0xee, 0x53, 0x3e, 0x51, 0xfc, 0x1a,
	0x1f, 0x42, 0x2f, 0x97, 0xe9, 0xf4, 0xae, 0xc8, 0x6e, 0xc3, 0x36, 0x41, 0xba, 0xb9, 0x7c, 0x55,
	0x7e, 0x62, 0x00, 0xad, 0x2b, 0xa1, 0xc3, 0x35, 0x52, 0xcb, 0x30, 0x7a, 0x49, 0x0d, 0x5f, 0xca,
	0x2b, 0x66, 0x78, 0xd5, 0x5d, 0x1d, 0xe0, 0xfd, 0x13, 0xe0, 0xff, 0x05, 0xec, 0xc0, 0xf6, 0x24,
	0x89, 0xc7, 0x5c, 0x1b, 0xa6, 0xaa, 0xf9, 0xa2, 0x9f, 0x3e, 0xa9, 0x95, 0x0f, 0x5a, 0x16, 0x42,
	0xd3, 0x64, 0xb5, 0x89, 0x29, 0xc6, 0x27, 0xb0, 0xa9, 0xb9, 0xca, 0xd9, 0x5d, 0x2a, 0xe6, 0xb3,
	0x29, 0x57, 0x0e, 0xbd, 0x61, 0xc5, 0x98, 0xb4, 0x86, 0x5d, 0xad, 0xff, 0xd8, 0xd5, 0xae, 0xd9,
	0x85, 0xd0, 0xce, 0xb8, 0x32, 0x6e, 0x4e, 0x8a, 0x71, 0x0f, 0xba, 0x19, 0x4b, 0x49, 0xee, 0x90,
	0xdc, 0xc9, 0xd8, 0x59, 0x99, 0x08, 0xa0, 0x25, 0xb8, 0x09, 0xbb, 0x76, 0x24, 0xc1, 0xe9, 0xf9,
	0x8c, 0xe9, 0xdb, 0xb0, 0x67, 0x9f, 0x97, 0x31, 0xee, 0x03, 0x64, 0x8a, 0x33, 0xc3, 0xaf, 0x52,
	0x66, 0xc2, 0x3e, 0x65, 0xfa, 0x4e, 0x39, 0x35, 0x38, 0x80, 0x35, 0x5a, 0x40, 0x08, 0x94, 0xb1,
	0x1f, 0x95, 0x5b, 0xeb, 0x4b, 0xb7, 0x4a, 0x0c, 0xff, 0x22, 0x73, 0xc5, 0x75, 0x89, 0xd9, 0xb0,
	0x18, 0xa7, 0x9c, 0x1a, 0x8c, 0x60, 0x33, 0x63, 0x69, 0xad, 0x62, 0x93, 0x2a, 0xd6, 0x33, 0x76,
	0x5e, 0xd5, 0x44, 0xdb, 0x70, 0x7f, 0x79, 0x4d, 0xd6, 0x58, 0xb7, 0x83, 0x6a, 0x89, 0x4e, 0x1c,
	0x00, 0xd6, 0x17, 0x63, 0xd5, 0xa3, 0x43, 0xe8, 0x55, 0x37, 0x83, 0x00, 0x9d, 0xf8, 0x7d, 0x32,
	0x3e, 0x7f, 0x13, 0xdc, 0xc3, 0x2e, 0xb4, 0x2e, 0x5f, 0x27, 0x81, 0x57, 0x06, 0x17, 0x67, 0x49,
	0xe0, 0x9f, 0xfc, 0xf2, 0x01, 0xca, 0x1d, 0x72, 0xb5, 0xc8, 0x33, 0x8e, 0x1f, 0xa0, 0x63, 0xd7,
	0x89, 0x03, 0x77, 0x78, 0x8d, 0x2b, 0x1f, 0xee, 0xae, 0xa8, 0xae, 0x8b, 0xe1, 0xd7, 0x1f, 0xbf,
	0xbf, 0xf9, 0x03, 0x44, 0xfa, 0x61, 0x16, 0xcf, 0x46, 0x0b, 0x29, 0x46, 0xda, 0x82, 0xde, 0x41,
	0xbb, 0x1c, 0x03, 0xd1, 0x3d, 0xad, 0xfd, 0x21, 0xc3, 0x9d, 0x86, 0xe6, 0x60, 0x8f, 0x08, 0xb6,
	0x1b, 0x05, 0x75, 0x58, 0x2e, 0x72, 0xf3, 0xc2, 0x3b, 0xc2, 0x0b, 0xe8, 0x58, 0x07, 0x96, 0x0d,
	0x36, 0xae, 0x7a, 0xd9, 0xe0, 0x8a, 0x4d, 0xfb, 0xc4, 0xdc, 0x8b, 0x1a, 0x0d, 0xce, 0xa9, 0xa6,
	0xa4, 0x5e, 0x42, 0xd7, 0x59, 0x88, 0x15, 0xa0, 0x79, 0xeb, 0xc3, 0x07, 0xab, 0xf2, 0x4a, 0xb3,
	0x3b, 0x75, 0xb0, 0xb2, 0x45, 0xd3, 0x0e, 0x1d, 0xc8, 0xf3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xd4, 0x74, 0x98, 0x4b, 0x5d, 0x04, 0x00, 0x00,
}
