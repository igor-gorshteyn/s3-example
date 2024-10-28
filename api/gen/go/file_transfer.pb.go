// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: file_transfer.proto

package filetransfer

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

type FileChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename    string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Chunk       []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
	ChunkNumber int32  `protobuf:"varint,3,opt,name=chunk_number,json=chunkNumber,proto3" json:"chunk_number,omitempty"`
	TotalChunks int32  `protobuf:"varint,4,opt,name=total_chunks,json=totalChunks,proto3" json:"total_chunks,omitempty"`
	ChunkHash   string `protobuf:"bytes,5,opt,name=chunk_hash,json=chunkHash,proto3" json:"chunk_hash,omitempty"`
	ServiceName string `protobuf:"bytes,6,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *FileChunk) Reset() {
	*x = FileChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunk) ProtoMessage() {}

func (x *FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_file_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunk.ProtoReflect.Descriptor instead.
func (*FileChunk) Descriptor() ([]byte, []int) {
	return file_file_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *FileChunk) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileChunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *FileChunk) GetChunkNumber() int32 {
	if x != nil {
		return x.ChunkNumber
	}
	return 0
}

func (x *FileChunk) GetTotalChunks() int32 {
	if x != nil {
		return x.TotalChunks
	}
	return 0
}

func (x *FileChunk) GetChunkHash() string {
	if x != nil {
		return x.ChunkHash
	}
	return ""
}

func (x *FileChunk) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type TransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TransferResponse) Reset() {
	*x = TransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResponse) ProtoMessage() {}

func (x *TransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResponse.ProtoReflect.Descriptor instead.
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return file_file_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *TransferResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename    string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	ChunkNumber int32  `protobuf:"varint,2,opt,name=chunk_number,json=chunkNumber,proto3" json:"chunk_number,omitempty"`
	ChunkHash   string `protobuf:"bytes,3,opt,name=chunk_hash,json=chunkHash,proto3" json:"chunk_hash,omitempty"`
}

func (x *ChunkRequest) Reset() {
	*x = ChunkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_transfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkRequest) ProtoMessage() {}

func (x *ChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_transfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkRequest.ProtoReflect.Descriptor instead.
func (*ChunkRequest) Descriptor() ([]byte, []int) {
	return file_file_transfer_proto_rawDescGZIP(), []int{2}
}

func (x *ChunkRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *ChunkRequest) GetChunkNumber() int32 {
	if x != nil {
		return x.ChunkNumber
	}
	return 0
}

func (x *ChunkRequest) GetChunkHash() string {
	if x != nil {
		return x.ChunkHash
	}
	return ""
}

type ChunkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *ChunkResponse) Reset() {
	*x = ChunkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_transfer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkResponse) ProtoMessage() {}

func (x *ChunkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_transfer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkResponse.ProtoReflect.Descriptor instead.
func (*ChunkResponse) Descriptor() ([]byte, []int) {
	return file_file_transfer_proto_rawDescGZIP(), []int{3}
}

func (x *ChunkResponse) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

var File_file_transfer_proto protoreflect.FileDescriptor

var file_file_transfer_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x22, 0xc5, 0x01, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x10, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x6c, 0x0a, 0x0c, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x22, 0x25, 0x0a, 0x0d, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0xa9, 0x01, 0x0a,
	0x13, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x1e, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x12, 0x45, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x1a, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x73, 0x33, 0x2d, 0x74,
	0x61, 0x73, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x3b, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_transfer_proto_rawDescOnce sync.Once
	file_file_transfer_proto_rawDescData = file_file_transfer_proto_rawDesc
)

func file_file_transfer_proto_rawDescGZIP() []byte {
	file_file_transfer_proto_rawDescOnce.Do(func() {
		file_file_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_transfer_proto_rawDescData)
	})
	return file_file_transfer_proto_rawDescData
}

var file_file_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_file_transfer_proto_goTypes = []any{
	(*FileChunk)(nil),        // 0: filetransfer.FileChunk
	(*TransferResponse)(nil), // 1: filetransfer.TransferResponse
	(*ChunkRequest)(nil),     // 2: filetransfer.ChunkRequest
	(*ChunkResponse)(nil),    // 3: filetransfer.ChunkResponse
}
var file_file_transfer_proto_depIdxs = []int32{
	0, // 0: filetransfer.FileTransferService.TransferFile:input_type -> filetransfer.FileChunk
	2, // 1: filetransfer.FileTransferService.GetChunk:input_type -> filetransfer.ChunkRequest
	1, // 2: filetransfer.FileTransferService.TransferFile:output_type -> filetransfer.TransferResponse
	3, // 3: filetransfer.FileTransferService.GetChunk:output_type -> filetransfer.ChunkResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_transfer_proto_init() }
func file_file_transfer_proto_init() {
	if File_file_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_transfer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FileChunk); i {
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
		file_file_transfer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TransferResponse); i {
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
		file_file_transfer_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ChunkRequest); i {
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
		file_file_transfer_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ChunkResponse); i {
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
			RawDescriptor: file_file_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_transfer_proto_goTypes,
		DependencyIndexes: file_file_transfer_proto_depIdxs,
		MessageInfos:      file_file_transfer_proto_msgTypes,
	}.Build()
	File_file_transfer_proto = out.File
	file_file_transfer_proto_rawDesc = nil
	file_file_transfer_proto_goTypes = nil
	file_file_transfer_proto_depIdxs = nil
}
