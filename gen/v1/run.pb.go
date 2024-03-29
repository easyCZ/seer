// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/v1/run.proto

package apiv1

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

type Run struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SyntheticId string        `protobuf:"bytes,2,opt,name=synthetic_id,json=syntheticId,proto3" json:"synthetic_id,omitempty"`
	Results     []*StepResult `protobuf:"bytes,3,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *Run) Reset() {
	*x = Run{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_run_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Run) ProtoMessage() {}

func (x *Run) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_run_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Run.ProtoReflect.Descriptor instead.
func (*Run) Descriptor() ([]byte, []int) {
	return file_proto_v1_run_proto_rawDescGZIP(), []int{0}
}

func (x *Run) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Run) GetSyntheticId() string {
	if x != nil {
		return x.SyntheticId
	}
	return ""
}

func (x *Run) GetResults() []*StepResult {
	if x != nil {
		return x.Results
	}
	return nil
}

type StepResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StepName string `protobuf:"bytes,1,opt,name=step_name,json=stepName,proto3" json:"step_name,omitempty"`
	// Types that are assignable to Outcome:
	//	*StepResult_Response
	//	*StepResult_Error
	Outcome   isStepResult_Outcome `protobuf_oneof:"outcome"`
	Variables []*Variable          `protobuf:"bytes,4,rep,name=variables,proto3" json:"variables,omitempty"`
}

func (x *StepResult) Reset() {
	*x = StepResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_run_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StepResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepResult) ProtoMessage() {}

func (x *StepResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_run_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepResult.ProtoReflect.Descriptor instead.
func (*StepResult) Descriptor() ([]byte, []int) {
	return file_proto_v1_run_proto_rawDescGZIP(), []int{1}
}

func (x *StepResult) GetStepName() string {
	if x != nil {
		return x.StepName
	}
	return ""
}

func (m *StepResult) GetOutcome() isStepResult_Outcome {
	if m != nil {
		return m.Outcome
	}
	return nil
}

func (x *StepResult) GetResponse() *Response {
	if x, ok := x.GetOutcome().(*StepResult_Response); ok {
		return x.Response
	}
	return nil
}

func (x *StepResult) GetError() *StepError {
	if x, ok := x.GetOutcome().(*StepResult_Error); ok {
		return x.Error
	}
	return nil
}

func (x *StepResult) GetVariables() []*Variable {
	if x != nil {
		return x.Variables
	}
	return nil
}

type isStepResult_Outcome interface {
	isStepResult_Outcome()
}

type StepResult_Response struct {
	// response is present if the step completed succesfully
	Response *Response `protobuf:"bytes,2,opt,name=response,proto3,oneof"`
}

type StepResult_Error struct {
	// error is if the step failed to execute
	Error *StepError `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*StepResult_Response) isStepResult_Outcome() {}

func (*StepResult_Error) isStepResult_Outcome() {}

type StepError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Details string `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *StepError) Reset() {
	*x = StepError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_run_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StepError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepError) ProtoMessage() {}

func (x *StepError) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_run_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepError.ProtoReflect.Descriptor instead.
func (*StepError) Descriptor() ([]byte, []int) {
	return file_proto_v1_run_proto_rawDescGZIP(), []int{2}
}

func (x *StepError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StepError) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Headers []*Header `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	Body    string    `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_run_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_run_proto_msgTypes[3]
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
	return file_proto_v1_run_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Response) GetHeaders() []*Header {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Response) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_run_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_run_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_proto_v1_run_proto_rawDescGZIP(), []int{4}
}

func (x *Header) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Header) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Run *Run `protobuf:"bytes,1,opt,name=run,proto3" json:"run,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_run_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_run_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_run_proto_rawDescGZIP(), []int{5}
}

