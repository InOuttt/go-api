package domain

import (
	"context"
	"time"
)

type (
	Customer struct {
		Id        int64     `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Password  string    `json:"-"`
		AuthLevel int64     `json:"auth_level"`
		CreatedAt time.Time `json:"-"`
		CreatedBy string    `json:"-"`
		UpdatedAt time.Time `json:"-"`
		UpdatedBy string    `json:"-"`
	}
	AuthResponse struct {
		Token   string  `json:"token"`
		Expired float64 `json:"expired_in"`
	}
)

type (
	CustomerUsecase interface {
		Auth(ctx context.Context, email, password string) (AuthResponse, error)
	}

	CustomerRepository interface {
		GetByEmail(ctx context.Context, email string) (Customer, error)
	}
)
