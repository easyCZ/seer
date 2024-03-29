// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/v1/synthetics.proto

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

type Synthetic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Spec *SyntheticSpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *Synthetic) Reset() {
	*x = Synthetic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_synthetics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Synthetic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Synthetic) ProtoMessage() {}

func (x *Synthetic) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_synthetics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Synthetic.ProtoReflect.Descriptor instead.
func (*Synthetic) Descriptor() ([]byte, []int) {
	return file_proto_v1_synthetics_proto_rawDescGZIP(), []int{0}
}

func (x *Synthetic) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Synthetic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Synthetic) GetSpec() *SyntheticSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type SyntheticSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Variables []*Variable `protobuf:"bytes,1,rep,name=variables,proto3" json:"variables,omitempty"`
	Steps     []*Step     `protobuf:"bytes,2,rep,name=steps,proto3" json:"steps,omitempty"`
}

func (x *SyntheticSpec) Reset() {
	*x = SyntheticSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_synthetics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyntheticSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyntheticSpec) ProtoMessage() {}

func (x *SyntheticSpec) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_synthetics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyntheticSpec.ProtoReflect.Descriptor instead.
func (*SyntheticSpec) Descriptor() ([]byte, []int) {
	return file_proto_v1_synthetics_proto_rawDescGZIP(), []int{1}
}

func (x *SyntheticSpec) GetVariables() []*Variable {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *SyntheticSpec) GetSteps() []*Step {
	if x != nil {
		return x.Steps
	}
	return nil
}

type Variable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Variable) Reset() {
	*x = Variable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_synthetics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Variable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Variable) ProtoMessage() {}

func (x *Variable) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_synthetics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Variable.ProtoReflect.Descriptor instead.
func (*Variable) Descriptor() ([]byte, []int) {
	return file_proto_v1_synthetics_proto_rawDescGZIP(), []int{2}
}

func (x *Variable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Variable) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Spec *StepSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_synthetics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_synthetics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_proto_v1_synthetics_proto_rawDescGZIP(), []int{3}
}

func (x *Step) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Step) GetSpec() *StepSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type StepSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string     `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Method   string     `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Body     string     `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Extracts []*Extract `protobuf:"bytes,4,rep,name=extracts,proto3" json:"extracts,omitempty"`
}

func (x *StepSpec) Reset() {
	*x = StepSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_synthetics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StepSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepSpec) ProtoMessage() {}

func (x *StepSpec) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_synthetics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepSpec.ProtoReflect.Descriptor instead.
func (*StepSpec) Descriptor() ([]byte, []int) {
	return file_proto_v1_synthetics_proto_rawDescGZIP(), []int{4}
}

func (x *StepSpec) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *StepSpec) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *StepSpec) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *StepSpec) GetExtracts() []*Extract {
	if x != nil {
		return x.Extracts
	}
	return nil
}

type Extract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of this extract
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to From:
	//	*Extract_Header
	//	*Extract_Body
	From isExtract_From `protobuf_oneof:"from"`
}

func (x *Extract) Reset() {
	*x = Extract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_synthetics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extract) ProtoMessage() {}

func (x *Extract) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_synthetics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extract.ProtoReflect.Descriptor instead.
func (*Extract) Descriptor() ([]byte, []int) {
	return file_proto_v1_synthetics_proto_rawDescGZIP(), []int{5}
}

func (x *Extract) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *Extract) GetFrom() isExtract_From {
	if m != nil {
		return m.From
	}
	return nil
}

func (x *Extract) GetHeader() *HeaderExtract {
	if x, ok := x.GetFrom().(*Extract_Header); ok {
		return x.Header
	}
	return nil
}

func (x *Extract) GetBody() *BodyExtract {
	if x, ok := x.GetFrom().(*Extract_Body); ok {
		return x.Body
	}
	return nil
}

type isExtract_From interface {
	isExtract_From()
}

type Extract_Header struct {
	Header *HeaderExtract `protobuf:"bytes,2,opt,name=header,proto3,oneof"`
}

type Extract_Body struct {
	Body *BodyExtract `protobuf:"bytes,3,opt,name=body,proto3,oneof"`
}

func (*Extract_Header) isExtract_From() {}

func (*Extract_Body) isExtract_From() {}

type HeaderExtract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// header_name is the name of the header from which to extract data.
	HeaderName string `protobuf:"bytes,1,opt,name=header_name,json=headerName,proto3" json:"header_name,omitempty"`
	// query is the query to run against the header value to extract data.
	// when not specified, the whole header will be extracted.
	Query *ExtractQuery `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *HeaderExtract) Reset() {
	*x = HeaderExtract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_synthetics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderExtract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderExtract) ProtoMessage() {}

