package enums

type SessionType uint8

const (
	Unknown     SessionType = 0
	P1          SessionType = 1
	P2          SessionType = 2
	P3          SessionType = 3
	ShortP      SessionType = 4
	Q1          SessionType = 5
	Q2          SessionType = 6
	Q3          SessionType = 7
	ShortQ      SessionType = 8
	OSQ         SessionType = 9
	R           SessionType = 10
	R2          SessionType = 11
	R3          SessionType = 12
	TimeTrial_s SessionType = 13
)
