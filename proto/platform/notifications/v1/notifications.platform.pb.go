// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: notifications.platform.proto

package v1

import (
	_ "chainguard.dev/sdk/proto/annotations"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NotificationsFilter_Location int32

const (
	// The "unset" value.
	NotificationsFilter_UNSPECIFIED NotificationsFilter_Location = 0
	// Console.
	NotificationsFilter_CONSOLE NotificationsFilter_Location = 1
	// Directory.
	NotificationsFilter_DIRECTORY NotificationsFilter_Location = 2
	// Chainctl.
	NotificationsFilter_CHAINCTL NotificationsFilter_Location = 3
)

// Enum value maps for NotificationsFilter_Location.
var (
	NotificationsFilter_Location_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "CONSOLE",
		2: "DIRECTORY",
		3: "CHAINCTL",
	}
	NotificationsFilter_Location_value = map[string]int32{
		"UNSPECIFIED": 0,
		"CONSOLE":     1,
		"DIRECTORY":   2,
		"CHAINCTL":    3,
	}
)

func (x NotificationsFilter_Location) Enum() *NotificationsFilter_Location {
	p := new(NotificationsFilter_Location)
	*p = x
	return p
}

func (x NotificationsFilter_Location) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationsFilter_Location) Descriptor() protoreflect.EnumDescriptor {
	return file_notifications_platform_proto_enumTypes[0].Descriptor()
}

func (NotificationsFilter_Location) Type() protoreflect.EnumType {
	return &file_notifications_platform_proto_enumTypes[0]
}

func (x NotificationsFilter_Location) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationsFilter_Location.Descriptor instead.
func (NotificationsFilter_Location) EnumDescriptor() ([]byte, []int) {
	return file_notifications_platform_proto_rawDescGZIP(), []int{5, 0}
}

type NotificationsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Notification `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *NotificationsList) Reset() {
	*x = NotificationsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_platform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationsList) ProtoMessage() {}

func (x *NotificationsList) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_platform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationsList.ProtoReflect.Descriptor instead.
func (*NotificationsList) Descriptor() ([]byte, []int) {
	return file_notifications_platform_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationsList) GetItems() []*Notification {
	if x != nil {
		return x.Items
	}
	return nil
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the identifier of this specific notification.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// when the notification was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// when the notification was last updated.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// id of the category the notification belongs to.
	CategoryId string `protobuf:"bytes,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	// subject or title of the notification.
	Subject string `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	// the body of the notification.
	Note string `protobuf:"bytes,6,opt,name=note,proto3" json:"note,omitempty"`
	// date when the notification was created.
	NoteDate *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=note_date,json=noteDate,proto3" json:"note_date,omitempty"`
	// tags associated with the notification.
	Tags []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	// account associated with the notification.
	Account *Account `protobuf:"bytes,9,opt,name=account,proto3" json:"account,omitempty"`
	// author associated with the notification.
	Author *Author `protobuf:"bytes,10,opt,name=author,proto3" json:"author,omitempty"`
	// account id associated with the notification.
	AccountId string `protobuf:"bytes,11,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// author id associated with the notification.
	AuthorId string `protobuf:"bytes,12,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	// traits associated with the notification.
	Traits *Traits `protobuf:"bytes,13,opt,name=traits,proto3" json:"traits,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_platform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_platform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_notifications_platform_proto_rawDescGZIP(), []int{1}
}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Notification) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Notification) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *Notification) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Notification) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Notification) GetNoteDate() *timestamppb.Timestamp {
	if x != nil {
		return x.NoteDate
	}
	return nil
}

func (x *Notification) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Notification) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *Notification) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Notification) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Notification) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Notification) GetTraits() *Traits {
	if x != nil {
		return x.Traits
	}
	return nil
}

type Traits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// expiration date of the notification.
	Expires *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=expires,proto3" json:"expires,omitempty"`
	// show notification in console.
	Console bool `protobuf:"varint,2,opt,name=console,proto3" json:"console,omitempty"`
	// show notification in directory.
	Directory bool `protobuf:"varint,3,opt,name=directory,proto3" json:"directory,omitempty"`
	// start date of the notification.
	Starts *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=starts,proto3" json:"starts,omitempty"`
	// whether or not notification has been approved.
	Approved bool `protobuf:"varint,5,opt,name=approved,proto3" json:"approved,omitempty"`
	// show notification for chainctl users.
	Chainctl bool `protobuf:"varint,6,opt,name=chainctl,proto3" json:"chainctl,omitempty"`
}

func (x *Traits) Reset() {
	*x = Traits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_platform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Traits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Traits) ProtoMessage() {}

func (x *Traits) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_platform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Traits.ProtoReflect.Descriptor instead.
func (*Traits) Descriptor() ([]byte, []int) {
	return file_notifications_platform_proto_rawDescGZIP(), []int{2}
}

func (x *Traits) GetExpires() *timestamppb.Timestamp {
	if x != nil {
		return x.Expires
	}
	return nil
}

func (x *Traits) GetConsole() bool {
	if x != nil {
		return x.Console
	}
	return false
}

func (x *Traits) GetDirectory() bool {
	if x != nil {
		return x.Directory
	}
	return false
}

func (x *Traits) GetStarts() *timestamppb.Timestamp {
	if x != nil {
		return x.Starts
	}
	return nil
}

func (x *Traits) GetApproved() bool {
	if x != nil {
		return x.Approved
	}
	return false
}

func (x *Traits) GetChainctl() bool {
	if x != nil {
		return x.Chainctl
	}
	return false
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the account.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// when the account was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// when the account was last updated.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// name of the account.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_platform_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_platform_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_notifications_platform_proto_rawDescGZIP(), []int{3}
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Account) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the author.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name of the author.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// email of the author.
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_platform_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_platform_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_notifications_platform_proto_rawDescGZIP(), []int{4}
}

