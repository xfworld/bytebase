// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: v1/group_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GroupMember_Role int32

const (
	GroupMember_ROLE_UNSPECIFIED GroupMember_Role = 0
	GroupMember_OWNER            GroupMember_Role = 1
	GroupMember_MEMBER           GroupMember_Role = 2
)

// Enum value maps for GroupMember_Role.
var (
	GroupMember_Role_name = map[int32]string{
		0: "ROLE_UNSPECIFIED",
		1: "OWNER",
		2: "MEMBER",
	}
	GroupMember_Role_value = map[string]int32{
		"ROLE_UNSPECIFIED": 0,
		"OWNER":            1,
		"MEMBER":           2,
	}
)

func (x GroupMember_Role) Enum() *GroupMember_Role {
	p := new(GroupMember_Role)
	*p = x
	return p
}

func (x GroupMember_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupMember_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_group_service_proto_enumTypes[0].Descriptor()
}

func (GroupMember_Role) Type() protoreflect.EnumType {
	return &file_v1_group_service_proto_enumTypes[0]
}

func (x GroupMember_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupMember_Role.Descriptor instead.
func (GroupMember_Role) EnumDescriptor() ([]byte, []int) {
	return file_v1_group_service_proto_rawDescGZIP(), []int{6, 0}
}

type GetGroupRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the group to retrieve.
	// Format: groups/{email}
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetGroupRequest) Reset() {
	*x = GetGroupRequest{}
	mi := &file_v1_group_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupRequest) ProtoMessage() {}

func (x *GetGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_group_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupRequest.ProtoReflect.Descriptor instead.
func (*GetGroupRequest) Descriptor() ([]byte, []int) {
	return file_v1_group_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListGroupsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Not used.
	// The maximum number of groups to return. The service may return fewer than
	// this value.
	// If unspecified, at most 10 groups will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Not used.
	// A page token, received from a previous `ListGroups` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListGroups` must match
	// the call that provided the page token.
	PageToken     string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListGroupsRequest) Reset() {
	*x = ListGroupsRequest{}
	mi := &file_v1_group_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupsRequest) ProtoMessage() {}

func (x *ListGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_group_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupsRequest.ProtoReflect.Descriptor instead.
func (*ListGroupsRequest) Descriptor() ([]byte, []int) {
	return file_v1_group_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListGroupsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListGroupsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListGroupsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The groups from the specified request.
	Groups []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListGroupsResponse) Reset() {
	*x = ListGroupsResponse{}
	mi := &file_v1_group_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupsResponse) ProtoMessage() {}

func (x *ListGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_group_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupsResponse.ProtoReflect.Descriptor instead.
func (*ListGroupsResponse) Descriptor() ([]byte, []int) {
	return file_v1_group_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListGroupsResponse) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *ListGroupsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateGroupRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The group to create.
	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	// The email to use for the group, which will become the final component
	// of the group's resource name.
	GroupEmail    string `protobuf:"bytes,2,opt,name=group_email,json=groupEmail,proto3" json:"group_email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	mi := &file_v1_group_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_group_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_v1_group_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateGroupRequest) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *CreateGroupRequest) GetGroupEmail() string {
	if x != nil {
		return x.GroupEmail
	}
	return ""
}

type UpdateGroupRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The group to update.
	//
	// The group's `name` field is used to identify the group to update.
	// Format: groups/{email}
	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	// The list of fields to update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// If set to true, and the group is not found, a new group will be created.
	// In this situation, `update_mask` is ignored.
	AllowMissing  bool `protobuf:"varint,3,opt,name=allow_missing,json=allowMissing,proto3" json:"allow_missing,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateGroupRequest) Reset() {
	*x = UpdateGroupRequest{}
	mi := &file_v1_group_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupRequest) ProtoMessage() {}

func (x *UpdateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_group_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return file_v1_group_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateGroupRequest) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *UpdateGroupRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateGroupRequest) GetAllowMissing() bool {
	if x != nil {
		return x.AllowMissing
	}
	return false
}

type DeleteGroupRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the group to delete.
	// Format: groups/{email}
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteGroupRequest) Reset() {
	*x = DeleteGroupRequest{}
	mi := &file_v1_group_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupRequest) ProtoMessage() {}

func (x *DeleteGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_group_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return file_v1_group_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GroupMember struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Member is the principal who belong to this group.
	//
	// Format: users/hello@world.com
	Member        string           `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
	Role          GroupMember_Role `protobuf:"varint,2,opt,name=role,proto3,enum=bytebase.v1.GroupMember_Role" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GroupMember) Reset() {
	*x = GroupMember{}
	mi := &file_v1_group_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMember) ProtoMessage() {}

