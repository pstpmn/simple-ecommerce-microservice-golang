package productCore

import "time"

type Product struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"Price"`
	Stock       int64     `json:"stock"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ProductIntroduction struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type Category struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProductProfile struct {
	*Product
	*Category
}
