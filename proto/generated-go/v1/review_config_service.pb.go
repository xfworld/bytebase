// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: v1/review_config_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

type ListReviewConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of sql review to return. The service may return fewer than
	// this value.
	// If unspecified, at most 50 sql review will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListSQLReviews` must match
	// the call that provided the page token.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListReviewConfigsRequest) Reset() {
	*x = ListReviewConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_review_config_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReviewConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReviewConfigsRequest) ProtoMessage() {}

func (x *ListReviewConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_review_config_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReviewConfigsRequest.ProtoReflect.Descriptor instead.
func (*ListReviewConfigsRequest) Descriptor() ([]byte, []int) {
	return file_v1_review_config_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListReviewConfigsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListReviewConfigsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListReviewConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The sql review from the specified request.
	ReviewConfigs []*ReviewConfig `protobuf:"bytes,1,rep,name=review_configs,json=reviewConfigs,proto3" json:"review_configs,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListReviewConfigsResponse) Reset() {
	*x = ListReviewConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_review_config_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReviewConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReviewConfigsResponse) ProtoMessage() {}

func (x *ListReviewConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_review_config_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReviewConfigsResponse.ProtoReflect.Descriptor instead.
func (*ListReviewConfigsResponse) Descriptor() ([]byte, []int) {
	return file_v1_review_config_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListReviewConfigsResponse) GetReviewConfigs() []*ReviewConfig {
	if x != nil {
		return x.ReviewConfigs
	}
	return nil
}

func (x *ListReviewConfigsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateReviewConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The sql review to create.
	ReviewConfig *ReviewConfig `protobuf:"bytes,1,opt,name=review_config,json=reviewConfig,proto3" json:"review_config,omitempty"`
}

func (x *CreateReviewConfigRequest) Reset() {
	*x = CreateReviewConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_review_config_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReviewConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReviewConfigRequest) ProtoMessage() {}

func (x *CreateReviewConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_review_config_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReviewConfigRequest.ProtoReflect.Descriptor instead.
func (*CreateReviewConfigRequest) Descriptor() ([]byte, []int) {
	return file_v1_review_config_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateReviewConfigRequest) GetReviewConfig() *ReviewConfig {
	if x != nil {
		return x.ReviewConfig
	}
	return nil
}

type UpdateReviewConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The sql review to update.
	//
	// The name field is used to identify the sql review to update.
	ReviewConfig *ReviewConfig `protobuf:"bytes,1,opt,name=review_config,json=reviewConfig,proto3" json:"review_config,omitempty"`
	// The list of fields to update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateReviewConfigRequest) Reset() {
	*x = UpdateReviewConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_review_config_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReviewConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReviewConfigRequest) ProtoMessage() {}

func (x *UpdateReviewConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_review_config_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReviewConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateReviewConfigRequest) Descriptor() ([]byte, []int) {
	return file_v1_review_config_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateReviewConfigRequest) GetReviewConfig() *ReviewConfig {
	if x != nil {
		return x.ReviewConfig
	}
	return nil
}

func (x *UpdateReviewConfigRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type GetReviewConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the sql review to retrieve.
	// Format: reviewConfigs/{reviewConfig}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetReviewConfigRequest) Reset() {
	*x = GetReviewConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_review_config_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReviewConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReviewConfigRequest) ProtoMessage() {}

func (x *GetReviewConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_review_config_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReviewConfigRequest.ProtoReflect.Descriptor instead.
func (*GetReviewConfigRequest) Descriptor() ([]byte, []int) {
	return file_v1_review_config_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetReviewConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteReviewConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the sql review to delete.
	// Format: reviewConfigs/{reviewConfig}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteReviewConfigRequest) Reset() {
	*x = DeleteReviewConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_review_config_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReviewConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReviewConfigRequest) ProtoMessage() {}

func (x *DeleteReviewConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_review_config_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReviewConfigRequest.ProtoReflect.Descriptor instead.
func (*DeleteReviewConfigRequest) Descriptor() ([]byte, []int) {
	return file_v1_review_config_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteReviewConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ReviewConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the sql review to retrieve.
	// Format: reviewConfigs/{reviewConfig}
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Enabled bool   `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Format: users/hello@world.com
	Creator    string                 `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Rules      []*SQLReviewRule       `protobuf:"bytes,7,rep,name=rules,proto3" json:"rules,omitempty"`
	// resources using the config.
	// Format: {resurce}/{resource id}, for example, environments/test.
	Resources []string `protobuf:"bytes,8,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *ReviewConfig) Reset() {
	*x = ReviewConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_review_config_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewConfig) ProtoMessage() {}

func (x *ReviewConfig) ProtoReflect() protoreflect.Message {
	mi := &file_v1_review_config_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewConfig.ProtoReflect.Descriptor instead.
func (*ReviewConfig) Descriptor() ([]byte, []int) {
	return file_v1_review_config_service_proto_rawDescGZIP(), []int{6}
}

func (x *ReviewConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReviewConfig) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ReviewConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *ReviewConfig) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *ReviewConfig) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ReviewConfig) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ReviewConfig) GetRules() []*SQLReviewRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *ReviewConfig) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_v1_review_config_service_proto protoreflect.FileDescriptor

