# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: certs/certs.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from evecommon import evecommon_pb2 as evecommon_dot_evecommon__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='certs/certs.proto',
  package='org.lfedge.eve.certs',
  syntax='proto3',
  serialized_options=_b('\n\024org.lfedge.eve.certsZ#github.com/lf-edge/eve/api/go/certs'),
  serialized_pb=_b('\n\x11\x63\x65rts/certs.proto\x12\x14org.lfedge.eve.certs\x1a\x19\x65vecommon/evecommon.proto\"=\n\x0fZControllerCert\x12*\n\x05\x63\x65rts\x18\x01 \x03(\x0b\x32\x1b.org.lfedge.eve.certs.ZCert\"Y\n\rZCertMetaData\x12\x35\n\x04type\x18\x01 \x01(\x0e\x32\'.org.lfedge.eve.certs.ZCertMetaDataType\x12\x11\n\tmeta_data\x18\x02 \x01(\x0c\"\x81\x02\n\x05ZCert\x12\x36\n\x08hashAlgo\x18\x01 \x01(\x0e\x32$.org.lfedge.eve.common.HashAlgorithm\x12\x10\n\x08\x63\x65rtHash\x18\x02 \x01(\x0c\x12-\n\x04type\x18\x03 \x01(\x0e\x32\x1f.org.lfedge.eve.certs.ZCertType\x12\x0c\n\x04\x63\x65rt\x18\x04 \x01(\x0c\x12\x33\n\nattributes\x18\x05 \x01(\x0b\x32\x1f.org.lfedge.eve.certs.ZCertAttr\x12<\n\x0fmeta_data_items\x18\x06 \x03(\x0b\x32#.org.lfedge.eve.certs.ZCertMetaData\"/\n\tZCertAttr\x12\x12\n\nis_mutable\x18\x01 \x01(\x08\x12\x0e\n\x06is_tpm\x18\x02 \x01(\x08*]\n\x11ZCertMetaDataType\x12!\n\x1dZ_CERT_META_DATA_TYPE_INVALID\x10\x00\x12%\n!Z_CERT_META_DATA_TYPE_TPM2_PUBLIC\x10\x01*\xaf\x02\n\tZCertType\x12\x1d\n\x19\x43\x45RT_TYPE_CONTROLLER_NONE\x10\x00\x12 \n\x1c\x43\x45RT_TYPE_CONTROLLER_SIGNING\x10\x01\x12%\n!CERT_TYPE_CONTROLLER_INTERMEDIATE\x10\x02\x12&\n\"CERT_TYPE_CONTROLLER_ECDH_EXCHANGE\x10\x03\x12\x1f\n\x1b\x43\x45RT_TYPE_DEVICE_ONBOARDING\x10\n\x12\'\n#CERT_TYPE_DEVICE_RESTRICTED_SIGNING\x10\x0b\x12$\n CERT_TYPE_DEVICE_ENDORSEMENT_RSA\x10\x0c\x12\"\n\x1e\x43\x45RT_TYPE_DEVICE_ECDH_EXCHANGE\x10\rB;\n\x14org.lfedge.eve.certsZ#github.com/lf-edge/eve/api/go/certsb\x06proto3')
  ,
  dependencies=[evecommon_dot_evecommon__pb2.DESCRIPTOR,])

_ZCERTMETADATATYPE = _descriptor.EnumDescriptor(
  name='ZCertMetaDataType',
  full_name='org.lfedge.eve.certs.ZCertMetaDataType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Z_CERT_META_DATA_TYPE_INVALID', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Z_CERT_META_DATA_TYPE_TPM2_PUBLIC', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=533,
  serialized_end=626,
)
_sym_db.RegisterEnumDescriptor(_ZCERTMETADATATYPE)

ZCertMetaDataType = enum_type_wrapper.EnumTypeWrapper(_ZCERTMETADATATYPE)
_ZCERTTYPE = _descriptor.EnumDescriptor(
  name='ZCertType',
  full_name='org.lfedge.eve.certs.ZCertType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_CONTROLLER_NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_CONTROLLER_SIGNING', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_CONTROLLER_INTERMEDIATE', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_CONTROLLER_ECDH_EXCHANGE', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_DEVICE_ONBOARDING', index=4, number=10,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_DEVICE_RESTRICTED_SIGNING', index=5, number=11,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_DEVICE_ENDORSEMENT_RSA', index=6, number=12,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_DEVICE_ECDH_EXCHANGE', index=7, number=13,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=629,
  serialized_end=932,
)
_sym_db.RegisterEnumDescriptor(_ZCERTTYPE)

