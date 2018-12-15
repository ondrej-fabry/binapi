// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/avf.api.json

/*
 Package avf is a generated from VPP binary API module 'avf'.

 It contains following objects:
	  4 messages
	  2 services

*/
package avf

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
//	    "avf_create": {
//	        "reply": "avf_create_reply"
//	    },
//	    "avf_delete": {
//	        "reply": "avf_delete_reply"
//	    }
//	},
//
type Services interface {
	AvfCreate(*AvfCreate) (*AvfCreateReply, error)
	AvfDelete(*AvfDelete) (*AvfDeleteReply, error)
}

/* Messages */

// AvfCreate represents VPP binary API message 'avf_create':
//
//	"avf_create",
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
//	    "pci_addr"
//	],
//	[
//	    "i32",
//	    "enable_elog"
//	],
//	[
//	    "u16",
//	    "rxq_num"
//	],
//	[
//	    "u16",
//	    "rxq_size"
//	],
//	[
//	    "u16",
//	    "txq_size"
//	],
//	{
//	    "crc": "0xdaab8ae2"
//	}
//
type AvfCreate struct {
	PciAddr    uint32
	EnableElog int32
	RxqNum     uint16
	RxqSize    uint16
	TxqSize    uint16
}

func (*AvfCreate) GetMessageName() string {
	return "avf_create"
}
func (*AvfCreate) GetCrcString() string {
	return "daab8ae2"
}
func (*AvfCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AvfCreateReply represents VPP binary API message 'avf_create_reply':
//
//	"avf_create_reply",
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
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	{
//	    "crc": "0xfda5941f"
//	}
//
type AvfCreateReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*AvfCreateReply) GetMessageName() string {
	return "avf_create_reply"
}
func (*AvfCreateReply) GetCrcString() string {
	return "fda5941f"
}
func (*AvfCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AvfDelete represents VPP binary API message 'avf_delete':
//
//	"avf_delete",
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
//	    "sw_if_index"
//	],
//	{
//	    "crc": "0x529cb13f"
//	}
//
type AvfDelete struct {
	SwIfIndex uint32
}

func (*AvfDelete) GetMessageName() string {
	return "avf_delete"
}
func (*AvfDelete) GetCrcString() string {
	return "529cb13f"
}
func (*AvfDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AvfDeleteReply represents VPP binary API message 'avf_delete_reply':
//
//	"avf_delete_reply",
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
type AvfDeleteReply struct {
	Retval int32
}

func (*AvfDeleteReply) GetMessageName() string {
	return "avf_delete_reply"
}
func (*AvfDeleteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*AvfDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*AvfCreate)(nil), "avf.AvfCreate")
	api.RegisterMessage((*AvfCreateReply)(nil), "avf.AvfCreateReply")
	api.RegisterMessage((*AvfDelete)(nil), "avf.AvfDelete")
	api.RegisterMessage((*AvfDeleteReply)(nil), "avf.AvfDeleteReply")
}
