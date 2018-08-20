// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notify.proto

package notify

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_notify_279df0951896b475, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Request struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_notify_279df0951896b475, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ForUserRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForUserRequest) Reset()         { *m = ForUserRequest{} }
func (m *ForUserRequest) String() string { return proto.CompactTextString(m) }
func (*ForUserRequest) ProtoMessage()    {}
func (*ForUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_notify_279df0951896b475, []int{2}
}
func (m *ForUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForUserRequest.Unmarshal(m, b)
}
func (m *ForUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForUserRequest.Marshal(b, m, deterministic)
}
func (dst *ForUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForUserRequest.Merge(dst, src)
}
func (m *ForUserRequest) XXX_Size() int {
	return xxx_messageInfo_ForUserRequest.Size(m)
}
func (m *ForUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ForUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ForUserRequest proto.InternalMessageInfo

func (m *ForUserRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ServiceRequest struct {
	ServiceId            string   `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceRequest) Reset()         { *m = ServiceRequest{} }
func (m *ServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ServiceRequest) ProtoMessage()    {}
func (*ServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_notify_279df0951896b475, []int{3}
}
func (m *ServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceRequest.Unmarshal(m, b)
}
func (m *ServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceRequest.Marshal(b, m, deterministic)
}
func (dst *ServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceRequest.Merge(dst, src)
}
func (m *ServiceRequest) XXX_Size() int {
	return xxx_messageInfo_ServiceRequest.Size(m)
}
func (m *ServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceRequest proto.InternalMessageInfo

func (m *ServiceRequest) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

type TemplatesRequest struct {
	ServiceId            string   `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemplatesRequest) Reset()         { *m = TemplatesRequest{} }
func (m *TemplatesRequest) String() string { return proto.CompactTextString(m) }
func (*TemplatesRequest) ProtoMessage()    {}
func (*TemplatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_notify_279df0951896b475, []int{4}
}
func (m *TemplatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplatesRequest.Unmarshal(m, b)
}
func (m *TemplatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplatesRequest.Marshal(b, m, deterministic)
}
func (dst *TemplatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplatesRequest.Merge(dst, src)
}
func (m *TemplatesRequest) XXX_Size() int {
	return xxx_messageInfo_TemplatesRequest.Size(m)
}
func (m *TemplatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TemplatesRequest proto.InternalMessageInfo

func (m *TemplatesRequest) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

type CreateTemplateRequest struct {
	ServiceId            string   `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	CreatedBy            string   `protobuf:"bytes,2,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Subject              string   `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTemplateRequest) Reset()         { *m = CreateTemplateRequest{} }
func (m *CreateTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTemplateRequest) ProtoMessage()    {}
func (*CreateTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_notify_279df0951896b475, []int{5}
}
func (m *CreateTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTemplateRequest.Unmarshal(m, b)
}
func (m *CreateTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTemplateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTemplateRequest.Merge(dst, src)
}
func (m *CreateTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTemplateRequest.Size(m)
}
func (m *CreateTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTemplateRequest proto.InternalMessageInfo

func (m *CreateTemplateRequest) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *CreateTemplateRequest) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *CreateTemplateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateTemplateRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *CreateTemplateRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_notify_279df0951896b475, []int{6}
}
func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (dst *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(dst, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type Services struct {
	Services             []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Services) Reset()         { *m = Services{} }
func (m *Services) String() string { return proto.CompactTextString(m) }
func (*Services) ProtoMessage()    {}
func (*Services) Descriptor() ([]byte, []int) {
	return fileDescriptor_notify_279df0951896b475, []int{7}
}
func (m *Services) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Services.Unmarshal(m, b)
}
func (m *Services) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Services.Marshal(b, m, deterministic)
}
func (dst *Services) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Services.Merge(dst, src)
}
func (m *Services) XXX_Size() int {
	return xxx_messageInfo_Services.Size(m)
}
func (m *Services) XXX_DiscardUnknown() {
	xxx_messageInfo_Services.DiscardUnknown(m)
}

var xxx_messageInfo_Services proto.InternalMessageInfo

func (m *Services) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type Templates struct {
	Templates            []*Template `protobuf:"bytes,1,rep,name=templates,proto3" json:"templates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Templates) Reset()         { *m = Templates{} }
func (m *Templates) String() string { return proto.CompactTextString(m) }
func (*Templates) ProtoMessage()    {}
func (*Templates) Descriptor() ([]byte, []int) {
	return fileDescriptor_notify_279df0951896b475, []int{8}
}
func (m *Templates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Templates.Unmarshal(m, b)
}
func (m *Templates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Templates.Marshal(b, m, deterministic)
}
func (dst *Templates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Templates.Merge(dst, src)
}
func (m *Templates) XXX_Size() int {
	return xxx_messageInfo_Templates.Size(m)
}
func (m *Templates) XXX_DiscardUnknown() {
	xxx_messageInfo_Templates.DiscardUnknown(m)
}

var xxx_messageInfo_Templates proto.InternalMessageInfo

func (m *Templates) GetTemplates() []*Template {
	if m != nil {
		return m.Templates
	}
	return nil
}

type Permissions struct {
	Permissions          []string `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permissions) Reset()         { *m = Permissions{} }
func (m *Permissions) String() string { return proto.CompactTextString(m) }
func (*Permissions) ProtoMessage()    {}
func (*Permissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_notify_279df0951896b475, []int{9}
}
func (m *Permissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permissions.Unmarshal(m, b)
}
func (m *Permissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permissions.Marshal(b, m, deterministic)
}
func (dst *Permissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permissions.Merge(dst, src)
}
func (m *Permissions) XXX_Size() int {
	return xxx_messageInfo_Permissions.Size(m)
}
func (m *Permissions) XXX_DiscardUnknown() {
	xxx_messageInfo_Permissions.DiscardUnknown(m)
}

var xxx_messageInfo_Permissions proto.InternalMessageInfo

func (m *Permissions) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type User struct {
	Id                   string                  `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AuthType             string                  `protobuf:"bytes,2,opt,name=auth_type,json=authType,proto3" json:"auth_type,omitempty"`
	EmailAddress         string                  `protobuf:"bytes,3,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	FailedLoginCount     int64                   `protobuf:"varint,4,opt,name=failed_login_count,json=failedLoginCount,proto3" json:"failed_login_count,omitempty"`
	PasswordChangedAt    string                  `protobuf:"bytes,5,opt,name=password_changed_at,json=passwordChangedAt,proto3" json:"password_changed_at,omitempty"`
	Permissions          map[string]*Permissions `protobuf:"bytes,6,rep,name=permissions,proto3" json:"permissions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_notify_279df0951896b475, []int{10}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAuthType() string {
	if m != nil {
		return m.AuthType
	}
	return ""
}

func (m *User) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *User) GetFailedLoginCount() int64 {
	if m != nil {
		return m.FailedLoginCount
	}
	return 0
}

func (m *User) GetPasswordChangedAt() string {
	if m != nil {
		return m.PasswordChangedAt
	}
	return ""
}

func (m *User) GetPermissions() map[string]*Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type Service struct {
	Active               bool     `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	Branding             string   `protobuf:"bytes,2,opt,name=branding,proto3" json:"branding,omitempty"`
	CreatedBy            string   `protobuf:"bytes,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	EmailFrom            string   `protobuf:"bytes,4,opt,name=email_from,json=emailFrom,proto3" json:"email_from,omitempty"`
	Id                   string   `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	OrganisationType     string   `protobuf:"bytes,8,opt,name=organisation_type,json=organisationType,proto3" json:"organisation_type,omitempty"`
	RateLimit            int64    `protobuf:"varint,9,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	AnnualBilling        []string `protobuf:"bytes,10,rep,name=annual_billing,json=annualBilling,proto3" json:"annual_billing,omitempty"`
	Permissions          []string `protobuf:"bytes,11,rep,name=permissions,proto3" json:"permissions,omitempty"`
	UserToService        []string `protobuf:"bytes,12,rep,name=user_to_service,json=userToService,proto3" json:"user_to_service,omitempty"`
	Users                []string `protobuf:"bytes,13,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_notify_279df0951896b475, []int{11}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (dst *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(dst, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Service) GetBranding() string {
	if m != nil {
		return m.Branding
	}
	return ""
}

func (m *Service) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Service) GetEmailFrom() string {
	if m != nil {
		return m.EmailFrom
	}
	return ""
}

func (m *Service) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetOrganisationType() string {
	if m != nil {
		return m.OrganisationType
	}
	return ""
}

func (m *Service) GetRateLimit() int64 {
	if m != nil {
		return m.RateLimit
	}
	return 0
}

func (m *Service) GetAnnualBilling() []string {
	if m != nil {
		return m.AnnualBilling
	}
	return nil
}

func (m *Service) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *Service) GetUserToService() []string {
	if m != nil {
		return m.UserToService
	}
	return nil
}