func (x *HeaderExtract) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_synthetics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderExtract.ProtoReflect.Descriptor instead.
func (*HeaderExtract) Descriptor() ([]byte, []int) {
	return file_proto_v1_synthetics_proto_rawDescGZIP(), []int{6}
}

func (x *HeaderExtract) GetHeaderName() string {
	if x != nil {
		return x.HeaderName
	}
	return ""
}

func (x *HeaderExtract) GetQuery() *ExtractQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

type BodyExtract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// query is the query to run against the body to extract data.
	Query *ExtractQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *BodyExtract) Reset() {
	*x = BodyExtract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_synthetics_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyExtract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyExtract) ProtoMessage() {}

func (x *BodyExtract) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_synthetics_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyExtract.ProtoReflect.Descriptor instead.
func (*BodyExtract) Descriptor() ([]byte, []int) {
	return file_proto_v1_synthetics_proto_rawDescGZIP(), []int{7}
}

func (x *BodyExtract) GetQuery() *ExtractQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

type ExtractQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Expression:
	//	*ExtractQuery_Jq
	//	*ExtractQuery_Regexp
	//	*ExtractQuery_Yq
	Expression isExtractQuery_Expression `protobuf_oneof:"expression"`
}

func (x *ExtractQuery) Reset() {
	*x = ExtractQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_synthetics_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractQuery) ProtoMessage() {}

func (x *ExtractQuery) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_synthetics_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractQuery.ProtoReflect.Descriptor instead.
func (*ExtractQuery) Descriptor() ([]byte, []int) {
	return file_proto_v1_synthetics_proto_rawDescGZIP(), []int{8}
}

func (m *ExtractQuery) GetExpression() isExtractQuery_Expression {
	if m != nil {
		return m.Expression
	}
	return nil
}

func (x *ExtractQuery) GetJq() string {
	if x, ok := x.GetExpression().(*ExtractQuery_Jq); ok {
		return x.Jq
	}
	return ""
}

func (x *ExtractQuery) GetRegexp() string {
	if x, ok := x.GetExpression().(*ExtractQuery_Regexp); ok {
		return x.Regexp
	}
	return ""
}

func (x *ExtractQuery) GetYq() string {
	if x, ok := x.GetExpression().(*ExtractQuery_Yq); ok {
		return x.Yq
	}
	return ""
}

type isExtractQuery_Expression interface {
	isExtractQuery_Expression()
}

type ExtractQuery_Jq struct {
	// The expression is a JSON Query expression.
	Jq string `protobuf:"bytes,1,opt,name=jq,proto3,oneof"`
}

type ExtractQuery_Regexp struct {
	// The expression is a Regular expression.
	Regexp string `protobuf:"bytes,2,opt,name=regexp,proto3,oneof"`
}

type ExtractQuery_Yq struct {
	// The expression is a Yaml Query expression.
	Yq string `protobuf:"bytes,3,opt,name=yq,proto3,oneof"`
}

func (*ExtractQuery_Jq) isExtractQuery_Expression() {}

func (*ExtractQuery_Regexp) isExtractQuery_Expression() {}

func (*ExtractQuery_Yq) isExtractQuery_Expression() {}

type ListSyntheticsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSyntheticsRequest) Reset() {
	*x = ListSyntheticsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_synthetics_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSyntheticsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSyntheticsRequest) ProtoMessage() {}

func (x *ListSyntheticsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_synthetics_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSyntheticsRequest.ProtoReflect.Descriptor instead.
func (*ListSyntheticsRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_synthetics_proto_rawDescGZIP(), []int{9}
}

var File_proto_v1_synthetics_proto protoreflect.FileDescriptor

var file_proto_v1_synthetics_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x6e, 0x74, 0x68,
	0x65, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x79, 0x6e,
	0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x61, 0x0a, 0x09, 0x53, 0x79,
	0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x79, 0x6e, 0x74,
	0x68, 0x65, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65,
	0x74, 0x69, 0x63, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x71, 0x0a,
	0x0d, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x53, 0x70, 0x65, 0x63, 0x12, 0x35,
	0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73,
	0x22, 0x34, 0x0a, 0x08, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x47, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x65, 0x70, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22,
	0x7c, 0x0a, 0x08, 0x53, 0x74, 0x65, 0x70, 0x53, 0x70, 0x65, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x32, 0x0a, 0x08, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x79,
	0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x52, 0x08, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x22, 0x8f, 0x01,
	0x0a, 0x07, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x48, 0x00, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x48,
	0x00, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x06, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x22,
	0x63, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x22, 0x40, 0x0a, 0x0b, 0x42, 0x6f, 0x64, 0x79, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x5a, 0x0a, 0x0c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x02, 0x6a, 0x71, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x6a, 0x71, 0x12, 0x18, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x65,
	0x78, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x67, 0x65,
	0x78, 0x70, 0x12, 0x10, 0x0a, 0x02, 0x79, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x02, 0x79, 0x71, 0x42, 0x0c, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65,
	0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x67, 0x0a, 0x11, 0x53,
	0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x52, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69,
	0x63, 0x73, 0x12, 0x24, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x68,
	0x65, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74,
	0x69, 0x63, 0x30, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x61, 0x73, 0x79, 0x43, 0x5a, 0x2f, 0x73, 0x65, 0x65, 0x72, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_synthetics_proto_rawDescOnce sync.Once
	file_proto_v1_synthetics_proto_rawDescData = file_proto_v1_synthetics_proto_rawDesc
)

