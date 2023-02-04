package entity

type Starship struct {
	Id                   int     `json:"id" gorm:"column:id"`
	Name                 string  `json:"name" gorm:"column:name"`
	Model                string  `json:"model" gorm:"column:model"`
	Manufacturer         string  `json:"manufacturer" gorm:"column:manufacturer"`
	CostInCredits        int64   `json:"costInCredits" gorm:"column:costInCredits"`
	Length               float32 `json:"length" gorm:"column:length"`
	MaxAtmospheringSpeed int     `json:"maxAtmospheringSpeed" gorm:"column:maxAtmospheringSpeed"`
	Crew                 int     `json:"crew" gorm:"column:crew"`
	Passengers           int     `json:"passengers" gorm:"column:passengers"`
	CargoCapacity        int64   `json:"cargoCapacity" gorm:"column:cargoCapacity"`
	Consumables          string  `json:"consumables" gorm:"column:consumables"`
	HyperdriveRating     float64 `json:"hyperdriveRating" gorm:"column:hyperdrive_rating"`
	Mglt                 int     `json:"mglt" gorm:"column:mglt"`
	StarshipClass        string  `json:"starshipClass" gorm:"column:starship_class"`
}
