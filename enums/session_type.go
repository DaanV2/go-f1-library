package enums

type SessionType uint8

const (
	SE_Unknown   SessionType = 0  // Unknown
	SE_P1        SessionType = 1  // Practice 1
	SE_P2        SessionType = 2  // Practice 2
	SE_P3        SessionType = 3  // Practice 3
	SE_ShortP    SessionType = 4  // Short Practice
	SE_Q1        SessionType = 5  // Qualifying 1
	SE_Q2        SessionType = 6  // Qualifying 2
	SE_Q3        SessionType = 7  // Qualifying 3
	SE_ShortQ    SessionType = 8  // Short Qualifying
	SE_OSQ       SessionType = 9  // OSQ
	SE_R         SessionType = 10 // Race
	SE_R2        SessionType = 11 // Race 2
	SE_R3        SessionType = 12 // Race 3
	SE_TimeTrial SessionType = 13 // Time trails
)
