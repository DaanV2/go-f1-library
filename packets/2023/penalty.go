package f1_2023

import (
	"github.com/DaanV2/go-f1-library/constants"
	"github.com/DaanV2/go-f1-library/game"
)

var _ game.Penalties = &Penalties{}

type (
	PenaltyType uint8
	Penalties   struct{}
)

func (d *Penalties) Get(track uint8) game.Penalty {
	return PenaltyType(track)
}

func (d *Penalties) Max() uint8 {
	return P_BlackFlagTimer.Id()
}

const (
	P_DriveThrough                               PenaltyType = 0
	P_StopGo                                     PenaltyType = 1
	P_GridPenalty                                PenaltyType = 2
	P_PenaltyReminder                            PenaltyType = 3
	P_TimePenalty                                PenaltyType = 4
	P_Warning                                    PenaltyType = 5
	P_Disqualified                               PenaltyType = 6
	P_RemovedFromFormationLap                    PenaltyType = 7
	P_ParkedTooLongTimer                         PenaltyType = 8
	P_TyreRegulations                            PenaltyType = 9
	P_ThisLapInvalidated                         PenaltyType = 10
	P_ThisAndNextLapInvalidated                  PenaltyType = 11
	P_ThisLapInvalidatedWithoutReason            PenaltyType = 12
	P_ThisAndNextLapInvalidatedWithoutReason     PenaltyType = 13
	P_ThisAndPreviousLapInvalidated              PenaltyType = 14
	P_ThisAndPreviousLapInvalidatedWithoutReason PenaltyType = 15
	P_Retired                                    PenaltyType = 16
	P_BlackFlagTimer                             PenaltyType = 17
)

func (p PenaltyType) Id() uint8 {
	return uint8(p)
}

func (p PenaltyType) String() string {
	switch p {
	case P_DriveThrough:
		return "drive through"
	case P_StopGo:
		return "stop go"
	case P_GridPenalty:
		return "grid penalty"
	case P_PenaltyReminder:
		return "penalty reminder"
	case P_TimePenalty:
		return "time penalty"
	case P_Warning:
		return "warning"
	case P_Disqualified:
		return "disqualified"
	case P_RemovedFromFormationLap:
		return "removed from formation lap"
	case P_ParkedTooLongTimer:
		return "parked too long timer"
	case P_TyreRegulations:
		return "tyre regulations"
	case P_ThisLapInvalidated:
		return "this lap invalidated"
	case P_ThisAndNextLapInvalidated:
		return "this and next lap invalidated"
	case P_ThisLapInvalidatedWithoutReason:
		return "this lap invalidated without reason"
	case P_ThisAndNextLapInvalidatedWithoutReason:
		return "this and next lap invalidated without reason"
	case P_ThisAndPreviousLapInvalidated:
		return "this and previous lap invalidated"
	case P_ThisAndPreviousLapInvalidatedWithoutReason:
		return "this and previous lap invalidated without reason"
	case P_Retired:
		return "retired"
	case P_BlackFlagTimer:
		return "black flag timer"
	}

	return constants.UNKNOWN
}
