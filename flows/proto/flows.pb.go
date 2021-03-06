// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flows.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	proto1 "www.velocidex.com/golang/velociraptor/crypto/proto"
	_ "www.velocidex.com/golang/velociraptor/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FlowContext_State int32

const (
	FlowContext_UNSET      FlowContext_State = 0
	FlowContext_RUNNING    FlowContext_State = 1
	FlowContext_TERMINATED FlowContext_State = 2
	FlowContext_ERROR      FlowContext_State = 3
	FlowContext_ARCHIVED   FlowContext_State = 4
)

var FlowContext_State_name = map[int32]string{
	0: "UNSET",
	1: "RUNNING",
	2: "TERMINATED",
	3: "ERROR",
	4: "ARCHIVED",
}

var FlowContext_State_value = map[string]int32{
	"UNSET":      0,
	"RUNNING":    1,
	"TERMINATED": 2,
	"ERROR":      3,
	"ARCHIVED":   4,
}

func (x FlowContext_State) String() string {
	return proto.EnumName(FlowContext_State_name, int32(x))
}

func (FlowContext_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b6bd6ebccf1a7f18, []int{1, 0}
}

type UploadedFileInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VfsPath              string   `protobuf:"bytes,2,opt,name=vfs_path,json=vfsPath,proto3" json:"vfs_path,omitempty"`
	Size                 uint64   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadedFileInfo) Reset()         { *m = UploadedFileInfo{} }
func (m *UploadedFileInfo) String() string { return proto.CompactTextString(m) }
func (*UploadedFileInfo) ProtoMessage()    {}
func (*UploadedFileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6bd6ebccf1a7f18, []int{0}
}

func (m *UploadedFileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadedFileInfo.Unmarshal(m, b)
}
func (m *UploadedFileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadedFileInfo.Marshal(b, m, deterministic)
}
func (m *UploadedFileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadedFileInfo.Merge(m, src)
}
func (m *UploadedFileInfo) XXX_Size() int {
	return xxx_messageInfo_UploadedFileInfo.Size(m)
}
func (m *UploadedFileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadedFileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UploadedFileInfo proto.InternalMessageInfo

func (m *UploadedFileInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UploadedFileInfo) GetVfsPath() string {
	if m != nil {
		return m.VfsPath
	}
	return ""
}

func (m *UploadedFileInfo) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

// The flow context.
// Next field: 19
type FlowContext struct {
	Backtrace        string `protobuf:"bytes,1,opt,name=backtrace,proto3" json:"backtrace,omitempty"`
	CreateTime       uint64 `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	KillTimestamp    uint64 `protobuf:"varint,6,opt,name=kill_timestamp,json=killTimestamp,proto3" json:"kill_timestamp,omitempty"`
	NetworkBytesSent uint64 `protobuf:"varint,7,opt,name=network_bytes_sent,json=networkBytesSent,proto3" json:"network_bytes_sent,omitempty"`
	// Uploads are now permanently stored in a csv file. This field is
	// never serialized - it is just a place holder during processing.
	UploadedFiles              []*UploadedFileInfo `protobuf:"bytes,24,rep,name=uploaded_files,json=uploadedFiles,proto3" json:"uploaded_files,omitempty"`
	TotalUploadedFiles         uint64              `protobuf:"varint,23,opt,name=total_uploaded_files,json=totalUploadedFiles,proto3" json:"total_uploaded_files,omitempty"`
	TotalExpectedUploadedBytes uint64              `protobuf:"varint,25,opt,name=total_expected_uploaded_bytes,json=totalExpectedUploadedBytes,proto3" json:"total_expected_uploaded_bytes,omitempty"`
	TotalUploadedBytes         uint64              `protobuf:"varint,26,opt,name=total_uploaded_bytes,json=totalUploadedBytes,proto3" json:"total_uploaded_bytes,omitempty"`
	// Logs are stored in their own CSV file. This is just a
	// placeholder during processing.
	Logs                 []*proto1.LogMessage `protobuf:"bytes,20,rep,name=logs,proto3" json:"logs,omitempty"`
	SessionId            string               `protobuf:"bytes,13,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	State                FlowContext_State    `protobuf:"varint,14,opt,name=state,proto3,enum=proto.FlowContext_State" json:"state,omitempty"`
	Status               string               `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	UserNotified         bool                 `protobuf:"varint,16,opt,name=user_notified,json=userNotified,proto3" json:"user_notified,omitempty"`
	ActiveTime           uint64               `protobuf:"varint,17,opt,name=active_time,json=activeTime,proto3" json:"active_time,omitempty"`
	Artifacts            []string             `protobuf:"bytes,21,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	ArtifactsWithResults []string             `protobuf:"bytes,22,rep,name=artifacts_with_results,json=artifactsWithResults,proto3" json:"artifacts_with_results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FlowContext) Reset()         { *m = FlowContext{} }
func (m *FlowContext) String() string { return proto.CompactTextString(m) }
func (*FlowContext) ProtoMessage()    {}
func (*FlowContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6bd6ebccf1a7f18, []int{1}
}

func (m *FlowContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowContext.Unmarshal(m, b)
}
func (m *FlowContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowContext.Marshal(b, m, deterministic)
}
func (m *FlowContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowContext.Merge(m, src)
}
func (m *FlowContext) XXX_Size() int {
	return xxx_messageInfo_FlowContext.Size(m)
}
func (m *FlowContext) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowContext.DiscardUnknown(m)
}