ZCertType = enum_type_wrapper.EnumTypeWrapper(_ZCERTTYPE)
Z_CERT_META_DATA_TYPE_INVALID = 0
Z_CERT_META_DATA_TYPE_TPM2_PUBLIC = 1
CERT_TYPE_CONTROLLER_NONE = 0
CERT_TYPE_CONTROLLER_SIGNING = 1
CERT_TYPE_CONTROLLER_INTERMEDIATE = 2
CERT_TYPE_CONTROLLER_ECDH_EXCHANGE = 3
CERT_TYPE_DEVICE_ONBOARDING = 10
CERT_TYPE_DEVICE_RESTRICTED_SIGNING = 11
CERT_TYPE_DEVICE_ENDORSEMENT_RSA = 12
CERT_TYPE_DEVICE_ECDH_EXCHANGE = 13



_ZCONTROLLERCERT = _descriptor.Descriptor(
  name='ZControllerCert',
  full_name='org.lfedge.eve.certs.ZControllerCert',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='certs', full_name='org.lfedge.eve.certs.ZControllerCert.certs', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=70,
  serialized_end=131,
)


_ZCERTMETADATA = _descriptor.Descriptor(
  name='ZCertMetaData',
  full_name='org.lfedge.eve.certs.ZCertMetaData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='org.lfedge.eve.certs.ZCertMetaData.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='meta_data', full_name='org.lfedge.eve.certs.ZCertMetaData.meta_data', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
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
  serialized_start=133,
  serialized_end=222,
)


_ZCERT = _descriptor.Descriptor(
  name='ZCert',
  full_name='org.lfedge.eve.certs.ZCert',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='hashAlgo', full_name='org.lfedge.eve.certs.ZCert.hashAlgo', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='certHash', full_name='org.lfedge.eve.certs.ZCert.certHash', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='org.lfedge.eve.certs.ZCert.type', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cert', full_name='org.lfedge.eve.certs.ZCert.cert', index=3,
      number=4, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='attributes', full_name='org.lfedge.eve.certs.ZCert.attributes', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='meta_data_items', full_name='org.lfedge.eve.certs.ZCert.meta_data_items', index=5,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=225,
  serialized_end=482,
)


_ZCERTATTR = _descriptor.Descriptor(
  name='ZCertAttr',
  full_name='org.lfedge.eve.certs.ZCertAttr',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='is_mutable', full_name='org.lfedge.eve.certs.ZCertAttr.is_mutable', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='is_tpm', full_name='org.lfedge.eve.certs.ZCertAttr.is_tpm', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=484,
  serialized_end=531,
)

_ZCONTROLLERCERT.fields_by_name['certs'].message_type = _ZCERT
_ZCERTMETADATA.fields_by_name['type'].enum_type = _ZCERTMETADATATYPE
_ZCERT.fields_by_name['hashAlgo'].enum_type = evecommon_dot_evecommon__pb2._HASHALGORITHM
_ZCERT.fields_by_name['type'].enum_type = _ZCERTTYPE
_ZCERT.fields_by_name['attributes'].message_type = _ZCERTATTR
_ZCERT.fields_by_name['meta_data_items'].message_type = _ZCERTMETADATA
DESCRIPTOR.message_types_by_name['ZControllerCert'] = _ZCONTROLLERCERT
DESCRIPTOR.message_types_by_name['ZCertMetaData'] = _ZCERTMETADATA
DESCRIPTOR.message_types_by_name['ZCert'] = _ZCERT
DESCRIPTOR.message_types_by_name['ZCertAttr'] = _ZCERTATTR
DESCRIPTOR.enum_types_by_name['ZCertMetaDataType'] = _ZCERTMETADATATYPE
DESCRIPTOR.enum_types_by_name['ZCertType'] = _ZCERTTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ZControllerCert = _reflection.GeneratedProtocolMessageType('ZControllerCert', (_message.Message,), dict(
  DESCRIPTOR = _ZCONTROLLERCERT,
  __module__ = 'certs.certs_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.certs.ZControllerCert)
  ))
_sym_db.RegisterMessage(ZControllerCert)

ZCertMetaData = _reflection.GeneratedProtocolMessageType('ZCertMetaData', (_message.Message,), dict(
  DESCRIPTOR = _ZCERTMETADATA,
  __module__ = 'certs.certs_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.certs.ZCertMetaData)
  ))
_sym_db.RegisterMessage(ZCertMetaData)

ZCert = _reflection.GeneratedProtocolMessageType('ZCert', (_message.Message,), dict(
  DESCRIPTOR = _ZCERT,
  __module__ = 'certs.certs_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.certs.ZCert)
  ))
_sym_db.RegisterMessage(ZCert)

ZCertAttr = _reflection.GeneratedProtocolMessageType('ZCertAttr', (_message.Message,), dict(
  DESCRIPTOR = _ZCERTATTR,
  __module__ = 'certs.certs_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.certs.ZCertAttr)
  ))
_sym_db.RegisterMessage(ZCertAttr)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
