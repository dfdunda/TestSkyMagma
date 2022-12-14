// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.10.0
// source: lte/protos/oai/std_3gpp_types.proto

package oai

import (
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

// parameters_list
type ParametersList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parameters    []*Parameter `protobuf:"bytes,1,rep,name=parameters,proto3" json:"parameters,omitempty"`
	NumParameters uint32       `protobuf:"varint,2,opt,name=num_parameters,json=numParameters,proto3" json:"num_parameters,omitempty"`
}

func (x *ParametersList) Reset() {
	*x = ParametersList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParametersList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParametersList) ProtoMessage() {}

func (x *ParametersList) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParametersList.ProtoReflect.Descriptor instead.
func (*ParametersList) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_std_3gpp_types_proto_rawDescGZIP(), []int{0}
}

func (x *ParametersList) GetParameters() []*Parameter {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *ParametersList) GetNumParameters() uint32 {
	if x != nil {
		return x.NumParameters
	}
	return 0
}

// parameter_t
type Parameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParameterIdentifier uint32 `protobuf:"varint,1,opt,name=parameter_identifier,json=parameterIdentifier,proto3" json:"parameter_identifier,omitempty"`
	Length              uint32 `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	Contents            []byte `protobuf:"bytes,3,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (x *Parameter) Reset() {
	*x = Parameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Parameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parameter) ProtoMessage() {}

func (x *Parameter) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parameter.ProtoReflect.Descriptor instead.
func (*Parameter) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_std_3gpp_types_proto_rawDescGZIP(), []int{1}
}

func (x *Parameter) GetParameterIdentifier() uint32 {
	if x != nil {
		return x.ParameterIdentifier
	}
	return 0
}

func (x *Parameter) GetLength() uint32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *Parameter) GetContents() []byte {
	if x != nil {
		return x.Contents
	}
	return nil
}

// protocol_configuration_options_t
type Pco struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ext                      uint32         `protobuf:"varint,1,opt,name=ext,proto3" json:"ext,omitempty"`
	Spare                    uint32         `protobuf:"varint,2,opt,name=spare,proto3" json:"spare,omitempty"`
	ConfigurationProtocol    uint32         `protobuf:"varint,3,opt,name=configuration_protocol,json=configurationProtocol,proto3" json:"configuration_protocol,omitempty"`
	NumProtocolOrContainerId uint32         `protobuf:"varint,4,opt,name=num_protocol_or_container_id,json=numProtocolOrContainerId,proto3" json:"num_protocol_or_container_id,omitempty"`
	PcoProtocol              []*PcoProtocol `protobuf:"bytes,5,rep,name=pco_protocol,json=pcoProtocol,proto3" json:"pco_protocol,omitempty"`
}

func (x *Pco) Reset() {
	*x = Pco{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pco) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pco) ProtoMessage() {}

func (x *Pco) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pco.ProtoReflect.Descriptor instead.
func (*Pco) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_std_3gpp_types_proto_rawDescGZIP(), []int{2}
}

func (x *Pco) GetExt() uint32 {
	if x != nil {
		return x.Ext
	}
	return 0
}

func (x *Pco) GetSpare() uint32 {
	if x != nil {
		return x.Spare
	}
	return 0
}

func (x *Pco) GetConfigurationProtocol() uint32 {
	if x != nil {
		return x.ConfigurationProtocol
	}
	return 0
}

func (x *Pco) GetNumProtocolOrContainerId() uint32 {
	if x != nil {
		return x.NumProtocolOrContainerId
	}
	return 0
}

func (x *Pco) GetPcoProtocol() []*PcoProtocol {
	if x != nil {
		return x.PcoProtocol
	}
	return nil
}

// pco_protocol_or_container_id
type PcoProtocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Length   uint32 `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	Contents []byte `protobuf:"bytes,3,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (x *PcoProtocol) Reset() {
	*x = PcoProtocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PcoProtocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PcoProtocol) ProtoMessage() {}

func (x *PcoProtocol) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PcoProtocol.ProtoReflect.Descriptor instead.
func (*PcoProtocol) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_std_3gpp_types_proto_rawDescGZIP(), []int{3}
}

func (x *PcoProtocol) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PcoProtocol) GetLength() uint32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *PcoProtocol) GetContents() []byte {
	if x != nil {
		return x.Contents
	}
	return nil
}

// fteid_t
type Fteid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ipv4Address   uint32 `protobuf:"varint,1,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address   []byte `protobuf:"bytes,2,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	InterfaceType uint32 `protobuf:"varint,3,opt,name=interface_type,json=interfaceType,proto3" json:"interface_type,omitempty"`
	Teid          uint32 `protobuf:"varint,4,opt,name=teid,proto3" json:"teid,omitempty"`
}

