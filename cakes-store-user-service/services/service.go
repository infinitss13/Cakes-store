package services

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/infinitss13/Cakes-store/entities"
	"github.com/infinitss13/Cakes-store/repository"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Database *repository.Database
}

func NewService(database *repository.Database) Service {
	return Service{Database: database}
}

type ServiceAuth interface {
	CreateUser(userData *entities.UserPersonalData) error
	CheckUser(signInData entities.UserSignInData) (int, error)
}

func (srv Service) CreateUser(userData *entities.UserPersonalData) error {
	isExist, err := srv.Database.IsUserExist(userData.Email, userData.PhoneNumber)
	if isExist {
		return ErrUserExists
	}
	if err != nil {
		return err
	}
	fmt.Println(isExist)
	err = isUserDataValid(userData)
	hashPassword, err := HashPassword(userData.Password)
	if err != nil {
		return fmt.Errorf("error in 'service.HashPassword' : %w", err)
	}
	userData.Password = hashPassword
	err = srv.Database.CreateUser(userData)
	if err != nil {
		return fmt.Errorf("error in database.CreateUser : %w", err)
	}
	return nil
}

func (srv Service) CheckUser(signInData entities.UserSignInData) (int, error) {
	if err := isRequestedUserValid(signInData); err != nil {
		return 0, err
	}
	passwordHash, err := srv.Database.GetPasswordByNumber(signInData.PhoneNumber)
	if err != nil {
		return 0, err
	}
	err = isPasswordCorrect(signInData.Password, passwordHash)
	if err != nil {
		return 0, err
	}
	id, err := srv.Database.GetUserIDByNumber(signInData.PhoneNumber)
	if err != nil {
		return 0, errors.New("trouble with getting id")
	}
	fmt.Println("ID in service user service : ", id)
	return id, nil
}

func isUserDataValid(userData *entities.UserPersonalData) error {
	if ok := isValidPhoneNumber(userData.PhoneNumber); !ok {
		return ErrInvalidPhone
	}
	if ok := isValidEmail(userData.Email); !ok {
		return ErrInvalidEmail
	}
	if ok := isValidRole(entities.Role(userData.Role)); !ok {
		return ErrRoleInvalid
	}
	return nil
}

func isRequestedUserValid(signInReq entities.UserSignInData) error {
	if ok := isValidPhoneNumber(signInReq.PhoneNumber); !ok {
		return ErrInvalidPhone
	}
	if signInReq.Password == "" {
		return ErrPasswordEmpty
	}
	return nil
}

func isPasswordCorrect(password, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return ErrPasswordIncorrect
		} else {
			return err
		}
	}
	return nil
}

func HashPassword(password string) (string, error) {
	saltedBytes := []byte(password)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashPassword := string(hashedBytes)
	return hashPassword, nil
}
func isValidEmail(email string) bool {
	// Define the regular expression pattern for a valid email address
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression pattern
	regex := regexp.MustCompile(pattern)

	// Use the regular expression to test if the email address matches the pattern
	return regex.MatchString(email)
}

func isValidPhoneNumber(phoneNumber string) bool {
	// Define a regular expression for phone numbers in the format "+375123456789"
	regex := regexp.MustCompile(`^\+375\d{9}$`)

	// Check if the phone number matches the regular expression
	return regex.MatchString(phoneNumber)
}

func isValidRole(role entities.Role) bool {
	if role != entities.UserRole && role != entities.AuthorRole && role != entities.AdminRole {
		return false
	} else {
		return true
	}
}
