package enums

// RuleSet represents the rule set.
type RuleSet int8

const (
	RS_PracticeQualifying  RuleSet = 0 // Practice/Qualifying
	RS_Race                RuleSet = 1 // Race
	RS_TimeTrial           RuleSet = 2 // Time trial
	RS_TimeAttack          RuleSet = 4 // Time attack
	RS_CheckpointChallenge RuleSet = 6 // Checkpoint challenge
	RS_Autocross           RuleSet = 8 // Autocross
	RS_Drift               RuleSet = 9 // Drift
	RS_AverageSpeedZone    RuleSet = 10 // Average speed zone
	RS_RivalDuel           RuleSet = 11 // Rival duel
)
