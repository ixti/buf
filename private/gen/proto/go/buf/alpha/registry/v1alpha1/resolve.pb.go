// Copyright 2020-2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: buf/alpha/registry/v1alpha1/resolve.proto

package registryv1alpha1

import (
	v1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/module/v1alpha1"
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

type ResolvedReferenceType int32

const (
	ResolvedReferenceType_RESOLVED_REFERENCE_TYPE_UNSPECIFIED ResolvedReferenceType = 0
	ResolvedReferenceType_RESOLVED_REFERENCE_TYPE_COMMIT      ResolvedReferenceType = 1
	ResolvedReferenceType_RESOLVED_REFERENCE_TYPE_TAG         ResolvedReferenceType = 3
	ResolvedReferenceType_RESOLVED_REFERENCE_TYPE_DRAFT       ResolvedReferenceType = 5
)

// Enum value maps for ResolvedReferenceType.
var (
	ResolvedReferenceType_name = map[int32]string{
		0: "RESOLVED_REFERENCE_TYPE_UNSPECIFIED",
		1: "RESOLVED_REFERENCE_TYPE_COMMIT",
		3: "RESOLVED_REFERENCE_TYPE_TAG",
		5: "RESOLVED_REFERENCE_TYPE_DRAFT",
	}
	ResolvedReferenceType_value = map[string]int32{
		"RESOLVED_REFERENCE_TYPE_UNSPECIFIED": 0,
		"RESOLVED_REFERENCE_TYPE_COMMIT":      1,
		"RESOLVED_REFERENCE_TYPE_TAG":         3,
		"RESOLVED_REFERENCE_TYPE_DRAFT":       5,
	}
)

func (x ResolvedReferenceType) Enum() *ResolvedReferenceType {
	p := new(ResolvedReferenceType)
	*p = x
	return p
}

func (x ResolvedReferenceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResolvedReferenceType) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_alpha_registry_v1alpha1_resolve_proto_enumTypes[0].Descriptor()
}

func (ResolvedReferenceType) Type() protoreflect.EnumType {
	return &file_buf_alpha_registry_v1alpha1_resolve_proto_enumTypes[0]
}

func (x ResolvedReferenceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResolvedReferenceType.Descriptor instead.
func (ResolvedReferenceType) EnumDescriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_resolve_proto_rawDescGZIP(), []int{0}
}

type GetModulePinsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleReferences []*v1alpha1.ModuleReference `protobuf:"bytes,1,rep,name=module_references,json=moduleReferences,proto3" json:"module_references,omitempty"`
	// current_module_pins allows for partial dependency updates by letting clients
	// send a request with the pins for their current module and only the
	// identities of the dependencies they want to update in module_references.
	//
	// When resolving, if a client supplied module pin is:
	//   - in the transitive closure of pins resolved from the module_references,
	//     the client supplied module pin will be an extra candidate for tie
	//     breaking.
	//   - NOT in the in the transitive closure of pins resolved from the
	//     module_references, it will be returned as is.
	CurrentModulePins []*v1alpha1.ModulePin `protobuf:"bytes,2,rep,name=current_module_pins,json=currentModulePins,proto3" json:"current_module_pins,omitempty"`
}

func (x *GetModulePinsRequest) Reset() {
	*x = GetModulePinsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModulePinsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModulePinsRequest) ProtoMessage() {}

func (x *GetModulePinsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModulePinsRequest.ProtoReflect.Descriptor instead.
func (*GetModulePinsRequest) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_resolve_proto_rawDescGZIP(), []int{0}
}

func (x *GetModulePinsRequest) GetModuleReferences() []*v1alpha1.ModuleReference {
	if x != nil {
		return x.ModuleReferences
	}
	return nil
}

func (x *GetModulePinsRequest) GetCurrentModulePins() []*v1alpha1.ModulePin {
	if x != nil {
		return x.CurrentModulePins
	}
	return nil
}

type GetModulePinsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModulePins []*v1alpha1.ModulePin `protobuf:"bytes,1,rep,name=module_pins,json=modulePins,proto3" json:"module_pins,omitempty"`
}

