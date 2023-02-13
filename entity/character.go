package entity

type Character struct {
	Id          int    `json:"id" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	Height      int32  `json:"height" gorm:"column:height"`
	Mass        int32  `json:"mass" gorm:"column:mass"`
	HairColor   string `json:"hairColor" gorm:"column:hair_color"`
	SkinColor   string `json:"skinColor" gorm:"column:skin_color"`
	EyeColor    string `json:"eyeColor" gorm:"column:eye_color"`
	BirthYear   string `json:"birthYear" gorm:"column:birth_year"`
	Gender      string `json:"gender" gorm:"column:gender"`
	HomeworldId int    `json:"homeworld" gorm:"column:homeworld_id"`
	SpeciesId   int    `json:"species" gorm:"column:species_id"`
}
