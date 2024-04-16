package main

import (
	"errors"
	"fmt"
)

type NonCriticalError struct{}

func (e NonCriticalError) Error() string {
	return "validation error"
}

var (
	ErrBadConnection = errors.New("bad connection")
	ErrBadRequest    = errors.New("bad request")
)

const UnknownErrorMsg = "unknown error"

// Реализуйте функцию GetErrorMsg(err error) string, которая возвращает текст ошибки,
//если она критичная. Если ошибка некритичная, то возвращается пустая строка.
//В случае неизвестной ошибки возвращается строка unknown error:

// BEGIN (write your solution here)
func GetErrorMsg(err error) string {
	switch {
	case errors.Is(err, ErrBadConnection):
		return ErrBadConnection.Error()
	case errors.Is(err, ErrBadRequest):
		return ErrBadRequest.Error()
	case errors.As(err, &NonCriticalError{}):
		return ""
	default:
		return UnknownErrorMsg
	}

}

func main() {
	fmt.Println(GetErrorMsg(ErrBadConnection))
	fmt.Println(GetErrorMsg(ErrBadRequest))
	fmt.Println(GetErrorMsg(NonCriticalError{}))
	fmt.Println(GetErrorMsg(errors.New("random error")))
}

// END