func file_proto_v1_synthetics_proto_rawDescGZIP() []byte {
	file_proto_v1_synthetics_proto_rawDescOnce.Do(func() {
		file_proto_v1_synthetics_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_synthetics_proto_rawDescData)
	})
	return file_proto_v1_synthetics_proto_rawDescData
}

var file_proto_v1_synthetics_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_v1_synthetics_proto_goTypes = []interface{}{
	(*Synthetic)(nil),             // 0: synthetics.v1.Synthetic
	(*SyntheticSpec)(nil),         // 1: synthetics.v1.SyntheticSpec
	(*Variable)(nil),              // 2: synthetics.v1.Variable
	(*Step)(nil),                  // 3: synthetics.v1.Step
	(*StepSpec)(nil),              // 4: synthetics.v1.StepSpec
	(*Extract)(nil),               // 5: synthetics.v1.Extract
	(*HeaderExtract)(nil),         // 6: synthetics.v1.HeaderExtract
	(*BodyExtract)(nil),           // 7: synthetics.v1.BodyExtract
	(*ExtractQuery)(nil),          // 8: synthetics.v1.ExtractQuery
	(*ListSyntheticsRequest)(nil), // 9: synthetics.v1.ListSyntheticsRequest
}
var file_proto_v1_synthetics_proto_depIdxs = []int32{
	1,  // 0: synthetics.v1.Synthetic.spec:type_name -> synthetics.v1.SyntheticSpec
	2,  // 1: synthetics.v1.SyntheticSpec.variables:type_name -> synthetics.v1.Variable
	3,  // 2: synthetics.v1.SyntheticSpec.steps:type_name -> synthetics.v1.Step
	4,  // 3: synthetics.v1.Step.spec:type_name -> synthetics.v1.StepSpec
	5,  // 4: synthetics.v1.StepSpec.extracts:type_name -> synthetics.v1.Extract
	6,  // 5: synthetics.v1.Extract.header:type_name -> synthetics.v1.HeaderExtract
	7,  // 6: synthetics.v1.Extract.body:type_name -> synthetics.v1.BodyExtract
	8,  // 7: synthetics.v1.HeaderExtract.query:type_name -> synthetics.v1.ExtractQuery
	8,  // 8: synthetics.v1.BodyExtract.query:type_name -> synthetics.v1.ExtractQuery
	9,  // 9: synthetics.v1.SyntheticsService.ListSynthetics:input_type -> synthetics.v1.ListSyntheticsRequest
	0,  // 10: synthetics.v1.SyntheticsService.ListSynthetics:output_type -> synthetics.v1.Synthetic
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_proto_v1_synthetics_proto_init() }
func file_proto_v1_synthetics_proto_init() {
	if File_proto_v1_synthetics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_synthetics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Synthetic); i {
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
		file_proto_v1_synthetics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyntheticSpec); i {
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
		file_proto_v1_synthetics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Variable); i {
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
		file_proto_v1_synthetics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Step); i {
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
		file_proto_v1_synthetics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StepSpec); i {
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
		file_proto_v1_synthetics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extract); i {
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
		file_proto_v1_synthetics_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderExtract); i {
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
		file_proto_v1_synthetics_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyExtract); i {
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
		file_proto_v1_synthetics_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractQuery); i {
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
		file_proto_v1_synthetics_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSyntheticsRequest); i {
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
	file_proto_v1_synthetics_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Extract_Header)(nil),
		(*Extract_Body)(nil),
	}
	file_proto_v1_synthetics_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*ExtractQuery_Jq)(nil),
		(*ExtractQuery_Regexp)(nil),
		(*ExtractQuery_Yq)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_v1_synthetics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_synthetics_proto_goTypes,
		DependencyIndexes: file_proto_v1_synthetics_proto_depIdxs,
		MessageInfos:      file_proto_v1_synthetics_proto_msgTypes,
	}.Build()
	File_proto_v1_synthetics_proto = out.File
	file_proto_v1_synthetics_proto_rawDesc = nil
	file_proto_v1_synthetics_proto_goTypes = nil
	file_proto_v1_synthetics_proto_depIdxs = nil
}
