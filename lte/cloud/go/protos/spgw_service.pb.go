//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.10.0
// source: lte/protos/spgw_service.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateBearerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid    *SubscriberID `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	IpAddr string        `protobuf:"bytes,2,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	// default bearer ID
	LinkBearerId uint32        `protobuf:"varint,3,opt,name=link_bearer_id,json=linkBearerId,proto3" json:"link_bearer_id,omitempty"`
	PolicyRules  []*PolicyRule `protobuf:"bytes,4,rep,name=policy_rules,json=policyRules,proto3" json:"policy_rules,omitempty"`
}

func (x *CreateBearerRequest) Reset() {
	*x = CreateBearerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_spgw_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBearerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBearerRequest) ProtoMessage() {}

func (x *CreateBearerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_spgw_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBearerRequest.ProtoReflect.Descriptor instead.
func (*CreateBearerRequest) Descriptor() ([]byte, []int) {
	return file_lte_protos_spgw_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBearerRequest) GetSid() *SubscriberID {
	if x != nil {
		return x.Sid
	}
	return nil
}

func (x *CreateBearerRequest) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

func (x *CreateBearerRequest) GetLinkBearerId() uint32 {
	if x != nil {
		return x.LinkBearerId
	}
	return 0
}

func (x *CreateBearerRequest) GetPolicyRules() []*PolicyRule {
	if x != nil {
		return x.PolicyRules
	}
	return nil
}

type CreateBearerResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateBearerResult) Reset() {
	*x = CreateBearerResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_spgw_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBearerResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBearerResult) ProtoMessage() {}

func (x *CreateBearerResult) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_spgw_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBearerResult.ProtoReflect.Descriptor instead.
func (*CreateBearerResult) Descriptor() ([]byte, []int) {
	return file_lte_protos_spgw_service_proto_rawDescGZIP(), []int{1}
}

type DeleteBearerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid    *SubscriberID `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	IpAddr string        `protobuf:"bytes,2,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	// default bearer ID
	LinkBearerId uint32 `protobuf:"varint,3,opt,name=link_bearer_id,json=linkBearerId,proto3" json:"link_bearer_id,omitempty"`
	// dedicated bearer IDs
	EpsBearerIds []uint32 `protobuf:"varint,4,rep,packed,name=eps_bearer_ids,json=epsBearerIds,proto3" json:"eps_bearer_ids,omitempty"`
}

func (x *DeleteBearerRequest) Reset() {
	*x = DeleteBearerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_spgw_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBearerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBearerRequest) ProtoMessage() {}

func (x *DeleteBearerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_spgw_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBearerRequest.ProtoReflect.Descriptor instead.
func (*DeleteBearerRequest) Descriptor() ([]byte, []int) {
	return file_lte_protos_spgw_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteBearerRequest) GetSid() *SubscriberID {
	if x != nil {
		return x.Sid
	}
	return nil
}

func (x *DeleteBearerRequest) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

func (x *DeleteBearerRequest) GetLinkBearerId() uint32 {
	if x != nil {
		return x.LinkBearerId
	}
	return 0
}

func (x *DeleteBearerRequest) GetEpsBearerIds() []uint32 {
	if x != nil {
		return x.EpsBearerIds
	}
	return nil
}

type DeleteBearerResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBearerResult) Reset() {
	*x = DeleteBearerResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_spgw_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBearerResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBearerResult) ProtoMessage() {}

func (x *DeleteBearerResult) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_spgw_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBearerResult.ProtoReflect.Descriptor instead.
func (*DeleteBearerResult) Descriptor() ([]byte, []int) {
	return file_lte_protos_spgw_service_proto_rawDescGZIP(), []int{3}
}

var File_lte_protos_spgw_service_proto protoreflect.FileDescriptor

var file_lte_protos_spgw_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6c, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x70, 0x67,
	0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x1a, 0x1d, 0x6c, 0x74, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6c, 0x74, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x64, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x03,
	0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x49, 0x44, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x70, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72,
	0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6c, 0x69, 0x6e, 0x6b, 0x42, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x73,
	0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xa5, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x61,
	0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x49, 0x44, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x70, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x70, 0x41, 0x64,
	0x64, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x62, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6c, 0x69, 0x6e, 0x6b,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x70, 0x73, 0x5f,
	0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0c, 0x65, 0x70, 0x73, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x14,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x32, 0xaf, 0x01, 0x0a, 0x0b, 0x53, 0x70, 0x67, 0x77, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f,
	0x6c, 0x74, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lte_protos_spgw_service_proto_rawDescOnce sync.Once
	file_lte_protos_spgw_service_proto_rawDescData = file_lte_protos_spgw_service_proto_rawDesc
)

func file_lte_protos_spgw_service_proto_rawDescGZIP() []byte {
	file_lte_protos_spgw_service_proto_rawDescOnce.Do(func() {
		file_lte_protos_spgw_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_lte_protos_spgw_service_proto_rawDescData)
	})
	return file_lte_protos_spgw_service_proto_rawDescData
}

