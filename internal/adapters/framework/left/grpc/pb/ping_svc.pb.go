// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: ping_svc.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_ping_svc_proto protoreflect.FileDescriptor

var file_ping_svc_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x0e, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x3a, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ping_svc_proto_goTypes = []interface{}{
	(*PingRequest)(nil),  // 0: pb.PingRequest
	(*PingResponse)(nil), // 1: pb.PingResponse
}
var file_ping_svc_proto_depIdxs = []int32{
	0, // 0: pb.PingService.Ping:input_type -> pb.PingRequest
	1, // 1: pb.PingService.Ping:output_type -> pb.PingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ping_svc_proto_init() }
func file_ping_svc_proto_init() {
	if File_ping_svc_proto != nil {
		return
	}
	file_ping_msg_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ping_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ping_svc_proto_goTypes,
		DependencyIndexes: file_ping_svc_proto_depIdxs,
	}.Build()
	File_ping_svc_proto = out.File
	file_ping_svc_proto_rawDesc = nil
	file_ping_svc_proto_goTypes = nil
	file_ping_svc_proto_depIdxs = nil
}
