// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/dictionaries/dictionaries.proto

package go_micro_srv_systemSetup

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 状态结构
type Status struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_881fdb2702eabd49, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Status) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// 列表页返回页面信息
type ResponsePage struct {
	Size                 int32    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Total                int32    `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	SumPage              int32    `protobuf:"varint,4,opt,name=sumPage,proto3" json:"sumPage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponsePage) Reset()         { *m = ResponsePage{} }
func (m *ResponsePage) String() string { return proto.CompactTextString(m) }
func (*ResponsePage) ProtoMessage()    {}
func (*ResponsePage) Descriptor() ([]byte, []int) {
	return fileDescriptor_881fdb2702eabd49, []int{1}
}

func (m *ResponsePage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponsePage.Unmarshal(m, b)
}
func (m *ResponsePage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponsePage.Marshal(b, m, deterministic)
}
func (m *ResponsePage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponsePage.Merge(m, src)
}
func (m *ResponsePage) XXX_Size() int {
	return xxx_messageInfo_ResponsePage.Size(m)
}
func (m *ResponsePage) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponsePage.DiscardUnknown(m)
}

var xxx_messageInfo_ResponsePage proto.InternalMessageInfo

func (m *ResponsePage) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ResponsePage) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ResponsePage) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ResponsePage) GetSumPage() int32 {
	if m != nil {
		return m.SumPage
	}
	return 0
}

// 请求列表页需要相关信息
type RequestPage struct {
	Size                 int32    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestPage) Reset()         { *m = RequestPage{} }
func (m *RequestPage) String() string { return proto.CompactTextString(m) }
func (*RequestPage) ProtoMessage()    {}
func (*RequestPage) Descriptor() ([]byte, []int) {
	return fileDescriptor_881fdb2702eabd49, []int{2}
}

func (m *RequestPage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestPage.Unmarshal(m, b)
}
func (m *RequestPage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestPage.Marshal(b, m, deterministic)
}
func (m *RequestPage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestPage.Merge(m, src)
}
func (m *RequestPage) XXX_Size() int {
	return xxx_messageInfo_RequestPage.Size(m)
}
func (m *RequestPage) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestPage.DiscardUnknown(m)
}

var xxx_messageInfo_RequestPage proto.InternalMessageInfo

func (m *RequestPage) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *RequestPage) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

// 请求字典列表需要的参数
type ListRequest struct {
	Page                 *RequestPage `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_881fdb2702eabd49, []int{3}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetPage() *RequestPage {
	if m != nil {
		return m.Page
	}
	return nil
}

// 请求字典列表的参数
type ListResponse struct {
	// 页面信息
	Page *ResponsePage `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	// 状态
	Status *Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// 字典数据(数组)
	Data                 []*DictionaryType `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_881fdb2702eabd49, []int{4}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetPage() *ResponsePage {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *ListResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListResponse) GetData() []*DictionaryType {
	if m != nil {
		return m.Data
	}
	return nil
}

// 删除字典请求参数
type DelRequest struct {
	Ids                  string   `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelRequest) Reset()         { *m = DelRequest{} }
func (m *DelRequest) String() string { return proto.CompactTextString(m) }
func (*DelRequest) ProtoMessage()    {}
func (*DelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_881fdb2702eabd49, []int{5}
}

