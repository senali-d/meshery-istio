
//Package meshes ...
package meshes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type OpCategory int32

const (
	OpCategory_INSTALL            OpCategory = 0
	OpCategory_SAMPLE_APPLICATION OpCategory = 1
	OpCategory_CONFIGURE          OpCategory = 2
	OpCategory_VALIDATE           OpCategory = 3
	OpCategory_CUSTOM             OpCategory = 4
)

var OpCategory_name = map[int32]string{
	0: "INSTALL",
	1: "SAMPLE_APPLICATION",
	2: "CONFIGURE",
	3: "VALIDATE",
	4: "CUSTOM",
}
var OpCategory_value = map[string]int32{
	"INSTALL":            0,
	"SAMPLE_APPLICATION": 1,
	"CONFIGURE":          2,
	"VALIDATE":           3,
	"CUSTOM":             4,
}

func (x OpCategory) String() string {
	return proto.EnumName(OpCategory_name, int32(x))
}
func (OpCategory) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_meshops_70742207f69c0a34, []int{0}
}

type EventType int32

const (
	EventType_INFO  EventType = 0
	EventType_WARN  EventType = 1
	EventType_ERROR EventType = 2
)

var EventType_name = map[int32]string{
	0: "INFO",
	1: "WARN",
	2: "ERROR",
}
var EventType_value = map[string]int32{
	"INFO":  0,
	"WARN":  1,
	"ERROR": 2,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}
func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_meshops_70742207f69c0a34, []int{1}
}

