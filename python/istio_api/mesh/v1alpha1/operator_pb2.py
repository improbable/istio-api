# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: mesh/v1alpha1/operator.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from mesh.v1alpha1 import config_pb2 as mesh_dot_v1alpha1_dot_config__pb2
from mesh.v1alpha1 import component_pb2 as mesh_dot_v1alpha1_dot_component__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='mesh/v1alpha1/operator.proto',
  package='istio.mesh.v1alpha1',
  syntax='proto3',
  serialized_options=_b('Z\032istio.io/api/mesh/v1alpha1'),
  serialized_pb=_b('\n\x1cmesh/v1alpha1/operator.proto\x12\x13istio.mesh.v1alpha1\x1a\x1amesh/v1alpha1/config.proto\x1a\x1dmesh/v1alpha1/component.proto\"\xca\x06\n\x11IstioOperatorSpec\x12\x0f\n\x07profile\x18\n \x01(\t\x12\x1c\n\x14install_package_path\x18\x0b \x01(\t\x12\x0b\n\x03hub\x18\x0c \x01(\t\x12\x0b\n\x03tag\x18\r \x01(\t\x12\x17\n\x0fresource_suffix\x18\x0e \x01(\t\x12\x34\n\x0bmesh_config\x18( \x01(\x0b\x32\x1f.istio.mesh.v1alpha1.MeshConfig\x12>\n\ncomponents\x18\x32 \x01(\x0b\x32*.istio.mesh.v1alpha1.IstioComponentSetSpec\x12;\n\x06values\x18\x64 \x01(\x0b\x32+.istio.mesh.v1alpha1.TypeMapStringInterface\x12G\n\x12unvalidated_values\x18\x65 \x01(\x0b\x32+.istio.mesh.v1alpha1.TypeMapStringInterface\x12>\n\x06status\x18\xc8\x01 \x01(\x0e\x32-.istio.mesh.v1alpha1.IstioOperatorSpec.Status\x12V\n\x10\x63omponent_status\x18\xc9\x01 \x03(\x0b\x32;.istio.mesh.v1alpha1.IstioOperatorSpec.ComponentStatusEntry\x1a\x85\x01\n\rVersionStatus\x12\x0f\n\x07version\x18\x01 \x01(\t\x12=\n\x06status\x18\x02 \x01(\x0e\x32-.istio.mesh.v1alpha1.IstioOperatorSpec.Status\x12\x15\n\rstatus_string\x18\x03 \x01(\t\x12\r\n\x05\x65rror\x18\x04 \x01(\t\x1al\n\x14\x43omponentStatusEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x43\n\x05value\x18\x02 \x01(\x0b\x32\x34.istio.mesh.v1alpha1.IstioOperatorSpec.VersionStatus:\x02\x38\x01\"I\n\x06Status\x12\x08\n\x04NONE\x10\x00\x12\x0c\n\x08UPDATING\x10\x01\x12\x0f\n\x0bRECONCILING\x10\x02\x12\x0b\n\x07HEALTHY\x10\x03\x12\t\n\x05\x45RROR\x10\x04\x42\x1cZ\x1aistio.io/api/mesh/v1alpha1b\x06proto3')
  ,
  dependencies=[mesh_dot_v1alpha1_dot_config__pb2.DESCRIPTOR,mesh_dot_v1alpha1_dot_component__pb2.DESCRIPTOR,])



_ISTIOOPERATORSPEC_STATUS = _descriptor.EnumDescriptor(
  name='Status',
  full_name='istio.mesh.v1alpha1.IstioOperatorSpec.Status',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UPDATING', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RECONCILING', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HEALTHY', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ERROR', index=4, number=4,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=882,
  serialized_end=955,
)
_sym_db.RegisterEnumDescriptor(_ISTIOOPERATORSPEC_STATUS)


_ISTIOOPERATORSPEC_VERSIONSTATUS = _descriptor.Descriptor(
  name='VersionStatus',
  full_name='istio.mesh.v1alpha1.IstioOperatorSpec.VersionStatus',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='version', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.VersionStatus.version', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='status', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.VersionStatus.status', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='status_string', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.VersionStatus.status_string', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.VersionStatus.error', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=637,
  serialized_end=770,
)

_ISTIOOPERATORSPEC_COMPONENTSTATUSENTRY = _descriptor.Descriptor(
  name='ComponentStatusEntry',
  full_name='istio.mesh.v1alpha1.IstioOperatorSpec.ComponentStatusEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.ComponentStatusEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.ComponentStatusEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=772,
  serialized_end=880,
)

