package models

type WantedPost struct {
	BaseModel
	AuthorID    string  `gorm:"type:char(36);not null"`
	Title       string  `gorm:"type:varchar(255);not null"`
	Content     string  `gorm:"type:longtext"`
	BudgetMin   float64 `gorm:"type:decimal(15,2)"`
	BudgetMax   float64 `gorm:"type:decimal(15,2)"`
	Province    string  `gorm:"type:varchar(100)"`
	District    string  `gorm:"type:varchar(100)"`
	PeopleCount int
	Author      User `gorm:"foreignKey:AuthorID"`
}
