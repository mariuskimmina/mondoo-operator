// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: captain.proto

package captain

import (
	api "go.mondoo.com/mondoo-operator/tests/framework/nexus/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_captain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_captain_proto_rawDescGZIP(), []int{0}
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mrn string `protobuf:"bytes,1,opt,name=mrn,proto3" json:"mrn,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_captain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_captain_proto_rawDescGZIP(), []int{1}
}

func (x *ID) GetMrn() string {
	if x != nil {
		return x.Mrn
	}
	return ""
}

type Organization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string            `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name         string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mrn          string            `protobuf:"bytes,3,opt,name=mrn,proto3" json:"mrn,omitempty"`
	Capabilities map[string]string `protobuf:"bytes,4,rep,name=capabilities,proto3" json:"capabilities,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Organization) Reset() {
	*x = Organization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Organization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Organization) ProtoMessage() {}

func (x *Organization) ProtoReflect() protoreflect.Message {
	mi := &file_captain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Organization.ProtoReflect.Descriptor instead.
func (*Organization) Descriptor() ([]byte, []int) {
	return file_captain_proto_rawDescGZIP(), []int{2}
}

func (x *Organization) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Organization) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Organization) GetMrn() string {
	if x != nil {
		return x.Mrn
	}
	return ""
}

func (x *Organization) GetCapabilities() map[string]string {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

type Space struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string        `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name         string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mrn          string        `protobuf:"bytes,3,opt,name=mrn,proto3" json:"mrn,omitempty"`
	Organization *Organization `protobuf:"bytes,4,opt,name=organization,proto3" json:"organization,omitempty"`
	// string title = 5;
	// string description = 6;
	Settings *SpaceSettings `protobuf:"bytes,5,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *Space) Reset() {
	*x = Space{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Space) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Space) ProtoMessage() {}

func (x *Space) ProtoReflect() protoreflect.Message {
	mi := &file_captain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Space.ProtoReflect.Descriptor instead.
func (*Space) Descriptor() ([]byte, []int) {
	return file_captain_proto_rawDescGZIP(), []int{3}
}

func (x *Space) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Space) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Space) GetMrn() string {
	if x != nil {
		return x.Mrn
	}
	return ""
}

func (x *Space) GetOrganization() *Organization {
	if x != nil {
		return x.Organization
	}
	return nil
}

func (x *Space) GetSettings() *SpaceSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type CreateSpaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string        `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name         string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mrn          string        `protobuf:"bytes,3,opt,name=mrn,proto3" json:"mrn,omitempty"`
	Organization *Organization `protobuf:"bytes,4,opt,name=organization,proto3" json:"organization,omitempty"`
	// string title = 6;
	// string description = 7;
	OwnerMrn string         `protobuf:"bytes,7,opt,name=ownerMrn,proto3" json:"ownerMrn,omitempty"`
	Settings *SpaceSettings `protobuf:"bytes,8,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *CreateSpaceRequest) Reset() {
	*x = CreateSpaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSpaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSpaceRequest) ProtoMessage() {}

func (x *CreateSpaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_captain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSpaceRequest.ProtoReflect.Descriptor instead.
func (*CreateSpaceRequest) Descriptor() ([]byte, []int) {
	return file_captain_proto_rawDescGZIP(), []int{4}
}

func (x *CreateSpaceRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *CreateSpaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSpaceRequest) GetMrn() string {
	if x != nil {
		return x.Mrn
	}
	return ""
}

func (x *CreateSpaceRequest) GetOrganization() *Organization {
	if x != nil {
		return x.Organization
	}
	return nil
}

func (x *CreateSpaceRequest) GetOwnerMrn() string {
	if x != nil {
		return x.OwnerMrn
	}
	return ""
}

func (x *CreateSpaceRequest) GetSettings() *SpaceSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type SpaceSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TerminatedAssetsConfiguration *TerminatedAssetsConfiguration `protobuf:"bytes,1,opt,name=terminatedAssetsConfiguration,proto3" json:"terminatedAssetsConfiguration,omitempty"`
}

func (x *SpaceSettings) Reset() {
	*x = SpaceSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceSettings) ProtoMessage() {}

func (x *SpaceSettings) ProtoReflect() protoreflect.Message {
	mi := &file_captain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceSettings.ProtoReflect.Descriptor instead.
func (*SpaceSettings) Descriptor() ([]byte, []int) {
	return file_captain_proto_rawDescGZIP(), []int{5}
}

func (x *SpaceSettings) GetTerminatedAssetsConfiguration() *TerminatedAssetsConfiguration {
	if x != nil {
		return x.TerminatedAssetsConfiguration
	}
	return nil
}

type TerminatedAssetsConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cleanup bool  `protobuf:"varint,1,opt,name=cleanup,proto3" json:"cleanup,omitempty"`
	After   int32 `protobuf:"varint,2,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *TerminatedAssetsConfiguration) Reset() {
	*x = TerminatedAssetsConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminatedAssetsConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminatedAssetsConfiguration) ProtoMessage() {}

