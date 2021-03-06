// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/mactime.api.json

/*
 Package mactime is a generated from VPP binary API module 'mactime'.

 It contains following objects:
	  4 messages
	  1 type
	  2 services

*/
package mactime

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
//	    "mactime_add_del_range": {
//	        "reply": "mactime_add_del_range_reply"
//	    },
//	    "mactime_enable_disable": {
//	        "reply": "mactime_enable_disable_reply"
//	    }
//	},
//
type Services interface {
	MactimeAddDelRange(*MactimeAddDelRange) (*MactimeAddDelRangeReply, error)
	MactimeEnableDisable(*MactimeEnableDisable) (*MactimeEnableDisableReply, error)
}

/* Types */

// TimeRange represents VPP binary API type 'time_range':
//
//	"time_range",
//	[
//	    "f64",
//	    "start"
//	],
//	[
//	    "f64",
//	    "end"
//	],
//	{
//	    "crc": "0x941c191b"
//	}
//
type TimeRange struct {
	Start float64
	End   float64
}

func (*TimeRange) GetTypeName() string {
	return "time_range"
}
func (*TimeRange) GetCrcString() string {
	return "941c191b"
}

/* Messages */

// MactimeEnableDisable represents VPP binary API message 'mactime_enable_disable':
//
//	"mactime_enable_disable",
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
type MactimeEnableDisable struct {
	EnableDisable uint8
	SwIfIndex     uint32
}

func (*MactimeEnableDisable) GetMessageName() string {
	return "mactime_enable_disable"
}
func (*MactimeEnableDisable) GetCrcString() string {
	return "57298519"
}
func (*MactimeEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MactimeEnableDisableReply represents VPP binary API message 'mactime_enable_disable_reply':
//
//	"mactime_enable_disable_reply",
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
type MactimeEnableDisableReply struct {
	Retval int32
}

func (*MactimeEnableDisableReply) GetMessageName() string {
	return "mactime_enable_disable_reply"
}
func (*MactimeEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MactimeEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MactimeAddDelRange represents VPP binary API message 'mactime_add_del_range':
//
//	"mactime_add_del_range",
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
//	    "drop"
//	],
//	[
//	    "u8",
//	    "allow"
//	],
//	[
//	    "u8",
//	    "mac_address",
//	    6
//	],
//	[
//	    "u8",
//	    "device_name",
//	    64
//	],
//	[
//	    "u32",
//	    "count"
//	],
//	[
//	    "vl_api_time_range_t",
//	    "ranges",
//	    0,
//	    "count"
//	],
//	{
//	    "crc": "0x2b1821d5"
//	}
//
type MactimeAddDelRange struct {
	IsAdd      uint8
	Drop       uint8
	Allow      uint8
	MacAddress []byte `struc:"[6]byte"`
	DeviceName []byte `struc:"[64]byte"`
	Count      uint32 `struc:"sizeof=Ranges"`
	Ranges     []TimeRange
}

func (*MactimeAddDelRange) GetMessageName() string {
	return "mactime_add_del_range"
}
func (*MactimeAddDelRange) GetCrcString() string {
	return "2b1821d5"
}
func (*MactimeAddDelRange) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MactimeAddDelRangeReply represents VPP binary API message 'mactime_add_del_range_reply':
//
//	"mactime_add_del_range_reply",
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
type MactimeAddDelRangeReply struct {
	Retval int32
}

func (*MactimeAddDelRangeReply) GetMessageName() string {
	return "mactime_add_del_range_reply"
}
func (*MactimeAddDelRangeReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MactimeAddDelRangeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*MactimeEnableDisable)(nil), "mactime.MactimeEnableDisable")
	api.RegisterMessage((*MactimeEnableDisableReply)(nil), "mactime.MactimeEnableDisableReply")
	api.RegisterMessage((*MactimeAddDelRange)(nil), "mactime.MactimeAddDelRange")
	api.RegisterMessage((*MactimeAddDelRangeReply)(nil), "mactime.MactimeAddDelRangeReply")
}
