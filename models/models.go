package models

import "time"

type User struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
}