// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/ethernet_types.api.json

/*
 Package ethernet_types is a generated from VPP binary API module 'ethernet_types'.

 It contains following objects:
	  1 type

*/
package ethernet_types

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

/* Types */

// MacAddress represents VPP binary API type 'mac_address':
//
//	"mac_address",
//	[
//	    "u8",
//	    "bytes",
//	    6
//	],
//	{
//	    "crc": "0xefdbdddc"
//	}
//
type MacAddress struct {
	Bytes []byte `struc:"[6]byte"`
}

func (*MacAddress) GetTypeName() string {
	return "mac_address"
}
func (*MacAddress) GetCrcString() string {
	return "efdbdddc"
}

func init() {
}
