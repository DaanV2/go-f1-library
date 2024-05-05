package f1_2023

import (
	"github.com/DaanV2/go-f1-library/constants"
	"github.com/DaanV2/go-f1-library/game"
)

var _ game.Drivers = &Drivers{}

type (
	Drivers struct{}
	Driver  uint8
)

// Get implements game.Drivers.
func (d *Drivers) Get(id uint8) game.Driver {
	return Driver(id)
}

// Max implements game.Drivers.
func (d *Drivers) Max() uint8 {
	return Driver_Max.Id()
}

const (
	CarlosSainz         Driver = 0   // Carlos Sainz
	DaniilKvyat         Driver = 1   // Daniil Kvyat
	DanielRicciardo     Driver = 2   // Daniel Ricciardo
	FernandoAlonso      Driver = 3   // Fernando Alonso
	FelipeMassa         Driver = 4   // Felipe Massa
	KimiRäikkönen       Driver = 6   // Kimi Räikkönen
	LewisHamilton       Driver = 7   // Lewis Hamilton
	MaxVerstappen       Driver = 9   // Max Verstappen
	NicoHulkenburg      Driver = 10  // Nico Hulkenburg
	KevinMagnussen      Driver = 11  // Kevin Magnussen
	RomainGrosjean      Driver = 12  // Romain Grosjean
	SebastianVettel     Driver = 13  // Sebastian Vettel
	SergioPerez         Driver = 14  // Sergio Perez
	ValtteriBottas      Driver = 15  // Valtteri Bottas
	EstebanOcon         Driver = 17  // Esteban Ocon
	LanceStroll         Driver = 19  // Lance Stroll
	ArronBarnes         Driver = 20  // Arron Barnes
	MartinGiles         Driver = 21  // Martin Giles
	AlexMurray          Driver = 22  // Alex Murray
	LucasRoth           Driver = 23  // Lucas Roth
	IgorCorreia         Driver = 24  // Igor Correia
	SophieLevasseur     Driver = 25  // Sophie Levasseur
	JonasSchiffer       Driver = 26  // Jonas Schiffer
	AlainForest         Driver = 27  // Alain Forest
	JayLetourneau       Driver = 28  // Jay Letourneau
	EstoSaari           Driver = 29  // Esto Saari
	YasarAtiyeh         Driver = 30  // Yasar Atiyeh
	CallistoCalabresi   Driver = 31  // Callisto Calabresi
	NaotaIzum           Driver = 32  // Naota Izum
	HowardClarke        Driver = 33  // Howard Clarke
	WilheimKaufmann     Driver = 34  // Wilheim Kaufmann
	MarieLaursen        Driver = 35  // Marie Laursen
	FlavioNieves        Driver = 36  // Flavio Nieves
	PeterBelousov       Driver = 37  // Peter Belousov
	KlimekMichalski     Driver = 38  // Klimek Michalski
	SantiagoMoreno      Driver = 39  // Santiago Moreno
	BenjaminCoppens     Driver = 40  // Benjamin Coppens
	NoahVisser          Driver = 41  // Noah Visser
	GertWaldmuller      Driver = 42  // Gert Waldmuller
	JulianQuesada       Driver = 43  // Julian Quesada
	DanielJones         Driver = 44  // Daniel Jones
	ArtemMarkelov       Driver = 45  // Artem Markelov
	TadasukeMakino      Driver = 46  // Tadasuke Makino
	SeanGelael          Driver = 47  // Sean Gelael
	NyckDeVries         Driver = 48  // Nyck De Vries
	JackAitken          Driver = 49  // Jack Aitken
	GeorgeRussell       Driver = 50  // George Russell
	MaximilianGünther   Driver = 51  // Maximilian Günther
	NireiFukuzumi       Driver = 52  // Nirei Fukuzumi
	LucaGhiotto         Driver = 53  // Luca Ghiotto
	LandoNorris         Driver = 54  // Lando Norris
	SérgioSetteCâmara   Driver = 55  // Sérgio Sette Câmara
	LouisDelétraz       Driver = 56  // Louis Delétraz
	AntonioFuoco        Driver = 57  // Antonio Fuoco
	CharlesLeclerc      Driver = 58  // Charles Leclerc
	PierreGasly         Driver = 59  // Pierre Gasly
	AlexanderAlbon      Driver = 62  // Alexander Albon
	NicholasLatifi      Driver = 63  // Nicholas Latifi
	DorianBoccolacci    Driver = 64  // Dorian Boccolacci
	NikoKari            Driver = 65  // Niko Kari
	RobertoMerhi        Driver = 66  // Roberto Merhi
	ArjunMaini          Driver = 67  // Arjun Maini
	AlessioLorandi      Driver = 68  // Alessio Lorandi
	RubenMeijer         Driver = 69  // Ruben Meijer
	RashidNair          Driver = 70  // Rashid Nair
	JackTremblay        Driver = 71  // Jack Tremblay
	DevonButler         Driver = 72  // Devon Butler
	LukasWeber          Driver = 73  // Lukas Weber
	AntonioGiovinazzi   Driver = 74  // Antonio Giovinazzi
	RobertKubica        Driver = 75  // Robert Kubica
	AlainProst          Driver = 76  // Alain Prost
	AyrtonSenna         Driver = 77  // Ayrton Senna
	NobuharuMatsushita  Driver = 78  // Nobuharu Matsushita
	NikitaMazepin       Driver = 79  // Nikita Mazepin
	GuanyaZhou          Driver = 80  // Guanya Zhou
	MickSchumacher      Driver = 81  // Mick Schumacher
	CallumIlott         Driver = 82  // Callum Ilott
	JuanManuelCorrea    Driver = 83  // Juan Manuel Correa
	JordanKing          Driver = 84  // Jordan King
	MahaveerRaghunathan Driver = 85  // Mahaveer Raghunathan
	TatianaCalderon     Driver = 86  // Tatiana Calderon
	AnthoineHubert      Driver = 87  // Anthoine Hubert
	GuilianoAlesi       Driver = 88  // Guiliano Alesi
	RalphBoschung       Driver = 89  // Ralph Boschung
	MichaelSchumacher   Driver = 90  // Michael Schumacher
	DanTicktum          Driver = 91  // Dan Ticktum
	MarcusArmstrong     Driver = 92  // Marcus Armstrong
	ChristianLundgaard  Driver = 93  // Christian Lundgaard
	YukiTsunoda         Driver = 94  // Yuki Tsunoda
	JehanDaruvala       Driver = 95  // Jehan Daruvala
	GulhermeSamaia      Driver = 96  // Gulherme Samaia
	PedroPiquet         Driver = 97  // Pedro Piquet
	FelipeDrugovich     Driver = 98  // Felipe Drugovich
	RobertSchwartzman   Driver = 99  // Robert Schwartzman
	RoyNissany          Driver = 100 // Roy Nissany
	MarinoSato          Driver = 101 // Marino Sato
	AidanJackson        Driver = 102 // Aidan Jackson
	CasperAkkerman      Driver = 103 // Casper Akkerman
	JensonButton        Driver = 109 // Jenson Button
	DavidCoulthard      Driver = 110 // David Coulthard
	NicoRosberg         Driver = 111 // Nico Rosberg
	OscarPiastri        Driver = 112 // Oscar Piastri
	LiamLawson          Driver = 113 // Liam Lawson
	JuriVips            Driver = 114 // Juri Vips
	TheoPourchaire      Driver = 115 // Theo Pourchaire
	RichardVerschoor    Driver = 116 // Richard Verschoor
	LirimZendeli        Driver = 117 // Lirim Zendeli
	DavidBeckmann       Driver = 118 // David Beckmann
	AlessioDeledda      Driver = 121 // Alessio Deledda
	BentViscaal         Driver = 122 // Bent Viscaal
	EnzoFittipaldi      Driver = 123 // Enzo Fittipaldi
	MarkWebber          Driver = 125 // Mark Webber
	JacquesVilleneuve   Driver = 126 // Jacques Villeneuve
	CallieMayer         Driver = 127 // Callie Mayer
	NoahBell            Driver = 128 // Noah Bell
	JakeHughes          Driver = 129 // Jake Hughes
	FrederikVesti       Driver = 130 // Frederik Vesti
	OlliCaldwell        Driver = 131 // Olli Caldwell
	LoganSargeant       Driver = 132 // Logan Sargeant
	CemBolukbasi        Driver = 133 // Cem Bolukbasi
	AyumuIwasa          Driver = 134 // Ayumu Iwasa
	ClementNovalak      Driver = 135 // Clement Novalak
	JackDoohan          Driver = 136 // Jack Doohan
	AmauryCordeel       Driver = 137 // Amaury Cordeel
	DennisHauger        Driver = 138 // Dennis Hauger
	CalanWilliams       Driver = 139 // Calan Williams
	JamieChadwick       Driver = 140 // Jamie Chadwick
	KamuiKobayashi      Driver = 141 // Kamui Kobayashi
	PastorMaldonado     Driver = 142 // Pastor Maldonado
	MikaHakkinen        Driver = 143 // Mika Hakkinen
	NigelMansell        Driver = 144 // Nigel Mansell

	Driver_Human Driver = 255 

	Driver_Max = NigelMansell
)

