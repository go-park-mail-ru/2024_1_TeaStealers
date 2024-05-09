// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: complex.proto

package gen

import (
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

type CreateCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	YearFounded int32  `protobuf:"varint,2,opt,name=YearFounded,proto3" json:"YearFounded,omitempty"`
	Phone       string `protobuf:"bytes,3,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *CreateCompanyRequest) Reset() {
	*x = CreateCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCompanyRequest) ProtoMessage() {}

func (x *CreateCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_complex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCompanyRequest.ProtoReflect.Descriptor instead.
func (*CreateCompanyRequest) Descriptor() ([]byte, []int) {
	return file_complex_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCompanyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCompanyRequest) GetYearFounded() int32 {
	if x != nil {
		return x.YearFounded
	}
	return 0
}

func (x *CreateCompanyRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateCompanyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateCompanyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Photo        string `protobuf:"bytes,3,opt,name=Photo,proto3" json:"Photo,omitempty"`
	YearFounded  int32  `protobuf:"varint,4,opt,name=YearFounded,proto3" json:"YearFounded,omitempty"`
	Phone        string `protobuf:"bytes,5,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Description  string `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
	DateCreation string `protobuf:"bytes,7,opt,name=DateCreation,proto3" json:"DateCreation,omitempty"`
	IsDeleted    string `protobuf:"bytes,8,opt,name=IsDeleted,proto3" json:"IsDeleted,omitempty"`
}

func (x *CreateCompanyResponse) Reset() {
	*x = CreateCompanyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCompanyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCompanyResponse) ProtoMessage() {}

func (x *CreateCompanyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_complex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCompanyResponse.ProtoReflect.Descriptor instead.
func (*CreateCompanyResponse) Descriptor() ([]byte, []int) {
	return file_complex_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCompanyResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateCompanyResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCompanyResponse) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *CreateCompanyResponse) GetYearFounded() int32 {
	if x != nil {
		return x.YearFounded
	}
	return 0
}

func (x *CreateCompanyResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateCompanyResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateCompanyResponse) GetDateCreation() string {
	if x != nil {
		return x.DateCreation
	}
	return ""
}

func (x *CreateCompanyResponse) GetIsDeleted() string {
	if x != nil {
		return x.IsDeleted
	}
	return ""
}

type GetCompanyByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetCompanyByIdRequest) Reset() {
	*x = GetCompanyByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompanyByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyByIdRequest) ProtoMessage() {}

