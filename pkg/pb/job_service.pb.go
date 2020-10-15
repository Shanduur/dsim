// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: job_service.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type JobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*JobRequest_Job
	//	*JobRequest_ChunkData
	Data isJobRequest_Data `protobuf_oneof:"data"`
}

func (x *JobRequest) Reset() {
	*x = JobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobRequest) ProtoMessage() {}

func (x *JobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobRequest.ProtoReflect.Descriptor instead.
func (*JobRequest) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{0}
}

func (m *JobRequest) GetData() isJobRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *JobRequest) GetJob() *Job {
	if x, ok := x.GetData().(*JobRequest_Job); ok {
		return x.Job
	}
	return nil
}

func (x *JobRequest) GetChunkData() *Chunk {
	if x, ok := x.GetData().(*JobRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isJobRequest_Data interface {
	isJobRequest_Data()
}

type JobRequest_Job struct {
	Job *Job `protobuf:"bytes,1,opt,name=job,proto3,oneof"`
}

type JobRequest_ChunkData struct {
	ChunkData *Chunk `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*JobRequest_Job) isJobRequest_Data() {}

func (*JobRequest_ChunkData) isJobRequest_Data() {}

type JobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       []int64   `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
	Response *Response `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *JobResponse) Reset() {
	*x = JobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobResponse) ProtoMessage() {}

func (x *JobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobResponse.ProtoReflect.Descriptor instead.
func (*JobResponse) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{1}
}

func (x *JobResponse) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *JobResponse) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_job_service_proto protoreflect.FileDescriptor

var file_job_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x1a, 0x11, 0x6a, 0x6f, 0x62, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x48, 0x00, 0x52, 0x03, 0x6a,
	0x6f, 0x62, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x57, 0x0a, 0x0b, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x5e, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a,
	0x09, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x1e, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42,
	0x08, 0x5a, 0x06, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_job_service_proto_rawDescOnce sync.Once
	file_job_service_proto_rawDescData = file_job_service_proto_rawDesc
)

func file_job_service_proto_rawDescGZIP() []byte {
	file_job_service_proto_rawDescOnce.Do(func() {
		file_job_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_service_proto_rawDescData)
	})
	return file_job_service_proto_rawDescData
}

var file_job_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_job_service_proto_goTypes = []interface{}{
	(*JobRequest)(nil),  // 0: pluggable.transfer.JobRequest
	(*JobResponse)(nil), // 1: pluggable.transfer.JobResponse
	(*Job)(nil),         // 2: pluggable.transfer.Job
	(*Chunk)(nil),       // 3: pluggable.transfer.Chunk
	(*Response)(nil),    // 4: pluggable.transfer.Response
}
var file_job_service_proto_depIdxs = []int32{
	2, // 0: pluggable.transfer.JobRequest.job:type_name -> pluggable.transfer.Job
	3, // 1: pluggable.transfer.JobRequest.chunk_data:type_name -> pluggable.transfer.Chunk
	4, // 2: pluggable.transfer.JobResponse.response:type_name -> pluggable.transfer.Response
	0, // 3: pluggable.transfer.JobService.SubmitJob:input_type -> pluggable.transfer.JobRequest
	1, // 4: pluggable.transfer.JobService.SubmitJob:output_type -> pluggable.transfer.JobResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_job_service_proto_init() }
func file_job_service_proto_init() {
	if File_job_service_proto != nil {
		return
	}
	file_job_message_proto_init()
	file_response_message_proto_init()
	file_chunk_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_job_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobRequest); i {
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
		file_job_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobResponse); i {
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
	file_job_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*JobRequest_Job)(nil),
		(*JobRequest_ChunkData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_job_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_job_service_proto_goTypes,
		DependencyIndexes: file_job_service_proto_depIdxs,
		MessageInfos:      file_job_service_proto_msgTypes,
	}.Build()
	File_job_service_proto = out.File
	file_job_service_proto_rawDesc = nil
	file_job_service_proto_goTypes = nil
	file_job_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobServiceClient interface {
	SubmitJob(ctx context.Context, opts ...grpc.CallOption) (JobService_SubmitJobClient, error)
}

type jobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobServiceClient(cc grpc.ClientConnInterface) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) SubmitJob(ctx context.Context, opts ...grpc.CallOption) (JobService_SubmitJobClient, error) {
	stream, err := c.cc.NewStream(ctx, &_JobService_serviceDesc.Streams[0], "/pluggable.transfer.JobService/SubmitJob", opts...)
	if err != nil {
		return nil, err
	}
	x := &jobServiceSubmitJobClient{stream}
	return x, nil
}

type JobService_SubmitJobClient interface {
	Send(*JobRequest) error
	CloseAndRecv() (*JobResponse, error)
	grpc.ClientStream
}

type jobServiceSubmitJobClient struct {
	grpc.ClientStream
}

func (x *jobServiceSubmitJobClient) Send(m *JobRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *jobServiceSubmitJobClient) CloseAndRecv() (*JobResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(JobResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// JobServiceServer is the server API for JobService service.
type JobServiceServer interface {
	SubmitJob(JobService_SubmitJobServer) error
}

// UnimplementedJobServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJobServiceServer struct {
}

func (*UnimplementedJobServiceServer) SubmitJob(JobService_SubmitJobServer) error {
	return status.Errorf(codes.Unimplemented, "method SubmitJob not implemented")
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_SubmitJob_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(JobServiceServer).SubmitJob(&jobServiceSubmitJobServer{stream})
}

type JobService_SubmitJobServer interface {
	SendAndClose(*JobResponse) error
	Recv() (*JobRequest, error)
	grpc.ServerStream
}

type jobServiceSubmitJobServer struct {
	grpc.ServerStream
}

func (x *jobServiceSubmitJobServer) SendAndClose(m *JobResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *jobServiceSubmitJobServer) Recv() (*JobRequest, error) {
	m := new(JobRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pluggable.transfer.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubmitJob",
			Handler:       _JobService_SubmitJob_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "job_service.proto",
}