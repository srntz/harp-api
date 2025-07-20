package response

type Success struct {
	Data any `json:"data"`
}

func NewSuccess(data any) Success {
	return Success{Data: data}
}
