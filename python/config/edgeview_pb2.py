# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: config/edgeview.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15\x63onfig/edgeview.proto\x12\x15org.lfedge.eve.config\"\x92\x02\n\x0e\x45\x64geViewConfig\x12\r\n\x05token\x18\x01 \x01(\t\x12\x15\n\rdisp_cert_pem\x18\x02 \x03(\x0c\x12?\n\ndev_policy\x18\x03 \x01(\x0b\x32+.org.lfedge.eve.config.DevDebugAccessPolicy\x12?\n\napp_policy\x18\x04 \x01(\x0b\x32+.org.lfedge.eve.config.AppDebugAccessPolicy\x12\x41\n\next_policy\x18\x05 \x01(\x0b\x32-.org.lfedge.eve.config.ExternalEndPointPolicy\x12\x15\n\rgeneration_id\x18\x06 \x01(\r\")\n\x14\x44\x65vDebugAccessPolicy\x12\x11\n\tallow_dev\x18\x01 \x01(\x08\")\n\x14\x41ppDebugAccessPolicy\x12\x11\n\tallow_app\x18\x01 \x01(\x08\"+\n\x16\x45xternalEndPointPolicy\x12\x11\n\tallow_ext\x18\x01 \x01(\x08\x42=\n\x15org.lfedge.eve.configZ$github.com/lf-edge/eve-api/go/configb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'config.edgeview_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\025org.lfedge.eve.configZ$github.com/lf-edge/eve-api/go/config'
  _globals['_EDGEVIEWCONFIG']._serialized_start=49
  _globals['_EDGEVIEWCONFIG']._serialized_end=323
  _globals['_DEVDEBUGACCESSPOLICY']._serialized_start=325
  _globals['_DEVDEBUGACCESSPOLICY']._serialized_end=366
  _globals['_APPDEBUGACCESSPOLICY']._serialized_start=368
  _globals['_APPDEBUGACCESSPOLICY']._serialized_end=409
  _globals['_EXTERNALENDPOINTPOLICY']._serialized_start=411
  _globals['_EXTERNALENDPOINTPOLICY']._serialized_end=454
# @@protoc_insertion_point(module_scope)