func (m *Service) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

type Template struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Subject              string   `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	TemplateType         string   `protobuf:"bytes,5,opt,name=template_type,json=templateType,proto3" json:"template_type,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy            string   `protobuf:"bytes,7,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Service              string   `protobuf:"bytes,8,opt,name=service,proto3" json:"service,omitempty"`
	ProcessType          string   `protobuf:"bytes,9,opt,name=process_type,json=processType,proto3" json:"process_type,omitempty"`
	Archived             bool     `protobuf:"varint,10,opt,name=archived,proto3" json:"archived,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Template) Reset()         { *m = Template{} }
func (m *Template) String() string { return proto.CompactTextString(m) }
func (*Template) ProtoMessage()    {}
func (*Template) Descriptor() ([]byte, []int) {
	return fileDescriptor_notify_279df0951896b475, []int{12}
}
func (m *Template) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Template.Unmarshal(m, b)
}
func (m *Template) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Template.Marshal(b, m, deterministic)
}
func (dst *Template) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Template.Merge(dst, src)
}
func (m *Template) XXX_Size() int {
	return xxx_messageInfo_Template.Size(m)
}
func (m *Template) XXX_DiscardUnknown() {
	xxx_messageInfo_Template.DiscardUnknown(m)
}

var xxx_messageInfo_Template proto.InternalMessageInfo

func (m *Template) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Template) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Template) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Template) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Template) GetTemplateType() string {
	if m != nil {
		return m.TemplateType
	}
	return ""
}

func (m *Template) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Template) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Template) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Template) GetProcessType() string {
	if m != nil {
		return m.ProcessType
	}
	return ""
}

func (m *Template) GetArchived() bool {
	if m != nil {
		return m.Archived
	}
	return false
}

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*ForUserRequest)(nil), "ForUserRequest")
	proto.RegisterType((*ServiceRequest)(nil), "ServiceRequest")
	proto.RegisterType((*TemplatesRequest)(nil), "TemplatesRequest")
	proto.RegisterType((*CreateTemplateRequest)(nil), "CreateTemplateRequest")
	proto.RegisterType((*Users)(nil), "Users")
	proto.RegisterType((*Services)(nil), "Services")
	proto.RegisterType((*Templates)(nil), "Templates")
	proto.RegisterType((*Permissions)(nil), "Permissions")
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterMapType((map[string]*Permissions)(nil), "User.PermissionsEntry")
	proto.RegisterType((*Service)(nil), "Service")
	proto.RegisterType((*Template)(nil), "Template")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotifyClient is the client API for Notify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotifyClient interface {
	GetUser(ctx context.Context, in *ForUserRequest, opts ...grpc.CallOption) (*User, error)
	GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Users, error)
	GetService(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*Service, error)
	ServicesForUser(ctx context.Context, in *ForUserRequest, opts ...grpc.CallOption) (*Services, error)
	TemplatesForService(ctx context.Context, in *TemplatesRequest, opts ...grpc.CallOption) (*Templates, error)
	CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*Template, error)
}

type notifyClient struct {
	cc *grpc.ClientConn
}

func NewNotifyClient(cc *grpc.ClientConn) NotifyClient {
	return &notifyClient{cc}
}

func (c *notifyClient) GetUser(ctx context.Context, in *ForUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/Notify/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/Notify/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) GetService(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/Notify/GetService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) ServicesForUser(ctx context.Context, in *ForUserRequest, opts ...grpc.CallOption) (*Services, error) {
	out := new(Services)
	err := c.cc.Invoke(ctx, "/Notify/ServicesForUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) TemplatesForService(ctx context.Context, in *TemplatesRequest, opts ...grpc.CallOption) (*Templates, error) {
	out := new(Templates)
	err := c.cc.Invoke(ctx, "/Notify/TemplatesForService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*Template, error) {
	out := new(Template)
	err := c.cc.Invoke(ctx, "/Notify/CreateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifyServer is the server API for Notify service.
type NotifyServer interface {
	GetUser(context.Context, *ForUserRequest) (*User, error)
	GetUsers(context.Context, *Empty) (*Users, error)
	GetService(context.Context, *ServiceRequest) (*Service, error)
	ServicesForUser(context.Context, *ForUserRequest) (*Services, error)
	TemplatesForService(context.Context, *TemplatesRequest) (*Templates, error)
	CreateTemplate(context.Context, *CreateTemplateRequest) (*Template, error)
}

func RegisterNotifyServer(s *grpc.Server, srv NotifyServer) {
	s.RegisterService(&_Notify_serviceDesc, srv)
}

func _Notify_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Notify/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).GetUser(ctx, req.(*ForUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Notify/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).GetUsers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Notify/GetService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).GetService(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_ServicesForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).ServicesForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Notify/ServicesForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).ServicesForUser(ctx, req.(*ForUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_TemplatesForService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).TemplatesForService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Notify/TemplatesForService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).TemplatesForService(ctx, req.(*TemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_CreateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).CreateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Notify/CreateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).CreateTemplate(ctx, req.(*CreateTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Notify",
	HandlerType: (*NotifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Notify_GetUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Notify_GetUsers_Handler,
		},
		{
			MethodName: "GetService",
			Handler:    _Notify_GetService_Handler,
		},
		{
			MethodName: "ServicesForUser",
			Handler:    _Notify_ServicesForUser_Handler,
		},
		{
			MethodName: "TemplatesForService",
			Handler:    _Notify_TemplatesForService_Handler,
		},
		{
			MethodName: "CreateTemplate",
			Handler:    _Notify_CreateTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notify.proto",
}

func init() { proto.RegisterFile("notify.proto", fileDescriptor_notify_279df0951896b475) }

var fileDescriptor_notify_279df0951896b475 = []byte{
	// 820 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6e, 0xe4, 0x34,
	0x14, 0x9e, 0xc9, 0x74, 0x26, 0xc9, 0x99, 0xe9, 0x74, 0xea, 0x85, 0x12, 0x65, 0xa9, 0x54, 0xbc,
	0x0b, 0xec, 0x0a, 0x94, 0x42, 0xf9, 0xd1, 0x8a, 0xbb, 0x6e, 0xb5, 0x5d, 0xad, 0x54, 0x21, 0x34,
	0x94, 0xeb, 0xc8, 0x93, 0xb8, 0xad, 0x21, 0x89, 0x83, 0xed, 0x29, 0xca, 0x8b, 0x70, 0x07, 0x0f,
	0xc5, 0xa3, 0xf0, 0x04, 0xc8, 0x7f, 0xf3, 0xd7, 0x45, 0xbb, 0x77, 0x39, 0xdf, 0x77, 0x1c, 0xdb,
	0xdf, 0x77, 0x8e, 0x0f, 0x4c, 0x1a, 0xae, 0xd8, 0x4d, 0x97, 0xb5, 0x82, 0x2b, 0x8e, 0x43, 0x18,
	0xbe, 0xaa, 0x5b, 0xd5, 0x61, 0x0c, 0xe1, 0x9c, 0xfe, 0xbe, 0xa4, 0x52, 0xa1, 0x8f, 0x20, 0x5c,
	0x4a, 0x2a, 0x72, 0x56, 0x26, 0xfd, 0x93, 0xfe, 0xb3, 0x78, 0x3e, 0xd2, 0xe1, 0x9b, 0x12, 0x3f,
	0x87, 0xe9, 0x25, 0x17, 0xbf, 0x48, 0x2a, 0xde, 0x99, 0x7a, 0x0a, 0xd3, 0x9f, 0xa9, 0xb8, 0x67,
	0x05, 0xf5, 0xa9, 0xc7, 0x00, 0xd2, 0x22, 0xeb, 0xec, 0xd8, 0x21, 0x6f, 0x4a, 0xfc, 0x35, 0xcc,
	0xae, 0x69, 0xdd, 0x56, 0x44, 0x51, 0xf9, 0x9e, 0x4b, 0xfe, 0xea, 0xc3, 0x87, 0x17, 0x82, 0x12,
	0x45, 0xfd, 0xca, 0xf7, 0x5b, 0xa8, 0xe9, 0xc2, 0xac, 0x2b, 0xf3, 0x45, 0x97, 0x04, 0x96, 0x76,
	0xc8, 0xcb, 0x0e, 0x21, 0xd8, 0x6b, 0x48, 0x4d, 0x93, 0x81, 0x21, 0xcc, 0x37, 0x4a, 0x20, 0x94,
	0xcb, 0xc5, 0xaf, 0xb4, 0x50, 0xc9, 0x9e, 0x81, 0x7d, 0xa8, 0x99, 0x82, 0x37, 0x8a, 0x36, 0x2a,
	0x19, 0x5a, 0xc6, 0x85, 0xf8, 0x29, 0x0c, 0xb5, 0x56, 0x12, 0x3d, 0x86, 0xa1, 0x96, 0x45, 0x26,
	0xfd, 0x93, 0xc1, 0xb3, 0xf1, 0xd9, 0x30, 0x33, 0x12, 0x5a, 0x0c, 0x7f, 0x05, 0x91, 0x53, 0x4a,
	0xa2, 0xa7, 0x10, 0xb9, 0x53, 0xfa, 0xdc, 0x28, 0xf3, 0x32, 0xae, 0x18, 0xfc, 0x2d, 0xc4, 0x2b,
	0xa9, 0xd0, 0xe7, 0x10, 0x2b, 0x1f, 0xb8, 0x35, 0x71, 0xb6, 0xd2, 0x63, 0xcd, 0xe1, 0x53, 0x18,
	0xff, 0x44, 0x45, 0xcd, 0xa4, 0x64, 0xbc, 0x91, 0xe8, 0x04, 0xc6, 0xed, 0x3a, 0x34, 0x2b, 0xe3,
	0xf9, 0x26, 0x84, 0xff, 0x09, 0x60, 0x4f, 0x1f, 0x14, 0x4d, 0x21, 0x60, 0x65, 0x12, 0x9a, 0xcb,
	0x05, 0xac, 0x5c, 0xe9, 0xd3, 0xdf, 0xd0, 0xe7, 0x31, 0xc4, 0x64, 0xa9, 0xee, 0x72, 0xd5, 0xb5,
	0xd4, 0x29, 0x1a, 0x69, 0xe0, 0xba, 0x6b, 0x29, 0x7a, 0x02, 0xfb, 0xb4, 0x26, 0xac, 0xca, 0x49,
	0x59, 0x0a, 0x2a, 0xa5, 0x53, 0x76, 0x62, 0xc0, 0x73, 0x8b, 0xa1, 0x2f, 0x01, 0xdd, 0x10, 0x56,
	0xd1, 0x32, 0xaf, 0xf8, 0x2d, 0x6b, 0xf2, 0x82, 0x2f, 0x1b, 0x2b, 0xf6, 0x60, 0x3e, 0xb3, 0xcc,
	0x95, 0x26, 0x2e, 0x34, 0x8e, 0x32, 0x78, 0xd4, 0x12, 0x29, 0xff, 0xe0, 0xa2, 0xcc, 0x8b, 0x3b,
	0xd2, 0xdc, 0xd2, 0x32, 0x27, 0xde, 0x81, 0x43, 0x4f, 0x5d, 0x58, 0xe6, 0x5c, 0xa1, 0x17, 0xdb,
	0xd7, 0x1d, 0x19, 0xa1, 0x8e, 0x8c, 0x11, 0xd9, 0x86, 0x2c, 0xaf, 0x1a, 0x25, 0xba, 0x2d, 0x19,
	0xd2, 0x2b, 0x98, 0xed, 0x26, 0xa0, 0x19, 0x0c, 0x7e, 0xa3, 0x9d, 0x13, 0x40, 0x7f, 0x22, 0x0c,
	0xc3, 0x7b, 0x52, 0x2d, 0xed, 0xdd, 0xc7, 0x67, 0x93, 0xcd, 0x9f, 0xce, 0x2d, 0xf5, 0x43, 0xf0,
	0xa2, 0x8f, 0xff, 0x0d, 0x20, 0x74, 0x8e, 0xa2, 0x23, 0x18, 0x91, 0x42, 0xb1, 0x7b, 0xab, 0x64,
	0x34, 0x77, 0x11, 0x4a, 0x21, 0x5a, 0x08, 0xd2, 0x94, 0xac, 0xb9, 0xf5, 0x52, 0xfa, 0x78, 0xa7,
	0x74, 0x07, 0xbb, 0xa5, 0x7b, 0x0c, 0x60, 0x95, 0xbe, 0x11, 0xbc, 0x76, 0x95, 0x1a, 0x1b, 0xe4,
	0x52, 0xf0, 0xda, 0x39, 0x39, 0x7c, 0xe0, 0xe4, 0x68, 0xc3, 0xc9, 0x2f, 0xe0, 0x90, 0x8b, 0x5b,
	0xd2, 0x30, 0x49, 0x14, 0xe3, 0x8d, 0x75, 0x34, 0x32, 0x09, 0xb3, 0x4d, 0xc2, 0x38, 0x7b, 0x0c,
	0x20, 0x88, 0xa2, 0x79, 0xc5, 0x6a, 0xa6, 0x92, 0xd8, 0x98, 0x15, 0x6b, 0xe4, 0x4a, 0x03, 0xe8,
	0x53, 0x98, 0x92, 0xa6, 0x59, 0x92, 0x2a, 0x5f, 0xb0, 0xaa, 0xd2, 0xf7, 0x01, 0x53, 0x67, 0xfb,
	0x16, 0x7d, 0x69, 0xc1, 0xdd, 0x5a, 0x1c, 0x3f, 0xa8, 0x45, 0xf4, 0x19, 0x1c, 0x98, 0x77, 0x46,
	0xf1, 0xdc, 0xb5, 0x41, 0x32, 0xb1, 0x7f, 0xd2, 0xf0, 0x35, 0xf7, 0x92, 0x7e, 0xe0, 0x3b, 0x6d,
	0xdf, 0xb0, 0xae, 0xc5, 0xfe, 0x0e, 0x20, 0xf2, 0x2d, 0xe1, 0x34, 0xe8, 0x3f, 0xd0, 0x20, 0x78,
	0x7b, 0xb7, 0x0f, 0xfe, 0xb7, 0xdb, 0xf7, 0xb6, 0xba, 0x5d, 0x17, 0xb9, 0x6f, 0x36, 0xab, 0x99,
	0x95, 0x79, 0xe2, 0x41, 0xaf, 0x97, 0xb7, 0x8f, 0x28, 0x27, 0xbb, 0xb7, 0xef, 0x5c, 0xed, 0xb8,
	0x1b, 0xee, 0xba, 0xab, 0x8f, 0xe5, 0x6e, 0x1f, 0xb9, 0x63, 0xb9, 0x7b, 0x7f, 0x02, 0x93, 0x56,
	0xf0, 0x82, 0x4a, 0x69, 0xf7, 0x8e, 0x0d, 0x3d, 0x76, 0x98, 0xd9, 0x3a, 0x85, 0x88, 0x88, 0xe2,
	0x8e, 0xdd, 0xd3, 0x32, 0x01, 0x53, 0x6f, 0xab, 0xf8, 0xec, 0xcf, 0x00, 0x46, 0x3f, 0x9a, 0xb1,
	0x80, 0x9e, 0x40, 0xf8, 0x9a, 0x2a, 0xd3, 0xf7, 0x07, 0xd9, 0xf6, 0x6b, 0x9f, 0xda, 0x87, 0x0b,
	0xf7, 0xd0, 0xc7, 0x10, 0xb9, 0x24, 0x89, 0x46, 0x99, 0x19, 0x20, 0xe9, 0xc8, 0x90, 0x12, 0xf7,
	0xd0, 0x73, 0x80, 0xd7, 0x54, 0x79, 0x4b, 0x0e, 0xb2, 0xed, 0x41, 0x90, 0xae, 0x9e, 0x34, 0xdc,
	0x43, 0xa7, 0x70, 0xe0, 0x1f, 0x3f, 0xb7, 0xd7, 0xc3, 0x5d, 0x63, 0x9f, 0xaf, 0xff, 0xfd, 0x3d,
	0x3c, 0x5a, 0xbd, 0x7d, 0x97, 0x5c, 0xf8, 0x4d, 0x0e, 0xb3, 0xdd, 0xe1, 0x91, 0xc2, 0x1a, 0xc2,
	0x3d, 0xf4, 0x1d, 0x4c, 0xb7, 0x47, 0x05, 0x3a, 0xca, 0xde, 0x3a, 0x3b, 0xd2, 0xf5, 0xeb, 0x89,
	0x7b, 0x8b, 0x91, 0x99, 0x92, 0xdf, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xf9, 0xfd, 0x4d,
	0x35, 0x07, 0x00, 0x00,
}