type CreateMeshInstanceRequest struct {
	K8SConfig            []byte   `protobuf:"bytes,1,opt,name=k8sConfig,proto3" json:"k8sConfig,omitempty"`
	ContextName          string   `protobuf:"bytes,2,opt,name=contextName,proto3" json:"contextName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMeshInstanceRequest) Reset()         { *m = CreateMeshInstanceRequest{} }
func (m *CreateMeshInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMeshInstanceRequest) ProtoMessage()    {}
func (*CreateMeshInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_meshops_70742207f69c0a34, []int{0}
}
func (m *CreateMeshInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMeshInstanceRequest.Unmarshal(m, b)
}
func (m *CreateMeshInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMeshInstanceRequest.Marshal(b, m, deterministic)
}
func (dst *CreateMeshInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMeshInstanceRequest.Merge(dst, src)
}
func (m *CreateMeshInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMeshInstanceRequest.Size(m)
}
func (m *CreateMeshInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMeshInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMeshInstanceRequest proto.InternalMessageInfo

func (m *CreateMeshInstanceRequest) GetK8SConfig() []byte {
	if m != nil {
		return m.K8SConfig
	}
	return nil
}

func (m *CreateMeshInstanceRequest) GetContextName() string {
	if m != nil {
		return m.ContextName
	}
	return ""
}

type CreateMeshInstanceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMeshInstanceResponse) Reset()         { *m = CreateMeshInstanceResponse{} }
func (m *CreateMeshInstanceResponse) String() string { return proto.CompactTextString(m) }
func (*CreateMeshInstanceResponse) ProtoMessage()    {}
func (*CreateMeshInstanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_meshops_70742207f69c0a34, []int{1}
}
func (m *CreateMeshInstanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMeshInstanceResponse.Unmarshal(m, b)
}
func (m *CreateMeshInstanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMeshInstanceResponse.Marshal(b, m, deterministic)
}
func (dst *CreateMeshInstanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMeshInstanceResponse.Merge(dst, src)
}
func (m *CreateMeshInstanceResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMeshInstanceResponse.Size(m)
}
func (m *CreateMeshInstanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMeshInstanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMeshInstanceResponse proto.InternalMessageInfo

type MeshNameRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshNameRequest) Reset()         { *m = MeshNameRequest{} }
func (m *MeshNameRequest) String() string { return proto.CompactTextString(m) }
func (*MeshNameRequest) ProtoMessage()    {}
func (*MeshNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_meshops_70742207f69c0a34, []int{2}
}
func (m *MeshNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshNameRequest.Unmarshal(m, b)
}
func (m *MeshNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshNameRequest.Marshal(b, m, deterministic)
}
func (dst *MeshNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshNameRequest.Merge(dst, src)
}
func (m *MeshNameRequest) XXX_Size() int {
	return xxx_messageInfo_MeshNameRequest.Size(m)
}
func (m *MeshNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MeshNameRequest proto.InternalMessageInfo

type MeshNameResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshNameResponse) Reset()         { *m = MeshNameResponse{} }
func (m *MeshNameResponse) String() string { return proto.CompactTextString(m) }
func (*MeshNameResponse) ProtoMessage()    {}
func (*MeshNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_meshops_70742207f69c0a34, []int{3}
}
func (m *MeshNameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshNameResponse.Unmarshal(m, b)
}
func (m *MeshNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshNameResponse.Marshal(b, m, deterministic)
}
func (dst *MeshNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshNameResponse.Merge(dst, src)
}
func (m *MeshNameResponse) XXX_Size() int {
	return xxx_messageInfo_MeshNameResponse.Size(m)
}
func (m *MeshNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MeshNameResponse proto.InternalMessageInfo

func (m *MeshNameResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ApplyRuleRequest struct {
	OpName               string   `protobuf:"bytes,1,opt,name=opName,proto3" json:"opName,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	CustomBody           string   `protobuf:"bytes,4,opt,name=custom_body,json=customBody,proto3" json:"custom_body,omitempty"`
	DeleteOp             bool     `protobuf:"varint,5,opt,name=delete_op,json=deleteOp,proto3" json:"delete_op,omitempty"`
	OperationId          string   `protobuf:"bytes,6,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyRuleRequest) Reset()         { *m = ApplyRuleRequest{} }
func (m *ApplyRuleRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyRuleRequest) ProtoMessage()    {}
func (*ApplyRuleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_meshops_70742207f69c0a34, []int{4}
}
func (m *ApplyRuleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyRuleRequest.Unmarshal(m, b)
}
func (m *ApplyRuleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyRuleRequest.Marshal(b, m, deterministic)
}
func (dst *ApplyRuleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyRuleRequest.Merge(dst, src)
}
func (m *ApplyRuleRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyRuleRequest.Size(m)
}
func (m *ApplyRuleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyRuleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyRuleRequest proto.InternalMessageInfo

func (m *ApplyRuleRequest) GetOpName() string {
	if m != nil {
		return m.OpName
	}
	return ""
}

func (m *ApplyRuleRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ApplyRuleRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ApplyRuleRequest) GetCustomBody() string {
	if m != nil {
		return m.CustomBody
	}
	return ""
}

func (m *ApplyRuleRequest) GetDeleteOp() bool {
	if m != nil {
		return m.DeleteOp
	}
	return false
}

func (m *ApplyRuleRequest) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

type ApplyRuleResponse struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	OperationId          string   `protobuf:"bytes,2,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyRuleResponse) Reset()         { *m = ApplyRuleResponse{} }
func (m *ApplyRuleResponse) String() string { return proto.CompactTextString(m) }
func (*ApplyRuleResponse) ProtoMessage()    {}
func (*ApplyRuleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_meshops_70742207f69c0a34, []int{5}
}
func (m *ApplyRuleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyRuleResponse.Unmarshal(m, b)
}
func (m *ApplyRuleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyRuleResponse.Marshal(b, m, deterministic)
}
func (dst *ApplyRuleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyRuleResponse.Merge(dst, src)
}
func (m *ApplyRuleResponse) XXX_Size() int {
	return xxx_messageInfo_ApplyRuleResponse.Size(m)
}
func (m *ApplyRuleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyRuleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyRuleResponse proto.InternalMessageInfo

func (m *ApplyRuleResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ApplyRuleResponse) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

type SupportedOperationsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupportedOperationsRequest) Reset()         { *m = SupportedOperationsRequest{} }
func (m *SupportedOperationsRequest) String() string { return proto.CompactTextString(m) }
func (*SupportedOperationsRequest) ProtoMessage()    {}
func (*SupportedOperationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_meshops_70742207f69c0a34, []int{6}
}
func (m *SupportedOperationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupportedOperationsRequest.Unmarshal(m, b)
}
func (m *SupportedOperationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupportedOperationsRequest.Marshal(b, m, deterministic)
}
func (dst *SupportedOperationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupportedOperationsRequest.Merge(dst, src)
}
func (m *SupportedOperationsRequest) XXX_Size() int {
	return xxx_messageInfo_SupportedOperationsRequest.Size(m)
}
func (m *SupportedOperationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SupportedOperationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SupportedOperationsRequest proto.InternalMessageInfo

type SupportedOperationsResponse struct {
	Ops                  []*SupportedOperation `protobuf:"bytes,1,rep,name=ops,proto3" json:"ops,omitempty"`
	Error                string                `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SupportedOperationsResponse) Reset()         { *m = SupportedOperationsResponse{} }
func (m *SupportedOperationsResponse) String() string { return proto.CompactTextString(m) }
func (*SupportedOperationsResponse) ProtoMessage()    {}
func (*SupportedOperationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_meshops_70742207f69c0a34, []int{7}
}
func (m *SupportedOperationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupportedOperationsResponse.Unmarshal(m, b)
}
func (m *SupportedOperationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupportedOperationsResponse.Marshal(b, m, deterministic)
}
func (dst *SupportedOperationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupportedOperationsResponse.Merge(dst, src)
}
func (m *SupportedOperationsResponse) XXX_Size() int {
	return xxx_messageInfo_SupportedOperationsResponse.Size(m)
}
func (m *SupportedOperationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SupportedOperationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SupportedOperationsResponse proto.InternalMessageInfo

func (m *SupportedOperationsResponse) GetOps() []*SupportedOperation {
	if m != nil {
		return m.Ops
	}
	return nil
}

func (m *SupportedOperationsResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type SupportedOperation struct {
	Key                  string     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string     `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Category             OpCategory `protobuf:"varint,3,opt,name=category,proto3,enum=meshes.OpCategory" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SupportedOperation) Reset()         { *m = SupportedOperation{} }
func (m *SupportedOperation) String() string { return proto.CompactTextString(m) }
func (*SupportedOperation) ProtoMessage()    {}
func (*SupportedOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_meshops_70742207f69c0a34, []int{8}
}
func (m *SupportedOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupportedOperation.Unmarshal(m, b)
}
func (m *SupportedOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupportedOperation.Marshal(b, m, deterministic)
}
func (dst *SupportedOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupportedOperation.Merge(dst, src)
}
func (m *SupportedOperation) XXX_Size() int {
	return xxx_messageInfo_SupportedOperation.Size(m)
}
func (m *SupportedOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_SupportedOperation.DiscardUnknown(m)
}

var xxx_messageInfo_SupportedOperation proto.InternalMessageInfo

func (m *SupportedOperation) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SupportedOperation) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *SupportedOperation) GetCategory() OpCategory {
	if m != nil {
		return m.Category
	}
	return OpCategory_INSTALL
}

type EventsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventsRequest) Reset()         { *m = EventsRequest{} }
func (m *EventsRequest) String() string { return proto.CompactTextString(m) }
func (*EventsRequest) ProtoMessage()    {}
func (*EventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_meshops_70742207f69c0a34, []int{9}
}
func (m *EventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventsRequest.Unmarshal(m, b)
}
func (m *EventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventsRequest.Marshal(b, m, deterministic)
}
func (dst *EventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsRequest.Merge(dst, src)
}
func (m *EventsRequest) XXX_Size() int {
	return xxx_messageInfo_EventsRequest.Size(m)
}
func (m *EventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventsRequest proto.InternalMessageInfo

type EventsResponse struct {
	EventType            EventType `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3,enum=meshes.EventType" json:"event_type,omitempty"`
	Summary              string    `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	Details              string    `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	OperationId          string    `protobuf:"bytes,4,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EventsResponse) Reset()         { *m = EventsResponse{} }
func (m *EventsResponse) String() string { return proto.CompactTextString(m) }
func (*EventsResponse) ProtoMessage()    {}
func (*EventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_meshops_70742207f69c0a34, []int{10}
}
func (m *EventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventsResponse.Unmarshal(m, b)
}
func (m *EventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventsResponse.Marshal(b, m, deterministic)
}
func (dst *EventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsResponse.Merge(dst, src)
}
func (m *EventsResponse) XXX_Size() int {
	return xxx_messageInfo_EventsResponse.Size(m)
}
func (m *EventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventsResponse proto.InternalMessageInfo

func (m *EventsResponse) GetEventType() EventType {
	if m != nil {
		return m.EventType
	}
	return EventType_INFO
}

func (m *EventsResponse) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *EventsResponse) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *EventsResponse) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateMeshInstanceRequest)(nil), "meshes.CreateMeshInstanceRequest")
	proto.RegisterType((*CreateMeshInstanceResponse)(nil), "meshes.CreateMeshInstanceResponse")
	proto.RegisterType((*MeshNameRequest)(nil), "meshes.MeshNameRequest")
	proto.RegisterType((*MeshNameResponse)(nil), "meshes.MeshNameResponse")
	proto.RegisterType((*ApplyRuleRequest)(nil), "meshes.ApplyRuleRequest")
	proto.RegisterType((*ApplyRuleResponse)(nil), "meshes.ApplyRuleResponse")
	proto.RegisterType((*SupportedOperationsRequest)(nil), "meshes.SupportedOperationsRequest")
	proto.RegisterType((*SupportedOperationsResponse)(nil), "meshes.SupportedOperationsResponse")
	proto.RegisterType((*SupportedOperation)(nil), "meshes.SupportedOperation")
	proto.RegisterType((*EventsRequest)(nil), "meshes.EventsRequest")
	proto.RegisterType((*EventsResponse)(nil), "meshes.EventsResponse")
	proto.RegisterEnum("meshes.OpCategory", OpCategory_name, OpCategory_value)
	proto.RegisterEnum("meshes.EventType", EventType_name, EventType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MeshServiceClient is the client API for MeshService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MeshServiceClient interface {
	CreateMeshInstance(ctx context.Context, in *CreateMeshInstanceRequest, opts ...grpc.CallOption) (*CreateMeshInstanceResponse, error)
	MeshName(ctx context.Context, in *MeshNameRequest, opts ...grpc.CallOption) (*MeshNameResponse, error)
	ApplyOperation(ctx context.Context, in *ApplyRuleRequest, opts ...grpc.CallOption) (*ApplyRuleResponse, error)
	SupportedOperations(ctx context.Context, in *SupportedOperationsRequest, opts ...grpc.CallOption) (*SupportedOperationsResponse, error)
	StreamEvents(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (MeshService_StreamEventsClient, error)
}

type meshServiceClient struct {
	cc *grpc.ClientConn
}

func NewMeshServiceClient(cc *grpc.ClientConn) MeshServiceClient {
	return &meshServiceClient{cc}
}

func (c *meshServiceClient) CreateMeshInstance(ctx context.Context, in *CreateMeshInstanceRequest, opts ...grpc.CallOption) (*CreateMeshInstanceResponse, error) {
	out := new(CreateMeshInstanceResponse)
	err := c.cc.Invoke(ctx, "/meshes.MeshService/CreateMeshInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) MeshName(ctx context.Context, in *MeshNameRequest, opts ...grpc.CallOption) (*MeshNameResponse, error) {
	out := new(MeshNameResponse)
	err := c.cc.Invoke(ctx, "/meshes.MeshService/MeshName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) ApplyOperation(ctx context.Context, in *ApplyRuleRequest, opts ...grpc.CallOption) (*ApplyRuleResponse, error) {
	out := new(ApplyRuleResponse)
	err := c.cc.Invoke(ctx, "/meshes.MeshService/ApplyOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) SupportedOperations(ctx context.Context, in *SupportedOperationsRequest, opts ...grpc.CallOption) (*SupportedOperationsResponse, error) {
	out := new(SupportedOperationsResponse)
	err := c.cc.Invoke(ctx, "/meshes.MeshService/SupportedOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) StreamEvents(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (MeshService_StreamEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MeshService_serviceDesc.Streams[0], "/meshes.MeshService/StreamEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &meshServiceStreamEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MeshService_StreamEventsClient interface {
	Recv() (*EventsResponse, error)
	grpc.ClientStream
}

type meshServiceStreamEventsClient struct {
	grpc.ClientStream
}

func (x *meshServiceStreamEventsClient) Recv() (*EventsResponse, error) {
	m := new(EventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MeshServiceServer is the server API for MeshService service.
type MeshServiceServer interface {
	CreateMeshInstance(context.Context, *CreateMeshInstanceRequest) (*CreateMeshInstanceResponse, error)
	MeshName(context.Context, *MeshNameRequest) (*MeshNameResponse, error)
	ApplyOperation(context.Context, *ApplyRuleRequest) (*ApplyRuleResponse, error)
	SupportedOperations(context.Context, *SupportedOperationsRequest) (*SupportedOperationsResponse, error)
	StreamEvents(*EventsRequest, MeshService_StreamEventsServer) error
}

func RegisterMeshServiceServer(s *grpc.Server, srv MeshServiceServer) {
	s.RegisterService(&_MeshService_serviceDesc, srv)
}

func _MeshService_CreateMeshInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMeshInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).CreateMeshInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshes.MeshService/CreateMeshInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).CreateMeshInstance(ctx, req.(*CreateMeshInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_MeshName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeshNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).MeshName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshes.MeshService/MeshName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).MeshName(ctx, req.(*MeshNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_ApplyOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).ApplyOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshes.MeshService/ApplyOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).ApplyOperation(ctx, req.(*ApplyRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_SupportedOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportedOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).SupportedOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshes.MeshService/SupportedOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).SupportedOperations(ctx, req.(*SupportedOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MeshServiceServer).StreamEvents(m, &meshServiceStreamEventsServer{stream})
}

type MeshService_StreamEventsServer interface {
	Send(*EventsResponse) error
	grpc.ServerStream
}

type meshServiceStreamEventsServer struct {
	grpc.ServerStream
}

func (x *meshServiceStreamEventsServer) Send(m *EventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MeshService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "meshes.MeshService",
	HandlerType: (*MeshServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMeshInstance",
			Handler:    _MeshService_CreateMeshInstance_Handler,
		},
		{
			MethodName: "MeshName",
			Handler:    _MeshService_MeshName_Handler,
		},
		{
			MethodName: "ApplyOperation",
			Handler:    _MeshService_ApplyOperation_Handler,
		},
		{
			MethodName: "SupportedOperations",
			Handler:    _MeshService_SupportedOperations_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEvents",
			Handler:       _MeshService_StreamEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "meshops.proto",
}

func init() { proto.RegisterFile("meshops.proto", fileDescriptor_meshops_70742207f69c0a34) }

var fileDescriptor_meshops_70742207f69c0a34 = []byte{
	// 676 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x61, 0x6f, 0xda, 0x3c,
	0x10, 0x6e, 0x80, 0x52, 0x38, 0x28, 0x4d, 0xfd, 0xbe, 0xeb, 0xd2, 0xb4, 0xd2, 0x68, 0x26, 0x4d,
	0xa8, 0x9a, 0x50, 0xc5, 0xbe, 0xec, 0xdb, 0x94, 0x31, 0x5a, 0x45, 0xa2, 0xa4, 0x0a, 0x74, 0x93,
	0x36, 0x4d, 0x2c, 0x85, 0x5b, 0x8b, 0x0a, 0xb1, 0x17, 0x9b, 0x6a, 0xf9, 0x29, 0xd3, 0x7e, 0xcf,
	0xfe, 0xd7, 0xe4, 0x24, 0x0e, 0x5d, 0x43, 0xfb, 0xcd, 0xf7, 0xdc, 0xf9, 0xf1, 0x3d, 0xf6, 0x73,
	0x86, 0xed, 0x05, 0xf2, 0x1b, 0xca, 0x78, 0x9b, 0x85, 0x54, 0x50, 0x52, 0x96, 0x21, 0x72, 0xeb,
	0x0b, 0xec, 0x77, 0x43, 0xf4, 0x05, 0x9e, 0x23, 0xbf, 0x71, 0x02, 0x2e, 0xfc, 0x60, 0x82, 0x1e,
	0xfe, 0x58, 0x22, 0x17, 0xe4, 0x10, 0xaa, 0xb7, 0x6f, 0x79, 0x97, 0x06, 0xdf, 0x67, 0xd7, 0x86,
	0xd6, 0xd4, 0x5a, 0x75, 0x6f, 0x05, 0x90, 0x26, 0xd4, 0x26, 0x34, 0x10, 0xf8, 0x53, 0x0c, 0xfc,
	0x05, 0x1a, 0x85, 0xa6, 0xd6, 0xaa, 0x7a, 0xf7, 0x21, 0xeb, 0x10, 0xcc, 0x75, 0xe4, 0x9c, 0xd1,
	0x80, 0xa3, 0xb5, 0x0b, 0x3b, 0x12, 0x97, 0x95, 0xe9, 0x81, 0xd6, 0x2b, 0xd0, 0x57, 0x50, 0x52,
	0x46, 0x08, 0x94, 0x02, 0xc9, 0xaf, 0xc5, 0xfc, 0xf1, 0xda, 0xfa, 0xa3, 0x81, 0x6e, 0x33, 0x36,
	0x8f, 0xbc, 0xe5, 0x3c, 0xeb, 0x76, 0x0f, 0xca, 0x94, 0x0d, 0x56, 0xa5, 0x69, 0x24, 0x55, 0xc8,
	0x4d, 0x9c, 0xf9, 0x13, 0xd5, 0xe5, 0x0a, 0x20, 0x26, 0x54, 0x96, 0x1c, 0xc3, 0xf8, 0x88, 0x62,
	0x9c, 0xcc, 0x62, 0xf2, 0x02, 0x6a, 0x93, 0x25, 0x17, 0x74, 0x31, 0xbe, 0xa2, 0xd3, 0xc8, 0x28,
	0xc5, 0x69, 0x48, 0xa0, 0xf7, 0x74, 0x1a, 0x91, 0x03, 0xa8, 0x4e, 0x71, 0x8e, 0x02, 0xc7, 0x94,
	0x19, 0x9b, 0x4d, 0xad, 0x55, 0xf1, 0x2a, 0x09, 0xe0, 0x32, 0x72, 0x04, 0x75, 0xca, 0x30, 0xf4,
	0xc5, 0x8c, 0x06, 0xe3, 0xd9, 0xd4, 0x28, 0x27, 0x17, 0x94, 0x61, 0xce, 0xd4, 0xea, 0xc3, 0xee,
	0x3d, 0x19, 0xa9, 0xe0, 0xff, 0x61, 0x13, 0xc3, 0x90, 0x86, 0xa9, 0x8c, 0x24, 0xc8, 0xb1, 0x15,
	0xf2, 0x6c, 0x87, 0x60, 0x0e, 0x97, 0x8c, 0xd1, 0x50, 0xe0, 0xd4, 0x55, 0x38, 0x57, 0x77, 0xeb,
	0xc3, 0xc1, 0xda, 0x6c, 0x7a, 0xea, 0x6b, 0x28, 0x52, 0xc6, 0x0d, 0xad, 0x59, 0x6c, 0xd5, 0x3a,
	0x66, 0x3b, 0xb1, 0x47, 0x3b, 0xbf, 0xc3, 0x93, 0x65, 0xab, 0x1e, 0x0b, 0xf7, 0x7a, 0xb4, 0xe6,
	0x40, 0xf2, 0x1b, 0x88, 0x0e, 0xc5, 0x5b, 0x8c, 0x52, 0x35, 0x72, 0x29, 0x77, 0xdf, 0xf9, 0xf3,
	0xa5, 0x7a, 0x8d, 0x24, 0x20, 0x6d, 0xa8, 0x4c, 0x7c, 0x81, 0xd7, 0x34, 0x8c, 0xe2, 0x97, 0x68,
	0x74, 0x88, 0x6a, 0xc3, 0x65, 0xdd, 0x34, 0xe3, 0x65, 0x35, 0xd6, 0x0e, 0x6c, 0xf7, 0xee, 0x30,
	0x10, 0x99, 0xc2, 0x5f, 0x1a, 0x34, 0x14, 0x92, 0xaa, 0x3a, 0x01, 0x40, 0x89, 0x8c, 0x45, 0xc4,
	0x12, 0x5f, 0x34, 0x3a, 0xbb, 0x8a, 0x35, 0xae, 0x1d, 0x45, 0x0c, 0xbd, 0x2a, 0xaa, 0x25, 0x31,
	0x60, 0x8b, 0x2f, 0x17, 0x0b, 0x3f, 0x8c, 0xd2, 0xee, 0x54, 0x28, 0x33, 0x53, 0x14, 0xfe, 0x6c,
	0xce, 0x53, 0xa3, 0xa8, 0x30, 0xf7, 0x36, 0xa5, 0xdc, 0xdb, 0x1c, 0x7f, 0x06, 0x58, 0x89, 0x20,
	0x35, 0xd8, 0x72, 0x06, 0xc3, 0x91, 0xdd, 0xef, 0xeb, 0x1b, 0x64, 0x0f, 0xc8, 0xd0, 0x3e, 0xbf,
	0xe8, 0xf7, 0xc6, 0xf6, 0xc5, 0x45, 0xdf, 0xe9, 0xda, 0x23, 0xc7, 0x1d, 0xe8, 0x1a, 0xd9, 0x86,
	0x6a, 0xd7, 0x1d, 0x9c, 0x3a, 0x67, 0x97, 0x5e, 0x4f, 0x2f, 0x90, 0x3a, 0x54, 0x3e, 0xda, 0x7d,
	0xe7, 0x83, 0x3d, 0xea, 0xe9, 0x45, 0x02, 0x50, 0xee, 0x5e, 0x0e, 0x47, 0xee, 0xb9, 0x5e, 0x3a,
	0x3e, 0x86, 0x6a, 0x26, 0x85, 0x54, 0xa0, 0xe4, 0x0c, 0x4e, 0x5d, 0x7d, 0x43, 0xae, 0x3e, 0xd9,
	0x9e, 0x64, 0xaa, 0xc2, 0x66, 0xcf, 0xf3, 0x5c, 0x4f, 0x2f, 0x74, 0x7e, 0x17, 0xa1, 0x26, 0x47,
	0x6c, 0x88, 0xe1, 0xdd, 0x6c, 0x82, 0xe4, 0x2b, 0x90, 0xfc, 0x88, 0x92, 0x23, 0x75, 0x45, 0x8f,
	0xfe, 0x0d, 0xa6, 0xf5, 0x54, 0x49, 0x3a, 0xe1, 0x1b, 0xe4, 0x1d, 0x54, 0xd4, 0x40, 0x93, 0xe7,
	0x6a, 0xc7, 0x83, 0xa9, 0x37, 0x8d, 0x7c, 0x22, 0x23, 0x38, 0x83, 0x46, 0x3c, 0x21, 0x2b, 0x3b,
	0x65, 0xd5, 0x0f, 0x3f, 0x00, 0x73, 0x7f, 0x4d, 0x26, 0x23, 0xfa, 0x06, 0xff, 0xad, 0xb1, 0x3f,
	0xb1, 0x1e, 0x77, 0xba, 0xf2, 0x95, 0xf9, 0xf2, 0xc9, 0x9a, 0xec, 0x04, 0x1b, 0xea, 0x43, 0x11,
	0xa2, 0xbf, 0x48, 0x3c, 0x48, 0x9e, 0xfd, 0xe3, 0xb3, 0x8c, 0x6d, 0xef, 0x21, 0xac, 0x08, 0x4e,
	0xb4, 0xab, 0x72, 0xfc, 0x39, 0xbf, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x27, 0xa3, 0x6e, 0x5a,
	0xad, 0x05, 0x00, 0x00,
}

