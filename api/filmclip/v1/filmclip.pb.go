// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.3
// source: filmclip/v1/filmclip.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Intent int32

const (
	// 默认直接调用LLM
	Intent_Default Intent = 0
	// 自动识别意图
	Intent_AutoClassify Intent = 1
	// 生成分镜脚本
	Intent_GenClipScript Intent = 2
)

// Enum value maps for Intent.
var (
	Intent_name = map[int32]string{
		0: "Default",
		1: "AutoClassify",
		2: "GenClipScript",
	}
	Intent_value = map[string]int32{
		"Default":       0,
		"AutoClassify":  1,
		"GenClipScript": 2,
	}
)

func (x Intent) Enum() *Intent {
	p := new(Intent)
	*p = x
	return p
}

func (x Intent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Intent) Descriptor() protoreflect.EnumDescriptor {
	return file_filmclip_v1_filmclip_proto_enumTypes[0].Descriptor()
}

func (Intent) Type() protoreflect.EnumType {
	return &file_filmclip_v1_filmclip_proto_enumTypes[0]
}

func (x Intent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Intent.Descriptor instead.
func (Intent) EnumDescriptor() ([]byte, []int) {
	return file_filmclip_v1_filmclip_proto_rawDescGZIP(), []int{0}
}

type UploadImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Base64 string `protobuf:"bytes,2,opt,name=base64,proto3" json:"base64,omitempty"`
}

func (x *UploadImageRequest) Reset() {
	*x = UploadImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filmclip_v1_filmclip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageRequest) ProtoMessage() {}

func (x *UploadImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filmclip_v1_filmclip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageRequest.ProtoReflect.Descriptor instead.
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return file_filmclip_v1_filmclip_proto_rawDescGZIP(), []int{0}
}

func (x *UploadImageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UploadImageRequest) GetBase64() string {
	if x != nil {
		return x.Base64
	}
	return ""
}

type UploadImageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UploadImageReply) Reset() {
	*x = UploadImageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filmclip_v1_filmclip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageReply) ProtoMessage() {}

func (x *UploadImageReply) ProtoReflect() protoreflect.Message {
	mi := &file_filmclip_v1_filmclip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageReply.ProtoReflect.Descriptor instead.
func (*UploadImageReply) Descriptor() ([]byte, []int) {
	return file_filmclip_v1_filmclip_proto_rawDescGZIP(), []int{1}
}

func (x *UploadImageReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GenClipScriptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClothingImage string `protobuf:"bytes,1,opt,name=clothing_image,json=clothingImage,proto3" json:"clothing_image,omitempty"`
	ModelImage    string `protobuf:"bytes,2,opt,name=model_image,json=modelImage,proto3" json:"model_image,omitempty"`
	Prompt        string `protobuf:"bytes,3,opt,name=prompt,proto3" json:"prompt,omitempty"`
}

func (x *GenClipScriptRequest) Reset() {
	*x = GenClipScriptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filmclip_v1_filmclip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenClipScriptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenClipScriptRequest) ProtoMessage() {}

func (x *GenClipScriptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filmclip_v1_filmclip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenClipScriptRequest.ProtoReflect.Descriptor instead.
func (*GenClipScriptRequest) Descriptor() ([]byte, []int) {
	return file_filmclip_v1_filmclip_proto_rawDescGZIP(), []int{2}
}

func (x *GenClipScriptRequest) GetClothingImage() string {
	if x != nil {
		return x.ClothingImage
	}
	return ""
}

func (x *GenClipScriptRequest) GetModelImage() string {
	if x != nil {
		return x.ModelImage
	}
	return ""
}

func (x *GenClipScriptRequest) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

type GenClipScriptReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string         `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string         `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Scenes  []*SceneScript `protobuf:"bytes,3,rep,name=scenes,proto3" json:"scenes,omitempty"`
}

func (x *GenClipScriptReply) Reset() {
	*x = GenClipScriptReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filmclip_v1_filmclip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenClipScriptReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenClipScriptReply) ProtoMessage() {}

func (x *GenClipScriptReply) ProtoReflect() protoreflect.Message {
	mi := &file_filmclip_v1_filmclip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenClipScriptReply.ProtoReflect.Descriptor instead.
func (*GenClipScriptReply) Descriptor() ([]byte, []int) {
	return file_filmclip_v1_filmclip_proto_rawDescGZIP(), []int{3}
}

func (x *GenClipScriptReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GenClipScriptReply) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *GenClipScriptReply) GetScenes() []*SceneScript {
	if x != nil {
		return x.Scenes
	}
	return nil
}

type SceneScript struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Actions     string `protobuf:"bytes,2,opt,name=actions,proto3" json:"actions,omitempty"`
	ShotType    string `protobuf:"bytes,3,opt,name=shotType,proto3" json:"shotType,omitempty"`
}

func (x *SceneScript) Reset() {
	*x = SceneScript{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filmclip_v1_filmclip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneScript) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneScript) ProtoMessage() {}

func (x *SceneScript) ProtoReflect() protoreflect.Message {
	mi := &file_filmclip_v1_filmclip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneScript.ProtoReflect.Descriptor instead.
func (*SceneScript) Descriptor() ([]byte, []int) {
	return file_filmclip_v1_filmclip_proto_rawDescGZIP(), []int{4}
}

func (x *SceneScript) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SceneScript) GetActions() string {
	if x != nil {
		return x.Actions
	}
	return ""
}

func (x *SceneScript) GetShotType() string {
	if x != nil {
		return x.ShotType
	}
	return ""
}

var File_filmclip_v1_filmclip_proto protoreflect.FileDescriptor

