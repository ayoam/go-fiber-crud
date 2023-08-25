package model

type Note struct {
	ID      uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Content string `gorm:"not null" json:"content"`
}