var xxx_messageInfo_FlowContext proto.InternalMessageInfo

func (m *FlowContext) GetBacktrace() string {
	if m != nil {
		return m.Backtrace
	}
	return ""
}

func (m *FlowContext) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *FlowContext) GetKillTimestamp() uint64 {
	if m != nil {
		return m.KillTimestamp
	}
	return 0
}

func (m *FlowContext) GetNetworkBytesSent() uint64 {
	if m != nil {
		return m.NetworkBytesSent
	}
	return 0
}

func (m *FlowContext) GetUploadedFiles() []*UploadedFileInfo {
	if m != nil {
		return m.UploadedFiles
	}
	return nil
}

func (m *FlowContext) GetTotalUploadedFiles() uint64 {
	if m != nil {
		return m.TotalUploadedFiles
	}
	return 0
}

func (m *FlowContext) GetTotalExpectedUploadedBytes() uint64 {
	if m != nil {
		return m.TotalExpectedUploadedBytes
	}
	return 0
}

func (m *FlowContext) GetTotalUploadedBytes() uint64 {
	if m != nil {
		return m.TotalUploadedBytes
	}
	return 0
}

func (m *FlowContext) GetLogs() []*proto1.LogMessage {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *FlowContext) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *FlowContext) GetState() FlowContext_State {
	if m != nil {
		return m.State
	}
	return FlowContext_UNSET
}

func (m *FlowContext) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *FlowContext) GetUserNotified() bool {
	if m != nil {
		return m.UserNotified
	}
	return false
}

func (m *FlowContext) GetActiveTime() uint64 {
	if m != nil {
		return m.ActiveTime
	}
	return 0
}

