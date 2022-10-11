package gen

import (
	"github.com/AlexSafatli/Saber/rng"
	"github.com/AlexSafatli/Saber/rpg"
)

func GenerateCharacter(l *Language, gender string) *rpg.Character {
	char := rpg.NewCharacter(l.Name(), gender)
	char.Sexuality = rng.ChooseWithWeights(rpg.CharacterSexualities,
		rpg.CharacterSexualitiesWeights)
	char.Profession = TableProfessions.Roll()
	return char
}
