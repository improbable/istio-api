// Code generated by protoc-gen-go. DO NOT EDIT.
// source: operator/v1alpha1/operator.proto

// Configuration affecting Istio control plane installation version and shape.

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1alpha1 "istio.io/api/mesh/v1alpha1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3	// please upgrade the proto package

// Status describes the current state of a component.
type InstallStatus_Status int32

const (
	// Component is not present.
	InstallStatus_NONE	InstallStatus_Status	= 0
	// Component is being updated to a different version.
	InstallStatus_UPDATING	InstallStatus_Status	= 1
	// Controller has started but not yet completed reconciliation loop for the component.
	InstallStatus_RECONCILING	InstallStatus_Status	= 2
	// Component is healthy.
	InstallStatus_HEALTHY	InstallStatus_Status	= 3
	// Component is in an error state.
	InstallStatus_ERROR	InstallStatus_Status	= 4
)

var InstallStatus_Status_name = map[int32]string{
	0:	"NONE",
	1:	"UPDATING",
	2:	"RECONCILING",
	3:	"HEALTHY",
	4:	"ERROR",
}

var InstallStatus_Status_value = map[string]int32{
	"NONE":		0,
	"UPDATING":	1,
	"RECONCILING":	2,
	"HEALTHY":	3,
	"ERROR":	4,
}

func (x InstallStatus_Status) String() string {
	return proto.EnumName(InstallStatus_Status_name, int32(x))
}

func (InstallStatus_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8023ebf2dcfea843, []int{1, 0}
}

// IstioOperatorSpec defines the desired installed state of Istio components.
// The spec is a used to define a customization of the default profile values that are supplied with each Istio release.
// Because the spec is a customization API, specifying an empty IstioOperatorSpec results in a default Istio
// component values.
type IstioOperatorSpec struct {
	// Path or name for the profile e.g.
	//     - minimal (looks in profiles dir for a file called minimal.yaml)
	//     - /tmp/istio/install/values/custom/custom-install.yaml (local file path)
	// default profile is used if this field is unset.
	Profile	string	`protobuf:"bytes,10,opt,name=profile,proto3" json:"profile,omitempty"`
	// Path for the install package. e.g.
	//     - /tmp/istio-installer/nightly (local file path)
	InstallPackagePath	string	`protobuf:"bytes,11,opt,name=install_package_path,json=installPackagePath,proto3" json:"installPackagePath,omitempty"`
	// Root for docker image paths e.g. docker.io/istio
	Hub	string	`protobuf:"bytes,12,opt,name=hub,proto3" json:"hub,omitempty"`
	// Version tag for docker images e.g. 1.0.6
	Tag	interface{}	`protobuf:"bytes,13,opt,name=tag,proto3" json:"tag,omitempty"`
	// $hide_from_docs
	// Resource suffix is appended to all resources installed by each component.
	// Never implemented; replaced by revision.
	ResourceSuffix	string	`protobuf:"bytes,14,opt,name=resource_suffix,json=resourceSuffix,proto3" json:"resourceSuffix,omitempty"`	// Deprecated: Do not use.
	// Namespace to install control plane resources into. If unset, Istio will be installed into the same namespace
	// as the IstioOperator CR.
	Namespace	string	`protobuf:"bytes,15,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Identify the revision this installation is associated with.
	// This option is currently experimental.
	Revision	string	`protobuf:"bytes,16,opt,name=revision,proto3" json:"revision,omitempty"`
	// Config used by control plane components internally.
	MeshConfig	*v1alpha1.MeshConfig	`protobuf:"bytes,40,opt,name=mesh_config,json=meshConfig,proto3" json:"meshConfig,omitempty"`
	// Kubernetes resource settings, enablement and component-specific settings that are not internal to the
	// component.
	Components	*IstioComponentSetSpec	`protobuf:"bytes,50,opt,name=components,proto3" json:"components,omitempty"`
	// Extra addon components which are not explicitly specified above.
	AddonComponents	map[string]*ExternalComponentSpec	`protobuf:"bytes,51,rep,name=addon_components,json=addonComponents,proto3" json:"addonComponents,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Overrides for default values.yaml. This is a validated pass-through to Helm templates.
	// See the Helm installation options for schema details: https://istio.io/docs/reference/config/installation-options/.
	// Anything that is available in IstioOperatorSpec should be set above rather than using the passthrough. This
	// includes Kubernetes resource settings for components in KubernetesResourcesSpec.
	Values	map[string]interface{}	`protobuf:"bytes,100,opt,name=values,proto3" json:"values,omitempty"`
	// Unvalidated overrides for default values.yaml. Used for custom templates where new parameters are added.
	UnvalidatedValues	map[string]interface{}	`protobuf:"bytes,101,opt,name=unvalidated_values,json=unvalidatedValues,proto3" json:"unvalidatedValues,omitempty"`
	XXX_NoUnkeyedLiteral	struct{}		`json:"-"`
	XXX_unrecognized	[]byte			`json:"-"`
	XXX_sizecache		int32			`json:"-"`
}