var file_lte_protos_spgw_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_lte_protos_spgw_service_proto_goTypes = []interface{}{
	(*CreateBearerRequest)(nil), // 0: magma.lte.CreateBearerRequest
	(*CreateBearerResult)(nil),  // 1: magma.lte.CreateBearerResult
	(*DeleteBearerRequest)(nil), // 2: magma.lte.DeleteBearerRequest
	(*DeleteBearerResult)(nil),  // 3: magma.lte.DeleteBearerResult
	(*SubscriberID)(nil),        // 4: magma.lte.SubscriberID
	(*PolicyRule)(nil),          // 5: magma.lte.PolicyRule
}
var file_lte_protos_spgw_service_proto_depIdxs = []int32{
	4, // 0: magma.lte.CreateBearerRequest.sid:type_name -> magma.lte.SubscriberID
	5, // 1: magma.lte.CreateBearerRequest.policy_rules:type_name -> magma.lte.PolicyRule
	4, // 2: magma.lte.DeleteBearerRequest.sid:type_name -> magma.lte.SubscriberID
	0, // 3: magma.lte.SpgwService.CreateBearer:input_type -> magma.lte.CreateBearerRequest
	2, // 4: magma.lte.SpgwService.DeleteBearer:input_type -> magma.lte.DeleteBearerRequest
	1, // 5: magma.lte.SpgwService.CreateBearer:output_type -> magma.lte.CreateBearerResult
	3, // 6: magma.lte.SpgwService.DeleteBearer:output_type -> magma.lte.DeleteBearerResult
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_lte_protos_spgw_service_proto_init() }
func file_lte_protos_spgw_service_proto_init() {
	if File_lte_protos_spgw_service_proto != nil {
		return
	}
	file_lte_protos_subscriberdb_proto_init()
	file_lte_protos_policydb_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lte_protos_spgw_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBearerRequest); i {
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
		file_lte_protos_spgw_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBearerResult); i {
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
		file_lte_protos_spgw_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBearerRequest); i {
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
		file_lte_protos_spgw_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBearerResult); i {
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
			RawDescriptor: file_lte_protos_spgw_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lte_protos_spgw_service_proto_goTypes,
		DependencyIndexes: file_lte_protos_spgw_service_proto_depIdxs,
		MessageInfos:      file_lte_protos_spgw_service_proto_msgTypes,
	}.Build()
	File_lte_protos_spgw_service_proto = out.File
	file_lte_protos_spgw_service_proto_rawDesc = nil
	file_lte_protos_spgw_service_proto_goTypes = nil
	file_lte_protos_spgw_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SpgwServiceClient is the client API for SpgwService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpgwServiceClient interface {
	CreateBearer(ctx context.Context, in *CreateBearerRequest, opts ...grpc.CallOption) (*CreateBearerResult, error)
	DeleteBearer(ctx context.Context, in *DeleteBearerRequest, opts ...grpc.CallOption) (*DeleteBearerResult, error)
}

type spgwServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpgwServiceClient(cc grpc.ClientConnInterface) SpgwServiceClient {
	return &spgwServiceClient{cc}
}

func (c *spgwServiceClient) CreateBearer(ctx context.Context, in *CreateBearerRequest, opts ...grpc.CallOption) (*CreateBearerResult, error) {
	out := new(CreateBearerResult)
	err := c.cc.Invoke(ctx, "/magma.lte.SpgwService/CreateBearer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spgwServiceClient) DeleteBearer(ctx context.Context, in *DeleteBearerRequest, opts ...grpc.CallOption) (*DeleteBearerResult, error) {
	out := new(DeleteBearerResult)
	err := c.cc.Invoke(ctx, "/magma.lte.SpgwService/DeleteBearer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpgwServiceServer is the server API for SpgwService service.
type SpgwServiceServer interface {
	CreateBearer(context.Context, *CreateBearerRequest) (*CreateBearerResult, error)
	DeleteBearer(context.Context, *DeleteBearerRequest) (*DeleteBearerResult, error)
}

// UnimplementedSpgwServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSpgwServiceServer struct {
}

func (*UnimplementedSpgwServiceServer) CreateBearer(context.Context, *CreateBearerRequest) (*CreateBearerResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBearer not implemented")
}
func (*UnimplementedSpgwServiceServer) DeleteBearer(context.Context, *DeleteBearerRequest) (*DeleteBearerResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBearer not implemented")
}

func RegisterSpgwServiceServer(s *grpc.Server, srv SpgwServiceServer) {
	s.RegisterService(&_SpgwService_serviceDesc, srv)
}

func _SpgwService_CreateBearer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBearerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpgwServiceServer).CreateBearer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.SpgwService/CreateBearer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpgwServiceServer).CreateBearer(ctx, req.(*CreateBearerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpgwService_DeleteBearer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBearerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpgwServiceServer).DeleteBearer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.SpgwService/DeleteBearer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpgwServiceServer).DeleteBearer(ctx, req.(*DeleteBearerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SpgwService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.lte.SpgwService",
	HandlerType: (*SpgwServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBearer",
			Handler:    _SpgwService_CreateBearer_Handler,
		},
		{
			MethodName: "DeleteBearer",
			Handler:    _SpgwService_DeleteBearer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lte/protos/spgw_service.proto",
}
