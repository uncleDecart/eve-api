# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: config/devconfig.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from config import acipherinfo_pb2 as config_dot_acipherinfo__pb2
from config import appconfig_pb2 as config_dot_appconfig__pb2
from config import baseosconfig_pb2 as config_dot_baseosconfig__pb2
from config import devcommon_pb2 as config_dot_devcommon__pb2
from config import devmodel_pb2 as config_dot_devmodel__pb2
from config import netconfig_pb2 as config_dot_netconfig__pb2
from config import netinst_pb2 as config_dot_netinst__pb2
from config import storage_pb2 as config_dot_storage__pb2
from config import edgeview_pb2 as config_dot_edgeview__pb2
from certs import certs_pb2 as certs_dot_certs__pb2
from auth import auth_pb2 as auth_dot_auth__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x16\x63onfig/devconfig.proto\x12\x15org.lfedge.eve.config\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x18\x63onfig/acipherinfo.proto\x1a\x16\x63onfig/appconfig.proto\x1a\x19\x63onfig/baseosconfig.proto\x1a\x16\x63onfig/devcommon.proto\x1a\x15\x63onfig/devmodel.proto\x1a\x16\x63onfig/netconfig.proto\x1a\x14\x63onfig/netinst.proto\x1a\x14\x63onfig/storage.proto\x1a\x15\x63onfig/edgeview.proto\x1a\x11\x63\x65rts/certs.proto\x1a\x0f\x61uth/auth.proto\"\x1c\n\tLOCConfig\x12\x0f\n\x07loc_url\x18\x01 \x01(\t\"\x95\x0c\n\rEdgeDevConfig\x12\x31\n\x02id\x18\x01 \x01(\x0b\x32%.org.lfedge.eve.config.UUIDandVersion\x12\x36\n\x04\x61pps\x18\x04 \x03(\x0b\x32(.org.lfedge.eve.config.AppInstanceConfig\x12\x36\n\x08networks\x18\x05 \x03(\x0b\x32$.org.lfedge.eve.config.NetworkConfig\x12:\n\ndatastores\x18\x06 \x03(\x0b\x32&.org.lfedge.eve.config.DatastoreConfig\x12\x31\n\x04\x62\x61se\x18\x08 \x03(\x0b\x32#.org.lfedge.eve.config.BaseOSConfig\x12\x33\n\x06reboot\x18\t \x01(\x0b\x32#.org.lfedge.eve.config.DeviceOpsCmd\x12\x33\n\x06\x62\x61\x63kup\x18\n \x01(\x0b\x32#.org.lfedge.eve.config.DeviceOpsCmd\x12\x36\n\x0b\x63onfigItems\x18\x0b \x03(\x0b\x32!.org.lfedge.eve.config.ConfigItem\x12?\n\x11systemAdapterList\x18\x0c \x03(\x0b\x32$.org.lfedge.eve.config.SystemAdapter\x12\x37\n\x0c\x64\x65viceIoList\x18\r \x03(\x0b\x32!.org.lfedge.eve.config.PhysicalIO\x12\x14\n\x0cmanufacturer\x18\x0e \x01(\t\x12\x13\n\x0bproductName\x18\x0f \x01(\t\x12\x46\n\x10networkInstances\x18\x10 \x03(\x0b\x32,.org.lfedge.eve.config.NetworkInstanceConfig\x12<\n\x0e\x63ipherContexts\x18\x13 \x03(\x0b\x32$.org.lfedge.eve.config.CipherContext\x12\x37\n\x0b\x63ontentInfo\x18\x14 \x03(\x0b\x32\".org.lfedge.eve.config.ContentTree\x12.\n\x07volumes\x18\x15 \x03(\x0b\x32\x1d.org.lfedge.eve.config.Volume\x12!\n\x19\x63ontrollercert_confighash\x18\x16 \x01(\t\x12\x18\n\x10maintenance_mode\x18\x18 \x01(\x08\x12\x18\n\x10\x63ontroller_epoch\x18\x19 \x01(\x03\x12-\n\x06\x62\x61seos\x18\x1a \x01(\x0b\x32\x1d.org.lfedge.eve.config.BaseOS\x12\x16\n\x0eglobal_profile\x18\x1b \x01(\t\x12\x1c\n\x14local_profile_server\x18\x1c \x01(\t\x12\x1c\n\x14profile_server_token\x18\x1d \x01(\t\x12\x31\n\x05vlans\x18\x1e \x03(\x0b\x32\".org.lfedge.eve.config.VlanAdapter\x12\x31\n\x05\x62onds\x18\x1f \x03(\x0b\x32\".org.lfedge.eve.config.BondAdapter\x12\x37\n\x08\x65\x64geview\x18  \x01(\x0b\x32%.org.lfedge.eve.config.EdgeViewConfig\x12\x31\n\x05\x64isks\x18! \x01(\x0b\x32\".org.lfedge.eve.config.DisksConfig\x12\x35\n\x08shutdown\x18\" \x01(\x0b\x32#.org.lfedge.eve.config.DeviceOpsCmd\x12\x13\n\x0b\x64\x65vice_name\x18# \x01(\t\x12\x14\n\x0cproject_name\x18$ \x01(\t\x12\x12\n\nproject_id\x18% \x01(\t\x12\x17\n\x0f\x65nterprise_name\x18& \x01(\t\x12\x15\n\renterprise_id\x18\' \x01(\t\x12\x34\n\x10\x63onfig_timestamp\x18( \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x34\n\nloc_config\x18) \x01(\x0b\x32 .org.lfedge.eve.config.LOCConfig\"<\n\rConfigRequest\x12\x12\n\nconfigHash\x18\x01 \x01(\t\x12\x17\n\x0fintegrity_token\x18\x02 \x01(\x0c\"Z\n\x0e\x43onfigResponse\x12\x34\n\x06\x63onfig\x18\x01 \x01(\x0b\x32$.org.lfedge.eve.config.EdgeDevConfig\x12\x12\n\nconfigHash\x18\x02 \x01(\t\"\x83\x01\n\x0f\x42ootstrapConfig\x12\x39\n\rsigned_config\x18\x01 \x01(\x0b\x32\".org.lfedge.eve.auth.AuthContainer\x12\x35\n\x10\x63ontroller_certs\x18\x02 \x03(\x0b\x32\x1b.org.lfedge.eve.certs.ZCertB=\n\x15org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/configb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'config.devconfig_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\025org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/config'
  _LOCCONFIG._serialized_start=333
  _LOCCONFIG._serialized_end=361
  _EDGEDEVCONFIG._serialized_start=364
  _EDGEDEVCONFIG._serialized_end=1921
  _CONFIGREQUEST._serialized_start=1923
  _CONFIGREQUEST._serialized_end=1983
  _CONFIGRESPONSE._serialized_start=1985
  _CONFIGRESPONSE._serialized_end=2075
  _BOOTSTRAPCONFIG._serialized_start=2078
  _BOOTSTRAPCONFIG._serialized_end=2209
# @@protoc_insertion_point(module_scope)
