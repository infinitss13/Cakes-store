package entities

type UserPersonalData struct {
	ID          int
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Password    string `json:"password" binding:"required"`
	DateOfBirth string `json:"dateOfBirth"`
	Role        int    `json:"role"`
	ImgUrl      string `json:"imgUrl"`
}

func NewUserPersonalData() *UserPersonalData {
	return &UserPersonalData{}
}

type UserSignInData struct {
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

func NewUserSignInData() UserSignInData {
	return UserSignInData{}
}
