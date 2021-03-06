// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/cop.api.json

/*
 Package cop is a generated from VPP binary API module 'cop'.

 It contains following objects:
	  4 messages
	  2 services

*/
package cop

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
//	    "cop_interface_enable_disable": {
//	        "reply": "cop_interface_enable_disable_reply"
//	    },
//	    "cop_whitelist_enable_disable": {
//	        "reply": "cop_whitelist_enable_disable_reply"
//	    }
//	},
//
type Services interface {
	CopInterfaceEnableDisable(*CopInterfaceEnableDisable) (*CopInterfaceEnableDisableReply, error)
	CopWhitelistEnableDisable(*CopWhitelistEnableDisable) (*CopWhitelistEnableDisableReply, error)
}

/* Messages */

// CopInterfaceEnableDisable represents VPP binary API message 'cop_interface_enable_disable':
//
//	"cop_interface_enable_disable",
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
//	    "enable_disable"
//	],
//	{
//	    "crc": "0x69d24598"
//	}
//
type CopInterfaceEnableDisable struct {
	SwIfIndex     uint32
	EnableDisable uint8
}

func (*CopInterfaceEnableDisable) GetMessageName() string {
	return "cop_interface_enable_disable"
}
func (*CopInterfaceEnableDisable) GetCrcString() string {
	return "69d24598"
}
func (*CopInterfaceEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// CopInterfaceEnableDisableReply represents VPP binary API message 'cop_interface_enable_disable_reply':
//
//	"cop_interface_enable_disable_reply",
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
type CopInterfaceEnableDisableReply struct {
	Retval int32
}

func (*CopInterfaceEnableDisableReply) GetMessageName() string {
	return "cop_interface_enable_disable_reply"
}
func (*CopInterfaceEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*CopInterfaceEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// CopWhitelistEnableDisable represents VPP binary API message 'cop_whitelist_enable_disable':
//
//	"cop_whitelist_enable_disable",
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
//	    "u32",
//	    "fib_id"
//	],
//	[
//	    "u8",
//	    "ip4"
//	],
//	[
//	    "u8",
//	    "ip6"
//	],
//	[
//	    "u8",
//	    "default_cop"
//	],
//	{
//	    "crc": "0x8bb8f6dc"
//	}
//
type CopWhitelistEnableDisable struct {
	SwIfIndex  uint32
	FibID      uint32
	IP4        uint8
	IP6        uint8
	DefaultCop uint8
}

func (*CopWhitelistEnableDisable) GetMessageName() string {
	return "cop_whitelist_enable_disable"
}
func (*CopWhitelistEnableDisable) GetCrcString() string {
	return "8bb8f6dc"
}
func (*CopWhitelistEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// CopWhitelistEnableDisableReply represents VPP binary API message 'cop_whitelist_enable_disable_reply':
//
//	"cop_whitelist_enable_disable_reply",
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
type CopWhitelistEnableDisableReply struct {
	Retval int32
}

func (*CopWhitelistEnableDisableReply) GetMessageName() string {
	return "cop_whitelist_enable_disable_reply"
}
func (*CopWhitelistEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*CopWhitelistEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*CopInterfaceEnableDisable)(nil), "cop.CopInterfaceEnableDisable")
	api.RegisterMessage((*CopInterfaceEnableDisableReply)(nil), "cop.CopInterfaceEnableDisableReply")
	api.RegisterMessage((*CopWhitelistEnableDisable)(nil), "cop.CopWhitelistEnableDisable")
	api.RegisterMessage((*CopWhitelistEnableDisableReply)(nil), "cop.CopWhitelistEnableDisableReply")
}
