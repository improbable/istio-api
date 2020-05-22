// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: type/v1beta1/selector.proto

package v1beta1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"

	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for WorkloadSelector
func (this *WorkloadSelector) MarshalJSON() ([]byte, error) {
	str, err := SelectorMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WorkloadSelector
func (this *WorkloadSelector) UnmarshalJSON(b []byte) error {
	return SelectorUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	SelectorMarshaler	= &github_com_gogo_protobuf_jsonpb.Marshaler{}
	SelectorUnmarshaler	= &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
