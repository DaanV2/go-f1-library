package f1_2023

import (
	"github.com/DaanV2/go-f1-library/constants"
	"github.com/DaanV2/go-f1-library/game"
)

var _ game.Infringements = &Infringements{}

type (
	InfringementType uint8
	Infringements    struct{}
)

func (d *Infringements) Get(track uint8) game.Infringement {
	return InfringementType(track)
}

func (d *Infringements) Max() uint8 {
	return IN_AttributeAssigned.Id()
}

const (
	IN_BlockingBySlowDriving                     InfringementType = 0  // Blocking by slow driving
	IN_BlockingByWrongWayDriving                 InfringementType = 1  // Blocking by wrong way driving
	IN_ReversingOffTheStartLine                  InfringementType = 2  // Reversing off the start line
	IN_BigCollision                              InfringementType = 3  // Big Collision
	IN_SmallCollision                            InfringementType = 4  // Small Collision
	IN_CollisionFailedToHandBackPositionSingle   InfringementType = 5  // Collision failed to hand back position single
	IN_CollisionFailedToHandBackPositionMultiple InfringementType = 6  // Collision failed to hand back position multiple
	IN_CornerCuttingGainedTime                   InfringementType = 7  // Corner cutting gained time
	IN_CornerCuttingOvertakeSingle               InfringementType = 8  // Corner cutting overtake single
	IN_CornerCuttingOvertakeMultiple             InfringementType = 9  // Corner cutting overtake multiple
	IN_CrossedPitExitLane                        InfringementType = 10 // Crossed pit exit lane
	IN_IgnoringBlueFlags                         InfringementType = 11 // Ignoring blue flags
	IN_IgnoringYellowFlags                       InfringementType = 12 // Ignoring yellow flags
	IN_IgnoringDriveThrough                      InfringementType = 13 // Ignoring drive through
	IN_TooManyDriveThroughs                      InfringementType = 14 // Too many drive throughs
	IN_DriveThroughReminderServeWithinNLaps      InfringementType = 15 // Drive through reminder serve within n laps
	IN_DriveThroughReminderServeThisLap          InfringementType = 16 // Drive through reminder serve this lap
	IN_PitLaneSpeeding                           InfringementType = 17 // Pit lane speeding
	IN_ParkedForTooLong                          InfringementType = 18 // Parked for too long
	IN_IgnoringTyreRegulations                   InfringementType = 19 // Ignoring tyre regulations
	IN_TooManyPenalties                          InfringementType = 20 // Too many penalties
	IN_MultipleWarnings                          InfringementType = 21 // Multiple warnings
	IN_ApproachingDisqualification               InfringementType = 22 // Approaching disqualification
	IN_TyreRegulationsSelectSingle               InfringementType = 23 // Tyre regulations select single
	IN_TyreRegulationsSelectMultiple             InfringementType = 24 // Tyre regulations select multiple
	IN_LapInvalidatedCornerCutting               InfringementType = 25 // Lap invalidated corner cutting
	IN_LapInvalidatedRunningWide                 InfringementType = 26 // Lap invalidated running wide
	IN_CornerCuttingRanWideGainedTimeMinor       InfringementType = 27 // Corner cutting ran wide gained time minor
	IN_CornerCuttingRanWideGainedTimeSignificant InfringementType = 28 // Corner cutting ran wide gained time significant
	IN_CornerCuttingRanWideGainedTimeExtreme     InfringementType = 29 // Corner cutting ran wide gained time extreme
	IN_LapInvalidatedWallRiding                  InfringementType = 30 // Lap invalidated wall riding
	IN_LapInvalidatedFlashbackUsed               InfringementType = 31 // Lap invalidated flashback used
	IN_LapInvalidatedResetToTrack                InfringementType = 32 // Lap invalidated reset to track
	IN_BlockingThePitlane                        InfringementType = 33 // Blocking the pitlane
	IN_JumpStart                                 InfringementType = 34 // Jump start
	IN_SafetyCarToCarCollision                   InfringementType = 35 // Safety car to car collision
	IN_SafetyCarIllegalOvertake                  InfringementType = 36 // Safety car illegal overtake
	IN_SafetyCarExceedingAllowedPace             InfringementType = 37 // Safety car exceeding allowed pace
	IN_VirtualSafetyCarExceedingAllowedPace      InfringementType = 38 // Virtual safety car exceeding allowed pace
	IN_FormationLapBelowAllowedSpeed             InfringementType = 39 // Formation lap below allowed speed
	IN_FormationLapParking                       InfringementType = 40 // Formation lap parking
	IN_RetiredMechanicalFailure                  InfringementType = 41 // Retired mechanical failure
	IN_RetiredTerminallyDamaged                  InfringementType = 42 // Retired terminally damaged
	IN_SafetyCarFallingTooFarBack                InfringementType = 43 // Safety car falling too far back
	IN_BlackFlagTimer                            InfringementType = 44 // Black flag timer
	IN_UnservedStopGoPenalty                     InfringementType = 45 // Unserved stop go penalty
	IN_UnservedDriveThroughPenalty               InfringementType = 46 // Unserved drive through penalty
	IN_EngineComponentChange                     InfringementType = 47 // Engine component change
	IN_GearboxChange                             InfringementType = 48 // Gearbox change
	IN_ParcFerméChange                           InfringementType = 49 // Parc Fermé change
	IN_LeagueGridPenalty                         InfringementType = 50 // League grid penalty
	IN_RetryPenalty                              InfringementType = 51 // Retry penalty
	IN_IllegalTimeGain                           InfringementType = 52 // Illegal time gain
	IN_MandatoryPitstop                          InfringementType = 53 // Mandatory pitstop
	IN_AttributeAssigned                         InfringementType = 54 // Attribute assigned
)

