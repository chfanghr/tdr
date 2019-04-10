// Code generated by protoc-gen-go. DO NOT EDIT.
// source: playlist4content.proto

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

type Item struct {
	Uri                  *string         `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	Attributes           *ItemAttributes `protobuf:"bytes,2,opt,name=attributes" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_c17ee0ddf0ac79b8, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *Item) GetAttributes() *ItemAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type ListItems struct {
	Pos                  *int32   `protobuf:"varint,1,opt,name=pos" json:"pos,omitempty"`
	Truncated            *bool    `protobuf:"varint,2,opt,name=truncated" json:"truncated,omitempty"`
	Items                []*Item  `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListItems) Reset()         { *m = ListItems{} }
func (m *ListItems) String() string { return proto.CompactTextString(m) }
func (*ListItems) ProtoMessage()    {}
func (*ListItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_c17ee0ddf0ac79b8, []int{1}
}

func (m *ListItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListItems.Unmarshal(m, b)
}
func (m *ListItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListItems.Marshal(b, m, deterministic)
}
func (m *ListItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListItems.Merge(m, src)
}
func (m *ListItems) XXX_Size() int {
	return xxx_messageInfo_ListItems.Size(m)
}
func (m *ListItems) XXX_DiscardUnknown() {
	xxx_messageInfo_ListItems.DiscardUnknown(m)
}

var xxx_messageInfo_ListItems proto.InternalMessageInfo

func (m *ListItems) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func (m *ListItems) GetTruncated() bool {
	if m != nil && m.Truncated != nil {
		return *m.Truncated
	}
	return false
}

func (m *ListItems) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type ContentRange struct {
	Pos                  *int32   `protobuf:"varint,1,opt,name=pos" json:"pos,omitempty"`
	Length               *int32   `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentRange) Reset()         { *m = ContentRange{} }
func (m *ContentRange) String() string { return proto.CompactTextString(m) }
func (*ContentRange) ProtoMessage()    {}
func (*ContentRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_c17ee0ddf0ac79b8, []int{2}
}

func (m *ContentRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentRange.Unmarshal(m, b)
}
func (m *ContentRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentRange.Marshal(b, m, deterministic)
}
func (m *ContentRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentRange.Merge(m, src)
}
func (m *ContentRange) XXX_Size() int {
	return xxx_messageInfo_ContentRange.Size(m)
}
func (m *ContentRange) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentRange.DiscardUnknown(m)
}

var xxx_messageInfo_ContentRange proto.InternalMessageInfo

func (m *ContentRange) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func (m *ContentRange) GetLength() int32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

