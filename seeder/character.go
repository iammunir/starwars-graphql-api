package seeder

import (
	"io/ioutil"
)

// CharacterSeed seeds product data
func (s Seed) CharacterSeed() {
	q, err := ioutil.ReadFile(GetSourcePath() + "/scripts/characters.sql")
	if err != nil {
		panic(err)
	}

	_, err = s.db.Exec(string(q))
	if err != nil {
		panic(err)
	}
}
