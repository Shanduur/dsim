// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: credentials_message.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserKey []byte `protobuf:"bytes,2,opt,name=user_key,json=userKey,proto3" json:"user_key,omitempty"`
}

func (x *Credentials) Reset() {
	*x = Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credentials_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credentials) ProtoMessage() {}

func (x *Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_credentials_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credentials.ProtoReflect.Descriptor instead.
func (*Credentials) Descriptor() ([]byte, []int) {
	return file_credentials_message_proto_rawDescGZIP(), []int{0}
}

func (x *Credentials) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Credentials) GetUserKey() []byte {
	if x != nil {
		return x.UserKey
	}
	return nil
}

var File_credentials_message_proto protoreflect.FileDescriptor

var file_credentials_message_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x6c, 0x75,
	0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x22,
	0x41, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4b,
	0x65, 0x79, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_credentials_message_proto_rawDescOnce sync.Once
	file_credentials_message_proto_rawDescData = file_credentials_message_proto_rawDesc
)

func file_credentials_message_proto_rawDescGZIP() []byte {
	file_credentials_message_proto_rawDescOnce.Do(func() {
		file_credentials_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_credentials_message_proto_rawDescData)
	})
	return file_credentials_message_proto_rawDescData
}

var file_credentials_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_credentials_message_proto_goTypes = []interface{}{
	(*Credentials)(nil), // 0: pluggable.transfer.Credentials
}
var file_credentials_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_credentials_message_proto_init() }
func file_credentials_message_proto_init() {
	if File_credentials_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_credentials_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credentials); i {
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
			RawDescriptor: file_credentials_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_credentials_message_proto_goTypes,
		DependencyIndexes: file_credentials_message_proto_depIdxs,
		MessageInfos:      file_credentials_message_proto_msgTypes,
	}.Build()
	File_credentials_message_proto = out.File
	file_credentials_message_proto_rawDesc = nil
	file_credentials_message_proto_goTypes = nil
	file_credentials_message_proto_depIdxs = nil
}