func (x *GetModulePinsResponse) Reset() {
	*x = GetModulePinsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModulePinsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModulePinsResponse) ProtoMessage() {}

func (x *GetModulePinsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModulePinsResponse.ProtoReflect.Descriptor instead.
func (*GetModulePinsResponse) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_resolve_proto_rawDescGZIP(), []int{1}
}

func (x *GetModulePinsResponse) GetModulePins() []*v1alpha1.ModulePin {
	if x != nil {
		return x.ModulePins
	}
	return nil
}

type GetLocalModulePinsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalModuleReferences []*LocalModuleReference `protobuf:"bytes,1,rep,name=local_module_references,json=localModuleReferences,proto3" json:"local_module_references,omitempty"`
}

func (x *GetLocalModulePinsRequest) Reset() {
	*x = GetLocalModulePinsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocalModulePinsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocalModulePinsRequest) ProtoMessage() {}

func (x *GetLocalModulePinsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocalModulePinsRequest.ProtoReflect.Descriptor instead.
func (*GetLocalModulePinsRequest) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_resolve_proto_rawDescGZIP(), []int{2}
}

func (x *GetLocalModulePinsRequest) GetLocalModuleReferences() []*LocalModuleReference {
	if x != nil {
		return x.LocalModuleReferences
	}
	return nil
}

type LocalModuleResolveResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A copy of the reference that was resolved.
	Reference *LocalModuleReference `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`
	// The pin the reference resolved to.
	Pin *LocalModulePin `protobuf:"bytes,2,opt,name=pin,proto3" json:"pin,omitempty"`
	// The type the reference resolved as.
	ResolvedReferenceType ResolvedReferenceType `protobuf:"varint,3,opt,name=resolved_reference_type,json=resolvedReferenceType,proto3,enum=buf.alpha.registry.v1alpha1.ResolvedReferenceType" json:"resolved_reference_type,omitempty"`
}

func (x *LocalModuleResolveResult) Reset() {
	*x = LocalModuleResolveResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalModuleResolveResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalModuleResolveResult) ProtoMessage() {}

func (x *LocalModuleResolveResult) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalModuleResolveResult.ProtoReflect.Descriptor instead.
func (*LocalModuleResolveResult) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_resolve_proto_rawDescGZIP(), []int{3}
}

func (x *LocalModuleResolveResult) GetReference() *LocalModuleReference {
	if x != nil {
		return x.Reference
	}
	return nil
}

func (x *LocalModuleResolveResult) GetPin() *LocalModulePin {
	if x != nil {
		return x.Pin
	}
	return nil
}

func (x *LocalModuleResolveResult) GetResolvedReferenceType() ResolvedReferenceType {
	if x != nil {
		return x.ResolvedReferenceType
	}
	return ResolvedReferenceType_RESOLVED_REFERENCE_TYPE_UNSPECIFIED
}

type GetLocalModulePinsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalModuleResolveResults []*LocalModuleResolveResult `protobuf:"bytes,1,rep,name=local_module_resolve_results,json=localModuleResolveResults,proto3" json:"local_module_resolve_results,omitempty"`
	// dependencies are the dependencies of the LocalModulePins.
	//
	// This includes the transitive deps.
	Dependencies []*v1alpha1.ModulePin `protobuf:"bytes,2,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
}

func (x *GetLocalModulePinsResponse) Reset() {
	*x = GetLocalModulePinsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocalModulePinsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocalModulePinsResponse) ProtoMessage() {}

func (x *GetLocalModulePinsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocalModulePinsResponse.ProtoReflect.Descriptor instead.
func (*GetLocalModulePinsResponse) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_resolve_proto_rawDescGZIP(), []int{4}
}

func (x *GetLocalModulePinsResponse) GetLocalModuleResolveResults() []*LocalModuleResolveResult {
	if x != nil {
		return x.LocalModuleResolveResults
	}
	return nil
}

func (x *GetLocalModulePinsResponse) GetDependencies() []*v1alpha1.ModulePin {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

var File_buf_alpha_registry_v1alpha1_resolve_proto protoreflect.FileDescriptor

var file_buf_alpha_registry_v1alpha1_resolve_proto_rawDesc = []byte{
	0x0a, 0x29, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x62, 0x75, 0x66,
	0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x26, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x28, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x57, 0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x10, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x13,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x70,
	0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x52,
	0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x69,
	0x6e, 0x73, 0x22, 0x5e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50,
	0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x69,
	0x6e, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x69, 0x0a, 0x17, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x15, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x96, 0x02, 0x0a, 0x18,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x4f, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x09,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x03, 0x70, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x50, 0x69, 0x6e, 0x52, 0x03, 0x70, 0x69, 0x6e, 0x12, 0x6a, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x15, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x22, 0xde, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x76, 0x0a, 0x1c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x19, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x48, 0x0a, 0x0c, 0x64,
	0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x2a, 0xf3, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x64, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x27, 0x0a, 0x23, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x46, 0x45,
	0x52, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x52, 0x45, 0x53, 0x4f,
	0x4c, 0x56, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b,
	0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e,
	0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x41, 0x47, 0x10, 0x03, 0x12, 0x21, 0x0a,
	0x1d, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45,
	0x4e, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x05,
	0x22, 0x04, 0x08, 0x02, 0x10, 0x02, 0x22, 0x04, 0x08, 0x04, 0x10, 0x04, 0x2a, 0x1e, 0x52, 0x45,
	0x53, 0x4f, 0x4c, 0x56, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x52, 0x41, 0x4e, 0x43, 0x48, 0x2a, 0x1d, 0x52, 0x45,
	0x53, 0x4f, 0x4c, 0x56, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x32, 0x88, 0x01, 0x0a, 0x0e,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x76,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x73, 0x12,
	0x31, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x32, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9d, 0x01, 0x0a, 0x13, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x85,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x50, 0x69, 0x6e, 0x73, 0x12, 0x36, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x50, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e,
	0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x99, 0x02, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f,
	0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x3b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x41, 0x52, 0xaa, 0x02, 0x1b, 0x42, 0x75,
	0x66, 0x2e, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1b, 0x42, 0x75, 0x66, 0x5c,
	0x41, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x27, 0x42, 0x75, 0x66, 0x5c, 0x41, 0x6c,
	0x70, 0x68, 0x61, 0x5c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x3a,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_alpha_registry_v1alpha1_resolve_proto_rawDescOnce sync.Once
	file_buf_alpha_registry_v1alpha1_resolve_proto_rawDescData = file_buf_alpha_registry_v1alpha1_resolve_proto_rawDesc
)

func file_buf_alpha_registry_v1alpha1_resolve_proto_rawDescGZIP() []byte {
	file_buf_alpha_registry_v1alpha1_resolve_proto_rawDescOnce.Do(func() {
		file_buf_alpha_registry_v1alpha1_resolve_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_alpha_registry_v1alpha1_resolve_proto_rawDescData)
	})
	return file_buf_alpha_registry_v1alpha1_resolve_proto_rawDescData
}

var file_buf_alpha_registry_v1alpha1_resolve_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_buf_alpha_registry_v1alpha1_resolve_proto_goTypes = []interface{}{
	(ResolvedReferenceType)(0),         // 0: buf.alpha.registry.v1alpha1.ResolvedReferenceType
	(*GetModulePinsRequest)(nil),       // 1: buf.alpha.registry.v1alpha1.GetModulePinsRequest
	(*GetModulePinsResponse)(nil),      // 2: buf.alpha.registry.v1alpha1.GetModulePinsResponse
	(*GetLocalModulePinsRequest)(nil),  // 3: buf.alpha.registry.v1alpha1.GetLocalModulePinsRequest
	(*LocalModuleResolveResult)(nil),   // 4: buf.alpha.registry.v1alpha1.LocalModuleResolveResult
	(*GetLocalModulePinsResponse)(nil), // 5: buf.alpha.registry.v1alpha1.GetLocalModulePinsResponse
	(*v1alpha1.ModuleReference)(nil),   // 6: buf.alpha.module.v1alpha1.ModuleReference
	(*v1alpha1.ModulePin)(nil),         // 7: buf.alpha.module.v1alpha1.ModulePin
	(*LocalModuleReference)(nil),       // 8: buf.alpha.registry.v1alpha1.LocalModuleReference
	(*LocalModulePin)(nil),             // 9: buf.alpha.registry.v1alpha1.LocalModulePin
}
var file_buf_alpha_registry_v1alpha1_resolve_proto_depIdxs = []int32{
	6,  // 0: buf.alpha.registry.v1alpha1.GetModulePinsRequest.module_references:type_name -> buf.alpha.module.v1alpha1.ModuleReference
	7,  // 1: buf.alpha.registry.v1alpha1.GetModulePinsRequest.current_module_pins:type_name -> buf.alpha.module.v1alpha1.ModulePin
	7,  // 2: buf.alpha.registry.v1alpha1.GetModulePinsResponse.module_pins:type_name -> buf.alpha.module.v1alpha1.ModulePin
	8,  // 3: buf.alpha.registry.v1alpha1.GetLocalModulePinsRequest.local_module_references:type_name -> buf.alpha.registry.v1alpha1.LocalModuleReference
	8,  // 4: buf.alpha.registry.v1alpha1.LocalModuleResolveResult.reference:type_name -> buf.alpha.registry.v1alpha1.LocalModuleReference
	9,  // 5: buf.alpha.registry.v1alpha1.LocalModuleResolveResult.pin:type_name -> buf.alpha.registry.v1alpha1.LocalModulePin
	0,  // 6: buf.alpha.registry.v1alpha1.LocalModuleResolveResult.resolved_reference_type:type_name -> buf.alpha.registry.v1alpha1.ResolvedReferenceType
	4,  // 7: buf.alpha.registry.v1alpha1.GetLocalModulePinsResponse.local_module_resolve_results:type_name -> buf.alpha.registry.v1alpha1.LocalModuleResolveResult
	7,  // 8: buf.alpha.registry.v1alpha1.GetLocalModulePinsResponse.dependencies:type_name -> buf.alpha.module.v1alpha1.ModulePin
	1,  // 9: buf.alpha.registry.v1alpha1.ResolveService.GetModulePins:input_type -> buf.alpha.registry.v1alpha1.GetModulePinsRequest
	3,  // 10: buf.alpha.registry.v1alpha1.LocalResolveService.GetLocalModulePins:input_type -> buf.alpha.registry.v1alpha1.GetLocalModulePinsRequest
	2,  // 11: buf.alpha.registry.v1alpha1.ResolveService.GetModulePins:output_type -> buf.alpha.registry.v1alpha1.GetModulePinsResponse
	5,  // 12: buf.alpha.registry.v1alpha1.LocalResolveService.GetLocalModulePins:output_type -> buf.alpha.registry.v1alpha1.GetLocalModulePinsResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_buf_alpha_registry_v1alpha1_resolve_proto_init() }
func file_buf_alpha_registry_v1alpha1_resolve_proto_init() {
	if File_buf_alpha_registry_v1alpha1_resolve_proto != nil {
		return
	}
	file_buf_alpha_registry_v1alpha1_module_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModulePinsRequest); i {
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
		file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModulePinsResponse); i {
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
		file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLocalModulePinsRequest); i {
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
		file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalModuleResolveResult); i {
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
		file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLocalModulePinsResponse); i {
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
			RawDescriptor: file_buf_alpha_registry_v1alpha1_resolve_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_buf_alpha_registry_v1alpha1_resolve_proto_goTypes,
		DependencyIndexes: file_buf_alpha_registry_v1alpha1_resolve_proto_depIdxs,
		EnumInfos:         file_buf_alpha_registry_v1alpha1_resolve_proto_enumTypes,
		MessageInfos:      file_buf_alpha_registry_v1alpha1_resolve_proto_msgTypes,
	}.Build()
	File_buf_alpha_registry_v1alpha1_resolve_proto = out.File
	file_buf_alpha_registry_v1alpha1_resolve_proto_rawDesc = nil
	file_buf_alpha_registry_v1alpha1_resolve_proto_goTypes = nil
	file_buf_alpha_registry_v1alpha1_resolve_proto_depIdxs = nil
}
