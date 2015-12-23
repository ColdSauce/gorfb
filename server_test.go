package main
import (
	"testing"
	"os"
	"fmt"
)
/*
	Helper functions
 */

// TODO: Refactor this! Not very clean...
func testLoadConfigFromFile(file_name string, test *testing.T) {
	configFileName := file_name
	configFile, err := os.Open(configFileName)
	if err != nil {
		//TODO: Fix this terrible code.
		actual_settings, err := loadDefaultConfig()
		if err != nil {
			test.Error("Could not load default config! Error: %q", err)
		}
		const (
			EXPECTED_IP = "127.0.0.1"
			EXPECTED_PORT = ":8000"
			EXPECTED_MAX_CONNECTIONS = "50"
		)
		// I made this a memory address instead of just a regular old struct so that the if statement following
		// it is consistent for the condition. Both variables are dereferenced.
		expected_settings := &Settings{EXPECTED_IP, EXPECTED_PORT, EXPECTED_MAX_CONNECTIONS}

		// Both are pointers to memory locations. Need to dereference them in order to check if equal.
		if *actual_settings != *expected_settings {
			test.Error("Default values in struct are not the same as the expected default values!")
		}


	} else {
		actual_settings, err := loadConfigFromFile(configFile)
		if err != nil {
			errorMessage := fmt.Sprintf("Could not load config from file! Error: %q", err)
			test.Error(errorMessage)
		}
		// These will ALWAYS be the same. No matter if the actual config changes. If you change the config_test
		// file, CHANGE THESE! Or else tests will fail. I'm not sure how to automate that... Perhaps it is best
		// to make the config_test.json file read only...
		const (
			EXPECTED_IP = "127.0.0.1"
			EXPECTED_PORT = ":8080"
			EXPECTED_MAX_CONNECTIONS = "50"
		)
		// I made this a memory address instead of just a regular old struct so that the if statement following
		// it is consistent for the condition. Both variables are dereferenced.
		expected_settings := &Settings{EXPECTED_IP, EXPECTED_PORT, EXPECTED_MAX_CONNECTIONS}

		// Both are pointers to memory locations. Need to dereference them in order to check if equal.
		if *actual_settings != *expected_settings {
			test.Error("Config values in struct are not the same as the expected config values!")
		}
	}
}


/*
	Protocol version tests
 */

/*
	Positive tests:
 */

func TestGetProtocolVersionMessageOneDigit(test *testing.T) {
	protocol_version := "3.3"
	expected_result := "RFB 003.003\n"
	actual_result, err :=   getProtocolVersionMessage(protocol_version)

	if err != nil {
		test.Error(err)
	}

	if actual_result != expected_result {
		test.Errorf("Test failed. Actual result %s does not match expected result %s.", actual_result, expected_result)
	}
}

func TestGetProtocolVersionMessageThreeDigits(test *testing.T) {
	protocol_version := "123.999"
	expected_result := "RFB 123.999\n"
	actual_result, err :=   getProtocolVersionMessage(protocol_version)

	if err != nil {
		test.Error(err)
	}

	if actual_result != expected_result {
		test.Errorf("Test failed. Actual result %s does not match expected result %s.", actual_result, expected_result)
	}
}

func TestGetProtocolVersionMessageTwoDigits(test *testing.T) {
	protocol_version := "11.99"
	expected_result := "RFB 011.099\n"
	actual_result, err :=   getProtocolVersionMessage(protocol_version)

	if err != nil {
		test.Error(err)
	}

	if actual_result != expected_result {
		test.Errorf("Test failed. Actual result %s does not match expected result %s.", actual_result, expected_result)
	}
}

/*
	Negative tests:
 */

func TestGetProtocolVersionMessageEmptyString(test *testing.T) {
	protocol_version := ""
	expected_result := ""
	actual_result, err := getProtocolVersionMessage(protocol_version)

	if err == nil {
		test.Error("Test failed! There was no error!")
	}

	if actual_result != expected_result {
		test.Error("Test failed. Actual result (empty string) does not match expected result (empty string).")
	}
}

func TestGetProtocolVersionMessageNoDot(test *testing.T) {
	protocol_version := "23"
	expected_result := ""
	actual_result, err := getProtocolVersionMessage(protocol_version)

	if err == nil {
		test.Error("Test failed! There was no error!")
	}

	if actual_result != expected_result {
		test.Error("Test failed. Actual result (empty string) does not match expected result (empty string).")
	}
}

func TestGetProtocolVersionMessageTooBig(test *testing.T) {
	protocol_version := "234234234234234234.23234234234"
	expected_result := ""
	actual_result, err := getProtocolVersionMessage(protocol_version)

	if err == nil {
		test.Error("Test failed! There was no error!")
	}

	if actual_result != expected_result {
		test.Error("Test failed. Actual result (empty string) does not match expected result (empty string).")
	}
}

func TestGetProtocolVersionMessageNotDigits(test *testing.T) {
	protocol_version := "sdf.0sd"
	expected_result := ""
	actual_result, err := getProtocolVersionMessage(protocol_version)

	if err == nil {
		test.Error("Test failed! There was no error!")
	}

	if actual_result != expected_result {
		test.Error("Test failed. Actual result (empty string) does not match expected result (empty string).")
	}
}

func TestGetProtocolVersionMessageManyDots(test *testing.T) {
	protocol_version := "3.3.."
	expected_result := ""
	actual_result, err := getProtocolVersionMessage(protocol_version)

	if err == nil {
		test.Error("Test failed! There was no error!")
	}

	if actual_result != expected_result {
		test.Error("Test failed. Actual result (empty string) does not match expected result (empty string).")
	}
}

/*
	Tests for config loading
 */


func TestLoadConfigFromInvalidFile(test *testing.T) {
	testLoadConfigFromFile("some_random_file.json", test)
}

func TestLoadConfigFromValid(test *testing.T) {
	testLoadConfigFromFile("config_test.json", test)
}
