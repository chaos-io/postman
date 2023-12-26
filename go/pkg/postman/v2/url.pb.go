// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: postman/v2/url.proto

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

type Url struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Raw      string            `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	Protocol string            `protobuf:"bytes,10,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Host     []string          `protobuf:"bytes,15,rep,name=host,proto3" json:"host,omitempty"`
	Path     []string          `protobuf:"bytes,18,rep,name=path,proto3" json:"path,omitempty"`
	Port     string            `protobuf:"bytes,20,opt,name=port,proto3" json:"port,omitempty"`
	Query    []*Url_QueryParam `protobuf:"bytes,21,rep,name=query,proto3" json:"query,omitempty"`
	Hash     string            `protobuf:"bytes,22,opt,name=hash,proto3" json:"hash,omitempty"`
	Variable []*Variable       `protobuf:"bytes,30,rep,name=variable,proto3" json:"variable,omitempty"`
}

func (x *Url) Reset() {
	*x = Url{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postman_v2_url_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Url) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Url) ProtoMessage() {}

func (x *Url) ProtoReflect() protoreflect.Message {
	mi := &file_postman_v2_url_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Url.ProtoReflect.Descriptor instead.
func (*Url) Descriptor() ([]byte, []int) {
	return file_postman_v2_url_proto_rawDescGZIP(), []int{0}
}

func (x *Url) GetRaw() string {
	if x != nil {
		return x.Raw
	}
	return ""
}

func (x *Url) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Url) GetHost() []string {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *Url) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *Url) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *Url) GetQuery() []*Url_QueryParam {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *Url) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Url) GetVariable() []*Variable {
	if x != nil {
		return x.Variable
	}
	return nil
}

type Url_QueryParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string       `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value       string       `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Disable     bool         `protobuf:"varint,5,opt,name=disable,proto3" json:"disable,omitempty"`
	Description *Description `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Url_QueryParam) Reset() {
	*x = Url_QueryParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postman_v2_url_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Url_QueryParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Url_QueryParam) ProtoMessage() {}

func (x *Url_QueryParam) ProtoReflect() protoreflect.Message {
	mi := &file_postman_v2_url_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Url_QueryParam.ProtoReflect.Descriptor instead.
func (*Url_QueryParam) Descriptor() ([]byte, []int) {
	return file_postman_v2_url_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Url_QueryParam) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Url_QueryParam) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Url_QueryParam) GetDisable() bool {
	if x != nil {
		return x.Disable
	}
	return false
}

func (x *Url_QueryParam) GetDescription() *Description {
	if x != nil {
		return x.Description
	}
	return nil
}

var File_postman_v2_url_proto protoreflect.FileDescriptor

var file_postman_v2_url_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x72, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2e,
	0x76, 0x32, 0x1a, 0x1c, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x02, 0x0a, 0x03,
	0x55, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x12, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x30, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x72, 0x6c, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x12, 0x30, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x1e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2e,
	0x76, 0x32, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x08, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x89, 0x01, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x4f, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2e, 0x76, 0x32, 0x42,
	0x08, 0x55, 0x72, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2d, 0x69, 0x6f,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2f, 0x76, 0x32, 0x3b, 0x70, 0x6f, 0x73, 0x74, 0x6d,
	0x61, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_postman_v2_url_proto_rawDescOnce sync.Once
	file_postman_v2_url_proto_rawDescData = file_postman_v2_url_proto_rawDesc
)

func file_postman_v2_url_proto_rawDescGZIP() []byte {
	file_postman_v2_url_proto_rawDescOnce.Do(func() {
		file_postman_v2_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_postman_v2_url_proto_rawDescData)
	})
	return file_postman_v2_url_proto_rawDescData
}

var file_postman_v2_url_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_postman_v2_url_proto_goTypes = []interface{}{
	(*Url)(nil),            // 0: postman.v2.Url
	(*Url_QueryParam)(nil), // 1: postman.v2.Url.QueryParam
	(*Variable)(nil),       // 2: postman.v2.Variable
	(*Description)(nil),    // 3: postman.v2.Description
}
var file_postman_v2_url_proto_depIdxs = []int32{
	1, // 0: postman.v2.Url.query:type_name -> postman.v2.Url.QueryParam
	2, // 1: postman.v2.Url.variable:type_name -> postman.v2.Variable
	3, // 2: postman.v2.Url.QueryParam.description:type_name -> postman.v2.Description
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_postman_v2_url_proto_init() }
func file_postman_v2_url_proto_init() {
	if File_postman_v2_url_proto != nil {
		return
	}
	file_postman_v2_description_proto_init()
	file_postman_v2_variable_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_postman_v2_url_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Url); i {
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
		file_postman_v2_url_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Url_QueryParam); i {
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
			RawDescriptor: file_postman_v2_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_postman_v2_url_proto_goTypes,
		DependencyIndexes: file_postman_v2_url_proto_depIdxs,
		MessageInfos:      file_postman_v2_url_proto_msgTypes,
	}.Build()
	File_postman_v2_url_proto = out.File
	file_postman_v2_url_proto_rawDesc = nil
	file_postman_v2_url_proto_goTypes = nil
	file_postman_v2_url_proto_depIdxs = nil
}
