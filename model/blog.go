package model

type Blog struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Title string `json:"title" gorm:"not null;coloumn:title;size:255"`

	Post string `json:"title" gorm:"not null;coloumn:title;size:255"`
}