func (x *GetCompanyByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_complex_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyByIdRequest.ProtoReflect.Descriptor instead.
func (*GetCompanyByIdRequest) Descriptor() ([]byte, []int) {
	return file_complex_proto_rawDescGZIP(), []int{2}
}

func (x *GetCompanyByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCompanyByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Photo       string `protobuf:"bytes,3,opt,name=Photo,proto3" json:"Photo,omitempty"`
	YearFounded int32  `protobuf:"varint,4,opt,name=YearFounded,proto3" json:"YearFounded,omitempty"`
	Phone       string `protobuf:"bytes,5,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *GetCompanyByIdResponse) Reset() {
	*x = GetCompanyByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompanyByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyByIdResponse) ProtoMessage() {}

func (x *GetCompanyByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_complex_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyByIdResponse.ProtoReflect.Descriptor instead.
func (*GetCompanyByIdResponse) Descriptor() ([]byte, []int) {
	return file_complex_proto_rawDescGZIP(), []int{3}
}

func (x *GetCompanyByIdResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetCompanyByIdResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetCompanyByIdResponse) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *GetCompanyByIdResponse) GetYearFounded() int32 {
	if x != nil {
		return x.YearFounded
	}
	return 0
}

func (x *GetCompanyByIdResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetCompanyByIdResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateCompanyPhotoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File      []byte `protobuf:"bytes,1,opt,name=File,proto3" json:"File,omitempty"`
	FileType  string `protobuf:"bytes,2,opt,name=FileType,proto3" json:"FileType,omitempty"`
	CompanyId int64  `protobuf:"varint,3,opt,name=CompanyId,proto3" json:"CompanyId,omitempty"`
}

func (x *UpdateCompanyPhotoRequest) Reset() {
	*x = UpdateCompanyPhotoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCompanyPhotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCompanyPhotoRequest) ProtoMessage() {}

func (x *UpdateCompanyPhotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_complex_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCompanyPhotoRequest.ProtoReflect.Descriptor instead.
func (*UpdateCompanyPhotoRequest) Descriptor() ([]byte, []int) {
	return file_complex_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCompanyPhotoRequest) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *UpdateCompanyPhotoRequest) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *UpdateCompanyPhotoRequest) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

type UpdateCompanyPhotoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=Filename,proto3" json:"Filename,omitempty"`
}

func (x *UpdateCompanyPhotoResponse) Reset() {
	*x = UpdateCompanyPhotoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCompanyPhotoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCompanyPhotoResponse) ProtoMessage() {}

func (x *UpdateCompanyPhotoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_complex_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCompanyPhotoResponse.ProtoReflect.Descriptor instead.
func (*UpdateCompanyPhotoResponse) Descriptor() ([]byte, []int) {
	return file_complex_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCompanyPhotoResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

var File_complex_proto protoreflect.FileDescriptor

var file_complex_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x22, 0x84, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x59, 0x65, 0x61, 0x72, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x59, 0x65, 0x61, 0x72,
	0x46, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xed, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68,
	0x6f, 0x74, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x59, 0x65, 0x61, 0x72, 0x46, 0x6f, 0x75, 0x6e, 0x64,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x59, 0x65, 0x61, 0x72, 0x46, 0x6f,
	0x75, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a,
	0x0c, 0x44, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x44, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22,
	0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0xac, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x0a,
	0x0b, 0x59, 0x65, 0x61, 0x72, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x59, 0x65, 0x61, 0x72, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x69, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x49, 0x64, 0x22, 0x38, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x91, 0x02, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x78, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5f, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x68, 0x6f,
	0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x78, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x31, 0x5a, 0x2f, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x3b,
	0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_complex_proto_rawDescOnce sync.Once
	file_complex_proto_rawDescData = file_complex_proto_rawDesc
)

func file_complex_proto_rawDescGZIP() []byte {
	file_complex_proto_rawDescOnce.Do(func() {
		file_complex_proto_rawDescData = protoimpl.X.CompressGZIP(file_complex_proto_rawDescData)
	})
	return file_complex_proto_rawDescData
}

var file_complex_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_complex_proto_goTypes = []interface{}{
	(*CreateCompanyRequest)(nil),       // 0: complex.CreateCompanyRequest
	(*CreateCompanyResponse)(nil),      // 1: complex.CreateCompanyResponse
	(*GetCompanyByIdRequest)(nil),      // 2: complex.GetCompanyByIdRequest
	(*GetCompanyByIdResponse)(nil),     // 3: complex.GetCompanyByIdResponse
	(*UpdateCompanyPhotoRequest)(nil),  // 4: complex.UpdateCompanyPhotoRequest
	(*UpdateCompanyPhotoResponse)(nil), // 5: complex.UpdateCompanyPhotoResponse
}
var file_complex_proto_depIdxs = []int32{
	0, // 0: complex.Complex.CreateCompany:input_type -> complex.CreateCompanyRequest
	2, // 1: complex.Complex.GetCompanyById:input_type -> complex.GetCompanyByIdRequest
	4, // 2: complex.Complex.UpdateCompanyPhoto:input_type -> complex.UpdateCompanyPhotoRequest
	1, // 3: complex.Complex.CreateCompany:output_type -> complex.CreateCompanyResponse
	3, // 4: complex.Complex.GetCompanyById:output_type -> complex.GetCompanyByIdResponse
	5, // 5: complex.Complex.UpdateCompanyPhoto:output_type -> complex.UpdateCompanyPhotoResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_complex_proto_init() }
func file_complex_proto_init() {
	if File_complex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_complex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCompanyRequest); i {
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
		file_complex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCompanyResponse); i {
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
		file_complex_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompanyByIdRequest); i {
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
		file_complex_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompanyByIdResponse); i {
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
		file_complex_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCompanyPhotoRequest); i {
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
		file_complex_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCompanyPhotoResponse); i {
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
			RawDescriptor: file_complex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_complex_proto_goTypes,
		DependencyIndexes: file_complex_proto_depIdxs,
		MessageInfos:      file_complex_proto_msgTypes,
	}.Build()
	File_complex_proto = out.File
	file_complex_proto_rawDesc = nil
	file_complex_proto_goTypes = nil
	file_complex_proto_depIdxs = nil
}
