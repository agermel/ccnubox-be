// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.4
// source: elecprice/v1/elecprice.proto

package elecpricev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetArchitectureRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AreaName      string                 `protobuf:"bytes,1,opt,name=AreaName,proto3" json:"AreaName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetArchitectureRequest) Reset() {
	*x = GetArchitectureRequest{}
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetArchitectureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArchitectureRequest) ProtoMessage() {}

func (x *GetArchitectureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArchitectureRequest.ProtoReflect.Descriptor instead.
func (*GetArchitectureRequest) Descriptor() ([]byte, []int) {
	return file_elecprice_v1_elecprice_proto_rawDescGZIP(), []int{0}
}

func (x *GetArchitectureRequest) GetAreaName() string {
	if x != nil {
		return x.AreaName
	}
	return ""
}

type GetArchitectureResponse struct {
	state            protoimpl.MessageState                  `protogen:"open.v1"`
	ArchitectureList []*GetArchitectureResponse_Architecture `protobuf:"bytes,1,rep,name=ArchitectureList,proto3" json:"ArchitectureList,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *GetArchitectureResponse) Reset() {
	*x = GetArchitectureResponse{}
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetArchitectureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArchitectureResponse) ProtoMessage() {}

func (x *GetArchitectureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArchitectureResponse.ProtoReflect.Descriptor instead.
func (*GetArchitectureResponse) Descriptor() ([]byte, []int) {
	return file_elecprice_v1_elecprice_proto_rawDescGZIP(), []int{1}
}

func (x *GetArchitectureResponse) GetArchitectureList() []*GetArchitectureResponse_Architecture {
	if x != nil {
		return x.ArchitectureList
	}
	return nil
}

type GetRoomInfoRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	ArchitectureID string                 `protobuf:"bytes,1,opt,name=ArchitectureID,proto3" json:"ArchitectureID,omitempty"`
	Floor          string                 `protobuf:"bytes,2,opt,name=Floor,proto3" json:"Floor,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GetRoomInfoRequest) Reset() {
	*x = GetRoomInfoRequest{}
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoomInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomInfoRequest) ProtoMessage() {}

func (x *GetRoomInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomInfoRequest.ProtoReflect.Descriptor instead.
func (*GetRoomInfoRequest) Descriptor() ([]byte, []int) {
	return file_elecprice_v1_elecprice_proto_rawDescGZIP(), []int{2}
}

func (x *GetRoomInfoRequest) GetArchitectureID() string {
	if x != nil {
		return x.ArchitectureID
	}
	return ""
}

func (x *GetRoomInfoRequest) GetFloor() string {
	if x != nil {
		return x.Floor
	}
	return ""
}

type GetRoomInfoResponse struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	RoomList      []*GetRoomInfoResponse_Room `protobuf:"bytes,1,rep,name=RoomList,proto3" json:"RoomList,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRoomInfoResponse) Reset() {
	*x = GetRoomInfoResponse{}
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoomInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomInfoResponse) ProtoMessage() {}

func (x *GetRoomInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomInfoResponse.ProtoReflect.Descriptor instead.
func (*GetRoomInfoResponse) Descriptor() ([]byte, []int) {
	return file_elecprice_v1_elecprice_proto_rawDescGZIP(), []int{3}
}

func (x *GetRoomInfoResponse) GetRoomList() []*GetRoomInfoResponse_Room {
	if x != nil {
		return x.RoomList
	}
	return nil
}

type GetPriceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPriceRequest) Reset() {
	*x = GetPriceRequest{}
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPriceRequest) ProtoMessage() {}

