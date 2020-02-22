package domain

import "errors"

var (
	ErrUserWithEmailAlreadyExist = errors.New("User With Email Already Exists")
	ErrUserWithUsernameAlreadyExist = errors.New("User With Username Already Exists")
)