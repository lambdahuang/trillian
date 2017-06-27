// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trillian_map_api.proto

package trillian

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

// MapLeaf represents the data behind Map leaves.
type MapLeaf struct {
	// index is the location of this leaf.
	// All indexes for a given Map must contain a constant number of bits.
	Index []byte `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	// leaf_hash is the tree hash of leaf_value.
	LeafHash []byte `protobuf:"bytes,2,opt,name=leaf_hash,json=leafHash,proto3" json:"leaf_hash,omitempty"`
	// leaf_value is the data the tree commits to.
	LeafValue []byte `protobuf:"bytes,3,opt,name=leaf_value,json=leafValue,proto3" json:"leaf_value,omitempty"`
	// extra_data holds related contextual data, but is not covered by any hash.
	ExtraData []byte `protobuf:"bytes,4,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
}

func (m *MapLeaf) Reset()                    { *m = MapLeaf{} }
func (m *MapLeaf) String() string            { return proto.CompactTextString(m) }
func (*MapLeaf) ProtoMessage()               {}
func (*MapLeaf) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *MapLeaf) GetIndex() []byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *MapLeaf) GetLeafHash() []byte {
	if m != nil {
		return m.LeafHash
	}
	return nil
}

func (m *MapLeaf) GetLeafValue() []byte {
	if m != nil {
		return m.LeafValue
	}
	return nil
}

func (m *MapLeaf) GetExtraData() []byte {
	if m != nil {
		return m.ExtraData
	}
	return nil
}

type MapLeafInclusion struct {
	Leaf      *MapLeaf `protobuf:"bytes,1,opt,name=leaf" json:"leaf,omitempty"`
	Inclusion [][]byte `protobuf:"bytes,2,rep,name=inclusion,proto3" json:"inclusion,omitempty"`
}

func (m *MapLeafInclusion) Reset()                    { *m = MapLeafInclusion{} }
func (m *MapLeafInclusion) String() string            { return proto.CompactTextString(m) }
func (*MapLeafInclusion) ProtoMessage()               {}
func (*MapLeafInclusion) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *MapLeafInclusion) GetLeaf() *MapLeaf {
	if m != nil {
		return m.Leaf
	}
	return nil
}

func (m *MapLeafInclusion) GetInclusion() [][]byte {
	if m != nil {
		return m.Inclusion
	}
	return nil
}

type GetMapLeavesRequest struct {
	MapId    int64    `protobuf:"varint,1,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
	Index    [][]byte `protobuf:"bytes,2,rep,name=index,proto3" json:"index,omitempty"`
	Revision int64    `protobuf:"varint,3,opt,name=revision" json:"revision,omitempty"`
}

func (m *GetMapLeavesRequest) Reset()                    { *m = GetMapLeavesRequest{} }
func (m *GetMapLeavesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetMapLeavesRequest) ProtoMessage()               {}
func (*GetMapLeavesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetMapLeavesRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *GetMapLeavesRequest) GetIndex() [][]byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *GetMapLeavesRequest) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

type GetMapLeavesResponse struct {
	MapLeafInclusion []*MapLeafInclusion `protobuf:"bytes,2,rep,name=map_leaf_inclusion,json=mapLeafInclusion" json:"map_leaf_inclusion,omitempty"`
	MapRoot          *SignedMapRoot      `protobuf:"bytes,3,opt,name=map_root,json=mapRoot" json:"map_root,omitempty"`
}

func (m *GetMapLeavesResponse) Reset()                    { *m = GetMapLeavesResponse{} }
func (m *GetMapLeavesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetMapLeavesResponse) ProtoMessage()               {}
func (*GetMapLeavesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GetMapLeavesResponse) GetMapLeafInclusion() []*MapLeafInclusion {
	if m != nil {
		return m.MapLeafInclusion
	}
	return nil
}

func (m *GetMapLeavesResponse) GetMapRoot() *SignedMapRoot {
	if m != nil {
		return m.MapRoot
	}
	return nil
}

type SetMapLeavesRequest struct {
	MapId      int64           `protobuf:"varint,1,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
	Leaves     []*MapLeaf      `protobuf:"bytes,2,rep,name=leaves" json:"leaves,omitempty"`
	MapperData *MapperMetadata `protobuf:"bytes,3,opt,name=mapper_data,json=mapperData" json:"mapper_data,omitempty"`
}

