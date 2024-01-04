package domain

import (
	"context"
	"time"
)

type (
	Order struct {
		Id          int64     `json:"id"`
		Id_customer int64     `json:"id_customer" validate:"numeric"`
		Date        time.Time `json:"date"`
		Status      int64     `json:"status" `
		Created_at  time.Time `json:"-"`
		Created_by  int64     `json:"-"`
		Updated_at  time.Time `json:"-"`
		Updated_by  int64     `json:"-"`
	}
	OrderProduct struct {
		Order
		Products []Product `json:"products"`
	}
	OrderReport struct {
		Id           int64
		CustomerName string
		Date         time.Time
		Price        float64
		Status       int64
	}
	OrderProductCustomer struct {
		OrderProduct
		CustomerName  string
		CustomerEmail string
	}
)

type (
	OrderUsecase interface {
		Create(ctx context.Context, order Order, products []int64) (Order, error)
		GetAll(ctx context.Context, customerId int64, paging Pagination) (response []OrderProduct, total int64, err error)
		GenerateReport(ctx context.Context) (response []OrderReport, err error)
	}

	OrderRepository interface {
		Create(ctx context.Context, order Order, products []int64) (Order, error)
		CountByCustomer(ctx context.Context, customerId int64) (total int64, err error)
		CountAll(ctx context.Context) (total int64, err error)
		GetAllByCustomer(ctx context.Context, customerId int64, paging Pagination) (response []OrderProduct, err error)
		GetOrderProductCustomerAllByStatus(ctx context.Context, status int64) (response []OrderProductCustomer, err error)
		GetAll(ctx context.Context, paging Pagination) (response []OrderProduct, err error)
		GetReport(ctx context.Context) ([]OrderReport, error)
	}
)
