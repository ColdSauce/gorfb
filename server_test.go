package main

func testProtocolVersionMessage() {
	protocol_version := "3.3"
	getProtocolVersionMessage(protocol_version)

	protocol_version = "123.999"
	getProtocolVersionMessage(protocol_version)

	protocol_version = "11.99"
	getProtocolVersionMessage(protocol_version)
}
