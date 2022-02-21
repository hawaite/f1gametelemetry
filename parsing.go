package f1gametelemetry

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func ParsePacketHeader(packetData []byte) PacketHeader {
	headerBuffer := bytes.NewBuffer(packetData[0:])
	headerReader := bufio.NewReader(headerBuffer)

	var parsedHeader PacketHeader
	binary.Read(headerReader, binary.LittleEndian, &parsedHeader)
	return parsedHeader
}

func ParseCarTelemetryPacket(packetData []byte) CarTelemetryPacket {
	packetBuffer := bytes.NewBuffer(packetData[0:])
	packetReader := bufio.NewReader(packetBuffer)

	var parsedCarTelemetryPacket CarTelemetryPacket
	binary.Read(packetReader, binary.LittleEndian, &parsedCarTelemetryPacket)
	return parsedCarTelemetryPacket
}
