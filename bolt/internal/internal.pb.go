// Code generated by protoc-gen-gogo.
// source: internal.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal.proto

It has these top-level messages:
	Source
	Dashboard
	DashboardCell
	Template
	TemplateValue
	TemplateQuery
	Server
	Layout
	Cell
	Query
	Range
	AlertRule
	User
*/
package internal

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Source struct {
	ID                 int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name               string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Type               string `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Username           string `protobuf:"bytes,4,opt,name=Username,proto3" json:"Username,omitempty"`
	Password           string `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	URL                string `protobuf:"bytes,6,opt,name=URL,proto3" json:"URL,omitempty"`
	Default            bool   `protobuf:"varint,7,opt,name=Default,proto3" json:"Default,omitempty"`
	Telegraf           string `protobuf:"bytes,8,opt,name=Telegraf,proto3" json:"Telegraf,omitempty"`
	InsecureSkipVerify bool   `protobuf:"varint,9,opt,name=InsecureSkipVerify,proto3" json:"InsecureSkipVerify,omitempty"`
	MetaURL            string `protobuf:"bytes,10,opt,name=MetaURL,proto3" json:"MetaURL,omitempty"`
}

func (m *Source) Reset()                    { *m = Source{} }
func (m *Source) String() string            { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()               {}
func (*Source) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0} }

type Dashboard struct {
	ID        int64            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name      string           `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Cells     []*DashboardCell `protobuf:"bytes,3,rep,name=cells" json:"cells,omitempty"`
	Templates []*Template      `protobuf:"bytes,4,rep,name=templates" json:"templates,omitempty"`
}

func (m *Dashboard) Reset()                    { *m = Dashboard{} }
func (m *Dashboard) String() string            { return proto.CompactTextString(m) }
func (*Dashboard) ProtoMessage()               {}
func (*Dashboard) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{1} }

func (m *Dashboard) GetCells() []*DashboardCell {
	if m != nil {
		return m.Cells
	}
	return nil
}

func (m *Dashboard) GetTemplates() []*Template {
	if m != nil {
		return m.Templates
	}
	return nil
}

type DashboardCell struct {
	X       int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y       int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	W       int32    `protobuf:"varint,3,opt,name=w,proto3" json:"w,omitempty"`
	H       int32    `protobuf:"varint,4,opt,name=h,proto3" json:"h,omitempty"`
	Queries []*Query `protobuf:"bytes,5,rep,name=queries" json:"queries,omitempty"`
	Name    string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Type    string   `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	ID      string   `protobuf:"bytes,8,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (m *DashboardCell) Reset()                    { *m = DashboardCell{} }
func (m *DashboardCell) String() string            { return proto.CompactTextString(m) }
func (*DashboardCell) ProtoMessage()               {}
func (*DashboardCell) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{2} }

func (m *DashboardCell) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type Template struct {
	ID      string           `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TempVar string           `protobuf:"bytes,2,opt,name=temp_var,json=tempVar,proto3" json:"temp_var,omitempty"`
	Values  []*TemplateValue `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
	Type    string           `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Label   string           `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	Query   *TemplateQuery   `protobuf:"bytes,6,opt,name=query" json:"query,omitempty"`
}

func (m *Template) Reset()                    { *m = Template{} }
func (m *Template) String() string            { return proto.CompactTextString(m) }
func (*Template) ProtoMessage()               {}
func (*Template) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{3} }

func (m *Template) GetValues() []*TemplateValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Template) GetQuery() *TemplateQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

type TemplateValue struct {
	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Selected bool   `protobuf:"varint,3,opt,name=selected,proto3" json:"selected,omitempty"`
}

func (m *TemplateValue) Reset()                    { *m = TemplateValue{} }
func (m *TemplateValue) String() string            { return proto.CompactTextString(m) }
func (*TemplateValue) ProtoMessage()               {}
func (*TemplateValue) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{4} }

type TemplateQuery struct {
	Command     string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Db          string `protobuf:"bytes,2,opt,name=db,proto3" json:"db,omitempty"`
	Rp          string `protobuf:"bytes,3,opt,name=rp,proto3" json:"rp,omitempty"`
	Measurement string `protobuf:"bytes,4,opt,name=measurement,proto3" json:"measurement,omitempty"`
	TagKey      string `protobuf:"bytes,5,opt,name=tag_key,json=tagKey,proto3" json:"tag_key,omitempty"`
	FieldKey    string `protobuf:"bytes,6,opt,name=field_key,json=fieldKey,proto3" json:"field_key,omitempty"`
}

func (m *TemplateQuery) Reset()                    { *m = TemplateQuery{} }
func (m *TemplateQuery) String() string            { return proto.CompactTextString(m) }
func (*TemplateQuery) ProtoMessage()               {}
func (*TemplateQuery) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{5} }

type Server struct {
	ID       int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	URL      string `protobuf:"bytes,5,opt,name=URL,proto3" json:"URL,omitempty"`
	SrcID    int64  `protobuf:"varint,6,opt,name=SrcID,proto3" json:"SrcID,omitempty"`
}

func (m *Server) Reset()                    { *m = Server{} }
func (m *Server) String() string            { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()               {}
func (*Server) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{6} }

type Layout struct {
	ID          string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Application string  `protobuf:"bytes,2,opt,name=Application,proto3" json:"Application,omitempty"`
	Measurement string  `protobuf:"bytes,3,opt,name=Measurement,proto3" json:"Measurement,omitempty"`
	Cells       []*Cell `protobuf:"bytes,4,rep,name=Cells" json:"Cells,omitempty"`
	Autoflow    bool    `protobuf:"varint,5,opt,name=Autoflow,proto3" json:"Autoflow,omitempty"`
}

func (m *Layout) Reset()                    { *m = Layout{} }
func (m *Layout) String() string            { return proto.CompactTextString(m) }
func (*Layout) ProtoMessage()               {}
func (*Layout) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{7} }

func (m *Layout) GetCells() []*Cell {
	if m != nil {
		return m.Cells
	}
	return nil
}

type Cell struct {
	X       int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y       int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	W       int32    `protobuf:"varint,3,opt,name=w,proto3" json:"w,omitempty"`
	H       int32    `protobuf:"varint,4,opt,name=h,proto3" json:"h,omitempty"`
	Queries []*Query `protobuf:"bytes,5,rep,name=queries" json:"queries,omitempty"`
	I       string   `protobuf:"bytes,6,opt,name=i,proto3" json:"i,omitempty"`
	Name    string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Yranges []int64  `protobuf:"varint,8,rep,name=yranges" json:"yranges,omitempty"`
	Ylabels []string `protobuf:"bytes,9,rep,name=ylabels" json:"ylabels,omitempty"`
	Type    string   `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *Cell) Reset()                    { *m = Cell{} }