func (d Driver) String() string {
	switch d {
	case CarlosSainz:
		return "Carlos Sainz"
	case DaniilKvyat:
		return "Daniil Kvyat"
	case DanielRicciardo:
		return "Daniel Ricciardo"
	case FernandoAlonso:
		return "Fernando Alonso"
	case FelipeMassa:
		return "Felipe Massa"
	case KimiRäikkönen:
		return "Kimi Räikkönen"
	case LewisHamilton:
		return "Lewis Hamilton"
	case MaxVerstappen:
		return "Max Verstappen"
	case NicoHulkenburg:
		return "Nico Hulkenburg"
	case KevinMagnussen:
		return "Kevin Magnussen"
	case RomainGrosjean:
		return "Romain Grosjean"
	case SebastianVettel:
		return "Sebastian Vettel"
	case SergioPerez:
		return "Sergio Perez"
	case ValtteriBottas:
		return "Valtteri Bottas"
	case EstebanOcon:
		return "Esteban Ocon"
	case LanceStroll:
		return "Lance Stroll"
	case ArronBarnes:
		return "Arron Barnes"
	case MartinGiles:
		return "Martin Giles"
	case AlexMurray:
		return "Alex Murray"
	case LucasRoth:
		return "Lucas Roth"
	case IgorCorreia:
		return "Igor Correia"
	case SophieLevasseur:
		return "Sophie Levasseur"
	case JonasSchiffer:
		return "Jonas Schiffer"
	case AlainForest:
		return "Alain Forest"
	case JayLetourneau:
		return "Jay Letourneau"
	case EstoSaari:
		return "Esto Saari"
	case YasarAtiyeh:
		return "Yasar Atiyeh"
	case CallistoCalabresi:
		return "Callisto Calabresi"
	case NaotaIzum:
		return "Naota Izum"
	case HowardClarke:
		return "Howard Clarke"
	case WilheimKaufmann:
		return "Wilheim Kaufmann"
	case MarieLaursen:
		return "Marie Laursen"
	case FlavioNieves:
		return "Flavio Nieves"
	case PeterBelousov:
		return "Peter Belousov"
	case KlimekMichalski:
		return "Klimek Michalski"
	case SantiagoMoreno:
		return "Santiago Moreno"
	case BenjaminCoppens:
		return "Benjamin Coppens"
	case NoahVisser:
		return "Noah Visser"
	case GertWaldmuller:
		return "Gert Waldmuller"
	case JulianQuesada:
		return "Julian Quesada"
	case DanielJones:
		return "Daniel Jones"
	case ArtemMarkelov:
		return "Artem Markelov"
	case TadasukeMakino:
		return "Tadasuke Makino"
	case SeanGelael:
		return "Sean Gelael"
	case NyckDeVries:
		return "Nyck De Vries"
	case JackAitken:
		return "Jack Aitken"
	case GeorgeRussell:
		return "George Russell"
	case MaximilianGünther:
		return "Maximilian Günther"
	case NireiFukuzumi:
		return "Nirei Fukuzumi"
	case LucaGhiotto:
		return "Luca Ghiotto"
	case LandoNorris:
		return "Lando Norris"
	case SérgioSetteCâmara:
		return "Sérgio Sette Câmara"
	case LouisDelétraz:
		return "Louis Delétraz"
	case AntonioFuoco:
		return "Antonio Fuoco"
	case CharlesLeclerc:
		return "Charles Leclerc"
	case PierreGasly:
		return "Pierre Gasly"
	case AlexanderAlbon:
		return "Alexander Albon"
	case NicholasLatifi:
		return "Nicholas Latifi"
	case DorianBoccolacci:
		return "Dorian Boccolacci"
	case NikoKari:
		return "Niko Kari"
	case RobertoMerhi:
		return "Roberto Merhi"
	case ArjunMaini:
		return "Arjun Maini"
	case AlessioLorandi:
		return "Alessio Lorandi"
	case RubenMeijer:
		return "Ruben Meijer"
	case RashidNair:
		return "Rashid Nair"
	case JackTremblay:
		return "Jack Tremblay"
	case DevonButler:
		return "Devon Butler"
	case LukasWeber:
		return "Lukas Weber"
	case AntonioGiovinazzi:
		return "Antonio Giovinazzi"
	case RobertKubica:
		return "Robert Kubica"
	case AlainProst:
		return "Alain Prost"
	case AyrtonSenna:
		return "Ayrton Senna"
	case NobuharuMatsushita:
		return "Nobuharu Matsushita"
	case NikitaMazepin:
		return "Nikita Mazepin"
	case GuanyaZhou:
		return "Guanya Zhou"
	case MickSchumacher:
		return "Mick Schumacher"
	case CallumIlott:
		return "Callum Ilott"
	case JuanManuelCorrea:
		return "Juan Manuel Correa"
	case JordanKing:
		return "Jordan King"
	case MahaveerRaghunathan:
		return "Mahaveer Raghunathan"
	case TatianaCalderon:
		return "Tatiana Calderon"
	case AnthoineHubert:
		return "Anthoine Hubert"
	case GuilianoAlesi:
		return "Guiliano Alesi"
	case RalphBoschung:
		return "Ralph Boschung"
	case MichaelSchumacher:
		return "Michael Schumacher"
	case DanTicktum:
		return "Dan Ticktum"
	case MarcusArmstrong:
		return "Marcus Armstrong"
	case ChristianLundgaard:
		return "Christian Lundgaard"
	case YukiTsunoda:
		return "Yuki Tsunoda"
	case JehanDaruvala:
		return "Jehan Daruvala"
	case GulhermeSamaia:
		return "Gulherme Samaia"
	case PedroPiquet:
		return "Pedro Piquet"
	case FelipeDrugovich:
		return "Felipe Drugovich"
	case RobertSchwartzman:
		return "Robert Schwartzman"
	case RoyNissany:
		return "Roy Nissany"
	case MarinoSato:
		return "Marino Sato"
	case AidanJackson:
		return "Aidan Jackson"
	case CasperAkkerman:
		return "Casper Akkerman"
	case JensonButton:
		return "Jenson Button"
	case DavidCoulthard:
		return "David Coulthard"
	case NicoRosberg:
		return "Nico Rosberg"
	case OscarPiastri:
		return "Oscar Piastri"
	case LiamLawson:
		return "Liam Lawson"
	case JuriVips:
		return "Juri Vips"
	case TheoPourchaire:
		return "Theo Pourchaire"
	case RichardVerschoor:
		return "Richard Verschoor"
	case LirimZendeli:
		return "Lirim Zendeli"
	case DavidBeckmann:
		return "David Beckmann"
	case AlessioDeledda:
		return "Alessio Deledda"
	case BentViscaal:
		return "Bent Viscaal"
	case EnzoFittipaldi:
		return "Enzo Fittipaldi"
	case MarkWebber:
		return "Mark Webber"
	case JacquesVilleneuve:
		return "Jacques Villeneuve"
	case CallieMayer:
		return "Callie Mayer"
	case NoahBell:
		return "Noah Bell"
	case JakeHughes:
		return "Jake Hughes"
	case FrederikVesti:
		return "Frederik Vesti"
	case OlliCaldwell:
		return "Olli Caldwell"
	case LoganSargeant:
		return "Logan Sargeant"
	case CemBolukbasi:
		return "Cem Bolukbasi"
	case AyumuIwasa:
		return "Ayumu Iwasa"
	case ClementNovalak:
		return "Clement Novalak"
	case JackDoohan:
		return "Jack Doohan"
	case AmauryCordeel:
		return "Amaury Cordeel"
	case DennisHauger:
		return "Dennis Hauger"
	case CalanWilliams:
		return "Calan Williams"
	case JamieChadwick:
		return "Jamie Chadwick"
	case KamuiKobayashi:
		return "Kamui Kobayashi"
	case PastorMaldonado:
		return "Pastor Maldonado"
	case MikaHakkinen:
		return "Mika Hakkinen"
	case NigelMansell:
		return "Nigel Mansell"
	case Driver_Human:
		return "Human"
	}

	return constants.UNKNOWN
}

func (d Driver) Id() uint8 {
	return uint8(d)
}
