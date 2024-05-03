package f1_2023

type Driver uint8

const (
	CarlosSainz         Driver = 0
	DaniilKvyat         Driver = 1
	DanielRicciardo     Driver = 2
	FernandoAlonso      Driver = 3
	FelipeMassa         Driver = 4
	KimiRäikkönen       Driver = 6
	LewisHamilton       Driver = 7
	MaxVerstappen       Driver = 9
	NicoHulkenburg      Driver = 10
	KevinMagnussen      Driver = 11
	RomainGrosjean      Driver = 12
	SebastianVettel     Driver = 13
	SergioPerez         Driver = 14
	ValtteriBottas      Driver = 15
	EstebanOcon         Driver = 17
	LanceStroll         Driver = 19
	ArronBarnes         Driver = 20
	MartinGiles         Driver = 21
	AlexMurray          Driver = 22
	LucasRoth           Driver = 23
	IgorCorreia         Driver = 24
	SophieLevasseur     Driver = 25
	JonasSchiffer       Driver = 26
	AlainForest         Driver = 27
	JayLetourneau       Driver = 28
	EstoSaari           Driver = 29
	YasarAtiyeh         Driver = 30
	CallistoCalabresi   Driver = 31
	NaotaIzum           Driver = 32
	HowardClarke        Driver = 33
	WilheimKaufmann     Driver = 34
	MarieLaursen        Driver = 35
	FlavioNieves        Driver = 36
	PeterBelousov       Driver = 37
	KlimekMichalski     Driver = 38
	SantiagoMoreno      Driver = 39
	BenjaminCoppens     Driver = 40
	NoahVisser          Driver = 41
	GertWaldmuller      Driver = 42
	JulianQuesada       Driver = 43
	DanielJones         Driver = 44
	ArtemMarkelov       Driver = 45
	TadasukeMakino      Driver = 46
	SeanGelael          Driver = 47
	NyckDeVries         Driver = 48
	JackAitken          Driver = 49
	GeorgeRussell       Driver = 50
	MaximilianGünther   Driver = 51
	NireiFukuzumi       Driver = 52
	LucaGhiotto         Driver = 53
	LandoNorris         Driver = 54
	SérgioSetteCâmara   Driver = 55
	LouisDelétraz       Driver = 56
	AntonioFuoco        Driver = 57
	CharlesLeclerc      Driver = 58
	PierreGasly         Driver = 59
	AlexanderAlbon      Driver = 62
	NicholasLatifi      Driver = 63
	DorianBoccolacci    Driver = 64
	NikoKari            Driver = 65
	RobertoMerhi        Driver = 66
	ArjunMaini          Driver = 67
	AlessioLorandi      Driver = 68
	RubenMeijer         Driver = 69
	RashidNair          Driver = 70
	JackTremblay        Driver = 71
	DevonButler         Driver = 72
	LukasWeber          Driver = 73
	AntonioGiovinazzi   Driver = 74
	RobertKubica        Driver = 75
	AlainProst          Driver = 76
	AyrtonSenna         Driver = 77
	NobuharuMatsushita  Driver = 78
	NikitaMazepin       Driver = 79
	GuanyaZhou          Driver = 80
	MickSchumacher      Driver = 81
	CallumIlott         Driver = 82
	JuanManuelCorrea    Driver = 83
	JordanKing          Driver = 84
	MahaveerRaghunathan Driver = 85
	TatianaCalderon     Driver = 86
	AnthoineHubert      Driver = 87
	GuilianoAlesi       Driver = 88
	RalphBoschung       Driver = 89
	MichaelSchumacher   Driver = 90
	DanTicktum          Driver = 91
	MarcusArmstrong     Driver = 92
	ChristianLundgaard  Driver = 93
	YukiTsunoda         Driver = 94
	JehanDaruvala       Driver = 95
	GulhermeSamaia      Driver = 96
	PedroPiquet         Driver = 97
	FelipeDrugovich     Driver = 98
	RobertSchwartzman   Driver = 99
	RoyNissany          Driver = 100
	MarinoSato          Driver = 101
	AidanJackson        Driver = 102
	CasperAkkerman      Driver = 103
	JensonButton        Driver = 109
	DavidCoulthard      Driver = 110
	NicoRosberg         Driver = 111
	OscarPiastri        Driver = 112
	LiamLawson          Driver = 113
	JuriVips            Driver = 114
	TheoPourchaire      Driver = 115
	RichardVerschoor    Driver = 116
	LirimZendeli        Driver = 117
	DavidBeckmann       Driver = 118
	AlessioDeledda      Driver = 121
	BentViscaal         Driver = 122
	EnzoFittipaldi      Driver = 123
	MarkWebber          Driver = 125
	JacquesVilleneuve   Driver = 126
	CallieMayer         Driver = 127
	NoahBell            Driver = 128
	JakeHughes          Driver = 129
	FrederikVesti       Driver = 130
	OlliCaldwell        Driver = 131
	LoganSargeant       Driver = 132
	CemBolukbasi        Driver = 133
	AyumuIwasa          Driver = 134
	ClementNovalak      Driver = 135
	JackDoohan          Driver = 136
	AmauryCordeel       Driver = 137
	DennisHauger        Driver = 138
	CalanWilliams       Driver = 139
	JamieChadwick       Driver = 140
	KamuiKobayashi      Driver = 141
	PastorMaldonado     Driver = 142
	MikaHakkinen        Driver = 143
	NigelMansell        Driver = 144
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
		return "Sérgio Sette Câmara	"
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
	default:
		return "UNKNOWN"
	}
}
