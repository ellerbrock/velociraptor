// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transport.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "www.velocidex.com/golang/velociraptor/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The Velociraptor client sends back the buffer and the filename and
// the server saves the entire file directly in the file storage
// filesystem. This allows easy recovery as well as data expiration
// policies (since the filestore is just a directory on disk with
// regular files and timestamps).
type FileBuffer struct {
	Pathspec         *PathSpec `protobuf:"bytes,1,opt,name=pathspec" json:"pathspec,omitempty"`
	Offset           *uint64   `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	Data             []byte    `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	FlowId           *string   `protobuf:"bytes,4,opt,name=flow_id,json=flowId" json:"flow_id,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *FileBuffer) Reset()                    { *m = FileBuffer{} }
func (m *FileBuffer) String() string            { return proto1.CompactTextString(m) }
func (*FileBuffer) ProtoMessage()               {}
func (*FileBuffer) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *FileBuffer) GetPathspec() *PathSpec {
	if m != nil {
		return m.Pathspec
	}
	return nil
}

func (m *FileBuffer) GetOffset() uint64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *FileBuffer) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *FileBuffer) GetFlowId() string {
	if m != nil && m.FlowId != nil {
		return *m.FlowId
	}
	return ""
}

func init() {
	proto1.RegisterType((*FileBuffer)(nil), "proto.FileBuffer")
}

func init() { proto1.RegisterFile("transport.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8c, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x89, 0xd6, 0x51, 0xa3, 0x32, 0x90, 0x8d, 0x65, 0x56, 0x51, 0x37, 0x15, 0x21, 0x05,
	0x97, 0xee, 0x9c, 0x85, 0xe0, 0x4a, 0xa9, 0x07, 0x90, 0x98, 0xbe, 0x4c, 0x03, 0x69, 0x5e, 0x48,
	0x9e, 0xd6, 0x2b, 0x79, 0x0f, 0x4f, 0xa2, 0xd7, 0x70, 0x21, 0xa6, 0x32, 0xab, 0xf7, 0xfe, 0x8f,
	0xef, 0xff, 0xf9, 0x92, 0x92, 0x0e, 0x39, 0x62, 0x22, 0x15, 0x13, 0x12, 0x8a, 0xbd, 0x72, 0x56,
	0x27, 0xda, 0x90, 0xc3, 0x90, 0x67, 0xba, 0xba, 0x99, 0xa6, 0x49, 0xbd, 0x81, 0x47, 0xe3, 0x7a,
	0x78, 0x57, 0x06, 0xc7, 0x76, 0x83, 0x5e, 0x87, 0x4d, 0x3b, 0xc3, 0xa4, 0x23, 0x61, 0x6a, 0x8b,
	0xdc, 0x66, 0x18, 0x75, 0x20, 0x67, 0xe6, 0xee, 0xf9, 0x07, 0xe3, 0xfc, 0xce, 0x79, 0x58, 0xbf,
	0x5a, 0x0b, 0x49, 0x5c, 0xf1, 0x83, 0xa8, 0x69, 0xc8, 0x11, 0x4c, 0xcd, 0x24, 0x6b, 0x8e, 0xae,
	0x97, 0xb3, 0xa8, 0x1e, 0x35, 0x0d, 0x4f, 0x11, 0x4c, 0xb7, 0x15, 0xc4, 0x2d, 0x5f, 0xa0, 0xb5,
	0x19, 0xa8, 0xde, 0x91, 0xac, 0xa9, 0xd6, 0x97, 0x5f, 0x3f, 0xdf, 0x9f, 0xec, 0x42, 0x9c, 0x3d,
	0x14, 0x2a, 0xd1, 0x4a, 0x1a, 0x40, 0xbe, 0x94, 0x69, 0xe9, 0x42, 0x49, 0xd6, 0x79, 0x50, 0xdd,
	0x7f, 0x51, 0x08, 0x5e, 0xf5, 0x9a, 0x74, 0xbd, 0x2b, 0x59, 0x73, 0xdc, 0x95, 0x5f, 0x9c, 0xf2,
	0x7d, 0xeb, 0x71, 0x7a, 0x76, 0x7d, 0x5d, 0x49, 0xd6, 0x1c, 0x76, 0x8b, 0xbf, 0x78, 0xdf, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x55, 0x17, 0xce, 0x0f, 0x01, 0x00, 0x00,
}
