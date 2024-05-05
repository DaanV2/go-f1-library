package f1_2023

import (
	"github.com/DaanV2/go-f1-library/constants"
	"github.com/DaanV2/go-f1-library/game"
)

var _ game.Teams = &Teams{}

type (
	TeamId uint16
	Teams struct{}
)

func (d *Teams) Get(id uint8) game.Team {
	return TeamId(id)
}

func (d *Teams) Max() uint8 {
	return Team_Max.Id()
}

const (
	Mercedes                          TeamId = 0 // Mercedes
	Ferrari                           TeamId = 1 // Ferrari
	RedBullRacing                     TeamId = 2 // Red Bull Racing
	Williams                          TeamId = 3 // Williams
	AstonMartin                       TeamId = 4 // Aston Martin
	Alpine                            TeamId = 5 // Alpine
	AlphaTauri                        TeamId = 6 // Alpha Tauri
	Haas                              TeamId = 7 // Haas
	McLaren                           TeamId = 8 // McLaren
	AlfaRomeo                         TeamId = 9 // Alfa Romeo
	Mercedes2020                      TeamId = 85 // Mercedes 2020
	Ferrari2020                       TeamId = 86 // Ferrari 2020
	RedBull2020                       TeamId = 87 // Red Bull 2020
	Williams2020                      TeamId = 88 // Williams 2020
	RacingPoint2020                   TeamId = 89 // Racing Point 2020
	Renault2020                       TeamId = 90 // Renault 2020
	AlphaTauri2020                    TeamId = 91 // Alpha Tauri 2020
	Haas2020                          TeamId = 92 // Haas 2020
	McLaren2020                       TeamId = 93 // McLaren 2020
	AlfaRomeo2020                     TeamId = 94 // Alfa Romeo 2020
	AstonMartinDB11V12                TeamId = 95 // Aston Martin DB11 V12
	AstonMartinVantageF1Edition       TeamId = 96 // Aston Martin Vantage F1 Edition
	AstonMartinVantageSafetyCar       TeamId = 97 // Aston Martin Vantage Safety Car
	FerrariF8Tributo                  TeamId = 98 // Ferrari F8 Tributo
	FerrariRoma                       TeamId = 99 // Ferrari Roma
	McLaren720S                       TeamId = 100 // McLaren 720S
	McLarenArtura                     TeamId = 101 // McLaren Artura
	MercedesAMGGTBlackSeriesSafetyCar TeamId = 102 // Mercedes AMG GT Black Series Safety Car
	MercedesAMGGTRPro                 TeamId = 103 // Mercedes AMG GTR Pro
	F1CustomTeam                      TeamId = 104 // F1 Custom Team
	Prema_21                          TeamId = 106 // Prema '21
	Uni_Virtuosi_21                   TeamId = 107 // Uni-Virtuosi '21
	Carlin_21                         TeamId = 108 // Carlin '21
	Hitech_21                         TeamId = 109 // Hitech '21
	ArtGP_21                          TeamId = 110 // Art GP '21
	MPMotorsport_21                   TeamId = 111 // MP Motorsport '21
	Charouz_21                        TeamId = 112 // Charouz '21
	Dams_21                           TeamId = 113 // Dams '21
	Campos_21                         TeamId = 114 // Campos '21
	BWT_21                            TeamId = 115 // BWT '21
	Trident_21                        TeamId = 116 // Trident '21
	MercedesAMGGTBlackSeries          TeamId = 117 // Mercedes AMG GT Black Series
	Mercedes_22                       TeamId = 118 // Mercedes '22
	Ferrari_22                        TeamId = 119 // Ferrari '22
	RedBullRacing_22                  TeamId = 120 // Red Bull Racing '22
	Williams_22                       TeamId = 121 // Williams '22
	AstonMartin_22                    TeamId = 122 // Aston Martin '22
	Alpine_22                         TeamId = 123 // Alpine '22
	AlphaTauri_22                     TeamId = 124 // Alpha Tauri '22
	Haas_22                           TeamId = 125 // Haas '22
	McLaren_22                        TeamId = 126 // McLaren '22
	AlfaRomeo_22                      TeamId = 127 // Alfa Romeo '22
	Konnersport_22                    TeamId = 128 // Konnersport '22
	Konnersport                       TeamId = 129 // Konnersport
	Prema_22                          TeamId = 130 // Prema '22
	Virtuosi_22                       TeamId = 131 // Virtuosi '22
	Carlin_22                         TeamId = 132 // Carlin '22
	MPMotorsport_22                   TeamId = 133 // MP Motorsport '22
	Charouz_22                        TeamId = 134 // Charouz '22
	Dams_22                           TeamId = 135 // Dams '22
	Campos_22                         TeamId = 136 // Campos '22
	VanAmersfoortRacing_22            TeamId = 137 // Van Amersfoort Racing '22
	Trident_22                        TeamId = 138 // Trident '22
	Hitech_22                         TeamId = 139 // Hitech '22
	ArtGP_22                          TeamId = 140 // Art GP '22

	Team_Max = ArtGP_22
)

