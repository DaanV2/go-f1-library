package f1_2023

import "github.com/DaanV2/go-f1-library/encoding"

const (
	PacketLapDataSize    = 1131
	PacketLapDataVersion = 1
)

type (
	// The lap data packet gives details of all the cars in the session.
	PacketLapData struct {
		Header               PacketHeader `json:"header"`                // Header
		LapData              [22]LapData  `json:"lap_data"`              // Lap data for all cars on track
		TimeTrialPBCarIdx    uint8        `json:"time_trial_pb_car_idx"`    // Index of Personal Best car in time trial (255 if invalid)
		TimeTrialRivalCarIdx uint8        `json:"time_trial_rival_car_idx"` // Index of Rival car in time trial (255 if invalid)
	}

	LapData struct {
		LastLapTimeInMS             uint32  `json:"last_lap_time_in_ms"`              // Last lap time in milliseconds
		CurrentLapTimeInMS          uint32  `json:"current_lap_time_in_ms"`           // Current time around the lap in milliseconds
		Sector1TimeInMS             uint16  `json:"sector_1_time_in_ms"`              // Sector 1 time in milliseconds
		Sector1TimeMinutes          uint8   `json:"sector_1_time_minutes"`            // Sector 1 whole minute part
		Sector2TimeInMS             uint16  `json:"sector_2_time_in_ms"`              // Sector 2 time in milliseconds
		Sector2TimeMinutes          uint8   `json:"sector_2_time_minutes"`            // Sector 2 whole minute part
		DeltaToCarInFrontInMS       uint16  `json:"delta_to_car_in_front_in_ms"`      // Time delta to car in front in milliseconds
		DeltaToRaceLeaderInMS       uint16  `json:"delta_to_race_leader_in_ms"`       // Time delta to race leader in milliseconds
		LapDistance                 float32 `json:"lap_distance"`                     // Distance vehicle is around current lap in metres – could,  be negative if line hasn’t been crossed yet
		TotalDistance               float32 `json:"total_distance"`                   // Total distance travelled in session in metres – could, be negative if line hasn’t been crossed yet
		SafetyCarDelta              float32 `json:"safety_car_delta"`                 // Delta in seconds for safety car
		CarPosition                 uint8   `json:"car_position"`                     // Car race position
		CurrentLapNum               uint8   `json:"current_lap_num"`                  // Current lap number
		PitStatus                   uint8   `json:"pit_status"`                       // 0 = none, 1 = pitting, 2 = in pit area
		NumPitStops                 uint8   `json:"num_pit_stops"`                    // Number of pit stops taken in this race
		Sector                      uint8   `json:"sector"`                           // 0 = sector1, 1 = sector2, 2 = sector3
		CurrentLapInvalid           uint8   `json:"current_lap_invalid"`              // Current lap invalid - 0 = valid, 1 = invalid
		Penalties                   uint8   `json:"penalties"`                        // Accumulated time penalties in seconds to be added
		TotalWarnings               uint8   `json:"total_warnings"`                   // Accumulated number of warnings issued
		CornerCuttingWarnings       uint8   `json:"corner_cutting_warnings"`           // Accumulated number of corner cutting warnings issued
		NumUnservedDriveThroughPens uint8   `json:"num_unserved_drive_through_pens"`   // Num drive through pens left to serve
		NumUnservedStopGoPens       uint8   `json:"num_unserved_stop_go_pens"`         // Num stop go pens left to serve
		GridPosition                uint8   `json:"grid_position"`                    // Grid position the vehicle started the race in
		DriverStatus                uint8   `json:"driver_status"`                    // Status of driver - 0 = in garage, 1 = flying lap, 2 = in lap, 3 = out lap, 4 = on track
		ResultStatus                uint8   `json:"result_status"`                    // Result status - 0 = invalid, 1 = inactive, 2 = active, 3 = finished, 4 = didnotfinish, 5 = disqualified, 6 = not classified, 7 = retired
		PitLaneTimerActive          uint8   `json:"pit_lane_timer_active"`            // Pit lane timing, 0 = inactive, 1 = active
		PitLaneTimeInLaneInMS       uint16  `json:"pit_lane_time_in_lane_in_ms"`       // If active, the current time spent in the pit lane in ms
		PitStopTimerInMS            uint16  `json:"pit_stop_timer_in_ms"`             // Time of the actual pit stop in ms
		PitStopShouldServePen       uint8   `json:"pit_stop_should_serve_pen"`        // Whether the car should serve a penalty at this stop
	}
)

// ParsePacketLapData will parse the given data into a packet
func ParsePacketLapData(decoder *encoding.Decoder) (PacketLapData, error) {
	header, err := ParsePacketHeader(decoder)
	if err != nil {
		return PacketLapData{}, err
	}

	return ParsePacketLapDataWithHeader(decoder, header)
}

// ParsePacketLapDataWithHeader will parse the given data into a packet, expected the decoder is past the header
func ParsePacketLapDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketLapData, error) {
	if decoder.LeftToRead() < PacketLapDataSize {
		return PacketLapData{}, encoding.ErrBufferNotLargeEnough
	}

	packet := PacketLapData{
		Header:               header,
		LapData:              parseLapData(decoder),
		TimeTrialPBCarIdx:    decoder.Uint8(),
		TimeTrialRivalCarIdx: decoder.Uint8(),
	}

	return packet, nil
}

func parseLapData(decoder *encoding.Decoder) [22]LapData {
	items := [22]LapData{}

	for i := range items {
		items[i] = LapData{
			LastLapTimeInMS:             decoder.Uint32(),
			CurrentLapTimeInMS:          decoder.Uint32(),
			Sector1TimeInMS:             decoder.Uint16(),
			Sector1TimeMinutes:          decoder.Uint8(),
			Sector2TimeInMS:             decoder.Uint16(),
			Sector2TimeMinutes:          decoder.Uint8(),
			DeltaToCarInFrontInMS:       decoder.Uint16(),
			DeltaToRaceLeaderInMS:       decoder.Uint16(),
			LapDistance:                 decoder.Float32(),
			TotalDistance:               decoder.Float32(),
			SafetyCarDelta:              decoder.Float32(),
			CarPosition:                 decoder.Uint8(),
			CurrentLapNum:               decoder.Uint8(),
			PitStatus:                   decoder.Uint8(),
			NumPitStops:                 decoder.Uint8(),
			Sector:                      decoder.Uint8(),
			CurrentLapInvalid:           decoder.Uint8(),
			Penalties:                   decoder.Uint8(),
			TotalWarnings:               decoder.Uint8(),
			CornerCuttingWarnings:       decoder.Uint8(),
			NumUnservedDriveThroughPens: decoder.Uint8(),
			NumUnservedStopGoPens:       decoder.Uint8(),
			GridPosition:                decoder.Uint8(),
			DriverStatus:                decoder.Uint8(),
			ResultStatus:                decoder.Uint8(),
			PitLaneTimerActive:          decoder.Uint8(),
			PitLaneTimeInLaneInMS:       decoder.Uint16(),
			PitStopTimerInMS:            decoder.Uint16(),
			PitStopShouldServePen:       decoder.Uint8(),
		}
	}

	return items
}