func (x *GroupMember) ProtoReflect() protoreflect.Message {
	mi := &file_v1_group_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMember.ProtoReflect.Descriptor instead.
func (*GroupMember) Descriptor() ([]byte, []int) {
	return file_v1_group_service_proto_rawDescGZIP(), []int{6}
}

func (x *GroupMember) GetMember() string {
	if x != nil {
		return x.Member
	}
	return ""
}

func (x *GroupMember) GetRole() GroupMember_Role {
	if x != nil {
		return x.Role
	}
	return GroupMember_ROLE_UNSPECIFIED
}

type Group struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the group to retrieve.
	// Format: groups/{group}, group is an email.
	Name        string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Title       string         `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Members     []*GroupMember `protobuf:"bytes,5,rep,name=members,proto3" json:"members,omitempty"`
	// source means where the group comes from. For now we support Entra ID SCIM sync, so the source could be Entra ID.
	Source        string `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Group) Reset() {
	*x = Group{}
	mi := &file_v1_group_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_v1_group_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_v1_group_service_proto_rawDescGZIP(), []int{7}
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Group) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Group) GetMembers() []*GroupMember {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *Group) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

var File_v1_group_service_proto protoreflect.FileDescriptor

const file_v1_group_service_proto_rawDesc = "" +
	"\n" +
	"\x16v1/group_service.proto\x12\vbytebase.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a google/protobuf/field_mask.proto\x1a\x13v1/annotation.proto\"A\n" +
	"\x0fGetGroupRequest\x12.\n" +
	"\x04name\x18\x01 \x01(\tB\x1a\xe0A\x02\xfaA\x14\n" +
	"\x12bytebase.com/GroupR\x04name\"O\n" +
	"\x11ListGroupsRequest\x12\x1b\n" +
	"\tpage_size\x18\x01 \x01(\x05R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x02 \x01(\tR\tpageToken\"h\n" +
	"\x12ListGroupsResponse\x12*\n" +
	"\x06groups\x18\x01 \x03(\v2\x12.bytebase.v1.GroupR\x06groups\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"i\n" +
	"\x12CreateGroupRequest\x12-\n" +
	"\x05group\x18\x01 \x01(\v2\x12.bytebase.v1.GroupB\x03\xe0A\x02R\x05group\x12$\n" +
	"\vgroup_email\x18\x02 \x01(\tB\x03\xe0A\x02R\n" +
	"groupEmail\"\xa5\x01\n" +
	"\x12UpdateGroupRequest\x12-\n" +
	"\x05group\x18\x01 \x01(\v2\x12.bytebase.v1.GroupB\x03\xe0A\x02R\x05group\x12;\n" +
	"\vupdate_mask\x18\x02 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMask\x12#\n" +
	"\rallow_missing\x18\x03 \x01(\bR\fallowMissing\"D\n" +
	"\x12DeleteGroupRequest\x12.\n" +
	"\x04name\x18\x01 \x01(\tB\x1a\xe0A\x02\xfaA\x14\n" +
	"\x12bytebase.com/GroupR\x04name\"\x8d\x01\n" +
	"\vGroupMember\x12\x16\n" +
	"\x06member\x18\x01 \x01(\tR\x06member\x121\n" +
	"\x04role\x18\x02 \x01(\x0e2\x1d.bytebase.v1.GroupMember.RoleR\x04role\"3\n" +
	"\x04Role\x12\x14\n" +
	"\x10ROLE_UNSPECIFIED\x10\x00\x12\t\n" +
	"\x05OWNER\x10\x01\x12\n" +
	"\n" +
	"\x06MEMBER\x10\x02\"\xcd\x01\n" +
	"\x05Group\x12\x17\n" +
	"\x04name\x18\x01 \x01(\tB\x03\xe0A\x03R\x04name\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x122\n" +
	"\amembers\x18\x05 \x03(\v2\x18.bytebase.v1.GroupMemberR\amembers\x12\x16\n" +
	"\x06source\x18\a \x01(\tR\x06source:'\xeaA$\n" +
	"\x12bytebase.com/Group\x12\x0egroups/{group}2\xad\x05\n" +
	"\fGroupService\x12u\n" +
	"\bGetGroup\x12\x1c.bytebase.v1.GetGroupRequest\x1a\x12.bytebase.v1.Group\"7\xdaA\x04name\x8a\xea0\rbb.groups.get\x90\xea0\x01\x82\xd3\xe4\x93\x02\x15\x12\x13/v1/{name=groups/*}\x12z\n" +
	"\n" +
	"ListGroups\x12\x1e.bytebase.v1.ListGroupsRequest\x1a\x1f.bytebase.v1.ListGroupsResponse\"+\xdaA\x00\x8a\xea0\x0ebb.groups.list\x90\xea0\x01\x82\xd3\xe4\x93\x02\f\x12\n" +
	"/v1/groups\x12\x81\x01\n" +
	"\vCreateGroup\x12\x1f.bytebase.v1.CreateGroupRequest\x1a\x12.bytebase.v1.Group\"=\xdaA\x05group\x8a\xea0\x10bb.groups.create\x90\xea0\x01\x98\xea0\x01\x82\xd3\xe4\x93\x02\x13:\x05group\"\n" +
	"/v1/groups\x12\x9c\x01\n" +
	"\vUpdateGroup\x12\x1f.bytebase.v1.UpdateGroupRequest\x1a\x12.bytebase.v1.Group\"X\xdaA\x11group,update_mask\x8a\xea0\x10bb.groups.update\x90\xea0\x02\x98\xea0\x01\x82\xd3\xe4\x93\x02\":\x05group2\x19/v1/{group.name=groups/*}\x12\x86\x01\n" +
	"\vDeleteGroup\x12\x1f.bytebase.v1.DeleteGroupRequest\x1a\x16.google.protobuf.Empty\">\xdaA\x04name\x8a\xea0\x10bb.groups.delete\x90\xea0\x02\x98\xea0\x01\x82\xd3\xe4\x93\x02\x15*\x13/v1/{name=groups/*}B6Z4github.com/bytebase/bytebase/backend/generated-go/v1b\x06proto3"

