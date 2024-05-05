package f1_2023

import "github.com/DaanV2/go-f1-library/constants"

type ResultStatus uint8

const (
	Invalid       ResultStatus = 0 // Invalid
	Inactive      ResultStatus = 1 // Inactive
	Active        ResultStatus = 2 // Active
	Finished      ResultStatus = 3 // Finished
	DidNotFinish  ResultStatus = 4 // Did not finish
	Disqualified  ResultStatus = 5 // Disqualified
	NotClassified ResultStatus = 6 // Not classified
	Retired       ResultStatus = 7 // Retired
)

func (r ResultStatus) String() string {
	switch r {
	case Invalid:
		return "Invalid"
	case Inactive:
		return "Inactive"
	case Active:
		return "Active"
	case Finished:
		return "Finished"
	case DidNotFinish:
		return "Did not finish"
	case Disqualified:
		return "Disqualified"
	case NotClassified:
		return "Not classified"
	case Retired:
		return "Retired"
	}

	return constants.UNKNOWN
}

func (r ResultStatus) Id() uint8 {
	return uint8(r)
}