// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/v1/config/client/client_config.proto

// Describes the configuration state for the Mixer client library that's built into Envoy.

package client

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/improbable/istio-api/mixer/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for NetworkFailPolicy
func (this *NetworkFailPolicy) MarshalJSON() ([]byte, error) {
	str, err := ClientConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for NetworkFailPolicy
func (this *NetworkFailPolicy) UnmarshalJSON(b []byte) error {
	return ClientConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ServiceConfig
func (this *ServiceConfig) MarshalJSON() ([]byte, error) {
	str, err := ClientConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ServiceConfig
func (this *ServiceConfig) UnmarshalJSON(b []byte) error {
	return ClientConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TransportConfig
func (this *TransportConfig) MarshalJSON() ([]byte, error) {
	str, err := ClientConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TransportConfig
func (this *TransportConfig) UnmarshalJSON(b []byte) error {
	return ClientConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HttpClientConfig
func (this *HttpClientConfig) MarshalJSON() ([]byte, error) {
	str, err := ClientConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HttpClientConfig
func (this *HttpClientConfig) UnmarshalJSON(b []byte) error {
	return ClientConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TcpClientConfig
func (this *TcpClientConfig) MarshalJSON() ([]byte, error) {
	str, err := ClientConfigMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TcpClientConfig
func (this *TcpClientConfig) UnmarshalJSON(b []byte) error {
	return ClientConfigUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ClientConfigMarshaler	= &github_com_gogo_protobuf_jsonpb.Marshaler{}
	ClientConfigUnmarshaler	= &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
