package ability

type Generator func() Ability

func GenerateAbilities(abilityGenerators []Generator) []*Ability {
	var abilities = make([]*Ability, 0)
	for _, abilityFunction := range abilityGenerators {
		var a = abilityFunction()
		abilities = append(abilities, &a)
	}

	return abilities
}
