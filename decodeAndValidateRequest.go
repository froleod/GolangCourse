package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CreateUserRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

var (
	ErrEmailRequired                = errors.New("email is required")
	ErrPasswordRequired             = errors.New("password is required")
	ErrPasswordConfirmationRequired = errors.New("password confirmation is required")
	ErrPasswordDoesNotMatch         = errors.New("password does not match with the confirmation")
)

// DecodeAndValidateRequest BEGIN (write your solution here)
func DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error) {
	req := CreateUserRequest{}
	json.Unmarshal(requestBody, &req)
	switch {
	case req.Email == "":
		return CreateUserRequest{}, ErrEmailRequired
	case req.Password == "":
		return CreateUserRequest{}, ErrPasswordRequired
	case req.PasswordConfirmation == "":
		return CreateUserRequest{}, ErrPasswordConfirmationRequired
	case req.Password != req.PasswordConfirmation:
		return CreateUserRequest{}, ErrPasswordDoesNotMatch
	default:
		return req, nil
	}

}

func main() {
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"\",\"password\":\"test\",\"password_confirmation\":\"test\"}")))
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"\",\"password_confirmation\":\"test\"}")))
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"email@test.com\",\"password\":\"passwordtest\",\"password_confirmation\":\"passwordtest\"}")))
}

// END
