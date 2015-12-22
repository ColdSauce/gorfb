package main
import "testing"

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

