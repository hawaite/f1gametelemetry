package f1gametelemetry

import "encoding/json"

const (
	MotionPacketId              = 0
	SessionPacketId             = 1
	LapDataPacketId             = 2
	EventPacketId               = 3
	ParticipantsPacketId        = 4
	CarSetupsPacketId           = 5
	CarTelemetryPacketId        = 6
	CarStatusPacketId           = 7
	FinalClassificationPacketId = 8
	LobbyInfoPacketId           = 9
)

type PacketHeader struct {
	PacketFormat            uint16
	GameMajorVersion        uint8
	GameMinorVersion        uint8
	PacketVersion           uint8
	PacketId                uint8
	SessionId               uint64
	SessionTime             float32
	FrameIdentifier         uint32
	PlayerCarIndex          uint8
	SecondaryPlayerCarIndex uint8
}

type CarTelemetryData struct {
	Speed                  uint16
	Throttle               float32
	Steer                  float32
	Brake                  float32
	Clutch                 uint8
	Gear                   int8
	EngineRPM              uint16
	DRS                    uint8
	RevLightPercent        uint8
	BrakesTemperature      U16Wheel
	TyreSurfaceTemperature U8Wheel
	TyreInnerTemperature   U8Wheel
	EngineTemperature      uint16
	TyrePressure           F32Wheel
	SurfaceType            U8Wheel
}

type CarTelemetryPacket struct {
	PacketHeader
	CarTelemetryData
	ButtonStatus              uint32
	MFDPanelIndex             uint8
	SecondPlayerMFDPanelIndex uint8
	SuggestedGear             int8
}

func (ctp CarTelemetryPacket) String() string {
	byteArray, _ := json.MarshalIndent(ctp, "", "  ")
	return string(byteArray[:])
}

func (ph PacketHeader) String() string {
	byteArray, _ := json.MarshalIndent(ph, "", "  ")
	return string(byteArray[:])
}

func (ctd CarTelemetryData) String() string {
	byteArray, _ := json.MarshalIndent(ctd, "", "  ")
	return string(byteArray[:])
}
