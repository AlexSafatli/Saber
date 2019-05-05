package gen

import (
	"../entities"
)

func GenerateCharacter(l *Language, gender string) *entities.Character {
	return entities.NewCharacter(l.Name(), gender)
}
