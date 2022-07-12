// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/proto/user/user.proto

package go_micro_service_user

import (
	fmt "fmt"
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

type UserInfoRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoRequest) Reset()         { *m = UserInfoRequest{} }
func (m *UserInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UserInfoRequest) ProtoMessage()    {}
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7f67f7fef7ad23f, []int{0}
}

func (m *UserInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoRequest.Unmarshal(m, b)
}
func (m *UserInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoRequest.Marshal(b, m, deterministic)
}
func (m *UserInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoRequest.Merge(m, src)
}
func (m *UserInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UserInfoRequest.Size(m)
}
func (m *UserInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoRequest proto.InternalMessageInfo

func (m *UserInfoRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type UserInfoResponse struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoResponse) Reset()         { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()    {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7f67f7fef7ad23f, []int{1}
}

func (m *UserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResponse.Unmarshal(m, b)
}
func (m *UserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResponse.Marshal(b, m, deterministic)
}
func (m *UserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResponse.Merge(m, src)
}
func (m *UserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserInfoResponse.Size(m)
}
func (m *UserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResponse proto.InternalMessageInfo

func (m *UserInfoResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserInfoResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserInfoResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

type LoginRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Pwd                  string   `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7f67f7fef7ad23f, []int{2}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type LoginResponse struct {
	IsSuccess            bool     `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7f67f7fef7ad23f, []int{3}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

type RegisterRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	Pwd                  string   `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7f67f7fef7ad23f, []int{4}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RegisterRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *RegisterRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type RegisterResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7f67f7fef7ad23f, []int{5}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*UserInfoRequest)(nil), "go.micro.service.user.UserInfoRequest")
	proto.RegisterType((*UserInfoResponse)(nil), "go.micro.service.user.UserInfoResponse")
	proto.RegisterType((*LoginRequest)(nil), "go.micro.service.user.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "go.micro.service.user.LoginResponse")
	proto.RegisterType((*RegisterRequest)(nil), "go.micro.service.user.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "go.micro.service.user.RegisterResponse")
}

func init() { proto.RegisterFile("user/proto/user/user.proto", fileDescriptor_a7f67f7fef7ad23f) }

var fileDescriptor_a7f67f7fef7ad23f = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0xb7, 0xf5, 0x7d, 0xb7, 0xf5, 0xa8, 0x6c, 0x04, 0xc4, 0x31, 0x19, 0x48, 0x14, 0xf5,
	0x42, 0x32, 0xd0, 0x6b, 0xaf, 0x65, 0x20, 0x5e, 0x54, 0xbd, 0x94, 0x59, 0xdb, 0xb3, 0x92, 0x8b,
	0x36, 0x35, 0xa7, 0xd5, 0x4f, 0xe3, 0x77, 0x95, 0xa4, 0xad, 0xb5, 0xc5, 0xea, 0x6e, 0x42, 0xce,
	0xbf, 0xe7, 0xf9, 0x91, 0x1c, 0x98, 0xe7, 0x84, 0x7a, 0x99, 0x6a, 0x95, 0xa9, 0xa5, 0xbd, 0x9a,
	0x43, 0xd8, 0x98, 0xed, 0x47, 0x4a, 0xc4, 0x32, 0xd0, 0x4a, 0x10, 0xea, 0x37, 0x19, 0xa0, 0x30,
	0x45, 0x2e, 0x60, 0xf2, 0x48, 0xa8, 0x57, 0xc9, 0x46, 0x79, 0xf8, 0x9a, 0x23, 0x65, 0xec, 0x10,
	0x5c, 0x53, 0x5a, 0x27, 0x7e, 0x8c, 0xb3, 0xfe, 0x51, 0xff, 0xdc, 0xf5, 0xc6, 0x26, 0x71, 0xe7,
	0xc7, 0xc8, 0x23, 0x98, 0xd6, 0xfd, 0x94, 0xaa, 0x84, 0x90, 0x1d, 0xc0, 0xc8, 0x0e, 0xc8, 0xd0,
	0xb6, 0x3b, 0xde, 0xd0, 0x84, 0xab, 0xb0, 0xa9, 0x34, 0x68, 0x2a, 0xb1, 0x05, 0xc0, 0x46, 0x6a,
	0xca, 0x8a, 0xaa, 0x63, 0xab, 0xae, 0xcd, 0x58, 0xa3, 0x6b, 0xd8, 0xbd, 0x55, 0x91, 0x4c, 0xb6,
	0xa1, 0x62, 0x53, 0x70, 0xd2, 0xf7, 0xb0, 0xb4, 0x30, 0x57, 0x2e, 0x60, 0xaf, 0x1c, 0x2f, 0x21,
	0x17, 0x00, 0x92, 0xd6, 0x94, 0x07, 0x01, 0x12, 0x59, 0x81, 0xb1, 0xe7, 0x4a, 0xba, 0x2f, 0x12,
	0x7c, 0x0d, 0x13, 0x0f, 0x23, 0x49, 0x19, 0xea, 0xad, 0x1c, 0x9b, 0xf4, 0x83, 0x16, 0x7d, 0x05,
	0xe4, 0xd4, 0x40, 0x17, 0x30, 0xad, 0x0d, 0x4a, 0xa6, 0x19, 0x8c, 0x62, 0x24, 0xf2, 0xa3, 0x4a,
	0xbf, 0x0a, 0x2f, 0x3f, 0x06, 0xf0, 0xcf, 0xbc, 0x33, 0x7b, 0x82, 0x71, 0x35, 0xc6, 0x4e, 0xc5,
	0x8f, 0x7f, 0x28, 0x5a, 0xe0, 0xf3, 0xb3, 0x3f, 0xfb, 0x0a, 0x7f, 0xde, 0x63, 0x0f, 0xf0, 0xdf,
	0x3e, 0x13, 0x3b, 0xee, 0x98, 0xf9, 0xfe, 0x07, 0xf3, 0x93, 0xdf, 0x9b, 0xbe, 0x54, 0x9f, 0x61,
	0xe7, 0x06, 0xb3, 0x6a, 0x4f, 0x3a, 0xb9, 0x5b, 0x8b, 0xd7, 0xc9, 0xdd, 0x5e, 0x38, 0xde, 0x7b,
	0x19, 0xda, 0xa5, 0xbe, 0xfa, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x32, 0x57, 0x9b, 0xf2, 0x02,
	0x00, 0x00,
}