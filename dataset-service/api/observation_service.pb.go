// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: caraml/timber/v1/observation_service.proto

package api

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

// Data source where Observation logs should be retrieved from
type ObservationServiceDataSourceType int32

const (
	ObservationServiceDataSourceType_OBSERVATION_SERVICE_DATA_SOURCE_TYPE_UNSPECIFIED ObservationServiceDataSourceType = 0
	// No-Op represents no need to fetch logs from any data source, this should be selected if
	// Observation Service should be deployed for just the eager API
	ObservationServiceDataSourceType_OBSERVATION_SERVICE_DATA_SOURCE_TYPE_EAGER ObservationServiceDataSourceType = 1
	// Observation Service will poll logs from a Kafka source
	ObservationServiceDataSourceType_OBSERVATION_SERVICE_DATA_SOURCE_TYPE_KAFKA ObservationServiceDataSourceType = 2
)

// Enum value maps for ObservationServiceDataSourceType.
var (
	ObservationServiceDataSourceType_name = map[int32]string{
		0: "OBSERVATION_SERVICE_DATA_SOURCE_TYPE_UNSPECIFIED",
		1: "OBSERVATION_SERVICE_DATA_SOURCE_TYPE_EAGER",
		2: "OBSERVATION_SERVICE_DATA_SOURCE_TYPE_KAFKA",
	}
	ObservationServiceDataSourceType_value = map[string]int32{
		"OBSERVATION_SERVICE_DATA_SOURCE_TYPE_UNSPECIFIED": 0,
		"OBSERVATION_SERVICE_DATA_SOURCE_TYPE_EAGER":       1,
		"OBSERVATION_SERVICE_DATA_SOURCE_TYPE_KAFKA":       2,
	}
)

func (x ObservationServiceDataSourceType) Enum() *ObservationServiceDataSourceType {
	p := new(ObservationServiceDataSourceType)
	*p = x
	return p
}

func (x ObservationServiceDataSourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObservationServiceDataSourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_caraml_timber_v1_observation_service_proto_enumTypes[0].Descriptor()
}

func (ObservationServiceDataSourceType) Type() protoreflect.EnumType {
	return &file_caraml_timber_v1_observation_service_proto_enumTypes[0]
}

