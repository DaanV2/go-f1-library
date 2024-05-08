package f1_2023

import (
	"github.com/DaanV2/go-f1-library/encoding"
	"github.com/DaanV2/go-f1-library/enums"
)

const (
	PacketSessionDataFrequency = 2 // Frequency: 2 per second
	PacketSessionDataSize      = 644
	PacketSessionDataVersion   = 1
)

type (
	// The session packet includes details about the current session in progress.
	PacketSessionData struct {
		Header                          PacketHeader              `json:"header"`                             // Header
		Weather                         enums.WeatherType         `json:"weather"`                            // uint8, Weather
		TrackTemperature                int8                      `json:"track_temperature"`                  // Track temp. in degrees celsius
		AirTemperature                  int8                      `json:"air_temperature"`                    // Air temp. in degrees celsius
		TotalLaps                       uint8                     `json:"total_laps"`                         // Total number of laps in this race
		TrackLength                     uint16                    `json:"track_length"`                       // Track length in metres
		SessionType                     enums.SessionType         `json:"session_type"`                       // uint8,
		TrackId                         int8                      `json:"track_id"`                           // -1 for unknown, see appendix
		Formula                         enums.Formula             `json:"formula"`                            // uint8,
		SessionTimeLeft                 uint16                    `json:"session_time_left"`                  // Time left in session in seconds
		SessionDuration                 uint16                    `json:"session_duration"`                   // Session duration in seconds
		PitSpeedLimit                   uint8                     `json:"pit_speed_limit"`                    // Pit speed limit in kilometres per hour
		GamePaused                      uint8                     `json:"game_paused"`                        // Whether the game is paused – network game only
		IsSpectating                    uint8                     `json:"is_spectating"`                      // Whether the player is spectating
		SpectatorCarIndex               uint8                     `json:"spectator_car_index"`                // Index of the car being spectated
		SliProNativeSupport             uint8                     `json:"sli_pro_native_support"`             // SLI Pro support, 0 = inactive, 1 = active
		NumMarshalZones                 uint8                     `json:"num_marshal_zones"`                  // Number of marshal zones to follow
		MarshalZones                    [21]MarshalZone           `json:"marshal_zones"`                      // List of marshal zones – max 21
		SafetyCarStatus                 uint8                     `json:"safety_car_status"`                  // 0 = no safety car, 1 = full, 2 = virtual, 3 = formation lap
		NetworkGame                     uint8                     `json:"network_game"`                       // 0 = offline, 1 = online
		NumWeatherForecastSamples       uint8                     `json:"num_weather_forecast_samples"`       // Number of weather samples to follow
		WeatherForecastSamples          [56]WeatherForecastSample `json:"weather_forecast_samples"`           // Array of weather forecast samples
		ForecastAccuracy                uint8                     `json:"forecast_accuracy"`                  // 0 = Perfect, 1 = Approximate
		AiDifficulty                    uint8                     `json:"ai_difficulty"`                      // AI Difficulty rating – 0-110
		SeasonLinkIdentifier            uint32                    `json:"season_link_identifier"`             // Identifier for season - persists across saves
		WeekendLinkIdentifier           uint32                    `json:"weekend_link_identifier"`            // Identifier for weekend - persists across saves
		SessionLinkIdentifier           uint32                    `json:"session_link_identifier"`            // Identifier for session - persists across saves
		PitStopWindowIdealLap           uint8                     `json:"pit_stop_window_ideal_lap"`          // Ideal lap to pit on for current strategy (player)
		PitStopWindowLatestLap          uint8                     `json:"pit_stop_window_latest_lap"`         // Latest lap to pit on for current strategy (player)
		PitStopRejoinPosition           uint8                     `json:"pit_stop_rejoin_position"`           // Predicted position to rejoin at (player)
		SteeringAssist                  uint8                     `json:"steering_assist"`                    // 0 = off, 1 = on
		BrakingAssist                   uint8                     `json:"braking_assist"`                     // 0 = off, 1 = low, 2 = medium, 3 = high
		GearboxAssist                   uint8                     `json:"gearbox_assist"`                     // 1 = manual, 2 = manual & suggested gear, 3 = auto
		PitAssist                       uint8                     `json:"pit_assist"`                         // 0 = off, 1 = on
		PitReleaseAssist                uint8                     `json:"pit_release_assist"`                 // 0 = off, 1 = on
		ERSAssist                       uint8                     `json:"ers_assist"`                         // 0 = off, 1 = on
		DRSAssist                       uint8                     `json:"drs_assist"`                         // 0 = off, 1 = on
		DynamicRacingLine               uint8                     `json:"dynamic_racing_line"`                // 0 = off, 1 = corners only, 2 = full
		DynamicRacingLineType           uint8                     `json:"dynamic_racing_line_type"`           // 0 = 2D, 1 = 3D
		GameMode                        uint8                     `json:"game_mode"`                          // Game mode id - see appendix
		RuleSet                         uint8                     `json:"rule_set"`                           // Ruleset - see appendix
		TimeOfDay                       uint32                    `json:"time_of_day"`                        // Local time of day - minutes since midnight
		SessionLength                   uint8                     `json:"session_length"`                     // 0 = None, 2 = Very Short, 3 = Short, 4 = Medium  5 = Medium Long, 6 = Long, 7 = Full
		SpeedUnitsLeadPlayer            enums.SpeedFormat         `json:"speed_units_lead_player"`            // uint8,
		TemperatureUnitsLeadPlayer      enums.TemperatureFormat   `json:"temperature_units_lead_player"`      // uint8,
		SpeedUnitsSecondaryPlayer       enums.SpeedFormat         `json:"speed_units_secondary_player"`       // uint8,
		TemperatureUnitsSecondaryPlayer enums.TemperatureFormat   `json:"temperature_units_secondary_player"` // uint8,
		NumSafetyCarPeriods             uint8                     `json:"num_safety_car_periods"`             // Number of safety cars called during session
		NumVirtualSafetyCarPeriods      uint8                     `json:"num_virtual_safety_car_periods"`     // Number of virtual safety cars called
		NumRedFlagPeriods               uint8                     `json:"num_red_flag_periods"`               // Number of red flags called during session
	}

	MarshalZone struct {
		ZoneStart float32        `json:"zone_start"` // Fraction (0..1) of way through the lap the marshal zone starts
		ZoneFlag  enums.ZoneFlag `json:"zone_flag"`  // int8,
	}

	WeatherForecastSample struct {
		SessionType            enums.SessionType       `json:"session_type"`             // uint8 The type of session the forecast is for
		TimeOffset             uint8                   `json:"time_offset"`              // Time in minutes the forecast is for
		WeatherType            enums.WeatherType       `json:"weather_type"`             // uint8, Weather
		TrackTemperature       int8                    `json:"track_temperature"`        // Track temp. in degrees Celsius
		TrackTemperatureChange enums.TemperatureChange `json:"track_temperature_change"` // int8, Track temp. change
		AirTemperature         int8                    `json:"air_temperature"`          // Air temp. in degrees celsius
		AirTemperatureChange   enums.TemperatureChange `json:"air_temperature_change"`   // int8, Air temp
		RainPercentage         uint8                   `json:"rain_percentage"`          // Rain percentage (0-100)
	}
)

