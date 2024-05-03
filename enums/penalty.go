package enums

type PenaltyType int8

const (
	DriveThrough                               PenaltyType = 0
	StopGo                                     PenaltyType = 1
	GridPenalty                                PenaltyType = 2
	PenaltyReminder                            PenaltyType = 3
	TimePenalty                                PenaltyType = 4
	Warning                                    PenaltyType = 5
	Disqualified                               PenaltyType = 6
	RemovedFromFormationLap                    PenaltyType = 7
	ParkedTooLongTimer                         PenaltyType = 8
	TyreRegulations                            PenaltyType = 9
	ThisLapInvalidated                         PenaltyType = 10
	ThisAndNextLapInvalidated                  PenaltyType = 11
	ThisLapInvalidatedWithoutReason            PenaltyType = 12
	ThisAndNextLapInvalidatedWithoutReason     PenaltyType = 13
	ThisAndPreviousLapInvalidated              PenaltyType = 14
	ThisAndPreviousLapInvalidatedWithoutReason PenaltyType = 15
	Retired                                    PenaltyType = 16
	BlackFlagTimer                             PenaltyType = 17
)
