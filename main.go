package main


import (
)



func doSecurityHandshake() {

	const (
		INVALID = 0
		NONE = 1
		VNC_AUTH = 2
		RA2 = 5
		RA2ne = 6
		Tight = 16
		Ultra = 17
		TLS = 18
		VENCRYPT = 19
		GTK_VNC_SASL = 20
		MD5_HASH = 21
		COLIN_DEAN_XVP = 22
	)

	type securityResult struct {
		message uint32 // can be either 0 or 1.
			       // 0 = OK
			       // 1 = Failed
	}

	type reasonForFailure struct {
		reasonLength uint32
		reasonString []byte
	}


}

func main() {


}

