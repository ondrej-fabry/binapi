// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/ipfix_export.api.json

/*
 Package ipfix_export is a generated from VPP binary API module 'ipfix_export'.

 It contains following objects:
	 12 messages
	  6 services

*/
package ipfix_export

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// Services represents VPP binary API services:
//
//	"services": {
//	    "ipfix_exporter_dump": {
//	        "reply": "ipfix_exporter_details",
//	        "stream": true
//	    },
//	    "ipfix_classify_stream_dump": {
//	        "reply": "ipfix_classify_stream_details",
//	        "stream": true
//	    },
//	    "ipfix_classify_table_add_del": {
//	        "reply": "ipfix_classify_table_add_del_reply"
//	    },
//	    "set_ipfix_classify_stream": {
//	        "reply": "set_ipfix_classify_stream_reply"
//	    },
//	    "set_ipfix_exporter": {
//	        "reply": "set_ipfix_exporter_reply"
//	    },
//	    "ipfix_classify_table_dump": {
//	        "reply": "ipfix_classify_table_details",
//	        "stream": true
//	    }
//	},
//
type Services interface {
	DumpIpfixClassifyStream(*IpfixClassifyStreamDump) ([]*IpfixClassifyStreamDetails, error)
	DumpIpfixClassifyTable(*IpfixClassifyTableDump) ([]*IpfixClassifyTableDetails, error)
	DumpIpfixExporter(*IpfixExporterDump) ([]*IpfixExporterDetails, error)
	IpfixClassifyTableAddDel(*IpfixClassifyTableAddDel) (*IpfixClassifyTableAddDelReply, error)
	SetIpfixClassifyStream(*SetIpfixClassifyStream) (*SetIpfixClassifyStreamReply, error)
	SetIpfixExporter(*SetIpfixExporter) (*SetIpfixExporterReply, error)
}

/* Messages */

// SetIpfixExporter represents VPP binary API message 'set_ipfix_exporter':
//
//	"set_ipfix_exporter",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u8",
//	    "collector_address",
//	    16
//	],
//	[
//	    "u16",
//	    "collector_port"
//	],
//	[
//	    "u8",
//	    "src_address",
//	    16
//	],
//	[
//	    "u32",
//	    "vrf_id"
//	],
//	[
//	    "u32",
//	    "path_mtu"
//	],
//	[
//	    "u32",
//	    "template_interval"
//	],
//	[
//	    "u8",
//	    "udp_checksum"
//	],
//	{
//	    "crc": "0x4ff71dea"
//	}
//
type SetIpfixExporter struct {
	CollectorAddress []byte `struc:"[16]byte"`
	CollectorPort    uint16
	SrcAddress       []byte `struc:"[16]byte"`
	VrfID            uint32
	PathMtu          uint32
	TemplateInterval uint32
	UDPChecksum      uint8
}

func (*SetIpfixExporter) GetMessageName() string {
	return "set_ipfix_exporter"
}
func (*SetIpfixExporter) GetCrcString() string {
	return "4ff71dea"
}
func (*SetIpfixExporter) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SetIpfixExporterReply represents VPP binary API message 'set_ipfix_exporter_reply':
//
//	"set_ipfix_exporter_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xe8d4e804"
//	}
//
type SetIpfixExporterReply struct {
	Retval int32
}

