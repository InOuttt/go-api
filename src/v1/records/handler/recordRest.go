package handler

import (
	"net/http"

	"github.com/inouttt/test-go-mezink/src/v1/domain"
	"github.com/labstack/echo/v4"
)

type httpOrderHandler struct {
	Usecase domain.RecordUsecase
}

func NewHttpOrderHandler(e *echo.Echo, us domain.RecordUsecase) {
	handler := &httpOrderHandler{
		Usecase: us,
	}

	orderAuth := e.Group("/v1/records")
	orderAuth.GET("", handler.Get)
}

func (h *httpOrderHandler) Get(c echo.Context) (err error) {
	var request domain.FetchRecordRequest
	err = c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.GenerateResponseError(domain.ErrBadRequest, "wrong datatype", err))
	}

	if err = c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, domain.GenerateResponseError(domain.ErrBadRequest, "validation failed", err))
	}

	resp, err := h.Usecase.GetAll(c.Request().Context(), request)
	if err != nil {
		return c.JSON(domain.GetStatusCode(err), domain.GenerateResponseError(err, "", nil))
	}

	return c.JSON(http.StatusOK, domain.GenerateReponseSuccess(resp))
}
