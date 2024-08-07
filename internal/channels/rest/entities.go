package rest

import "time"

type productRequest struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float32 `json:"price"`
	Stock    int     `json:"stock"`
}

type productResponse struct {
	Id        string    `json:"_id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Price     float32   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
}
