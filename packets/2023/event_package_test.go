package f1_2023_test

import (
	"encoding/json"
	"testing"

	"github.com/DaanV2/go-f1-library/enums"
	f1_2023 "github.com/DaanV2/go-f1-library/packets/2023"
	"github.com/stretchr/testify/require"
)

func Test_PacketEventData_Json_Marshal(t *testing.T) {
	cases := []f1_2023.PacketEventData{
		{
			EventStringCode: f1_2023.EC_FastestLap,
			EventDetails: f1_2023.FastestLap{
				VehicleIdx: 1,
				LapTime:    1.2,
			},
		},
		{
			EventStringCode: f1_2023.EC_Retirement,
			EventDetails: f1_2023.Retirement{
				VehicleIdx: 1,
			},
		},
		{
			EventStringCode: f1_2023.EC_TeamMateInPits,
			EventDetails: f1_2023.TeamMateInPits{
				VehicleIdx: 1,
			},
		},
		{
			EventStringCode: f1_2023.EC_RaceWinner,
			EventDetails: f1_2023.RaceWinner{
				VehicleIdx: 1,
			},
		},
		{
			EventStringCode: f1_2023.EC_PenaltyIssued,
			EventDetails: f1_2023.Penalty{
				PenaltyType:      f1_2023.P_BlackFlagTimer,
				InfringementType: f1_2023.IN_BigCollision,
				VehicleIdx:       1,
				OtherVehicleIdx:  2,
				Time:             5,
				LapNum:           1,
				PlacesGained:     1,
			},
		},
		{
			EventStringCode: f1_2023.EC_SpeedTrapTriggered,
			EventDetails: f1_2023.SpeedTrap{
				VehicleIdx:                 1,
				Speed:                      200,
				IsOverallFastestInSession:  1,
				IsDriverFastestInSession:   1,
				FastestVehicleIdxInSession: 1,
				FastestSpeedInSession:      1,
			},
		},
		{
			EventStringCode: f1_2023.EC_StartLights,
			EventDetails: f1_2023.StartLights{
				NumLights: 5,
			},
		},
		{
			EventStringCode: f1_2023.EC_DriveThroughServed,
			EventDetails: f1_2023.DriveThroughPenaltyServed{
				VehicleIdx: 1,
			},
		},
		{
			EventStringCode: f1_2023.EC_StopGoServed,
			EventDetails: f1_2023.StopGoPenaltyServed{
				VehicleIdx: 1,
			},
		},
		{
			EventStringCode: f1_2023.EC_Flashback,
			EventDetails: f1_2023.Flashback{
				FlashbackFrameIdentifier: 1,
				FlashbackSessionTime:     1.2,
			},
		},
		{
			EventStringCode: f1_2023.EC_ButtonStatus,
			EventDetails: f1_2023.Buttons{
				ButtonStatus: enums.BUT_DPadLeft,
			},
		},
		{
			EventStringCode: f1_2023.EC_Overtake,
			EventDetails: f1_2023.Overtake{
				OvertakingVehicleIdx:     1,
				BeingOvertakenVehicleIdx: 2,
			},
		},
	}

	for _, c := range cases {
		t.Run(string(c.EventStringCode), func(t *testing.T) {
			// Marshal to json and back, should be the same

			data, err := json.Marshal(c)
			require.NoError(t, err)

			var c2 f1_2023.PacketEventData
			err = json.Unmarshal(data, &c2)
			require.NoError(t, err)

			require.Equal(t, c, c2)

			require.Equal(t, c.Header, c2.Header)
			require.Equal(t, c.EventStringCode, c2.EventStringCode)
			require.Equal(t, c.EventDetails, c2.EventDetails)
		})
	}
}
