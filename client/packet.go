package client

import (
	"encoding/binary"
	"bytes"
	"fmt"
)

type Packet struct {
	BuildVersionNumber         uint16
	RawSequencePacket          uint8
	RawSessionState            uint8
	ViewedParticipantIndex     int8
	NumParticipants            int8
	UnfilteredThrottle         uint8
	UnfilteredBrake            uint8
	UnfilteredSteering         int8
	UnfilteredClutch           uint8
	RaceStateFlags             uint8
	LapsInEvent                uint8
	BestLapTime                float32
	LastLapTime                float32
	CurrentTime                float32
	SplitTimeAhead             float32
	SplitTimeBehind            float32
	SplitTime                  float32
	EventTimeRemaining         float32
	PersonalFastestLapTime     float32
	WorldFastestLapTime        float32
	CurrentSector1Time         float32
	CurrentSector2Time         float32
	CurrentSector3Time         float32
	FastestSector1Time         float32
	FastestSector2Time         float32
	FastestSector3Time         float32
	PersonalFastestSector1Time float32
	PersonalFastestSector2Time float32
	PersonalFastestSector3Time float32
	WorldFastestSector1Time    float32
	WorldFastestSector2Time    float32
	WorldFastestSector3Time    float32
	JoyPad                     uint16
	HighestFlag                uint8
	PitModeSchedule            uint8
	OilTempCelsius             int16
	OilPressureKPa             uint16
	WaterTempCelsius           int16
	WaterPressureKPa           uint16
	FuelPressureKPa            uint16
	CarFlags                   uint8
	FuelCapacity               uint8
	Brake                      uint8
	Throttle                   uint8
	Clutch                     uint8
	Steering                   int8
	FuelLevel                  float32
	Speed                      float32
	Rpm                        uint16
	MaxRpm                     uint16
	GearNumGears               uint8
	BoostAmount                uint8
	EnforcedPitStopLap         int8
	CrashState                 uint8
	OdometerKM                 float32
	OrientationX               float32
	OrientationY               float32
	OrientationZ               float32
	LocalVelocityX             float32
	LocalVelocityY             float32
	LocalVelocityZ             float32
	WorldVelocityX             float32
	WorldVelocityY             float32
	WorldVelocityZ             float32
	AngularVelocityX           float32
	AngularVelocityY           float32
	AngularVelocityZ           float32
	LocalAccelerationX         float32
	LocalAccelerationY         float32
	LocalAccelerationZ         float32
	WorldAccelerationX         float32
	WorldAccelerationY         float32
	WorldAccelerationZ         float32
	ExtentsCentreX             float32
	ExtentsCentreY             float32
	ExtentsCentreZ             float32
}

func (p Packet) GetSequenceNumber() uint8 {
	return (p.RawSequencePacket & 0xFC) >> 2
}

func (p Packet) GetPacketType() uint8 {
	return p.RawSequencePacket & 0x3
}

type GameState uint8

const (
	GameState_EXITED GameState = 0
	GameState_FRONT_END GameState = 1
	GameState_INGAME_PLAYING GameState = 2
	GameState_INGAME_PAUSED GameState = 3
)

type SessionState uint8

const (
	SessionState_INVALID SessionState = 0
	SessionState_PRACTICE SessionState = 1
	SessionState_TEST SessionState = 2
	SessionState_QUALIFY SessionState = 3
	SessionState_FORMATION_LAP SessionState = 4
	SessionState_RACE SessionState = 5
	SessionState_TIME_ATTACK SessionState = 6
)

func (p Packet) GetGameState() GameState {
	return GameState((p.RawSessionState) & 0x07)
}

func (p Packet) GetSessionState() SessionState {
	return SessionState((p.RawSessionState & 0x38) >> 2)
}

func Parse(raw_message []byte) (Packet, error) {
	packet := Packet{}
	buf := bytes.NewReader(raw_message)
	err := binary.Read(buf, binary.LittleEndian, &packet)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	return packet, err
}