package main


import (
	"strings"
	"fmt"
	"errors"
	"strconv"
	"os"
	"encoding/json"
	"net"
)

const (
	PROTOCOL_VERSION = "3.8"
	CONFIG_FILE_NAME = "config.json"
)

type Settings struct {
	Ip_address string
	Port string
	Max_connections string
}




// The parameter protocol_version is meant to be a string that is the protocol version you would like to get the message of.
// It returns the message and then an error if one is present.
func getProtocolVersionMessage(protocol_version string) (string, error) {
	if protocol_version == "" {
		return "", errors.New("The protocol version was an empty string")
	}


	if strings.Count(protocol_version, ".") > 1 {
		return "", errors.New("Too many (.) period characters! Expected just one")
	}

	if !strings.Contains(protocol_version, ".") {
		return "", errors.New("Protocol version invalid. Does not contain a period (.) Expected one")
	}

	const (
		MAX_PADDING = 3
	)

	splitString := strings.Split(protocol_version, ".")

	majorVersion := splitString[0]
	// It is needed to check if the major version and minor version are indeed integers.
	if _, err := strconv.Atoi(majorVersion); err != nil {
		// If err is not nil, strconv is not an integer.
		return "", errors.New("Version number to the left of the period (.) is not a valid integer")
	}
	if len(majorVersion) > MAX_PADDING {
		return "", errors.New("The version number to the left of the (.) had more than the max padding of characters!")
	}

	minorVersion := splitString[1]
	if _, err := strconv.Atoi(majorVersion); err != nil {
		// If err is not nil, strconv is not an integer.
		return "", errors.New("Version number to the right of the period (.) is not a valid integer")
	}
	if len(minorVersion) > MAX_PADDING {
		return "", errors.New("The version number to the right of the (.) had more than the max padding 3 characters!")
	}

	minorPaddingLength := MAX_PADDING - len(minorVersion)
	majorPaddingLength := MAX_PADDING - len(majorVersion)

	minorPadding := strings.Repeat("0", minorPaddingLength)
	majorPadding := strings.Repeat("0", majorPaddingLength)

	return fmt.Sprintf("RFB %s%s.%s%s\n", majorPadding, majorVersion, minorPadding, minorVersion), nil
}

func doProtocolVersionHandshake(conn net.Conn) error {
	// Not error checking right now.
	// TODO: Set up error checking
	protocol_version_message, _ := getProtocolVersionMessage(PROTOCOL_VERSION)
	fmt.Fprint(conn, protocol_version_message)
	return nil
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


func loadConfigFromFile(file *os.File) (settings *Settings, err error) {
	if file == nil {
		return nil, errors.New("loadConfigFromFile: file is nil!")
	}
	jsonParser := json.NewDecoder(file)
	if err := jsonParser.Decode(&settings); err != nil {
		errorMessage := fmt.Sprintf("Could not decode file %s, check if it is valid JSON. Error message: %q\n", CONFIG_FILE_NAME, err)
		return nil, errors.New(errorMessage)
	}
	return settings, nil
}

// Sets up the default settings. To be used as a fallback for when settings cannot be obtained from the CONFIG_FILE_NAME file
// error is there in case it is needed in the future.
func loadDefaultConfig() (settings *Settings, err error) {
	const (
		DEFAULT_IP_ADDRESS = "127.0.0.1"
		DEFAULT_PORT = ":8000"
		DEFAULT_MAX_CONNECTIONS = "50"
	)
	settings = &Settings{DEFAULT_IP_ADDRESS, DEFAULT_PORT, DEFAULT_MAX_CONNECTIONS}
	return settings, nil
}

func handleConnection(connection net.Conn) {
	fmt.Print("Connection is set up... Now listening...")
	defer connection.Close()
	doProtocolVersionHandshake(connection)
}

func main() {
	configFile, err := os.Open(CONFIG_FILE_NAME)
	var settings *Settings
	if err != nil {
		fmt.Errorf("Could not open the file %s. Are you sure it's really there? Reason: %q\n", CONFIG_FILE_NAME)
		fmt.Errorf("Going to be using the default config, instead...\n")
		settings, _ = loadDefaultConfig()
	} else {
		settings, _ = loadConfigFromFile(configFile)
	}
	fmt.Printf("Settings loaded! They are: %q\n", *settings)
	ln, err := net.Listen("tcp", settings.Port)

	if err != nil {
		// handle error
	}
	for {
		connection, err := ln.Accept()
		if err != nil {

		}
		go handleConnection(connection)
	}
}