func (x *TerminatedAssetsConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_captain_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminatedAssetsConfiguration.ProtoReflect.Descriptor instead.
func (*TerminatedAssetsConfiguration) Descriptor() ([]byte, []int) {
	return file_captain_proto_rawDescGZIP(), []int{6}
}

func (x *TerminatedAssetsConfiguration) GetCleanup() bool {
	if x != nil {
		return x.Cleanup
	}
	return false
}

func (x *TerminatedAssetsConfiguration) GetAfter() int32 {
	if x != nil {
		return x.After
	}
	return 0
}

type SpacesQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page            *api.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	OrganizationMrn string           `protobuf:"bytes,2,opt,name=organizationMrn,proto3" json:"organizationMrn,omitempty"`
}

func (x *SpacesQuery) Reset() {
	*x = SpacesQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captain_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpacesQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpacesQuery) ProtoMessage() {}

func (x *SpacesQuery) ProtoReflect() protoreflect.Message {
	mi := &file_captain_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpacesQuery.ProtoReflect.Descriptor instead.
func (*SpacesQuery) Descriptor() ([]byte, []int) {
	return file_captain_proto_rawDescGZIP(), []int{7}
}

func (x *SpacesQuery) GetPage() *api.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *SpacesQuery) GetOrganizationMrn() string {
	if x != nil {
		return x.OrganizationMrn
	}
	return ""
}

type SpacesPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *api.PageInfo `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	List []*Space      `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *SpacesPage) Reset() {
	*x = SpacesPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captain_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpacesPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpacesPage) ProtoMessage() {}

func (x *SpacesPage) ProtoReflect() protoreflect.Message {
	mi := &file_captain_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpacesPage.ProtoReflect.Descriptor instead.
func (*SpacesPage) Descriptor() ([]byte, []int) {
	return file_captain_proto_rawDescGZIP(), []int{8}
}

func (x *SpacesPage) GetPage() *api.PageInfo {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *SpacesPage) GetList() []*Space {
	if x != nil {
		return x.List
	}
	return nil
}

var File_captain_proto protoreflect.FileDescriptor

var file_captain_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x1a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x16, 0x0a,
	0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x72, 0x6e, 0x22, 0xdc, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x72,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x72, 0x6e, 0x12, 0x55, 0x0a, 0x0c,
	0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x61, 0x70, 0x74,
	0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x1a, 0x3f, 0x0a, 0x11, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xc0, 0x01, 0x0a, 0x05, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x72, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x72, 0x6e, 0x12, 0x43, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x6f, 0x6e,
	0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x08, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x6f,
	0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xe9, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x72, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x72, 0x6e, 0x12, 0x43, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x6f, 0x6e,
	0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x4d, 0x72, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x4d, 0x72, 0x6e, 0x12, 0x3c, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f,
	0x2e, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x0d, 0x53, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x76, 0x0a, 0x1d, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6d,
	0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1d,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4f, 0x0a,
	0x1d, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0x6b,
	0x0a, 0x0b, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x32, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x6f,
	0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x72, 0x6e, 0x22, 0x6b, 0x0a, 0x0a, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x50, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f,
	0x2e, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f,
	0x6f, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xab, 0x02, 0x0a, 0x07, 0x43, 0x61, 0x70,
	0x74, 0x61, 0x69, 0x6e, 0x12, 0x50, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x61, 0x70,
	0x74, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x6f, 0x6e,
	0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x15, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x61, 0x70, 0x74,
	0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x44, 0x1a, 0x18, 0x2e, 0x6d, 0x6f, 0x6e, 0x64,
	0x6f, 0x6f, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x61, 0x70,
	0x74, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x1a, 0x1d, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x61, 0x70,
	0x74, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x50, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x15, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x61, 0x70,
	0x74, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x44, 0x1a, 0x18, 0x2e, 0x6d, 0x6f, 0x6e,
	0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x6f, 0x2e, 0x6d, 0x6f, 0x6e,
	0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2d, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_captain_proto_rawDescOnce sync.Once
	file_captain_proto_rawDescData = file_captain_proto_rawDesc
)

