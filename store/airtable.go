package store

import (
	"../entities"
	"github.com/fabioberger/airtable-go"
	"github.com/spf13/viper"
	"strconv"
)

var (
	apiKey = viper.GetString("airtableApiKey")
	baseID = viper.GetString("airtableBase")
)

func NewAirtableConnection() (*airtable.Client, error) {
	return airtable.New(apiKey, baseID)
}

func ReadCharacterFromAirtable(client *airtable.Client, id string) *entities.Character {
	var char map[string]string
	if err := client.RetrieveRecord("Characters", id, &char); err != nil {
		return nil
	}
	level, err := strconv.ParseUint(char["Level"], 10, 64)
	if err != nil {
		level = 0
	}
	return &entities.Character{
		Name:  char["Name"],
		Race:  char["Race"],
		Level: level,
		Characteristics: map[string]entities.EntityTag{
			"AC":          {Name: "AC", Value: char["AC"]},
			"Initiative":  {Name: "Initiative", Value: char["Initiative"]},
			"Dex ST":      {Name: "Dex ST", Value: char["Dex ST"]},
			"Passive Per": {Name: "Passive Per", Value: char["Passive Per"]},
		},
	}
}
