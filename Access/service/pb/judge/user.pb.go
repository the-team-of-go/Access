// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: user.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Cpuused  float64 `protobuf:"fixed64,2,opt,name=cpuused,proto3" json:"cpuused,omitempty"`
	Memused  float64 `protobuf:"fixed64,3,opt,name=memused,proto3" json:"memused,omitempty"`
	DiskUsed float64 `protobuf:"fixed64,4,opt,name=DiskUsed,proto3" json:"DiskUsed,omitempty"`
	Timeout  int64   `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserRequest) GetCpuused() float64 {
	if x != nil {
		return x.Cpuused
	}
	return 0
}

func (x *UserRequest) GetMemused() float64 {
	if x != nil {
		return x.Memused
	}
	return 0
}

func (x *UserRequest) GetDiskUsed() float64 {
	if x != nil {
		return x.DiskUsed
	}
	return 0
}

func (x *UserRequest) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Mesg string `protobuf:"bytes,2,opt,name=mesg,proto3" json:"mesg,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UserResponse) GetMesg() string {
	if x != nil {
		return x.Mesg
	}
	return ""
}

type SqlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Timeout        int32  `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	MaxValueCpu    int32  `protobuf:"varint,3,opt,name=maxValueCpu,proto3" json:"maxValueCpu,omitempty"`
	MinValueCpu    int32  `protobuf:"varint,4,opt,name=minValueCpu,proto3" json:"minValueCpu,omitempty"`
	AvergValueCpue int32  `protobuf:"varint,5,opt,name=avergValueCpue,proto3" json:"avergValueCpue,omitempty"`
	MaxValueMem    int32  `protobuf:"varint,6,opt,name=maxValueMem,proto3" json:"maxValueMem,omitempty"`
	MinValueMem    int32  `protobuf:"varint,7,opt,name=minValueMem,proto3" json:"minValueMem,omitempty"`
	AvergValueMem  int32  `protobuf:"varint,8,opt,name=avergValueMem,proto3" json:"avergValueMem,omitempty"`
}

func (x *SqlRequest) Reset() {
	*x = SqlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqlRequest) ProtoMessage() {}

func (x *SqlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqlRequest.ProtoReflect.Descriptor instead.
func (*SqlRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *SqlRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SqlRequest) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *SqlRequest) GetMaxValueCpu() int32 {
	if x != nil {
		return x.MaxValueCpu
	}
	return 0
}

func (x *SqlRequest) GetMinValueCpu() int32 {
	if x != nil {
		return x.MinValueCpu
	}
	return 0
}

func (x *SqlRequest) GetAvergValueCpue() int32 {
	if x != nil {
		return x.AvergValueCpue
	}
	return 0
}

func (x *SqlRequest) GetMaxValueMem() int32 {
	if x != nil {
		return x.MaxValueMem
	}
	return 0
}

func (x *SqlRequest) GetMinValueMem() int32 {
	if x != nil {
		return x.MinValueMem
	}
	return 0
}

func (x *SqlRequest) GetAvergValueMem() int32 {
	if x != nil {
		return x.AvergValueMem
	}
	return 0
}

type SqlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Mesg string `protobuf:"bytes,2,opt,name=mesg,proto3" json:"mesg,omitempty"`
}

func (x *SqlResponse) Reset() {
	*x = SqlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqlResponse) ProtoMessage() {}

func (x *SqlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqlResponse.ProtoReflect.Descriptor instead.
func (*SqlResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *SqlResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SqlResponse) GetMesg() string {
	if x != nil {
		return x.Mesg
	}
	return ""
}

type AdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  []string `protobuf:"bytes,1,rep,name=name,proto3" json:"name,omitempty"`
	Email []string `protobuf:"bytes,2,rep,name=email,proto3" json:"email,omitempty"`
}

func (x *AdminRequest) Reset() {
	*x = AdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRequest) ProtoMessage() {}

func (x *AdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRequest.ProtoReflect.Descriptor instead.
func (*AdminRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *AdminRequest) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *AdminRequest) GetEmail() []string {
	if x != nil {
		return x.Email
	}
	return nil
}

type AdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Mesg string `protobuf:"bytes,2,opt,name=mesg,proto3" json:"mesg,omitempty"`
}

func (x *AdminResponse) Reset() {
	*x = AdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminResponse) ProtoMessage() {}

func (x *AdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminResponse.ProtoReflect.Descriptor instead.
func (*AdminResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *AdminResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AdminResponse) GetMesg() string {
	if x != nil {
		return x.Mesg
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x70, 0x75, 0x75, 0x73, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x63, 0x70, 0x75, 0x75, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x6d, 0x75, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07,
	0x6d, 0x65, 0x6d, 0x75, 0x73, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x6b, 0x55,
	0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x44, 0x69, 0x73, 0x6b, 0x55,
	0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x36, 0x0a,
	0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x65, 0x73, 0x67, 0x22, 0x90, 0x02, 0x0a, 0x0a, 0x53, 0x71, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x43, 0x70,
	0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x43, 0x70, 0x75, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x43, 0x70, 0x75, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x43, 0x70, 0x75, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x76, 0x65, 0x72, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x43, 0x70, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x61, 0x76, 0x65, 0x72, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x43, 0x70, 0x75, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x65, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x65, 0x6d,
	0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x65, 0x6d, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d,
	0x65, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x76, 0x65, 0x72, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4d, 0x65, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x76, 0x65, 0x72, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x65, 0x6d, 0x22, 0x35, 0x0a, 0x0b, 0x53, 0x71, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x65, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x73, 0x67, 0x22,
	0x38, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x37, 0x0a, 0x0d, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x65, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65,
	0x73, 0x67, 0x32, 0x4b, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32,
	0x4a, 0x0a, 0x11, 0x53, 0x71, 0x6c, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x71, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x71, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x71,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x4e, 0x0a, 0x0f, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_proto_goTypes = []interface{}{
	(*UserRequest)(nil),   // 0: proto.UserRequest
	(*UserResponse)(nil),  // 1: proto.UserResponse
	(*SqlRequest)(nil),    // 2: proto.SqlRequest
	(*SqlResponse)(nil),   // 3: proto.SqlResponse
	(*AdminRequest)(nil),  // 4: proto.AdminRequest
	(*AdminResponse)(nil), // 5: proto.AdminResponse
}
var file_user_proto_depIdxs = []int32{
	0, // 0: proto.UserInfoService.GetUserInfo:input_type -> proto.UserRequest
	2, // 1: proto.SqlDefaultService.GetSqlInfo:input_type -> proto.SqlRequest
	4, // 2: proto.AdminGetService.GetAdminInfo:input_type -> proto.AdminRequest
	1, // 3: proto.UserInfoService.GetUserInfo:output_type -> proto.UserResponse
	3, // 4: proto.SqlDefaultService.GetSqlInfo:output_type -> proto.SqlResponse
	5, // 5: proto.AdminGetService.GetAdminInfo:output_type -> proto.AdminResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqlRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqlResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserInfoServiceClient is the client API for UserInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserInfoServiceClient interface {
	GetUserInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInfoServiceClient(cc grpc.ClientConnInterface) UserInfoServiceClient {
	return &userInfoServiceClient{cc}
}

func (c *userInfoServiceClient) GetUserInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/proto.UserInfoService/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoServiceServer is the server API for UserInfoService service.
type UserInfoServiceServer interface {
	GetUserInfo(context.Context, *UserRequest) (*UserResponse, error)
}

// UnimplementedUserInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserInfoServiceServer struct {
}

func (*UnimplementedUserInfoServiceServer) GetUserInfo(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}

func RegisterUserInfoServiceServer(s *grpc.Server, srv UserInfoServiceServer) {
	s.RegisterService(&_UserInfoService_serviceDesc, srv)
}

func _UserInfoService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserInfoService/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServiceServer).GetUserInfo(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserInfoService",
	HandlerType: (*UserInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _UserInfoService_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

// SqlDefaultServiceClient is the client API for SqlDefaultService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SqlDefaultServiceClient interface {
	GetSqlInfo(ctx context.Context, in *SqlRequest, opts ...grpc.CallOption) (*SqlResponse, error)
}

type sqlDefaultServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSqlDefaultServiceClient(cc grpc.ClientConnInterface) SqlDefaultServiceClient {
	return &sqlDefaultServiceClient{cc}
}

func (c *sqlDefaultServiceClient) GetSqlInfo(ctx context.Context, in *SqlRequest, opts ...grpc.CallOption) (*SqlResponse, error) {
	out := new(SqlResponse)
	err := c.cc.Invoke(ctx, "/proto.SqlDefaultService/GetSqlInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SqlDefaultServiceServer is the server API for SqlDefaultService service.
type SqlDefaultServiceServer interface {
	GetSqlInfo(context.Context, *SqlRequest) (*SqlResponse, error)
}

// UnimplementedSqlDefaultServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSqlDefaultServiceServer struct {
}

func (*UnimplementedSqlDefaultServiceServer) GetSqlInfo(context.Context, *SqlRequest) (*SqlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSqlInfo not implemented")
}

func RegisterSqlDefaultServiceServer(s *grpc.Server, srv SqlDefaultServiceServer) {
	s.RegisterService(&_SqlDefaultService_serviceDesc, srv)
}

func _SqlDefaultService_GetSqlInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqlDefaultServiceServer).GetSqlInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SqlDefaultService/GetSqlInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqlDefaultServiceServer).GetSqlInfo(ctx, req.(*SqlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SqlDefaultService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SqlDefaultService",
	HandlerType: (*SqlDefaultServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSqlInfo",
			Handler:    _SqlDefaultService_GetSqlInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

// AdminGetServiceClient is the client API for AdminGetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminGetServiceClient interface {
	GetAdminInfo(ctx context.Context, in *AdminRequest, opts ...grpc.CallOption) (*AdminResponse, error)
}

type adminGetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminGetServiceClient(cc grpc.ClientConnInterface) AdminGetServiceClient {
	return &adminGetServiceClient{cc}
}

func (c *adminGetServiceClient) GetAdminInfo(ctx context.Context, in *AdminRequest, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, "/proto.AdminGetService/GetAdminInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminGetServiceServer is the server API for AdminGetService service.
type AdminGetServiceServer interface {
	GetAdminInfo(context.Context, *AdminRequest) (*AdminResponse, error)
}

// UnimplementedAdminGetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdminGetServiceServer struct {
}

func (*UnimplementedAdminGetServiceServer) GetAdminInfo(context.Context, *AdminRequest) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminInfo not implemented")
}

func RegisterAdminGetServiceServer(s *grpc.Server, srv AdminGetServiceServer) {
	s.RegisterService(&_AdminGetService_serviceDesc, srv)
}

func _AdminGetService_GetAdminInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminGetServiceServer).GetAdminInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AdminGetService/GetAdminInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminGetServiceServer).GetAdminInfo(ctx, req.(*AdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminGetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AdminGetService",
	HandlerType: (*AdminGetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdminInfo",
			Handler:    _AdminGetService_GetAdminInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
