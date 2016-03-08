package client

import (
	"encoding/binary"
	"bytes"
)

type Packet struct {
	BuildVersionNumber         uint16
	RawSequence                uint8
	RawGameSessionState        uint8
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
	Tires                      [4]Tire
}

type Tire struct {
	TireFlags             uint8
	Terrain               uint8
	TireY                 float32
	TireRPS               float32
	TireSlipSpeed         float32
	TireTemp              uint8
	TireGrip              uint8
	TireHeightAboveGround float32
	TireLateralStiffness  float32
	TireWear              uint8
	BrakeDamage           uint8
	SuspensionDamage      uint8
	BrakeTempCelsius      int16
	TireTreadTemp         uint16
	TireLayerTemp         uint16
	TireCarcassTemp       uint16
	TireRimTemp           uint16
	TireInternalAirTemp   uint16
	WheelLocalPositionY   float32
	RideHeight            float32
	SuspensionTravel      float32
	SuspensionVelocity    float32
	AirPressure           uint16
}

type ExtrasWeather struct {
	EngineSpeed        float32
	EngineTorque       float32
	AeroDamage         uint8
	EngineDamage       uint8
	AmbientTemperature int8
	TrackTemperature   int8
	RainDensity        uint8
	WindSpeed          int8
	WindDirectionX     int8
	WindDirectionY     int8
}

type ParticipantInfo struct {
	WorldPositionX     int16
	WorldPositionY     int16
	WorldPositionZ     int16
	CurrentLapDistance uint16
	RacePosition       uint8
	LapsCompleted      uint8
	CurrentLap         uint8
	Sector             uint8
	LastSectorTime     float32
}

type Epilogue struct {
	TrackLength float32
	Wings1      uint8
	Wings2      uint8
	DPad        uint8
	Padding     uint16
}

type GameState uint8

const (
	GameState_EXITED GameState = iota
	GameState_FRONT_END GameState = iota
	GameState_INGAME_PLAYING GameState = iota
	GameState_INGAME_PAUSED GameState = iota
)

type SessionState uint8

const (
	SessionState_INVALID SessionState = iota
	SessionState_PRACTICE SessionState = iota
	SessionState_TEST SessionState = iota
	SessionState_QUALIFY SessionState = iota
	SessionState_FORMATION_LAP SessionState = iota
	SessionState_RACE SessionState = iota
	SessionState_TIME_ATTACK SessionState = iota
)

const (
	RaceState_INVALID = iota
	RaceState_NOT_STARTED = iota
	RaceState_RACING = iota
	RaceState_FINISHED = iota
	RaceState_DISQUALIFIED = iota
	RaceState_RETIRED = iota
	RaceState_DNF = iota
)

const (
	Sector_INVALID = iota
	Sector_START = iota
	Sector_SECTOR_1 = iota
	Sector_SECTOR_2 = iota
	Sector_FINISH = iota
	Sector_STOP = iota
)

const (
	FlagColour_NONE = iota           // Not used for actual flags, only for some query functions
	FlagColour_GREEN = iota          // End of danger zone, or race started
	FlagColour_BLUE = iota           // Faster car wants to overtake the participant
	FlagColour_WHITE = iota          // Approaching a slow car
	FlagColour_YELLOW = iota         // Danger on the racing surface itself
	FlagColour_DOUBLE_YELLOW = iota  // Danger that wholly or partly blocks the racing surface
	FlagColour_BLACK = iota          // Participant disqualified
	FlagColour_CHEQUERED = iota      // Chequered flag
)

const (
	FlagReason_NONE = iota
	FlagReason_SOLO_CRASH = iota
	FlagReason_VEHICLE_CRASH = iota
	FlagReason_VEHICLE_OBSTRUCTION = iota
)

const (
	PitMode_NONE = iota
	PitMode_DRIVING_INTO_PITS = iota
	PitMode_IN_PIT = iota
	PitMode_DRIVING_OUT_OF_PITS = iota
	PitMode_IN_GARAGE = iota
)

