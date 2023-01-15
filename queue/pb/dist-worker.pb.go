// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: dist-worker.proto

package pb

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

type Language int32

const (
	Language_nodejs Language = 0
	Language_ruby   Language = 1
	Language_python Language = 2
	Language_cpp    Language = 3
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "nodejs",
		1: "ruby",
		2: "python",
		3: "cpp",
	}
	Language_value = map[string]int32{
		"nodejs": 0,
		"ruby":   1,
		"python": 2,
		"cpp":    3,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_dist_worker_proto_enumTypes[0].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_dist_worker_proto_enumTypes[0]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_dist_worker_proto_rawDescGZIP(), []int{0}
}

type JudgeResult int32

const (
	JudgeResult_ac  JudgeResult = 0
	JudgeResult_wa  JudgeResult = 1
	JudgeResult_ce  JudgeResult = 2
	JudgeResult_re  JudgeResult = 3
	JudgeResult_mle JudgeResult = 4
)

// Enum value maps for JudgeResult.
var (
	JudgeResult_name = map[int32]string{
		0: "ac",
		1: "wa",
		2: "ce",
		3: "re",
		4: "mle",
	}
	JudgeResult_value = map[string]int32{
		"ac":  0,
		"wa":  1,
		"ce":  2,
		"re":  3,
		"mle": 4,
	}
)

func (x JudgeResult) Enum() *JudgeResult {
	p := new(JudgeResult)
	*p = x
	return p
}

func (x JudgeResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JudgeResult) Descriptor() protoreflect.EnumDescriptor {
	return file_dist_worker_proto_enumTypes[1].Descriptor()
}

func (JudgeResult) Type() protoreflect.EnumType {
	return &file_dist_worker_proto_enumTypes[1]
}

func (x JudgeResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JudgeResult.Descriptor instead.
func (JudgeResult) EnumDescriptor() ([]byte, []int) {
	return file_dist_worker_proto_rawDescGZIP(), []int{1}
}

type Submission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language Language `protobuf:"varint,1,opt,name=language,proto3,enum=Language" json:"language,omitempty"`
	Code     string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Id       uint64   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Submission) Reset() {
	*x = Submission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dist_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Submission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submission) ProtoMessage() {}

func (x *Submission) ProtoReflect() protoreflect.Message {
	mi := &file_dist_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Submission.ProtoReflect.Descriptor instead.
func (*Submission) Descriptor() ([]byte, []int) {
	return file_dist_worker_proto_rawDescGZIP(), []int{0}
}

func (x *Submission) GetLanguage() Language {
	if x != nil {
		return x.Language
	}
	return Language_nodejs
}

func (x *Submission) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Submission) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode JudgeResult `protobuf:"varint,1,opt,name=result_code,json=resultCode,proto3,enum=JudgeResult" json:"result_code,omitempty"`
	Submission *Submission `protobuf:"bytes,2,opt,name=submission,proto3" json:"submission,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dist_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_dist_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_dist_worker_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetResultCode() JudgeResult {
	if x != nil {
		return x.ResultCode
	}
	return JudgeResult_ac
}

func (x *Result) GetSubmission() *Submission {
	if x != nil {
		return x.Submission
	}
	return nil
}

var File_dist_worker_proto protoreflect.FileDescriptor

var file_dist_worker_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x69, 0x73, 0x74, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x64, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x4a, 0x75,
	0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2b, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2a, 0x35, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x6a, 0x73, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x72, 0x75,
	0x62, 0x79, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x10, 0x02,
	0x12, 0x07, 0x0a, 0x03, 0x63, 0x70, 0x70, 0x10, 0x03, 0x2a, 0x36, 0x0a, 0x0b, 0x4a, 0x75, 0x64,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x61, 0x63, 0x10, 0x00,
	0x12, 0x06, 0x0a, 0x02, 0x77, 0x61, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x63, 0x65, 0x10, 0x02,
	0x12, 0x06, 0x0a, 0x02, 0x72, 0x65, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x6d, 0x6c, 0x65, 0x10,
	0x04, 0x42, 0x05, 0x5a, 0x03, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dist_worker_proto_rawDescOnce sync.Once
	file_dist_worker_proto_rawDescData = file_dist_worker_proto_rawDesc
)

func file_dist_worker_proto_rawDescGZIP() []byte {
	file_dist_worker_proto_rawDescOnce.Do(func() {
		file_dist_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_dist_worker_proto_rawDescData)
	})
	return file_dist_worker_proto_rawDescData
}

var file_dist_worker_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_dist_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dist_worker_proto_goTypes = []interface{}{
	(Language)(0),      // 0: Language
	(JudgeResult)(0),   // 1: JudgeResult
	(*Submission)(nil), // 2: Submission
	(*Result)(nil),     // 3: Result
}
var file_dist_worker_proto_depIdxs = []int32{
	0, // 0: Submission.language:type_name -> Language
	1, // 1: Result.result_code:type_name -> JudgeResult
	2, // 2: Result.submission:type_name -> Submission
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_dist_worker_proto_init() }
func file_dist_worker_proto_init() {
	if File_dist_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dist_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Submission); i {
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
		file_dist_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_dist_worker_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dist_worker_proto_goTypes,
		DependencyIndexes: file_dist_worker_proto_depIdxs,
		EnumInfos:         file_dist_worker_proto_enumTypes,
		MessageInfos:      file_dist_worker_proto_msgTypes,
	}.Build()
	File_dist_worker_proto = out.File
	file_dist_worker_proto_rawDesc = nil
	file_dist_worker_proto_goTypes = nil
	file_dist_worker_proto_depIdxs = nil
}
