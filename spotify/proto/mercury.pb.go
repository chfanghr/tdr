// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mercury.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type MercuryReply_CachePolicy int32

const (
	MercuryReply_CACHE_NO      MercuryReply_CachePolicy = 1
	MercuryReply_CACHE_PRIVATE MercuryReply_CachePolicy = 2
	MercuryReply_CACHE_PUBLIC  MercuryReply_CachePolicy = 3
)

var MercuryReply_CachePolicy_name = map[int32]string{
	1: "CACHE_NO",
	2: "CACHE_PRIVATE",
	3: "CACHE_PUBLIC",
}

var MercuryReply_CachePolicy_value = map[string]int32{
	"CACHE_NO":      1,
	"CACHE_PRIVATE": 2,
	"CACHE_PUBLIC":  3,
}

func (x MercuryReply_CachePolicy) Enum() *MercuryReply_CachePolicy {
	p := new(MercuryReply_CachePolicy)
	*p = x
	return p
}

func (x MercuryReply_CachePolicy) String() string {
	return proto.EnumName(MercuryReply_CachePolicy_name, int32(x))
}

func (x *MercuryReply_CachePolicy) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MercuryReply_CachePolicy_value, data, "MercuryReply_CachePolicy")
	if err != nil {
		return err
	}
	*x = MercuryReply_CachePolicy(value)
	return nil
}

func (MercuryReply_CachePolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0265ec996b5e6b6, []int{3, 0}
}

