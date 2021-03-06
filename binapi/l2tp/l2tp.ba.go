// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/l2tp.api.json

/*
 Package l2tp is a generated from VPP binary API module 'l2tp'.

 It contains following objects:
	 10 messages
	  5 services

*/
package l2tp

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
//	    "l2tpv3_interface_enable_disable": {
//	        "reply": "l2tpv3_interface_enable_disable_reply"
//	    },
//	    "l2tpv3_set_tunnel_cookies": {
//	        "reply": "l2tpv3_set_tunnel_cookies_reply"
//	    },
//	    "l2tpv3_create_tunnel": {
//	        "reply": "l2tpv3_create_tunnel_reply"
//	    },
//	    "l2tpv3_set_lookup_key": {
//	        "reply": "l2tpv3_set_lookup_key_reply"
//	    },
//	    "sw_if_l2tpv3_tunnel_dump": {
//	        "reply": "sw_if_l2tpv3_tunnel_details",
//	        "stream": true
//	    }
//	},
//
type Services interface {
	DumpSwIfL2tpv3Tunnel(*SwIfL2tpv3TunnelDump) ([]*SwIfL2tpv3TunnelDetails, error)
	L2tpv3CreateTunnel(*L2tpv3CreateTunnel) (*L2tpv3CreateTunnelReply, error)
	L2tpv3InterfaceEnableDisable(*L2tpv3InterfaceEnableDisable) (*L2tpv3InterfaceEnableDisableReply, error)
	L2tpv3SetLookupKey(*L2tpv3SetLookupKey) (*L2tpv3SetLookupKeyReply, error)
	L2tpv3SetTunnelCookies(*L2tpv3SetTunnelCookies) (*L2tpv3SetTunnelCookiesReply, error)
}

/* Messages */

// L2tpv3CreateTunnel represents VPP binary API message 'l2tpv3_create_tunnel':
//
//	"l2tpv3_create_tunnel",
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
//	    "client_address",
//	    16
//	],
//	[
//	    "u8",
//	    "our_address",
//	    16
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u32",
//	    "local_session_id"
//	],
//	[
//	    "u32",
//	    "remote_session_id"
//	],
//	[
//	    "u64",
//	    "local_cookie"
//	],
//	[
//	    "u64",
//	    "remote_cookie"
//	],
//	[
//	    "u8",
//	    "l2_sublayer_present"
//	],
//	[
//	    "u32",
//	    "encap_vrf_id"
//	],
//	{
//	    "crc": "0x11bafa9f"
//	}
//
type L2tpv3CreateTunnel struct {
	ClientAddress     []byte `struc:"[16]byte"`
	OurAddress        []byte `struc:"[16]byte"`
	IsIPv6            uint8
	LocalSessionID    uint32
	RemoteSessionID   uint32
	LocalCookie       uint64
	RemoteCookie      uint64
	L2SublayerPresent uint8
	EncapVrfID        uint32
}

