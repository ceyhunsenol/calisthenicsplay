package model

type ServiceError struct {
	Code    int
	Message string
}

func (h *ServiceError) Error() string {
	return h.Message
}
