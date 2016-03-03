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
		getter   func(Packet) interface{}
		expected interface{}
	}{
		{func(p Packet) interface{} {
			return p.BuildVersionNumber
		}, uint16(1182)},

		{func(p Packet) interface{} {
			return p.GetGameState()
		}, GameState_INGAME_PLAYING},

		{func(p Packet) interface{} {
			return p.GetSessionState()
		}, SessionState_FORMATION_LAP},

		{func(p Packet) interface{} {
			return p.ViewedParticipantIndex
		}, int8(0)},

		{func(p Packet) interface{} {
			return p.NumParticipants
		}, int8(16)},
	}
	
	for _, tt := range fmtTests {
		actual := tt.getter(packet)

		if actual != tt.expected {
			t.Error(fmt.Sprintf("actual %T(%v) != expected %T(%v)", actual, actual, tt.expected, tt.expected))
		}
	}

}