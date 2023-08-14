// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.5
// source: github.com/cloudprober/cloudprober/validators/http/proto/config.proto

package proto

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

// HTTP validator configuration. For HTTP validator to succeed, all conditions
// specified in the validator should succeed. Note that failures conditions are
// evaluated before success conditions.
type Validator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Comma-separated list of success status codes and code ranges.
	// Example: success_stauts_codes: 200-299,301,302
	SuccessStatusCodes *string `protobuf:"bytes,1,opt,name=success_status_codes,json=successStatusCodes,proto3,oneof" json:"success_status_codes,omitempty"`
	// Comma-separated list of failure status codes and code ranges. If HTTP
	// status code matches failure_status_codes, validator fails.
	FailureStatusCodes *string `protobuf:"bytes,2,opt,name=failure_status_codes,json=failureStatusCodes,proto3,oneof" json:"failure_status_codes,omitempty"`
	// Header based validations.
	// TODO(manugarg): Add support for specifying multiple success and failure
	// headers.
	//
	// Success Header:
	//
	//	If specified, HTTP response headers should match the success_header for
	//	validation to succeed. Example:
	//	  success_header: {
	//	    name: "Strict-Transport-Security"
	//	    value_regex: "max-age=31536000"
	//	  }
	SuccessHeader *Validator_Header `protobuf:"bytes,3,opt,name=success_header,json=successHeader,proto3,oneof" json:"success_header,omitempty"`
	// Failure Header:
	//
	//	If HTTP response headers match failure_header, validation fails.
	FailureHeader *Validator_Header `protobuf:"bytes,4,opt,name=failure_header,json=failureHeader,proto3,oneof" json:"failure_header,omitempty"`
}

func (x *Validator) Reset() {
	*x = Validator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validator) ProtoMessage() {}

func (x *Validator) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Validator.ProtoReflect.Descriptor instead.
func (*Validator) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *Validator) GetSuccessStatusCodes() string {
	if x != nil && x.SuccessStatusCodes != nil {
		return *x.SuccessStatusCodes
	}
	return ""
}

func (x *Validator) GetFailureStatusCodes() string {
	if x != nil && x.FailureStatusCodes != nil {
		return *x.FailureStatusCodes
	}
	return ""
}

func (x *Validator) GetSuccessHeader() *Validator_Header {
	if x != nil {
		return x.SuccessHeader
	}
	return nil
}

func (x *Validator) GetFailureHeader() *Validator_Header {
	if x != nil {
		return x.FailureHeader
	}
	return nil
}

type Validator_Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Header name to look for
	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Header value to match. If omited - check for header existence
	ValueRegex *string `protobuf:"bytes,2,opt,name=value_regex,json=valueRegex,proto3,oneof" json:"value_regex,omitempty"`
}

func (x *Validator_Header) Reset() {
	*x = Validator_Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validator_Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validator_Header) ProtoMessage() {}

func (x *Validator_Header) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Validator_Header.ProtoReflect.Descriptor instead.
func (*Validator_Header) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Validator_Header) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Validator_Header) GetValueRegex() string {
	if x != nil && x.ValueRegex != nil {
		return *x.ValueRegex
	}
	return ""
}

var File_github_com_cloudprober_cloudprober_validators_http_proto_config_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x22, 0xe9, 0x03, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x35, 0x0a, 0x14, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x12, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x14, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x12, 0x66, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x59, 0x0a, 0x0e, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x02, 0x52, 0x0d, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x59, 0x0a, 0x0e, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x72, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x48, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x1a, 0x60, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x67, 0x65, 0x78, 0x88, 0x01, 0x01, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x73, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x11, 0x0a,
	0x0f, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_rawDescData = file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_goTypes = []interface{}{
	(*Validator)(nil),        // 0: cloudprober.validators.http.Validator
	(*Validator_Header)(nil), // 1: cloudprober.validators.http.Validator.Header
}
var file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_depIdxs = []int32{
	1, // 0: cloudprober.validators.http.Validator.success_header:type_name -> cloudprober.validators.http.Validator.Header
	1, // 1: cloudprober.validators.http.Validator.failure_header:type_name -> cloudprober.validators.http.Validator.Header
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_init() }
func file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_init() {
	if File_github_com_cloudprober_cloudprober_validators_http_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Validator); i {
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
		file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Validator_Header); i {
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
	file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_depIdxs,
		MessageInfos:      file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_validators_http_proto_config_proto = out.File
	file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_validators_http_proto_config_proto_depIdxs = nil
}