func (*L2tpv3CreateTunnel) GetMessageName() string {
	return "l2tpv3_create_tunnel"
}
func (*L2tpv3CreateTunnel) GetCrcString() string {
	return "11bafa9f"
}
func (*L2tpv3CreateTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2tpv3CreateTunnelReply represents VPP binary API message 'l2tpv3_create_tunnel_reply':
//
//	"l2tpv3_create_tunnel_reply",
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
type L2tpv3CreateTunnelReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*L2tpv3CreateTunnelReply) GetMessageName() string {
	return "l2tpv3_create_tunnel_reply"
}
func (*L2tpv3CreateTunnelReply) GetCrcString() string {
	return "fda5941f"
}
func (*L2tpv3CreateTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// L2tpv3SetTunnelCookies represents VPP binary API message 'l2tpv3_set_tunnel_cookies':
//
//	"l2tpv3_set_tunnel_cookies",
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
//	    "u64",
//	    "new_local_cookie"
//	],
//	[
//	    "u64",
//	    "new_remote_cookie"
//	],
//	{
//	    "crc": "0xb0ce1366"
//	}
//
type L2tpv3SetTunnelCookies struct {
	SwIfIndex       uint32
	NewLocalCookie  uint64
	NewRemoteCookie uint64
}

func (*L2tpv3SetTunnelCookies) GetMessageName() string {
	return "l2tpv3_set_tunnel_cookies"
}
func (*L2tpv3SetTunnelCookies) GetCrcString() string {
	return "b0ce1366"
}
func (*L2tpv3SetTunnelCookies) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2tpv3SetTunnelCookiesReply represents VPP binary API message 'l2tpv3_set_tunnel_cookies_reply':
//
//	"l2tpv3_set_tunnel_cookies_reply",
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
type L2tpv3SetTunnelCookiesReply struct {
	Retval int32
}

func (*L2tpv3SetTunnelCookiesReply) GetMessageName() string {
	return "l2tpv3_set_tunnel_cookies_reply"
}
func (*L2tpv3SetTunnelCookiesReply) GetCrcString() string {
	return "e8d4e804"
}
func (*L2tpv3SetTunnelCookiesReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SwIfL2tpv3TunnelDetails represents VPP binary API message 'sw_if_l2tpv3_tunnel_details':
//
//	"sw_if_l2tpv3_tunnel_details",
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
//	    "interface_name",
//	    64
//	],
//	[
//	    "u8",
//	    "client_address",
//	    16
//	],
//	[
//	    "u8",
//	    "our_address",
//	    16
//	],
//	[
//	    "u32",
//	    "local_session_id"
//	],
//	[
//	    "u32",
//	    "remote_session_id"
//	],
//	[
//	    "u64",
//	    "local_cookie",
//	    2
//	],
//	[
//	    "u64",
//	    "remote_cookie"
//	],
//	[
//	    "u8",
//	    "l2_sublayer_present"
//	],
//	{
//	    "crc": "0x7b22bb34"
//	}
//
type SwIfL2tpv3TunnelDetails struct {
	SwIfIndex         uint32
	InterfaceName     []byte `struc:"[64]byte"`
	ClientAddress     []byte `struc:"[16]byte"`
	OurAddress        []byte `struc:"[16]byte"`
	LocalSessionID    uint32
	RemoteSessionID   uint32
	LocalCookie       []uint64 `struc:"[2]uint64"`
	RemoteCookie      uint64
	L2SublayerPresent uint8
}

func (*SwIfL2tpv3TunnelDetails) GetMessageName() string {
	return "sw_if_l2tpv3_tunnel_details"
}
func (*SwIfL2tpv3TunnelDetails) GetCrcString() string {
	return "7b22bb34"
}
func (*SwIfL2tpv3TunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SwIfL2tpv3TunnelDump represents VPP binary API message 'sw_if_l2tpv3_tunnel_dump':
//
//	"sw_if_l2tpv3_tunnel_dump",
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
type SwIfL2tpv3TunnelDump struct{}

func (*SwIfL2tpv3TunnelDump) GetMessageName() string {
	return "sw_if_l2tpv3_tunnel_dump"
}
func (*SwIfL2tpv3TunnelDump) GetCrcString() string {
	return "51077d14"
}
func (*SwIfL2tpv3TunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2tpv3InterfaceEnableDisable represents VPP binary API message 'l2tpv3_interface_enable_disable':
//
//	"l2tpv3_interface_enable_disable",
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
//	    "enable_disable"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	{
//	    "crc": "0x57298519"
//	}
//
type L2tpv3InterfaceEnableDisable struct {
	EnableDisable uint8
	SwIfIndex     uint32
}

func (*L2tpv3InterfaceEnableDisable) GetMessageName() string {
	return "l2tpv3_interface_enable_disable"
}
func (*L2tpv3InterfaceEnableDisable) GetCrcString() string {
	return "57298519"
}
func (*L2tpv3InterfaceEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2tpv3InterfaceEnableDisableReply represents VPP binary API message 'l2tpv3_interface_enable_disable_reply':
//
//	"l2tpv3_interface_enable_disable_reply",
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
type L2tpv3InterfaceEnableDisableReply struct {
	Retval int32
}

func (*L2tpv3InterfaceEnableDisableReply) GetMessageName() string {
	return "l2tpv3_interface_enable_disable_reply"
}
func (*L2tpv3InterfaceEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*L2tpv3InterfaceEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// L2tpv3SetLookupKey represents VPP binary API message 'l2tpv3_set_lookup_key':
//
//	"l2tpv3_set_lookup_key",
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
//	    "key"
//	],
//	{
//	    "crc": "0xd7736d45"
//	}
//
type L2tpv3SetLookupKey struct {
	Key uint8
}

func (*L2tpv3SetLookupKey) GetMessageName() string {
	return "l2tpv3_set_lookup_key"
}
func (*L2tpv3SetLookupKey) GetCrcString() string {
	return "d7736d45"
}
func (*L2tpv3SetLookupKey) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2tpv3SetLookupKeyReply represents VPP binary API message 'l2tpv3_set_lookup_key_reply':
//
//	"l2tpv3_set_lookup_key_reply",
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
type L2tpv3SetLookupKeyReply struct {
	Retval int32
}

func (*L2tpv3SetLookupKeyReply) GetMessageName() string {
	return "l2tpv3_set_lookup_key_reply"
}
func (*L2tpv3SetLookupKeyReply) GetCrcString() string {
	return "e8d4e804"
}
func (*L2tpv3SetLookupKeyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*L2tpv3CreateTunnel)(nil), "l2tp.L2tpv3CreateTunnel")
	api.RegisterMessage((*L2tpv3CreateTunnelReply)(nil), "l2tp.L2tpv3CreateTunnelReply")
	api.RegisterMessage((*L2tpv3SetTunnelCookies)(nil), "l2tp.L2tpv3SetTunnelCookies")
	api.RegisterMessage((*L2tpv3SetTunnelCookiesReply)(nil), "l2tp.L2tpv3SetTunnelCookiesReply")
	api.RegisterMessage((*SwIfL2tpv3TunnelDetails)(nil), "l2tp.SwIfL2tpv3TunnelDetails")
	api.RegisterMessage((*SwIfL2tpv3TunnelDump)(nil), "l2tp.SwIfL2tpv3TunnelDump")
	api.RegisterMessage((*L2tpv3InterfaceEnableDisable)(nil), "l2tp.L2tpv3InterfaceEnableDisable")
	api.RegisterMessage((*L2tpv3InterfaceEnableDisableReply)(nil), "l2tp.L2tpv3InterfaceEnableDisableReply")
	api.RegisterMessage((*L2tpv3SetLookupKey)(nil), "l2tp.L2tpv3SetLookupKey")
	api.RegisterMessage((*L2tpv3SetLookupKeyReply)(nil), "l2tp.L2tpv3SetLookupKeyReply")
}