func (*SetIpfixExporterReply) GetMessageName() string {
	return "set_ipfix_exporter_reply"
}
func (*SetIpfixExporterReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SetIpfixExporterReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpfixExporterDump represents VPP binary API message 'ipfix_exporter_dump':
//
//	"ipfix_exporter_dump",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	{
//	    "crc": "0x51077d14"
//	}
//
type IpfixExporterDump struct{}

func (*IpfixExporterDump) GetMessageName() string {
	return "ipfix_exporter_dump"
}
func (*IpfixExporterDump) GetCrcString() string {
	return "51077d14"
}
func (*IpfixExporterDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpfixExporterDetails represents VPP binary API message 'ipfix_exporter_details':
//
//	"ipfix_exporter_details",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u8",
//	    "collector_address",
//	    16
//	],
//	[
//	    "u16",
//	    "collector_port"
//	],
//	[
//	    "u8",
//	    "src_address",
//	    16
//	],
//	[
//	    "u32",
//	    "vrf_id"
//	],
//	[
//	    "u32",
//	    "path_mtu"
//	],
//	[
//	    "u32",
//	    "template_interval"
//	],
//	[
//	    "u8",
//	    "udp_checksum"
//	],
//	{
//	    "crc": "0x742dddee"
//	}
//
type IpfixExporterDetails struct {
	CollectorAddress []byte `struc:"[16]byte"`
	CollectorPort    uint16
	SrcAddress       []byte `struc:"[16]byte"`
	VrfID            uint32
	PathMtu          uint32
	TemplateInterval uint32
	UDPChecksum      uint8
}

func (*IpfixExporterDetails) GetMessageName() string {
	return "ipfix_exporter_details"
}
func (*IpfixExporterDetails) GetCrcString() string {
	return "742dddee"
}
func (*IpfixExporterDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SetIpfixClassifyStream represents VPP binary API message 'set_ipfix_classify_stream':
//
//	"set_ipfix_classify_stream",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u32",
//	    "domain_id"
//	],
//	[
//	    "u16",
//	    "src_port"
//	],
//	{
//	    "crc": "0xc9cbe053"
//	}
//
type SetIpfixClassifyStream struct {
	DomainID uint32
	SrcPort  uint16
}

func (*SetIpfixClassifyStream) GetMessageName() string {
	return "set_ipfix_classify_stream"
}
func (*SetIpfixClassifyStream) GetCrcString() string {
	return "c9cbe053"
}
func (*SetIpfixClassifyStream) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SetIpfixClassifyStreamReply represents VPP binary API message 'set_ipfix_classify_stream_reply':
//
//	"set_ipfix_classify_stream_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xe8d4e804"
//	}
//
type SetIpfixClassifyStreamReply struct {
	Retval int32
}

func (*SetIpfixClassifyStreamReply) GetMessageName() string {
	return "set_ipfix_classify_stream_reply"
}
func (*SetIpfixClassifyStreamReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SetIpfixClassifyStreamReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpfixClassifyStreamDump represents VPP binary API message 'ipfix_classify_stream_dump':
//
//	"ipfix_classify_stream_dump",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	{
//	    "crc": "0x51077d14"
//	}
//
type IpfixClassifyStreamDump struct{}

func (*IpfixClassifyStreamDump) GetMessageName() string {
	return "ipfix_classify_stream_dump"
}
func (*IpfixClassifyStreamDump) GetCrcString() string {
	return "51077d14"
}
func (*IpfixClassifyStreamDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpfixClassifyStreamDetails represents VPP binary API message 'ipfix_classify_stream_details':
//
//	"ipfix_classify_stream_details",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u32",
//	    "domain_id"
//	],
//	[
//	    "u16",
//	    "src_port"
//	],
//	{
//	    "crc": "0x2903539d"
//	}
//
type IpfixClassifyStreamDetails struct {
	DomainID uint32
	SrcPort  uint16
}

func (*IpfixClassifyStreamDetails) GetMessageName() string {
	return "ipfix_classify_stream_details"
}
func (*IpfixClassifyStreamDetails) GetCrcString() string {
	return "2903539d"
}
func (*IpfixClassifyStreamDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpfixClassifyTableAddDel represents VPP binary API message 'ipfix_classify_table_add_del':
//
//	"ipfix_classify_table_add_del",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u32",
//	    "table_id"
//	],
//	[
//	    "u8",
//	    "ip_version"
//	],
//	[
//	    "u8",
//	    "transport_protocol"
//	],
//	[
//	    "u8",
//	    "is_add"
//	],
//	{
//	    "crc": "0x48efe167"
//	}
//
type IpfixClassifyTableAddDel struct {
	TableID           uint32
	IPVersion         uint8
	TransportProtocol uint8
	IsAdd             uint8
}

func (*IpfixClassifyTableAddDel) GetMessageName() string {
	return "ipfix_classify_table_add_del"
}
func (*IpfixClassifyTableAddDel) GetCrcString() string {
	return "48efe167"
}
func (*IpfixClassifyTableAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpfixClassifyTableAddDelReply represents VPP binary API message 'ipfix_classify_table_add_del_reply':
//
//	"ipfix_classify_table_add_del_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xe8d4e804"
//	}
//
type IpfixClassifyTableAddDelReply struct {
	Retval int32
}

func (*IpfixClassifyTableAddDelReply) GetMessageName() string {
	return "ipfix_classify_table_add_del_reply"
}
func (*IpfixClassifyTableAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IpfixClassifyTableAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IpfixClassifyTableDump represents VPP binary API message 'ipfix_classify_table_dump':
//
//	"ipfix_classify_table_dump",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	{
//	    "crc": "0x51077d14"
//	}
//
type IpfixClassifyTableDump struct{}

func (*IpfixClassifyTableDump) GetMessageName() string {
	return "ipfix_classify_table_dump"
}
func (*IpfixClassifyTableDump) GetCrcString() string {
	return "51077d14"
}
func (*IpfixClassifyTableDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IpfixClassifyTableDetails represents VPP binary API message 'ipfix_classify_table_details':
//
//	"ipfix_classify_table_details",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u32",
//	    "table_id"
//	],
//	[
//	    "u8",
//	    "ip_version"
//	],
//	[
//	    "u8",
//	    "transport_protocol"
//	],
//	{
//	    "crc": "0x973d0d5b"
//	}
//
type IpfixClassifyTableDetails struct {
	TableID           uint32
	IPVersion         uint8
	TransportProtocol uint8
}

func (*IpfixClassifyTableDetails) GetMessageName() string {
	return "ipfix_classify_table_details"
}
func (*IpfixClassifyTableDetails) GetCrcString() string {
	return "973d0d5b"
}
func (*IpfixClassifyTableDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*SetIpfixExporter)(nil), "ipfix_export.SetIpfixExporter")
	api.RegisterMessage((*SetIpfixExporterReply)(nil), "ipfix_export.SetIpfixExporterReply")
	api.RegisterMessage((*IpfixExporterDump)(nil), "ipfix_export.IpfixExporterDump")
	api.RegisterMessage((*IpfixExporterDetails)(nil), "ipfix_export.IpfixExporterDetails")
	api.RegisterMessage((*SetIpfixClassifyStream)(nil), "ipfix_export.SetIpfixClassifyStream")
	api.RegisterMessage((*SetIpfixClassifyStreamReply)(nil), "ipfix_export.SetIpfixClassifyStreamReply")
	api.RegisterMessage((*IpfixClassifyStreamDump)(nil), "ipfix_export.IpfixClassifyStreamDump")
	api.RegisterMessage((*IpfixClassifyStreamDetails)(nil), "ipfix_export.IpfixClassifyStreamDetails")
	api.RegisterMessage((*IpfixClassifyTableAddDel)(nil), "ipfix_export.IpfixClassifyTableAddDel")
	api.RegisterMessage((*IpfixClassifyTableAddDelReply)(nil), "ipfix_export.IpfixClassifyTableAddDelReply")
	api.RegisterMessage((*IpfixClassifyTableDump)(nil), "ipfix_export.IpfixClassifyTableDump")
	api.RegisterMessage((*IpfixClassifyTableDetails)(nil), "ipfix_export.IpfixClassifyTableDetails")
}
