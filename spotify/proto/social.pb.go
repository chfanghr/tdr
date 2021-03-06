// Code generated by protoc-gen-go. DO NOT EDIT.
// source: social.proto

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

type DecorationData struct {
	Username             *string  `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	FullName             *string  `protobuf:"bytes,2,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	ImageUrl             *string  `protobuf:"bytes,3,opt,name=image_url,json=imageUrl" json:"image_url,omitempty"`
	LargeImageUrl        *string  `protobuf:"bytes,5,opt,name=large_image_url,json=largeImageUrl" json:"large_image_url,omitempty"`
	FirstName            *string  `protobuf:"bytes,6,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName             *string  `protobuf:"bytes,7,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	FacebookUid          *string  `protobuf:"bytes,8,opt,name=facebook_uid,json=facebookUid" json:"facebook_uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecorationData) Reset()         { *m = DecorationData{} }
func (m *DecorationData) String() string { return proto.CompactTextString(m) }
func (*DecorationData) ProtoMessage()    {}
func (*DecorationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c23ae38a95c1cd21, []int{0}
}

func (m *DecorationData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecorationData.Unmarshal(m, b)
}
func (m *DecorationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecorationData.Marshal(b, m, deterministic)
}
func (m *DecorationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecorationData.Merge(m, src)
}
func (m *DecorationData) XXX_Size() int {
	return xxx_messageInfo_DecorationData.Size(m)
}
func (m *DecorationData) XXX_DiscardUnknown() {
	xxx_messageInfo_DecorationData.DiscardUnknown(m)
}

var xxx_messageInfo_DecorationData proto.InternalMessageInfo

func (m *DecorationData) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *DecorationData) GetFullName() string {
	if m != nil && m.FullName != nil {
		return *m.FullName
	}
	return ""
}

func (m *DecorationData) GetImageUrl() string {
	if m != nil && m.ImageUrl != nil {
		return *m.ImageUrl
	}
	return ""
}

func (m *DecorationData) GetLargeImageUrl() string {
	if m != nil && m.LargeImageUrl != nil {
		return *m.LargeImageUrl
	}
	return ""
}

func (m *DecorationData) GetFirstName() string {
	if m != nil && m.FirstName != nil {
		return *m.FirstName
	}
	return ""
}

func (m *DecorationData) GetLastName() string {
	if m != nil && m.LastName != nil {
		return *m.LastName
	}
	return ""
}

func (m *DecorationData) GetFacebookUid() string {
	if m != nil && m.FacebookUid != nil {
		return *m.FacebookUid
	}
	return ""
}

func init() {
	proto.RegisterType((*DecorationData)(nil), "proto.DecorationData")
}

func init() { proto.RegisterFile("social.proto", fileDescriptor_c23ae38a95c1cd21) }

var fileDescriptor_c23ae38a95c1cd21 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8d, 0xb1, 0x0a, 0xc2, 0x30,
	0x10, 0x86, 0xa9, 0x52, 0x6d, 0xcf, 0xaa, 0x90, 0xa9, 0x28, 0x82, 0x3a, 0x88, 0x93, 0x4f, 0xd1,
	0xc5, 0xc5, 0x41, 0xe8, 0x1c, 0xce, 0x36, 0x2d, 0xc1, 0xb4, 0x91, 0x34, 0x79, 0x6c, 0xdf, 0x41,
	0x7a, 0x31, 0x3a, 0x1d, 0xf7, 0x7d, 0x1f, 0xfc, 0x90, 0x0d, 0xba, 0x92, 0xa8, 0x2e, 0x2f, 0xa3,
	0xad, 0x66, 0x31, 0x9d, 0xe3, 0x3b, 0x82, 0x55, 0x21, 0x2a, 0x6d, 0xd0, 0x4a, 0xdd, 0x17, 0x68,
	0x91, 0x6d, 0x20, 0x71, 0x83, 0x30, 0x3d, 0x76, 0x22, 0x8f, 0xf6, 0xd1, 0x39, 0xbd, 0xff, 0x7e,
	0xb6, 0x85, 0xb4, 0x71, 0x4a, 0x71, 0x92, 0x13, 0x2f, 0x47, 0x70, 0xfb, 0x4a, 0xd9, 0x61, 0x2b,
	0xb8, 0x33, 0x2a, 0x9f, 0x7a, 0x49, 0xa0, 0x34, 0x8a, 0x9d, 0x60, 0xad, 0xd0, 0xb4, 0x82, 0xff,
	0x93, 0x98, 0x92, 0x25, 0xe1, 0x6b, 0xe8, 0x76, 0x00, 0x8d, 0x34, 0x83, 0xf5, 0x13, 0x33, 0x4a,
	0x52, 0x22, 0x61, 0x43, 0x61, 0xb0, 0x73, 0xbf, 0x31, 0x02, 0x92, 0x07, 0xc8, 0x1a, 0xac, 0xc4,
	0x43, 0xeb, 0x27, 0x77, 0xb2, 0xce, 0x13, 0xf2, 0x8b, 0xc0, 0x4a, 0x59, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x4d, 0x1f, 0x26, 0xff, 0x05, 0x01, 0x00, 0x00,
}
