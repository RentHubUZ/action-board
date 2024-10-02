// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: tariff.proto

package tariff

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

type GetTariffReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Tarif ID
}

func (x *GetTariffReq) Reset() {
	*x = GetTariffReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTariffReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTariffReq) ProtoMessage() {}

func (x *GetTariffReq) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTariffReq.ProtoReflect.Descriptor instead.
func (*GetTariffReq) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{0}
}

func (x *GetTariffReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTariffRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tariff *Tariff `protobuf:"bytes,1,opt,name=tariff,proto3" json:"tariff,omitempty"`
}

func (x *GetTariffRes) Reset() {
	*x = GetTariffRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTariffRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTariffRes) ProtoMessage() {}

func (x *GetTariffRes) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTariffRes.ProtoReflect.Descriptor instead.
func (*GetTariffRes) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{1}
}

func (x *GetTariffRes) GetTariff() *Tariff {
	if x != nil {
		return x.Tariff
	}
	return nil
}

type GetAllTariffReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetAllTariffReq) Reset() {
	*x = GetAllTariffReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTariffReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTariffReq) ProtoMessage() {}

func (x *GetAllTariffReq) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTariffReq.ProtoReflect.Descriptor instead.
func (*GetAllTariffReq) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllTariffReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllTariffReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAllTariffRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tariffs []*Tariff `protobuf:"bytes,1,rep,name=tariffs,proto3" json:"tariffs,omitempty"`
}

func (x *GetAllTariffRes) Reset() {
	*x = GetAllTariffRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTariffRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTariffRes) ProtoMessage() {}

func (x *GetAllTariffRes) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTariffRes.ProtoReflect.Descriptor instead.
func (*GetAllTariffRes) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllTariffRes) GetTariffs() []*Tariff {
	if x != nil {
		return x.Tariffs
	}
	return nil
}

type CreateTariffReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Days   int32   `protobuf:"varint,2,opt,name=days,proto3" json:"days,omitempty"`
	Price  float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Offers string  `protobuf:"bytes,4,opt,name=offers,proto3" json:"offers,omitempty"`
}

func (x *CreateTariffReq) Reset() {
	*x = CreateTariffReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTariffReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTariffReq) ProtoMessage() {}

func (x *CreateTariffReq) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTariffReq.ProtoReflect.Descriptor instead.
func (*CreateTariffReq) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTariffReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTariffReq) GetDays() int32 {
	if x != nil {
		return x.Days
	}
	return 0
}

func (x *CreateTariffReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateTariffReq) GetOffers() string {
	if x != nil {
		return x.Offers
	}
	return ""
}

type CreateTariffRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateTariffRes) Reset() {
	*x = CreateTariffRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTariffRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTariffRes) ProtoMessage() {}

func (x *CreateTariffRes) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTariffRes.ProtoReflect.Descriptor instead.
func (*CreateTariffRes) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTariffRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateTariffReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Days   int32   `protobuf:"varint,3,opt,name=days,proto3" json:"days,omitempty"`
	Price  float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Offers string  `protobuf:"bytes,5,opt,name=offers,proto3" json:"offers,omitempty"`
}

func (x *UpdateTariffReq) Reset() {
	*x = UpdateTariffReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTariffReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTariffReq) ProtoMessage() {}

func (x *UpdateTariffReq) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTariffReq.ProtoReflect.Descriptor instead.
func (*UpdateTariffReq) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTariffReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTariffReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateTariffReq) GetDays() int32 {
	if x != nil {
		return x.Days
	}
	return 0
}

func (x *UpdateTariffReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateTariffReq) GetOffers() string {
	if x != nil {
		return x.Offers
	}
	return ""
}

type UpdateTariffRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UpdateTariffRes) Reset() {
	*x = UpdateTariffRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTariffRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTariffRes) ProtoMessage() {}

func (x *UpdateTariffRes) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTariffRes.ProtoReflect.Descriptor instead.
func (*UpdateTariffRes) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTariffRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteTariffReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTariffReq) Reset() {
	*x = DeleteTariffReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTariffReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTariffReq) ProtoMessage() {}

func (x *DeleteTariffReq) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTariffReq.ProtoReflect.Descriptor instead.
func (*DeleteTariffReq) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteTariffReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTariffRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteTariffRes) Reset() {
	*x = DeleteTariffRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTariffRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTariffRes) ProtoMessage() {}

func (x *DeleteTariffRes) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTariffRes.ProtoReflect.Descriptor instead.
func (*DeleteTariffRes) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteTariffRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Tariff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Days      int32   `protobuf:"varint,3,opt,name=days,proto3" json:"days,omitempty"`
	Price     float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Offers    string  `protobuf:"bytes,5,opt,name=offers,proto3" json:"offers,omitempty"`
	CreatedAt string  `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string  `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Tariff) Reset() {
	*x = Tariff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tariff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tariff) ProtoMessage() {}