func (x *CreateRequest) GetRun() *Run {
	if x != nil {
		return x.Run
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Run *Run `protobuf:"bytes,1,opt,name=run,proto3" json:"run,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_run_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_run_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_run_proto_rawDescGZIP(), []int{6}
}

func (x *CreateResponse) GetRun() *Run {
	if x != nil {
		return x.Run
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_run_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_run_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_run_proto_rawDescGZIP(), []int{7}
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runs []*Run `protobuf:"bytes,1,rep,name=runs,proto3" json:"runs,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_run_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_run_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_run_proto_rawDescGZIP(), []int{8}
}

func (x *ListResponse) GetRuns() []*Run {
	if x != nil {
		return x.Runs
	}
	return nil
}

var File_proto_v1_run_proto protoreflect.FileDescriptor

var file_proto_v1_run_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x49,
	0x64, 0x12, 0x2c, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x65, 0x70,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22,
	0xc6, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x74, 0x65, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x75, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x35, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x79, 0x6e, 0x74,
	0x68, 0x65, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x42, 0x09, 0x0a,
	0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x09, 0x53, 0x74, 0x65, 0x70,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x60, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x30, 0x0a, 0x06, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2e, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x03, 0x72, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x75,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x03, 0x72, 0x75, 0x6e, 0x22, 0x2f, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x03, 0x72, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72,
	0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x03, 0x72, 0x75, 0x6e, 0x22, 0x0d,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2f, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x72, 0x75, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x75,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x32, 0x7c,
	0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x13, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x61, 0x73, 0x79, 0x43,
	0x5a, 0x2f, 0x73, 0x65, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_run_proto_rawDescOnce sync.Once
	file_proto_v1_run_proto_rawDescData = file_proto_v1_run_proto_rawDesc
)

func file_proto_v1_run_proto_rawDescGZIP() []byte {
	file_proto_v1_run_proto_rawDescOnce.Do(func() {
		file_proto_v1_run_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_run_proto_rawDescData)
	})
	return file_proto_v1_run_proto_rawDescData
}

var file_proto_v1_run_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_v1_run_proto_goTypes = []interface{}{
	(*Run)(nil),            // 0: run.v1.Run
	(*StepResult)(nil),     // 1: run.v1.StepResult
	(*StepError)(nil),      // 2: run.v1.StepError
	(*Response)(nil),       // 3: run.v1.Response
	(*Header)(nil),         // 4: run.v1.Header
	(*CreateRequest)(nil),  // 5: run.v1.CreateRequest
	(*CreateResponse)(nil), // 6: run.v1.CreateResponse
	(*ListRequest)(nil),    // 7: run.v1.ListRequest
	(*ListResponse)(nil),   // 8: run.v1.ListResponse
	(*Variable)(nil),       // 9: synthetics.v1.Variable
}
var file_proto_v1_run_proto_depIdxs = []int32{
	1,  // 0: run.v1.Run.results:type_name -> run.v1.StepResult
	3,  // 1: run.v1.StepResult.response:type_name -> run.v1.Response
	2,  // 2: run.v1.StepResult.error:type_name -> run.v1.StepError
	9,  // 3: run.v1.StepResult.variables:type_name -> synthetics.v1.Variable
	4,  // 4: run.v1.Response.headers:type_name -> run.v1.Header
	0,  // 5: run.v1.CreateRequest.run:type_name -> run.v1.Run
	0,  // 6: run.v1.CreateResponse.run:type_name -> run.v1.Run
	0,  // 7: run.v1.ListResponse.runs:type_name -> run.v1.Run
	5,  // 8: run.v1.RunService.Create:input_type -> run.v1.CreateRequest
	7,  // 9: run.v1.RunService.List:input_type -> run.v1.ListRequest
	6,  // 10: run.v1.RunService.Create:output_type -> run.v1.CreateResponse
	8,  // 11: run.v1.RunService.List:output_type -> run.v1.ListResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_v1_run_proto_init() }
func file_proto_v1_run_proto_init() {
	if File_proto_v1_run_proto != nil {
		return
	}
	file_proto_v1_synthetics_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_run_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Run); i {
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
		file_proto_v1_run_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StepResult); i {
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
		file_proto_v1_run_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StepError); i {
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
		file_proto_v1_run_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_v1_run_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_proto_v1_run_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_proto_v1_run_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_proto_v1_run_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_proto_v1_run_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
	file_proto_v1_run_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*StepResult_Response)(nil),
		(*StepResult_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_v1_run_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_run_proto_goTypes,
		DependencyIndexes: file_proto_v1_run_proto_depIdxs,
		MessageInfos:      file_proto_v1_run_proto_msgTypes,
	}.Build()
	File_proto_v1_run_proto = out.File
	file_proto_v1_run_proto_rawDesc = nil
	file_proto_v1_run_proto_goTypes = nil
	file_proto_v1_run_proto_depIdxs = nil
}
