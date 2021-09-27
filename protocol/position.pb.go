// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: protocol/position.proto

package protocol

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

type BuyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price []byte `protobuf:"bytes,1,req,name=price" json:"price,omitempty"`
}

func (x *BuyRequest) Reset() {
	*x = BuyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_position_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyRequest) ProtoMessage() {}

func (x *BuyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_position_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyRequest.ProtoReflect.Descriptor instead.
func (*BuyRequest) Descriptor() ([]byte, []int) {
	return file_protocol_position_proto_rawDescGZIP(), []int{0}
}

func (x *BuyRequest) GetPrice() []byte {
	if x != nil {
		return x.Price
	}
	return nil
}

type BuyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *string `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
}

func (x *BuyReply) Reset() {
	*x = BuyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_position_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyReply) ProtoMessage() {}

func (x *BuyReply) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_position_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyReply.ProtoReflect.Descriptor instead.
func (*BuyReply) Descriptor() ([]byte, []int) {
	return file_protocol_position_proto_rawDescGZIP(), []int{1}
}

func (x *BuyReply) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type SellRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price []byte `protobuf:"bytes,1,req,name=price" json:"price,omitempty"`
}

func (x *SellRequest) Reset() {
	*x = SellRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_position_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellRequest) ProtoMessage() {}

func (x *SellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_position_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellRequest.ProtoReflect.Descriptor instead.
func (*SellRequest) Descriptor() ([]byte, []int) {
	return file_protocol_position_proto_rawDescGZIP(), []int{2}
}

func (x *SellRequest) GetPrice() []byte {
	if x != nil {
		return x.Price
	}
	return nil
}

type SellReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *string `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
}

func (x *SellReply) Reset() {
	*x = SellReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_position_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellReply) ProtoMessage() {}

func (x *SellReply) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_position_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellReply.ProtoReflect.Descriptor instead.
func (*SellReply) Descriptor() ([]byte, []int) {
	return file_protocol_position_proto_rawDescGZIP(), []int{3}
}

func (x *SellReply) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_position_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_position_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_protocol_position_proto_rawDescGZIP(), []int{4}
}

type GetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ButchOfPrices []byte `protobuf:"bytes,1,req,name=butchOfPrices" json:"butchOfPrices,omitempty"`
}

func (x *GetReply) Reset() {
	*x = GetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_position_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReply) ProtoMessage() {}

func (x *GetReply) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_position_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReply.ProtoReflect.Descriptor instead.
func (*GetReply) Descriptor() ([]byte, []int) {
	return file_protocol_position_proto_rawDescGZIP(), []int{5}
}

func (x *GetReply) GetButchOfPrices() []byte {
	if x != nil {
		return x.ButchOfPrices
	}
	return nil
}

var File_protocol_position_proto protoreflect.FileDescriptor

var file_protocol_position_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x22, 0x0a, 0x0a, 0x42, 0x75, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x42, 0x75, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x23, 0x0a, 0x0b, 0x53, 0x65,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22,
	0x25, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x24, 0x0a, 0x0d, 0x62, 0x75, 0x74, 0x63, 0x68, 0x4f, 0x66, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x0d, 0x62, 0x75, 0x74, 0x63, 0x68, 0x4f, 0x66,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x32, 0x9b, 0x01, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x03, 0x42, 0x75,
	0x79, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x75, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x75, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x04, 0x53, 0x65, 0x6c, 0x6c, 0x12,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6c, 0x6c,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c,
}

var (
	file_protocol_position_proto_rawDescOnce sync.Once
	file_protocol_position_proto_rawDescData = file_protocol_position_proto_rawDesc
)

func file_protocol_position_proto_rawDescGZIP() []byte {
	file_protocol_position_proto_rawDescOnce.Do(func() {
		file_protocol_position_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_position_proto_rawDescData)
	})
	return file_protocol_position_proto_rawDescData
}

var file_protocol_position_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protocol_position_proto_goTypes = []interface{}{
	(*BuyRequest)(nil),  // 0: proto.BuyRequest
	(*BuyReply)(nil),    // 1: proto.BuyReply
	(*SellRequest)(nil), // 2: proto.SellRequest
	(*SellReply)(nil),   // 3: proto.SellReply
	(*GetRequest)(nil),  // 4: proto.GetRequest
	(*GetReply)(nil),    // 5: proto.GetReply
}
var file_protocol_position_proto_depIdxs = []int32{
	0, // 0: proto.PositionService.Buy:input_type -> proto.BuyRequest
	2, // 1: proto.PositionService.Sell:input_type -> proto.SellRequest
	4, // 2: proto.PositionService.Get:input_type -> proto.GetRequest
	1, // 3: proto.PositionService.Buy:output_type -> proto.BuyReply
	3, // 4: proto.PositionService.Sell:output_type -> proto.SellReply
	5, // 5: proto.PositionService.Get:output_type -> proto.GetReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protocol_position_proto_init() }
func file_protocol_position_proto_init() {
	if File_protocol_position_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_position_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyRequest); i {
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
		file_protocol_position_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyReply); i {
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
		file_protocol_position_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellRequest); i {
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
		file_protocol_position_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellReply); i {
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
		file_protocol_position_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_protocol_position_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReply); i {
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
			RawDescriptor: file_protocol_position_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_position_proto_goTypes,
		DependencyIndexes: file_protocol_position_proto_depIdxs,
		MessageInfos:      file_protocol_position_proto_msgTypes,
	}.Build()
	File_protocol_position_proto = out.File
	file_protocol_position_proto_rawDesc = nil
	file_protocol_position_proto_goTypes = nil
	file_protocol_position_proto_depIdxs = nil
}
