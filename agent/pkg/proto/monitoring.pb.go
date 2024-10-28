// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: proto/monitoring.proto

package server

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

type ContainerLogMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerId   string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ContainerName string `protobuf:"bytes,2,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	Image         string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	State         string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	LogPath       string `protobuf:"bytes,5,opt,name=log_path,json=logPath,proto3" json:"log_path,omitempty"`
	LogDriver     string `protobuf:"bytes,6,opt,name=log_driver,json=logDriver,proto3" json:"log_driver,omitempty"`
}

func (x *ContainerLogMetadata) Reset() {
	*x = ContainerLogMetadata{}
	mi := &file_proto_monitoring_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerLogMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerLogMetadata) ProtoMessage() {}

func (x *ContainerLogMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_proto_monitoring_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerLogMetadata.ProtoReflect.Descriptor instead.
func (*ContainerLogMetadata) Descriptor() ([]byte, []int) {
	return file_proto_monitoring_proto_rawDescGZIP(), []int{0}
}

func (x *ContainerLogMetadata) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *ContainerLogMetadata) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *ContainerLogMetadata) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ContainerLogMetadata) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ContainerLogMetadata) GetLogPath() string {
	if x != nil {
		return x.LogPath
	}
	return ""
}

func (x *ContainerLogMetadata) GetLogDriver() string {
	if x != nil {
		return x.LogDriver
	}
	return ""
}

type LogData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *ContainerLogMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Log      string                `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *LogData) Reset() {
	*x = LogData{}
	mi := &file_proto_monitoring_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogData) ProtoMessage() {}

func (x *LogData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_monitoring_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogData.ProtoReflect.Descriptor instead.
func (*LogData) Descriptor() ([]byte, []int) {
	return file_proto_monitoring_proto_rawDescGZIP(), []int{1}
}

func (x *LogData) GetMetadata() *ContainerLogMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *LogData) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

type ContainerUsageStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerId    string  `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Timestamp      int64   `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	CpuPercent     float64 `protobuf:"fixed64,3,opt,name=cpu_percent,json=cpuPercent,proto3" json:"cpu_percent,omitempty"`
	CpuUsage       uint64  `protobuf:"varint,4,opt,name=cpu_usage,json=cpuUsage,proto3" json:"cpu_usage,omitempty"`
	SystemCpuUsage uint64  `protobuf:"varint,5,opt,name=system_cpu_usage,json=systemCpuUsage,proto3" json:"system_cpu_usage,omitempty"`
	MemoryUsage    uint64  `protobuf:"varint,6,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	MemoryLimit    uint64  `protobuf:"varint,7,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	MemoryPercent  float64 `protobuf:"fixed64,8,opt,name=memory_percent,json=memoryPercent,proto3" json:"memory_percent,omitempty"`
	MemoryCache    uint64  `protobuf:"varint,9,opt,name=memory_cache,json=memoryCache,proto3" json:"memory_cache,omitempty"`
}

func (x *ContainerUsageStats) Reset() {
	*x = ContainerUsageStats{}
	mi := &file_proto_monitoring_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerUsageStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerUsageStats) ProtoMessage() {}

func (x *ContainerUsageStats) ProtoReflect() protoreflect.Message {
	mi := &file_proto_monitoring_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerUsageStats.ProtoReflect.Descriptor instead.
func (*ContainerUsageStats) Descriptor() ([]byte, []int) {
	return file_proto_monitoring_proto_rawDescGZIP(), []int{2}
}

func (x *ContainerUsageStats) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *ContainerUsageStats) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ContainerUsageStats) GetCpuPercent() float64 {
	if x != nil {
		return x.CpuPercent
	}
	return 0
}

func (x *ContainerUsageStats) GetCpuUsage() uint64 {
	if x != nil {
		return x.CpuUsage
	}
	return 0
}

func (x *ContainerUsageStats) GetSystemCpuUsage() uint64 {
	if x != nil {
		return x.SystemCpuUsage
	}
	return 0
}

