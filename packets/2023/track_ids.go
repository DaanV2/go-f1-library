package f1_2023

type TrackId int8

const (
	Melbourne        TrackId = 0
	PaulRicard       TrackId = 1
	Shanghai         TrackId = 2
	SakhirBahrain    TrackId = 3
	Catalunya        TrackId = 4
	Monaco           TrackId = 5
	Montreal         TrackId = 6
	Silverstone      TrackId = 7
	Hockenheim       TrackId = 8
	Hungaroring      TrackId = 9
	Spa              TrackId = 10
	Monza            TrackId = 11
	Singapore        TrackId = 12
	Suzuka           TrackId = 13
	AbuDhabi         TrackId = 14
	Texas            TrackId = 15
	Brazil           TrackId = 16
	Austria          TrackId = 17
	Sochi            TrackId = 18
	Mexico           TrackId = 19
	BakuAzerbaijan   TrackId = 20
	SakhirShort      TrackId = 21
	SilverstoneShort TrackId = 22
	TexasShort       TrackId = 23
	SuzukaShort      TrackId = 24
	Hanoi            TrackId = 25
	Zandvoort        TrackId = 26
	Imola            TrackId = 27
	Portimão         TrackId = 28
	Jeddah           TrackId = 29
	Miami            TrackId = 30
	LasVegas         TrackId = 31
	Losail           TrackId = 32
)

func (t TrackId) String() string {
	switch t {
	case Melbourne:
		return "Melbourne"
	case PaulRicard:
		return "Paul Ricard"
	case Shanghai:
		return "Shanghai"
	case SakhirBahrain:
		return "Sakhir (Bahrain)"
	case Catalunya:
		return "Catalunya"
	case Monaco:
		return "Monaco"
	case Montreal:
		return "Montreal"
	case Silverstone:
		return "Silverstone"
	case Hockenheim:
		return "Hockenheim"
	case Hungaroring:
		return "Hungaroring"
	case Spa:
		return "Spa"
	case Monza:
		return "Monza"
	case Singapore:
		return "Singapore"
	case Suzuka:
		return "Suzuka"
	case AbuDhabi:
		return "Abu Dhabi"
	case Texas:
		return "Texas"
	case Brazil:
		return "Brazil"
	case Austria:
		return "Austria"
	case Sochi:
		return "Sochi"
	case Mexico:
		return "Mexico"
	case BakuAzerbaijan:
		return "Baku (Azerbaijan)"
	case SakhirShort:
		return "Sakhir Short"
	case SilverstoneShort:
		return "Silverstone Short"
	case TexasShort:
		return "Texas Short"
	case SuzukaShort:
		return "Suzuka Short"
	case Hanoi:
		return "Hanoi"
	case Zandvoort:
		return "Zandvoort"
	case Imola:
		return "Imola"
	case Portimão:
		return "Portimão"
	case Jeddah:
		return "Jeddah"
	case Miami:
		return "Miami"
	case LasVegas:
		return "Las Vegas"
	case Losail:
		return "Losail"
	}

	return "UNKNOWN"
}
