package f1_2023

type Nationality int8

const (
	American      Nationality = 1
	Argentinean   Nationality = 2
	Australian    Nationality = 3
	Austrian      Nationality = 4
	Azerbaijani   Nationality = 5
	Bahraini      Nationality = 6
	Belgian       Nationality = 7
	Bolivian      Nationality = 8
	Brazilian     Nationality = 9
	British       Nationality = 10
	Bulgarian     Nationality = 11
	Cameroonian   Nationality = 12
	Canadian      Nationality = 13
	Chilean       Nationality = 14
	Chinese       Nationality = 15
	Colombian     Nationality = 16
	CostaRican    Nationality = 17
	Croatian      Nationality = 18
	Cypriot       Nationality = 19
	Czech         Nationality = 20
	Danish        Nationality = 21
	Dutch         Nationality = 22
	Ecuadorian    Nationality = 23
	English       Nationality = 24
	Emirian       Nationality = 25
	Estonian      Nationality = 26
	Finnish       Nationality = 27
	French        Nationality = 28
	German        Nationality = 29
	Ghanaian      Nationality = 30
	Greek         Nationality = 31
	Guatemalan    Nationality = 32
	Honduran      Nationality = 33
	HongKonger    Nationality = 34
	Hungarian     Nationality = 35
	Icelander     Nationality = 36
	Indian        Nationality = 37
	Indonesian    Nationality = 38
	Irish         Nationality = 39
	Israeli       Nationality = 40
	Italian       Nationality = 41
	Jamaican      Nationality = 42
	Japanese      Nationality = 43
	Jordanian     Nationality = 44
	Kuwaiti       Nationality = 45
	Latvian       Nationality = 46
	Lebanese      Nationality = 47
	Lithuanian    Nationality = 48
	Luxembourger  Nationality = 49
	Malaysian     Nationality = 50
	Maltese       Nationality = 51
	Mexican       Nationality = 52
	Monegasque    Nationality = 53
	NewZealander  Nationality = 54
	Nicaraguan    Nationality = 55
	NorthernIrish Nationality = 56
	Norwegian     Nationality = 57
	Omani         Nationality = 58
	Pakistani     Nationality = 59
	Panamanian    Nationality = 60
	Paraguayan    Nationality = 61
	Peruvian      Nationality = 62
	Polish        Nationality = 63
	Portuguese    Nationality = 64
	Qatari        Nationality = 65
	Romanian      Nationality = 66
	Russian       Nationality = 67
	Salvadoran    Nationality = 68
	Saudi         Nationality = 69
	Scottish      Nationality = 70
	Serbian       Nationality = 71
	Singaporean   Nationality = 72
	Slovakian     Nationality = 73
	Slovenian     Nationality = 74
	SouthKorean   Nationality = 75
	SouthAfrican  Nationality = 76
	Spanish       Nationality = 77
	Swedish       Nationality = 78
	Swiss         Nationality = 79
	Thai          Nationality = 80
	Turkish       Nationality = 81
	Uruguayan     Nationality = 82
	Ukrainian     Nationality = 83
	Venezuelan    Nationality = 84
	Barbadian     Nationality = 85
	Welsh         Nationality = 86
	Vietnamese    Nationality = 87
)

func (n Nationality) String() string {
	switch n {
	case American:
		return "American"
	case Argentinean:
		return "Argentinean"
	case Australian:
		return "Australian"
	case Austrian:
		return "Austrian"
	case Azerbaijani:
		return "Azerbaijani"
	case Bahraini:
		return "Bahraini"
	case Belgian:
		return "Belgian"
	case Bolivian:
		return "Bolivian"
	case Brazilian:
		return "Brazilian"
	case British:
		return "British"
	case Bulgarian:
		return "Bulgarian"
	case Cameroonian:
		return "Cameroonian"
	case Canadian:
		return "Canadian"
	case Chilean:
		return "Chilean"
	case Chinese:
		return "Chinese"
	case Colombian:
		return "Colombian"
	case CostaRican:
		return "CostaRican"
	case Croatian:
		return "Croatian"
	case Cypriot:
		return "Cypriot"
	case Czech:
		return "Czech"
	case Danish:
		return "Danish"
	case Dutch:
		return "Dutch"
	case Ecuadorian:
		return "Ecuadorian"
	case English:
		return "English"
	case Emirian:
		return "Emirian"
	case Estonian:
		return "Estonian"
	case Finnish:
		return "Finnish"
	case French:
		return "French"
	case German:
		return "German"
	case Ghanaian:
		return "Ghanaian"
	case Greek:
		return "Greek"
	case Guatemalan:
		return "Guatemalan"
	case Honduran:
		return "Honduran"
	case HongKonger:
		return "HongKonger"
	case Hungarian:
		return "Hungarian"
	case Icelander:
		return "Icelander"
	case Indian:
		return "Indian"
	case Indonesian:
		return "Indonesian"
	case Irish:
		return "Irish"
	case Israeli:
		return "Israeli"
	case Italian:
		return "Italian"
	case Jamaican:
		return "Jamaican"
	case Japanese:
		return "Japanese"
	case Jordanian:
		return "Jordanian"
	case Kuwaiti:
		return "Kuwaiti"
	case Latvian:
		return "Latvian"
	case Lebanese:
		return "Lebanese"
	case Lithuanian:
		return "Lithuanian"
	case Luxembourger:
		return "Luxembourger"
	case Malaysian:
		return "Malaysian"
	case Maltese:
		return "Maltese"
	case Mexican:
		return "Mexican"
	case Monegasque:
		return "Monegasque"
	case NewZealander:
		return "NewZealander"
	case Nicaraguan:
		return "Nicaraguan"
	case NorthernIrish:
		return "NorthernIrish"
	case Norwegian:
		return "Norwegian"
	case Omani:
		return "Omani"
	case Pakistani:
		return "Pakistani"
	case Panamanian:
		return "Panamanian"
	case Paraguayan:
		return "Paraguayan"
	case Peruvian:
		return "Peruvian"
	case Polish:
		return "Polish"
	case Portuguese:
		return "Portuguese"
	case Qatari:
		return "Qatari"
	case Romanian:
		return "Romanian"
	case Russian:
		return "Russian"
	case Salvadoran:
		return "Salvadoran"
	case Saudi:
		return "Saudi"
	case Scottish:
		return "Scottish"
	case Serbian:
		return "Serbian"
	case Singaporean:
		return "Singaporean"
	case Slovakian:
		return "Slovakian"
	case Slovenian:
		return "Slovenian"
	case SouthKorean:
		return "SouthKorean"
	case SouthAfrican:
		return "SouthAfrican"
	case Spanish:
		return "Spanish"
	case Swedish:
		return "Swedish"
	case Swiss:
		return "Swiss"
	case Thai:
		return "Thai"
	case Turkish:
		return "Turkish"
	case Uruguayan:
		return "Uruguayan"
	case Ukrainian:
		return "Ukrainian"
	case Venezuelan:
		return "Venezuelan"
	case Barbadian:
		return "Barbadian"
	case Welsh:
		return "Welsh"
	case Vietnamese:
		return "Vietnamese"
	}

	return "UNKNOWN"
}
