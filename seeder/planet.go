package seeder

import (
	"io/ioutil"
)

// PlanetSeed seeds product data
func (s Seed) PlanetSeed() {
	q, err := ioutil.ReadFile(GetSourcePath() + "/scripts/planets.sql")
	if err != nil {
		panic(err)
	}

	_, err = s.db.Exec(string(q))
	if err != nil {
		panic(err)
	}
}
