// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comment.proto

package commentpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Comment struct {
	UUID                 string               `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Message              string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	UserUUID             string               `protobuf:"bytes,3,opt,name=userUUID,proto3" json:"userUUID,omitempty"`
	ParentCommentUUID    string               `protobuf:"bytes,4,opt,name=parentCommentUUID,proto3" json:"parentCommentUUID,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{0}
}

func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *Comment) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Comment) GetUserUUID() string {
	if m != nil {
		return m.UserUUID
	}
	return ""
}

func (m *Comment) GetParentCommentUUID() string {
	if m != nil {
		return m.ParentCommentUUID
	}
	return ""
}

func (m *Comment) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Comment) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Comment) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

type GetRequest struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

type GetResponse struct {
	Comment              *Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{2}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetComment() *Comment {
	if m != nil {
		return m.Comment
	}
	return nil
}

type SetRequest struct {
	Comment              *Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{3}
}

func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetComment() *Comment {
	if m != nil {
		return m.Comment
	}
	return nil
}

type SetResponse struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetResponse) Reset()         { *m = SetResponse{} }
func (m *SetResponse) String() string { return proto.CompactTextString(m) }
func (*SetResponse) ProtoMessage()    {}
func (*SetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{4}
}

func (m *SetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResponse.Unmarshal(m, b)
}
func (m *SetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResponse.Marshal(b, m, deterministic)
}
func (m *SetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResponse.Merge(m, src)
}
func (m *SetResponse) XXX_Size() int {
	return xxx_messageInfo_SetResponse.Size(m)
}
func (m *SetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetResponse proto.InternalMessageInfo

func (m *SetResponse) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

type UpdateRequest struct {
	Comment              *Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{5}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetComment() *Comment {
	if m != nil {
		return m.Comment
	}
	return nil
}

type DeleteRequest struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{6}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func init() {
	proto.RegisterType((*Comment)(nil), "commentpb.Comment")
	proto.RegisterType((*GetRequest)(nil), "commentpb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "commentpb.GetResponse")
	proto.RegisterType((*SetRequest)(nil), "commentpb.SetRequest")
	proto.RegisterType((*SetResponse)(nil), "commentpb.SetResponse")
	proto.RegisterType((*UpdateRequest)(nil), "commentpb.UpdateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "commentpb.DeleteRequest")
}

func init() { proto.RegisterFile("comment.proto", fileDescriptor_749aee09ea917828) }

var fileDescriptor_749aee09ea917828 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4d, 0x4f, 0xf2, 0x40,
	0x10, 0x80, 0xdf, 0x02, 0x2f, 0x95, 0x69, 0x38, 0xb8, 0x89, 0xa4, 0xa9, 0x07, 0xb1, 0x5e, 0x38,
	0x90, 0x92, 0x60, 0x62, 0x8c, 0x1f, 0x07, 0x22, 0x86, 0x70, 0x33, 0xad, 0xfc, 0x80, 0x02, 0x23,
	0x31, 0xa1, 0x74, 0xed, 0x6e, 0x0f, 0xfe, 0x2a, 0xff, 0x9e, 0x47, 0xd3, 0xd9, 0x6d, 0xa1, 0x80,
	0x12, 0xbd, 0x75, 0x77, 0x9e, 0x67, 0x76, 0x3e, 0x52, 0x68, 0xce, 0xe2, 0x28, 0xc2, 0x95, 0xf4,
	0x78, 0x12, 0xcb, 0x98, 0x35, 0xf4, 0x91, 0x4f, 0x9d, 0xd3, 0x45, 0x1c, 0x2f, 0x96, 0xd8, 0xa3,
	0xc0, 0x34, 0x7d, 0xe9, 0x61, 0xc4, 0xe5, 0xbb, 0xe2, 0x9c, 0xb3, 0xed, 0xa0, 0x7c, 0x8d, 0x50,
	0xc8, 0x30, 0xe2, 0x0a, 0x70, 0x3f, 0x2a, 0x60, 0x3e, 0xa8, 0x5c, 0x8c, 0x41, 0x6d, 0x32, 0x19,
	0x0f, 0x6d, 0xa3, 0x6d, 0x74, 0x1a, 0x3e, 0x7d, 0x33, 0x1b, 0xcc, 0x08, 0x85, 0x08, 0x17, 0x68,
	0x57, 0xe8, 0x3a, 0x3f, 0x32, 0x07, 0x8e, 0x52, 0x81, 0x09, 0x19, 0x55, 0x0a, 0x15, 0x67, 0xd6,
	0x85, 0x63, 0x1e, 0x26, 0xb8, 0x92, 0x3a, 0x35, 0x41, 0x35, 0x82, 0x76, 0x03, 0xec, 0x1a, 0x1a,
	0xb3, 0x04, 0x43, 0x89, 0xf3, 0x81, 0xb4, 0xff, 0xb7, 0x8d, 0x8e, 0xd5, 0x77, 0x3c, 0x55, 0xb8,
	0x97, 0x17, 0xee, 0x3d, 0xe7, 0x85, 0xfb, 0x6b, 0x38, 0x33, 0x53, 0x3e, 0xd7, 0x66, 0xfd, 0xb0,
	0x59, 0xc0, 0x99, 0x39, 0xc7, 0x25, 0x2a, 0xd3, 0x3c, 0x6c, 0x16, 0xb0, 0xdb, 0x06, 0x18, 0xa1,
	0xf4, 0xf1, 0x2d, 0x45, 0xb1, 0x77, 0x66, 0xee, 0x2d, 0x58, 0x44, 0x08, 0x1e, 0xaf, 0x04, 0xb2,
	0x2e, 0x98, 0x7a, 0x5b, 0x44, 0x59, 0x7d, 0xe6, 0x15, 0xdb, 0xf3, 0xf4, 0x1c, 0xfc, 0x1c, 0x71,
	0x6f, 0x00, 0x82, 0x75, 0xfa, 0xdf, 0xb9, 0xe7, 0x60, 0x05, 0x1b, 0x0f, 0xef, 0xab, 0xed, 0x1e,
	0x9a, 0x13, 0x1a, 0xc2, 0xdf, 0x5e, 0xb8, 0x80, 0xe6, 0x90, 0x26, 0xf1, 0x43, 0xff, 0xfd, 0x4f,
	0x03, 0x40, 0x9b, 0x83, 0xa7, 0x31, 0xbb, 0x82, 0xea, 0x08, 0x25, 0x3b, 0xd9, 0xc8, 0xbb, 0x1e,
	0xa0, 0xd3, 0xda, 0xbe, 0x56, 0xc5, 0xbb, 0xff, 0x32, 0x2f, 0xd8, 0xf2, 0x82, 0xfd, 0x5e, 0x50,
	0xf2, 0xee, 0xa0, 0xae, 0x5a, 0x64, 0xf6, 0x06, 0x53, 0xea, 0xda, 0x69, 0xed, 0xec, 0xfa, 0x31,
	0xfb, 0x6b, 0x94, 0xad, 0x3a, 0x2c, 0xd9, 0xa5, 0xa6, 0xbf, 0xb7, 0xa7, 0x75, 0xba, 0xb9, 0xfc,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0x01, 0xd9, 0xc9, 0xa0, 0xaf, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommentAPIClient is the client API for CommentAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommentAPIClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type commentAPIClient struct {
	cc *grpc.ClientConn
}

func NewCommentAPIClient(cc *grpc.ClientConn) CommentAPIClient {
	return &commentAPIClient{cc}
}

func (c *commentAPIClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/commentpb.CommentAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentAPIClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/commentpb.CommentAPI/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentAPIClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/commentpb.CommentAPI/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentAPIClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/commentpb.CommentAPI/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentAPIServer is the server API for CommentAPI service.
type CommentAPIServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Update(context.Context, *UpdateRequest) (*empty.Empty, error)
	Delete(context.Context, *DeleteRequest) (*empty.Empty, error)
}

// UnimplementedCommentAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCommentAPIServer struct {
}

func (*UnimplementedCommentAPIServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCommentAPIServer) Set(ctx context.Context, req *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedCommentAPIServer) Update(ctx context.Context, req *UpdateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedCommentAPIServer) Delete(ctx context.Context, req *DeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterCommentAPIServer(s *grpc.Server, srv CommentAPIServer) {
	s.RegisterService(&_CommentAPI_serviceDesc, srv)
}

func _CommentAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commentpb.CommentAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentAPIServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentAPI_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentAPIServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commentpb.CommentAPI/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentAPIServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentAPI_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentAPIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commentpb.CommentAPI/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentAPIServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commentpb.CommentAPI/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentAPIServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommentAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "commentpb.CommentAPI",
	HandlerType: (*CommentAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CommentAPI_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _CommentAPI_Set_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CommentAPI_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CommentAPI_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