func file_captain_proto_rawDescGZIP() []byte {
	file_captain_proto_rawDescOnce.Do(func() {
		file_captain_proto_rawDescData = protoimpl.X.CompressGZIP(file_captain_proto_rawDescData)
	})
	return file_captain_proto_rawDescData
}

var file_captain_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_captain_proto_goTypes = []interface{}{
	(*Empty)(nil),                         // 0: mondoo.captain.v1.Empty
	(*ID)(nil),                            // 1: mondoo.captain.v1.ID
	(*Organization)(nil),                  // 2: mondoo.captain.v1.Organization
	(*Space)(nil),                         // 3: mondoo.captain.v1.Space
	(*CreateSpaceRequest)(nil),            // 4: mondoo.captain.v1.CreateSpaceRequest
	(*SpaceSettings)(nil),                 // 5: mondoo.captain.v1.SpaceSettings
	(*TerminatedAssetsConfiguration)(nil), // 6: mondoo.captain.v1.TerminatedAssetsConfiguration
	(*SpacesQuery)(nil),                   // 7: mondoo.captain.v1.SpacesQuery
	(*SpacesPage)(nil),                    // 8: mondoo.captain.v1.SpacesPage
	nil,                                   // 9: mondoo.captain.v1.Organization.CapabilitiesEntry
	(*api.PageRequest)(nil),               // 10: mondoo.captain.v1.PageRequest
	(*api.PageInfo)(nil),                  // 11: mondoo.captain.v1.PageInfo
}
var file_captain_proto_depIdxs = []int32{
	9,  // 0: mondoo.captain.v1.Organization.capabilities:type_name -> mondoo.captain.v1.Organization.CapabilitiesEntry
	2,  // 1: mondoo.captain.v1.Space.organization:type_name -> mondoo.captain.v1.Organization
	5,  // 2: mondoo.captain.v1.Space.settings:type_name -> mondoo.captain.v1.SpaceSettings
	2,  // 3: mondoo.captain.v1.CreateSpaceRequest.organization:type_name -> mondoo.captain.v1.Organization
	5,  // 4: mondoo.captain.v1.CreateSpaceRequest.settings:type_name -> mondoo.captain.v1.SpaceSettings
	6,  // 5: mondoo.captain.v1.SpaceSettings.terminatedAssetsConfiguration:type_name -> mondoo.captain.v1.TerminatedAssetsConfiguration
	10, // 6: mondoo.captain.v1.SpacesQuery.page:type_name -> mondoo.captain.v1.PageRequest
	11, // 7: mondoo.captain.v1.SpacesPage.page:type_name -> mondoo.captain.v1.PageInfo
	3,  // 8: mondoo.captain.v1.SpacesPage.list:type_name -> mondoo.captain.v1.Space
	4,  // 9: mondoo.captain.v1.Captain.CreateSpace:input_type -> mondoo.captain.v1.CreateSpaceRequest
	1,  // 10: mondoo.captain.v1.Captain.GetSpace:input_type -> mondoo.captain.v1.ID
	7,  // 11: mondoo.captain.v1.Captain.ListSpaces:input_type -> mondoo.captain.v1.SpacesQuery
	1,  // 12: mondoo.captain.v1.Captain.DeleteSpace:input_type -> mondoo.captain.v1.ID
	3,  // 13: mondoo.captain.v1.Captain.CreateSpace:output_type -> mondoo.captain.v1.Space
	3,  // 14: mondoo.captain.v1.Captain.GetSpace:output_type -> mondoo.captain.v1.Space
	8,  // 15: mondoo.captain.v1.Captain.ListSpaces:output_type -> mondoo.captain.v1.SpacesPage
	0,  // 16: mondoo.captain.v1.Captain.DeleteSpace:output_type -> mondoo.captain.v1.Empty
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_captain_proto_init() }
func file_captain_proto_init() {
	if File_captain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_captain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_captain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_captain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Organization); i {
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
		file_captain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Space); i {
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
		file_captain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSpaceRequest); i {
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
		file_captain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceSettings); i {
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
		file_captain_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminatedAssetsConfiguration); i {
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
		file_captain_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpacesQuery); i {
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
		file_captain_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpacesPage); i {
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
			RawDescriptor: file_captain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_captain_proto_goTypes,
		DependencyIndexes: file_captain_proto_depIdxs,
		MessageInfos:      file_captain_proto_msgTypes,
	}.Build()
	File_captain_proto = out.File
	file_captain_proto_rawDesc = nil
	file_captain_proto_goTypes = nil
	file_captain_proto_depIdxs = nil
}
