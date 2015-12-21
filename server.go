package main


import (
	"strings"
	"fmt"
)

const (
	PROTOCOL_VERSION = "3.8"
)



func getProtocolVersionMessage(protocol_version string) string {
	const (
		MAX_PADDING = 3
	)
	splitString := strings.Split(protocol_version, ".")
	majorVersion := splitString[0]
	minorVersion := splitString[1]

	minorPaddingLength := MAX_PADDING - len(minorVersion)
	majorPaddingLength := MAX_PADDING - len(majorVersion)


	minorPadding := strings.Repeat("0", minorPaddingLength)
	majorPadding := strings.Repeat("0", majorPaddingLength)

	return fmt.Sprintf("RFB %s%s.%s%s\n", majorPadding, majorVersion, minorPadding, minorVersion)
}

func doProtocolVersionHandshake() string {
	// Nothing for right now...
	return "nothing #yoloswag.. soon this will be awesome..!"
}

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
	fmt.Println(doProtocolVersionHandshake())
}

