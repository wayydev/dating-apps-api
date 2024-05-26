package requests

import "time"

type Registration struct {
	Name           string    `json:"name" validate:"required,max=50"`
	Username       string    `json:"username" validate:"required,max=50,unique=users"`
	Email          string    `json:"email" validate:"required,email,unique=users"`
	Password       string    `json:"password,omitempty" validate:"required"`
	Phone          string    `json:"phone" validate:"required,min=10,max=20"`
	Latitude       float32   `json:"latitude" validate:"required,latitude"`
	Longitude      float32   `json:"longitude" validate:"required,longitude"`
	PlaceOfBirth   string    `json:"place_of_birth" validate:"required,max=256"`
	DateOfBirth    time.Time `json:"date_of_birth" validate:"required,date"`
	Gender         string    `json:"gender" validate:"required"`
	PhotoProfile   string    `json:"photo_profile" validate:"image"`
	FindOnDistance uint      `json:"find_on_distance" validate:"required,min:0"`
}
