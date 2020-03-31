// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calendar.proto

// protoc -I api/ api/calendar.proto --go_out=plugins=grpc:internal/pkg/client

package pkg

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Event struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	DateStarted          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=date_started,json=dateStarted,proto3" json:"date_started,omitempty"`
	DateComplete         *timestamp.Timestamp `protobuf:"bytes,4,opt,name=date_complete,json=dateComplete,proto3" json:"date_complete,omitempty"`
	Notice               string               `protobuf:"bytes,5,opt,name=notice,proto3" json:"notice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Event) GetDateStarted() *timestamp.Timestamp {
	if m != nil {
		return m.DateStarted
	}
	return nil
}

func (m *Event) GetDateComplete() *timestamp.Timestamp {
	if m != nil {
		return m.DateComplete
	}
	return nil
}

func (m *Event) GetNotice() string {
	if m != nil {
		return m.Notice
	}
	return ""
}

type EventAddRequest struct {
	Title                string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	DateStarted          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date_started,json=dateStarted,proto3" json:"date_started,omitempty"`
	DateComplete         *timestamp.Timestamp `protobuf:"bytes,3,opt,name=date_complete,json=dateComplete,proto3" json:"date_complete,omitempty"`
	Notice               string               `protobuf:"bytes,4,opt,name=notice,proto3" json:"notice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EventAddRequest) Reset()         { *m = EventAddRequest{} }
func (m *EventAddRequest) String() string { return proto.CompactTextString(m) }
func (*EventAddRequest) ProtoMessage()    {}
func (*EventAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}

func (m *EventAddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventAddRequest.Unmarshal(m, b)
}
func (m *EventAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventAddRequest.Marshal(b, m, deterministic)
}
func (m *EventAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAddRequest.Merge(m, src)
}
func (m *EventAddRequest) XXX_Size() int {
	return xxx_messageInfo_EventAddRequest.Size(m)
}
func (m *EventAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventAddRequest proto.InternalMessageInfo

func (m *EventAddRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EventAddRequest) GetDateStarted() *timestamp.Timestamp {
	if m != nil {
		return m.DateStarted
	}
	return nil
}

func (m *EventAddRequest) GetDateComplete() *timestamp.Timestamp {
	if m != nil {
		return m.DateComplete
	}
	return nil
}

func (m *EventAddRequest) GetNotice() string {
	if m != nil {
		return m.Notice
	}
	return ""
}

type EventAddResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventAddResponse) Reset()         { *m = EventAddResponse{} }
func (m *EventAddResponse) String() string { return proto.CompactTextString(m) }
func (*EventAddResponse) ProtoMessage()    {}
func (*EventAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{2}
}

func (m *EventAddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventAddResponse.Unmarshal(m, b)
}
func (m *EventAddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventAddResponse.Marshal(b, m, deterministic)
}
func (m *EventAddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAddResponse.Merge(m, src)
}
func (m *EventAddResponse) XXX_Size() int {
	return xxx_messageInfo_EventAddResponse.Size(m)
}
func (m *EventAddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventAddResponse proto.InternalMessageInfo

func (m *EventAddResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EventEditRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Event                *Event   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventEditRequest) Reset()         { *m = EventEditRequest{} }
func (m *EventEditRequest) String() string { return proto.CompactTextString(m) }
func (*EventEditRequest) ProtoMessage()    {}
func (*EventEditRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{3}
}

func (m *EventEditRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventEditRequest.Unmarshal(m, b)
}
func (m *EventEditRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventEditRequest.Marshal(b, m, deterministic)
}
func (m *EventEditRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventEditRequest.Merge(m, src)
}
func (m *EventEditRequest) XXX_Size() int {
	return xxx_messageInfo_EventEditRequest.Size(m)
}
func (m *EventEditRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventEditRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventEditRequest proto.InternalMessageInfo

func (m *EventEditRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EventEditRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type EventEditResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventEditResponse) Reset()         { *m = EventEditResponse{} }
func (m *EventEditResponse) String() string { return proto.CompactTextString(m) }
func (*EventEditResponse) ProtoMessage()    {}
func (*EventEditResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{4}
}

func (m *EventEditResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventEditResponse.Unmarshal(m, b)
}
func (m *EventEditResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventEditResponse.Marshal(b, m, deterministic)
}
func (m *EventEditResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventEditResponse.Merge(m, src)
}
func (m *EventEditResponse) XXX_Size() int {
	return xxx_messageInfo_EventEditResponse.Size(m)
}
func (m *EventEditResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventEditResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventEditResponse proto.InternalMessageInfo

type EventDeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventDeleteRequest) Reset()         { *m = EventDeleteRequest{} }
func (m *EventDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*EventDeleteRequest) ProtoMessage()    {}
func (*EventDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{5}
}

func (m *EventDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventDeleteRequest.Unmarshal(m, b)
}
func (m *EventDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventDeleteRequest.Marshal(b, m, deterministic)
}
func (m *EventDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDeleteRequest.Merge(m, src)
}
func (m *EventDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_EventDeleteRequest.Size(m)
}
func (m *EventDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventDeleteRequest proto.InternalMessageInfo

func (m *EventDeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EventDeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventDeleteResponse) Reset()         { *m = EventDeleteResponse{} }
func (m *EventDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*EventDeleteResponse) ProtoMessage()    {}
func (*EventDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{6}
}

func (m *EventDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventDeleteResponse.Unmarshal(m, b)
}
func (m *EventDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventDeleteResponse.Marshal(b, m, deterministic)
}
func (m *EventDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDeleteResponse.Merge(m, src)
}
func (m *EventDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_EventDeleteResponse.Size(m)
}
func (m *EventDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventDeleteResponse proto.InternalMessageInfo

type EventGetByIdRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventGetByIdRequest) Reset()         { *m = EventGetByIdRequest{} }
func (m *EventGetByIdRequest) String() string { return proto.CompactTextString(m) }
func (*EventGetByIdRequest) ProtoMessage()    {}
func (*EventGetByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{7}
}

func (m *EventGetByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventGetByIdRequest.Unmarshal(m, b)
}
func (m *EventGetByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventGetByIdRequest.Marshal(b, m, deterministic)
}
func (m *EventGetByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventGetByIdRequest.Merge(m, src)
}
func (m *EventGetByIdRequest) XXX_Size() int {
	return xxx_messageInfo_EventGetByIdRequest.Size(m)
}
func (m *EventGetByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventGetByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventGetByIdRequest proto.InternalMessageInfo

func (m *EventGetByIdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EventGetByIdResponse struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventGetByIdResponse) Reset()         { *m = EventGetByIdResponse{} }
func (m *EventGetByIdResponse) String() string { return proto.CompactTextString(m) }
func (*EventGetByIdResponse) ProtoMessage()    {}
func (*EventGetByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{8}
}

func (m *EventGetByIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventGetByIdResponse.Unmarshal(m, b)
}
func (m *EventGetByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventGetByIdResponse.Marshal(b, m, deterministic)
}
func (m *EventGetByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventGetByIdResponse.Merge(m, src)
}
func (m *EventGetByIdResponse) XXX_Size() int {
	return xxx_messageInfo_EventGetByIdResponse.Size(m)
}
func (m *EventGetByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventGetByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventGetByIdResponse proto.InternalMessageInfo

func (m *EventGetByIdResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type EventGetAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventGetAllRequest) Reset()         { *m = EventGetAllRequest{} }
func (m *EventGetAllRequest) String() string { return proto.CompactTextString(m) }
func (*EventGetAllRequest) ProtoMessage()    {}
func (*EventGetAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{9}
}

func (m *EventGetAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventGetAllRequest.Unmarshal(m, b)
}
func (m *EventGetAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventGetAllRequest.Marshal(b, m, deterministic)
}
func (m *EventGetAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventGetAllRequest.Merge(m, src)
}
func (m *EventGetAllRequest) XXX_Size() int {
	return xxx_messageInfo_EventGetAllRequest.Size(m)
}
func (m *EventGetAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventGetAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventGetAllRequest proto.InternalMessageInfo

type EventGetAllResponse struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventGetAllResponse) Reset()         { *m = EventGetAllResponse{} }
func (m *EventGetAllResponse) String() string { return proto.CompactTextString(m) }
func (*EventGetAllResponse) ProtoMessage()    {}
func (*EventGetAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{10}
}

func (m *EventGetAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventGetAllResponse.Unmarshal(m, b)
}
func (m *EventGetAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventGetAllResponse.Marshal(b, m, deterministic)
}
func (m *EventGetAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventGetAllResponse.Merge(m, src)
}
func (m *EventGetAllResponse) XXX_Size() int {
	return xxx_messageInfo_EventGetAllResponse.Size(m)
}
func (m *EventGetAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventGetAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventGetAllResponse proto.InternalMessageInfo

func (m *EventGetAllResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "events.Event")
	proto.RegisterType((*EventAddRequest)(nil), "events.EventAddRequest")
	proto.RegisterType((*EventAddResponse)(nil), "events.EventAddResponse")
	proto.RegisterType((*EventEditRequest)(nil), "events.EventEditRequest")
	proto.RegisterType((*EventEditResponse)(nil), "events.EventEditResponse")
	proto.RegisterType((*EventDeleteRequest)(nil), "events.EventDeleteRequest")
	proto.RegisterType((*EventDeleteResponse)(nil), "events.EventDeleteResponse")
	proto.RegisterType((*EventGetByIdRequest)(nil), "events.EventGetByIdRequest")
	proto.RegisterType((*EventGetByIdResponse)(nil), "events.EventGetByIdResponse")
	proto.RegisterType((*EventGetAllRequest)(nil), "events.EventGetAllRequest")
	proto.RegisterType((*EventGetAllResponse)(nil), "events.EventGetAllResponse")
}

func init() {
	proto.RegisterFile("calendar.proto", fileDescriptor_2d17a9d3f0ddf27e)
}

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x8f, 0xd3, 0x40,
	0x0c, 0x25, 0x49, 0x13, 0x09, 0xb7, 0xcb, 0xc7, 0x6c, 0x81, 0x30, 0x8b, 0x44, 0x35, 0xb0, 0xd2,
	0x9e, 0xb2, 0xd2, 0x72, 0x43, 0x54, 0xa8, 0x5d, 0xaa, 0x88, 0x6b, 0xca, 0x89, 0x0b, 0x2a, 0x1d,
	0x53, 0x45, 0x4a, 0x9b, 0xd0, 0xb8, 0x95, 0xf8, 0x75, 0x9c, 0xf8, 0x0d, 0xfc, 0x1d, 0x94, 0x99,
	0x49, 0xd3, 0x69, 0x53, 0x10, 0x70, 0x8a, 0xc6, 0xef, 0xd9, 0xcf, 0xcf, 0x76, 0xa0, 0x8b, 0x5b,
	0x5c, 0x51, 0x54, 0xac, 0x73, 0xca, 0x59, 0xa0, 0x1e, 0x25, 0x7f, 0xbe, 0xc8, 0xf3, 0x45, 0x86,
	0xd7, 0x2a, 0xfa, 0x79, 0xf3, 0xe5, 0x9a, 0xd2, 0x25, 0x96, 0x34, 0x5b, 0x16, 0x9a, 0x28, 0x7e,
	0x38, 0xe0, 0x4f, 0x2a, 0x2e, 0xbb, 0x07, 0x6e, 0x2a, 0x43, 0x67, 0xe0, 0x5c, 0xdd, 0x4d, 0xdc,
	0x54, 0xb2, 0x3e, 0xf8, 0x94, 0x52, 0x86, 0xa1, 0xab, 0x42, 0xfa, 0xc1, 0x86, 0xd0, 0x93, 0x33,
	0xc2, 0x4f, 0x25, 0xcd, 0xd6, 0x84, 0x32, 0xf4, 0x06, 0xce, 0x55, 0xf7, 0x86, 0x47, 0x5a, 0x27,
	0xaa, 0x75, 0xa2, 0x0f, 0xb5, 0x4e, 0xd2, 0xad, 0xf8, 0x53, 0x4d, 0x67, 0x6f, 0xe1, 0x4c, 0xa5,
	0xcf, 0xf3, 0x65, 0x91, 0x21, 0x61, 0xd8, 0xf9, 0x63, 0xbe, 0xd2, 0xbb, 0x35, 0x7c, 0xf6, 0x18,
	0x82, 0x55, 0x4e, 0xe9, 0x1c, 0x43, 0x5f, 0xb5, 0x65, 0x5e, 0xe2, 0xbb, 0x03, 0xf7, 0x95, 0x8f,
	0x91, 0x94, 0x09, 0x7e, 0xdd, 0x60, 0x49, 0x8d, 0x03, 0xe7, 0x77, 0x0e, 0xdc, 0xff, 0x74, 0xe0,
	0xfd, 0xb3, 0x83, 0x8e, 0xe5, 0x40, 0xc0, 0x83, 0xc6, 0x40, 0x59, 0xe4, 0xab, 0x12, 0x0f, 0x77,
	0x22, 0x62, 0xc3, 0x99, 0xc8, 0x94, 0x6a, 0x97, 0x87, 0x7b, 0x7b, 0x01, 0xbe, 0x5a, 0xbe, 0x31,
	0x76, 0x16, 0xe9, 0x53, 0x88, 0x54, 0x62, 0xa2, 0x31, 0x71, 0x0e, 0x0f, 0xf7, 0x0a, 0x69, 0x35,
	0xf1, 0x12, 0x98, 0x0a, 0xbe, 0xc3, 0xaa, 0xd1, 0x13, 0xf5, 0xc5, 0x23, 0x38, 0xb7, 0x58, 0x26,
	0xf9, 0xd2, 0x84, 0x63, 0xa4, 0xf1, 0xb7, 0xf7, 0xf2, 0x54, 0xf6, 0x10, 0xfa, 0x36, 0xcd, 0x38,
	0xbd, 0x04, 0x73, 0xb2, 0xa1, 0x33, 0xf0, 0x8e, 0xdb, 0x36, 0xa0, 0xe8, 0x9b, 0x16, 0x63, 0xa4,
	0x51, 0x96, 0x19, 0x11, 0xf1, 0xa6, 0xd1, 0x56, 0xd1, 0xbf, 0xaa, 0x79, 0xf3, 0xd3, 0x85, 0x9e,
	0x8a, 0x4c, 0x71, 0xbd, 0x4d, 0xe7, 0xc8, 0x5e, 0x83, 0x37, 0x92, 0x92, 0x3d, 0xb1, 0xe8, 0xcd,
	0x5d, 0xf1, 0xf0, 0x18, 0x30, 0x43, 0xb8, 0xc3, 0x86, 0xd0, 0xa9, 0x66, 0xca, 0x6c, 0xce, 0xde,
	0xbe, 0xf8, 0xd3, 0x16, 0x64, 0x97, 0x7e, 0x0b, 0x81, 0x9e, 0x2b, 0xe3, 0x16, 0xcd, 0x5a, 0x09,
	0xbf, 0x68, 0xc5, 0x76, 0x45, 0xc6, 0xe0, 0xc5, 0x48, 0xcc, 0x66, 0xd9, 0x7b, 0xe1, 0xcf, 0xda,
	0xc1, 0xfd, 0x46, 0xf4, 0x34, 0x0f, 0x1a, 0xb1, 0x06, 0xcf, 0x2f, 0x5a, 0xb1, 0xba, 0xc8, 0xb8,
	0xf7, 0x11, 0x14, 0xae, 0xff, 0x89, 0x40, 0x7d, 0x5e, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x94,
	0xdd, 0x0c, 0x5f, 0xa9, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	Add(ctx context.Context, in *EventAddRequest, opts ...grpc.CallOption) (*EventAddResponse, error)
	Edit(ctx context.Context, in *EventEditRequest, opts ...grpc.CallOption) (*EventEditResponse, error)
	Delete(ctx context.Context, in *EventDeleteRequest, opts ...grpc.CallOption) (*EventDeleteResponse, error)
	Get(ctx context.Context, in *EventGetByIdRequest, opts ...grpc.CallOption) (*EventGetByIdResponse, error)
	GetAll(ctx context.Context, in *EventGetAllRequest, opts ...grpc.CallOption) (*EventGetAllResponse, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) Add(ctx context.Context, in *EventAddRequest, opts ...grpc.CallOption) (*EventAddResponse, error) {
	out := new(EventAddResponse)
	err := c.cc.Invoke(ctx, "/events.EventService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Edit(ctx context.Context, in *EventEditRequest, opts ...grpc.CallOption) (*EventEditResponse, error) {
	out := new(EventEditResponse)
	err := c.cc.Invoke(ctx, "/events.EventService/Edit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Delete(ctx context.Context, in *EventDeleteRequest, opts ...grpc.CallOption) (*EventDeleteResponse, error) {
	out := new(EventDeleteResponse)
	err := c.cc.Invoke(ctx, "/events.EventService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Get(ctx context.Context, in *EventGetByIdRequest, opts ...grpc.CallOption) (*EventGetByIdResponse, error) {
	out := new(EventGetByIdResponse)
	err := c.cc.Invoke(ctx, "/events.EventService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetAll(ctx context.Context, in *EventGetAllRequest, opts ...grpc.CallOption) (*EventGetAllResponse, error) {
	out := new(EventGetAllResponse)
	err := c.cc.Invoke(ctx, "/events.EventService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	Add(context.Context, *EventAddRequest) (*EventAddResponse, error)
	Edit(context.Context, *EventEditRequest) (*EventEditResponse, error)
	Delete(context.Context, *EventDeleteRequest) (*EventDeleteResponse, error)
	Get(context.Context, *EventGetByIdRequest) (*EventGetByIdResponse, error)
	GetAll(context.Context, *EventGetAllRequest) (*EventGetAllResponse, error)
}

// UnimplementedEventServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (*UnimplementedEventServiceServer) Add(ctx context.Context, req *EventAddRequest) (*EventAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedEventServiceServer) Edit(ctx context.Context, req *EventEditRequest) (*EventEditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}
func (*UnimplementedEventServiceServer) Delete(ctx context.Context, req *EventDeleteRequest) (*EventDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedEventServiceServer) Get(ctx context.Context, req *EventGetByIdRequest) (*EventGetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedEventServiceServer) GetAll(ctx context.Context, req *EventGetAllRequest) (*EventGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.EventService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Add(ctx, req.(*EventAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventEditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.EventService/Edit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Edit(ctx, req.(*EventEditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.EventService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Delete(ctx, req.(*EventDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventGetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.EventService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Get(ctx, req.(*EventGetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventGetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.EventService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetAll(ctx, req.(*EventGetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "events.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _EventService_Add_Handler,
		},
		{
			MethodName: "Edit",
			Handler:    _EventService_Edit_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EventService_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _EventService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _EventService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calendar.proto",
}
