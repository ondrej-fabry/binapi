// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/geneve.api.json

/*
 Package geneve is a generated from VPP binary API module 'geneve'.

 It contains following objects:
	  6 messages
	  3 services

*/
package geneve

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
//	    "geneve_tunnel_dump": {
//	        "reply": "geneve_tunnel_details",
//	        "stream": true
//	    },
//	    "geneve_add_del_tunnel": {
//	        "reply": "geneve_add_del_tunnel_reply"
//	    },
//	    "sw_interface_set_geneve_bypass": {
//	        "reply": "sw_interface_set_geneve_bypass_reply"
//	    }
//	},
//
type Services interface {
	DumpGeneveTunnel(*GeneveTunnelDump) ([]*GeneveTunnelDetails, error)
	GeneveAddDelTunnel(*GeneveAddDelTunnel) (*GeneveAddDelTunnelReply, error)
	SwInterfaceSetGeneveBypass(*SwInterfaceSetGeneveBypass) (*SwInterfaceSetGeneveBypassReply, error)
}

/* Messages */

// GeneveAddDelTunnel represents VPP binary API message 'geneve_add_del_tunnel':
//
//	"geneve_add_del_tunnel",
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
//	    "is_add"
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "local_address",
//	    16
//	],
//	[
//	    "u8",
//	    "remote_address",
//	    16
//	],
//	[
//	    "u32",
//	    "mcast_sw_if_index"
//	],
//	[
//	    "u32",
//	    "encap_vrf_id"
//	],
//	[
//	    "u32",
//	    "decap_next_index"
//	],
//	[
//	    "u32",
//	    "vni"
//	],
//	{
//	    "crc": "0x403cf981"
//	}
//
type GeneveAddDelTunnel struct {
	IsAdd          uint8
	IsIPv6         uint8
	LocalAddress   []byte `struc:"[16]byte"`
	RemoteAddress  []byte `struc:"[16]byte"`
	McastSwIfIndex uint32
	EncapVrfID     uint32
	DecapNextIndex uint32
	Vni            uint32
}

func (*GeneveAddDelTunnel) GetMessageName() string {
	return "geneve_add_del_tunnel"
}
func (*GeneveAddDelTunnel) GetCrcString() string {
	return "403cf981"
}
func (*GeneveAddDelTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// GeneveAddDelTunnelReply represents VPP binary API message 'geneve_add_del_tunnel_reply':
//
//	"geneve_add_del_tunnel_reply",
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
type GeneveAddDelTunnelReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*GeneveAddDelTunnelReply) GetMessageName() string {
	return "geneve_add_del_tunnel_reply"
}
func (*GeneveAddDelTunnelReply) GetCrcString() string {
	return "fda5941f"
}
func (*GeneveAddDelTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// GeneveTunnelDump represents VPP binary API message 'geneve_tunnel_dump':
//
//	"geneve_tunnel_dump",
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
type GeneveTunnelDump struct {
	SwIfIndex uint32
}

func (*GeneveTunnelDump) GetMessageName() string {
	return "geneve_tunnel_dump"
}
func (*GeneveTunnelDump) GetCrcString() string {
	return "529cb13f"
}
func (*GeneveTunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// GeneveTunnelDetails represents VPP binary API message 'geneve_tunnel_details':
//
//	"geneve_tunnel_details",
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
//	    "src_address",
//	    16
//	],
//	[
//	    "u8",
//	    "dst_address",
//	    16
//	],
//	[
//	    "u32",
//	    "mcast_sw_if_index"
//	],
//	[
//	    "u32",
//	    "encap_vrf_id"
//	],
//	[
//	    "u32",
//	    "decap_next_index"
//	],
//	[
//	    "u32",
//	    "vni"
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	{
//	    "crc": "0x024fa31f"
//	}
//
type GeneveTunnelDetails struct {
	SwIfIndex      uint32
	SrcAddress     []byte `struc:"[16]byte"`
	DstAddress     []byte `struc:"[16]byte"`
	McastSwIfIndex uint32
	EncapVrfID     uint32
	DecapNextIndex uint32
	Vni            uint32
	IsIPv6         uint8
}

func (*GeneveTunnelDetails) GetMessageName() string {
	return "geneve_tunnel_details"
}
func (*GeneveTunnelDetails) GetCrcString() string {
	return "024fa31f"
}
func (*GeneveTunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SwInterfaceSetGeneveBypass represents VPP binary API message 'sw_interface_set_geneve_bypass':
//
//	"sw_interface_set_geneve_bypass",
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
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "enable"
//	],
//	{
//	    "crc": "0xe74ca095"
//	}
//
type SwInterfaceSetGeneveBypass struct {
	SwIfIndex uint32
	IsIPv6    uint8
	Enable    uint8
}

func (*SwInterfaceSetGeneveBypass) GetMessageName() string {
	return "sw_interface_set_geneve_bypass"
}
func (*SwInterfaceSetGeneveBypass) GetCrcString() string {
	return "e74ca095"
}
func (*SwInterfaceSetGeneveBypass) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SwInterfaceSetGeneveBypassReply represents VPP binary API message 'sw_interface_set_geneve_bypass_reply':
//
//	"sw_interface_set_geneve_bypass_reply",
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
type SwInterfaceSetGeneveBypassReply struct {
	Retval int32
}

func (*SwInterfaceSetGeneveBypassReply) GetMessageName() string {
	return "sw_interface_set_geneve_bypass_reply"
}
func (*SwInterfaceSetGeneveBypassReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SwInterfaceSetGeneveBypassReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*GeneveAddDelTunnel)(nil), "geneve.GeneveAddDelTunnel")
	api.RegisterMessage((*GeneveAddDelTunnelReply)(nil), "geneve.GeneveAddDelTunnelReply")
	api.RegisterMessage((*GeneveTunnelDump)(nil), "geneve.GeneveTunnelDump")
	api.RegisterMessage((*GeneveTunnelDetails)(nil), "geneve.GeneveTunnelDetails")
	api.RegisterMessage((*SwInterfaceSetGeneveBypass)(nil), "geneve.SwInterfaceSetGeneveBypass")
	api.RegisterMessage((*SwInterfaceSetGeneveBypassReply)(nil), "geneve.SwInterfaceSetGeneveBypassReply")
}
