// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/v1/mixer.proto

// This package defines the Mixer API that the sidecar proxy uses to perform
// precondition checks, manage quotas, and report telemetry.

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"

	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for CheckRequest
func (this *CheckRequest) MarshalJSON() ([]byte, error) {
	str, err := MixerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CheckRequest
func (this *CheckRequest) UnmarshalJSON(b []byte) error {
	return MixerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CheckRequest_QuotaParams
func (this *CheckRequest_QuotaParams) MarshalJSON() ([]byte, error) {
	str, err := MixerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CheckRequest_QuotaParams
func (this *CheckRequest_QuotaParams) UnmarshalJSON(b []byte) error {
	return MixerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CheckResponse
func (this *CheckResponse) MarshalJSON() ([]byte, error) {
	str, err := MixerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CheckResponse
func (this *CheckResponse) UnmarshalJSON(b []byte) error {
	return MixerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CheckResponse_PreconditionResult
func (this *CheckResponse_PreconditionResult) MarshalJSON() ([]byte, error) {
	str, err := MixerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CheckResponse_PreconditionResult
func (this *CheckResponse_PreconditionResult) UnmarshalJSON(b []byte) error {
	return MixerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CheckResponse_QuotaResult
func (this *CheckResponse_QuotaResult) MarshalJSON() ([]byte, error) {
	str, err := MixerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CheckResponse_QuotaResult
func (this *CheckResponse_QuotaResult) UnmarshalJSON(b []byte) error {
	return MixerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ReferencedAttributes
func (this *ReferencedAttributes) MarshalJSON() ([]byte, error) {
	str, err := MixerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ReferencedAttributes
func (this *ReferencedAttributes) UnmarshalJSON(b []byte) error {
	return MixerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ReferencedAttributes_AttributeMatch
func (this *ReferencedAttributes_AttributeMatch) MarshalJSON() ([]byte, error) {
	str, err := MixerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ReferencedAttributes_AttributeMatch
func (this *ReferencedAttributes_AttributeMatch) UnmarshalJSON(b []byte) error {
	return MixerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HeaderOperation
func (this *HeaderOperation) MarshalJSON() ([]byte, error) {
	str, err := MixerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HeaderOperation
func (this *HeaderOperation) UnmarshalJSON(b []byte) error {
	return MixerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RouteDirective
func (this *RouteDirective) MarshalJSON() ([]byte, error) {
	str, err := MixerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteDirective
func (this *RouteDirective) UnmarshalJSON(b []byte) error {
	return MixerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ReportRequest
func (this *ReportRequest) MarshalJSON() ([]byte, error) {
	str, err := MixerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ReportRequest
func (this *ReportRequest) UnmarshalJSON(b []byte) error {
	return MixerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ReportResponse
func (this *ReportResponse) MarshalJSON() ([]byte, error) {
	str, err := MixerMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ReportResponse
func (this *ReportResponse) UnmarshalJSON(b []byte) error {
	return MixerUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	MixerMarshaler		= &github_com_gogo_protobuf_jsonpb.Marshaler{}
	MixerUnmarshaler	= &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
