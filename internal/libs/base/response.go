package base

type BaseResponse struct {
	Data any `json:"data"`
}

type ErrorResponse struct {
	Data string `json:"data" example:"error message"`
}

func NewResponse(data any) *BaseResponse {
	return &BaseResponse{Data: data}
}
