// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/af_packet.api.json

/*
 Package af_packet is a generated from VPP binary API module 'af_packet'.

 It contains following objects:
	  8 messages
	  4 services

*/
package af_packet

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
//	    "af_packet_dump": {
//	        "reply": "af_packet_details",
//	        "stream": true
//	    },
//	    "af_packet_set_l4_cksum_offload": {
//	        "reply": "af_packet_set_l4_cksum_offload_reply"
//	    },
//	    "af_packet_delete": {
//	        "reply": "af_packet_delete_reply"
//	    },
//	    "af_packet_create": {
//	        "reply": "af_packet_create_reply"
//	    }
//	},
//
type Services interface {
	DumpAfPacket(*AfPacketDump) ([]*AfPacketDetails, error)
	AfPacketCreate(*AfPacketCreate) (*AfPacketCreateReply, error)
	AfPacketDelete(*AfPacketDelete) (*AfPacketDeleteReply, error)
	AfPacketSetL4CksumOffload(*AfPacketSetL4CksumOffload) (*AfPacketSetL4CksumOffloadReply, error)
}

/* Messages */

// AfPacketCreate represents VPP binary API message 'af_packet_create':
//
//	"af_packet_create",
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
//	    "host_if_name",
//	    64
//	],
//	[
//	    "u8",
//	    "hw_addr",
//	    6
//	],
//	[
//	    "u8",
//	    "use_random_hw_addr"
//	],
//	{
//	    "crc": "0x6d5d30d6"
//	}
//
type AfPacketCreate struct {
	HostIfName      []byte `struc:"[64]byte"`
	HwAddr          []byte `struc:"[6]byte"`
	UseRandomHwAddr uint8
}

func (*AfPacketCreate) GetMessageName() string {
	return "af_packet_create"
}
func (*AfPacketCreate) GetCrcString() string {
	return "6d5d30d6"
}
func (*AfPacketCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AfPacketCreateReply represents VPP binary API message 'af_packet_create_reply':
//
//	"af_packet_create_reply",
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
type AfPacketCreateReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*AfPacketCreateReply) GetMessageName() string {
	return "af_packet_create_reply"
}
func (*AfPacketCreateReply) GetCrcString() string {
	return "fda5941f"
}
func (*AfPacketCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AfPacketDelete represents VPP binary API message 'af_packet_delete':
//
//	"af_packet_delete",
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
//	    "host_if_name",
//	    64
//	],
//	{
//	    "crc": "0x3efceda3"
//	}
//
type AfPacketDelete struct {
	HostIfName []byte `struc:"[64]byte"`
}

func (*AfPacketDelete) GetMessageName() string {
	return "af_packet_delete"
}
func (*AfPacketDelete) GetCrcString() string {
	return "3efceda3"
}
func (*AfPacketDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AfPacketDeleteReply represents VPP binary API message 'af_packet_delete_reply':
//
//	"af_packet_delete_reply",
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
type AfPacketDeleteReply struct {
	Retval int32
}

func (*AfPacketDeleteReply) GetMessageName() string {
	return "af_packet_delete_reply"
}
func (*AfPacketDeleteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*AfPacketDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AfPacketSetL4CksumOffload represents VPP binary API message 'af_packet_set_l4_cksum_offload':
//
//	"af_packet_set_l4_cksum_offload",
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
//	    "sw_if_index"
//	],
//	[
//	    "u8",
//	    "set"
//	],
//	{
//	    "crc": "0x86538585"
//	}
//
type AfPacketSetL4CksumOffload struct {
	SwIfIndex uint8
	Set       uint8
}

func (*AfPacketSetL4CksumOffload) GetMessageName() string {
	return "af_packet_set_l4_cksum_offload"
}
func (*AfPacketSetL4CksumOffload) GetCrcString() string {
	return "86538585"
}
func (*AfPacketSetL4CksumOffload) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AfPacketSetL4CksumOffloadReply represents VPP binary API message 'af_packet_set_l4_cksum_offload_reply':
//
//	"af_packet_set_l4_cksum_offload_reply",
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
type AfPacketSetL4CksumOffloadReply struct {
	Retval int32
}

func (*AfPacketSetL4CksumOffloadReply) GetMessageName() string {
	return "af_packet_set_l4_cksum_offload_reply"
}
func (*AfPacketSetL4CksumOffloadReply) GetCrcString() string {
	return "e8d4e804"
}
func (*AfPacketSetL4CksumOffloadReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AfPacketDump represents VPP binary API message 'af_packet_dump':
//
//	"af_packet_dump",
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
type AfPacketDump struct{}

func (*AfPacketDump) GetMessageName() string {
	return "af_packet_dump"
}
func (*AfPacketDump) GetCrcString() string {
	return "51077d14"
}
func (*AfPacketDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AfPacketDetails represents VPP binary API message 'af_packet_details':
//
//	"af_packet_details",
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
//	    "sw_if_index"
//	],
//	[
//	    "u8",
//	    "host_if_name",
//	    64
//	],
//	{
//	    "crc": "0x057205fa"
//	}
//
type AfPacketDetails struct {
	SwIfIndex  uint32
	HostIfName []byte `struc:"[64]byte"`
}

func (*AfPacketDetails) GetMessageName() string {
	return "af_packet_details"
}
func (*AfPacketDetails) GetCrcString() string {
	return "057205fa"
}
func (*AfPacketDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*AfPacketCreate)(nil), "af_packet.AfPacketCreate")
	api.RegisterMessage((*AfPacketCreateReply)(nil), "af_packet.AfPacketCreateReply")
	api.RegisterMessage((*AfPacketDelete)(nil), "af_packet.AfPacketDelete")
	api.RegisterMessage((*AfPacketDeleteReply)(nil), "af_packet.AfPacketDeleteReply")
	api.RegisterMessage((*AfPacketSetL4CksumOffload)(nil), "af_packet.AfPacketSetL4CksumOffload")
	api.RegisterMessage((*AfPacketSetL4CksumOffloadReply)(nil), "af_packet.AfPacketSetL4CksumOffloadReply")
	api.RegisterMessage((*AfPacketDump)(nil), "af_packet.AfPacketDump")
	api.RegisterMessage((*AfPacketDetails)(nil), "af_packet.AfPacketDetails")
}