// ParsePacketSessionData will parse the given data into a packet
func ParsePacketSessionData(decoder *encoding.Decoder) (PacketSessionData, error) {
	header, err := ParsePacketHeader(decoder)
	if err != nil {
		return PacketSessionData{}, err
	}

	return ParsePacketSessionDataWithHeader(decoder, header)
}

// ParsePacketSessionDataWithHeader will parse the given data into a packet, expected the decoder is past the header
func ParsePacketSessionDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketSessionData, error) {
	if decoder.LeftToRead() < PacketSessionDataSize {
		return PacketSessionData{}, encoding.ErrBufferNotLargeEnough
	}

	return PacketSessionData{
		Header:                          header,
		Weather:                         enums.WeatherType(decoder.Uint8()),
		TrackTemperature:                decoder.Int8(),
		AirTemperature:                  decoder.Int8(),
		TotalLaps:                       decoder.Uint8(),
		TrackLength:                     decoder.Uint16(),
		SessionType:                     enums.SessionType(decoder.Uint8()),
		TrackId:                         decoder.Int8(),
		Formula:                         enums.Formula(decoder.Uint8()),
		SessionTimeLeft:                 decoder.Uint16(),
		SessionDuration:                 decoder.Uint16(),
		PitSpeedLimit:                   decoder.Uint8(),
		GamePaused:                      decoder.Uint8(),
		IsSpectating:                    decoder.Uint8(),
		SpectatorCarIndex:               decoder.Uint8(),
		SliProNativeSupport:             decoder.Uint8(),
		NumMarshalZones:                 decoder.Uint8(),
		MarshalZones:                    parseMarshalZone(decoder),
		SafetyCarStatus:                 decoder.Uint8(),
		NetworkGame:                     decoder.Uint8(),
		NumWeatherForecastSamples:       decoder.Uint8(),
		WeatherForecastSamples:          parseWeatherForecastSample(decoder),
		ForecastAccuracy:                decoder.Uint8(),
		AiDifficulty:                    decoder.Uint8(),
		SeasonLinkIdentifier:            decoder.Uint32(),
		WeekendLinkIdentifier:           decoder.Uint32(),
		SessionLinkIdentifier:           decoder.Uint32(),
		PitStopWindowIdealLap:           decoder.Uint8(),
		PitStopWindowLatestLap:          decoder.Uint8(),
		PitStopRejoinPosition:           decoder.Uint8(),
		SteeringAssist:                  decoder.Uint8(),
		BrakingAssist:                   decoder.Uint8(),
		GearboxAssist:                   decoder.Uint8(),
		PitAssist:                       decoder.Uint8(),
		PitReleaseAssist:                decoder.Uint8(),
		ERSAssist:                       decoder.Uint8(),
		DRSAssist:                       decoder.Uint8(),
		DynamicRacingLine:               decoder.Uint8(),
		DynamicRacingLineType:           decoder.Uint8(),
		GameMode:                        decoder.Uint8(),
		RuleSet:                         decoder.Uint8(),
		TimeOfDay:                       decoder.Uint32(),
		SessionLength:                   decoder.Uint8(),
		SpeedUnitsLeadPlayer:            enums.SpeedFormat(decoder.Uint8()),
		TemperatureUnitsLeadPlayer:      enums.TemperatureFormat(decoder.Uint8()),
		SpeedUnitsSecondaryPlayer:       enums.SpeedFormat(decoder.Uint8()),
		TemperatureUnitsSecondaryPlayer: enums.TemperatureFormat(decoder.Uint8()),
		NumSafetyCarPeriods:             decoder.Uint8(),
		NumVirtualSafetyCarPeriods:      decoder.Uint8(),
		NumRedFlagPeriods:               decoder.Uint8(),
	}, nil
}

func parseMarshalZone(decoder *encoding.Decoder) [21]MarshalZone {
	items := [21]MarshalZone{}

	for i := range items {
		items[i] = MarshalZone{
			ZoneStart: decoder.Float32(),
			ZoneFlag:  enums.ZoneFlag(decoder.Int8()),
		}
	}

	return items
}

func parseWeatherForecastSample(decoder *encoding.Decoder) [56]WeatherForecastSample {
	items := [56]WeatherForecastSample{}

	for i := range items {
		items[i] = WeatherForecastSample{
			SessionType:            enums.SessionType(decoder.Uint8()),
			TimeOffset:             decoder.Uint8(),
			WeatherType:            enums.WeatherType(decoder.Uint8()),
			TrackTemperature:       decoder.Int8(),
			TrackTemperatureChange: enums.TemperatureChange(decoder.Int8()),
			AirTemperature:         decoder.Int8(),
			AirTemperatureChange:   enums.TemperatureChange(decoder.Int8()),
			RainPercentage:         decoder.Uint8(),
		}
	}

	return items
}
