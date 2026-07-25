package models

type Report struct {
	BaseModel
	ReporterID string `gorm:"type:char(36);not null"`
	PostID     string `gorm:"type:char(36);not null"`
	Reason     string `gorm:"type:text;not null"`
	Status     string `gorm:"type:enum('pending', 'reviewing', 'resolved', 'rejected' );default:'pending'"`
	Reporter   User   `gorm:"foreignKey:ReporterID"`
	Post       Post   `gorm:"foreignKey:PostID"`
}
