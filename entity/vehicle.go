package entity

type Vehicle struct {
	Id                   int     `json:"id" gorm:"column:id"`
	Name                 string  `json:"name" gorm:"column:name"`
	Model                string  `json:"model" gorm:"column:model"`
	Manufacturer         string  `json:"manufacturer" gorm:"column:manufacturer"`
	CostInCredits        int     `json:"costInCredits" gorm:"column:costInCredits"`
	Length               float32 `json:"length" gorm:"column:length"`
	MaxAtmospheringSpeed int     `json:"maxAtmospheringSpeed" gorm:"column:maxAtmospheringSpeed"`
	Crew                 int     `json:"crew" gorm:"column:crew"`
	Passengers           int     `json:"passengers" gorm:"column:passengers"`
	CargoCapacity        int     `json:"cargoCapacity" gorm:"column:cargoCapacity"`
	Consumables          string  `json:"consumables" gorm:"column:consumables"`
	VehicleClass         string  `json:"vehicleClass" gorm:"column:vehicleClass"`
}
