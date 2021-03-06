// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/lldp.api.json

/*
 Package lldp is a generated from VPP binary API module 'lldp'.

 It contains following objects:
	  4 messages
	  2 services

*/
package lldp

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
//	    "lldp_config": {
//	        "reply": "lldp_config_reply"
//	    },
//	    "sw_interface_set_lldp": {
//	        "reply": "sw_interface_set_lldp_reply"
//	    }
//	},
//
type Services interface {
	LldpConfig(*LldpConfig) (*LldpConfigReply, error)
	SwInterfaceSetLldp(*SwInterfaceSetLldp) (*SwInterfaceSetLldpReply, error)
}

/* Messages */

// LldpConfig represents VPP binary API message 'lldp_config':
//
//	"lldp_config",
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
//	    "system_name",
//	    256
//	],
//	[
//	    "u32",
//	    "tx_hold"
//	],
//	[
//	    "u32",
//	    "tx_interval"
//	],
//	{
//	    "crc": "0x2410286f"
//	}
//
type LldpConfig struct {
	SystemName []byte `struc:"[256]byte"`
	TxHold     uint32
	TxInterval uint32
}

func (*LldpConfig) GetMessageName() string {
	return "lldp_config"
}
func (*LldpConfig) GetCrcString() string {
	return "2410286f"
}
func (*LldpConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// LldpConfigReply represents VPP binary API message 'lldp_config_reply':
//
//	"lldp_config_reply",
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
type LldpConfigReply struct {
	Retval int32
}

func (*LldpConfigReply) GetMessageName() string {
	return "lldp_config_reply"
}
func (*LldpConfigReply) GetCrcString() string {
	return "e8d4e804"
}
func (*LldpConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SwInterfaceSetLldp represents VPP binary API message 'sw_interface_set_lldp':
//
//	"sw_interface_set_lldp",
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
//	[
//	    "u8",
//	    "port_desc",
//	    256
//	],
//	[
//	    "u8",
//	    "mgmt_ip4",
//	    4
//	],
//	[
//	    "u8",
//	    "mgmt_ip6",
//	    16
//	],
//	[
//	    "u8",
//	    "mgmt_oid",
//	    128
//	],
//	[
//	    "u8",
//	    "enable"
//	],
//	{
//	    "crc": "0x2d85d156"
//	}
//
type SwInterfaceSetLldp struct {
	SwIfIndex uint32
	PortDesc  []byte `struc:"[256]byte"`
	MgmtIP4   []byte `struc:"[4]byte"`
	MgmtIP6   []byte `struc:"[16]byte"`
	MgmtOid   []byte `struc:"[128]byte"`
	Enable    uint8
}

func (*SwInterfaceSetLldp) GetMessageName() string {
	return "sw_interface_set_lldp"
}
func (*SwInterfaceSetLldp) GetCrcString() string {
	return "2d85d156"
}
func (*SwInterfaceSetLldp) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SwInterfaceSetLldpReply represents VPP binary API message 'sw_interface_set_lldp_reply':
//
//	"sw_interface_set_lldp_reply",
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
type SwInterfaceSetLldpReply struct {
	Retval int32
}

func (*SwInterfaceSetLldpReply) GetMessageName() string {
	return "sw_interface_set_lldp_reply"
}
func (*SwInterfaceSetLldpReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SwInterfaceSetLldpReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*LldpConfig)(nil), "lldp.LldpConfig")
	api.RegisterMessage((*LldpConfigReply)(nil), "lldp.LldpConfigReply")
	api.RegisterMessage((*SwInterfaceSetLldp)(nil), "lldp.SwInterfaceSetLldp")
	api.RegisterMessage((*SwInterfaceSetLldpReply)(nil), "lldp.SwInterfaceSetLldpReply")
}