func (p InfringementType) Id() uint8 {
	return uint8(p)
}

func (p InfringementType) String() string {
	switch p {
	case IN_BlockingBySlowDriving:
		return "blocking by slow driving"
	case IN_BlockingByWrongWayDriving:
		return "blocking by wrong way driving"
	case IN_ReversingOffTheStartLine:
		return "reversing off the start line"
	case IN_BigCollision:
		return "big Collision"
	case IN_SmallCollision:
		return "small Collision"
	case IN_CollisionFailedToHandBackPositionSingle:
		return "collision failed to hand back position single"
	case IN_CollisionFailedToHandBackPositionMultiple:
		return "collision failed to hand back position multiple"
	case IN_CornerCuttingGainedTime:
		return "corner cutting gained time"
	case IN_CornerCuttingOvertakeSingle:
		return "corner cutting overtake single"
	case IN_CornerCuttingOvertakeMultiple:
		return "corner cutting overtake multiple"
	case IN_CrossedPitExitLane:
		return "crossed pit exit lane"
	case IN_IgnoringBlueFlags:
		return "ignoring blue flags"
	case IN_IgnoringYellowFlags:
		return "ignoring yellow flags"
	case IN_IgnoringDriveThrough:
		return "ignoring drive through"
	case IN_TooManyDriveThroughs:
		return "too many drive throughs"
	case IN_DriveThroughReminderServeWithinNLaps:
		return "drive through reminder serve within n laps"
	case IN_DriveThroughReminderServeThisLap:
		return "drive through reminder serve this lap"
	case IN_PitLaneSpeeding:
		return "pit lane speeding"
	case IN_ParkedForTooLong:
		return "parked for too long"
	case IN_IgnoringTyreRegulations:
		return "ignoring tyre regulations"
	case IN_TooManyPenalties:
		return "too many penalties"
	case IN_MultipleWarnings:
		return "multiple warnings"
	case IN_ApproachingDisqualification:
		return "approaching disqualification"
	case IN_TyreRegulationsSelectSingle:
		return "tyre regulations select single"
	case IN_TyreRegulationsSelectMultiple:
		return "tyre regulations select multiple"
	case IN_LapInvalidatedCornerCutting:
		return "lap invalidated corner cutting"
	case IN_LapInvalidatedRunningWide:
		return "lap invalidated running wide"
	case IN_CornerCuttingRanWideGainedTimeMinor:
		return "corner cutting ran wide gained time minor"
	case IN_CornerCuttingRanWideGainedTimeSignificant:
		return "corner cutting ran wide gained time significant"
	case IN_CornerCuttingRanWideGainedTimeExtreme:
		return "corner cutting ran wide gained time extreme"
	case IN_LapInvalidatedWallRiding:
		return "lap invalidated wall riding"
	case IN_LapInvalidatedFlashbackUsed:
		return "lap invalidated flashback used"
	case IN_LapInvalidatedResetToTrack:
		return "lap invalidated reset to track"
	case IN_BlockingThePitlane:
		return "blocking the pitlane"
	case IN_JumpStart:
		return "jump start"
	case IN_SafetyCarToCarCollision:
		return "safety car to car collision"
	case IN_SafetyCarIllegalOvertake:
		return "safety car illegal overtake"
	case IN_SafetyCarExceedingAllowedPace:
		return "safety car exceeding allowed pace"
	case IN_VirtualSafetyCarExceedingAllowedPace:
		return "virtual safety car exceeding allowed pace"
	case IN_FormationLapBelowAllowedSpeed:
		return "formation lap below allowed speed"
	case IN_FormationLapParking:
		return "formation lap parking"
	case IN_RetiredMechanicalFailure:
		return "retired mechanical failure"
	case IN_RetiredTerminallyDamaged:
		return "retired terminally damaged"
	case IN_SafetyCarFallingTooFarBack:
		return "safety car falling too far back"
	case IN_BlackFlagTimer:
		return "black flag timer"
	case IN_UnservedStopGoPenalty:
		return "unserved stop go penalty"
	case IN_UnservedDriveThroughPenalty:
		return "unserved drive through penalty"
	case IN_EngineComponentChange:
		return "engine component change"
	case IN_GearboxChange:
		return "gearbox change"
	case IN_ParcFerméChange:
		return "parc Fermé change"
	case IN_LeagueGridPenalty:
		return "league grid penalty"
	case IN_RetryPenalty:
		return "retry penalty"
	case IN_IllegalTimeGain:
		return "illegal time gain"
	case IN_MandatoryPitstop:
		return "mandatory pitstop"
	case IN_AttributeAssigned:
		return "attribute assigned"
	}

	return constants.UNKNOWN
}
