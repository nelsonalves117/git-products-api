package rest

type productRequest struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float32 `json:"price"`
}

type productResponse struct {
	Id        string  `json:"_id"`
	Name      string  `json:"name"`
	Category  string  `json:"category"`
	Price     float32 `json:"price"`
	CreatedAt string  `json:"created_at"`
}