var file_v1_review_config_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x56, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x85, 0x01, 0x0a, 0x19, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x61, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x44, 0x0a, 0x0d, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x9e, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0d, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x0c, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x50, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22,
	0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x1b, 0x0a, 0x19, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x53, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x22, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x1b, 0x0a, 0x19, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8c, 0x03,
	0x0a, 0x0c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41,
	0x01, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41,
	0x01, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30,
	0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x51, 0x4c, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3a, 0x3c,
	0xea, 0x41, 0x39, 0x0a, 0x19, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x7b, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x7d, 0x32, 0xed, 0x06, 0x0a,
	0x13, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xa3, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x26, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x4a,
	0xda, 0x41, 0x00, 0x8a, 0xea, 0x30, 0x17, 0x62, 0x62, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x90, 0xea,
	0x30, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x0d, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x9d, 0x01, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x12, 0x25, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x39, 0xda, 0x41, 0x00, 0x8a, 0xea, 0x30, 0x15, 0x62, 0x62, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x90, 0xea, 0x30,
	0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x98, 0x01, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23,
	0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x45,
	0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x8a, 0xea, 0x30, 0x14, 0x62, 0x62, 0x2e, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x90,
	0xea, 0x30, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xd3, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x26, 0x2e, 0x62,
	0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x7a, 0xda, 0x41, 0x19, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x8a, 0xea, 0x30,
	0x17, 0x62, 0x62, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x90, 0xea, 0x30, 0x01, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x39, 0x3a, 0x0d, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x32, 0x28, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x9e, 0x01, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x26, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x48, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x8a, 0xea, 0x30, 0x17, 0x62,
	0x62, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2e,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x90, 0xea, 0x30, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c,
	0x2a, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x42, 0x11, 0x5a, 0x0f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_review_config_service_proto_rawDescOnce sync.Once
	file_v1_review_config_service_proto_rawDescData = file_v1_review_config_service_proto_rawDesc
)

func file_v1_review_config_service_proto_rawDescGZIP() []byte {
	file_v1_review_config_service_proto_rawDescOnce.Do(func() {
		file_v1_review_config_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_review_config_service_proto_rawDescData)
	})
	return file_v1_review_config_service_proto_rawDescData
}