func (m *SetMapLeavesRequest) Reset()                    { *m = SetMapLeavesRequest{} }
func (m *SetMapLeavesRequest) String() string            { return proto.CompactTextString(m) }
func (*SetMapLeavesRequest) ProtoMessage()               {}
func (*SetMapLeavesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *SetMapLeavesRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *SetMapLeavesRequest) GetLeaves() []*MapLeaf {
	if m != nil {
		return m.Leaves
	}
	return nil
}

func (m *SetMapLeavesRequest) GetMapperData() *MapperMetadata {
	if m != nil {
		return m.MapperData
	}
	return nil
}

type SetMapLeavesResponse struct {
	MapRoot *SignedMapRoot `protobuf:"bytes,2,opt,name=map_root,json=mapRoot" json:"map_root,omitempty"`
}

func (m *SetMapLeavesResponse) Reset()                    { *m = SetMapLeavesResponse{} }
func (m *SetMapLeavesResponse) String() string            { return proto.CompactTextString(m) }
func (*SetMapLeavesResponse) ProtoMessage()               {}
func (*SetMapLeavesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *SetMapLeavesResponse) GetMapRoot() *SignedMapRoot {
	if m != nil {
		return m.MapRoot
	}
	return nil
}

type GetSignedMapRootRequest struct {
	MapId int64 `protobuf:"varint,1,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
}

func (m *GetSignedMapRootRequest) Reset()                    { *m = GetSignedMapRootRequest{} }
func (m *GetSignedMapRootRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSignedMapRootRequest) ProtoMessage()               {}
func (*GetSignedMapRootRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *GetSignedMapRootRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

type GetSignedMapRootByRevisionRequest struct {
	MapId    int64 `protobuf:"varint,1,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
	Revision int64 `protobuf:"varint,2,opt,name=revision" json:"revision,omitempty"`
}

func (m *GetSignedMapRootByRevisionRequest) Reset()         { *m = GetSignedMapRootByRevisionRequest{} }
func (m *GetSignedMapRootByRevisionRequest) String() string { return proto.CompactTextString(m) }
func (*GetSignedMapRootByRevisionRequest) ProtoMessage()    {}
func (*GetSignedMapRootByRevisionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{7}
}

func (m *GetSignedMapRootByRevisionRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *GetSignedMapRootByRevisionRequest) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

type GetSignedMapRootResponse struct {
	MapRoot *SignedMapRoot `protobuf:"bytes,2,opt,name=map_root,json=mapRoot" json:"map_root,omitempty"`
}

func (m *GetSignedMapRootResponse) Reset()                    { *m = GetSignedMapRootResponse{} }
func (m *GetSignedMapRootResponse) String() string            { return proto.CompactTextString(m) }
func (*GetSignedMapRootResponse) ProtoMessage()               {}
func (*GetSignedMapRootResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *GetSignedMapRootResponse) GetMapRoot() *SignedMapRoot {
	if m != nil {
		return m.MapRoot
	}
	return nil
}

