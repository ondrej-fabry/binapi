// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/dhcp6_pd_client_cp.api.json

/*
 Package dhcp6_pd_client_cp is a generated from VPP binary API module 'dhcp6_pd_client_cp'.

 It contains following objects:
	  4 messages
	  2 services

*/
package dhcp6_pd_client_cp

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
//	    "dhcp6_pd_client_enable_disable": {
//	        "reply": "dhcp6_pd_client_enable_disable_reply"
//	    },
//	    "ip6_add_del_address_using_prefix": {
//	        "reply": "ip6_add_del_address_using_prefix_reply"
//	    }
//	},
//
type Services interface {
	DHCP6PdClientEnableDisable(*DHCP6PdClientEnableDisable) (*DHCP6PdClientEnableDisableReply, error)
	IP6AddDelAddressUsingPrefix(*IP6AddDelAddressUsingPrefix) (*IP6AddDelAddressUsingPrefixReply, error)
}

/* Messages */

// DHCP6PdClientEnableDisable represents VPP binary API message 'dhcp6_pd_client_enable_disable':
//
//	"dhcp6_pd_client_enable_disable",
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
//	    "prefix_group",
//	    64
//	],
//	[
//	    "u8",
//	    "enable"
//	],
//	{
//	    "crc": "0xd2caf7ee"
//	}
//
type DHCP6PdClientEnableDisable struct {
	SwIfIndex   uint32
	PrefixGroup []byte `struc:"[64]byte"`
	Enable      uint8
}

func (*DHCP6PdClientEnableDisable) GetMessageName() string {
	return "dhcp6_pd_client_enable_disable"
}
func (*DHCP6PdClientEnableDisable) GetCrcString() string {
	return "d2caf7ee"
}
func (*DHCP6PdClientEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCP6PdClientEnableDisableReply represents VPP binary API message 'dhcp6_pd_client_enable_disable_reply':
//
//	"dhcp6_pd_client_enable_disable_reply",
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
type DHCP6PdClientEnableDisableReply struct {
	Retval int32
}

func (*DHCP6PdClientEnableDisableReply) GetMessageName() string {
	return "dhcp6_pd_client_enable_disable_reply"
}
func (*DHCP6PdClientEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCP6PdClientEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IP6AddDelAddressUsingPrefix represents VPP binary API message 'ip6_add_del_address_using_prefix':
//
//	"ip6_add_del_address_using_prefix",
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
//	    "prefix_group",
//	    64
//	],
//	[
//	    "u8",
//	    "address",
//	    16
//	],
//	[
//	    "u8",
//	    "prefix_length"
//	],
//	[
//	    "u8",
//	    "is_add"
//	],
//	{
//	    "crc": "0x3216fd48"
//	}
//
type IP6AddDelAddressUsingPrefix struct {
	SwIfIndex    uint32
	PrefixGroup  []byte `struc:"[64]byte"`
	Address      []byte `struc:"[16]byte"`
	PrefixLength uint8
	IsAdd        uint8
}

func (*IP6AddDelAddressUsingPrefix) GetMessageName() string {
	return "ip6_add_del_address_using_prefix"
}
func (*IP6AddDelAddressUsingPrefix) GetCrcString() string {
	return "3216fd48"
}
func (*IP6AddDelAddressUsingPrefix) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IP6AddDelAddressUsingPrefixReply represents VPP binary API message 'ip6_add_del_address_using_prefix_reply':
//
//	"ip6_add_del_address_using_prefix_reply",
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
type IP6AddDelAddressUsingPrefixReply struct {
	Retval int32
}

func (*IP6AddDelAddressUsingPrefixReply) GetMessageName() string {
	return "ip6_add_del_address_using_prefix_reply"
}
func (*IP6AddDelAddressUsingPrefixReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IP6AddDelAddressUsingPrefixReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*DHCP6PdClientEnableDisable)(nil), "dhcp6_pd_client_cp.DHCP6PdClientEnableDisable")
	api.RegisterMessage((*DHCP6PdClientEnableDisableReply)(nil), "dhcp6_pd_client_cp.DHCP6PdClientEnableDisableReply")
	api.RegisterMessage((*IP6AddDelAddressUsingPrefix)(nil), "dhcp6_pd_client_cp.IP6AddDelAddressUsingPrefix")
	api.RegisterMessage((*IP6AddDelAddressUsingPrefixReply)(nil), "dhcp6_pd_client_cp.IP6AddDelAddressUsingPrefixReply")
}