type MercuryMultiGetRequest struct {
	Request              []*MercuryRequest `protobuf:"bytes,1,rep,name=request" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MercuryMultiGetRequest) Reset()         { *m = MercuryMultiGetRequest{} }
func (m *MercuryMultiGetRequest) String() string { return proto.CompactTextString(m) }
func (*MercuryMultiGetRequest) ProtoMessage()    {}
func (*MercuryMultiGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0265ec996b5e6b6, []int{0}
}

func (m *MercuryMultiGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MercuryMultiGetRequest.Unmarshal(m, b)
}
func (m *MercuryMultiGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MercuryMultiGetRequest.Marshal(b, m, deterministic)
}
func (m *MercuryMultiGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MercuryMultiGetRequest.Merge(m, src)
}
func (m *MercuryMultiGetRequest) XXX_Size() int {
	return xxx_messageInfo_MercuryMultiGetRequest.Size(m)
}
func (m *MercuryMultiGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MercuryMultiGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MercuryMultiGetRequest proto.InternalMessageInfo

func (m *MercuryMultiGetRequest) GetRequest() []*MercuryRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

type MercuryMultiGetReply struct {
	Reply                []*MercuryReply `protobuf:"bytes,1,rep,name=reply" json:"reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MercuryMultiGetReply) Reset()         { *m = MercuryMultiGetReply{} }
func (m *MercuryMultiGetReply) String() string { return proto.CompactTextString(m) }
func (*MercuryMultiGetReply) ProtoMessage()    {}
func (*MercuryMultiGetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0265ec996b5e6b6, []int{1}
}

func (m *MercuryMultiGetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MercuryMultiGetReply.Unmarshal(m, b)
}
func (m *MercuryMultiGetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MercuryMultiGetReply.Marshal(b, m, deterministic)
}
func (m *MercuryMultiGetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MercuryMultiGetReply.Merge(m, src)
}
func (m *MercuryMultiGetReply) XXX_Size() int {
	return xxx_messageInfo_MercuryMultiGetReply.Size(m)
}
func (m *MercuryMultiGetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MercuryMultiGetReply.DiscardUnknown(m)
}

var xxx_messageInfo_MercuryMultiGetReply proto.InternalMessageInfo

func (m *MercuryMultiGetReply) GetReply() []*MercuryReply {
	if m != nil {
		return m.Reply
	}
	return nil
}

type MercuryRequest struct {
	Uri                  *string  `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	ContentType          *string  `protobuf:"bytes,2,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Body                 []byte   `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
	Etag                 []byte   `protobuf:"bytes,4,opt,name=etag" json:"etag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MercuryRequest) Reset()         { *m = MercuryRequest{} }
func (m *MercuryRequest) String() string { return proto.CompactTextString(m) }
func (*MercuryRequest) ProtoMessage()    {}
func (*MercuryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0265ec996b5e6b6, []int{2}
}

func (m *MercuryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MercuryRequest.Unmarshal(m, b)
}
func (m *MercuryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MercuryRequest.Marshal(b, m, deterministic)
}
func (m *MercuryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MercuryRequest.Merge(m, src)
}
func (m *MercuryRequest) XXX_Size() int {
	return xxx_messageInfo_MercuryRequest.Size(m)
}
func (m *MercuryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MercuryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MercuryRequest proto.InternalMessageInfo

func (m *MercuryRequest) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *MercuryRequest) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func (m *MercuryRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *MercuryRequest) GetEtag() []byte {
	if m != nil {
		return m.Etag
	}
	return nil
}

type MercuryReply struct {
	StatusCode           *int32                    `protobuf:"zigzag32,1,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	StatusMessage        *string                   `protobuf:"bytes,2,opt,name=status_message,json=statusMessage" json:"status_message,omitempty"`
	CachePolicy          *MercuryReply_CachePolicy `protobuf:"varint,3,opt,name=cache_policy,json=cachePolicy,enum=proto.MercuryReply_CachePolicy" json:"cache_policy,omitempty"`
	Ttl                  *int32                    `protobuf:"zigzag32,4,opt,name=ttl" json:"ttl,omitempty"`
	Etag                 []byte                    `protobuf:"bytes,5,opt,name=etag" json:"etag,omitempty"`
	ContentType          *string                   `protobuf:"bytes,6,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Body                 []byte                    `protobuf:"bytes,7,opt,name=body" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MercuryReply) Reset()         { *m = MercuryReply{} }
func (m *MercuryReply) String() string { return proto.CompactTextString(m) }
func (*MercuryReply) ProtoMessage()    {}
func (*MercuryReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0265ec996b5e6b6, []int{3}
}

func (m *MercuryReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MercuryReply.Unmarshal(m, b)
}
func (m *MercuryReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MercuryReply.Marshal(b, m, deterministic)
}
func (m *MercuryReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MercuryReply.Merge(m, src)
}
func (m *MercuryReply) XXX_Size() int {
	return xxx_messageInfo_MercuryReply.Size(m)
}
func (m *MercuryReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MercuryReply.DiscardUnknown(m)
}

var xxx_messageInfo_MercuryReply proto.InternalMessageInfo

func (m *MercuryReply) GetStatusCode() int32 {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return 0
}

func (m *MercuryReply) GetStatusMessage() string {
	if m != nil && m.StatusMessage != nil {
		return *m.StatusMessage
	}
	return ""
}

func (m *MercuryReply) GetCachePolicy() MercuryReply_CachePolicy {
	if m != nil && m.CachePolicy != nil {
		return *m.CachePolicy
	}
	return MercuryReply_CACHE_NO
}

func (m *MercuryReply) GetTtl() int32 {
	if m != nil && m.Ttl != nil {
		return *m.Ttl
	}
	return 0
}

func (m *MercuryReply) GetEtag() []byte {
	if m != nil {
		return m.Etag
	}
	return nil
}

func (m *MercuryReply) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func (m *MercuryReply) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Header struct {
	Uri                  *string      `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	ContentType          *string      `protobuf:"bytes,2,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Method               *string      `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
	StatusCode           *int32       `protobuf:"zigzag32,4,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	UserFields           []*UserField `protobuf:"bytes,6,rep,name=user_fields,json=userFields" json:"user_fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0265ec996b5e6b6, []int{4}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *Header) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func (m *Header) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *Header) GetStatusCode() int32 {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return 0
}

func (m *Header) GetUserFields() []*UserField {
	if m != nil {
		return m.UserFields
	}
	return nil
}

type UserField struct {
	Key                  *string  `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserField) Reset()         { *m = UserField{} }
func (m *UserField) String() string { return proto.CompactTextString(m) }
func (*UserField) ProtoMessage()    {}
func (*UserField) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0265ec996b5e6b6, []int{5}
}

func (m *UserField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserField.Unmarshal(m, b)
}
func (m *UserField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserField.Marshal(b, m, deterministic)
}
func (m *UserField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserField.Merge(m, src)
}
func (m *UserField) XXX_Size() int {
	return xxx_messageInfo_UserField.Size(m)
}
func (m *UserField) XXX_DiscardUnknown() {
	xxx_messageInfo_UserField.DiscardUnknown(m)
}

var xxx_messageInfo_UserField proto.InternalMessageInfo

func (m *UserField) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *UserField) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.MercuryReply_CachePolicy", MercuryReply_CachePolicy_name, MercuryReply_CachePolicy_value)
	proto.RegisterType((*MercuryMultiGetRequest)(nil), "proto.MercuryMultiGetRequest")
	proto.RegisterType((*MercuryMultiGetReply)(nil), "proto.MercuryMultiGetReply")
	proto.RegisterType((*MercuryRequest)(nil), "proto.MercuryRequest")
	proto.RegisterType((*MercuryReply)(nil), "proto.MercuryReply")
	proto.RegisterType((*Header)(nil), "proto.Header")
	proto.RegisterType((*UserField)(nil), "proto.UserField")
}

func init() { proto.RegisterFile("mercury.proto", fileDescriptor_a0265ec996b5e6b6) }

var fileDescriptor_a0265ec996b5e6b6 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x73, 0x2b, 0x19, 0x3b, 0x91, 0xb3, 0x94, 0xca, 0x6f, 0x0d, 0x96, 0x90, 0xc2, 0x4b,
	0x10, 0xe5, 0x07, 0x48, 0xad, 0x42, 0x23, 0x11, 0xa8, 0x56, 0x2d, 0xaf, 0x96, 0xb1, 0x87, 0x36,
	0xaa, 0xdd, 0x35, 0x7b, 0x41, 0xda, 0x4f, 0xe2, 0x13, 0xf8, 0x3b, 0xb4, 0x17, 0x37, 0xa5, 0x81,
	0x97, 0x3e, 0xf9, 0xcc, 0x99, 0xe3, 0x99, 0x33, 0x33, 0x0b, 0x93, 0x06, 0x79, 0xa9, 0xb8, 0x5e,
	0xb6, 0x9c, 0x49, 0x46, 0x86, 0xf6, 0x93, 0xae, 0xe1, 0x68, 0xe3, 0xf8, 0x8d, 0xaa, 0xe5, 0xf6,
	0x23, 0x4a, 0x8a, 0x3f, 0x14, 0x0a, 0x49, 0xde, 0xc0, 0x01, 0x77, 0x30, 0x09, 0xe6, 0xfd, 0x45,
	0x78, 0xf2, 0xc2, 0xfd, 0xb9, 0xf4, 0x7a, 0xaf, 0xa3, 0x9d, 0x2a, 0x5d, 0xc1, 0xe1, 0x5e, 0xa9,
	0xb6, 0xd6, 0xe4, 0x35, 0x0c, 0xb9, 0x01, 0xbe, 0xcc, 0xf3, 0xc7, 0x65, 0xda, 0x5a, 0x53, 0xa7,
	0x48, 0x1b, 0x98, 0xfe, 0x5d, 0x9d, 0xc4, 0xd0, 0x57, 0x7c, 0x9b, 0x04, 0xf3, 0x60, 0x31, 0xa6,
	0x06, 0x92, 0x97, 0x10, 0x95, 0xec, 0x4e, 0xe2, 0x9d, 0xcc, 0xa5, 0x6e, 0x31, 0xe9, 0xd9, 0x54,
	0xe8, 0xb9, 0x4b, 0xdd, 0x22, 0x21, 0x30, 0xf8, 0xc6, 0x2a, 0x9d, 0xf4, 0xe7, 0xc1, 0x22, 0xa2,
	0x16, 0x1b, 0x0e, 0x65, 0x71, 0x9d, 0x0c, 0x1c, 0x67, 0x70, 0xfa, 0xbb, 0x07, 0xd1, 0x43, 0x1b,
	0xe4, 0x18, 0x42, 0x21, 0x0b, 0xa9, 0x44, 0x5e, 0xb2, 0x0a, 0x6d, 0xd7, 0x19, 0x05, 0x47, 0x65,
	0xac, 0x42, 0xf2, 0x0a, 0xa6, 0x5e, 0xd0, 0xa0, 0x10, 0xc5, 0x75, 0xd7, 0x7e, 0xe2, 0xd8, 0x8d,
	0x23, 0xc9, 0x29, 0x44, 0x65, 0x51, 0xde, 0x60, 0xde, 0xb2, 0x7a, 0x5b, 0x3a, 0x23, 0xd3, 0x93,
	0xe3, 0x7f, 0x4c, 0xbe, 0xcc, 0x8c, 0xee, 0xc2, 0xca, 0x68, 0x58, 0xee, 0x02, 0x33, 0xb9, 0x94,
	0xb5, 0xf5, 0x3b, 0xa3, 0x06, 0xde, 0x8f, 0x30, 0xdc, 0x8d, 0xb0, 0xb7, 0x8d, 0xd1, 0xff, 0xb7,
	0x71, 0xb0, 0xdb, 0x46, 0xfa, 0x1e, 0xc2, 0x07, 0x8d, 0x49, 0x04, 0xcf, 0xb2, 0x55, 0x76, 0x7e,
	0x96, 0x7f, 0xfe, 0x12, 0x07, 0x64, 0x06, 0x13, 0x17, 0x5d, 0xd0, 0xf5, 0xd7, 0xd5, 0xe5, 0x59,
	0xdc, 0x23, 0x31, 0x44, 0x9e, 0xba, 0x3a, 0xfd, 0xb4, 0xce, 0xe2, 0x7e, 0xfa, 0x2b, 0x80, 0xd1,
	0x39, 0x16, 0x15, 0xf2, 0xa7, 0xdd, 0xe8, 0x08, 0x46, 0x0d, 0xca, 0x1b, 0x56, 0xd9, 0xe5, 0x8c,
	0xa9, 0x8f, 0x1e, 0x9f, 0x60, 0xb0, 0x77, 0x82, 0xb7, 0x10, 0x2a, 0x81, 0x3c, 0xff, 0xbe, 0xc5,
	0xba, 0x12, 0xc9, 0xc8, 0x3e, 0xaa, 0xd8, 0xaf, 0xf6, 0x4a, 0x20, 0xff, 0x60, 0x12, 0x14, 0x54,
	0x07, 0x45, 0xfa, 0x0e, 0xc6, 0xf7, 0x09, 0xe3, 0xf6, 0x16, 0x75, 0xe7, 0xf6, 0x16, 0x35, 0x39,
	0x84, 0xe1, 0xcf, 0xa2, 0x56, 0xce, 0x66, 0x44, 0x5d, 0xf0, 0x27, 0x00, 0x00, 0xff, 0xff, 0x6b,
	0x13, 0x8c, 0x6a, 0x30, 0x03, 0x00, 0x00,
}
