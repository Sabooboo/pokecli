package typecolor

type Name string
type Value string

const (
	NormalName   Name = "normal"
	FireName     Name = "fire"
	WaterName    Name = "water"
	ElectricName Name = "electric"
	GrassName    Name = "grass"
	IceName      Name = "ice"
	FightingName Name = "fighting"
	PoisonName   Name = "poison"
	GroundName   Name = "ground"
	FlyingName   Name = "flying"
	PsychicName  Name = "psychic"
	BugName      Name = "bug"
	RockName     Name = "rock"
	GhostName    Name = "ghost"
	DragonName   Name = "dragon"
	DarkName     Name = "dark"
	SteelName    Name = "steel"
	FairyName    Name = "fairy"
)

const (
	NormalColor   Value = "#c9c0bf"
	FireColor     Value = "#ff9900"
	WaterColor    Value = "#0099cc"
	ElectricColor Value = "#ffff00"
	GrassColor    Value = "#33cc33"
	IceColor      Value = "#ccffff"
	FightingColor Value = "#ff3300"
	PoisonColor   Value = "#9933ff"
	GroundColor   Value = "#993300"
	FlyingColor   Value = "#6699ff"
	PsychicColor  Value = "#ff0066"
	BugColor      Value = "#99ff66"
	RockColor     Value = "#e09952"
	GhostColor    Value = "#6600cc"
	DragonColor   Value = "#b388dd"
	DarkColor     Value = "#31311b"
	SteelColor    Value = "#b9b9b9"
	FairyColor    Value = "#ff99ff"
	BaseColor     Value = "#999999"
)

type TypeColor struct {
	Name  string
	Color string
}

func Get(name Name) Value {
	switch name {
	case NormalName:
		return NormalColor
	case FireName:
		return FireColor
	case WaterName:
		return WaterColor
	case ElectricName:
		return ElectricColor
	case GrassName:
		return GrassColor
	case IceName:
		return IceColor
	case FightingName:
		return FightingColor
	case PoisonName:
		return PoisonColor
	case GroundName:
		return GroundColor
	case FlyingName:
		return FlyingColor
	case PsychicName:
		return PsychicColor
	case BugName:
		return BugColor
	case RockName:
		return RockColor
	case GhostName:
		return GhostColor
	case DragonName:
		return DragonColor
	case DarkName:
		return DarkColor
	case SteelName:
		return SteelColor
	case FairyName:
		return FairyColor
	default:
		return BaseColor
	}
}
