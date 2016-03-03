package client

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	const filename = "sample_data/2016-03-02T19:02:33-06:00"

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error("read error", err)
	}

	packet, err := Parse(contents)
	if err != nil {
		t.Error("parse error", err)
	}

	var fmtTests = []struct {
		name     string
		getter   func(Packet) interface{}
		expected interface{}
	}{

		{"BuildVersionNumber", func(p Packet) interface{} {
			return p.BuildVersionNumber
		}, uint16(1182)},

		{"SequenceNumber", func(p Packet) interface{} {
			return p.GetSequenceNumber()
		}, uint8(15)},

		{"PacketType", func(p Packet) interface{} {
			return p.GetPacketType()
		}, uint8(0)},

		{"GameState", func(p Packet) interface{} {
			return p.GetGameState()
		}, GameState_INGAME_PLAYING},

		{"SessionState", func(p Packet) interface{} {
			return p.GetSessionState()
		}, SessionState_FORMATION_LAP},

		{"ViewedParticipantIndex", func(p Packet) interface{} {
			return p.ViewedParticipantIndex
		}, int8(0)},

		{"NumParticipants", func(p Packet) interface{} {
			return p.NumParticipants
		}, int8(16)},

		{"UnfilteredThrottle", func(p Packet) interface{} {
			return p.UnfilteredThrottle
		}, uint8(0)},

		{"UnfilteredBrake", func(p Packet) interface{} {
			return p.UnfilteredBrake
		}, uint8(0)},

		{"UnfilteredSteering", func(p Packet) interface{} {
			return p.UnfilteredSteering
		}, int8(0)},

		{"UnfilteredClutch", func(p Packet) interface{} {
			return p.UnfilteredClutch
		}, uint8(251)},

		{"RaceStateFlags", func(p Packet) interface{} {
			return p.RaceStateFlags
		}, uint8(10)},

		{"LapsInEvent", func(p Packet) interface{} {
			return p.LapsInEvent
		}, uint8(2)},

		{"BestLapTime", func(p Packet) interface{} {
			return p.BestLapTime
		}, float32(-1)},

		{"LastLapTime", func(p Packet) interface{} {
			return p.LastLapTime
		}, float32(-1)},

		{"CurrentTime", func(p Packet) interface{} {
			return p.CurrentTime
		}, float32(73.679)},

		{"SplitTimeAhead", func(p Packet) interface{} {
			return p.SplitTimeAhead
		}, float32(43.541267)},

		{"SplitTimeBehind", func(p Packet) interface{} {
			return p.SplitTimeBehind
		}, float32(-1)},

		{"SplitTime", func(p Packet) interface{} {
			return p.SplitTime
		}, float32(0)},

		{"EventTimeRemaining", func(p Packet) interface{} {
			return p.EventTimeRemaining
		}, float32(-1)},

		{"PersonalFastestLapTime", func(p Packet) interface{} {
			return p.PersonalFastestLapTime
		}, float32(-1)},

		{"WorldFastestLapTime", func(p Packet) interface{} {
			return p.WorldFastestLapTime
		}, float32(84.614006)},

		{"CurrentSector1Time", func(p Packet) interface{} {
			return p.CurrentSector1Time
		}, float32(-1)},

		{"CurrentSector2Time", func(p Packet) interface{} {
			return p.CurrentSector2Time
		}, float32(-1)},

		{"CurrentSector3Time", func(p Packet) interface{} {
			return p.CurrentSector3Time
		}, float32(-1)},

		{"FastestSector1Time", func(p Packet) interface{} {
			return p.FastestSector1Time
		}, float32(-1)},

		{"FastestSector2Time", func(p Packet) interface{} {
			return p.FastestSector2Time
		}, float32(-1)},

		{"FastestSector3Time", func(p Packet) interface{} {
			return p.FastestSector3Time
		}, float32(-1)},

		{"PersonalFastestSector1Time", func(p Packet) interface{} {
			return p.PersonalFastestSector1Time
		}, float32(-1)},

		{"PersonalFastestSector2Time", func(p Packet) interface{} {
			return p.PersonalFastestSector2Time
		}, float32(-1)},

		{"PersonalFastestSector3Time", func(p Packet) interface{} {
			return p.PersonalFastestSector3Time
		}, float32(-1)},

		{"WorldFastestSector1Time", func(p Packet) interface{} {
			return p.WorldFastestSector1Time
		}, float32(22.218)},

		{"WorldFastestSector2Time", func(p Packet) interface{} {
			return p.WorldFastestSector2Time
		}, float32(35.412003)},

		{"WorldFastestSector3Time", func(p Packet) interface{} {
			return p.WorldFastestSector3Time
		}, float32(26.984001)},

		{"JoyPad", func(p Packet) interface{} {
			return p.JoyPad
		}, uint16(0)},

		{"HighestFlag", func(p Packet) interface{} {
			return p.HighestFlag
		}, uint8(0)},

		{"PitModeSchedule", func(p Packet) interface{} {
			return p.PitModeSchedule
		}, uint8(0)},

		{"OilTempCelsius", func(p Packet) interface{} {
			return p.OilTempCelsius
		}, int16(49)},

		{"OilPressureKPa", func(p Packet) interface{} {
			return p.OilPressureKPa
		}, uint16(282)},

		{"WaterTempCelsius", func(p Packet) interface{} {
			return p.WaterTempCelsius
		}, int16(49)},

		{"WaterPressureKPa", func(p Packet) interface{} {
			return p.WaterPressureKPa
		}, uint16(158)},

		{"FuelPressureKPa", func(p Packet) interface{} {
			return p.FuelPressureKPa
		}, uint16(11)},

		{"CarFlags", func(p Packet) interface{} {
			return p.CarFlags
		}, uint8(2)},

		{"FuelCapacity", func(p Packet) interface{} {
			return p.FuelCapacity
		}, uint8(17)},

		{"Brake", func(p Packet) interface{} {
			return p.Brake
		}, uint8(0)},

		{"Throttle", func(p Packet) interface{} {
			return p.Throttle
		}, uint8(0)},

		{"Clutch", func(p Packet) interface{} {
			return p.Clutch
		}, uint8(254)},

		{"Steering", func(p Packet) interface{} {
			return p.Steering
		}, int8(0)},

		{"FuelLevel", func(p Packet) interface{} {
			return p.FuelLevel
		}, float32(0.068779856)},

		{"Speed", func(p Packet) interface{} {
			return p.Speed
		}, float32(0.0059392797)},

		{"Rpm", func(p Packet) interface{} {
			return p.Rpm
		}, uint16(5021)},

		{"MaxRpm", func(p Packet) interface{} {
			return p.MaxRpm
		}, uint16(14500)},

		{"GearNumGears", func(p Packet) interface{} {
			return p.GearNumGears
		}, uint8(97)},

		{"BoostAmount", func(p Packet) interface{} {
			return p.BoostAmount
		}, uint8(0)},

		{"EnforcedPitStopLap", func(p Packet) interface{} {
			return p.EnforcedPitStopLap
		}, int8(-1)},

		{"CrashState", func(p Packet) interface{} {
			return p.CrashState
		}, uint8(0)},

		{"OdometerKM", func(p Packet) interface{} {
			return p.OdometerKM
		}, float32(254.89038)},

		{"OrientationX", func(p Packet) interface{} {
			return p.OrientationX
		}, float32(-0.0058781323)},

		{"OrientationY", func(p Packet) interface{} {
			return p.OrientationY
		}, float32(-0.80131507)},

		{"OrientationZ", func(p Packet) interface{} {
			return p.OrientationZ
		}, float32(0.0054766205)},

		{"LocalVelocityX", func(p Packet) interface{} {
			return p.LocalVelocityX
		}, float32(6.0650007e-05)},

		{"LocalVelocity", func(p Packet) interface{} {
			return p.LocalVelocityY
		}, float32(-0.00022578704)},

		{"LocalVelocityZ", func(p Packet) interface{} {
			return p.LocalVelocityZ
		}, float32(0.00032369886)},

		{"WorldVelocityX", func(p Packet) interface{} {
			return p.WorldVelocityX
		}, float32(-0.0001903444)},

		{"WorldVelocityY", func(p Packet) interface{} {
			return p.WorldVelocityY
		}, float32(-0.00022351979)},

		{"WorldVelocityZ", func(p Packet) interface{} {
			return p.WorldVelocityZ
		}, float32(0.00027063675)},

		{"AngularVelocityX", func(p Packet) interface{} {
			return p.AngularVelocityX
		}, float32(-0.0020025268)},

		{"AngularVelocityY", func(p Packet) interface{} {
			return p.AngularVelocityY
		}, float32(0.00412653)},

		{"AngularVelocityZ", func(p Packet) interface{} {
			return p.AngularVelocityZ
		}, float32(-6.9707676e-06)},

		{"LocalAccelerationX", func(p Packet) interface{} {
			return p.LocalAccelerationX
		}, float32(-0.00020823441)},

		{"LocalAccelerationY", func(p Packet) interface{} {
			return p.LocalAccelerationY
		}, float32(0.0031503327)},

		{"LocalAccelerationZ", func(p Packet) interface{} {
			return p.LocalAccelerationZ
		}, float32(-1.2588134)},

		{"WorldAccelerationX", func(p Packet) interface{} {
			return p.WorldAccelerationX
		}, float32(0.9039378)},

		{"WorldAccelerationY", func(p Packet) interface{} {
			return p.WorldAccelerationY
		}, float32(-0.0042903796)},

		{"WorldAccelerationZ", func(p Packet) interface{} {
			return p.WorldAccelerationZ
		}, float32(-0.87606335)},

		{"ExtentsCentreX", func(p Packet) interface{} {
			return p.ExtentsCentreX
		}, float32(0)},

		{"ExtentsCentreY", func(p Packet) interface{} {
			return p.ExtentsCentreY
		}, float32(0.4531746)},

		{"ExtentsCentreZ", func(p Packet) interface{} {
			return p.ExtentsCentreZ
		}, float32(-0.112976015)},
	}

	for _, tt := range fmtTests {
		actual := tt.getter(packet)

		if actual != tt.expected {
			t.Error(fmt.Sprintf("%v: actual %T(%v) != expected %T(%v)", tt.name, actual, actual, tt.expected, tt.expected))
		}
	}

}