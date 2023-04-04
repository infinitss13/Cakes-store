package services

import "errors"

var ErrInvalidPhone = errors.New("phone number is in invalid format. Phone should consist of +375, then 2 digits of operator and 7 digits of number")
var ErrInvalidEmail = errors.New("email is in invalid format")
var ErrRoleInvalid = errors.New("role is invalid. It can be 1 - user, 2 - author, 3 - admin")
var ErrUserExists = errors.New("user with this email or phone number already exists")
var ErrPasswordEmpty = errors.New("password can't be empty")
var ErrPasswordIncorrect = errors.New("password isn't correct")