const (
	PitSchedule_NONE = iota
	PitSchedule_STANDARD = iota
	PitSchedule_DRIVE_THROUGH = iota
	PitSchedule_STOP_GO = iota
)

const (
	CarFlags_NONE = 0
	CarFlags_HEADLIGHT = 1
	CarFlags_ENGINE_ACTIVE = 2
	CarFlags_ENGINE_WARNING = 4
	CarFlags_SPEED_LIMITER = 8
	CarFlags_ABS = 16
	CarFlags_HANDBRAKE = 32
)

const (
	Tires_FRONT_LEFT = iota
	Tires_FRONT_RIGHT = iota
	Tires_REAR_LEFT = iota
	Tires_REAR_RIGHT = iota
)

const (
	TireFlags_NONE = 0
	TireFlags_ATTACHED = 1
	TireFlags_INFLATED = 2
	TireFlags_IS_ON_GROUND = 4
)

const (
	Terrain_TERRAIN_ROAD = iota
	Terrain_TERRAIN_LOW_GRIP_ROAD = iota
	Terrain_TERRAIN_BUMPY_ROAD1 = iota
	Terrain_TERRAIN_BUMPY_ROAD2 = iota
	Terrain_TERRAIN_BUMPY_ROAD3 = iota
	Terrain_TERRAIN_MARBLES = iota
	Terrain_TERRAIN_GRASSY_BERMS = iota
	Terrain_TERRAIN_GRASS = iota
	Terrain_TERRAIN_GRAVEL = iota
	Terrain_TERRAIN_BUMPY_GRAVEL = iota
	Terrain_TERRAIN_RUMBLE_STRIPS = iota
	Terrain_TERRAIN_DRAINS = iota
	Terrain_TERRAIN_TireWALLS = iota
	Terrain_TERRAIN_CEMENTWALLS = iota
	Terrain_TERRAIN_GUARDRAILS = iota
	Terrain_TERRAIN_SAND = iota
	Terrain_TERRAIN_BUMPY_SAND = iota
	Terrain_TERRAIN_DIRT = iota
	Terrain_TERRAIN_BUMPY_DIRT = iota
	Terrain_TERRAIN_DIRT_ROAD = iota
	Terrain_TERRAIN_BUMPY_DIRT_ROAD = iota
	Terrain_TERRAIN_PAVEMENT = iota
	Terrain_TERRAIN_DIRT_BANK = iota
	Terrain_TERRAIN_WOOD = iota
	Terrain_TERRAIN_DRY_VERGE = iota
	Terrain_TERRAIN_EXIT_RUMBLE_STRIPS = iota
	Terrain_TERRAIN_GRASSCRETE = iota
	Terrain_TERRAIN_LONG_GRASS = iota
	Terrain_TERRAIN_SLOPE_GRASS = iota
	Terrain_TERRAIN_COBBLES = iota
	Terrain_TERRAIN_SAND_ROAD = iota
	Terrain_TERRAIN_BAKED_CLAY = iota
	Terrain_TERRAIN_ASTROTURF = iota
	Terrain_TERRAIN_SNOWHALF = iota
	Terrain_TERRAIN_SNOWFULL = iota
)

func (p Packet) GetSequenceNumber() uint8 {
	return (p.RawSequence & 0xFC) >> 2
}

func (p Packet) GetPacketType() uint8 {
	return p.RawSequence & 0x3
}

func (p Packet) GetGameState() GameState {
	return GameState((p.RawGameSessionState) & 0x07)
}

func (p Packet) GetSessionState() SessionState {
	return SessionState((p.RawGameSessionState & 0x38) >> 2)
}

func Parse(raw_message []byte) (Packet, error) {
	packet := Packet{}
	buf := bytes.NewReader(raw_message)
	err := binary.Read(buf, binary.LittleEndian, &packet)
	return packet, err
}