//
// Copyright 2020 The Magma Authors.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.10.0
// source: lte/cloud/go/services/smsd/storage/storage.proto

package storage

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type MessageStatus int32

const (
	MessageStatus_WAITING   MessageStatus = 0
	MessageStatus_DELIVERED MessageStatus = 1
	MessageStatus_FAILED    MessageStatus = 2
)

// Enum value maps for MessageStatus.
var (
	MessageStatus_name = map[int32]string{
		0: "WAITING",
		1: "DELIVERED",
		2: "FAILED",
	}
	MessageStatus_value = map[string]int32{
		"WAITING":   0,
		"DELIVERED": 1,
		"FAILED":    2,
	}
)

func (x MessageStatus) Enum() *MessageStatus {
	p := new(MessageStatus)
	*p = x
	return p
}

func (x MessageStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_lte_cloud_go_services_smsd_storage_storage_proto_enumTypes[0].Descriptor()
}

func (MessageStatus) Type() protoreflect.EnumType {
	return &file_lte_cloud_go_services_smsd_storage_storage_proto_enumTypes[0]
}

func (x MessageStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageStatus.Descriptor instead.
func (MessageStatus) EnumDescriptor() ([]byte, []int) {
	return file_lte_cloud_go_services_smsd_storage_storage_proto_rawDescGZIP(), []int{0}
}

// SMS represents a message tracked by the smsd service
type SMS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pk uniquely identifies an SMS message (generated unique key)
	Pk string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk,omitempty"`
	// delivery status of the SMS
	Status MessageStatus `protobuf:"varint,2,opt,name=status,proto3,enum=magma.lte.smsd.storage.MessageStatus" json:"status,omitempty"`
	// destination for the message
	Imsi string `protobuf:"bytes,10,opt,name=imsi,proto3" json:"imsi,omitempty"`
	// source MSISDN for the mesasge, to be encoded into the SMS
	SourceMsisdn string `protobuf:"bytes,11,opt,name=sourceMsisdn,proto3" json:"sourceMsisdn,omitempty"`
	// the desired message content of the SMS
	Message string `protobuf:"bytes,12,opt,name=message,proto3" json:"message,omitempty"`
	// time at which the message was created in the system
	CreatedTime *timestamp.Timestamp `protobuf:"bytes,20,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	// time that we last tried delivering this message. if status is delivered,
	// this will be the delivery time
	LastDeliveryAttemptTime *timestamp.Timestamp `protobuf:"bytes,21,opt,name=lastDeliveryAttemptTime,proto3" json:"lastDeliveryAttemptTime,omitempty"`
	// number of times we've attempted to send this SMS
	AttemptCount uint32 `protobuf:"varint,22,opt,name=attemptCount,proto3" json:"attemptCount,omitempty"`
	// error message from the most recent failed delivery attempt
	DeliveryError string `protobuf:"bytes,23,opt,name=deliveryError,proto3" json:"deliveryError,omitempty"`
	// Internal field which holds the reference numbers assigned to an SMS
	// which is in flight.
	// Value is a bytearray because one message could result in multiple SMSs
	// being sent to the UE.
	// If this field is non-empty, the message is in flight. Otherwise, it has
	// yet to be delivered to an AGW.
	RefNums []byte `protobuf:"bytes,30,opt,name=refNums,proto3" json:"refNums,omitempty"`
}

func (x *SMS) Reset() {
	*x = SMS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_cloud_go_services_smsd_storage_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SMS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SMS) ProtoMessage() {}

func (x *SMS) ProtoReflect() protoreflect.Message {
	mi := &file_lte_cloud_go_services_smsd_storage_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SMS.ProtoReflect.Descriptor instead.
func (*SMS) Descriptor() ([]byte, []int) {
	return file_lte_cloud_go_services_smsd_storage_storage_proto_rawDescGZIP(), []int{0}
}

func (x *SMS) GetPk() string {
	if x != nil {
		return x.Pk
	}
	return ""
}

func (x *SMS) GetStatus() MessageStatus {
	if x != nil {
		return x.Status
	}
	return MessageStatus_WAITING
}

func (x *SMS) GetImsi() string {
	if x != nil {
		return x.Imsi
	}
	return ""
}

func (x *SMS) GetSourceMsisdn() string {
	if x != nil {
		return x.SourceMsisdn
	}
	return ""
}

func (x *SMS) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SMS) GetCreatedTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *SMS) GetLastDeliveryAttemptTime() *timestamp.Timestamp {
	if x != nil {
		return x.LastDeliveryAttemptTime
	}
	return nil
}

func (x *SMS) GetAttemptCount() uint32 {
	if x != nil {
		return x.AttemptCount
	}
	return 0
}

func (x *SMS) GetDeliveryError() string {
	if x != nil {
		return x.DeliveryError
	}
	return ""
}

func (x *SMS) GetRefNums() []byte {
	if x != nil {
		return x.RefNums
	}
	return nil
}

// MutableSMS encapsulates the state that service clients are allowed to set.
type MutableSMS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imsi         string `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	SourceMsisdn string `protobuf:"bytes,2,opt,name=sourceMsisdn,proto3" json:"sourceMsisdn,omitempty"`
	Message      string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MutableSMS) Reset() {
	*x = MutableSMS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_cloud_go_services_smsd_storage_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutableSMS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutableSMS) ProtoMessage() {}

