package model

type Menu struct {
	Status *bool  `json:"status" 	gorm:"not null;index;" `
	Id     string `json:"id" 	gorm:"primaryKey;autoIncrement;" `
	enetity.Global
}