func (m *FlowContext) GetArtifacts() []string {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

func (m *FlowContext) GetArtifactsWithResults() []string {
	if m != nil {
		return m.ArtifactsWithResults
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.FlowContext_State", FlowContext_State_name, FlowContext_State_value)
	proto.RegisterType((*UploadedFileInfo)(nil), "proto.UploadedFileInfo")
	proto.RegisterType((*FlowContext)(nil), "proto.FlowContext")
}

func init() { proto.RegisterFile("flows.proto", fileDescriptor_b6bd6ebccf1a7f18) }

var fileDescriptor_b6bd6ebccf1a7f18 = []byte{
	// 867 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xff, 0x8e, 0x1b, 0x35,
	0x10, 0xc7, 0x49, 0xef, 0x67, 0x9c, 0xe6, 0x48, 0xcd, 0xd1, 0x6e, 0x23, 0x2a, 0xcc, 0x51, 0x41,
	0x84, 0xca, 0x46, 0x14, 0x10, 0x52, 0x25, 0x54, 0x5d, 0xee, 0x72, 0x5c, 0xa4, 0x36, 0x45, 0xbe,
	0x04, 0x84, 0x84, 0xb4, 0x72, 0x76, 0x67, 0x13, 0x73, 0xde, 0x75, 0x64, 0xcf, 0x66, 0xef, 0x78,
	0x19, 0x5e, 0x86, 0x27, 0x81, 0xd7, 0x00, 0x09, 0xd9, 0xbb, 0x97, 0x94, 0xab, 0x40, 0xf0, 0x57,
	0x9c, 0x99, 0xef, 0x7c, 0xbe, 0xf6, 0x78, 0xbc, 0xa4, 0x95, 0x2a, 0x5d, 0xda, 0x70, 0x69, 0x34,
	0x6a, 0xba, 0xe3, 0x7f, 0xba, 0xcf, 0xca, 0xb2, 0x0c, 0x57, 0xa0, 0x74, 0x2c, 0x13, 0xb8, 0x0a,
	0x63, 0x9d, 0xf5, 0xe7, 0x5a, 0x89, 0x7c, 0xde, 0xaf, 0x82, 0x46, 0x2c, 0x51, 0x9b, 0xbe, 0x17,
	0xf7, 0x2d, 0x64, 0x22, 0x47, 0x19, 0x57, 0x88, 0xee, 0xf3, 0xff, 0x53, 0xeb, 0xbc, 0xa3, 0x0c,
	0x50, 0x24, 0x02, 0x45, 0x0d, 0xf8, 0xfa, 0xbf, 0x01, 0x62, 0x73, 0xbd, 0x44, 0x5d, 0x73, 0x7e,
	0xd2, 0xb3, 0xfa, 0x08, 0x47, 0x53, 0xd2, 0x99, 0x2e, 0x95, 0x16, 0x09, 0x24, 0x67, 0x52, 0xc1,
	0x28, 0x4f, 0x35, 0xa5, 0x64, 0x3b, 0x17, 0x19, 0x04, 0x0d, 0xd6, 0xe8, 0x35, 0xb9, 0x5f, 0xd3,
	0x87, 0x64, 0x7f, 0x95, 0xda, 0x68, 0x29, 0x70, 0x11, 0xdc, 0xf1, 0xf1, 0xbd, 0x55, 0x6a, 0xbf,
	0x15, 0xb8, 0x70, 0x72, 0x2b, 0x7f, 0x86, 0x60, 0x8b, 0x35, 0x7a, 0xdb, 0xdc, 0xaf, 0x8f, 0xfe,
	0x6c, 0x92, 0xd6, 0x99, 0xd2, 0xe5, 0x89, 0xce, 0x11, 0xae, 0x90, 0xbe, 0x47, 0x9a, 0x33, 0x11,
	0x5f, 0xa2, 0x11, 0xf1, 0x0d, 0x77, 0x13, 0xa0, 0x5f, 0x90, 0x56, 0x6c, 0x40, 0x20, 0x44, 0x28,
	0xb3, 0x1a, 0x34, 0x78, 0xe7, 0xb7, 0x3f, 0x7e, 0xff, 0xb5, 0xd1, 0x26, 0x2d, 0x7e, 0x7a, 0x76,
	0x2a, 0x10, 0x5c, 0x8a, 0x93, 0x4a, 0x37, 0x91, 0x19, 0xd0, 0x67, 0xe4, 0xe0, 0x52, 0x2a, 0xe5,
	0x6b, 0x2c, 0x8a, 0x6c, 0x19, 0xec, 0xfe, 0x73, 0x61, 0xdb, 0x49, 0x27, 0x37, 0x4a, 0xfa, 0x84,
	0xd0, 0x1c, 0xb0, 0xd4, 0xe6, 0x32, 0x9a, 0x5d, 0x23, 0xd8, 0xc8, 0x42, 0x8e, 0xc1, 0x9e, 0x3f,
	0x41, 0xa7, 0xce, 0x0c, 0x5c, 0xe2, 0x02, 0x72, 0xa4, 0x05, 0x39, 0x28, 0xea, 0x26, 0x45, 0xa9,
	0x54, 0x60, 0x83, 0x80, 0x6d, 0xf5, 0x5a, 0x4f, 0x1f, 0x54, 0x4d, 0x0c, 0x6f, 0x77, 0x70, 0xf0,
	0xa5, 0xdf, 0x42, 0x9f, 0x7e, 0x7a, 0xcc, 0x94, 0xb4, 0xc8, 0x74, 0xca, 0x7c, 0x1d, 0xbb, 0xc1,
	0x30, 0x61, 0xd9, 0x52, 0x18, 0x9f, 0xc1, 0x05, 0x30, 0x77, 0xb3, 0x21, 0x6f, 0x17, 0xaf, 0x81,
	0x2c, 0xfd, 0x81, 0x1c, 0xa2, 0x46, 0xa1, 0xa2, 0x5b, 0xe6, 0x0f, 0xfc, 0x31, 0x3f, 0xf6, 0x1e,
	0x1f, 0xd0, 0xf7, 0x27, 0x4e, 0xc3, 0xf2, 0x22, 0x9b, 0x81, 0x71, 0xbc, 0xb5, 0x87, 0x57, 0x87,
	0x9c, 0x7a, 0xc8, 0xf4, 0x6f, 0x68, 0x4d, 0x1e, 0x55, 0x68, 0xb8, 0x5a, 0x42, 0x8c, 0x90, 0x6c,
	0x3c, 0x7c, 0x3f, 0x82, 0x87, 0xde, 0xe3, 0x89, 0xf7, 0xf8, 0x88, 0x3e, 0x3e, 0xd7, 0x25, 0xcb,
	0x44, 0x7e, 0xcd, 0x7c, 0x96, 0x95, 0xc0, 0xaa, 0x42, 0x86, 0x9a, 0x19, 0x88, 0x41, 0xae, 0x20,
	0xe4, 0x5d, 0x8f, 0x1c, 0xd6, 0xc4, 0x1b, 0x43, 0xdf, 0x46, 0xfa, 0xe3, 0x1b, 0x67, 0xa9, 0x7c,
	0xba, 0xde, 0xe7, 0x13, 0xef, 0xf3, 0x98, 0x1e, 0xbd, 0xe9, 0x53, 0xd3, 0x13, 0x66, 0x35, 0x4b,
	0x85, 0xb9, 0x7d, 0x9c, 0x8a, 0xfe, 0x9c, 0x6c, 0x2b, 0x3d, 0xb7, 0xc1, 0xa1, 0xbf, 0x96, 0x7b,
	0xf5, 0xb5, 0xbc, 0xd0, 0xf3, 0x97, 0x60, 0xad, 0x98, 0xc3, 0x20, 0xf0, 0x06, 0x94, 0xb6, 0x5f,
	0xd4, 0xd7, 0xe1, 0xe4, 0x61, 0xb7, 0x71, 0x87, 0xfb, 0x42, 0xfa, 0x88, 0x10, 0x0b, 0xd6, 0x4a,
	0x9d, 0x47, 0x32, 0x09, 0xda, 0xd5, 0x80, 0xd6, 0x91, 0x51, 0x42, 0x43, 0xb2, 0x63, 0x51, 0x20,
	0x04, 0x07, 0xac, 0xd1, 0x3b, 0x78, 0x1a, 0xd4, 0x06, 0xaf, 0x4d, 0x78, 0x78, 0xe1, 0xf2, 0xbc,
	0x92, 0xd1, 0x57, 0x64, 0xd7, 0x2d, 0x0a, 0x1b, 0xbc, 0xed, 0x50, 0x83, 0xaf, 0xbc, 0xfd, 0x67,
	0xb4, 0xef, 0xd5, 0x39, 0x5a, 0xb7, 0x05, 0x91, 0x33, 0x30, 0x46, 0x1b, 0x56, 0x49, 0x99, 0x1b,
	0x41, 0x36, 0xbb, 0xf6, 0xf3, 0x10, 0x2b, 0x09, 0x39, 0x86, 0xbc, 0xc6, 0xd0, 0x0f, 0x49, 0xbb,
	0xb0, 0x60, 0xa2, 0x5c, 0xa3, 0x4c, 0x25, 0x24, 0x41, 0x87, 0x35, 0x7a, 0xfb, 0xfc, 0xae, 0x0b,
	0x8e, 0xeb, 0x98, 0x7b, 0x46, 0x22, 0x46, 0xb9, 0xaa, 0x9f, 0xd1, 0xbd, 0x7f, 0x79, 0x46, 0x95,
	0xce, 0x3f, 0xa3, 0x39, 0x69, 0x0a, 0x83, 0x32, 0x15, 0x31, 0xda, 0xe0, 0x5d, 0xb6, 0xd5, 0x6b,
	0x0e, 0x46, 0xbe, 0xe6, 0x84, 0x1e, 0x8f, 0x36, 0xd3, 0xc9, 0x0c, 0x60, 0x61, 0x72, 0xcb, 0x04,
	0xb3, 0xe0, 0x5b, 0xb8, 0xae, 0x62, 0xb8, 0x90, 0x96, 0xa5, 0x12, 0x54, 0xc2, 0x4a, 0xa9, 0x54,
	0x35, 0xf3, 0xb8, 0x80, 0x2c, 0xe4, 0x1b, 0x36, 0xfd, 0xa5, 0x41, 0xee, 0xaf, 0xff, 0x45, 0xa5,
	0xc4, 0x45, 0x64, 0xc0, 0x16, 0x0a, 0x6d, 0x70, 0xdf, 0xdb, 0x4a, 0x6f, 0x1b, 0x53, 0x31, 0x71,
	0x9e, 0x85, 0x52, 0xcc, 0x7d, 0x6e, 0xdc, 0x88, 0xb9, 0x4d, 0x6c, 0xec, 0x5c, 0xa9, 0x0b, 0x49,
	0xc3, 0x6a, 0x40, 0xc8, 0x26, 0x6e, 0x0b, 0x28, 0x2e, 0xc1, 0x56, 0x3b, 0xd6, 0x26, 0x5b, 0x57,
	0x30, 0xf7, 0x05, 0xeb, 0x5b, 0x5d, 0x98, 0x18, 0xfc, 0x3a, 0xe4, 0x87, 0x6b, 0xda, 0xf7, 0x12,
	0x17, 0xbc, 0xa2, 0x1c, 0x9d, 0x93, 0x1d, 0x7f, 0x8d, 0xb4, 0x49, 0x76, 0xa6, 0xe3, 0x8b, 0xe1,
	0xa4, 0xf3, 0x16, 0x6d, 0x91, 0x3d, 0x3e, 0x1d, 0x8f, 0x47, 0xe3, 0x6f, 0x3a, 0x0d, 0x7a, 0x40,
	0xc8, 0x64, 0xc8, 0x5f, 0x8e, 0xc6, 0xc7, 0x93, 0xe1, 0x69, 0xe7, 0x8e, 0xd3, 0x0d, 0x39, 0x7f,
	0xc5, 0x3b, 0x5b, 0xf4, 0x2e, 0xd9, 0x3f, 0xe6, 0x27, 0xe7, 0xa3, 0xef, 0x86, 0xa7, 0x9d, 0xed,
	0xd9, 0xae, 0x1f, 0x90, 0xcf, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x17, 0x76, 0xa9, 0xf6, 0x2f,
	0x06, 0x00, 0x00,
}
