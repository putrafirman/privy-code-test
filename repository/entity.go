package repository

import "time"

type Cake struct {
	ID          int       `json:"id" column:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Rating      float32   `json:"rating"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CakeCollection struct {
	Cakes []Cake `json:"items"`
}