var file_filmclip_v1_filmclip_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x66, 0x69, 0x6c, 0x6d, 0x63, 0x6c, 0x69, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69,
	0x6c, 0x6d, 0x63, 0x6c, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66, 0x69,
	0x6c, 0x6d, 0x63, 0x6c, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x52, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x62, 0x61,
	0x73, 0x65, 0x36, 0x34, 0x22, 0x24, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x88, 0x01, 0x0a, 0x14, 0x47,
	0x65, 0x6e, 0x43, 0x6c, 0x69, 0x70, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0e, 0x63, 0x6c, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x63, 0x6c, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x22, 0x76, 0x0a, 0x12, 0x47, 0x65, 0x6e, 0x43, 0x6c, 0x69, 0x70,
	0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x66, 0x69,
	0x6c, 0x6d, 0x63, 0x6c, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x53,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x06, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x73, 0x22, 0x65, 0x0a,
	0x0b, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x2a, 0x3a, 0x0a, 0x06, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0b,
	0x0a, 0x07, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41,
	0x75, 0x74, 0x6f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x10, 0x01, 0x12, 0x11, 0x0a,
	0x0d, 0x47, 0x65, 0x6e, 0x43, 0x6c, 0x69, 0x70, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x10, 0x02,
	0x32, 0xf0, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x6d, 0x63, 0x6c, 0x69, 0x70, 0x12, 0x6d, 0x0a,
	0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x66,
	0x69, 0x6c, 0x6d, 0x63, 0x6c, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x66, 0x69, 0x6c, 0x6d, 0x63, 0x6c, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x75, 0x0a, 0x0d,
	0x47, 0x65, 0x6e, 0x43, 0x6c, 0x69, 0x70, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x21, 0x2e,
	0x66, 0x69, 0x6c, 0x6d, 0x63, 0x6c, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x43,
	0x6c, 0x69, 0x70, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x63, 0x6c, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x6e, 0x43, 0x6c, 0x69, 0x70, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x63,
	0x6c, 0x69, 0x70, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x42, 0x4d, 0x0a, 0x1a, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x63, 0x6c, 0x69, 0x70, 0x2e, 0x76,
	0x31, 0x42, 0x0f, 0x46, 0x69, 0x6c, 0x6d, 0x63, 0x6c, 0x69, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x56, 0x31, 0x50, 0x01, 0x5a, 0x1c, 0x61, 0x69, 0x2d, 0x6d, 0x6b, 0x74, 0x2d, 0x62, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c, 0x6d, 0x63, 0x6c, 0x69, 0x70, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filmclip_v1_filmclip_proto_rawDescOnce sync.Once
	file_filmclip_v1_filmclip_proto_rawDescData = file_filmclip_v1_filmclip_proto_rawDesc
)

func file_filmclip_v1_filmclip_proto_rawDescGZIP() []byte {
	file_filmclip_v1_filmclip_proto_rawDescOnce.Do(func() {
		file_filmclip_v1_filmclip_proto_rawDescData = protoimpl.X.CompressGZIP(file_filmclip_v1_filmclip_proto_rawDescData)
	})
	return file_filmclip_v1_filmclip_proto_rawDescData
}

var file_filmclip_v1_filmclip_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_filmclip_v1_filmclip_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_filmclip_v1_filmclip_proto_goTypes = []any{
	(Intent)(0),                  // 0: filmclip.v1.Intent
	(*UploadImageRequest)(nil),   // 1: filmclip.v1.UploadImageRequest
	(*UploadImageReply)(nil),     // 2: filmclip.v1.UploadImageReply
	(*GenClipScriptRequest)(nil), // 3: filmclip.v1.GenClipScriptRequest
	(*GenClipScriptReply)(nil),   // 4: filmclip.v1.GenClipScriptReply
	(*SceneScript)(nil),          // 5: filmclip.v1.SceneScript
}
var file_filmclip_v1_filmclip_proto_depIdxs = []int32{
	5, // 0: filmclip.v1.GenClipScriptReply.scenes:type_name -> filmclip.v1.SceneScript
	1, // 1: filmclip.v1.Filmclip.UploadImage:input_type -> filmclip.v1.UploadImageRequest
	3, // 2: filmclip.v1.Filmclip.GenClipScript:input_type -> filmclip.v1.GenClipScriptRequest
	2, // 3: filmclip.v1.Filmclip.UploadImage:output_type -> filmclip.v1.UploadImageReply
	4, // 4: filmclip.v1.Filmclip.GenClipScript:output_type -> filmclip.v1.GenClipScriptReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_filmclip_v1_filmclip_proto_init() }
func file_filmclip_v1_filmclip_proto_init() {
	if File_filmclip_v1_filmclip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filmclip_v1_filmclip_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UploadImageRequest); i {
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
		file_filmclip_v1_filmclip_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UploadImageReply); i {
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
		file_filmclip_v1_filmclip_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GenClipScriptRequest); i {
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
		file_filmclip_v1_filmclip_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GenClipScriptReply); i {
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
		file_filmclip_v1_filmclip_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SceneScript); i {
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
			RawDescriptor: file_filmclip_v1_filmclip_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_filmclip_v1_filmclip_proto_goTypes,
		DependencyIndexes: file_filmclip_v1_filmclip_proto_depIdxs,
		EnumInfos:         file_filmclip_v1_filmclip_proto_enumTypes,
		MessageInfos:      file_filmclip_v1_filmclip_proto_msgTypes,
	}.Build()
	File_filmclip_v1_filmclip_proto = out.File
	file_filmclip_v1_filmclip_proto_rawDesc = nil
	file_filmclip_v1_filmclip_proto_goTypes = nil
	file_filmclip_v1_filmclip_proto_depIdxs = nil
}