func (x *GetPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPriceRequest.ProtoReflect.Descriptor instead.
func (*GetPriceRequest) Descriptor() ([]byte, []int) {
	return file_elecprice_v1_elecprice_proto_rawDescGZIP(), []int{4}
}

func (x *GetPriceRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type GetPriceResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Price         *GetPriceResponse_Price `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPriceResponse) Reset() {
	*x = GetPriceResponse{}
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPriceResponse) ProtoMessage() {}

func (x *GetPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPriceResponse.ProtoReflect.Descriptor instead.
func (*GetPriceResponse) Descriptor() ([]byte, []int) {
	return file_elecprice_v1_elecprice_proto_rawDescGZIP(), []int{5}
}

func (x *GetPriceResponse) GetPrice() *GetPriceResponse_Price {
	if x != nil {
		return x.Price
	}
	return nil
}

// 有一个id对应两个状态和一个状态
type Standard struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         int64                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	RoomId        string                 `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RoomName      string                 `protobuf:"bytes,3,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Standard) Reset() {
	*x = Standard{}
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Standard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Standard) ProtoMessage() {}

func (x *Standard) ProtoReflect() protoreflect.Message {
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Standard.ProtoReflect.Descriptor instead.
func (*Standard) Descriptor() ([]byte, []int) {
	return file_elecprice_v1_elecprice_proto_rawDescGZIP(), []int{6}
}

func (x *Standard) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Standard) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *Standard) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

type SetStandardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StudentId     string                 `protobuf:"bytes,1,opt,name=studentId,proto3" json:"studentId,omitempty"`
	Standard      *Standard              `protobuf:"bytes,2,opt,name=standard,proto3" json:"standard,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetStandardRequest) Reset() {
	*x = SetStandardRequest{}
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetStandardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStandardRequest) ProtoMessage() {}

func (x *SetStandardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStandardRequest.ProtoReflect.Descriptor instead.
func (*SetStandardRequest) Descriptor() ([]byte, []int) {
	return file_elecprice_v1_elecprice_proto_rawDescGZIP(), []int{7}
}

func (x *SetStandardRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *SetStandardRequest) GetStandard() *Standard {
	if x != nil {
		return x.Standard
	}
	return nil
}

type SetStandardResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetStandardResponse) Reset() {
	*x = SetStandardResponse{}
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetStandardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStandardResponse) ProtoMessage() {}

func (x *SetStandardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStandardResponse.ProtoReflect.Descriptor instead.
func (*SetStandardResponse) Descriptor() ([]byte, []int) {
	return file_elecprice_v1_elecprice_proto_rawDescGZIP(), []int{8}
}

type GetStandardListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StudentId     string                 `protobuf:"bytes,1,opt,name=studentId,proto3" json:"studentId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStandardListRequest) Reset() {
	*x = GetStandardListRequest{}
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStandardListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStandardListRequest) ProtoMessage() {}

func (x *GetStandardListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStandardListRequest.ProtoReflect.Descriptor instead.
func (*GetStandardListRequest) Descriptor() ([]byte, []int) {
	return file_elecprice_v1_elecprice_proto_rawDescGZIP(), []int{9}
}

func (x *GetStandardListRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

type GetStandardListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Standards     []*Standard            `protobuf:"bytes,1,rep,name=standards,proto3" json:"standards,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStandardListResponse) Reset() {
	*x = GetStandardListResponse{}
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStandardListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStandardListResponse) ProtoMessage() {}

func (x *GetStandardListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStandardListResponse.ProtoReflect.Descriptor instead.
func (*GetStandardListResponse) Descriptor() ([]byte, []int) {
	return file_elecprice_v1_elecprice_proto_rawDescGZIP(), []int{10}
}

func (x *GetStandardListResponse) GetStandards() []*Standard {
	if x != nil {
		return x.Standards
	}
	return nil
}

type CancelStandardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StudentId     string                 `protobuf:"bytes,1,opt,name=studentId,proto3" json:"studentId,omitempty"`
	RoomId        string                 `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelStandardRequest) Reset() {
	*x = CancelStandardRequest{}
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelStandardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelStandardRequest) ProtoMessage() {}

func (x *CancelStandardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelStandardRequest.ProtoReflect.Descriptor instead.
func (*CancelStandardRequest) Descriptor() ([]byte, []int) {
	return file_elecprice_v1_elecprice_proto_rawDescGZIP(), []int{11}
}

func (x *CancelStandardRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *CancelStandardRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type CancelStandardResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelStandardResponse) Reset() {
	*x = CancelStandardResponse{}
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelStandardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelStandardResponse) ProtoMessage() {}

func (x *CancelStandardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelStandardResponse.ProtoReflect.Descriptor instead.
func (*CancelStandardResponse) Descriptor() ([]byte, []int) {
	return file_elecprice_v1_elecprice_proto_rawDescGZIP(), []int{12}
}

type GetArchitectureResponse_Architecture struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	ArchitectureID   string                 `protobuf:"bytes,1,opt,name=ArchitectureID,proto3" json:"ArchitectureID,omitempty"`
	ArchitectureName string                 `protobuf:"bytes,2,opt,name=ArchitectureName,proto3" json:"ArchitectureName,omitempty"`
	BaseFloor        string                 `protobuf:"bytes,3,opt,name=BaseFloor,proto3" json:"BaseFloor,omitempty"`
	TopFloor         string                 `protobuf:"bytes,4,opt,name=TopFloor,proto3" json:"TopFloor,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *GetArchitectureResponse_Architecture) Reset() {
	*x = GetArchitectureResponse_Architecture{}
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetArchitectureResponse_Architecture) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArchitectureResponse_Architecture) ProtoMessage() {}