func (x *Fteid) Reset() {
	*x = Fteid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fteid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fteid) ProtoMessage() {}

func (x *Fteid) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fteid.ProtoReflect.Descriptor instead.
func (*Fteid) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_std_3gpp_types_proto_rawDescGZIP(), []int{4}
}

func (x *Fteid) GetIpv4Address() uint32 {
	if x != nil {
		return x.Ipv4Address
	}
	return 0
}

func (x *Fteid) GetIpv6Address() []byte {
	if x != nil {
		return x.Ipv6Address
	}
	return nil
}

func (x *Fteid) GetInterfaceType() uint32 {
	if x != nil {
		return x.InterfaceType
	}
	return 0
}

func (x *Fteid) GetTeid() uint32 {
	if x != nil {
		return x.Teid
	}
	return 0
}

// port_range
type PortRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LowLimit  uint32 `protobuf:"varint,1,opt,name=low_limit,json=lowLimit,proto3" json:"low_limit,omitempty"`
	HighLimit uint32 `protobuf:"varint,2,opt,name=high_limit,json=highLimit,proto3" json:"high_limit,omitempty"`
}

func (x *PortRange) Reset() {
	*x = PortRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortRange) ProtoMessage() {}

func (x *PortRange) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortRange.ProtoReflect.Descriptor instead.
func (*PortRange) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_std_3gpp_types_proto_rawDescGZIP(), []int{5}
}

func (x *PortRange) GetLowLimit() uint32 {
	if x != nil {
		return x.LowLimit
	}
	return 0
}

func (x *PortRange) GetHighLimit() uint32 {
	if x != nil {
		return x.HighLimit
	}
	return 0
}

// typdeofservice_trafficclass
type TypeOfServiceTrafficClass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Mask  uint32 `protobuf:"varint,2,opt,name=mask,proto3" json:"mask,omitempty"`
}

func (x *TypeOfServiceTrafficClass) Reset() {
	*x = TypeOfServiceTrafficClass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeOfServiceTrafficClass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeOfServiceTrafficClass) ProtoMessage() {}

func (x *TypeOfServiceTrafficClass) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeOfServiceTrafficClass.ProtoReflect.Descriptor instead.
func (*TypeOfServiceTrafficClass) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_std_3gpp_types_proto_rawDescGZIP(), []int{6}
}

func (x *TypeOfServiceTrafficClass) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *TypeOfServiceTrafficClass) GetMask() uint32 {
	if x != nil {
		return x.Mask
	}
	return 0
}

