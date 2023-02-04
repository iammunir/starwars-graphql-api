package entity

type Planet struct {
	Id             int    `json:"id" gorm:"column:id"`
	Name           string `json:"name" gorm:"column:name"`
	RotationPeriod int32  `json:"rotationPeriod" gorm:"column:rotation_period"`
	OrbitalPeriod  int32  `json:"orbitalPeriod" gorm:"column:orbital_period"`
	Diameter       int    `json:"diameter" gorm:"column:diameter"`
	Climate        string `json:"climate" gorm:"column:climate"`
	Gravity        string `json:"gravity" gorm:"column:gravity"`
	Terrain        string `json:"terrain" gorm:"column:terrain"`
	SurfaceWater   int32  `json:"surfaceWater" gorm:"column:surface_water"`
	Population     int64  `json:"population" gorm:"column:population"`
}