var file_v1_review_config_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_review_config_service_proto_goTypes = []any{
	(*ListReviewConfigsRequest)(nil),  // 0: bytebase.v1.ListReviewConfigsRequest
	(*ListReviewConfigsResponse)(nil), // 1: bytebase.v1.ListReviewConfigsResponse
	(*CreateReviewConfigRequest)(nil), // 2: bytebase.v1.CreateReviewConfigRequest
	(*UpdateReviewConfigRequest)(nil), // 3: bytebase.v1.UpdateReviewConfigRequest
	(*GetReviewConfigRequest)(nil),    // 4: bytebase.v1.GetReviewConfigRequest
	(*DeleteReviewConfigRequest)(nil), // 5: bytebase.v1.DeleteReviewConfigRequest
	(*ReviewConfig)(nil),              // 6: bytebase.v1.ReviewConfig
	(*fieldmaskpb.FieldMask)(nil),     // 7: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil),     // 8: google.protobuf.Timestamp
	(*SQLReviewRule)(nil),             // 9: bytebase.v1.SQLReviewRule
	(*emptypb.Empty)(nil),             // 10: google.protobuf.Empty
}
var file_v1_review_config_service_proto_depIdxs = []int32{
	6,  // 0: bytebase.v1.ListReviewConfigsResponse.review_configs:type_name -> bytebase.v1.ReviewConfig
	6,  // 1: bytebase.v1.CreateReviewConfigRequest.review_config:type_name -> bytebase.v1.ReviewConfig
	6,  // 2: bytebase.v1.UpdateReviewConfigRequest.review_config:type_name -> bytebase.v1.ReviewConfig
	7,  // 3: bytebase.v1.UpdateReviewConfigRequest.update_mask:type_name -> google.protobuf.FieldMask
	8,  // 4: bytebase.v1.ReviewConfig.create_time:type_name -> google.protobuf.Timestamp
	8,  // 5: bytebase.v1.ReviewConfig.update_time:type_name -> google.protobuf.Timestamp
	9,  // 6: bytebase.v1.ReviewConfig.rules:type_name -> bytebase.v1.SQLReviewRule
	2,  // 7: bytebase.v1.ReviewConfigService.CreateReviewConfig:input_type -> bytebase.v1.CreateReviewConfigRequest
	0,  // 8: bytebase.v1.ReviewConfigService.ListReviewConfigs:input_type -> bytebase.v1.ListReviewConfigsRequest
	4,  // 9: bytebase.v1.ReviewConfigService.GetReviewConfig:input_type -> bytebase.v1.GetReviewConfigRequest
	3,  // 10: bytebase.v1.ReviewConfigService.UpdateReviewConfig:input_type -> bytebase.v1.UpdateReviewConfigRequest
	5,  // 11: bytebase.v1.ReviewConfigService.DeleteReviewConfig:input_type -> bytebase.v1.DeleteReviewConfigRequest
	6,  // 12: bytebase.v1.ReviewConfigService.CreateReviewConfig:output_type -> bytebase.v1.ReviewConfig
	1,  // 13: bytebase.v1.ReviewConfigService.ListReviewConfigs:output_type -> bytebase.v1.ListReviewConfigsResponse
	6,  // 14: bytebase.v1.ReviewConfigService.GetReviewConfig:output_type -> bytebase.v1.ReviewConfig
	6,  // 15: bytebase.v1.ReviewConfigService.UpdateReviewConfig:output_type -> bytebase.v1.ReviewConfig
	10, // 16: bytebase.v1.ReviewConfigService.DeleteReviewConfig:output_type -> google.protobuf.Empty
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_v1_review_config_service_proto_init() }
func file_v1_review_config_service_proto_init() {
	if File_v1_review_config_service_proto != nil {
		return
	}
	file_v1_annotation_proto_init()
	file_v1_org_policy_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_review_config_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListReviewConfigsRequest); i {
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
		file_v1_review_config_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListReviewConfigsResponse); i {
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
		file_v1_review_config_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateReviewConfigRequest); i {
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
		file_v1_review_config_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateReviewConfigRequest); i {
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
		file_v1_review_config_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetReviewConfigRequest); i {
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
		file_v1_review_config_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteReviewConfigRequest); i {
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
		file_v1_review_config_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ReviewConfig); i {
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
			RawDescriptor: file_v1_review_config_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_review_config_service_proto_goTypes,
		DependencyIndexes: file_v1_review_config_service_proto_depIdxs,
		MessageInfos:      file_v1_review_config_service_proto_msgTypes,
	}.Build()
	File_v1_review_config_service_proto = out.File
	file_v1_review_config_service_proto_rawDesc = nil
	file_v1_review_config_service_proto_goTypes = nil
	file_v1_review_config_service_proto_depIdxs = nil
}
