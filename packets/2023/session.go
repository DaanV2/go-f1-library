package f1_2023

import "github.com/DaanV2/go-f1-library/enums"

const (
	PacketSessionDataFrequency = 2
	PacketSessionDataSize      = 644
	PacketSessionDataVersion   = 1
)

type (
	MarshalZone struct {
		ZoneStart float32        // Fraction (0..1) of way through the lap the marshal zone starts
		ZoneFlag  enums.ZoneFlag // int8,
	}

	WeatherForecastSample struct {
		SessionType            enums.SessionType       // uint8 The type of session the forecast is for
		TimeOffset             uint8                   // Time in minutes the forecast is for
		WeatherType            enums.WeatherType       // uint8, Weather
		TrackTemperature       int8                    // Track temp. in degrees Celsius
		TrackTemperatureChange enums.TemperatureChange // int8, Track temp. change
		AirTemperature         int8                    // Air temp. in degrees celsius
		AirTemperatureChange   enums.TemperatureChange // int8, Air temp
		RainPercentage         uint8                   // Rain percentage (0-100)
	}

	// The session packet includes details about the current session in progress.
	PacketSessionData struct {
		Header                          PacketHeader              // Header
		Weather                         enums.WeatherType         // Weather
		TrackTemperature                int8                      // Track temp. in degrees celsius
		AirTemperature                  int8                      // Air temp. in degrees celsius
		TotalLaps                       uint8                     // Total number of laps in this race
		TrackLength                     uint16                    // Track length in metres
		SessionType                     enums.SessionType         //uint8,
		TrackId                         int8                      // -1 for unknown, see appendix
		Formula                         enums.Formula             //uint8,
		SessionTimeLeft                 uint16                    // Time left in session in seconds
		SessionDuration                 uint16                    // Session duration in seconds
		PitSpeedLimit                   uint8                     // Pit speed limit in kilometres per hour
		GamePaused                      uint8                     // Whether the game is paused – network game only
		IsSpectating                    uint8                     // Whether the player is spectating
		SpectatorCarIndex               uint8                     // Index of the car being spectated
		SliProNativeSupport             uint8                     // SLI Pro support, 0 = inactive, 1 = active
		NumMarshalZones                 uint8                     // Number of marshal zones to follow
		MarshalZones                    [21]MarshalZone           // List of marshal zones – max 21
		SafetyCarStatus                 uint8                     // 0 = no safety car, 1 = full, 2 = virtual, 3 = formation lap
		NetworkGame                     uint8                     // 0 = offline, 1 = online
		NumWeatherForecastSamples       uint8                     // Number of weather samples to follow
		WeatherForecastSamples          [56]WeatherForecastSample // Array of weather forecast samples
		ForecastAccuracy                uint8                     // 0 = Perfect, 1 = Approximate
		AiDifficulty                    uint8                     // AI Difficulty rating – 0-110
		SeasonLinkIdentifier            uint32                    // Identifier for season - persists across saves
		WeekendLinkIdentifier           uint32                    // Identifier for weekend - persists across saves
		SessionLinkIdentifier           uint32                    // Identifier for session - persists across saves
		PitStopWindowIdealLap           uint8                     // Ideal lap to pit on for current strategy (player)
		PitStopWindowLatestLap          uint8                     // Latest lap to pit on for current strategy (player)
		PitStopRejoinPosition           uint8                     // Predicted position to rejoin at (player)
		SteeringAssist                  enums.OnOff               // 0 = off, 1 = on
		BrakingAssist                   uint8                     // 0 = off, 1 = low, 2 = medium, 3 = high
		GearboxAssist                   uint8                     // 1 = manual, 2 = manual & suggested gear, 3 = auto
		PitAssist                       enums.OnOff               // 0 = off, 1 = on
		PitReleaseAssist                enums.OnOff               // 0 = off, 1 = on
		ERSAssist                       enums.OnOff               // 0 = off, 1 = on
		DRSAssist                       enums.OnOff               // 0 = off, 1 = on
		DynamicRacingLine               uint8                     // 0 = off, 1 = corners only, 2 = full
		DynamicRacingLineType           uint8                     // 0 = 2D, 1 = 3D
		GameMode                        uint8                     // Game mode id - see appendix
		RuleSet                         uint8                     // Ruleset - see appendix
		TimeOfDay                       uint32                    // Local time of day - minutes since midnight
		SessionLength                   uint8                     // 0 = None, 2 = Very Short, 3 = Short, 4 = Medium  5 = Medium Long, 6 = Long, 7 = Full
		SpeedUnitsLeadPlayer            enums.SpeedFormat         // uint8,
		TemperatureUnitsLeadPlayer      enums.TemperatureFormat   // uint8,
		SpeedUnitsSecondaryPlayer       enums.SpeedFormat         // uint8,
		TemperatureUnitsSecondaryPlayer enums.TemperatureFormat   // uint8,
		NumSafetyCarPeriods             uint8                     // Number of safety cars called during session
		NumVirtualSafetyCarPeriods      uint8                     // Number of virtual safety cars called
		NumRedFlagPeriods               uint8                     // Number of red flags called during session
	}
)