// ipvremoteaddr
type IpRemoteAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr uint32 `protobuf:"varint,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Mask uint32 `protobuf:"varint,2,opt,name=mask,proto3" json:"mask,omitempty"`
}

func (x *IpRemoteAddress) Reset() {
	*x = IpRemoteAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpRemoteAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpRemoteAddress) ProtoMessage() {}

func (x *IpRemoteAddress) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpRemoteAddress.ProtoReflect.Descriptor instead.
func (*IpRemoteAddress) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_std_3gpp_types_proto_rawDescGZIP(), []int{7}
}

func (x *IpRemoteAddress) GetAddr() uint32 {
	if x != nil {
		return x.Addr
	}
	return 0
}

func (x *IpRemoteAddress) GetMask() uint32 {
	if x != nil {
		return x.Mask
	}
	return 0
}

// packet_filter_contents
type PacketFilterContents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flags                        uint32                     `protobuf:"varint,1,opt,name=flags,proto3" json:"flags,omitempty"`
	Ipv4RemoteAddresses          []*IpRemoteAddress         `protobuf:"bytes,2,rep,name=ipv4_remote_addresses,json=ipv4RemoteAddresses,proto3" json:"ipv4_remote_addresses,omitempty"`
	Ipv6RemoteAddresses          []*IpRemoteAddress         `protobuf:"bytes,3,rep,name=ipv6_remote_addresses,json=ipv6RemoteAddresses,proto3" json:"ipv6_remote_addresses,omitempty"`
	ProtocolIdentifierNextheader uint32                     `protobuf:"varint,4,opt,name=protocol_identifier_nextheader,json=protocolIdentifierNextheader,proto3" json:"protocol_identifier_nextheader,omitempty"`
	SingleLocalPort              uint32                     `protobuf:"varint,10,opt,name=single_local_port,json=singleLocalPort,proto3" json:"single_local_port,omitempty"`
	LocalPortRange               *PortRange                 `protobuf:"bytes,11,opt,name=local_port_range,json=localPortRange,proto3" json:"local_port_range,omitempty"`
	SingleRemotePort             uint32                     `protobuf:"varint,12,opt,name=single_remote_port,json=singleRemotePort,proto3" json:"single_remote_port,omitempty"`
	RemotePortRange              *PortRange                 `protobuf:"bytes,13,opt,name=remote_port_range,json=remotePortRange,proto3" json:"remote_port_range,omitempty"`
	SecurityParameterIndex       uint32                     `protobuf:"varint,20,opt,name=security_parameter_index,json=securityParameterIndex,proto3" json:"security_parameter_index,omitempty"`
	TypeOfServiceTrafficClass    *TypeOfServiceTrafficClass `protobuf:"bytes,21,opt,name=type_of_service_traffic_class,json=typeOfServiceTrafficClass,proto3" json:"type_of_service_traffic_class,omitempty"`
	FlowLabel                    uint32                     `protobuf:"varint,22,opt,name=flow_label,json=flowLabel,proto3" json:"flow_label,omitempty"`
}

func (x *PacketFilterContents) Reset() {
	*x = PacketFilterContents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PacketFilterContents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketFilterContents) ProtoMessage() {}

func (x *PacketFilterContents) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_std_3gpp_types_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketFilterContents.ProtoReflect.Descriptor instead.
func (*PacketFilterContents) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_std_3gpp_types_proto_rawDescGZIP(), []int{8}
}

func (x *PacketFilterContents) GetFlags() uint32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *PacketFilterContents) GetIpv4RemoteAddresses() []*IpRemoteAddress {
	if x != nil {
		return x.Ipv4RemoteAddresses
	}
	return nil
}

func (x *PacketFilterContents) GetIpv6RemoteAddresses() []*IpRemoteAddress {
	if x != nil {
		return x.Ipv6RemoteAddresses
	}
	return nil
}

func (x *PacketFilterContents) GetProtocolIdentifierNextheader() uint32 {
	if x != nil {
		return x.ProtocolIdentifierNextheader
	}
	return 0
}

func (x *PacketFilterContents) GetSingleLocalPort() uint32 {
	if x != nil {
		return x.SingleLocalPort
	}
	return 0
}

func (x *PacketFilterContents) GetLocalPortRange() *PortRange {
	if x != nil {
		return x.LocalPortRange
	}
	return nil
}

func (x *PacketFilterContents) GetSingleRemotePort() uint32 {
	if x != nil {
		return x.SingleRemotePort
	}
	return 0
}

func (x *PacketFilterContents) GetRemotePortRange() *PortRange {
	if x != nil {
		return x.RemotePortRange
	}
	return nil
}

func (x *PacketFilterContents) GetSecurityParameterIndex() uint32 {
	if x != nil {
		return x.SecurityParameterIndex
	}
	return 0
}

func (x *PacketFilterContents) GetTypeOfServiceTrafficClass() *TypeOfServiceTrafficClass {
	if x != nil {
		return x.TypeOfServiceTrafficClass
	}
	return nil
}

func (x *PacketFilterContents) GetFlowLabel() uint32 {
	if x != nil {
		return x.FlowLabel
	}
	return 0
}

var File_lte_protos_oai_std_3gpp_types_proto protoreflect.FileDescriptor

var file_lte_protos_oai_std_3gpp_types_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6c, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x61, 0x69,
	0x2f, 0x73, 0x74, 0x64, 0x5f, 0x33, 0x67, 0x70, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65,
	0x2e, 0x6f, 0x61, 0x69, 0x22, 0x71, 0x0a, 0x0e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x61, 0x67,
	0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x72, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x14, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x13, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xe3, 0x01, 0x0a, 0x03,
	0x50, 0x63, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x70, 0x61, 0x72, 0x65, 0x12, 0x35, 0x0a, 0x16, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x12, 0x3e, 0x0a, 0x1c, 0x6e, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x5f, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x70, 0x63, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61,
	0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x50, 0x63, 0x6f, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x0b, 0x70, 0x63, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x22, 0x51, 0x0a, 0x0b, 0x50, 0x63, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x05, 0x46, 0x74, 0x65, 0x69, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x69, 0x70, 0x76, 0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x69, 0x70, 0x76, 0x36, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x65, 0x69, 0x64, 0x22,
	0x47, 0x0a, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x69, 0x67,
	0x68, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x68,
	0x69, 0x67, 0x68, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x45, 0x0a, 0x19, 0x54, 0x79, 0x70, 0x65,
	0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x22,
	0x39, 0x0a, 0x0f, 0x49, 0x70, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x22, 0xc3, 0x05, 0x0a, 0x14, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x52, 0x0a, 0x15, 0x69, 0x70, 0x76,
	0x34, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61,
	0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x49, 0x70, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x13, 0x69, 0x70, 0x76, 0x34, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x52, 0x0a,
	0x15, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x49, 0x70, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x13, 0x69, 0x70,
	0x76, 0x36, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x12, 0x44, 0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4e, 0x65, 0x78,
	0x74, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x42, 0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x50, 0x6f,
	0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6f,
	0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x10, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x44, 0x0a, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69,
	0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x6a, 0x0a, 0x1d, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6f, 0x66,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x19, 0x74, 0x79, 0x70, 0x65, 0x4f, 0x66, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x66, 0x6c, 0x6f, 0x77, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x42, 0x1f, 0x5a, 0x1d, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6c, 0x74, 0x65, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x61,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lte_protos_oai_std_3gpp_types_proto_rawDescOnce sync.Once
	file_lte_protos_oai_std_3gpp_types_proto_rawDescData = file_lte_protos_oai_std_3gpp_types_proto_rawDesc
)

func file_lte_protos_oai_std_3gpp_types_proto_rawDescGZIP() []byte {
	file_lte_protos_oai_std_3gpp_types_proto_rawDescOnce.Do(func() {
		file_lte_protos_oai_std_3gpp_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_lte_protos_oai_std_3gpp_types_proto_rawDescData)
	})
	return file_lte_protos_oai_std_3gpp_types_proto_rawDescData
}

var file_lte_protos_oai_std_3gpp_types_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_lte_protos_oai_std_3gpp_types_proto_goTypes = []interface{}{
	(*ParametersList)(nil),            // 0: magma.lte.oai.ParametersList
	(*Parameter)(nil),                 // 1: magma.lte.oai.Parameter
	(*Pco)(nil),                       // 2: magma.lte.oai.Pco
	(*PcoProtocol)(nil),               // 3: magma.lte.oai.PcoProtocol
	(*Fteid)(nil),                     // 4: magma.lte.oai.Fteid
	(*PortRange)(nil),                 // 5: magma.lte.oai.PortRange
	(*TypeOfServiceTrafficClass)(nil), // 6: magma.lte.oai.TypeOfServiceTrafficClass
	(*IpRemoteAddress)(nil),           // 7: magma.lte.oai.IpRemoteAddress
	(*PacketFilterContents)(nil),      // 8: magma.lte.oai.PacketFilterContents
}
var file_lte_protos_oai_std_3gpp_types_proto_depIdxs = []int32{
	1, // 0: magma.lte.oai.ParametersList.parameters:type_name -> magma.lte.oai.Parameter
	3, // 1: magma.lte.oai.Pco.pco_protocol:type_name -> magma.lte.oai.PcoProtocol
	7, // 2: magma.lte.oai.PacketFilterContents.ipv4_remote_addresses:type_name -> magma.lte.oai.IpRemoteAddress
	7, // 3: magma.lte.oai.PacketFilterContents.ipv6_remote_addresses:type_name -> magma.lte.oai.IpRemoteAddress
	5, // 4: magma.lte.oai.PacketFilterContents.local_port_range:type_name -> magma.lte.oai.PortRange
	5, // 5: magma.lte.oai.PacketFilterContents.remote_port_range:type_name -> magma.lte.oai.PortRange
	6, // 6: magma.lte.oai.PacketFilterContents.type_of_service_traffic_class:type_name -> magma.lte.oai.TypeOfServiceTrafficClass
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_lte_protos_oai_std_3gpp_types_proto_init() }
func file_lte_protos_oai_std_3gpp_types_proto_init() {
	if File_lte_protos_oai_std_3gpp_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lte_protos_oai_std_3gpp_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParametersList); i {
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
		file_lte_protos_oai_std_3gpp_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Parameter); i {
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
		file_lte_protos_oai_std_3gpp_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pco); i {
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
		file_lte_protos_oai_std_3gpp_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PcoProtocol); i {
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
		file_lte_protos_oai_std_3gpp_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fteid); i {
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
		file_lte_protos_oai_std_3gpp_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortRange); i {
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
		file_lte_protos_oai_std_3gpp_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeOfServiceTrafficClass); i {
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
		file_lte_protos_oai_std_3gpp_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpRemoteAddress); i {
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
		file_lte_protos_oai_std_3gpp_types_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PacketFilterContents); i {
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
			RawDescriptor: file_lte_protos_oai_std_3gpp_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lte_protos_oai_std_3gpp_types_proto_goTypes,
		DependencyIndexes: file_lte_protos_oai_std_3gpp_types_proto_depIdxs,
		MessageInfos:      file_lte_protos_oai_std_3gpp_types_proto_msgTypes,
	}.Build()
	File_lte_protos_oai_std_3gpp_types_proto = out.File
	file_lte_protos_oai_std_3gpp_types_proto_rawDesc = nil
	file_lte_protos_oai_std_3gpp_types_proto_goTypes = nil
	file_lte_protos_oai_std_3gpp_types_proto_depIdxs = nil
}
