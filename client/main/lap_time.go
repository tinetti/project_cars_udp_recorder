package main

type LapTime struct {
    CarName         string
    CarClassName    string
    TrackLocation   string
    TrackVariation  string
    Names           []string
    FastestLapTimes []float32
}

func CreateLapTime(packet Packet) LapTime {
    lapTime := LapTime{}
    lapTime.CarName = packet.ParticipantInfo.GetCarName()
    lapTime.CarClassName = packet.ParticipantInfo.GetCarClassName()
    lapTime.TrackLocation = packet.ParticipantInfo.GetTrackLocation()
    lapTime.TrackVariation = packet.ParticipantInfo.GetTrackVariation()
    lapTime.Names = packet.ParticipantInfo.GetNames()
    lapTime.FastestLapTimes = packet.ParticipantInfo.GetFastestLapTimes()

    return lapTime
}
