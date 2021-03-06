// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/ip_types.api.json

/*
 Package ip_types is a generated from VPP binary API module 'ip_types'.

 It contains following objects:
	  5 types
	  2 aliases
	  1 enum
	  1 union

*/
package ip_types

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

/* Enums */

// AddressFamily represents VPP binary API enum 'address_family':
//
//	"address_family",
//	[
//	    "ADDRESS_IP4",
//	    0
//	],
//	[
//	    "ADDRESS_IP6",
//	    1
//	],
//	{
//	    "enumtype": "u32"
//	}
//
type AddressFamily uint32

const (
	ADDRESS_IP4 AddressFamily = 0
	ADDRESS_IP6 AddressFamily = 1
)

/* Aliases */

// IP4Address represents VPP binary API alias 'ip4_address':
//
//	"ip4_address": {
//	    "length": 4,
//	    "type": "u8"
//	}
//
type IP4Address [4]uint8

// IP6Address represents VPP binary API alias 'ip6_address':
//
//	"ip6_address": {
//	    "length": 16,
//	    "type": "u8"
//	},
//
type IP6Address [16]uint8

/* Types */

// Address represents VPP binary API type 'address':
//
//	"address",
//	[
//	    "vl_api_address_family_t",
//	    "af"
//	],
//	[
//	    "vl_api_address_union_t",
//	    "un"
//	],
//	{
//	    "crc": "0x09f11671"
//	}
//
type Address struct {
	Af AddressFamily
	Un AddressUnion
}

func (*Address) GetTypeName() string {
	return "address"
}
func (*Address) GetCrcString() string {
	return "09f11671"
}

// Prefix represents VPP binary API type 'prefix':
//
//	"prefix",
//	[
//	    "vl_api_address_t",
//	    "address"
//	],
//	[
//	    "u8",
//	    "address_length"
//	],
//	{
//	    "crc": "0x0403aebc"
//	}
//
type Prefix struct {
	Address       Address
	AddressLength uint8
}

func (*Prefix) GetTypeName() string {
	return "prefix"
}
func (*Prefix) GetCrcString() string {
	return "0403aebc"
}

// Mprefix represents VPP binary API type 'mprefix':
//
//	"mprefix",
//	[
//	    "vl_api_address_family_t",
//	    "af"
//	],
//	[
//	    "u16",
//	    "grp_address_length"
//	],
//	[
//	    "vl_api_address_union_t",
//	    "grp_address"
//	],
//	[
//	    "vl_api_address_union_t",
//	    "src_address"
//	],
//	{
//	    "crc": "0x1c4cba05"
//	}
//
type Mprefix struct {
	Af               AddressFamily
	GrpAddressLength uint16
	GrpAddress       AddressUnion
	SrcAddress       AddressUnion
}

func (*Mprefix) GetTypeName() string {
	return "mprefix"
}
func (*Mprefix) GetCrcString() string {
	return "1c4cba05"
}

// IP6Prefix represents VPP binary API type 'ip6_prefix':
//
//	"ip6_prefix",
//	[
//	    "vl_api_ip6_address_t",
//	    "prefix"
//	],
//	[
//	    "u8",
//	    "len"
//	],
//	{
//	    "crc": "0x779fd64f"
//	}
//
type IP6Prefix struct {
	Prefix IP6Address
	Len    uint8
}

func (*IP6Prefix) GetTypeName() string {
	return "ip6_prefix"
}
func (*IP6Prefix) GetCrcString() string {
	return "779fd64f"
}

// IP4Prefix represents VPP binary API type 'ip4_prefix':
//
//	"ip4_prefix",
//	[
//	    "vl_api_ip4_address_t",
//	    "prefix"
//	],
//	[
//	    "u8",
//	    "len"
//	],
//	{
//	    "crc": "0xea8dc11d"
//	}
//
type IP4Prefix struct {
	Prefix IP4Address
	Len    uint8
}

func (*IP4Prefix) GetTypeName() string {
	return "ip4_prefix"
}
func (*IP4Prefix) GetCrcString() string {
	return "ea8dc11d"
}

/* Unions */

// AddressUnion represents VPP binary API union 'address_union':
//
//	"address_union",
//	[
//	    "vl_api_ip4_address_t",
//	    "ip4"
//	],
//	[
//	    "vl_api_ip6_address_t",
//	    "ip6"
//	],
//	{
//	    "crc": "0xd68a2fb4"
//	}
//
type AddressUnion struct {
	Union_data [16]byte
}

func (*AddressUnion) GetTypeName() string {
	return "address_union"
}
func (*AddressUnion) GetCrcString() string {
	return "d68a2fb4"
}

func (u *AddressUnion) SetIP4(a IP4Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.Union_data[:], b.Bytes())
}
func (u *AddressUnion) GetIP4() (a IP4Address) {
	var b = bytes.NewReader(u.Union_data[:])
	struc.Unpack(b, &a)
	return
}

func (u *AddressUnion) SetIP6(a IP6Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.Union_data[:], b.Bytes())
}
func (u *AddressUnion) GetIP6() (a IP6Address) {
	var b = bytes.NewReader(u.Union_data[:])
	struc.Unpack(b, &a)
	return
}

func init() {
}
