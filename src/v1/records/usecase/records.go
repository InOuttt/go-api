package usecase

import (
	"context"
	"time"

	"github.com/inouttt/test-go-mezink/src/v1/domain"
)

type recordUsecase struct {
	recordRepo     domain.RecordRepository
	contextTimeout time.Duration
}

func NewRecordUsecase(a domain.RecordRepository, timeout time.Duration) domain.RecordUsecase {
	return &recordUsecase{
		recordRepo:     a,
		contextTimeout: timeout,
	}
}

func (u *recordUsecase) GetAll(ctx context.Context, req domain.FetchRecordRequest) (response []domain.Record, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if req.EndDate == "" {
		req.EndDate = time.Now().Format("2006-01-02")
	}

	response, err = u.recordRepo.GetAll(ctx, req)
	if err != nil {
		return nil, domain.ErrBadParamInput
	}

	return
}
