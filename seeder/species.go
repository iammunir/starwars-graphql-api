package seeder

import (
	"io/ioutil"
)

// SpeciesSeed seeds product data
func (s Seed) SpeciesSeed() {
	q, err := ioutil.ReadFile(GetSourcePath() + "/scripts/species.sql")
	if err != nil {
		panic(err)
	}

	_, err = s.db.Exec(string(q))
	if err != nil {
		panic(err)
	}
}