func (x ObservationServiceDataSourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObservationServiceDataSourceType.Descriptor instead.
func (ObservationServiceDataSourceType) EnumDescriptor() ([]byte, []int) {
	return file_caraml_timber_v1_observation_service_proto_rawDescGZIP(), []int{0}
}

// Data sink where Observation logs would be flushed to
type ObservationServiceDataSinkType int32

const (
	ObservationServiceDataSinkType_OBSERVATION_SERVICE_DATA_SINK_TYPE_UNSPECIFIED ObservationServiceDataSinkType = 0
	// No-Op represents no need to flush logs to any data sink
	ObservationServiceDataSinkType_OBSERVATION_SERVICE_DATA_SINK_TYPE_NOOP ObservationServiceDataSinkType = 1
	// Observation Service will publish logs to standard output
	ObservationServiceDataSinkType_OBSERVATION_SERVICE_DATA_SINK_TYPE_STDOUT ObservationServiceDataSinkType = 2
	// Observation Service will flush logs to a Kafka sink
	ObservationServiceDataSinkType_OBSERVATION_SERVICE_DATA_SINK_TYPE_KAFKA ObservationServiceDataSinkType = 3
	// Observation Service will flush logs to Fluentd
	ObservationServiceDataSinkType_OBSERVATION_SERVICE_DATA_SINK_TYPE_FLUENTD ObservationServiceDataSinkType = 4
)

// Enum value maps for ObservationServiceDataSinkType.
var (
	ObservationServiceDataSinkType_name = map[int32]string{
		0: "OBSERVATION_SERVICE_DATA_SINK_TYPE_UNSPECIFIED",
		1: "OBSERVATION_SERVICE_DATA_SINK_TYPE_NOOP",
		2: "OBSERVATION_SERVICE_DATA_SINK_TYPE_STDOUT",
		3: "OBSERVATION_SERVICE_DATA_SINK_TYPE_KAFKA",
		4: "OBSERVATION_SERVICE_DATA_SINK_TYPE_FLUENTD",
	}
	ObservationServiceDataSinkType_value = map[string]int32{
		"OBSERVATION_SERVICE_DATA_SINK_TYPE_UNSPECIFIED": 0,
		"OBSERVATION_SERVICE_DATA_SINK_TYPE_NOOP":        1,
		"OBSERVATION_SERVICE_DATA_SINK_TYPE_STDOUT":      2,
		"OBSERVATION_SERVICE_DATA_SINK_TYPE_KAFKA":       3,
		"OBSERVATION_SERVICE_DATA_SINK_TYPE_FLUENTD":     4,
	}
)

func (x ObservationServiceDataSinkType) Enum() *ObservationServiceDataSinkType {
	p := new(ObservationServiceDataSinkType)
	*p = x
	return p
}

func (x ObservationServiceDataSinkType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObservationServiceDataSinkType) Descriptor() protoreflect.EnumDescriptor {
	return file_caraml_timber_v1_observation_service_proto_enumTypes[1].Descriptor()
}

func (ObservationServiceDataSinkType) Type() protoreflect.EnumType {
	return &file_caraml_timber_v1_observation_service_proto_enumTypes[1]
}

func (x ObservationServiceDataSinkType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObservationServiceDataSinkType.Descriptor instead.
func (ObservationServiceDataSinkType) EnumDescriptor() ([]byte, []int) {
	return file_caraml_timber_v1_observation_service_proto_rawDescGZIP(), []int{1}
}

// Configurations of Data source where Observation logs should be retrieved from
type ObservationServiceDataSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        ObservationServiceDataSourceType `protobuf:"varint,1,opt,name=type,proto3,enum=caraml.timber.v1.ObservationServiceDataSourceType" json:"type,omitempty"`
	KafkaConfig *KafkaConfig                     `protobuf:"bytes,2,opt,name=kafka_config,json=kafkaConfig,proto3" json:"kafka_config,omitempty"`
}

func (x *ObservationServiceDataSource) Reset() {
	*x = ObservationServiceDataSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caraml_timber_v1_observation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObservationServiceDataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservationServiceDataSource) ProtoMessage() {}

func (x *ObservationServiceDataSource) ProtoReflect() protoreflect.Message {
	mi := &file_caraml_timber_v1_observation_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservationServiceDataSource.ProtoReflect.Descriptor instead.
func (*ObservationServiceDataSource) Descriptor() ([]byte, []int) {
	return file_caraml_timber_v1_observation_service_proto_rawDescGZIP(), []int{0}
}

func (x *ObservationServiceDataSource) GetType() ObservationServiceDataSourceType {
	if x != nil {
		return x.Type
	}
	return ObservationServiceDataSourceType_OBSERVATION_SERVICE_DATA_SOURCE_TYPE_UNSPECIFIED
}

func (x *ObservationServiceDataSource) GetKafkaConfig() *KafkaConfig {
	if x != nil {
		return x.KafkaConfig
	}
	return nil
}

// Configurations of Data sink where Observation logs would be flushed to
type ObservationServiceDataSink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          ObservationServiceDataSinkType `protobuf:"varint,1,opt,name=type,proto3,enum=caraml.timber.v1.ObservationServiceDataSinkType" json:"type,omitempty"`
	KafkaConfig   *KafkaConfig                   `protobuf:"bytes,2,opt,name=kafka_config,json=kafkaConfig,proto3" json:"kafka_config,omitempty"`
	FluentdConfig *FluentdConfig                 `protobuf:"bytes,3,opt,name=fluentd_config,json=fluentdConfig,proto3" json:"fluentd_config,omitempty"`
}

func (x *ObservationServiceDataSink) Reset() {
	*x = ObservationServiceDataSink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caraml_timber_v1_observation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObservationServiceDataSink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservationServiceDataSink) ProtoMessage() {}

func (x *ObservationServiceDataSink) ProtoReflect() protoreflect.Message {
	mi := &file_caraml_timber_v1_observation_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservationServiceDataSink.ProtoReflect.Descriptor instead.
func (*ObservationServiceDataSink) Descriptor() ([]byte, []int) {
	return file_caraml_timber_v1_observation_service_proto_rawDescGZIP(), []int{1}
}

func (x *ObservationServiceDataSink) GetType() ObservationServiceDataSinkType {
	if x != nil {
		return x.Type
	}
	return ObservationServiceDataSinkType_OBSERVATION_SERVICE_DATA_SINK_TYPE_UNSPECIFIED
}

func (x *ObservationServiceDataSink) GetKafkaConfig() *KafkaConfig {
	if x != nil {
		return x.KafkaConfig
	}
	return nil
}

func (x *ObservationServiceDataSink) GetFluentdConfig() *FluentdConfig {
	if x != nil {
		return x.FluentdConfig
	}
	return nil
}

// ObservationServiceConfig describes details of a Observation Service
type ObservationServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required to postfix helm release name with provided Service name as a CaraML project could have multiple different services
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Configuration for pull-based Observation Service data source.
	Source *ObservationServiceDataSource `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	// Configuration for Observation Service data sink.
	Sink *ObservationServiceDataSink `protobuf:"bytes,3,opt,name=sink,proto3" json:"sink,omitempty"`
}

func (x *ObservationServiceConfig) Reset() {
	*x = ObservationServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caraml_timber_v1_observation_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObservationServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservationServiceConfig) ProtoMessage() {}

func (x *ObservationServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_caraml_timber_v1_observation_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservationServiceConfig.ProtoReflect.Descriptor instead.
func (*ObservationServiceConfig) Descriptor() ([]byte, []int) {
	return file_caraml_timber_v1_observation_service_proto_rawDescGZIP(), []int{2}
}

func (x *ObservationServiceConfig) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *ObservationServiceConfig) GetSource() *ObservationServiceDataSource {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *ObservationServiceConfig) GetSink() *ObservationServiceDataSink {
	if x != nil {
		return x.Sink
	}
	return nil
}

// ObservationServiceResponse describes details of a deployed Observation Service
type ObservationServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of an Observation Service deployed by Dataset Service.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ObservationServiceResponse) Reset() {
	*x = ObservationServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caraml_timber_v1_observation_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObservationServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservationServiceResponse) ProtoMessage() {}

func (x *ObservationServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_caraml_timber_v1_observation_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservationServiceResponse.ProtoReflect.Descriptor instead.
func (*ObservationServiceResponse) Descriptor() ([]byte, []int) {
	return file_caraml_timber_v1_observation_service_proto_rawDescGZIP(), []int{3}
}

func (x *ObservationServiceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_caraml_timber_v1_observation_service_proto protoreflect.FileDescriptor

var file_caraml_timber_v1_observation_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2f, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x61,
	0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1a,
	0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2f, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x1c, 0x4f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x63, 0x61, 0x72, 0x61,
	0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x61,
	0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x66,
	0x6b, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xec, 0x01, 0x0a, 0x1a, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x69, 0x6e, 0x6b, 0x12, 0x44, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x30, 0x2e, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x69, 0x6e, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x46, 0x0a, 0x0e,
	0x66, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x64, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x64, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0xc7, 0x01, 0x0a, 0x18, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x04,
	0x73, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x61, 0x72,
	0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x04, 0x73, 0x69, 0x6e, 0x6b, 0x22, 0x2c,
	0x0a, 0x1a, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a, 0xb8, 0x01, 0x0a,
	0x20, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x34, 0x0a, 0x30, 0x4f, 0x42, 0x53, 0x45, 0x52, 0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2e, 0x0a, 0x2a, 0x4f, 0x42, 0x53, 0x45, 0x52,
	0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x44,
	0x41, 0x54, 0x41, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x45, 0x41, 0x47, 0x45, 0x52, 0x10, 0x01, 0x12, 0x2e, 0x0a, 0x2a, 0x4f, 0x42, 0x53, 0x45, 0x52,
	0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x44,
	0x41, 0x54, 0x41, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4b, 0x41, 0x46, 0x4b, 0x41, 0x10, 0x02, 0x2a, 0x8e, 0x02, 0x0a, 0x1e, 0x4f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x69, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x2e, 0x4f, 0x42,
	0x53, 0x45, 0x52, 0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x49, 0x4e, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2b,
	0x0a, 0x27, 0x4f, 0x42, 0x53, 0x45, 0x52, 0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x49, 0x4e, 0x4b, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4f, 0x50, 0x10, 0x01, 0x12, 0x2d, 0x0a, 0x29, 0x4f,
	0x42, 0x53, 0x45, 0x52, 0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x49, 0x4e, 0x4b, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x54, 0x44, 0x4f, 0x55, 0x54, 0x10, 0x02, 0x12, 0x2c, 0x0a, 0x28, 0x4f, 0x42,
	0x53, 0x45, 0x52, 0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x49, 0x4e, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4b, 0x41, 0x46, 0x4b, 0x41, 0x10, 0x03, 0x12, 0x2e, 0x0a, 0x2a, 0x4f, 0x42, 0x53, 0x45,
	0x52, 0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x49, 0x4e, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46,
	0x4c, 0x55, 0x45, 0x4e, 0x54, 0x44, 0x10, 0x04, 0x42, 0xd8, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x17, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2d,
	0x64, 0x65, 0x76, 0x2f, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2f, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b,
	0x61, 0x70, 0x69, 0xa2, 0x02, 0x03, 0x43, 0x54, 0x58, 0xaa, 0x02, 0x10, 0x43, 0x61, 0x72, 0x61,
	0x6d, 0x6c, 0x2e, 0x54, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x43,
	0x61, 0x72, 0x61, 0x6d, 0x6c, 0x5c, 0x54, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1c, 0x43, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x5c, 0x54, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x12, 0x43, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x3a, 0x3a, 0x54, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_caraml_timber_v1_observation_service_proto_rawDescOnce sync.Once
	file_caraml_timber_v1_observation_service_proto_rawDescData = file_caraml_timber_v1_observation_service_proto_rawDesc
)

func file_caraml_timber_v1_observation_service_proto_rawDescGZIP() []byte {
	file_caraml_timber_v1_observation_service_proto_rawDescOnce.Do(func() {
		file_caraml_timber_v1_observation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_caraml_timber_v1_observation_service_proto_rawDescData)
	})
	return file_caraml_timber_v1_observation_service_proto_rawDescData
}

var file_caraml_timber_v1_observation_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_caraml_timber_v1_observation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_caraml_timber_v1_observation_service_proto_goTypes = []interface{}{
	(ObservationServiceDataSourceType)(0), // 0: caraml.timber.v1.ObservationServiceDataSourceType
	(ObservationServiceDataSinkType)(0),   // 1: caraml.timber.v1.ObservationServiceDataSinkType
	(*ObservationServiceDataSource)(nil),  // 2: caraml.timber.v1.ObservationServiceDataSource
	(*ObservationServiceDataSink)(nil),    // 3: caraml.timber.v1.ObservationServiceDataSink
	(*ObservationServiceConfig)(nil),      // 4: caraml.timber.v1.ObservationServiceConfig
	(*ObservationServiceResponse)(nil),    // 5: caraml.timber.v1.ObservationServiceResponse
	(*KafkaConfig)(nil),                   // 6: caraml.timber.v1.KafkaConfig
	(*FluentdConfig)(nil),                 // 7: caraml.timber.v1.FluentdConfig
}
var file_caraml_timber_v1_observation_service_proto_depIdxs = []int32{
	0, // 0: caraml.timber.v1.ObservationServiceDataSource.type:type_name -> caraml.timber.v1.ObservationServiceDataSourceType
	6, // 1: caraml.timber.v1.ObservationServiceDataSource.kafka_config:type_name -> caraml.timber.v1.KafkaConfig
	1, // 2: caraml.timber.v1.ObservationServiceDataSink.type:type_name -> caraml.timber.v1.ObservationServiceDataSinkType
	6, // 3: caraml.timber.v1.ObservationServiceDataSink.kafka_config:type_name -> caraml.timber.v1.KafkaConfig
	7, // 4: caraml.timber.v1.ObservationServiceDataSink.fluentd_config:type_name -> caraml.timber.v1.FluentdConfig
	2, // 5: caraml.timber.v1.ObservationServiceConfig.source:type_name -> caraml.timber.v1.ObservationServiceDataSource
	3, // 6: caraml.timber.v1.ObservationServiceConfig.sink:type_name -> caraml.timber.v1.ObservationServiceDataSink
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_caraml_timber_v1_observation_service_proto_init() }
func file_caraml_timber_v1_observation_service_proto_init() {
	if File_caraml_timber_v1_observation_service_proto != nil {
		return
	}
	file_caraml_timber_v1_log_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_caraml_timber_v1_observation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObservationServiceDataSource); i {
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
		file_caraml_timber_v1_observation_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObservationServiceDataSink); i {
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
		file_caraml_timber_v1_observation_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObservationServiceConfig); i {
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
		file_caraml_timber_v1_observation_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObservationServiceResponse); i {
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
			RawDescriptor: file_caraml_timber_v1_observation_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_caraml_timber_v1_observation_service_proto_goTypes,
		DependencyIndexes: file_caraml_timber_v1_observation_service_proto_depIdxs,
		EnumInfos:         file_caraml_timber_v1_observation_service_proto_enumTypes,
		MessageInfos:      file_caraml_timber_v1_observation_service_proto_msgTypes,
	}.Build()
	File_caraml_timber_v1_observation_service_proto = out.File
	file_caraml_timber_v1_observation_service_proto_rawDesc = nil
	file_caraml_timber_v1_observation_service_proto_goTypes = nil
	file_caraml_timber_v1_observation_service_proto_depIdxs = nil
}