func (t TeamId) Id() uint8 {
	return uint8(t)
}

func (t TeamId) String() string {
	switch t {
	case Mercedes:
		return "Mercedes"
	case Ferrari:
		return "Ferrari"
	case RedBullRacing:
		return "Red Bull Racing"
	case Williams:
		return "Williams"
	case AstonMartin:
		return "Aston Martin"
	case Alpine:
		return "Alpine"
	case AlphaTauri:
		return "Alpha Tauri"
	case Haas:
		return "Haas"
	case McLaren:
		return "McLaren"
	case AlfaRomeo:
		return "Alfa Romeo"
	case Mercedes2020:
		return "Mercedes 2020"
	case Ferrari2020:
		return "Ferrari 2020"
	case RedBull2020:
		return "Red Bull 2020"
	case Williams2020:
		return "Williams 2020"
	case RacingPoint2020:
		return "Racing Point 2020"
	case Renault2020:
		return "Renault 2020"
	case AlphaTauri2020:
		return "Alpha Tauri 2020"
	case Haas2020:
		return "Haas 2020"
	case McLaren2020:
		return "McLaren 2020"
	case AlfaRomeo2020:
		return "Alfa Romeo 2020"
	case AstonMartinDB11V12:
		return "Aston Martin DB11 V12"
	case AstonMartinVantageF1Edition:
		return "Aston Martin Vantage F1 Edition"
	case AstonMartinVantageSafetyCar:
		return "Aston Martin Vantage Safety Car"
	case FerrariF8Tributo:
		return "Ferrari F8 Tributo"
	case FerrariRoma:
		return "Ferrari Roma"
	case McLaren720S:
		return "McLaren 720S"
	case McLarenArtura:
		return "McLaren Artura"
	case MercedesAMGGTBlackSeriesSafetyCar:
		return "Mercedes AMG GT Black Series Safety Car"
	case MercedesAMGGTRPro:
		return "Mercedes AMG GTR Pro"
	case F1CustomTeam:
		return "F1 Custom Team"
	case Prema_21:
		return "Prema '21"
	case Uni_Virtuosi_21:
		return "Uni-Virtuosi '21"
	case Carlin_21:
		return "Carlin '21"
	case Hitech_21:
		return "Hitech '21"
	case ArtGP_21:
		return "Art GP '21"
	case MPMotorsport_21:
		return "MP Motorsport '21"
	case Charouz_21:
		return "Charouz '21"
	case Dams_21:
		return "Dams '21"
	case Campos_21:
		return "Campos '21"
	case BWT_21:
		return "BWT '21"
	case Trident_21:
		return "Trident '21"
	case MercedesAMGGTBlackSeries:
		return "Mercedes AMG GT Black Series"
	case Mercedes_22:
		return "Mercedes '22"
	case Ferrari_22:
		return "Ferrari '22"
	case RedBullRacing_22:
		return "Red Bull Racing '22"
	case Williams_22:
		return "Williams '22"
	case AstonMartin_22:
		return "Aston Martin '22"
	case Alpine_22:
		return "Alpine '22"
	case AlphaTauri_22:
		return "Alpha Tauri '22"
	case Haas_22:
		return "Haas '22"
	case McLaren_22:
		return "McLaren '22"
	case AlfaRomeo_22:
		return "Alfa Romeo '22"
	case Konnersport_22:
		return "Konnersport '22"
	case Konnersport:
		return "Konnersport"
	case Prema_22:
		return "Prema '22"
	case Virtuosi_22:
		return "Virtuosi '22"
	case Carlin_22:
		return "Carlin '22"
	case MPMotorsport_22:
		return "MP Motorsport '22"
	case Charouz_22:
		return "Charouz '22"
	case Dams_22:
		return "Dams '22"
	case Campos_22:
		return "Campos '22"
	case VanAmersfoortRacing_22:
		return "Van Amersfoort Racing '22"
	case Trident_22:
		return "Trident '22"
	case Hitech_22:
		return "Hitech '22"
	case ArtGP_22:
		return "Art GP '22"
	}

	return constants.UNKNOWN
}