func (m *IstioOperatorSpec) Reset()		{ *m = IstioOperatorSpec{} }
func (m *IstioOperatorSpec) String() string	{ return proto.CompactTextString(m) }
func (*IstioOperatorSpec) ProtoMessage()	{}
func (*IstioOperatorSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_8023ebf2dcfea843, []int{0}
}

func (m *IstioOperatorSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IstioOperatorSpec.Unmarshal(m, b)
}
func (m *IstioOperatorSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IstioOperatorSpec.Marshal(b, m, deterministic)
}
func (m *IstioOperatorSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioOperatorSpec.Merge(m, src)
}
func (m *IstioOperatorSpec) XXX_Size() int {
	return xxx_messageInfo_IstioOperatorSpec.Size(m)
}
func (m *IstioOperatorSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioOperatorSpec.DiscardUnknown(m)
}

var xxx_messageInfo_IstioOperatorSpec proto.InternalMessageInfo

func (m *IstioOperatorSpec) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *IstioOperatorSpec) GetInstallPackagePath() string {
	if m != nil {
		return m.InstallPackagePath
	}
	return ""
}

func (m *IstioOperatorSpec) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}

// Deprecated: Do not use.
func (m *IstioOperatorSpec) GetResourceSuffix() string {
	if m != nil {
		return m.ResourceSuffix
	}
	return ""
}

func (m *IstioOperatorSpec) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *IstioOperatorSpec) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *IstioOperatorSpec) GetMeshConfig() *v1alpha1.MeshConfig {
	if m != nil {
		return m.MeshConfig
	}
	return nil
}

func (m *IstioOperatorSpec) GetComponents() *IstioComponentSetSpec {
	if m != nil {
		return m.Components
	}
	return nil
}

func (m *IstioOperatorSpec) GetAddonComponents() map[string]*ExternalComponentSpec {
	if m != nil {
		return m.AddonComponents
	}
	return nil
}

