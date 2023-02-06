package entity

type Species struct {
	Id              int    `json:"id" gorm:"column:id"`
	Name            string `json:"name" gorm:"column:name"`
	Classification  string `json:"classification" gorm:"column:classification"`
	Designation     string `json:"designation" gorm:"column:designation"`
	AverageHeight   int32  `json:"averageHeight" gorm:"column:average_height"`
	SkinColors      string `json:"skinColors" gorm:"column:skin_colors"`
	HairColors      string `json:"hairColors" gorm:"column:hair_colors"`
	EyeColors       string `json:"eyeColors" gorm:"column:eye_colors"`
	AverageLifespan int32  `json:"averageLifespan" gorm:"column:average_lifespan"`
	Language        string `json:"language" gorm:"column:language"`
	HomeworldId     int    `json:"homeworld" gorm:"column:homeworld_id"`
}
