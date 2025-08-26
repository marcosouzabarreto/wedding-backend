package models

import "gorm.io/gorm"

type Gift struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	GifterName  string `json:"gifterName"`
	Message     string `json:"message"`
}