_ISTIOOPERATORSPEC = _descriptor.Descriptor(
  name='IstioOperatorSpec',
  full_name='istio.mesh.v1alpha1.IstioOperatorSpec',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='profile', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.profile', index=0,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='install_package_path', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.install_package_path', index=1,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='hub', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.hub', index=2,
      number=12, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tag', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.tag', index=3,
      number=13, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resource_suffix', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.resource_suffix', index=4,
      number=14, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='mesh_config', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.mesh_config', index=5,
      number=40, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='components', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.components', index=6,
      number=50, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='values', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.values', index=7,
      number=100, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='unvalidated_values', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.unvalidated_values', index=8,
      number=101, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='status', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.status', index=9,
      number=200, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='component_status', full_name='istio.mesh.v1alpha1.IstioOperatorSpec.component_status', index=10,
      number=201, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_ISTIOOPERATORSPEC_VERSIONSTATUS, _ISTIOOPERATORSPEC_COMPONENTSTATUSENTRY, ],
  enum_types=[
    _ISTIOOPERATORSPEC_STATUS,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=113,
  serialized_end=955,
)

_ISTIOOPERATORSPEC_VERSIONSTATUS.fields_by_name['status'].enum_type = _ISTIOOPERATORSPEC_STATUS
_ISTIOOPERATORSPEC_VERSIONSTATUS.containing_type = _ISTIOOPERATORSPEC
_ISTIOOPERATORSPEC_COMPONENTSTATUSENTRY.fields_by_name['value'].message_type = _ISTIOOPERATORSPEC_VERSIONSTATUS
_ISTIOOPERATORSPEC_COMPONENTSTATUSENTRY.containing_type = _ISTIOOPERATORSPEC
_ISTIOOPERATORSPEC.fields_by_name['mesh_config'].message_type = mesh_dot_v1alpha1_dot_config__pb2._MESHCONFIG
_ISTIOOPERATORSPEC.fields_by_name['components'].message_type = mesh_dot_v1alpha1_dot_component__pb2._ISTIOCOMPONENTSETSPEC
_ISTIOOPERATORSPEC.fields_by_name['values'].message_type = mesh_dot_v1alpha1_dot_component__pb2._TYPEMAPSTRINGINTERFACE
_ISTIOOPERATORSPEC.fields_by_name['unvalidated_values'].message_type = mesh_dot_v1alpha1_dot_component__pb2._TYPEMAPSTRINGINTERFACE
_ISTIOOPERATORSPEC.fields_by_name['status'].enum_type = _ISTIOOPERATORSPEC_STATUS
_ISTIOOPERATORSPEC.fields_by_name['component_status'].message_type = _ISTIOOPERATORSPEC_COMPONENTSTATUSENTRY
_ISTIOOPERATORSPEC_STATUS.containing_type = _ISTIOOPERATORSPEC
DESCRIPTOR.message_types_by_name['IstioOperatorSpec'] = _ISTIOOPERATORSPEC
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

IstioOperatorSpec = _reflection.GeneratedProtocolMessageType('IstioOperatorSpec', (_message.Message,), {

  'VersionStatus' : _reflection.GeneratedProtocolMessageType('VersionStatus', (_message.Message,), {
    'DESCRIPTOR' : _ISTIOOPERATORSPEC_VERSIONSTATUS,
    '__module__' : 'mesh.v1alpha1.operator_pb2'
    # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.IstioOperatorSpec.VersionStatus)
    })
  ,

  'ComponentStatusEntry' : _reflection.GeneratedProtocolMessageType('ComponentStatusEntry', (_message.Message,), {
    'DESCRIPTOR' : _ISTIOOPERATORSPEC_COMPONENTSTATUSENTRY,
    '__module__' : 'mesh.v1alpha1.operator_pb2'
    # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.IstioOperatorSpec.ComponentStatusEntry)
    })
  ,
  'DESCRIPTOR' : _ISTIOOPERATORSPEC,
  '__module__' : 'mesh.v1alpha1.operator_pb2'
  # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.IstioOperatorSpec)
  })
_sym_db.RegisterMessage(IstioOperatorSpec)
_sym_db.RegisterMessage(IstioOperatorSpec.VersionStatus)
_sym_db.RegisterMessage(IstioOperatorSpec.ComponentStatusEntry)


DESCRIPTOR._options = None
_ISTIOOPERATORSPEC_COMPONENTSTATUSENTRY._options = None
# @@protoc_insertion_point(module_scope)
