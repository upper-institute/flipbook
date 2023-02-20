// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api/v1/event.proto

package flipbookv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SortingKeyType int32

const (
	SortingKeyType_SORTING_KEY_INCREASING_SEQUENCE SortingKeyType = 0
	SortingKeyType_SORTING_KEY_ARBITRARY_NUMBER    SortingKeyType = 1
)

// Enum value maps for SortingKeyType.
var (
	SortingKeyType_name = map[int32]string{
		0: "SORTING_KEY_INCREASING_SEQUENCE",
		1: "SORTING_KEY_ARBITRARY_NUMBER",
	}
	SortingKeyType_value = map[string]int32{
		"SORTING_KEY_INCREASING_SEQUENCE": 0,
		"SORTING_KEY_ARBITRARY_NUMBER":    1,
	}
)

func (x SortingKeyType) Enum() *SortingKeyType {
	p := new(SortingKeyType)
	*p = x
	return p
}

func (x SortingKeyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortingKeyType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_event_proto_enumTypes[0].Descriptor()
}

func (SortingKeyType) Type() protoreflect.EnumType {
	return &file_api_v1_event_proto_enumTypes[0]
}

func (x SortingKeyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortingKeyType.Descriptor instead.
func (SortingKeyType) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_event_proto_rawDescGZIP(), []int{0}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This key is used to route events to partitions. It must be unique for the
	// source generating it. We recommend using any well known standard like UUIDs
	// with a meaningful representation for business domainds
	PartitionKey string `protobuf:"bytes,1,opt,name=partition_key,json=partitionKey,proto3" json:"partition_key,omitempty"`
	// This key is used to control two things:
	// 1. The payload stored in this exact sorting key (a sorting key can only have 1 payload)
	// 2. The sorting order for EventStore.Iterate()
	//
	// An event "id" is the combination of the partition_key+sorting_key
	SortingKey     int64          `protobuf:"varint,2,opt,name=sorting_key,json=sortingKey,proto3" json:"sorting_key,omitempty"`
	EventPayload   *anypb.Any     `protobuf:"bytes,3,opt,name=event_payload,json=eventPayload,proto3" json:"event_payload,omitempty"`
	SortingKeyType SortingKeyType `protobuf:"varint,4,opt,name=sorting_key_type,json=sortingKeyType,proto3,enum=flipbook.v1.SortingKeyType" json:"sorting_key_type,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_api_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetPartitionKey() string {
	if x != nil {
		return x.PartitionKey
	}
	return ""
}

func (x *Event) GetSortingKey() int64 {
	if x != nil {
		return x.SortingKey
	}
	return 0
}

func (x *Event) GetEventPayload() *anypb.Any {
	if x != nil {
		return x.EventPayload
	}
	return nil
}

func (x *Event) GetSortingKeyType() SortingKeyType {
	if x != nil {
		return x.SortingKeyType
	}
	return SortingKeyType_SORTING_KEY_INCREASING_SEQUENCE
}

type Commit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartitionKey string `protobuf:"bytes,1,opt,name=partition_key,json=partitionKey,proto3" json:"partition_key,omitempty"`
	SortingKey   int64  `protobuf:"varint,2,opt,name=sorting_key,json=sortingKey,proto3" json:"sorting_key,omitempty"`
}

func (x *Commit) Reset() {
	*x = Commit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commit) ProtoMessage() {}

func (x *Commit) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commit.ProtoReflect.Descriptor instead.
func (*Commit) Descriptor() ([]byte, []int) {
	return file_api_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *Commit) GetPartitionKey() string {
	if x != nil {
		return x.PartitionKey
	}
	return ""
}

func (x *Commit) GetSortingKey() int64 {
	if x != nil {
		return x.SortingKey
	}
	return 0
}

type Event_AppendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *Event_AppendRequest) Reset() {
	*x = Event_AppendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_AppendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_AppendRequest) ProtoMessage() {}

func (x *Event_AppendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_AppendRequest.ProtoReflect.Descriptor instead.
func (*Event_AppendRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_event_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Event_AppendRequest) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type Event_IterateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartitionKey string `protobuf:"bytes,1,opt,name=partition_key,json=partitionKey,proto3" json:"partition_key,omitempty"`
	Query        *Query `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	BatchSize    int64  `protobuf:"varint,3,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
}

func (x *Event_IterateRequest) Reset() {
	*x = Event_IterateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_IterateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_IterateRequest) ProtoMessage() {}

func (x *Event_IterateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_IterateRequest.ProtoReflect.Descriptor instead.
func (*Event_IterateRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_event_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Event_IterateRequest) GetPartitionKey() string {
	if x != nil {
		return x.PartitionKey
	}
	return ""
}

func (x *Event_IterateRequest) GetQuery() *Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *Event_IterateRequest) GetBatchSize() int64 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

type Event_GetLatestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartitionKey string `protobuf:"bytes,1,opt,name=partition_key,json=partitionKey,proto3" json:"partition_key,omitempty"`
}

func (x *Event_GetLatestRequest) Reset() {
	*x = Event_GetLatestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_GetLatestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_GetLatestRequest) ProtoMessage() {}

func (x *Event_GetLatestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_GetLatestRequest.ProtoReflect.Descriptor instead.
func (*Event_GetLatestRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_event_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Event_GetLatestRequest) GetPartitionKey() string {
	if x != nil {
		return x.PartitionKey
	}
	return ""
}

var File_api_v1_event_proto protoreflect.FileDescriptor

var file_api_v1_event_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc5, 0x03, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65,
	0x79, 0x12, 0x39, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0c,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x45, 0x0a, 0x10,
	0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0e, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x1a, 0x7e, 0x0a, 0x0e, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65,
	0x1a, 0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x22, 0x4e, 0x0a, 0x06, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x2a, 0x57, 0x0a, 0x0e, 0x53, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x53,
	0x4f, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x49, 0x4e, 0x43, 0x52, 0x45,
	0x41, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x00,
	0x12, 0x20, 0x0a, 0x1c, 0x53, 0x4f, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4b, 0x45, 0x59, 0x5f,
	0x41, 0x52, 0x42, 0x49, 0x54, 0x52, 0x41, 0x52, 0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52,
	0x10, 0x01, 0x42, 0xaa, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x75, 0x70, 0x70, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65,
	0x2f, 0x66, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x6c, 0x69, 0x70, 0x62, 0x6f,
	0x6f, 0x6b, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x46, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x46, 0x6c, 0x69,
	0x70, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x46, 0x6c, 0x69, 0x70, 0x62,
	0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x46, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x6f,
	0x6b, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0c, 0x46, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x6f, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_event_proto_rawDescOnce sync.Once
	file_api_v1_event_proto_rawDescData = file_api_v1_event_proto_rawDesc
)

func file_api_v1_event_proto_rawDescGZIP() []byte {
	file_api_v1_event_proto_rawDescOnce.Do(func() {
		file_api_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_event_proto_rawDescData)
	})
	return file_api_v1_event_proto_rawDescData
}

var file_api_v1_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_event_proto_goTypes = []interface{}{
	(SortingKeyType)(0),            // 0: flipbook.v1.SortingKeyType
	(*Event)(nil),                  // 1: flipbook.v1.Event
	(*Commit)(nil),                 // 2: flipbook.v1.Commit
	(*Event_AppendRequest)(nil),    // 3: flipbook.v1.Event.AppendRequest
	(*Event_IterateRequest)(nil),   // 4: flipbook.v1.Event.IterateRequest
	(*Event_GetLatestRequest)(nil), // 5: flipbook.v1.Event.GetLatestRequest
	(*anypb.Any)(nil),              // 6: google.protobuf.Any
	(*Query)(nil),                  // 7: flipbook.v1.Query
}
var file_api_v1_event_proto_depIdxs = []int32{
	6, // 0: flipbook.v1.Event.event_payload:type_name -> google.protobuf.Any
	0, // 1: flipbook.v1.Event.sorting_key_type:type_name -> flipbook.v1.SortingKeyType
	1, // 2: flipbook.v1.Event.AppendRequest.events:type_name -> flipbook.v1.Event
	7, // 3: flipbook.v1.Event.IterateRequest.query:type_name -> flipbook.v1.Query
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_v1_event_proto_init() }
func file_api_v1_event_proto_init() {
	if File_api_v1_event_proto != nil {
		return
	}
	file_api_v1_query_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_api_v1_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commit); i {
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
		file_api_v1_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_AppendRequest); i {
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
		file_api_v1_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_IterateRequest); i {
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
		file_api_v1_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_GetLatestRequest); i {
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
			RawDescriptor: file_api_v1_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_event_proto_goTypes,
		DependencyIndexes: file_api_v1_event_proto_depIdxs,
		EnumInfos:         file_api_v1_event_proto_enumTypes,
		MessageInfos:      file_api_v1_event_proto_msgTypes,
	}.Build()
	File_api_v1_event_proto = out.File
	file_api_v1_event_proto_rawDesc = nil
	file_api_v1_event_proto_goTypes = nil
	file_api_v1_event_proto_depIdxs = nil
}
