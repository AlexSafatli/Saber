package gen

import (
	"../entities"
)

func GenerateCharacter(l *Language, gender string) *entities.Character {
	char := entities.NewCharacter(l.Name(), gender)
	char.Profession = TableProfessions.Roll()
	return char
}
