package seeder

import (
	"io/ioutil"
)

// StarshipSeed seeds product data
func (s Seed) StarshipSeed() {
	q, err := ioutil.ReadFile(GetSourcePath() + "/scripts/starships.sql")
	if err != nil {
		panic(err)
	}

	_, err = s.db.Exec(string(q))
	if err != nil {
		panic(err)
	}
}
