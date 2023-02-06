package entity

type Vehicle struct {
	Id                   int     `json:"id" gorm:"column:id"`
	Name                 string  `json:"name" gorm:"column:name"`
	Model                string  `json:"model" gorm:"column:model"`
	Manufacturer         string  `json:"manufacturer" gorm:"column:manufacturer"`
	CostInCredits        int     `json:"costInCredits" gorm:"column:cost_in_credits"`
	Length               float32 `json:"length" gorm:"column:length"`
	MaxAtmospheringSpeed int     `json:"maxAtmospheringSpeed" gorm:"column:max_atmosphering_speed"`
	Crew                 int     `json:"crew" gorm:"column:crew"`
	Passengers           int     `json:"passengers" gorm:"column:passengers"`
	CargoCapacity        int     `json:"cargoCapacity" gorm:"column:cargo_capacity"`
	Consumables          string  `json:"consumables" gorm:"column:consumables"`
	VehicleClass         string  `json:"vehicleClass" gorm:"column:vehicle_class"`
}