func (x *Author) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type NotificationsFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the notification.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Location of where notifications are displayed.
	Location NotificationsFilter_Location `protobuf:"varint,2,opt,name=location,proto3,enum=chainguard.platform.notifications.NotificationsFilter_Location" json:"location,omitempty"`
}

func (x *NotificationsFilter) Reset() {
	*x = NotificationsFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_platform_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationsFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationsFilter) ProtoMessage() {}

func (x *NotificationsFilter) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_platform_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationsFilter.ProtoReflect.Descriptor instead.
func (*NotificationsFilter) Descriptor() ([]byte, []int) {
	return file_notifications_platform_proto_rawDescGZIP(), []int{5}
}

func (x *NotificationsFilter) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NotificationsFilter) GetLocation() NotificationsFilter_Location {
	if x != nil {
		return x.Location
	}
	return NotificationsFilter_UNSPECIFIED
}

var File_notifications_platform_proto protoreflect.FileDescriptor

var file_notifications_platform_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x45, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0xb8, 0x04, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x6e, 0x6f,
	0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x44, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x41, 0x0a,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x06,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x52, 0x06, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x22,
	0xe2, 0x01, 0x0a, 0x06, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x63, 0x74, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x63, 0x74, 0x6c, 0x22, 0xa3, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x06, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xc9,
	0x01, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x5b, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x53, 0x4f, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x48, 0x41, 0x49, 0x4e, 0x43, 0x54, 0x4c, 0x10, 0x03, 0x32, 0xbe, 0x01, 0x0a, 0x0d, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xac, 0x01, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61,
	0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x34, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x8a, 0xaf, 0xa8, 0xd2,
	0x05, 0x09, 0x12, 0x07, 0x0a, 0x03, 0x67, 0xc5, 0x0c, 0x10, 0x01, 0x42, 0x80, 0x01, 0x0a, 0x2c,
	0x64, 0x65, 0x76, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x1a, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notifications_platform_proto_rawDescOnce sync.Once
	file_notifications_platform_proto_rawDescData = file_notifications_platform_proto_rawDesc
)

func file_notifications_platform_proto_rawDescGZIP() []byte {
	file_notifications_platform_proto_rawDescOnce.Do(func() {
		file_notifications_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_notifications_platform_proto_rawDescData)
	})
	return file_notifications_platform_proto_rawDescData
}

var file_notifications_platform_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_notifications_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_notifications_platform_proto_goTypes = []interface{}{
	(NotificationsFilter_Location)(0), // 0: chainguard.platform.notifications.NotificationsFilter.Location
	(*NotificationsList)(nil),         // 1: chainguard.platform.notifications.NotificationsList
	(*Notification)(nil),              // 2: chainguard.platform.notifications.Notification
	(*Traits)(nil),                    // 3: chainguard.platform.notifications.Traits
	(*Account)(nil),                   // 4: chainguard.platform.notifications.Account
	(*Author)(nil),                    // 5: chainguard.platform.notifications.Author
	(*NotificationsFilter)(nil),       // 6: chainguard.platform.notifications.NotificationsFilter
	(*timestamppb.Timestamp)(nil),     // 7: google.protobuf.Timestamp
}
var file_notifications_platform_proto_depIdxs = []int32{
	2,  // 0: chainguard.platform.notifications.NotificationsList.items:type_name -> chainguard.platform.notifications.Notification
	7,  // 1: chainguard.platform.notifications.Notification.created_at:type_name -> google.protobuf.Timestamp
	7,  // 2: chainguard.platform.notifications.Notification.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 3: chainguard.platform.notifications.Notification.note_date:type_name -> google.protobuf.Timestamp
	4,  // 4: chainguard.platform.notifications.Notification.account:type_name -> chainguard.platform.notifications.Account
	5,  // 5: chainguard.platform.notifications.Notification.author:type_name -> chainguard.platform.notifications.Author
	3,  // 6: chainguard.platform.notifications.Notification.traits:type_name -> chainguard.platform.notifications.Traits
	7,  // 7: chainguard.platform.notifications.Traits.expires:type_name -> google.protobuf.Timestamp
	7,  // 8: chainguard.platform.notifications.Traits.starts:type_name -> google.protobuf.Timestamp
	7,  // 9: chainguard.platform.notifications.Account.created_at:type_name -> google.protobuf.Timestamp
	7,  // 10: chainguard.platform.notifications.Account.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 11: chainguard.platform.notifications.NotificationsFilter.location:type_name -> chainguard.platform.notifications.NotificationsFilter.Location
	6,  // 12: chainguard.platform.notifications.Notifications.List:input_type -> chainguard.platform.notifications.NotificationsFilter
	1,  // 13: chainguard.platform.notifications.Notifications.List:output_type -> chainguard.platform.notifications.NotificationsList
	13, // [13:14] is the sub-list for method output_type
	12, // [12:13] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_notifications_platform_proto_init() }
func file_notifications_platform_proto_init() {
	if File_notifications_platform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notifications_platform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationsList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_notifications_platform_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_notifications_platform_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Traits); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_notifications_platform_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_notifications_platform_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_notifications_platform_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationsFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notifications_platform_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notifications_platform_proto_goTypes,
		DependencyIndexes: file_notifications_platform_proto_depIdxs,
		EnumInfos:         file_notifications_platform_proto_enumTypes,
		MessageInfos:      file_notifications_platform_proto_msgTypes,
	}.Build()
	File_notifications_platform_proto = out.File
	file_notifications_platform_proto_rawDesc = nil
	file_notifications_platform_proto_goTypes = nil
	file_notifications_platform_proto_depIdxs = nil
}