func init() {
	proto.RegisterType((*MapLeaf)(nil), "trillian.MapLeaf")
	proto.RegisterType((*MapLeafInclusion)(nil), "trillian.MapLeafInclusion")
	proto.RegisterType((*GetMapLeavesRequest)(nil), "trillian.GetMapLeavesRequest")
	proto.RegisterType((*GetMapLeavesResponse)(nil), "trillian.GetMapLeavesResponse")
	proto.RegisterType((*SetMapLeavesRequest)(nil), "trillian.SetMapLeavesRequest")
	proto.RegisterType((*SetMapLeavesResponse)(nil), "trillian.SetMapLeavesResponse")
	proto.RegisterType((*GetSignedMapRootRequest)(nil), "trillian.GetSignedMapRootRequest")
	proto.RegisterType((*GetSignedMapRootByRevisionRequest)(nil), "trillian.GetSignedMapRootByRevisionRequest")
	proto.RegisterType((*GetSignedMapRootResponse)(nil), "trillian.GetSignedMapRootResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TrillianMap service

type TrillianMapClient interface {
	// GetLeaves returns an inclusion proof for each index requested.
	// For indexes that do not exist, the inclusion proof will use nil for the empty leaf value.
	GetLeaves(ctx context.Context, in *GetMapLeavesRequest, opts ...grpc.CallOption) (*GetMapLeavesResponse, error)
	SetLeaves(ctx context.Context, in *SetMapLeavesRequest, opts ...grpc.CallOption) (*SetMapLeavesResponse, error)
	GetSignedMapRoot(ctx context.Context, in *GetSignedMapRootRequest, opts ...grpc.CallOption) (*GetSignedMapRootResponse, error)
	GetSignedMapRootByRevision(ctx context.Context, in *GetSignedMapRootByRevisionRequest, opts ...grpc.CallOption) (*GetSignedMapRootResponse, error)
}

type trillianMapClient struct {
	cc *grpc.ClientConn
}

func NewTrillianMapClient(cc *grpc.ClientConn) TrillianMapClient {
	return &trillianMapClient{cc}
}

func (c *trillianMapClient) GetLeaves(ctx context.Context, in *GetMapLeavesRequest, opts ...grpc.CallOption) (*GetMapLeavesResponse, error) {
	out := new(GetMapLeavesResponse)
	err := grpc.Invoke(ctx, "/trillian.TrillianMap/GetLeaves", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianMapClient) SetLeaves(ctx context.Context, in *SetMapLeavesRequest, opts ...grpc.CallOption) (*SetMapLeavesResponse, error) {
	out := new(SetMapLeavesResponse)
	err := grpc.Invoke(ctx, "/trillian.TrillianMap/SetLeaves", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianMapClient) GetSignedMapRoot(ctx context.Context, in *GetSignedMapRootRequest, opts ...grpc.CallOption) (*GetSignedMapRootResponse, error) {
	out := new(GetSignedMapRootResponse)
	err := grpc.Invoke(ctx, "/trillian.TrillianMap/GetSignedMapRoot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianMapClient) GetSignedMapRootByRevision(ctx context.Context, in *GetSignedMapRootByRevisionRequest, opts ...grpc.CallOption) (*GetSignedMapRootResponse, error) {
	out := new(GetSignedMapRootResponse)
	err := grpc.Invoke(ctx, "/trillian.TrillianMap/GetSignedMapRootByRevision", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TrillianMap service

type TrillianMapServer interface {
	// GetLeaves returns an inclusion proof for each index requested.
	// For indexes that do not exist, the inclusion proof will use nil for the empty leaf value.
	GetLeaves(context.Context, *GetMapLeavesRequest) (*GetMapLeavesResponse, error)
	SetLeaves(context.Context, *SetMapLeavesRequest) (*SetMapLeavesResponse, error)
	GetSignedMapRoot(context.Context, *GetSignedMapRootRequest) (*GetSignedMapRootResponse, error)
	GetSignedMapRootByRevision(context.Context, *GetSignedMapRootByRevisionRequest) (*GetSignedMapRootResponse, error)
}

func RegisterTrillianMapServer(s *grpc.Server, srv TrillianMapServer) {
	s.RegisterService(&_TrillianMap_serviceDesc, srv)
}

func _TrillianMap_GetLeaves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMapLeavesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianMapServer).GetLeaves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianMap/GetLeaves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianMapServer).GetLeaves(ctx, req.(*GetMapLeavesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrillianMap_SetLeaves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMapLeavesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianMapServer).SetLeaves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianMap/SetLeaves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianMapServer).SetLeaves(ctx, req.(*SetMapLeavesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrillianMap_GetSignedMapRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignedMapRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianMapServer).GetSignedMapRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianMap/GetSignedMapRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianMapServer).GetSignedMapRoot(ctx, req.(*GetSignedMapRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrillianMap_GetSignedMapRootByRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignedMapRootByRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianMapServer).GetSignedMapRootByRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianMap/GetSignedMapRootByRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianMapServer).GetSignedMapRootByRevision(ctx, req.(*GetSignedMapRootByRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TrillianMap_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trillian.TrillianMap",
	HandlerType: (*TrillianMapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLeaves",
			Handler:    _TrillianMap_GetLeaves_Handler,
		},
		{
			MethodName: "SetLeaves",
			Handler:    _TrillianMap_SetLeaves_Handler,
		},
		{
			MethodName: "GetSignedMapRoot",
			Handler:    _TrillianMap_GetSignedMapRoot_Handler,
		},
		{
			MethodName: "GetSignedMapRootByRevision",
			Handler:    _TrillianMap_GetSignedMapRootByRevision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trillian_map_api.proto",
}

func init() { proto.RegisterFile("trillian_map_api.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x4e, 0xd4, 0x40,
	0x14, 0xb6, 0xbb, 0xfc, 0x2c, 0x67, 0x8d, 0xc1, 0x01, 0xa5, 0x56, 0x30, 0x50, 0x43, 0x94, 0x90,
	0x6c, 0xa5, 0x5e, 0xe9, 0x9d, 0xc4, 0x04, 0x30, 0xac, 0x21, 0xad, 0xc1, 0x3b, 0x37, 0x07, 0x76,
	0x80, 0x49, 0xda, 0xce, 0xd8, 0x0e, 0x1b, 0x94, 0x70, 0xe3, 0x85, 0x0f, 0xa0, 0x5e, 0xfb, 0x52,
	0xbc, 0x82, 0x0f, 0x62, 0x66, 0xa6, 0xfb, 0xd3, 0xdd, 0xee, 0xba, 0xf1, 0xae, 0x3d, 0xdf, 0x39,
	0xdf, 0xf9, 0xbe, 0x73, 0x4e, 0x06, 0x1e, 0xca, 0x94, 0x45, 0x11, 0xc3, 0xa4, 0x15, 0xa3, 0x68,
	0xa1, 0x60, 0x0d, 0x91, 0x72, 0xc9, 0x49, 0xad, 0x1b, 0x77, 0xee, 0x75, 0xbf, 0x0c, 0xe2, 0xac,
	0x9e, 0x73, 0x7e, 0x1e, 0x51, 0x0f, 0x05, 0xf3, 0x30, 0x49, 0xb8, 0x44, 0xc9, 0x78, 0x92, 0x19,
	0xd4, 0xfd, 0x0a, 0xf3, 0x4d, 0x14, 0x87, 0x14, 0xcf, 0xc8, 0x32, 0xcc, 0xb2, 0xa4, 0x4d, 0xaf,
	0x6c, 0x6b, 0xdd, 0x7a, 0x7e, 0x37, 0x30, 0x3f, 0xe4, 0x31, 0x2c, 0x44, 0x14, 0xcf, 0x5a, 0x17,
	0x98, 0x5d, 0xd8, 0x15, 0x8d, 0xd4, 0x54, 0x60, 0x1f, 0xb3, 0x0b, 0xb2, 0x06, 0xa0, 0xc1, 0x0e,
	0x46, 0x97, 0xd4, 0xae, 0x6a, 0x54, 0xa7, 0x1f, 0xab, 0x80, 0x82, 0xe9, 0x95, 0x4c, 0xb1, 0xd5,
	0x46, 0x89, 0xf6, 0x8c, 0x81, 0x75, 0xe4, 0x2d, 0x4a, 0x74, 0x3f, 0xc2, 0x62, 0xde, 0xfb, 0x20,
	0x39, 0x8d, 0x2e, 0x33, 0xc6, 0x13, 0xb2, 0x09, 0x33, 0xaa, 0x5e, 0x6b, 0xa8, 0xfb, 0xf7, 0x1b,
	0x3d, 0x33, 0x79, 0x66, 0xa0, 0x61, 0xb2, 0x0a, 0x0b, 0xac, 0x5b, 0x63, 0x57, 0xd6, 0xab, 0x8a,
	0xb8, 0x17, 0x70, 0x3f, 0xc1, 0xd2, 0x1e, 0x95, 0xa6, 0xa2, 0x43, 0xb3, 0x80, 0x7e, 0xbe, 0xa4,
	0x99, 0x24, 0x0f, 0x60, 0x4e, 0x0d, 0x8d, 0xb5, 0x35, 0x7b, 0x35, 0x98, 0x8d, 0x51, 0x1c, 0xb4,
	0xfb, 0xbe, 0x0d, 0x4f, 0xee, 0xdb, 0x81, 0x5a, 0x4a, 0x3b, 0x4c, 0x37, 0xa8, 0xea, 0xf4, 0xde,
	0xbf, 0xfb, 0xcb, 0x82, 0xe5, 0x62, 0x83, 0x4c, 0xf0, 0x24, 0xa3, 0x64, 0x1f, 0x88, 0xea, 0xa0,
	0x67, 0x52, 0xd4, 0x57, 0xf7, 0x9d, 0x11, 0x2f, 0x3d, 0xd7, 0xc1, 0x62, 0x3c, 0x3c, 0x07, 0x1f,
	0x6a, 0x8a, 0x29, 0xe5, 0x5c, 0xea, 0xf6, 0x75, 0x7f, 0xa5, 0x5f, 0x1f, 0xb2, 0xf3, 0x84, 0xb6,
	0x9b, 0x28, 0x02, 0xce, 0x65, 0x30, 0x1f, 0x9b, 0x0f, 0xf7, 0x87, 0x05, 0x4b, 0xe1, 0xf4, 0xbe,
	0xb7, 0x60, 0x2e, 0xd2, 0x79, 0xb9, 0xc0, 0x92, 0x61, 0xe7, 0x09, 0xe4, 0x15, 0xd4, 0x63, 0x14,
	0x82, 0xa6, 0x66, 0x93, 0x46, 0x90, 0x5d, 0xc8, 0x17, 0x34, 0x6d, 0x52, 0x89, 0x0a, 0x0f, 0xc0,
	0x24, 0xeb, 0x25, 0xbf, 0x83, 0xe5, 0xb0, 0x6c, 0x54, 0x83, 0x06, 0x2b, 0x53, 0x1a, 0x7c, 0x01,
	0x2b, 0x7b, 0x54, 0x16, 0xc1, 0x89, 0x1e, 0xdd, 0x63, 0xd8, 0x18, 0xae, 0xd8, 0xfd, 0x12, 0xe4,
	0x7b, 0xfc, 0xc7, 0x7c, 0x06, 0x2f, 0xa0, 0x32, 0x74, 0x01, 0xef, 0xc1, 0x1e, 0x55, 0xf2, 0xff,
	0xce, 0xfc, 0xdb, 0x2a, 0xd4, 0x3f, 0xe4, 0x39, 0x4d, 0x14, 0xe4, 0x10, 0x16, 0xf6, 0xa8, 0x34,
	0x23, 0x23, 0x6b, 0xfd, 0xf2, 0x92, 0xb3, 0x76, 0x9e, 0x8c, 0x83, 0x8d, 0x1e, 0xf7, 0x8e, 0x62,
	0x0b, 0xcb, 0xd8, 0xc2, 0xc9, 0x6c, 0x61, 0x39, 0xdb, 0x77, 0x0b, 0x16, 0x87, 0xcd, 0x93, 0x8d,
	0x82, 0x88, 0xb2, 0x15, 0x39, 0xee, 0xa4, 0x94, 0x9c, 0x7d, 0xfb, 0xdb, 0xed, 0x9f, 0x9f, 0x95,
	0x4d, 0xf2, 0xd4, 0xeb, 0xec, 0x9c, 0x50, 0x89, 0x3b, 0x5e, 0x8c, 0x22, 0xf3, 0xae, 0xcd, 0x7e,
	0x6e, 0x3c, 0x35, 0xd4, 0xec, 0x75, 0x84, 0x52, 0xed, 0xed, 0xb7, 0x05, 0xce, 0xf8, 0xed, 0x92,
	0xed, 0xf1, 0xfd, 0x46, 0x6e, 0x60, 0x2a, 0x71, 0x9e, 0x16, 0xb7, 0x45, 0x9e, 0x4d, 0x12, 0xe7,
	0x5d, 0x77, 0x8f, 0xe4, 0x66, 0xd7, 0x87, 0x47, 0xa7, 0x3c, 0x6e, 0x98, 0x07, 0xb8, 0x51, 0x7c,
	0x97, 0x77, 0x97, 0x06, 0xf6, 0xfd, 0x46, 0xb0, 0x23, 0x15, 0x3c, 0xb2, 0x4e, 0xe6, 0x34, 0xfa,
	0xf2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x75, 0x1f, 0xfa, 0xe9, 0x05, 0x00, 0x00,
}
