package domain

type (
	Pagination struct {
		Offset int64 `query:"offset"`
		Limit  int64 `query:"limit"`
	}

	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"msg"`
		Data    interface{} `json:"records"`
	}
)