func (m *DelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelRequest.Unmarshal(m, b)
}
func (m *DelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelRequest.Marshal(b, m, deterministic)
}
func (m *DelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelRequest.Merge(m, src)
}
func (m *DelRequest) XXX_Size() int {
	return xxx_messageInfo_DelRequest.Size(m)
}
func (m *DelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DelRequest proto.InternalMessageInfo

func (m *DelRequest) GetIds() string {
	if m != nil {
		return m.Ids
	}
	return ""
}

// 字典操作结果
type EditResponse struct {
	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// 状态
	Status               *Status  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditResponse) Reset()         { *m = EditResponse{} }
func (m *EditResponse) String() string { return proto.CompactTextString(m) }
func (*EditResponse) ProtoMessage()    {}
func (*EditResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_881fdb2702eabd49, []int{6}
}

func (m *EditResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditResponse.Unmarshal(m, b)
}
func (m *EditResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditResponse.Marshal(b, m, deterministic)
}
func (m *EditResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditResponse.Merge(m, src)
}
func (m *EditResponse) XXX_Size() int {
	return xxx_messageInfo_EditResponse.Size(m)
}
func (m *EditResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EditResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EditResponse proto.InternalMessageInfo

func (m *EditResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *EditResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// 字典类型的结构
type DictionaryType struct {
	Id                   int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ClassId              int32         `protobuf:"varint,2,opt,name=classId,proto3" json:"classId,omitempty"`
	Title                string        `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Key                  string        `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	IsUse                bool          `protobuf:"varint,5,opt,name=isUse,proto3" json:"isUse,omitempty"`
	Describe             string        `protobuf:"bytes,6,opt,name=describe,proto3" json:"describe,omitempty"`
	Dictionaries         []*Dictionary `protobuf:"bytes,7,rep,name=dictionaries,proto3" json:"dictionaries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DictionaryType) Reset()         { *m = DictionaryType{} }
func (m *DictionaryType) String() string { return proto.CompactTextString(m) }
func (*DictionaryType) ProtoMessage()    {}
func (*DictionaryType) Descriptor() ([]byte, []int) {
	return fileDescriptor_881fdb2702eabd49, []int{7}
}

func (m *DictionaryType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DictionaryType.Unmarshal(m, b)
}
func (m *DictionaryType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DictionaryType.Marshal(b, m, deterministic)
}
func (m *DictionaryType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DictionaryType.Merge(m, src)
}
func (m *DictionaryType) XXX_Size() int {
	return xxx_messageInfo_DictionaryType.Size(m)
}
func (m *DictionaryType) XXX_DiscardUnknown() {
	xxx_messageInfo_DictionaryType.DiscardUnknown(m)
}

var xxx_messageInfo_DictionaryType proto.InternalMessageInfo

func (m *DictionaryType) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DictionaryType) GetClassId() int32 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *DictionaryType) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *DictionaryType) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DictionaryType) GetIsUse() bool {
	if m != nil {
		return m.IsUse
	}
	return false
}

func (m *DictionaryType) GetDescribe() string {
	if m != nil {
		return m.Describe
	}
	return ""
}

func (m *DictionaryType) GetDictionaries() []*Dictionary {
	if m != nil {
		return m.Dictionaries
	}
	return nil
}

// 添加字典类型的结构
type AddDictionaryType struct {
	ClassId              int32            `protobuf:"varint,2,opt,name=classId,proto3" json:"classId,omitempty"`
	Title                string           `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Key                  string           `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	IsUse                bool             `protobuf:"varint,5,opt,name=isUse,proto3" json:"isUse,omitempty"`
	Describe             string           `protobuf:"bytes,6,opt,name=describe,proto3" json:"describe,omitempty"`
	Dictionaries         []*AddDictionary `protobuf:"bytes,7,rep,name=dictionaries,proto3" json:"dictionaries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AddDictionaryType) Reset()         { *m = AddDictionaryType{} }
func (m *AddDictionaryType) String() string { return proto.CompactTextString(m) }
func (*AddDictionaryType) ProtoMessage()    {}
func (*AddDictionaryType) Descriptor() ([]byte, []int) {
	return fileDescriptor_881fdb2702eabd49, []int{8}
}

func (m *AddDictionaryType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddDictionaryType.Unmarshal(m, b)
}
func (m *AddDictionaryType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddDictionaryType.Marshal(b, m, deterministic)
}
func (m *AddDictionaryType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddDictionaryType.Merge(m, src)
}
func (m *AddDictionaryType) XXX_Size() int {
	return xxx_messageInfo_AddDictionaryType.Size(m)
}
func (m *AddDictionaryType) XXX_DiscardUnknown() {
	xxx_messageInfo_AddDictionaryType.DiscardUnknown(m)
}

var xxx_messageInfo_AddDictionaryType proto.InternalMessageInfo

func (m *AddDictionaryType) GetClassId() int32 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *AddDictionaryType) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AddDictionaryType) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AddDictionaryType) GetIsUse() bool {
	if m != nil {
		return m.IsUse
	}
	return false
}

func (m *AddDictionaryType) GetDescribe() string {
	if m != nil {
		return m.Describe
	}
	return ""
}

func (m *AddDictionaryType) GetDictionaries() []*AddDictionary {
	if m != nil {
		return m.Dictionaries
	}
	return nil
}

// 字典结构
type Dictionary struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TypeId               int32    `protobuf:"varint,2,opt,name=typeId,proto3" json:"typeId,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Order                int32    `protobuf:"varint,5,opt,name=order,proto3" json:"order,omitempty"`
	Describe             string   `protobuf:"bytes,6,opt,name=describe,proto3" json:"describe,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dictionary) Reset()         { *m = Dictionary{} }
func (m *Dictionary) String() string { return proto.CompactTextString(m) }
func (*Dictionary) ProtoMessage()    {}
func (*Dictionary) Descriptor() ([]byte, []int) {
	return fileDescriptor_881fdb2702eabd49, []int{9}
}

func (m *Dictionary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dictionary.Unmarshal(m, b)
}
func (m *Dictionary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dictionary.Marshal(b, m, deterministic)
}
func (m *Dictionary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dictionary.Merge(m, src)
}
func (m *Dictionary) XXX_Size() int {
	return xxx_messageInfo_Dictionary.Size(m)
}
func (m *Dictionary) XXX_DiscardUnknown() {
	xxx_messageInfo_Dictionary.DiscardUnknown(m)
}

var xxx_messageInfo_Dictionary proto.InternalMessageInfo

func (m *Dictionary) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Dictionary) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *Dictionary) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Dictionary) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Dictionary) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *Dictionary) GetDescribe() string {
	if m != nil {
		return m.Describe
	}
	return ""
}