type ListContentSelection struct {
	WantRevision          *bool                  `protobuf:"varint,1,opt,name=wantRevision" json:"wantRevision,omitempty"`
	WantLength            *bool                  `protobuf:"varint,2,opt,name=wantLength" json:"wantLength,omitempty"`
	WantAttributes        *bool                  `protobuf:"varint,3,opt,name=wantAttributes" json:"wantAttributes,omitempty"`
	WantChecksum          *bool                  `protobuf:"varint,4,opt,name=wantChecksum" json:"wantChecksum,omitempty"`
	WantContent           *bool                  `protobuf:"varint,5,opt,name=wantContent" json:"wantContent,omitempty"`
	ContentRange          *ContentRange          `protobuf:"bytes,6,opt,name=contentRange" json:"contentRange,omitempty"`
	WantDiff              *bool                  `protobuf:"varint,7,opt,name=wantDiff" json:"wantDiff,omitempty"`
	BaseRevision          []byte                 `protobuf:"bytes,8,opt,name=baseRevision" json:"baseRevision,omitempty"`
	HintRevision          []byte                 `protobuf:"bytes,9,opt,name=hintRevision" json:"hintRevision,omitempty"`
	WantNothingIfUpToDate *bool                  `protobuf:"varint,10,opt,name=wantNothingIfUpToDate" json:"wantNothingIfUpToDate,omitempty"`
	WantResolveAction     *bool                  `protobuf:"varint,12,opt,name=wantResolveAction" json:"wantResolveAction,omitempty"`
	Issues                []*ClientIssue         `protobuf:"bytes,13,rep,name=issues" json:"issues,omitempty"`
	ResolveAction         []*ClientResolveAction `protobuf:"bytes,14,rep,name=resolveAction" json:"resolveAction,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *ListContentSelection) Reset()         { *m = ListContentSelection{} }
func (m *ListContentSelection) String() string { return proto.CompactTextString(m) }
func (*ListContentSelection) ProtoMessage()    {}
func (*ListContentSelection) Descriptor() ([]byte, []int) {
	return fileDescriptor_c17ee0ddf0ac79b8, []int{3}
}

func (m *ListContentSelection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListContentSelection.Unmarshal(m, b)
}
func (m *ListContentSelection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListContentSelection.Marshal(b, m, deterministic)
}
func (m *ListContentSelection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListContentSelection.Merge(m, src)
}
func (m *ListContentSelection) XXX_Size() int {
	return xxx_messageInfo_ListContentSelection.Size(m)
}
func (m *ListContentSelection) XXX_DiscardUnknown() {
	xxx_messageInfo_ListContentSelection.DiscardUnknown(m)
}

var xxx_messageInfo_ListContentSelection proto.InternalMessageInfo

func (m *ListContentSelection) GetWantRevision() bool {
	if m != nil && m.WantRevision != nil {
		return *m.WantRevision
	}
	return false
}

func (m *ListContentSelection) GetWantLength() bool {
	if m != nil && m.WantLength != nil {
		return *m.WantLength
	}
	return false
}

func (m *ListContentSelection) GetWantAttributes() bool {
	if m != nil && m.WantAttributes != nil {
		return *m.WantAttributes
	}
	return false
}

func (m *ListContentSelection) GetWantChecksum() bool {
	if m != nil && m.WantChecksum != nil {
		return *m.WantChecksum
	}
	return false
}

func (m *ListContentSelection) GetWantContent() bool {
	if m != nil && m.WantContent != nil {
		return *m.WantContent
	}
	return false
}

func (m *ListContentSelection) GetContentRange() *ContentRange {
	if m != nil {
		return m.ContentRange
	}
	return nil
}

func (m *ListContentSelection) GetWantDiff() bool {
	if m != nil && m.WantDiff != nil {
		return *m.WantDiff
	}
	return false
}

func (m *ListContentSelection) GetBaseRevision() []byte {
	if m != nil {
		return m.BaseRevision
	}
	return nil
}

func (m *ListContentSelection) GetHintRevision() []byte {
	if m != nil {
		return m.HintRevision
	}
	return nil
}

func (m *ListContentSelection) GetWantNothingIfUpToDate() bool {
	if m != nil && m.WantNothingIfUpToDate != nil {
		return *m.WantNothingIfUpToDate
	}
	return false
}

func (m *ListContentSelection) GetWantResolveAction() bool {
	if m != nil && m.WantResolveAction != nil {
		return *m.WantResolveAction
	}
	return false
}

func (m *ListContentSelection) GetIssues() []*ClientIssue {
	if m != nil {
		return m.Issues
	}
	return nil
}

func (m *ListContentSelection) GetResolveAction() []*ClientResolveAction {
	if m != nil {
		return m.ResolveAction
	}
	return nil
}

func init() {
	proto.RegisterType((*Item)(nil), "proto.Item")
	proto.RegisterType((*ListItems)(nil), "proto.ListItems")
	proto.RegisterType((*ContentRange)(nil), "proto.ContentRange")
	proto.RegisterType((*ListContentSelection)(nil), "proto.ListContentSelection")
}

func init() { proto.RegisterFile("playlist4content.proto", fileDescriptor_c17ee0ddf0ac79b8) }

var fileDescriptor_c17ee0ddf0ac79b8 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0x48, 0x1c, 0x9c, 0x89, 0x5b, 0xc1, 0x94, 0x54, 0xab, 0x08, 0x21, 0xe3, 0x03, 0x8a,
	0x10, 0xea, 0xa1, 0x2a, 0x82, 0x23, 0x55, 0x7b, 0x89, 0x54, 0x81, 0xb4, 0xc0, 0x15, 0xc9, 0x35,
	0x93, 0x64, 0x85, 0xb3, 0x1b, 0x79, 0xc7, 0x45, 0x3c, 0x19, 0xaf, 0x87, 0x76, 0x37, 0xf1, 0x0f,
	0xcd, 0xc9, 0x9e, 0xef, 0x6f, 0xd6, 0x9f, 0x17, 0xce, 0x77, 0x65, 0xfe, 0xa7, 0x54, 0x96, 0xaf,
	0x0a, 0xa3, 0x99, 0x34, 0x5f, 0xec, 0x2a, 0xc3, 0x06, 0x23, 0xff, 0x98, 0x9f, 0x35, 0xf4, 0x96,
	0x38, 0x0f, 0xdc, 0x7c, 0xd6, 0x80, 0xca, 0xda, 0x9a, 0x6c, 0x80, 0xb3, 0x2f, 0x30, 0x5a, 0x32,
	0x6d, 0xf1, 0x19, 0x0c, 0xeb, 0x4a, 0x89, 0x41, 0x3a, 0x58, 0x4c, 0xa4, 0x7b, 0xc5, 0xf7, 0x00,
	0x39, 0x73, 0xa5, 0xee, 0x6b, 0x26, 0x2b, 0x9e, 0xa4, 0x83, 0xc5, 0xf4, 0x72, 0x16, 0x5c, 0x17,
	0xce, 0x72, 0xdd, 0x90, 0xb2, 0x23, 0xcc, 0x7e, 0xc0, 0xe4, 0x4e, 0x59, 0x76, 0x0a, 0xeb, 0x52,
	0x77, 0xc6, 0xfa, 0xd4, 0x48, 0xba, 0x57, 0x7c, 0x09, 0x13, 0xae, 0x6a, 0x5d, 0xe4, 0x4c, 0x3f,
	0x7d, 0x68, 0x2c, 0x5b, 0x00, 0x5f, 0x43, 0xa4, 0x9c, 0x51, 0x0c, 0xd3, 0xe1, 0x62, 0x7a, 0x39,
	0xed, 0xac, 0x93, 0x81, 0xc9, 0x3e, 0x42, 0x72, 0x13, 0x3e, 0x5a, 0xe6, 0x7a, 0x4d, 0x47, 0x56,
	0x9c, 0xc3, 0xb8, 0x24, 0xbd, 0xe6, 0x8d, 0xcf, 0x8f, 0xe4, 0x7e, 0xca, 0xfe, 0x8e, 0xe0, 0x85,
	0x3b, 0xda, 0xde, 0xfe, 0x95, 0x4a, 0x2a, 0x58, 0x19, 0x8d, 0x19, 0x24, 0xbf, 0x73, 0xcd, 0x92,
	0x1e, 0x94, 0x55, 0x46, 0xfb, 0xac, 0x58, 0xf6, 0x30, 0x7c, 0x05, 0xe0, 0xe6, 0xbb, 0x36, 0x38,
	0x96, 0x1d, 0x04, 0xdf, 0xc0, 0xa9, 0x9b, 0xda, 0x52, 0xc4, 0xd0, 0x6b, 0xfe, 0x43, 0x0f, 0xbb,
	0x6e, 0x36, 0x54, 0xfc, 0xb2, 0xf5, 0x56, 0x8c, 0xda, 0x5d, 0x07, 0x0c, 0x53, 0x98, 0xfa, 0x39,
	0x9c, 0x53, 0x44, 0x5e, 0xd2, 0x85, 0xf0, 0x03, 0x24, 0x45, 0xa7, 0x04, 0x31, 0xf6, 0x7f, 0xe7,
	0x6c, 0x5f, 0x57, 0xb7, 0x1f, 0xd9, 0x13, 0xe2, 0x1c, 0x62, 0x97, 0x73, 0xab, 0x56, 0x2b, 0xf1,
	0xd4, 0xe7, 0x36, 0xb3, 0x3b, 0xda, 0x7d, 0x6e, 0xa9, 0xa9, 0x21, 0x4e, 0x07, 0x8b, 0x44, 0xf6,
	0x30, 0xa7, 0xd9, 0xa8, 0x4e, 0x55, 0x93, 0xa0, 0xe9, 0x62, 0x78, 0x05, 0x33, 0x97, 0xf9, 0xd9,
	0xf0, 0x46, 0xe9, 0xf5, 0x72, 0xf5, 0x7d, 0xf7, 0xcd, 0xdc, 0xe6, 0x4c, 0x02, 0xfc, 0xc2, 0xe3,
	0x24, 0xbe, 0x83, 0xe7, 0xa1, 0x70, 0x6b, 0xca, 0x07, 0xba, 0xf6, 0x7f, 0x46, 0x24, 0xde, 0xf1,
	0x98, 0xc0, 0xb7, 0x30, 0x0e, 0xd7, 0x58, 0x9c, 0xf8, 0x9b, 0x82, 0x87, 0x4f, 0x2f, 0x15, 0x69,
	0x5e, 0x3a, 0x4a, 0xee, 0x15, 0xf8, 0x09, 0x4e, 0xaa, 0x5e, 0xea, 0xa9, 0xb7, 0xcc, 0x7b, 0x96,
	0x5e, 0xbc, 0xec, 0x1b, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x98, 0x5e, 0xc2, 0xeb, 0x70, 0x03,
	0x00, 0x00,
}