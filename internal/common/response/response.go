package response

type Response[T any] struct {
	Status      string            `json:"status"`
	Data        *T                `json:"data,omitempty"`
	Message     string            `json:"message,omitempty"`
	Validations map[string]string `json:"validations,omitempty"`
}

type M map[string]any

func (r *Response[T]) ResponseFail(validations map[string]string) Response[T] {
	r.Status = StatusFail
	r.Validations = validations
	r.Data = nil
	r.Message = ""

	return *r
}

func (r *Response[T]) ResponseFailMessage(message string) Response[T] {
	r.Status = StatusFail
	r.Message = message
	r.Validations = nil
	r.Data = nil

	return *r
}

func (r *Response[T]) ResponseError(message string) Response[T] {

	r.Status = StatusError
	r.Message = message
	r.Validations = nil
	r.Data = nil

	return *r
}

func (r *Response[T]) ResponseSuccess(data T) Response[T] {
	r.Status = StatusSuccess
	r.Data = &data
	r.Message = ""
	r.Validations = nil

	return *r
}

func (r *Response[T]) ResponseNotFound() Response[T] {
	r.Status = StatusNotFound
	r.Message = ""
	r.Data = nil
	r.Validations = nil

	return *r
}

func (r *Response[T]) ResponseUnauthorized(message string) Response[T] {
	r.Status = StatusUnauthorized
	r.Message = message
	r.Validations = nil
	r.Data = nil

	return *r
}
