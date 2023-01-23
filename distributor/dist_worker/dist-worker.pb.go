// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: dist-worker.proto

package dist_worker

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	pb "distributor/pb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_dist_worker_proto protoreflect.FileDescriptor

var file_dist_worker_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x69, 0x73, 0x74, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64, 0x69, 0x73, 0x74, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x1a, 0x10, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2d, 0x64, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x4c, 0x0a, 0x0b, 0x4c, 0x61, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3d, 0x0a, 0x08, 0x52, 0x75, 0x6e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x12, 0x16, 0x2e,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x64, 0x69,
	0x73, 0x74, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x42, 0x0e, 0x5a, 0x0c, 0x64, 0x69, 0x73, 0x74, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_dist_worker_proto_goTypes = []interface{}{
	(*pb.Submission)(nil),  // 0: queue_dist.Submission
	(*pb.JudgeResult)(nil), // 1: queue_dist.JudgeResult
}
var file_dist_worker_proto_depIdxs = []int32{
	0, // 0: dist_worker.LangService.RunJudge:input_type -> queue_dist.Submission
	1, // 1: dist_worker.LangService.RunJudge:output_type -> queue_dist.JudgeResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dist_worker_proto_init() }
func file_dist_worker_proto_init() {
	if File_dist_worker_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dist_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dist_worker_proto_goTypes,
		DependencyIndexes: file_dist_worker_proto_depIdxs,
	}.Build()
	File_dist_worker_proto = out.File
	file_dist_worker_proto_rawDesc = nil
	file_dist_worker_proto_goTypes = nil
	file_dist_worker_proto_depIdxs = nil
}