var (
	file_v1_group_service_proto_rawDescOnce sync.Once
	file_v1_group_service_proto_rawDescData []byte
)

func file_v1_group_service_proto_rawDescGZIP() []byte {
	file_v1_group_service_proto_rawDescOnce.Do(func() {
		file_v1_group_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_group_service_proto_rawDesc), len(file_v1_group_service_proto_rawDesc)))
	})
	return file_v1_group_service_proto_rawDescData
}

var file_v1_group_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_group_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_v1_group_service_proto_goTypes = []any{
	(GroupMember_Role)(0),         // 0: bytebase.v1.GroupMember.Role
	(*GetGroupRequest)(nil),       // 1: bytebase.v1.GetGroupRequest
	(*ListGroupsRequest)(nil),     // 2: bytebase.v1.ListGroupsRequest
	(*ListGroupsResponse)(nil),    // 3: bytebase.v1.ListGroupsResponse
	(*CreateGroupRequest)(nil),    // 4: bytebase.v1.CreateGroupRequest
	(*UpdateGroupRequest)(nil),    // 5: bytebase.v1.UpdateGroupRequest
	(*DeleteGroupRequest)(nil),    // 6: bytebase.v1.DeleteGroupRequest
	(*GroupMember)(nil),           // 7: bytebase.v1.GroupMember
	(*Group)(nil),                 // 8: bytebase.v1.Group
	(*fieldmaskpb.FieldMask)(nil), // 9: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_v1_group_service_proto_depIdxs = []int32{
	8,  // 0: bytebase.v1.ListGroupsResponse.groups:type_name -> bytebase.v1.Group
	8,  // 1: bytebase.v1.CreateGroupRequest.group:type_name -> bytebase.v1.Group
	8,  // 2: bytebase.v1.UpdateGroupRequest.group:type_name -> bytebase.v1.Group
	9,  // 3: bytebase.v1.UpdateGroupRequest.update_mask:type_name -> google.protobuf.FieldMask
	0,  // 4: bytebase.v1.GroupMember.role:type_name -> bytebase.v1.GroupMember.Role
	7,  // 5: bytebase.v1.Group.members:type_name -> bytebase.v1.GroupMember
	1,  // 6: bytebase.v1.GroupService.GetGroup:input_type -> bytebase.v1.GetGroupRequest
	2,  // 7: bytebase.v1.GroupService.ListGroups:input_type -> bytebase.v1.ListGroupsRequest
	4,  // 8: bytebase.v1.GroupService.CreateGroup:input_type -> bytebase.v1.CreateGroupRequest
	5,  // 9: bytebase.v1.GroupService.UpdateGroup:input_type -> bytebase.v1.UpdateGroupRequest
	6,  // 10: bytebase.v1.GroupService.DeleteGroup:input_type -> bytebase.v1.DeleteGroupRequest
	8,  // 11: bytebase.v1.GroupService.GetGroup:output_type -> bytebase.v1.Group
	3,  // 12: bytebase.v1.GroupService.ListGroups:output_type -> bytebase.v1.ListGroupsResponse
	8,  // 13: bytebase.v1.GroupService.CreateGroup:output_type -> bytebase.v1.Group
	8,  // 14: bytebase.v1.GroupService.UpdateGroup:output_type -> bytebase.v1.Group
	10, // 15: bytebase.v1.GroupService.DeleteGroup:output_type -> google.protobuf.Empty
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_v1_group_service_proto_init() }
func file_v1_group_service_proto_init() {
	if File_v1_group_service_proto != nil {
		return
	}
	file_v1_annotation_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_group_service_proto_rawDesc), len(file_v1_group_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_group_service_proto_goTypes,
		DependencyIndexes: file_v1_group_service_proto_depIdxs,
		EnumInfos:         file_v1_group_service_proto_enumTypes,
		MessageInfos:      file_v1_group_service_proto_msgTypes,
	}.Build()
	File_v1_group_service_proto = out.File
	file_v1_group_service_proto_goTypes = nil
	file_v1_group_service_proto_depIdxs = nil
}
