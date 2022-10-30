package poketype

import (
	"github.com/Sabooboo/pokecli/ui/typdef/typecolor"
	apistructs "github.com/mtslzr/pokeapi-go/structs"
)

type DamageRelation float32

const (
	NoDamage      DamageRelation = 0
	QuarterDamage DamageRelation = 0.25
	HalfDamage    DamageRelation = 0.5
	DoubleDamage  DamageRelation = 2
	QuadDamage    DamageRelation = 4
)

// TypeMatchups represents a set of damage relations any one or two types may
// have. Generally instantiated through GetTypeMatchups.
type TypeMatchups struct {
	Immunities       []typecolor.Name
	Resistances      []typecolor.Name
	Weaknesses       []typecolor.Name
	MajorResistances []typecolor.Name
	MajorWeaknesses  []typecolor.Name
}

// FromMap takes a map[DamageRelation][]typecolor.Name and converts it to a
// TypeMatchups, returning the resulting struct.
func FromMap(m map[DamageRelation][]typecolor.Name) TypeMatchups {
	return TypeMatchups{
		Immunities:       m[NoDamage],
		Resistances:      m[HalfDamage],
		Weaknesses:       m[DoubleDamage],
		MajorResistances: m[QuarterDamage],
		MajorWeaknesses:  m[QuadDamage],
	}
}

// FromTypeIndexableMap takes a map[typecolor.Name]DamageRelation and converts
// it to a TypeMatchups, returning the resulting struct.
func FromTypeIndexableMap(m map[typecolor.Name]DamageRelation) TypeMatchups {
	matchups := TypeMatchups{}.AsMap()
	for name, dmg := range m {
		matchups[dmg] = append(matchups[dmg], name)
	}
	return FromMap(matchups)
}

// GetTypeMatchups takes one or two pokeapi-go Types and returns a relevant
// TypeMatchups. If one Type is given, MajorResistances and MajorWeaknesses
// will be nil. If two Types are given, this is the same as CompareTypeMatchups.
// If any amount other than one or two are given, an empty TypeMatchups will
// be returned.
func GetTypeMatchups(types ...apistructs.Type) TypeMatchups {
	if len(types) < 1 || len(types) > 2 {
		return TypeMatchups{}
	}
	if len(types) == 2 {
		return CompareTypeMatchups(types[0], types[1])
	}
	// length of types must therefore be 1
	return ConvertTypeMatchups(types[0])
}

// CompareTypeMatchups takes two pokeapi-go Types and compares them. This
// returns a TypeMatchups with relevant information.
func CompareTypeMatchups(a, b apistructs.Type) TypeMatchups {
	typesA := ConvertTypeMatchups(a).AsTypeIndexableMap()
	typesB := ConvertTypeMatchups(b).AsTypeIndexableMap()

	matchups := TypeMatchups{}.AsTypeIndexableMap()

	for name, val := range typesA {
		matchups[name] = val
	}

	for name, val := range typesB {
		if mapValue, exists := matchups[name]; exists {
			matchups[name] = mapValue * val
		} else {
			matchups[name] = val
		}
	}

	return FromTypeIndexableMap(matchups)
}

// AsMap returns a map with float32 indexable values. This is useful
// when creating type-matchups as you can multiply keys of one TypeMatchups
// field with another to find where they belong in the relations list.
func (t TypeMatchups) AsMap() map[DamageRelation][]typecolor.Name {
	matchups := make(map[DamageRelation][]typecolor.Name)
	matchups[NoDamage] = t.Immunities
	matchups[HalfDamage] = t.Resistances
	matchups[DoubleDamage] = t.Weaknesses
	matchups[QuarterDamage] = t.MajorResistances
	matchups[QuadDamage] = t.MajorWeaknesses
	return matchups
}

// AsTypeIndexableMap returns a map where the indexes are
func (t TypeMatchups) AsTypeIndexableMap() map[typecolor.Name]DamageRelation {
	matchups := make(map[typecolor.Name]DamageRelation)
	for _, v := range t.Immunities {
		matchups[v] = NoDamage
	}
	for _, v := range t.Resistances {
		matchups[v] = HalfDamage
	}
	for _, v := range t.Weaknesses {
		matchups[v] = DoubleDamage
	}
	for _, v := range t.MajorResistances {
		matchups[v] = QuarterDamage
	}
	for _, v := range t.MajorWeaknesses {
		matchups[v] = QuadDamage
	}
	return matchups
}

// ConvertTypeMatchups takes a pokeapi-go Type and returns the corresponding
// matchups as a TypeMatchups.
func ConvertTypeMatchups(t apistructs.Type) TypeMatchups {
	immunities := make([]typecolor.Name, 0)
	for _, v := range t.DamageRelations.NoDamageFrom {
		immunities = append(immunities, typecolor.Name(v.Name))
	}

	// HalfDamageFrom (as well as DoubleDamageTo) are both interface{}'s.
	// I suspect this is because the normal type has no resistances/strengths
	// and pokeapi-go had difficulties trying to unmarshal.
	resistances := make([]typecolor.Name, 0)
	for _, v := range t.DamageRelations.HalfDamageFrom {
		// Manual type checking to account for this nonsense lol
		if resistance, ok := v.(map[string]interface{}); ok {
			if name, reallyOk := resistance["name"].(string); reallyOk {
				resistances = append(resistances, typecolor.Name(name))
			}
		}
	}

	weaknesses := make([]typecolor.Name, 0)
	for _, v := range t.DamageRelations.DoubleDamageFrom {
		weaknesses = append(weaknesses, typecolor.Name(v.Name))
	}

	return TypeMatchups{
		Immunities:  immunities,
		Resistances: resistances,
		Weaknesses:  weaknesses,
	}
}
