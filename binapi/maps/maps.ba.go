// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/map.api.json

/*
 Package maps is a generated from VPP binary API module 'map'.

 It contains following objects:
	 28 messages
	  5 types
	  2 aliases
	  1 enum
	  1 union
	 14 services

*/
package maps

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
//	    "map_param_set_fragmentation": {
//	        "reply": "map_param_set_fragmentation_reply"
//	    },
//	    "map_param_add_del_pre_resolve": {
//	        "reply": "map_param_add_del_pre_resolve_reply"
//	    },
//	    "map_rule_dump": {
//	        "reply": "map_rule_details",
//	        "stream": true
//	    },
//	    "map_param_set_icmp6": {
//	        "reply": "map_param_set_icmp6_reply"
//	    },
//	    "map_add_del_rule": {
//	        "reply": "map_add_del_rule_reply"
//	    },
//	    "map_domain_dump": {
//	        "reply": "map_domain_details",
//	        "stream": true
//	    },
//	    "map_param_get": {
//	        "reply": "map_param_get_reply"
//	    },
//	    "map_param_set_icmp": {
//	        "reply": "map_param_set_icmp_reply"
//	    },
//	    "map_add_domain": {
//	        "reply": "map_add_domain_reply"
//	    },
//	    "map_summary_stats": {
//	        "reply": "map_summary_stats_reply"
//	    },
//	    "map_param_set_traffic_class": {
//	        "reply": "map_param_set_traffic_class_reply"
//	    },
//	    "map_del_domain": {
//	        "reply": "map_del_domain_reply"
//	    },
//	    "map_param_set_reassembly": {
//	        "reply": "map_param_set_reassembly_reply"
//	    },
//	    "map_param_set_security_check": {
//	        "reply": "map_param_set_security_check_reply"
//	    }
//	},
//
type Services interface {
	DumpMapDomain(*MapDomainDump) ([]*MapDomainDetails, error)
	DumpMapRule(*MapRuleDump) ([]*MapRuleDetails, error)
	MapAddDelRule(*MapAddDelRule) (*MapAddDelRuleReply, error)
	MapAddDomain(*MapAddDomain) (*MapAddDomainReply, error)
	MapDelDomain(*MapDelDomain) (*MapDelDomainReply, error)
	MapParamAddDelPreResolve(*MapParamAddDelPreResolve) (*MapParamAddDelPreResolveReply, error)
	MapParamGet(*MapParamGet) (*MapParamGetReply, error)
	MapParamSetFragmentation(*MapParamSetFragmentation) (*MapParamSetFragmentationReply, error)
	MapParamSetICMP(*MapParamSetICMP) (*MapParamSetICMPReply, error)
	MapParamSetICMP6(*MapParamSetICMP6) (*MapParamSetICMP6Reply, error)
	MapParamSetReassembly(*MapParamSetReassembly) (*MapParamSetReassemblyReply, error)
	MapParamSetSecurityCheck(*MapParamSetSecurityCheck) (*MapParamSetSecurityCheckReply, error)
	MapParamSetTrafficClass(*MapParamSetTrafficClass) (*MapParamSetTrafficClassReply, error)
	MapSummaryStats(*MapSummaryStats) (*MapSummaryStatsReply, error)
}

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

// IP6Address represents VPP binary API alias 'ip6_address':
//
//	"ip6_address": {
//	    "length": 16,
//	    "type": "u8"
//	},
//
type IP6Address [16]uint8

// IP4Address represents VPP binary API alias 'ip4_address':
//
//	"ip4_address": {
//	    "length": 4,
//	    "type": "u8"
//	}
//
type IP4Address [4]uint8

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

/* Messages */