// 添加字典结构
type AddDictionary struct {
	TypeId               int32    `protobuf:"varint,2,opt,name=typeId,proto3" json:"typeId,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Order                int32    `protobuf:"varint,5,opt,name=order,proto3" json:"order,omitempty"`
	Describe             string   `protobuf:"bytes,6,opt,name=describe,proto3" json:"describe,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddDictionary) Reset()         { *m = AddDictionary{} }
func (m *AddDictionary) String() string { return proto.CompactTextString(m) }
func (*AddDictionary) ProtoMessage()    {}
func (*AddDictionary) Descriptor() ([]byte, []int) {
	return fileDescriptor_881fdb2702eabd49, []int{10}
}

func (m *AddDictionary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddDictionary.Unmarshal(m, b)
}
func (m *AddDictionary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddDictionary.Marshal(b, m, deterministic)
}
func (m *AddDictionary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddDictionary.Merge(m, src)
}
func (m *AddDictionary) XXX_Size() int {
	return xxx_messageInfo_AddDictionary.Size(m)
}
func (m *AddDictionary) XXX_DiscardUnknown() {
	xxx_messageInfo_AddDictionary.DiscardUnknown(m)
}

var xxx_messageInfo_AddDictionary proto.InternalMessageInfo

func (m *AddDictionary) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *AddDictionary) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AddDictionary) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *AddDictionary) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *AddDictionary) GetDescribe() string {
	if m != nil {
		return m.Describe
	}
	return ""
}

func init() {
	proto.RegisterType((*Status)(nil), "go.micro.srv.systemSetup.Status")
	proto.RegisterType((*ResponsePage)(nil), "go.micro.srv.systemSetup.ResponsePage")
	proto.RegisterType((*RequestPage)(nil), "go.micro.srv.systemSetup.RequestPage")
	proto.RegisterType((*ListRequest)(nil), "go.micro.srv.systemSetup.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.srv.systemSetup.ListResponse")
	proto.RegisterType((*DelRequest)(nil), "go.micro.srv.systemSetup.DelRequest")
	proto.RegisterType((*EditResponse)(nil), "go.micro.srv.systemSetup.EditResponse")
	proto.RegisterType((*DictionaryType)(nil), "go.micro.srv.systemSetup.DictionaryType")
	proto.RegisterType((*AddDictionaryType)(nil), "go.micro.srv.systemSetup.AddDictionaryType")
	proto.RegisterType((*Dictionary)(nil), "go.micro.srv.systemSetup.Dictionary")
	proto.RegisterType((*AddDictionary)(nil), "go.micro.srv.systemSetup.AddDictionary")
}

