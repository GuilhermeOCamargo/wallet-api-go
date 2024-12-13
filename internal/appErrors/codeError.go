package apperrors

import "fmt"

type CodeError struct {
	Status int
	Msg    string
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ERRO: '%d': %s", e.Status, e.Msg)
}

func NewCodeError(status int, msg string) error {
	return &CodeError{
		Status: status,
		Msg:    msg,
	}
}