// MapAddDomain represents VPP binary API message 'map_add_domain':
//
//	"map_add_domain",
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
//	    "vl_api_ip6_prefix_t",
//	    "ip6_prefix"
//	],
//	[
//	    "vl_api_ip4_prefix_t",
//	    "ip4_prefix"
//	],
//	[
//	    "vl_api_ip6_prefix_t",
//	    "ip6_src"
//	],
//	[
//	    "u8",
//	    "ea_bits_len"
//	],
//	[
//	    "u8",
//	    "psid_offset"
//	],
//	[
//	    "u8",
//	    "psid_length"
//	],
//	[
//	    "bool",
//	    "is_translation"
//	],
//	[
//	    "bool",
//	    "is_rfc6052"
//	],
//	[
//	    "u16",
//	    "mtu"
//	],
//	{
//	    "crc": "0x7a64714e"
//	}
//
type MapAddDomain struct {
	IP6Prefix     IP6Prefix
	IP4Prefix     IP4Prefix
	IP6Src        IP6Prefix
	EaBitsLen     uint8
	PsidOffset    uint8
	PsidLength    uint8
	IsTranslation bool
	IsRfc6052     bool
	Mtu           uint16
}

func (*MapAddDomain) GetMessageName() string {
	return "map_add_domain"
}
func (*MapAddDomain) GetCrcString() string {
	return "7a64714e"
}
func (*MapAddDomain) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapAddDomainReply represents VPP binary API message 'map_add_domain_reply':
//
//	"map_add_domain_reply",
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
//	    "index"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0x3e6d4e2c"
//	}
//
type MapAddDomainReply struct {
	Index  uint32
	Retval int32
}