func (x *ContainerUsageStats) GetMemoryUsage() uint64 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *ContainerUsageStats) GetMemoryLimit() uint64 {
	if x != nil {
		return x.MemoryLimit
	}
	return 0
}

func (x *ContainerUsageStats) GetMemoryPercent() float64 {
	if x != nil {
		return x.MemoryPercent
	}
	return 0
}

func (x *ContainerUsageStats) GetMemoryCache() uint64 {
	if x != nil {
		return x.MemoryCache
	}
	return 0
}

type LogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LogResponse) Reset() {
	*x = LogResponse{}
	mi := &file_proto_monitoring_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogResponse) ProtoMessage() {}

func (x *LogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_monitoring_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogResponse.ProtoReflect.Descriptor instead.
func (*LogResponse) Descriptor() ([]byte, []int) {
	return file_proto_monitoring_proto_rawDescGZIP(), []int{3}
}

func (x *LogResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UsageResponse) Reset() {
	*x = UsageResponse{}
	mi := &file_proto_monitoring_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsageResponse) ProtoMessage() {}

func (x *UsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_monitoring_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsageResponse.ProtoReflect.Descriptor instead.
func (*UsageResponse) Descriptor() ([]byte, []int) {
	return file_proto_monitoring_proto_rawDescGZIP(), []int{4}
}

func (x *UsageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_monitoring_proto protoreflect.FileDescriptor

var file_proto_monitoring_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x22, 0xc6, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1d,
	0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x22, 0x59, 0x0a,
	0x07, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x22, 0xce, 0x02, 0x0a, 0x13, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x70, 0x75, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x70, 0x75, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x70, 0x75, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x28, 0x0a, 0x10, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x43, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x43, 0x61, 0x63, 0x68, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x55, 0x0a,
	0x13, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f,
	0x67, 0x73, 0x12, 0x13, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x17, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x30, 0x01, 0x32, 0x66, 0x0a, 0x15, 0x55, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a,
	0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x1a, 0x19, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x26, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x78, 0x66, 0x6c,
	0x6f, 0x77, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_monitoring_proto_rawDescOnce sync.Once
	file_proto_monitoring_proto_rawDescData = file_proto_monitoring_proto_rawDesc
)

func file_proto_monitoring_proto_rawDescGZIP() []byte {
	file_proto_monitoring_proto_rawDescOnce.Do(func() {
		file_proto_monitoring_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_monitoring_proto_rawDescData)
	})
	return file_proto_monitoring_proto_rawDescData
}

var file_proto_monitoring_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_monitoring_proto_goTypes = []any{
	(*ContainerLogMetadata)(nil), // 0: monitoring.ContainerLogMetadata
	(*LogData)(nil),              // 1: monitoring.LogData
	(*ContainerUsageStats)(nil),  // 2: monitoring.ContainerUsageStats
	(*LogResponse)(nil),          // 3: monitoring.LogResponse
	(*UsageResponse)(nil),        // 4: monitoring.UsageResponse
}
var file_proto_monitoring_proto_depIdxs = []int32{
	0, // 0: monitoring.LogData.metadata:type_name -> monitoring.ContainerLogMetadata
	1, // 1: monitoring.LogStreamingService.StreamLogs:input_type -> monitoring.LogData
	2, // 2: monitoring.UsageStreamingService.StreamUsage:input_type -> monitoring.ContainerUsageStats
	3, // 3: monitoring.LogStreamingService.StreamLogs:output_type -> monitoring.LogResponse
	4, // 4: monitoring.UsageStreamingService.StreamUsage:output_type -> monitoring.UsageResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_monitoring_proto_init() }
func file_proto_monitoring_proto_init() {
	if File_proto_monitoring_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_monitoring_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_monitoring_proto_goTypes,
		DependencyIndexes: file_proto_monitoring_proto_depIdxs,
		MessageInfos:      file_proto_monitoring_proto_msgTypes,
	}.Build()
	File_proto_monitoring_proto = out.File
	file_proto_monitoring_proto_rawDesc = nil
	file_proto_monitoring_proto_goTypes = nil
	file_proto_monitoring_proto_depIdxs = nil
}