func init() {
	proto.RegisterFile("proto/dictionaries/dictionaries.proto", fileDescriptor_881fdb2702eabd49)
}

var fileDescriptor_881fdb2702eabd49 = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0xe3, 0xd8, 0x6d, 0x27, 0xa1, 0x2a, 0x2b, 0x84, 0xac, 0x1e, 0x50, 0x64, 0xd1, 0x12,
	0x09, 0xc9, 0x48, 0x41, 0x48, 0x80, 0xb8, 0x54, 0x0a, 0x52, 0x10, 0x1c, 0xd0, 0x16, 0x04, 0x42,
	0xbd, 0xb8, 0xde, 0x21, 0x5a, 0xe1, 0xd4, 0xc6, 0xb3, 0xa9, 0x14, 0x6e, 0xbc, 0x00, 0x0f, 0xc5,
	0x33, 0x70, 0xe6, 0x45, 0xb8, 0xa0, 0xfd, 0x71, 0x12, 0x17, 0x12, 0x52, 0x0e, 0x70, 0x9b, 0x6f,
	0x3d, 0xf3, 0xed, 0x7c, 0x33, 0xb3, 0x63, 0x38, 0x2c, 0xab, 0x42, 0x15, 0xf7, 0x84, 0xcc, 0x94,
	0x2c, 0xce, 0xd3, 0x4a, 0x22, 0x35, 0x40, 0x62, 0xbe, 0xb3, 0x68, 0x5c, 0x24, 0x13, 0x99, 0x55,
	0x45, 0x42, 0xd5, 0x45, 0x42, 0x33, 0x52, 0x38, 0x39, 0x41, 0x35, 0x2d, 0xe3, 0x11, 0x84, 0x27,
	0x2a, 0x55, 0x53, 0x62, 0x0c, 0xda, 0x59, 0x21, 0x30, 0xf2, 0x7a, 0x5e, 0x3f, 0xe0, 0xc6, 0x66,
	0x11, 0x6c, 0xd3, 0x34, 0xcb, 0x90, 0x28, 0x6a, 0xf5, 0xbc, 0xfe, 0x0e, 0xaf, 0x21, 0xdb, 0x07,
	0x7f, 0x42, 0xe3, 0xc8, 0xef, 0x79, 0xfd, 0x5d, 0xae, 0xcd, 0xf8, 0x3d, 0x74, 0x39, 0x52, 0x59,
	0x9c, 0x13, 0xbe, 0x4c, 0xc7, 0xa8, 0xf9, 0x48, 0x7e, 0x9a, 0xf3, 0x69, 0x5b, 0x9f, 0x95, 0xe9,
	0x18, 0x0d, 0x59, 0xc0, 0x8d, 0xcd, 0x6e, 0x40, 0xa0, 0x0a, 0x95, 0xe6, 0x86, 0x2b, 0xe0, 0x16,
	0xd8, 0x9b, 0x27, 0x9a, 0x28, 0x6a, 0x9b, 0xf3, 0x1a, 0xc6, 0x0f, 0xa0, 0xc3, 0xf1, 0xe3, 0x14,
	0x49, 0x5d, 0xe5, 0x9a, 0x78, 0x04, 0x9d, 0x17, 0x92, 0x94, 0x0b, 0x65, 0x8f, 0x9c, 0x8b, 0x0e,
	0xeb, 0x0c, 0x0e, 0x93, 0x55, 0x05, 0x4a, 0x96, 0xee, 0x72, 0x4c, 0x5f, 0x3d, 0xe8, 0x5a, 0x2a,
	0xab, 0x96, 0x3d, 0x6e, 0x70, 0x1d, 0xad, 0xe3, 0x5a, 0xd4, 0xc7, 0xa9, 0x7f, 0x08, 0x21, 0x99,
	0xfa, 0x9b, 0x64, 0x3b, 0x83, 0xde, 0xea, 0x68, 0xdb, 0x27, 0xee, 0xfc, 0xd9, 0x13, 0x68, 0x8b,
	0x54, 0xa5, 0x91, 0xdf, 0xf3, 0xfb, 0x9d, 0x41, 0x7f, 0x75, 0xdc, 0xb0, 0x9e, 0x87, 0xd9, 0xab,
	0x59, 0x89, 0xdc, 0x44, 0xc5, 0xb7, 0x00, 0x86, 0x98, 0xd7, 0xd5, 0xd8, 0x07, 0x5f, 0x0a, 0x32,
	0x02, 0x76, 0xb9, 0x36, 0xe3, 0x53, 0xe8, 0x3e, 0x15, 0x72, 0xa1, 0x91, 0xb9, 0xdb, 0xac, 0x8b,
	0xb1, 0xff, 0x3e, 0xf7, 0xf8, 0xbb, 0x07, 0x7b, 0xcd, 0xb4, 0xd8, 0x1e, 0xb4, 0xa4, 0x70, 0x5d,
	0x6c, 0x49, 0xa1, 0x07, 0x20, 0xcb, 0x53, 0xa2, 0x67, 0xc2, 0xb5, 0xb1, 0x86, 0x66, 0x60, 0xa4,
	0xca, 0xd1, 0x0d, 0x9f, 0x05, 0x5a, 0xc2, 0x07, 0x9c, 0x99, 0x61, 0xd9, 0xe5, 0xda, 0xd4, 0x7e,
	0x92, 0x5e, 0x13, 0x46, 0x81, 0x19, 0x5d, 0x0b, 0xd8, 0x01, 0xec, 0x08, 0xa4, 0xac, 0x92, 0x67,
	0x18, 0x85, 0xc6, 0x79, 0x8e, 0xd9, 0x08, 0xba, 0xcb, 0x8f, 0x27, 0xda, 0x36, 0xa5, 0xbd, 0xbd,
	0x49, 0x69, 0x79, 0x23, 0x32, 0xfe, 0xe6, 0xc1, 0xf5, 0x63, 0x21, 0x2e, 0x69, 0xfc, 0x1f, 0x9a,
	0x9e, 0xff, 0x56, 0xd3, 0x9d, 0xd5, 0x9a, 0x1a, 0x69, 0x5f, 0x92, 0xf5, 0xc5, 0x03, 0x58, 0x7c,
	0xfc, 0xa5, 0x67, 0x37, 0x21, 0x54, 0xb3, 0x12, 0xe7, 0xf2, 0x1c, 0xaa, 0x75, 0xf8, 0x0d, 0x1d,
	0x17, 0x69, 0x3e, 0x45, 0xa7, 0xcd, 0x02, 0x7d, 0x5a, 0x54, 0x02, 0x2b, 0xa3, 0x2e, 0xe0, 0x16,
	0xac, 0x53, 0x17, 0x7f, 0xf6, 0xe0, 0x5a, 0x23, 0xe1, 0x7f, 0x9f, 0xc3, 0xe0, 0x47, 0x0b, 0xba,
	0xc3, 0xa5, 0x2a, 0xb1, 0x37, 0xd0, 0xd6, 0xfb, 0x81, 0xad, 0xd9, 0x2a, 0x4b, 0xab, 0xe8, 0xe0,
	0xe8, 0x4f, 0x6e, 0xf6, 0x09, 0xc6, 0x5b, 0xec, 0x14, 0xfc, 0x63, 0x21, 0xd8, 0xdd, 0x0d, 0x9b,
	0xa7, 0x67, 0x6e, 0x1d, 0xfb, 0xf2, 0x03, 0x8f, 0xb7, 0xd8, 0x3b, 0x68, 0xeb, 0x13, 0xb6, 0xf1,
	0x2a, 0xb9, 0x02, 0xf7, 0x5b, 0x08, 0x87, 0x98, 0xa3, 0x42, 0xb6, 0xee, 0x35, 0xcd, 0x17, 0xd2,
	0xe6, 0xcc, 0x67, 0xa1, 0xf9, 0xc3, 0xdd, 0xff, 0x19, 0x00, 0x00, 0xff, 0xff, 0x21, 0x6d, 0x30,
	0x86, 0x0a, 0x07, 0x00, 0x00,
}