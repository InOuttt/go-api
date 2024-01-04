package domain

import "time"

type (
	Product struct {
		Id          string    `json:"id"`
		Name        string    `json:"name"`
		Price       float64   `json:"price"`
		Description string    `json:"description"`
		Image       string    `json:"image"`
		Created_at  time.Time `json:"-"`
		Created_by  int64     `json:"-"`
		Updated_at  time.Time `json:"-"`
		Updated_by  int64     `json:"-"`
	}
)

type (
	ProductUsecase struct {
	}
)