func (m *Cell) String() string            { return proto.CompactTextString(m) }
func (*Cell) ProtoMessage()               {}
func (*Cell) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{8} }

func (m *Cell) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type Query struct {
	Command  string   `protobuf:"bytes,1,opt,name=Command,proto3" json:"Command,omitempty"`
	DB       string   `protobuf:"bytes,2,opt,name=DB,proto3" json:"DB,omitempty"`
	RP       string   `protobuf:"bytes,3,opt,name=RP,proto3" json:"RP,omitempty"`
	GroupBys []string `protobuf:"bytes,4,rep,name=GroupBys" json:"GroupBys,omitempty"`
	Wheres   []string `protobuf:"bytes,5,rep,name=Wheres" json:"Wheres,omitempty"`
	Label    string   `protobuf:"bytes,6,opt,name=Label,proto3" json:"Label,omitempty"`
	Range    *Range   `protobuf:"bytes,7,opt,name=Range" json:"Range,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{9} }

func (m *Query) GetRange() *Range {
	if m != nil {
		return m.Range
	}
	return nil
}

type Range struct {
	Upper int64 `protobuf:"varint,1,opt,name=Upper,proto3" json:"Upper,omitempty"`
	Lower int64 `protobuf:"varint,2,opt,name=Lower,proto3" json:"Lower,omitempty"`
}

func (m *Range) Reset()                    { *m = Range{} }
func (m *Range) String() string            { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()               {}
func (*Range) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{10} }

type AlertRule struct {
	ID     string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	JSON   string `protobuf:"bytes,2,opt,name=JSON,proto3" json:"JSON,omitempty"`
	SrcID  int64  `protobuf:"varint,3,opt,name=SrcID,proto3" json:"SrcID,omitempty"`
	KapaID int64  `protobuf:"varint,4,opt,name=KapaID,proto3" json:"KapaID,omitempty"`
}

func (m *AlertRule) Reset()                    { *m = AlertRule{} }
func (m *AlertRule) String() string            { return proto.CompactTextString(m) }
func (*AlertRule) ProtoMessage()               {}
func (*AlertRule) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{11} }

type User struct {
	ID   uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{12} }

func init() {
	proto.RegisterType((*Source)(nil), "internal.Source")
	proto.RegisterType((*Dashboard)(nil), "internal.Dashboard")
	proto.RegisterType((*DashboardCell)(nil), "internal.DashboardCell")
	proto.RegisterType((*Template)(nil), "internal.Template")
	proto.RegisterType((*TemplateValue)(nil), "internal.TemplateValue")
	proto.RegisterType((*TemplateQuery)(nil), "internal.TemplateQuery")
	proto.RegisterType((*Server)(nil), "internal.Server")
	proto.RegisterType((*Layout)(nil), "internal.Layout")
	proto.RegisterType((*Cell)(nil), "internal.Cell")
	proto.RegisterType((*Query)(nil), "internal.Query")
	proto.RegisterType((*Range)(nil), "internal.Range")
	proto.RegisterType((*AlertRule)(nil), "internal.AlertRule")
	proto.RegisterType((*User)(nil), "internal.User")
}

func init() { proto.RegisterFile("internal.proto", fileDescriptorInternal) }

var fileDescriptorInternal = []byte{
	// 848 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xd6, 0xc4, 0x76, 0x62, 0x9f, 0xdd, 0x2d, 0x68, 0xb4, 0x62, 0x0d, 0xdc, 0x44, 0x16, 0x48,
	0x01, 0x89, 0x82, 0xd8, 0x27, 0x68, 0x6b, 0x09, 0x85, 0x76, 0x97, 0x32, 0x69, 0xcb, 0x15, 0x5a,
	0x4d, 0x92, 0x93, 0xd6, 0xda, 0x49, 0x6c, 0xc6, 0x76, 0xb3, 0x7e, 0x05, 0xc4, 0x05, 0x4f, 0x80,
	0xc4, 0x2d, 0x97, 0xbc, 0x00, 0x0f, 0xc1, 0x0b, 0xa1, 0x33, 0x33, 0xfe, 0x89, 0xb6, 0xa0, 0xbd,
	0xe2, 0x6e, 0xbe, 0x73, 0x26, 0xdf, 0x9c, 0x9f, 0xef, 0x73, 0xe0, 0x28, 0xdb, 0x55, 0xa8, 0x77,
	0x52, 0x1d, 0x17, 0x3a, 0xaf, 0x72, 0x1e, 0xb6, 0x38, 0xf9, 0x79, 0x04, 0xe3, 0x45, 0x5e, 0xeb,
	0x15, 0xf2, 0x23, 0x18, 0xcd, 0xd3, 0x98, 0x4d, 0xd9, 0xcc, 0x13, 0xa3, 0x79, 0xca, 0x39, 0xf8,
	0x2f, 0xe5, 0x16, 0xe3, 0xd1, 0x94, 0xcd, 0x22, 0x61, 0xce, 0x14, 0xbb, 0x6a, 0x0a, 0x8c, 0x3d,
	0x1b, 0xa3, 0x33, 0xff, 0x08, 0xc2, 0xeb, 0x92, 0xd8, 0xb6, 0x18, 0xfb, 0x26, 0xde, 0x61, 0xca,
	0x5d, 0xca, 0xb2, 0xdc, 0xe7, 0x7a, 0x1d, 0x07, 0x36, 0xd7, 0x62, 0xfe, 0x3e, 0x78, 0xd7, 0xe2,
	0x22, 0x1e, 0x9b, 0x30, 0x1d, 0x79, 0x0c, 0x93, 0x14, 0x37, 0xb2, 0x56, 0x55, 0x3c, 0x99, 0xb2,
	0x59, 0x28, 0x5a, 0x48, 0x3c, 0x57, 0xa8, 0xf0, 0x56, 0xcb, 0x4d, 0x1c, 0x5a, 0x9e, 0x16, 0xf3,
	0x63, 0xe0, 0xf3, 0x5d, 0x89, 0xab, 0x5a, 0xe3, 0xe2, 0x75, 0x56, 0xdc, 0xa0, 0xce, 0x36, 0x4d,
	0x1c, 0x19, 0x82, 0x07, 0x32, 0xf4, 0xca, 0x0b, 0xac, 0x24, 0xbd, 0x0d, 0x86, 0xaa, 0x85, 0xc9,
	0xaf, 0x0c, 0xa2, 0x54, 0x96, 0x77, 0xcb, 0x5c, 0xea, 0xf5, 0x3b, 0xcd, 0xe3, 0x0b, 0x08, 0x56,
	0xa8, 0x54, 0x19, 0x7b, 0x53, 0x6f, 0xf6, 0xe8, 0xeb, 0x67, 0xc7, 0xdd, 0xa0, 0x3b, 0x9e, 0x33,
	0x54, 0x4a, 0xd8, 0x5b, 0xfc, 0x2b, 0x88, 0x2a, 0xdc, 0x16, 0x4a, 0x56, 0x58, 0xc6, 0xbe, 0xf9,
	0x09, 0xef, 0x7f, 0x72, 0xe5, 0x52, 0xa2, 0xbf, 0x94, 0xfc, 0xc1, 0xe0, 0xc9, 0x01, 0x15, 0x7f,
	0x0c, 0xec, 0x8d, 0xa9, 0x2a, 0x10, 0xec, 0x0d, 0xa1, 0xc6, 0x54, 0x14, 0x08, 0xd6, 0x10, 0xda,
	0x9b, 0xdd, 0x04, 0x82, 0xed, 0x09, 0xdd, 0x99, 0x8d, 0x04, 0x82, 0xdd, 0xf1, 0xcf, 0x60, 0xf2,
	0x53, 0x8d, 0x3a, 0xc3, 0x32, 0x0e, 0xcc, 0xcb, 0xef, 0xf5, 0x2f, 0x7f, 0x5f, 0xa3, 0x6e, 0x44,
	0x9b, 0xa7, 0x4e, 0xcd, 0x36, 0xed, 0x6a, 0xcc, 0x99, 0x62, 0x15, 0x6d, 0x7e, 0x62, 0x63, 0x74,
	0x76, 0x13, 0xb2, 0xfb, 0x18, 0xcd, 0xd3, 0xe4, 0x2f, 0x46, 0x6b, 0xb2, 0xa5, 0x0f, 0xc6, 0x67,
	0x92, 0xfc, 0x43, 0x08, 0xa9, 0xad, 0x57, 0xf7, 0x52, 0xbb, 0x11, 0x4e, 0x08, 0xdf, 0x48, 0xcd,
	0xbf, 0x84, 0xf1, 0xbd, 0x54, 0x35, 0x3e, 0x30, 0xc6, 0x96, 0xee, 0x86, 0xf2, 0xc2, 0x5d, 0xeb,
	0x8a, 0xf1, 0x07, 0xc5, 0x3c, 0x85, 0x40, 0xc9, 0x25, 0x2a, 0xa7, 0x33, 0x0b, 0x68, 0x41, 0xd4,
	0x55, 0x63, 0x7a, 0x79, 0x90, 0xd9, 0xf6, 0x6e, 0x6f, 0x25, 0xd7, 0xf0, 0xe4, 0xe0, 0xc5, 0xee,
	0x25, 0x76, 0xf8, 0x92, 0xa9, 0xc3, 0xb5, 0x61, 0x01, 0x49, 0xb4, 0x44, 0x85, 0xab, 0x0a, 0xd7,
	0x66, 0x05, 0xa1, 0xe8, 0x70, 0xf2, 0x3b, 0xeb, 0x79, 0xcd, 0x7b, 0x24, 0xc2, 0x55, 0xbe, 0xdd,
	0xca, 0xdd, 0xda, 0x51, 0xb7, 0x90, 0xe6, 0xb6, 0x5e, 0x3a, 0xea, 0xd1, 0x7a, 0x49, 0x58, 0x17,
	0xce, 0x70, 0x23, 0x5d, 0xf0, 0x29, 0x3c, 0xda, 0xa2, 0x2c, 0x6b, 0x8d, 0x5b, 0xdc, 0x55, 0x6e,
	0x04, 0xc3, 0x10, 0x7f, 0x06, 0x93, 0x4a, 0xde, 0xbe, 0x7a, 0x8d, 0x8d, 0x9b, 0xc5, 0xb8, 0x92,
	0xb7, 0xe7, 0xd8, 0xf0, 0x8f, 0x21, 0xda, 0x64, 0xa8, 0xd6, 0x26, 0x65, 0x97, 0x1b, 0x9a, 0xc0,
	0x39, 0x36, 0xc9, 0x2f, 0x0c, 0xc6, 0x0b, 0xd4, 0xf7, 0xa8, 0xdf, 0x49, 0xf9, 0x43, 0xd7, 0x7b,
	0xff, 0xe1, 0x7a, 0xff, 0x61, 0xd7, 0x07, 0xbd, 0xeb, 0x9f, 0x42, 0xb0, 0xd0, 0xab, 0x79, 0x6a,
	0x2a, 0xf2, 0x84, 0x05, 0xc9, 0x6f, 0x0c, 0xc6, 0x17, 0xb2, 0xc9, 0xeb, 0xea, 0x2d, 0x25, 0x4d,
	0xe1, 0xd1, 0x49, 0x51, 0xa8, 0x6c, 0x25, 0xab, 0x2c, 0xdf, 0xb9, 0xaa, 0x86, 0x21, 0xba, 0xf1,
	0x62, 0x30, 0x23, 0x5b, 0xdf, 0x30, 0xc4, 0x3f, 0x81, 0xe0, 0xcc, 0x18, 0xd7, 0xba, 0xf0, 0xa8,
	0xd7, 0x85, 0xf5, 0xab, 0x49, 0x52, 0x23, 0x27, 0x75, 0x95, 0x6f, 0x54, 0xbe, 0x37, 0x15, 0x87,
	0xa2, 0xc3, 0xc9, 0xdf, 0x0c, 0xfc, 0xff, 0xcb, 0x90, 0x8f, 0x81, 0x65, 0x6e, 0x61, 0x2c, 0xeb,
	0xec, 0x39, 0x19, 0xd8, 0x33, 0x86, 0x49, 0xa3, 0xe5, 0xee, 0x16, 0xcb, 0x38, 0x9c, 0x7a, 0x33,
	0x4f, 0xb4, 0xd0, 0x64, 0x8c, 0x17, 0xca, 0x38, 0x9a, 0x7a, 0xa4, 0x34, 0x07, 0x3b, 0x6d, 0x43,
	0xaf, 0xed, 0xe4, 0x4f, 0x06, 0x41, 0xa7, 0xd0, 0xb3, 0x43, 0x85, 0x9e, 0xf5, 0x0a, 0x4d, 0x4f,
	0x5b, 0x85, 0xa6, 0xa7, 0x84, 0xc5, 0x65, 0xab, 0x50, 0x71, 0x49, 0x53, 0xfb, 0x46, 0xe7, 0x75,
	0x71, 0xda, 0xd8, 0xf1, 0x46, 0xa2, 0xc3, 0xfc, 0x03, 0x18, 0xff, 0x70, 0x87, 0xda, 0xf5, 0x1c,
	0x09, 0x87, 0x48, 0x04, 0x17, 0xc6, 0xbd, 0xb6, 0x4b, 0x0b, 0xf8, 0xa7, 0x10, 0x08, 0xea, 0xc2,
	0xb4, 0x7a, 0x30, 0x20, 0x13, 0x16, 0x36, 0x9b, 0x3c, 0x77, 0xd7, 0x88, 0xe5, 0xba, 0x28, 0x50,
	0x3b, 0xed, 0x5a, 0x60, 0xb8, 0xf3, 0x3d, 0xda, 0xcf, 0x8e, 0x27, 0x2c, 0x48, 0x7e, 0x84, 0xe8,
	0x44, 0xa1, 0xae, 0x44, 0xad, 0xde, 0xfe, 0x58, 0x71, 0xf0, 0xbf, 0x5d, 0x7c, 0xf7, 0xb2, 0x55,
	0x3c, 0x9d, 0x7b, 0x9d, 0x7a, 0x03, 0x9d, 0x52, 0x43, 0xe7, 0xb2, 0x90, 0xf3, 0xd4, 0x2c, 0xd6,
	0x13, 0x0e, 0x25, 0x9f, 0x83, 0x4f, 0x7e, 0x18, 0x30, 0xfb, 0xff, 0xe6, 0xa5, 0xe5, 0xd8, 0xfc,
	0x2b, 0x3f, 0xff, 0x27, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x04, 0x69, 0xb6, 0xa7, 0x07, 0x00, 0x00,
}
