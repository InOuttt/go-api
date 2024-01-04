package domain

import (
	"context"
	"time"
)

type (
	Record struct {
		Id         int64     `json:"id,omitempty" bson:"id"`
		Name       string    `json:"name,omitempty" bson:"name"`
		Marks      []int64   `json:"-,omitempty" bson:"marks"`
		CreatedAt  time.Time `json:"createdAt,omitempty" bson:"createdAt"`
		TotalMarks int64     `json:"totalMarks,omitempty" bson:"totalMarks"`
	}
	FetchRecordRequest struct {
		StartDate string `json:"startDate" validate:""`
		EndDate   string `json:"endDate" validate:""`
		MinCount  int64  `json:"minCount" validate:"number"`
		MaxCount  int64  `json:"maxCount" validate:"number"`
	}
)

type (
	RecordUsecase interface {
		GetAll(ctx context.Context, req FetchRecordRequest) (response []Record, err error)
	}

	RecordRepository interface {
		GetAll(ctx context.Context, req FetchRecordRequest) (response []Record, err error)
	}
)