func (x *MutableSMS) ProtoReflect() protoreflect.Message {
	mi := &file_lte_cloud_go_services_smsd_storage_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutableSMS.ProtoReflect.Descriptor instead.
func (*MutableSMS) Descriptor() ([]byte, []int) {
	return file_lte_cloud_go_services_smsd_storage_storage_proto_rawDescGZIP(), []int{1}
}

func (x *MutableSMS) GetImsi() string {
	if x != nil {
		return x.Imsi
	}
	return ""
}

func (x *MutableSMS) GetSourceMsisdn() string {
	if x != nil {
		return x.SourceMsisdn
	}
	return ""
}

func (x *MutableSMS) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_lte_cloud_go_services_smsd_storage_storage_proto protoreflect.FileDescriptor

var file_lte_cloud_go_services_smsd_storage_storage_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6c, 0x74, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x6d, 0x73, 0x64, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x73, 0x6d,
	0x73, 0x64, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x03, 0x0a, 0x03,
	0x53, 0x4d, 0x53, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x70, 0x6b, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e,
	0x73, 0x6d, 0x73, 0x64, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4d, 0x73, 0x69, 0x73, 0x64, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4d, 0x73, 0x69, 0x73, 0x64, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x54, 0x0a, 0x17, 0x6c, 0x61, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x17, 0x6c, 0x61, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x74, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x74, 0x74, 0x65,
	0x6d, 0x70, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x17, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x4e, 0x75, 0x6d, 0x73, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x66, 0x4e, 0x75, 0x6d, 0x73, 0x22, 0x5e, 0x0a, 0x0a,
	0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x4d, 0x53, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d,
	0x73, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x12, 0x22,
	0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x73, 0x69, 0x73, 0x64, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x73, 0x69, 0x73,
	0x64, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x37, 0x0a, 0x0d,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45,
	0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x42, 0x2a, 0x5a, 0x28, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6c,
	0x74, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x6d, 0x73, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lte_cloud_go_services_smsd_storage_storage_proto_rawDescOnce sync.Once
	file_lte_cloud_go_services_smsd_storage_storage_proto_rawDescData = file_lte_cloud_go_services_smsd_storage_storage_proto_rawDesc
)

func file_lte_cloud_go_services_smsd_storage_storage_proto_rawDescGZIP() []byte {
	file_lte_cloud_go_services_smsd_storage_storage_proto_rawDescOnce.Do(func() {
		file_lte_cloud_go_services_smsd_storage_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_lte_cloud_go_services_smsd_storage_storage_proto_rawDescData)
	})
	return file_lte_cloud_go_services_smsd_storage_storage_proto_rawDescData
}

var file_lte_cloud_go_services_smsd_storage_storage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_lte_cloud_go_services_smsd_storage_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_lte_cloud_go_services_smsd_storage_storage_proto_goTypes = []interface{}{
	(MessageStatus)(0),          // 0: magma.lte.smsd.storage.MessageStatus
	(*SMS)(nil),                 // 1: magma.lte.smsd.storage.SMS
	(*MutableSMS)(nil),          // 2: magma.lte.smsd.storage.MutableSMS
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_lte_cloud_go_services_smsd_storage_storage_proto_depIdxs = []int32{
	0, // 0: magma.lte.smsd.storage.SMS.status:type_name -> magma.lte.smsd.storage.MessageStatus
	3, // 1: magma.lte.smsd.storage.SMS.createdTime:type_name -> google.protobuf.Timestamp
	3, // 2: magma.lte.smsd.storage.SMS.lastDeliveryAttemptTime:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_lte_cloud_go_services_smsd_storage_storage_proto_init() }
func file_lte_cloud_go_services_smsd_storage_storage_proto_init() {
	if File_lte_cloud_go_services_smsd_storage_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lte_cloud_go_services_smsd_storage_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SMS); i {
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
		file_lte_cloud_go_services_smsd_storage_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutableSMS); i {
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
			RawDescriptor: file_lte_cloud_go_services_smsd_storage_storage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lte_cloud_go_services_smsd_storage_storage_proto_goTypes,
		DependencyIndexes: file_lte_cloud_go_services_smsd_storage_storage_proto_depIdxs,
		EnumInfos:         file_lte_cloud_go_services_smsd_storage_storage_proto_enumTypes,
		MessageInfos:      file_lte_cloud_go_services_smsd_storage_storage_proto_msgTypes,
	}.Build()
	File_lte_cloud_go_services_smsd_storage_storage_proto = out.File
	file_lte_cloud_go_services_smsd_storage_storage_proto_rawDesc = nil
	file_lte_cloud_go_services_smsd_storage_storage_proto_goTypes = nil
	file_lte_cloud_go_services_smsd_storage_storage_proto_depIdxs = nil
}
