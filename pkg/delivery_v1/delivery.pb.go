// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: delivery.proto

package delivery_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SetDeliveryStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`                  // Status to be set.
	OrderId string `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"` // ID of order to be udpated.
}

func (x *SetDeliveryStatusRequest) Reset() {
	*x = SetDeliveryStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDeliveryStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDeliveryStatusRequest) ProtoMessage() {}

func (x *SetDeliveryStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDeliveryStatusRequest.ProtoReflect.Descriptor instead.
func (*SetDeliveryStatusRequest) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{0}
}

func (x *SetDeliveryStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SetDeliveryStatusRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

var File_delivery_proto protoreflect.FileDescriptor

var file_delivery_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x18, 0x53, 0x65, 0x74, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x32, 0x5b, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x12, 0x4f, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x67, 0x6f, 0x72, 0x74, 0x6f, 0x69, 0x67, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x2f,
	0x73, 0x74, 0x75, 0x70, 0x65, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x62, 0x65, 0x6c, 0x6c, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x76, 0x31, 0x3b, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_delivery_proto_rawDescOnce sync.Once
	file_delivery_proto_rawDescData = file_delivery_proto_rawDesc
)

func file_delivery_proto_rawDescGZIP() []byte {
	file_delivery_proto_rawDescOnce.Do(func() {
		file_delivery_proto_rawDescData = protoimpl.X.CompressGZIP(file_delivery_proto_rawDescData)
	})
	return file_delivery_proto_rawDescData
}

var file_delivery_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_delivery_proto_goTypes = []any{
	(*SetDeliveryStatusRequest)(nil), // 0: delivery.SetDeliveryStatusRequest
	(*emptypb.Empty)(nil),            // 1: google.protobuf.Empty
}
var file_delivery_proto_depIdxs = []int32{
	0, // 0: delivery.Delivery.SetDeliveryStatus:input_type -> delivery.SetDeliveryStatusRequest
	1, // 1: delivery.Delivery.SetDeliveryStatus:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_delivery_proto_init() }
func file_delivery_proto_init() {
	if File_delivery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_delivery_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SetDeliveryStatusRequest); i {
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
			RawDescriptor: file_delivery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_delivery_proto_goTypes,
		DependencyIndexes: file_delivery_proto_depIdxs,
		MessageInfos:      file_delivery_proto_msgTypes,
	}.Build()
	File_delivery_proto = out.File
	file_delivery_proto_rawDesc = nil
	file_delivery_proto_goTypes = nil
	file_delivery_proto_depIdxs = nil
}