func (*MapAddDomainReply) GetMessageName() string {
	return "map_add_domain_reply"
}
func (*MapAddDomainReply) GetCrcString() string {
	return "3e6d4e2c"
}
func (*MapAddDomainReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapDelDomain represents VPP binary API message 'map_del_domain':
//
//	"map_del_domain",
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
//	    "index"
//	],
//	{
//	    "crc": "0x8ac76db6"
//	}
//
type MapDelDomain struct {
	Index uint32
}

func (*MapDelDomain) GetMessageName() string {
	return "map_del_domain"
}
func (*MapDelDomain) GetCrcString() string {
	return "8ac76db6"
}
func (*MapDelDomain) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapDelDomainReply represents VPP binary API message 'map_del_domain_reply':
//
//	"map_del_domain_reply",
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
type MapDelDomainReply struct {
	Retval int32
}

func (*MapDelDomainReply) GetMessageName() string {
	return "map_del_domain_reply"
}
func (*MapDelDomainReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapDelDomainReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapAddDelRule represents VPP binary API message 'map_add_del_rule':
//
//	"map_add_del_rule",
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
//	    "index"
//	],
//	[
//	    "bool",
//	    "is_add"
//	],
//	[
//	    "vl_api_ip6_address_t",
//	    "ip6_dst"
//	],
//	[
//	    "u16",
//	    "psid"
//	],
//	{
//	    "crc": "0xe6132040"
//	}
//
type MapAddDelRule struct {
	Index  uint32
	IsAdd  bool
	IP6Dst IP6Address
	Psid   uint16
}

func (*MapAddDelRule) GetMessageName() string {
	return "map_add_del_rule"
}
func (*MapAddDelRule) GetCrcString() string {
	return "e6132040"
}
func (*MapAddDelRule) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapAddDelRuleReply represents VPP binary API message 'map_add_del_rule_reply':
//
//	"map_add_del_rule_reply",
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
type MapAddDelRuleReply struct {
	Retval int32
}

func (*MapAddDelRuleReply) GetMessageName() string {
	return "map_add_del_rule_reply"
}
func (*MapAddDelRuleReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapAddDelRuleReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapDomainDump represents VPP binary API message 'map_domain_dump':
//
//	"map_domain_dump",
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
type MapDomainDump struct{}

func (*MapDomainDump) GetMessageName() string {
	return "map_domain_dump"
}
func (*MapDomainDump) GetCrcString() string {
	return "51077d14"
}
func (*MapDomainDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapDomainDetails represents VPP binary API message 'map_domain_details':
//
//	"map_domain_details",
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
//	    "domain_index"
//	],
//	[
//	    "vl_api_ip6_prefix_t",
//	    "ip6_prefix"
//	],
//	[
//	    "vl_api_ip4_prefix_t",
//	    "ip4_prefix"
//	],
//	[
//	    "vl_api_ip6_prefix_t",
//	    "ip6_src"
//	],
//	[
//	    "u8",
//	    "ea_bits_len"
//	],
//	[
//	    "u8",
//	    "psid_offset"
//	],
//	[
//	    "u8",
//	    "psid_length"
//	],
//	[
//	    "u8",
//	    "flags"
//	],
//	[
//	    "u16",
//	    "mtu"
//	],
//	[
//	    "bool",
//	    "is_translation"
//	],
//	{
//	    "crc": "0x7a986fe6"
//	}
//
type MapDomainDetails struct {
	DomainIndex   uint32
	IP6Prefix     IP6Prefix
	IP4Prefix     IP4Prefix
	IP6Src        IP6Prefix
	EaBitsLen     uint8
	PsidOffset    uint8
	PsidLength    uint8
	Flags         uint8
	Mtu           uint16
	IsTranslation bool
}

func (*MapDomainDetails) GetMessageName() string {
	return "map_domain_details"
}
func (*MapDomainDetails) GetCrcString() string {
	return "7a986fe6"
}
func (*MapDomainDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapRuleDump represents VPP binary API message 'map_rule_dump':
//
//	"map_rule_dump",
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
//	    "domain_index"
//	],
//	{
//	    "crc": "0xe43e6ff6"
//	}
//
type MapRuleDump struct {
	DomainIndex uint32
}

func (*MapRuleDump) GetMessageName() string {
	return "map_rule_dump"
}
func (*MapRuleDump) GetCrcString() string {
	return "e43e6ff6"
}
func (*MapRuleDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapRuleDetails represents VPP binary API message 'map_rule_details':
//
//	"map_rule_details",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "vl_api_ip6_address_t",
//	    "ip6_dst"
//	],
//	[
//	    "u16",
//	    "psid"
//	],
//	{
//	    "crc": "0x4f932665"
//	}
//
type MapRuleDetails struct {
	IP6Dst IP6Address
	Psid   uint16
}

func (*MapRuleDetails) GetMessageName() string {
	return "map_rule_details"
}
func (*MapRuleDetails) GetCrcString() string {
	return "4f932665"
}
func (*MapRuleDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapSummaryStats represents VPP binary API message 'map_summary_stats':
//
//	"map_summary_stats",
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
type MapSummaryStats struct{}

func (*MapSummaryStats) GetMessageName() string {
	return "map_summary_stats"
}
func (*MapSummaryStats) GetCrcString() string {
	return "51077d14"
}
func (*MapSummaryStats) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapSummaryStatsReply represents VPP binary API message 'map_summary_stats_reply':
//
//	"map_summary_stats_reply",
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
//	    "u64",
//	    "total_bindings"
//	],
//	[
//	    "u64",
//	    "total_pkts",
//	    2
//	],
//	[
//	    "u64",
//	    "total_bytes",
//	    2
//	],
//	[
//	    "u64",
//	    "total_ip4_fragments"
//	],
//	[
//	    "u64",
//	    "total_security_check",
//	    2
//	],
//	{
//	    "crc": "0x0e4ace0e"
//	}
//
type MapSummaryStatsReply struct {
	Retval             int32
	TotalBindings      uint64
	TotalPkts          []uint64 `struc:"[2]uint64"`
	TotalBytes         []uint64 `struc:"[2]uint64"`
	TotalIP4Fragments  uint64
	TotalSecurityCheck []uint64 `struc:"[2]uint64"`
}

func (*MapSummaryStatsReply) GetMessageName() string {
	return "map_summary_stats_reply"
}
func (*MapSummaryStatsReply) GetCrcString() string {
	return "0e4ace0e"
}
func (*MapSummaryStatsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamSetFragmentation represents VPP binary API message 'map_param_set_fragmentation':
//
//	"map_param_set_fragmentation",
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
//	    "bool",
//	    "inner"
//	],
//	[
//	    "bool",
//	    "ignore_df"
//	],
//	{
//	    "crc": "0x9ff54d90"
//	}
//
type MapParamSetFragmentation struct {
	Inner    bool
	IgnoreDf bool
}

func (*MapParamSetFragmentation) GetMessageName() string {
	return "map_param_set_fragmentation"
}
func (*MapParamSetFragmentation) GetCrcString() string {
	return "9ff54d90"
}
func (*MapParamSetFragmentation) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamSetFragmentationReply represents VPP binary API message 'map_param_set_fragmentation_reply':
//
//	"map_param_set_fragmentation_reply",
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
type MapParamSetFragmentationReply struct {
	Retval int32
}

func (*MapParamSetFragmentationReply) GetMessageName() string {
	return "map_param_set_fragmentation_reply"
}
func (*MapParamSetFragmentationReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapParamSetFragmentationReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamSetICMP represents VPP binary API message 'map_param_set_icmp':
//
//	"map_param_set_icmp",
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
//	    "vl_api_ip4_address_t",
//	    "ip4_err_relay_src"
//	],
//	{
//	    "crc": "0x4c0a4fd2"
//	}
//
type MapParamSetICMP struct {
	IP4ErrRelaySrc IP4Address
}

func (*MapParamSetICMP) GetMessageName() string {
	return "map_param_set_icmp"
}
func (*MapParamSetICMP) GetCrcString() string {
	return "4c0a4fd2"
}
func (*MapParamSetICMP) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamSetICMPReply represents VPP binary API message 'map_param_set_icmp_reply':
//
//	"map_param_set_icmp_reply",
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
type MapParamSetICMPReply struct {
	Retval int32
}

func (*MapParamSetICMPReply) GetMessageName() string {
	return "map_param_set_icmp_reply"
}
func (*MapParamSetICMPReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapParamSetICMPReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamSetICMP6 represents VPP binary API message 'map_param_set_icmp6':
//
//	"map_param_set_icmp6",
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
//	    "bool",
//	    "enable_unreachable"
//	],
//	{
//	    "crc": "0x5d01f8c1"
//	}
//
type MapParamSetICMP6 struct {
	EnableUnreachable bool
}

func (*MapParamSetICMP6) GetMessageName() string {
	return "map_param_set_icmp6"
}
func (*MapParamSetICMP6) GetCrcString() string {
	return "5d01f8c1"
}
func (*MapParamSetICMP6) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamSetICMP6Reply represents VPP binary API message 'map_param_set_icmp6_reply':
//
//	"map_param_set_icmp6_reply",
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
type MapParamSetICMP6Reply struct {
	Retval int32
}

func (*MapParamSetICMP6Reply) GetMessageName() string {
	return "map_param_set_icmp6_reply"
}
func (*MapParamSetICMP6Reply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapParamSetICMP6Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamAddDelPreResolve represents VPP binary API message 'map_param_add_del_pre_resolve':
//
//	"map_param_add_del_pre_resolve",
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
//	    "bool",
//	    "is_add"
//	],
//	[
//	    "vl_api_ip4_address_t",
//	    "ip4_nh_address"
//	],
//	[
//	    "vl_api_ip6_address_t",
//	    "ip6_nh_address"
//	],
//	{
//	    "crc": "0xea9a9a4a"
//	}
//
type MapParamAddDelPreResolve struct {
	IsAdd        bool
	IP4NhAddress IP4Address
	IP6NhAddress IP6Address
}

func (*MapParamAddDelPreResolve) GetMessageName() string {
	return "map_param_add_del_pre_resolve"
}
func (*MapParamAddDelPreResolve) GetCrcString() string {
	return "ea9a9a4a"
}
func (*MapParamAddDelPreResolve) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamAddDelPreResolveReply represents VPP binary API message 'map_param_add_del_pre_resolve_reply':
//
//	"map_param_add_del_pre_resolve_reply",
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
type MapParamAddDelPreResolveReply struct {
	Retval int32
}

func (*MapParamAddDelPreResolveReply) GetMessageName() string {
	return "map_param_add_del_pre_resolve_reply"
}
func (*MapParamAddDelPreResolveReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapParamAddDelPreResolveReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamSetReassembly represents VPP binary API message 'map_param_set_reassembly':
//
//	"map_param_set_reassembly",
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
//	    "bool",
//	    "is_ip6"
//	],
//	[
//	    "u16",
//	    "lifetime_ms"
//	],
//	[
//	    "u16",
//	    "pool_size"
//	],
//	[
//	    "u32",
//	    "buffers"
//	],
//	[
//	    "f64",
//	    "ht_ratio"
//	],
//	{
//	    "crc": "0x54172b10"
//	}
//
type MapParamSetReassembly struct {
	IsIP6      bool
	LifetimeMs uint16
	PoolSize   uint16
	Buffers    uint32
	HtRatio    float64
}

func (*MapParamSetReassembly) GetMessageName() string {
	return "map_param_set_reassembly"
}
func (*MapParamSetReassembly) GetCrcString() string {
	return "54172b10"
}
func (*MapParamSetReassembly) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamSetReassemblyReply represents VPP binary API message 'map_param_set_reassembly_reply':
//
//	"map_param_set_reassembly_reply",
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
type MapParamSetReassemblyReply struct {
	Retval int32
}

func (*MapParamSetReassemblyReply) GetMessageName() string {
	return "map_param_set_reassembly_reply"
}
func (*MapParamSetReassemblyReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapParamSetReassemblyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamSetSecurityCheck represents VPP binary API message 'map_param_set_security_check':
//
//	"map_param_set_security_check",
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
//	    "bool",
//	    "enable"
//	],
//	[
//	    "bool",
//	    "fragments"
//	],
//	{
//	    "crc": "0x6abe9836"
//	}
//
type MapParamSetSecurityCheck struct {
	Enable    bool
	Fragments bool
}

func (*MapParamSetSecurityCheck) GetMessageName() string {
	return "map_param_set_security_check"
}
func (*MapParamSetSecurityCheck) GetCrcString() string {
	return "6abe9836"
}
func (*MapParamSetSecurityCheck) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamSetSecurityCheckReply represents VPP binary API message 'map_param_set_security_check_reply':
//
//	"map_param_set_security_check_reply",
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
type MapParamSetSecurityCheckReply struct {
	Retval int32
}

func (*MapParamSetSecurityCheckReply) GetMessageName() string {
	return "map_param_set_security_check_reply"
}
func (*MapParamSetSecurityCheckReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapParamSetSecurityCheckReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamSetTrafficClass represents VPP binary API message 'map_param_set_traffic_class':
//
//	"map_param_set_traffic_class",
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
//	    "bool",
//	    "copy"
//	],
//	[
//	    "u8",
//	    "class"
//	],
//	{
//	    "crc": "0x007ee563"
//	}
//
type MapParamSetTrafficClass struct {
	Copy  bool
	Class uint8
}

func (*MapParamSetTrafficClass) GetMessageName() string {
	return "map_param_set_traffic_class"
}
func (*MapParamSetTrafficClass) GetCrcString() string {
	return "007ee563"
}
func (*MapParamSetTrafficClass) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamSetTrafficClassReply represents VPP binary API message 'map_param_set_traffic_class_reply':
//
//	"map_param_set_traffic_class_reply",
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
type MapParamSetTrafficClassReply struct {
	Retval int32
}

func (*MapParamSetTrafficClassReply) GetMessageName() string {
	return "map_param_set_traffic_class_reply"
}
func (*MapParamSetTrafficClassReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapParamSetTrafficClassReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapParamGet represents VPP binary API message 'map_param_get':
//
//	"map_param_get",
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
type MapParamGet struct{}

func (*MapParamGet) GetMessageName() string {
	return "map_param_get"
}
func (*MapParamGet) GetCrcString() string {
	return "51077d14"
}
func (*MapParamGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapParamGetReply represents VPP binary API message 'map_param_get_reply':
//
//	"map_param_get_reply",
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
//	    "u8",
//	    "frag_inner"
//	],
//	[
//	    "u8",
//	    "frag_ignore_df"
//	],
//	[
//	    "vl_api_ip4_address_t",
//	    "icmp_ip4_err_relay_src"
//	],
//	[
//	    "bool",
//	    "icmp6_enable_unreachable"
//	],
//	[
//	    "vl_api_ip4_address_t",
//	    "ip4_nh_address"
//	],
//	[
//	    "vl_api_ip6_address_t",
//	    "ip6_nh_address"
//	],
//	[
//	    "u16",
//	    "ip4_lifetime_ms"
//	],
//	[
//	    "u16",
//	    "ip4_pool_size"
//	],
//	[
//	    "u32",
//	    "ip4_buffers"
//	],
//	[
//	    "f64",
//	    "ip4_ht_ratio"
//	],
//	[
//	    "u16",
//	    "ip6_lifetime_ms"
//	],
//	[
//	    "u16",
//	    "ip6_pool_size"
//	],
//	[
//	    "u32",
//	    "ip6_buffers"
//	],
//	[
//	    "f64",
//	    "ip6_ht_ratio"
//	],
//	[
//	    "bool",
//	    "sec_check_enable"
//	],
//	[
//	    "bool",
//	    "sec_check_fragments"
//	],
//	[
//	    "bool",
//	    "tc_copy"
//	],
//	[
//	    "u8",
//	    "tc_class"
//	],
//	{
//	    "crc": "0xb40e9226"
//	}
//
type MapParamGetReply struct {
	Retval                 int32
	FragInner              uint8
	FragIgnoreDf           uint8
	ICMPIP4ErrRelaySrc     IP4Address
	ICMP6EnableUnreachable bool
	IP4NhAddress           IP4Address
	IP6NhAddress           IP6Address
	IP4LifetimeMs          uint16
	IP4PoolSize            uint16
	IP4Buffers             uint32
	IP4HtRatio             float64
	IP6LifetimeMs          uint16
	IP6PoolSize            uint16
	IP6Buffers             uint32
	IP6HtRatio             float64
	SecCheckEnable         bool
	SecCheckFragments      bool
	TcCopy                 bool
	TcClass                uint8
}

func (*MapParamGetReply) GetMessageName() string {
	return "map_param_get_reply"
}
func (*MapParamGetReply) GetCrcString() string {
	return "b40e9226"
}
func (*MapParamGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*MapAddDomain)(nil), "map.MapAddDomain")
	api.RegisterMessage((*MapAddDomainReply)(nil), "map.MapAddDomainReply")
	api.RegisterMessage((*MapDelDomain)(nil), "map.MapDelDomain")
	api.RegisterMessage((*MapDelDomainReply)(nil), "map.MapDelDomainReply")
	api.RegisterMessage((*MapAddDelRule)(nil), "map.MapAddDelRule")
	api.RegisterMessage((*MapAddDelRuleReply)(nil), "map.MapAddDelRuleReply")
	api.RegisterMessage((*MapDomainDump)(nil), "map.MapDomainDump")
	api.RegisterMessage((*MapDomainDetails)(nil), "map.MapDomainDetails")
	api.RegisterMessage((*MapRuleDump)(nil), "map.MapRuleDump")
	api.RegisterMessage((*MapRuleDetails)(nil), "map.MapRuleDetails")
	api.RegisterMessage((*MapSummaryStats)(nil), "map.MapSummaryStats")
	api.RegisterMessage((*MapSummaryStatsReply)(nil), "map.MapSummaryStatsReply")
	api.RegisterMessage((*MapParamSetFragmentation)(nil), "map.MapParamSetFragmentation")
	api.RegisterMessage((*MapParamSetFragmentationReply)(nil), "map.MapParamSetFragmentationReply")
	api.RegisterMessage((*MapParamSetICMP)(nil), "map.MapParamSetICMP")
	api.RegisterMessage((*MapParamSetICMPReply)(nil), "map.MapParamSetICMPReply")
	api.RegisterMessage((*MapParamSetICMP6)(nil), "map.MapParamSetICMP6")
	api.RegisterMessage((*MapParamSetICMP6Reply)(nil), "map.MapParamSetICMP6Reply")
	api.RegisterMessage((*MapParamAddDelPreResolve)(nil), "map.MapParamAddDelPreResolve")
	api.RegisterMessage((*MapParamAddDelPreResolveReply)(nil), "map.MapParamAddDelPreResolveReply")
	api.RegisterMessage((*MapParamSetReassembly)(nil), "map.MapParamSetReassembly")
	api.RegisterMessage((*MapParamSetReassemblyReply)(nil), "map.MapParamSetReassemblyReply")
	api.RegisterMessage((*MapParamSetSecurityCheck)(nil), "map.MapParamSetSecurityCheck")
	api.RegisterMessage((*MapParamSetSecurityCheckReply)(nil), "map.MapParamSetSecurityCheckReply")
	api.RegisterMessage((*MapParamSetTrafficClass)(nil), "map.MapParamSetTrafficClass")
	api.RegisterMessage((*MapParamSetTrafficClassReply)(nil), "map.MapParamSetTrafficClassReply")
	api.RegisterMessage((*MapParamGet)(nil), "map.MapParamGet")
	api.RegisterMessage((*MapParamGetReply)(nil), "map.MapParamGetReply")
}
