// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: api/v1/endpoint.proto

package flipbookv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EndpointStatus int32

const (
	EndpointStatus_ENDPOINT_STATUS_ACTIVE EndpointStatus = 0
	EndpointStatus_ENDPOINT_STATUS_HIDLE  EndpointStatus = 1
)

// Enum value maps for EndpointStatus.
var (
	EndpointStatus_name = map[int32]string{
		0: "ENDPOINT_STATUS_ACTIVE",
		1: "ENDPOINT_STATUS_HIDLE",
	}
	EndpointStatus_value = map[string]int32{
		"ENDPOINT_STATUS_ACTIVE": 0,
		"ENDPOINT_STATUS_HIDLE":  1,
	}
)

func (x EndpointStatus) Enum() *EndpointStatus {
	p := new(EndpointStatus)
	*p = x
	return p
}

func (x EndpointStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EndpointStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_endpoint_proto_enumTypes[0].Descriptor()
}

func (EndpointStatus) Type() protoreflect.EnumType {
	return &file_api_v1_endpoint_proto_enumTypes[0]
}

func (x EndpointStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EndpointStatus.Descriptor instead.
func (EndpointStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_endpoint_proto_rawDescGZIP(), []int{0}
}

type Endpoint_AddressFamily int32

const (
	Endpoint_ADDRESS_FAMILY_IPV4 Endpoint_AddressFamily = 0
	Endpoint_ADDRESS_FAMILY_IPV6 Endpoint_AddressFamily = 1
)

// Enum value maps for Endpoint_AddressFamily.
var (
	Endpoint_AddressFamily_name = map[int32]string{
		0: "ADDRESS_FAMILY_IPV4",
		1: "ADDRESS_FAMILY_IPV6",
	}
	Endpoint_AddressFamily_value = map[string]int32{
		"ADDRESS_FAMILY_IPV4": 0,
		"ADDRESS_FAMILY_IPV6": 1,
	}
)

func (x Endpoint_AddressFamily) Enum() *Endpoint_AddressFamily {
	p := new(Endpoint_AddressFamily)
	*p = x
	return p
}

func (x Endpoint_AddressFamily) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Endpoint_AddressFamily) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_endpoint_proto_enumTypes[1].Descriptor()
}

func (Endpoint_AddressFamily) Type() protoreflect.EnumType {
	return &file_api_v1_endpoint_proto_enumTypes[1]
}

func (x Endpoint_AddressFamily) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Endpoint_AddressFamily.Descriptor instead.
func (Endpoint_AddressFamily) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_endpoint_proto_rawDescGZIP(), []int{0, 0}
}

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port          uint32                 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	AddressFamily Endpoint_AddressFamily `protobuf:"varint,3,opt,name=address_family,json=addressFamily,proto3,enum=flipbook.v1.Endpoint_AddressFamily" json:"address_family,omitempty"`
	Status        EndpointStatus         `protobuf:"varint,4,opt,name=status,proto3,enum=flipbook.v1.EndpointStatus" json:"status,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_endpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_endpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_api_v1_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *Endpoint) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Endpoint) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Endpoint) GetAddressFamily() Endpoint_AddressFamily {
	if x != nil {
		return x.AddressFamily
	}
	return Endpoint_ADDRESS_FAMILY_IPV4
}

func (x *Endpoint) GetStatus() EndpointStatus {
	if x != nil {
		return x.Status
	}
	return EndpointStatus_ENDPOINT_STATUS_ACTIVE
}

func (x *Endpoint) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_api_v1_endpoint_proto protoreflect.FileDescriptor

var file_api_v1_endpoint_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x02, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x4a, 0x0a, 0x0e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x66, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x0d, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x33, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x66,
	0x6c, 0x69, 0x70, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x41, 0x0a, 0x0d,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x17, 0x0a,
	0x13, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x5f,
	0x49, 0x50, 0x56, 0x34, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53,
	0x53, 0x5f, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x49, 0x50, 0x56, 0x36, 0x10, 0x01, 0x2a,
	0x47, 0x0a, 0x0e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x19, 0x0a,
	0x15, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x48, 0x49, 0x44, 0x4c, 0x45, 0x10, 0x01, 0x42, 0xad, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x2e, 0x66, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x70, 0x70, 0x65, 0x72, 0x2d,
	0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x2f, 0x66, 0x6c, 0x69, 0x70, 0x62, 0x6f,
	0x6f, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x3b, 0x66, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x6f, 0x6b, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x46, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x46, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0b, 0x46, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x17, 0x46, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x46, 0x6c, 0x69, 0x70,
	0x62, 0x6f, 0x6f, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_endpoint_proto_rawDescOnce sync.Once
	file_api_v1_endpoint_proto_rawDescData = file_api_v1_endpoint_proto_rawDesc
)

func file_api_v1_endpoint_proto_rawDescGZIP() []byte {
	file_api_v1_endpoint_proto_rawDescOnce.Do(func() {
		file_api_v1_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_endpoint_proto_rawDescData)
	})
	return file_api_v1_endpoint_proto_rawDescData
}

var file_api_v1_endpoint_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_v1_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v1_endpoint_proto_goTypes = []interface{}{
	(EndpointStatus)(0),           // 0: flipbook.v1.EndpointStatus
	(Endpoint_AddressFamily)(0),   // 1: flipbook.v1.Endpoint.AddressFamily
	(*Endpoint)(nil),              // 2: flipbook.v1.Endpoint
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_api_v1_endpoint_proto_depIdxs = []int32{
	1, // 0: flipbook.v1.Endpoint.address_family:type_name -> flipbook.v1.Endpoint.AddressFamily
	0, // 1: flipbook.v1.Endpoint.status:type_name -> flipbook.v1.EndpointStatus
	3, // 2: flipbook.v1.Endpoint.updated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_v1_endpoint_proto_init() }
func file_api_v1_endpoint_proto_init() {
	if File_api_v1_endpoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_endpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
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
			RawDescriptor: file_api_v1_endpoint_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_endpoint_proto_goTypes,
		DependencyIndexes: file_api_v1_endpoint_proto_depIdxs,
		EnumInfos:         file_api_v1_endpoint_proto_enumTypes,
		MessageInfos:      file_api_v1_endpoint_proto_msgTypes,
	}.Build()
	File_api_v1_endpoint_proto = out.File
	file_api_v1_endpoint_proto_rawDesc = nil
	file_api_v1_endpoint_proto_goTypes = nil
	file_api_v1_endpoint_proto_depIdxs = nil
}
