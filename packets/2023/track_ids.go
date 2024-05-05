package f1_2023

import (
	"github.com/DaanV2/go-f1-library/constants"
	"github.com/DaanV2/go-f1-library/game"
)

var _ game.Tracks = &Tracks{}

type (
	TrackId uint8
	Tracks struct{}
)

func (d *Tracks) Get(track uint8) game.Track {
	return TrackId(track)
}

func (d *Tracks) Max() uint8 {
	return Track_Max.Id()
}

const (
	Melbourne        TrackId = 0  // Melbourne
	PaulRicard       TrackId = 1  // Paul Ricard
	Shanghai         TrackId = 2  // Shanghai
	SakhirBahrain    TrackId = 3  // Sakhir (Bahrain)
	Catalunya        TrackId = 4  // Catalunya
	Monaco           TrackId = 5  // Monaco
	Montreal         TrackId = 6  // Montreal
	Silverstone      TrackId = 7  // Silverstone
	Hockenheim       TrackId = 8  // Hockenheim
	Hungaroring      TrackId = 9  // Hungaroring
	Spa              TrackId = 10 // Spa
	Monza            TrackId = 11 // Monza
	Singapore        TrackId = 12 // Singapore
	Suzuka           TrackId = 13 // Suzuka
	AbuDhabi         TrackId = 14 // Abu Dhabi
	Texas            TrackId = 15 // Texas
	Brazil           TrackId = 16 // Brazil
	Austria          TrackId = 17 // Austria
	Sochi            TrackId = 18 // Sochi
	Mexico           TrackId = 19 // Mexico
	BakuAzerbaijan   TrackId = 20 // Baku (Azerbaijan)
	SakhirShort      TrackId = 21 // Sakhir Short
	SilverstoneShort TrackId = 22 // Silverstone Short
	TexasShort       TrackId = 23 // Texas Short
	SuzukaShort      TrackId = 24 // Suzuka Short
	Hanoi            TrackId = 25 // Hanoi
	Zandvoort        TrackId = 26 // Zandvoort
	Imola            TrackId = 27 // Imola
	Portimão         TrackId = 28 // Portimão
	Jeddah           TrackId = 29 // Jeddah
	Miami            TrackId = 30 // Miami
	LasVegas         TrackId = 31 // Las Vegas
	Losail           TrackId = 32 // Losail

	Track_Max = Losail // Maximum TrackId
)

func (t TrackId) Id() uint8 {
	return uint8(t)
}

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

	return constants.UNKNOWN
}