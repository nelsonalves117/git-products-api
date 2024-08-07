package canonical

import "time"

type Product struct {
	Id        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Category  string    `bson:"category"`
	Price     float32   `bson:"price"`
	Stock     int       `bson:"stock"`
	CreatedAt time.Time `bson:"created_at"`
}