func (x *Tariff) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tariff.ProtoReflect.Descriptor instead.
func (*Tariff) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{10}
}

func (x *Tariff) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tariff) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tariff) GetDays() int32 {
	if x != nil {
		return x.Days
	}
	return 0
}

func (x *Tariff) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Tariff) GetOffers() string {
	if x != nil {
		return x.Offers
	}
	return ""
}

func (x *Tariff) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Tariff) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_tariff_proto protoreflect.FileDescriptor

var file_tariff_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x22, 0x1e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e,
	0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x22, 0x3b,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x3b, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x12, 0x28,
	0x0a, 0x07, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52,
	0x07, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x22, 0x67, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64,
	0x61, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x65, 0x72,
	0x73, 0x22, 0x21, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x77, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x79, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x22, 0x2b, 0x0a,
	0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x21, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a,
	0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xac, 0x01, 0x0a, 0x06, 0x54,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x79,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xb2, 0x02, 0x0a, 0x0d, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x14, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x12, 0x3a,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65,
	0x71, 0x1a, 0x17, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e,
	0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x17, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x74, 0x61, 0x72, 0x69,
	0x66, 0x66, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52,
	0x65, 0x73, 0x12, 0x3a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x74,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69,
	0x66, 0x66, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x42, 0x11,
	0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tariff_proto_rawDescOnce sync.Once
	file_tariff_proto_rawDescData = file_tariff_proto_rawDesc
)

func file_tariff_proto_rawDescGZIP() []byte {
	file_tariff_proto_rawDescOnce.Do(func() {
		file_tariff_proto_rawDescData = protoimpl.X.CompressGZIP(file_tariff_proto_rawDescData)
	})
	return file_tariff_proto_rawDescData
}

var file_tariff_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_tariff_proto_goTypes = []any{
	(*GetTariffReq)(nil),    // 0: tariff.GetTariffReq
	(*GetTariffRes)(nil),    // 1: tariff.GetTariffRes
	(*GetAllTariffReq)(nil), // 2: tariff.GetAllTariffReq
	(*GetAllTariffRes)(nil), // 3: tariff.GetAllTariffRes
	(*CreateTariffReq)(nil), // 4: tariff.CreateTariffReq
	(*CreateTariffRes)(nil), // 5: tariff.CreateTariffRes
	(*UpdateTariffReq)(nil), // 6: tariff.UpdateTariffReq
	(*UpdateTariffRes)(nil), // 7: tariff.UpdateTariffRes
	(*DeleteTariffReq)(nil), // 8: tariff.DeleteTariffReq
	(*DeleteTariffRes)(nil), // 9: tariff.DeleteTariffRes
	(*Tariff)(nil),          // 10: tariff.Tariff
}
var file_tariff_proto_depIdxs = []int32{
	10, // 0: tariff.GetTariffRes.tariff:type_name -> tariff.Tariff
	10, // 1: tariff.GetAllTariffRes.tariffs:type_name -> tariff.Tariff
	0,  // 2: tariff.TariffService.Get:input_type -> tariff.GetTariffReq
	2,  // 3: tariff.TariffService.GetAll:input_type -> tariff.GetAllTariffReq
	4,  // 4: tariff.TariffService.Create:input_type -> tariff.CreateTariffReq
	6,  // 5: tariff.TariffService.Update:input_type -> tariff.UpdateTariffReq
	8,  // 6: tariff.TariffService.Delete:input_type -> tariff.DeleteTariffReq
	1,  // 7: tariff.TariffService.Get:output_type -> tariff.GetTariffRes
	3,  // 8: tariff.TariffService.GetAll:output_type -> tariff.GetAllTariffRes
	5,  // 9: tariff.TariffService.Create:output_type -> tariff.CreateTariffRes
	7,  // 10: tariff.TariffService.Update:output_type -> tariff.UpdateTariffRes
	9,  // 11: tariff.TariffService.Delete:output_type -> tariff.DeleteTariffRes
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_tariff_proto_init() }
func file_tariff_proto_init() {
	if File_tariff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tariff_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetTariffReq); i {
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
		file_tariff_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetTariffRes); i {
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
		file_tariff_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllTariffReq); i {
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
		file_tariff_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllTariffRes); i {
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
		file_tariff_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTariffReq); i {
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
		file_tariff_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTariffRes); i {
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
		file_tariff_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTariffReq); i {
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
		file_tariff_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTariffRes); i {
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
		file_tariff_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTariffReq); i {
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
		file_tariff_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTariffRes); i {
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
		file_tariff_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*Tariff); i {
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
			RawDescriptor: file_tariff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tariff_proto_goTypes,
		DependencyIndexes: file_tariff_proto_depIdxs,
		MessageInfos:      file_tariff_proto_msgTypes,
	}.Build()
	File_tariff_proto = out.File
	file_tariff_proto_rawDesc = nil
	file_tariff_proto_goTypes = nil
	file_tariff_proto_depIdxs = nil
}
