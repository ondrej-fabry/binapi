// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/sctp.api.json

/*
 Package sctp is a generated from VPP binary API module 'sctp'.

 It contains following objects:
	  6 messages
	  3 services

*/
package sctp

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
//	    "sctp_add_src_dst_connection": {
//	        "reply": "sctp_add_src_dst_connection_reply"
//	    },
//	    "sctp_config": {
//	        "reply": "sctp_config_reply"
//	    },
//	    "sctp_del_src_dst_connection": {
//	        "reply": "sctp_del_src_dst_connection_reply"
//	    }
//	},
//
type Services interface {
	SctpAddSrcDstConnection(*SctpAddSrcDstConnection) (*SctpAddSrcDstConnectionReply, error)
	SctpConfig(*SctpConfig) (*SctpConfigReply, error)
	SctpDelSrcDstConnection(*SctpDelSrcDstConnection) (*SctpDelSrcDstConnectionReply, error)
}

/* Messages */

// SctpAddSrcDstConnection represents VPP binary API message 'sctp_add_src_dst_connection':
//
//	"sctp_add_src_dst_connection",
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
//	    "is_ipv6"
//	],
//	[
//	    "u32",
//	    "vrf_id"
//	],
//	[
//	    "u8",
//	    "src_address",
//	    16
//	],
//	[
//	    "u8",
//	    "dst_address",
//	    16
//	],
//	{
//	    "crc": "0x45c59b2f"
//	}
//
type SctpAddSrcDstConnection struct {
	IsIPv6     uint8
	VrfID      uint32
	SrcAddress []byte `struc:"[16]byte"`
	DstAddress []byte `struc:"[16]byte"`
}

func (*SctpAddSrcDstConnection) GetMessageName() string {
	return "sctp_add_src_dst_connection"
}
func (*SctpAddSrcDstConnection) GetCrcString() string {
	return "45c59b2f"
}
func (*SctpAddSrcDstConnection) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SctpAddSrcDstConnectionReply represents VPP binary API message 'sctp_add_src_dst_connection_reply':
//
//	"sctp_add_src_dst_connection_reply",
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
type SctpAddSrcDstConnectionReply struct {
	Retval int32
}

func (*SctpAddSrcDstConnectionReply) GetMessageName() string {
	return "sctp_add_src_dst_connection_reply"
}
func (*SctpAddSrcDstConnectionReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SctpAddSrcDstConnectionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SctpDelSrcDstConnection represents VPP binary API message 'sctp_del_src_dst_connection':
//
//	"sctp_del_src_dst_connection",
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
//	    "is_ipv6"
//	],
//	[
//	    "u32",
//	    "vrf_id"
//	],
//	[
//	    "u8",
//	    "src_address",
//	    16
//	],
//	[
//	    "u8",
//	    "dst_address",
//	    16
//	],
//	{
//	    "crc": "0x45c59b2f"
//	}
//
type SctpDelSrcDstConnection struct {
	IsIPv6     uint8
	VrfID      uint32
	SrcAddress []byte `struc:"[16]byte"`
	DstAddress []byte `struc:"[16]byte"`
}

func (*SctpDelSrcDstConnection) GetMessageName() string {
	return "sctp_del_src_dst_connection"
}
func (*SctpDelSrcDstConnection) GetCrcString() string {
	return "45c59b2f"
}
func (*SctpDelSrcDstConnection) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SctpDelSrcDstConnectionReply represents VPP binary API message 'sctp_del_src_dst_connection_reply':
//
//	"sctp_del_src_dst_connection_reply",
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
type SctpDelSrcDstConnectionReply struct {
	Retval int32
}

func (*SctpDelSrcDstConnectionReply) GetMessageName() string {
	return "sctp_del_src_dst_connection_reply"
}
func (*SctpDelSrcDstConnectionReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SctpDelSrcDstConnectionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SctpConfig represents VPP binary API message 'sctp_config':
//
//	"sctp_config",
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
//	    "never_delay_sack"
//	],
//	[
//	    "u8",
//	    "never_bundle"
//	],
//	{
//	    "crc": "0x7617eced"
//	}
//
type SctpConfig struct {
	NeverDelaySack uint8
	NeverBundle    uint8
}

func (*SctpConfig) GetMessageName() string {
	return "sctp_config"
}
func (*SctpConfig) GetCrcString() string {
	return "7617eced"
}
func (*SctpConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SctpConfigReply represents VPP binary API message 'sctp_config_reply':
//
//	"sctp_config_reply",
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
type SctpConfigReply struct {
	Retval int32
}

func (*SctpConfigReply) GetMessageName() string {
	return "sctp_config_reply"
}
func (*SctpConfigReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SctpConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*SctpAddSrcDstConnection)(nil), "sctp.SctpAddSrcDstConnection")
	api.RegisterMessage((*SctpAddSrcDstConnectionReply)(nil), "sctp.SctpAddSrcDstConnectionReply")
	api.RegisterMessage((*SctpDelSrcDstConnection)(nil), "sctp.SctpDelSrcDstConnection")
	api.RegisterMessage((*SctpDelSrcDstConnectionReply)(nil), "sctp.SctpDelSrcDstConnectionReply")
	api.RegisterMessage((*SctpConfig)(nil), "sctp.SctpConfig")
	api.RegisterMessage((*SctpConfigReply)(nil), "sctp.SctpConfigReply")
}
