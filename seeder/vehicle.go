package seeder

import (
	"io/ioutil"
)

// VehicleSeed seeds product data
func (s Seed) VehicleSeed() {
	q, err := ioutil.ReadFile(GetSourcePath() + "/scripts/vehicles.sql")
	if err != nil {
		panic(err)
	}

	_, err = s.db.Exec(string(q))
	if err != nil {
		panic(err)
	}
}
