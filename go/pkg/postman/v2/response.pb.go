// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: postman/v2/response.proto

package postman

import (
	core "github.com/mojo-lang/core/go/pkg/mojo/core"
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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OriginalRequest *Request        `protobuf:"bytes,10,opt,name=original_request,json=originalRequest,proto3" json:"originalRequest,omitempty"`
	ResponseTime    *core.Timestamp `protobuf:"bytes,11,opt,name=response_time,json=responseTime,proto3" json:"responseTime,omitempty"`
	Timings         *core.Object    `protobuf:"bytes,12,opt,name=timings,proto3" json:"timings,omitempty"`
	Header          []string        `protobuf:"bytes,20,rep,name=header,proto3" json:"header,omitempty"`
	Body            string          `protobuf:"bytes,22,opt,name=body,proto3" json:"body,omitempty"`
	Status          string          `protobuf:"bytes,30,opt,name=status,proto3" json:"status,omitempty"`
	Code            int32           `protobuf:"varint,31,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postman_v2_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_postman_v2_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_postman_v2_response_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Response) GetOriginalRequest() *Request {
	if x != nil {
		return x.OriginalRequest
	}
	return nil
}

func (x *Response) GetResponseTime() *core.Timestamp {
	if x != nil {
		return x.ResponseTime
	}
	return nil
}

func (x *Response) GetTimings() *core.Object {
	if x != nil {
		return x.Timings
	}
	return nil
}

func (x *Response) GetHeader() []string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Response) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Response) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type Cookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain   string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Expires  string `protobuf:"bytes,5,opt,name=expires,proto3" json:"expires,omitempty"`
	MaxAge   string `protobuf:"bytes,10,opt,name=max_age,json=maxAge,proto3" json:"maxAge,omitempty"`
	HostOnly bool   `protobuf:"varint,11,opt,name=host_only,json=hostOnly,proto3" json:"hostOnly,omitempty"`
	HttpOnly bool   `protobuf:"varint,12,opt,name=http_only,json=httpOnly,proto3" json:"httpOnly,omitempty"`
	Name     string `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`
	Path     string `protobuf:"bytes,21,opt,name=path,proto3" json:"path,omitempty"`
	Secure   bool   `protobuf:"varint,25,opt,name=secure,proto3" json:"secure,omitempty"`
	Session  bool   `protobuf:"varint,26,opt,name=session,proto3" json:"session,omitempty"`
	Value    string `protobuf:"bytes,30,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Cookie) Reset() {
	*x = Cookie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postman_v2_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cookie) ProtoMessage() {}

func (x *Cookie) ProtoReflect() protoreflect.Message {
	mi := &file_postman_v2_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cookie.ProtoReflect.Descriptor instead.
func (*Cookie) Descriptor() ([]byte, []int) {
	return file_postman_v2_response_proto_rawDescGZIP(), []int{1}
}

func (x *Cookie) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Cookie) GetExpires() string {
	if x != nil {
		return x.Expires
	}
	return ""
}

func (x *Cookie) GetMaxAge() string {
	if x != nil {
		return x.MaxAge
	}
	return ""
}

func (x *Cookie) GetHostOnly() bool {
	if x != nil {
		return x.HostOnly
	}
	return false
}

func (x *Cookie) GetHttpOnly() bool {
	if x != nil {
		return x.HttpOnly
	}
	return false
}

func (x *Cookie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cookie) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Cookie) GetSecure() bool {
	if x != nil {
		return x.Secure
	}
	return false
}

func (x *Cookie) GetSession() bool {
	if x != nil {
		return x.Session
	}
	return false
}

func (x *Cookie) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_postman_v2_response_proto protoreflect.FileDescriptor

var file_postman_v2_response_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x6f, 0x73,
	0x74, 0x6d, 0x61, 0x6e, 0x2e, 0x76, 0x32, 0x1a, 0x14, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6d,
	0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2f, 0x76, 0x32,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a,
	0x02, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3e, 0x0a, 0x10, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2e,
	0x76, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0f, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x14, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x1f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xfd, 0x01, 0x0a, 0x06,
	0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f,
	0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18, 0x19, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x54, 0x0a, 0x0a, 0x70,
	0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2e, 0x76, 0x32, 0x42, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2d, 0x69, 0x6f, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2f, 0x76, 0x32, 0x3b, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_postman_v2_response_proto_rawDescOnce sync.Once
	file_postman_v2_response_proto_rawDescData = file_postman_v2_response_proto_rawDesc
)

func file_postman_v2_response_proto_rawDescGZIP() []byte {
	file_postman_v2_response_proto_rawDescOnce.Do(func() {
		file_postman_v2_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_postman_v2_response_proto_rawDescData)
	})
	return file_postman_v2_response_proto_rawDescData
}

var file_postman_v2_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_postman_v2_response_proto_goTypes = []interface{}{
	(*Response)(nil),       // 0: postman.v2.Response
	(*Cookie)(nil),         // 1: postman.v2.Cookie
	(*Request)(nil),        // 2: postman.v2.Request
	(*core.Timestamp)(nil), // 3: mojo.core.Timestamp
	(*core.Object)(nil),    // 4: mojo.core.Object
}
var file_postman_v2_response_proto_depIdxs = []int32{
	2, // 0: postman.v2.Response.original_request:type_name -> postman.v2.Request
	3, // 1: postman.v2.Response.response_time:type_name -> mojo.core.Timestamp
	4, // 2: postman.v2.Response.timings:type_name -> mojo.core.Object
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_postman_v2_response_proto_init() }
func file_postman_v2_response_proto_init() {
	if File_postman_v2_response_proto != nil {
		return
	}
	file_postman_v2_request_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_postman_v2_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_postman_v2_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cookie); i {
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
			RawDescriptor: file_postman_v2_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_postman_v2_response_proto_goTypes,
		DependencyIndexes: file_postman_v2_response_proto_depIdxs,
		MessageInfos:      file_postman_v2_response_proto_msgTypes,
	}.Build()
	File_postman_v2_response_proto = out.File
	file_postman_v2_response_proto_rawDesc = nil
	file_postman_v2_response_proto_goTypes = nil
	file_postman_v2_response_proto_depIdxs = nil
}