// Observed state of IstioOperator
type InstallStatus struct {
	// Overall status of all components controlled by the operator.
	// - If all components have status NONE, overall status is NONE.
	// - If all components are HEALTHY, overall status is HEALTHY.
	// - If one or more components are RECONCILING and others are HEALTHY, overall status is RECONCILING.
	// - If one or more components are UPDATING and others are HEALTHY, overall status is UPDATING.
	// - If components are a mix of RECONCILING, UPDATING and HEALTHY, overall status is UPDATING.
	// - If any component is in ERROR state, overall status is ERROR.
	Status	InstallStatus_Status	`protobuf:"varint,1,opt,name=status,proto3,enum=istio.operator.v1alpha1.InstallStatus_Status" json:"status,omitempty"`
	// Individual status of each component controlled by the operator. The map key is the name of the component.
	ComponentStatus		map[string]*InstallStatus_VersionStatus	`protobuf:"bytes,2,rep,name=component_status,json=componentStatus,proto3" json:"componentStatus,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral	struct{}				`json:"-"`
	XXX_unrecognized	[]byte					`json:"-"`
	XXX_sizecache		int32					`json:"-"`
}

func (m *InstallStatus) Reset()		{ *m = InstallStatus{} }
func (m *InstallStatus) String() string	{ return proto.CompactTextString(m) }
func (*InstallStatus) ProtoMessage()	{}
func (*InstallStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_8023ebf2dcfea843, []int{1}
}

func (m *InstallStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallStatus.Unmarshal(m, b)
}
func (m *InstallStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallStatus.Marshal(b, m, deterministic)
}
func (m *InstallStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallStatus.Merge(m, src)
}
func (m *InstallStatus) XXX_Size() int {
	return xxx_messageInfo_InstallStatus.Size(m)
}
func (m *InstallStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallStatus.DiscardUnknown(m)
}

var xxx_messageInfo_InstallStatus proto.InternalMessageInfo

func (m *InstallStatus) GetStatus() InstallStatus_Status {
	if m != nil {
		return m.Status
	}
	return InstallStatus_NONE
}

func (m *InstallStatus) GetComponentStatus() map[string]*InstallStatus_VersionStatus {
	if m != nil {
		return m.ComponentStatus
	}
	return nil
}

// VersionStatus is the status and version of a component.
type InstallStatus_VersionStatus struct {
	Version			string			`protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Status			InstallStatus_Status	`protobuf:"varint,2,opt,name=status,proto3,enum=istio.operator.v1alpha1.InstallStatus_Status" json:"status,omitempty"`
	StatusString		string			`protobuf:"bytes,3,opt,name=status_string,json=statusString,proto3" json:"statusString,omitempty"`
	Error			string			`protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral	struct{}		`json:"-"`
	XXX_unrecognized	[]byte			`json:"-"`
	XXX_sizecache		int32			`json:"-"`
}

func (m *InstallStatus_VersionStatus) Reset()		{ *m = InstallStatus_VersionStatus{} }
func (m *InstallStatus_VersionStatus) String() string	{ return proto.CompactTextString(m) }
func (*InstallStatus_VersionStatus) ProtoMessage()	{}
func (*InstallStatus_VersionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_8023ebf2dcfea843, []int{1, 0}
}

func (m *InstallStatus_VersionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallStatus_VersionStatus.Unmarshal(m, b)
}
func (m *InstallStatus_VersionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallStatus_VersionStatus.Marshal(b, m, deterministic)
}
func (m *InstallStatus_VersionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallStatus_VersionStatus.Merge(m, src)
}
func (m *InstallStatus_VersionStatus) XXX_Size() int {
	return xxx_messageInfo_InstallStatus_VersionStatus.Size(m)
}
func (m *InstallStatus_VersionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallStatus_VersionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_InstallStatus_VersionStatus proto.InternalMessageInfo

func (m *InstallStatus_VersionStatus) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *InstallStatus_VersionStatus) GetStatus() InstallStatus_Status {
	if m != nil {
		return m.Status
	}
	return InstallStatus_NONE
}

func (m *InstallStatus_VersionStatus) GetStatusString() string {
	if m != nil {
		return m.StatusString
	}
	return ""
}

func (m *InstallStatus_VersionStatus) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// This is required because synthetic type definition has file rather than package scope.

func init() {
	proto.RegisterEnum("istio.operator.v1alpha1.InstallStatus_Status", InstallStatus_Status_name, InstallStatus_Status_value)
	proto.RegisterType((*IstioOperatorSpec)(nil), "istio.operator.v1alpha1.IstioOperatorSpec")
	proto.RegisterMapType((map[string]*ExternalComponentSpec)(nil), "istio.operator.v1alpha1.IstioOperatorSpec.AddonComponentsEntry")
	proto.RegisterType((*InstallStatus)(nil), "istio.operator.v1alpha1.InstallStatus")
	proto.RegisterMapType((map[string]*InstallStatus_VersionStatus)(nil), "istio.operator.v1alpha1.InstallStatus.ComponentStatusEntry")
	proto.RegisterType((*InstallStatus_VersionStatus)(nil), "istio.operator.v1alpha1.InstallStatus.VersionStatus")
}

func init()	{ proto.RegisterFile("operator/v1alpha1/operator.proto", fileDescriptor_8023ebf2dcfea843) }

var fileDescriptor_8023ebf2dcfea843 = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5b, 0x6f, 0xda, 0x4a,
	0x10, 0x3e, 0x5c, 0x42, 0x92, 0x21, 0x80, 0xb3, 0x42, 0xca, 0x1e, 0x74, 0x74, 0x4a, 0xe9, 0x43,
	0x91, 0xaa, 0x9a, 0x84, 0xf4, 0xa1, 0x97, 0x87, 0x36, 0x21, 0x56, 0x43, 0x95, 0x40, 0x64, 0xd2,
	0x48, 0xed, 0x8b, 0xb5, 0x31, 0x0b, 0xb8, 0x31, 0x5e, 0x6b, 0x77, 0x41, 0xe1, 0x37, 0xf5, 0xa5,
	0xfd, 0x87, 0x95, 0x77, 0x0d, 0x18, 0x25, 0xa4, 0x51, 0xfb, 0xe4, 0x9d, 0x6f, 0x66, 0xbe, 0xb9,
	0x7c, 0x23, 0x43, 0x95, 0x85, 0x94, 0x13, 0xc9, 0x78, 0x63, 0x7a, 0x40, 0xfc, 0x70, 0x44, 0x0e,
	0x1a, 0x73, 0xc4, 0x0c, 0x39, 0x93, 0x0c, 0xed, 0x79, 0x42, 0x7a, 0xcc, 0x5c, 0xa0, 0xf3, 0xb8,
	0x4a, 0x65, 0x4c, 0xc5, 0x68, 0x99, 0xe6, 0xb2, 0x60, 0xe0, 0x0d, 0x75, 0x52, 0xe5, 0xe9, 0x5d,
	0x5a, 0x97, 0x8d, 0x43, 0x16, 0xd0, 0x40, 0xea, 0x90, 0xda, 0x8f, 0x1c, 0xec, 0xb6, 0x23, 0xea,
	0x6e, 0x1c, 0xda, 0x0b, 0xa9, 0x8b, 0x30, 0x6c, 0x86, 0x9c, 0x0d, 0x3c, 0x9f, 0x62, 0xa8, 0xa6,
	0xea, 0xdb, 0xf6, 0xdc, 0x44, 0xfb, 0x50, 0xf6, 0x02, 0x21, 0x89, 0xef, 0x3b, 0x21, 0x71, 0x6f,
	0xc8, 0x90, 0x3a, 0x21, 0x91, 0x23, 0x9c, 0x57, 0x61, 0x28, 0xf6, 0x5d, 0x68, 0xd7, 0x05, 0x91,
	0x23, 0x64, 0x40, 0x66, 0x34, 0xb9, 0xc6, 0x3b, 0x2a, 0x20, 0x7a, 0xa2, 0x37, 0x90, 0x91, 0x64,
	0x88, 0x0b, 0xd5, 0x54, 0x3d, 0xdf, 0x7c, 0x6e, 0xae, 0x99, 0xcc, 0xbc, 0x9c, 0x85, 0xb4, 0x1d,
	0x48, 0xca, 0x07, 0xc4, 0xa5, 0x4d, 0x3b, 0xca, 0x41, 0x2f, 0xa0, 0xc4, 0xa9, 0x60, 0x13, 0xee,
	0x52, 0x47, 0x4c, 0x06, 0x03, 0xef, 0x16, 0x17, 0x23, 0xe2, 0xe3, 0x34, 0x4e, 0xd9, 0xc5, 0xb9,
	0xab, 0xa7, 0x3c, 0xe8, 0x3f, 0xd8, 0x0e, 0xc8, 0x98, 0x8a, 0x90, 0xb8, 0x14, 0x97, 0x54, 0xfd,
	0x25, 0x80, 0x2a, 0xb0, 0xc5, 0xe9, 0xd4, 0x13, 0x1e, 0x0b, 0xb0, 0xa1, 0x9c, 0x0b, 0x1b, 0x7d,
	0x80, 0x7c, 0xb4, 0x56, 0x47, 0x6f, 0x13, 0xd7, 0x55, 0xa7, 0x4f, 0xe2, 0x4e, 0x23, 0xcf, 0xb2,
	0xcb, 0x73, 0x2a, 0x46, 0x2d, 0x15, 0x66, 0xc3, 0x78, 0xf1, 0x46, 0x1d, 0x80, 0xc5, 0xaa, 0x05,
	0x6e, 0x2a, 0x02, 0x73, 0xed, 0xa8, 0x4a, 0x81, 0xd6, 0x3c, 0xbe, 0x47, 0x65, 0xa4, 0x82, 0x9d,
	0x60, 0x40, 0xdf, 0xc0, 0x20, 0xfd, 0x3e, 0x0b, 0x9c, 0x04, 0xeb, 0x61, 0x35, 0x53, 0xcf, 0x37,
	0xdf, 0x3f, 0xcc, 0x9a, 0xd4, 0xd5, 0x3c, 0x8a, 0x28, 0x16, 0x75, 0x84, 0x15, 0x48, 0x3e, 0xb3,
	0x4b, 0x64, 0x15, 0x45, 0xa7, 0x90, 0x9b, 0x12, 0x7f, 0x42, 0x05, 0xee, 0xab, 0xbe, 0xf7, 0x1f,
	0x94, 0xe8, 0x9c, 0x84, 0x3d, 0xc9, 0xbd, 0x60, 0x98, 0xd0, 0x2a, 0xce, 0x47, 0x0e, 0xa0, 0x49,
	0x30, 0x25, 0xbe, 0xd7, 0x27, 0x92, 0xf6, 0x9d, 0x98, 0x95, 0xfe, 0x21, 0xeb, 0x6e, 0x82, 0xeb,
	0x4a, 0x51, 0x55, 0x38, 0x94, 0xef, 0x9b, 0x29, 0x3a, 0xba, 0x1b, 0x3a, 0xc3, 0x29, 0x7d, 0x74,
	0x37, 0x74, 0x86, 0x4e, 0x60, 0x43, 0x95, 0xc7, 0xe9, 0xdf, 0x68, 0x61, 0xdd, 0x4a, 0xca, 0x03,
	0xe2, 0x2f, 0xe5, 0x88, 0xb4, 0xd0, 0xc9, 0x6f, 0xd3, 0xaf, 0x53, 0xb5, 0x9f, 0x59, 0x28, 0xb4,
	0xf5, 0x9d, 0xf7, 0x24, 0x91, 0x13, 0x81, 0x2c, 0xc8, 0x09, 0xf5, 0x52, 0x05, 0x8b, 0xcd, 0x97,
	0xeb, 0x25, 0x49, 0xe6, 0x99, 0xfa, 0x63, 0xc7, 0xc9, 0x68, 0x00, 0xc6, 0x42, 0x5d, 0x27, 0x26,
	0x4c, 0x2b, 0x8d, 0xdf, 0x3d, 0x92, 0x70, 0xd9, 0xb3, 0xb2, 0x63, 0x7d, 0xdd, 0x55, 0xb4, 0xf2,
	0x3d, 0x05, 0x85, 0x2b, 0xca, 0xa3, 0x4b, 0x8f, 0x07, 0xc0, 0xb0, 0x39, 0xd5, 0x40, 0xbc, 0xb2,
	0xb9, 0x99, 0x18, 0x2d, 0xfd, 0x37, 0xa3, 0x3d, 0x83, 0x82, 0x7e, 0x39, 0x42, 0xa9, 0x8a, 0x33,
	0xaa, 0xcc, 0x8e, 0x06, 0xb5, 0xd2, 0xa8, 0x0c, 0x1b, 0x94, 0x73, 0xc6, 0x71, 0x56, 0x39, 0xb5,
	0x51, 0xb9, 0x85, 0xf2, 0x7d, 0x63, 0xdd, 0x23, 0xf1, 0xa7, 0x55, 0x89, 0x5f, 0x3d, 0xb2, 0xd5,
	0x95, 0x55, 0x24, 0x85, 0x6e, 0x43, 0x2e, 0xde, 0xcf, 0x16, 0x64, 0x3b, 0xdd, 0x8e, 0x65, 0xfc,
	0x83, 0x76, 0x60, 0xeb, 0xf3, 0xc5, 0xc9, 0xd1, 0x65, 0xbb, 0xf3, 0xd1, 0x48, 0xa1, 0x12, 0xe4,
	0x6d, 0xab, 0xd5, 0xed, 0xb4, 0xda, 0x67, 0x11, 0x90, 0x46, 0x79, 0xd8, 0x3c, 0xb5, 0x8e, 0xce,
	0x2e, 0x4f, 0xbf, 0x18, 0x19, 0xb4, 0x0d, 0x1b, 0x96, 0x6d, 0x77, 0x6d, 0x23, 0x5b, 0xfb, 0x17,
	0xf6, 0xd6, 0x5c, 0x75, 0xcd, 0x80, 0xe2, 0xea, 0x9f, 0xee, 0xb8, 0xfa, 0xf5, 0x7f, 0xdd, 0xb9,
	0xc7, 0x1a, 0x24, 0xf4, 0x1a, 0x77, 0xfe, 0xe2, 0xd7, 0x39, 0xf5, 0xf3, 0x3e, 0xfc, 0x15, 0x00,
	0x00, 0xff, 0xff, 0x79, 0xdf, 0xba, 0xde, 0x38, 0x06, 0x00, 0x00,
}
