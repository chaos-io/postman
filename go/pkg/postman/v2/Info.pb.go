// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: postman/v2/Info.proto

package postman

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

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description *Description `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Version     *Version     `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	Schema      string       `protobuf:"bytes,10,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postman_v2_Info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_postman_v2_Info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_postman_v2_Info_proto_rawDescGZIP(), []int{0}
}

func (x *Info) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Info) GetDescription() *Description {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Info) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *Info) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major      int32  `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor      int32  `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch      int32  `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	Identifier string `protobuf:"bytes,10,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postman_v2_Info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_postman_v2_Info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_postman_v2_Info_proto_rawDescGZIP(), []int{1}
}

func (x *Version) GetMajor() int32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *Version) GetMinor() int32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *Version) GetPatch() int32 {
	if x != nil {
		return x.Patch
	}
	return 0
}

func (x *Version) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

var File_postman_v2_Info_proto protoreflect.FileDescriptor

var file_postman_v2_Info_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e,
	0x2e, 0x76, 0x32, 0x1a, 0x13, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2f, 0x76, 0x32, 0x2f,
	0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x6d, 0x61, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x6b, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x42, 0x50, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2e,
	0x76, 0x32, 0x42, 0x09, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6f,
	0x73, 0x2d, 0x69, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2f, 0x76, 0x32, 0x3b, 0x70,
	0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_postman_v2_Info_proto_rawDescOnce sync.Once
	file_postman_v2_Info_proto_rawDescData = file_postman_v2_Info_proto_rawDesc
)

func file_postman_v2_Info_proto_rawDescGZIP() []byte {
	file_postman_v2_Info_proto_rawDescOnce.Do(func() {
		file_postman_v2_Info_proto_rawDescData = protoimpl.X.CompressGZIP(file_postman_v2_Info_proto_rawDescData)
	})
	return file_postman_v2_Info_proto_rawDescData
}

var file_postman_v2_Info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_postman_v2_Info_proto_goTypes = []interface{}{
	(*Info)(nil),        // 0: postman.v2.Info
	(*Version)(nil),     // 1: postman.v2.Version
	(*Description)(nil), // 2: postman.v2.Description
}
var file_postman_v2_Info_proto_depIdxs = []int32{
	2, // 0: postman.v2.Info.description:type_name -> postman.v2.Description
	1, // 1: postman.v2.Info.version:type_name -> postman.v2.Version
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_postman_v2_Info_proto_init() }
func file_postman_v2_Info_proto_init() {
	if File_postman_v2_Info_proto != nil {
		return
	}
	file_postman_v2_v2_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_postman_v2_Info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_postman_v2_Info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
			RawDescriptor: file_postman_v2_Info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_postman_v2_Info_proto_goTypes,
		DependencyIndexes: file_postman_v2_Info_proto_depIdxs,
		MessageInfos:      file_postman_v2_Info_proto_msgTypes,
	}.Build()
	File_postman_v2_Info_proto = out.File
	file_postman_v2_Info_proto_rawDesc = nil
	file_postman_v2_Info_proto_goTypes = nil
	file_postman_v2_Info_proto_depIdxs = nil
}
