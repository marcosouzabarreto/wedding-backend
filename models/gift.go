package models

type Gift struct {
	ID          string  `json:"id" gorm:"autoincrement"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"imageUrl"`
}
