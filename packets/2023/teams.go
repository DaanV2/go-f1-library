package f1_2023

type TeamId int16

const (
	Mercedes                          TeamId = 0
	Ferrari                           TeamId = 1
	RedBullRacing                     TeamId = 2
	Williams                          TeamId = 3
	AstonMartin                       TeamId = 4
	Alpine                            TeamId = 5
	AlphaTauri                        TeamId = 6
	Haas                              TeamId = 7
	McLaren                           TeamId = 8
	AlfaRomeo                         TeamId = 9
	Mercedes2020                      TeamId = 85
	Ferrari2020                       TeamId = 86
	RedBull2020                       TeamId = 87
	Williams2020                      TeamId = 88
	RacingPoint2020                   TeamId = 89
	Renault2020                       TeamId = 90
	AlphaTauri2020                    TeamId = 91
	Haas2020                          TeamId = 92
	McLaren2020                       TeamId = 93
	AlfaRomeo2020                     TeamId = 94
	AstonMartinDB11V12                TeamId = 95
	AstonMartinVantageF1Edition       TeamId = 96
	AstonMartinVantageSafetyCar       TeamId = 97
	FerrariF8Tributo                  TeamId = 98
	FerrariRoma                       TeamId = 99
	McLaren720S                       TeamId = 100
	McLarenArtura                     TeamId = 101
	MercedesAMGGTBlackSeriesSafetyCar TeamId = 102
	MercedesAMGGTRPro                 TeamId = 103
	F1CustomTeam                      TeamId = 104
	Prema_21                          TeamId = 106
	Uni_Virtuosi_21                   TeamId = 107
	Carlin_21                         TeamId = 108
	Hitech_21                         TeamId = 109
	ArtGP_21                          TeamId = 110
	MPMotorsport_21                   TeamId = 111
	Charouz_21                        TeamId = 112
	Dams_21                           TeamId = 113
	Campos_21                         TeamId = 114
	BWT_21                            TeamId = 115
	Trident_21                        TeamId = 116
	MercedesAMGGTBlackSeries          TeamId = 117
	Mercedes_22                       TeamId = 118
	Ferrari_22                        TeamId = 119
	RedBullRacing_22                  TeamId = 120
	Williams_22                       TeamId = 121
	AstonMartin_22                    TeamId = 122
	Alpine_22                         TeamId = 123
	AlphaTauri_22                     TeamId = 124
	Haas_22                           TeamId = 125
	McLaren_22                        TeamId = 126
	AlfaRomeo_22                      TeamId = 127
	Konnersport_22                    TeamId = 128
	Konnersport                       TeamId = 129
	Prema_22                          TeamId = 130
	Virtuosi_22                       TeamId = 131
	Carlin_22                         TeamId = 132
	MPMotorsport_22                   TeamId = 133
	Charouz_22                        TeamId = 134
	Dams_22                           TeamId = 135
	Campos_22                         TeamId = 136
	VanAmersfoortRacing_22            TeamId = 137
	Trident_22                        TeamId = 138
	Hitech_22                         TeamId = 139
	ArtGP_22                          TeamId = 140
)

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

	return "UNKNOWN"
}
