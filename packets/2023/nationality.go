package f1_2023

import (
	"github.com/DaanV2/go-f1-library/constants"
	"github.com/DaanV2/go-f1-library/game"
)

var _ game.Nationalities = &Nationalities{}

type (
	Nationality uint8
	Nationalities struct{}
)

func (n *Nationalities) Get(id uint8) game.Nationality {
	return Nationality(id)
}

func (n *Nationalities) Max() uint8 {
	return Nationality_Max.Id()
}

const (
	American      Nationality = 1 // American
	Argentinean   Nationality = 2 // Argentinean
	Australian    Nationality = 3 // Australian
	Austrian      Nationality = 4 // Austrian
	Azerbaijani   Nationality = 5 // Azerbaijani
	Bahraini      Nationality = 6 // Bahraini
	Belgian       Nationality = 7 // Belgian
	Bolivian      Nationality = 8 // Bolivian
	Brazilian     Nationality = 9 // Brazilian
	British       Nationality = 10 // British
	Bulgarian     Nationality = 11 // Bulgarian
	Cameroonian   Nationality = 12 // Cameroonian
	Canadian      Nationality = 13 // Canadian
	Chilean       Nationality = 14 // Chilean
	Chinese       Nationality = 15 // Chinese
	Colombian     Nationality = 16 // Colombian
	CostaRican    Nationality = 17 // CostaRican
	Croatian      Nationality = 18 // Croatian
	Cypriot       Nationality = 19 // Cypriot
	Czech         Nationality = 20 // Czech
	Danish        Nationality = 21 // Danish
	Dutch         Nationality = 22 // Dutch
	Ecuadorian    Nationality = 23 // Ecuadorian
	English       Nationality = 24 // English
	Emirian       Nationality = 25 // Emirian
	Estonian      Nationality = 26 // Estonian
	Finnish       Nationality = 27 // Finnish
	French        Nationality = 28 // French
	German        Nationality = 29 // German
	Ghanaian      Nationality = 30 // Ghanaian
	Greek         Nationality = 31 // Greek
	Guatemalan    Nationality = 32 // Guatemalan
	Honduran      Nationality = 33 // Honduran
	HongKonger    Nationality = 34 // HongKonger
	Hungarian     Nationality = 35 // Hungarian
	Icelander     Nationality = 36 // Icelander
	Indian        Nationality = 37 // Indian
	Indonesian    Nationality = 38 // Indonesian
	Irish         Nationality = 39 // Irish
	Israeli       Nationality = 40 // Israeli
	Italian       Nationality = 41 // Italian
	Jamaican      Nationality = 42 // Jamaican
	Japanese      Nationality = 43 // Japanese
	Jordanian     Nationality = 44 // Jordanian
	Kuwaiti       Nationality = 45 // Kuwaiti
	Latvian       Nationality = 46 // Latvian
	Lebanese      Nationality = 47 // Lebanese
	Lithuanian    Nationality = 48 // Lithuanian
	Luxembourger  Nationality = 49 // Luxembourger
	Malaysian     Nationality = 50 // Malaysian
	Maltese       Nationality = 51 // Maltese
	Mexican       Nationality = 52 // Mexican
	Monegasque    Nationality = 53 // Monegasque
	NewZealander  Nationality = 54 // NewZealander
	Nicaraguan    Nationality = 55 // Nicaraguan
	NorthernIrish Nationality = 56 // NorthernIrish
	Norwegian     Nationality = 57 // Norwegian
	Omani         Nationality = 58 // Omani
	Pakistani     Nationality = 59 // Pakistani
	Panamanian    Nationality = 60 // Panamanian
	Paraguayan    Nationality = 61 // Paraguayan
	Peruvian      Nationality = 62 // Peruvian
	Polish        Nationality = 63 // Polish
	Portuguese    Nationality = 64 // Portuguese
	Qatari        Nationality = 65 // Qatari
	Romanian      Nationality = 66 // Romanian
	Russian       Nationality = 67 // Russian
	Salvadoran    Nationality = 68 // Salvadoran
	Saudi         Nationality = 69 // Saudi
	Scottish      Nationality = 70 // Scottish
	Serbian       Nationality = 71 // Serbian
	Singaporean   Nationality = 72 // Singaporean
	Slovakian     Nationality = 73 // Slovakian
	Slovenian     Nationality = 74 // Slovenian
	SouthKorean   Nationality = 75 // SouthKorean
	SouthAfrican  Nationality = 76 // SouthAfrican
	Spanish       Nationality = 77 // Spanish
	Swedish       Nationality = 78 // Swedish
	Swiss         Nationality = 79 // Swiss
	Thai          Nationality = 80 // Thai
	Turkish       Nationality = 81 // Turkish
	Uruguayan     Nationality = 82 // Uruguayan
	Ukrainian     Nationality = 83 // Ukrainian
	Venezuelan    Nationality = 84 // Venezuelan
	Barbadian     Nationality = 85 // Barbadian
	Welsh         Nationality = 86 // Welsh
	Vietnamese    Nationality = 87 // Vietnamese

	Nationality_Max = Vietnamese
)

func (n Nationality) Id() uint8 {
	return uint8(n)
}

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

	return constants.UNKNOWN
}