func (x *GetArchitectureResponse_Architecture) ProtoReflect() protoreflect.Message {
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArchitectureResponse_Architecture.ProtoReflect.Descriptor instead.
func (*GetArchitectureResponse_Architecture) Descriptor() ([]byte, []int) {
	return file_elecprice_v1_elecprice_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetArchitectureResponse_Architecture) GetArchitectureID() string {
	if x != nil {
		return x.ArchitectureID
	}
	return ""
}

func (x *GetArchitectureResponse_Architecture) GetArchitectureName() string {
	if x != nil {
		return x.ArchitectureName
	}
	return ""
}

func (x *GetArchitectureResponse_Architecture) GetBaseFloor() string {
	if x != nil {
		return x.BaseFloor
	}
	return ""
}

func (x *GetArchitectureResponse_Architecture) GetTopFloor() string {
	if x != nil {
		return x.TopFloor
	}
	return ""
}

type GetRoomInfoResponse_Room struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomName      string                 `protobuf:"bytes,1,opt,name=RoomName,proto3" json:"RoomName,omitempty"`
	RoomID        string                 `protobuf:"bytes,2,opt,name=RoomID,proto3" json:"RoomID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRoomInfoResponse_Room) Reset() {
	*x = GetRoomInfoResponse_Room{}
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoomInfoResponse_Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomInfoResponse_Room) ProtoMessage() {}

func (x *GetRoomInfoResponse_Room) ProtoReflect() protoreflect.Message {
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomInfoResponse_Room.ProtoReflect.Descriptor instead.
func (*GetRoomInfoResponse_Room) Descriptor() ([]byte, []int) {
	return file_elecprice_v1_elecprice_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetRoomInfoResponse_Room) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *GetRoomInfoResponse_Room) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

type GetPriceResponse_Price struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	RemainMoney       string                 `protobuf:"bytes,1,opt,name=RemainMoney,proto3" json:"RemainMoney,omitempty"`             // 剩余电费
	YesterdayUseValue string                 `protobuf:"bytes,2,opt,name=YesterdayUseValue,proto3" json:"YesterdayUseValue,omitempty"` // 昨日花费电量
	YesterdayUseMoney string                 `protobuf:"bytes,3,opt,name=YesterdayUseMoney,proto3" json:"YesterdayUseMoney,omitempty"` // 昨日花费电费
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetPriceResponse_Price) Reset() {
	*x = GetPriceResponse_Price{}
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPriceResponse_Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPriceResponse_Price) ProtoMessage() {}

func (x *GetPriceResponse_Price) ProtoReflect() protoreflect.Message {
	mi := &file_elecprice_v1_elecprice_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPriceResponse_Price.ProtoReflect.Descriptor instead.
func (*GetPriceResponse_Price) Descriptor() ([]byte, []int) {
	return file_elecprice_v1_elecprice_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetPriceResponse_Price) GetRemainMoney() string {
	if x != nil {
		return x.RemainMoney
	}
	return ""
}

func (x *GetPriceResponse_Price) GetYesterdayUseValue() string {
	if x != nil {
		return x.YesterdayUseValue
	}
	return ""
}

func (x *GetPriceResponse_Price) GetYesterdayUseMoney() string {
	if x != nil {
		return x.YesterdayUseMoney
	}
	return ""
}

var File_elecprice_v1_elecprice_proto protoreflect.FileDescriptor

const file_elecprice_v1_elecprice_proto_rawDesc = "" +
	"\n" +
	"\x1celecprice/v1/elecprice.proto\x12\felecprice.v1\"4\n" +
	"\x16GetArchitectureRequest\x12\x1a\n" +
	"\bAreaName\x18\x01 \x01(\tR\bAreaName\"\x98\x02\n" +
	"\x17GetArchitectureResponse\x12^\n" +
	"\x10ArchitectureList\x18\x01 \x03(\v22.elecprice.v1.GetArchitectureResponse.ArchitectureR\x10ArchitectureList\x1a\x9c\x01\n" +
	"\fArchitecture\x12&\n" +
	"\x0eArchitectureID\x18\x01 \x01(\tR\x0eArchitectureID\x12*\n" +
	"\x10ArchitectureName\x18\x02 \x01(\tR\x10ArchitectureName\x12\x1c\n" +
	"\tBaseFloor\x18\x03 \x01(\tR\tBaseFloor\x12\x1a\n" +
	"\bTopFloor\x18\x04 \x01(\tR\bTopFloor\"R\n" +
	"\x12GetRoomInfoRequest\x12&\n" +
	"\x0eArchitectureID\x18\x01 \x01(\tR\x0eArchitectureID\x12\x14\n" +
	"\x05Floor\x18\x02 \x01(\tR\x05Floor\"\x95\x01\n" +
	"\x13GetRoomInfoResponse\x12B\n" +
	"\bRoomList\x18\x01 \x03(\v2&.elecprice.v1.GetRoomInfoResponse.RoomR\bRoomList\x1a:\n" +
	"\x04Room\x12\x1a\n" +
	"\bRoomName\x18\x01 \x01(\tR\bRoomName\x12\x16\n" +
	"\x06RoomID\x18\x02 \x01(\tR\x06RoomID\"*\n" +
	"\x0fGetPriceRequest\x12\x17\n" +
	"\aroom_id\x18\x01 \x01(\tR\x06roomId\"\xd6\x01\n" +
	"\x10GetPriceResponse\x12:\n" +
	"\x05price\x18\x01 \x01(\v2$.elecprice.v1.GetPriceResponse.PriceR\x05price\x1a\x85\x01\n" +
	"\x05Price\x12 \n" +
	"\vRemainMoney\x18\x01 \x01(\tR\vRemainMoney\x12,\n" +
	"\x11YesterdayUseValue\x18\x02 \x01(\tR\x11YesterdayUseValue\x12,\n" +
	"\x11YesterdayUseMoney\x18\x03 \x01(\tR\x11YesterdayUseMoney\"V\n" +
	"\bStandard\x12\x14\n" +
	"\x05limit\x18\x01 \x01(\x03R\x05limit\x12\x17\n" +
	"\aroom_id\x18\x02 \x01(\tR\x06roomId\x12\x1b\n" +
	"\troom_name\x18\x03 \x01(\tR\broomName\"f\n" +
	"\x12SetStandardRequest\x12\x1c\n" +
	"\tstudentId\x18\x01 \x01(\tR\tstudentId\x122\n" +
	"\bstandard\x18\x02 \x01(\v2\x16.elecprice.v1.StandardR\bstandard\"\x15\n" +
	"\x13SetStandardResponse\"6\n" +
	"\x16GetStandardListRequest\x12\x1c\n" +
	"\tstudentId\x18\x01 \x01(\tR\tstudentId\"O\n" +
	"\x17GetStandardListResponse\x124\n" +
	"\tstandards\x18\x01 \x03(\v2\x16.elecprice.v1.StandardR\tstandards\"N\n" +
	"\x15CancelStandardRequest\x12\x1c\n" +
	"\tstudentId\x18\x01 \x01(\tR\tstudentId\x12\x17\n" +
	"\aroom_id\x18\x02 \x01(\tR\x06roomId\"\x18\n" +
	"\x16CancelStandardResponse2\xa2\x04\n" +
	"\x10ElecpriceService\x12^\n" +
	"\x0fGetArchitecture\x12$.elecprice.v1.GetArchitectureRequest\x1a%.elecprice.v1.GetArchitectureResponse\x12R\n" +
	"\vGetRoomInfo\x12 .elecprice.v1.GetRoomInfoRequest\x1a!.elecprice.v1.GetRoomInfoResponse\x12I\n" +
	"\bGetPrice\x12\x1d.elecprice.v1.GetPriceRequest\x1a\x1e.elecprice.v1.GetPriceResponse\x12R\n" +
	"\vSetStandard\x12 .elecprice.v1.SetStandardRequest\x1a!.elecprice.v1.SetStandardResponse\x12^\n" +
	"\x0fGetStandardList\x12$.elecprice.v1.GetStandardListRequest\x1a%.elecprice.v1.GetStandardListResponse\x12[\n" +
	"\x0eCancelStandard\x12#.elecprice.v1.CancelStandardRequest\x1a$.elecprice.v1.CancelStandardResponseBJZHgithub.com/asynccnu/ccnubox-be/be-api/gen/proto/elecprice/v1;elecpricev1b\x06proto3"

var (
	file_elecprice_v1_elecprice_proto_rawDescOnce sync.Once
	file_elecprice_v1_elecprice_proto_rawDescData []byte
)

func file_elecprice_v1_elecprice_proto_rawDescGZIP() []byte {
	file_elecprice_v1_elecprice_proto_rawDescOnce.Do(func() {
		file_elecprice_v1_elecprice_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_elecprice_v1_elecprice_proto_rawDesc), len(file_elecprice_v1_elecprice_proto_rawDesc)))
	})
	return file_elecprice_v1_elecprice_proto_rawDescData
}

var file_elecprice_v1_elecprice_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_elecprice_v1_elecprice_proto_goTypes = []any{
	(*GetArchitectureRequest)(nil),               // 0: elecprice.v1.GetArchitectureRequest
	(*GetArchitectureResponse)(nil),              // 1: elecprice.v1.GetArchitectureResponse
	(*GetRoomInfoRequest)(nil),                   // 2: elecprice.v1.GetRoomInfoRequest
	(*GetRoomInfoResponse)(nil),                  // 3: elecprice.v1.GetRoomInfoResponse
	(*GetPriceRequest)(nil),                      // 4: elecprice.v1.GetPriceRequest
	(*GetPriceResponse)(nil),                     // 5: elecprice.v1.GetPriceResponse
	(*Standard)(nil),                             // 6: elecprice.v1.Standard
	(*SetStandardRequest)(nil),                   // 7: elecprice.v1.SetStandardRequest
	(*SetStandardResponse)(nil),                  // 8: elecprice.v1.SetStandardResponse
	(*GetStandardListRequest)(nil),               // 9: elecprice.v1.GetStandardListRequest
	(*GetStandardListResponse)(nil),              // 10: elecprice.v1.GetStandardListResponse
	(*CancelStandardRequest)(nil),                // 11: elecprice.v1.CancelStandardRequest
	(*CancelStandardResponse)(nil),               // 12: elecprice.v1.CancelStandardResponse
	(*GetArchitectureResponse_Architecture)(nil), // 13: elecprice.v1.GetArchitectureResponse.Architecture
	(*GetRoomInfoResponse_Room)(nil),             // 14: elecprice.v1.GetRoomInfoResponse.Room
	(*GetPriceResponse_Price)(nil),               // 15: elecprice.v1.GetPriceResponse.Price
}
var file_elecprice_v1_elecprice_proto_depIdxs = []int32{
	13, // 0: elecprice.v1.GetArchitectureResponse.ArchitectureList:type_name -> elecprice.v1.GetArchitectureResponse.Architecture
	14, // 1: elecprice.v1.GetRoomInfoResponse.RoomList:type_name -> elecprice.v1.GetRoomInfoResponse.Room
	15, // 2: elecprice.v1.GetPriceResponse.price:type_name -> elecprice.v1.GetPriceResponse.Price
	6,  // 3: elecprice.v1.SetStandardRequest.standard:type_name -> elecprice.v1.Standard
	6,  // 4: elecprice.v1.GetStandardListResponse.standards:type_name -> elecprice.v1.Standard
	0,  // 5: elecprice.v1.ElecpriceService.GetArchitecture:input_type -> elecprice.v1.GetArchitectureRequest
	2,  // 6: elecprice.v1.ElecpriceService.GetRoomInfo:input_type -> elecprice.v1.GetRoomInfoRequest
	4,  // 7: elecprice.v1.ElecpriceService.GetPrice:input_type -> elecprice.v1.GetPriceRequest
	7,  // 8: elecprice.v1.ElecpriceService.SetStandard:input_type -> elecprice.v1.SetStandardRequest
	9,  // 9: elecprice.v1.ElecpriceService.GetStandardList:input_type -> elecprice.v1.GetStandardListRequest
	11, // 10: elecprice.v1.ElecpriceService.CancelStandard:input_type -> elecprice.v1.CancelStandardRequest
	1,  // 11: elecprice.v1.ElecpriceService.GetArchitecture:output_type -> elecprice.v1.GetArchitectureResponse
	3,  // 12: elecprice.v1.ElecpriceService.GetRoomInfo:output_type -> elecprice.v1.GetRoomInfoResponse
	5,  // 13: elecprice.v1.ElecpriceService.GetPrice:output_type -> elecprice.v1.GetPriceResponse
	8,  // 14: elecprice.v1.ElecpriceService.SetStandard:output_type -> elecprice.v1.SetStandardResponse
	10, // 15: elecprice.v1.ElecpriceService.GetStandardList:output_type -> elecprice.v1.GetStandardListResponse
	12, // 16: elecprice.v1.ElecpriceService.CancelStandard:output_type -> elecprice.v1.CancelStandardResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_elecprice_v1_elecprice_proto_init() }
func file_elecprice_v1_elecprice_proto_init() {
	if File_elecprice_v1_elecprice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_elecprice_v1_elecprice_proto_rawDesc), len(file_elecprice_v1_elecprice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_elecprice_v1_elecprice_proto_goTypes,
		DependencyIndexes: file_elecprice_v1_elecprice_proto_depIdxs,
		MessageInfos:      file_elecprice_v1_elecprice_proto_msgTypes,
	}.Build()
	File_elecprice_v1_elecprice_proto = out.File
	file_elecprice_v1_elecprice_proto_goTypes = nil
	file_elecprice_v1_elecprice_proto_depIdxs = nil
}
