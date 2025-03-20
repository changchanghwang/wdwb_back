package base

type BaseResponse struct {
	Data any `json:"data"`
}

func NewResponse(data any) *BaseResponse {
	return &BaseResponse{Data: data}
}
