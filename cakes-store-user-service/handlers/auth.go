package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/infinitss13/Cakes-store/entities"
	"github.com/infinitss13/Cakes-store/services"
	"github.com/sirupsen/logrus"
)

// @Summary SignUp
// @Tags auth
// @Description handler for SignUp request, allows user to register in service
// @ID signup
// @Param input body entities.UserPersonalData true "account info"
// @Accept json
// @Produce json
// @Success 200 {object} docs.SuccessSignUp
// @Failure 400 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /auth/sign-up [post]
func (handlers AuthHandlers) signUp(context *gin.Context) {
	userData := entities.NewUserPersonalData()
	err := context.ShouldBindJSON(&userData)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logrus.Error(err)
		return
	}
	err = handlers.ServiceAuth.CreateUser(userData)
	if err != nil {
		if err == services.ErrUserExists {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			logrus.Error(err)
			return
		} else {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			logrus.Error(err)
			return
		}

	}
	context.JSON(http.StatusOK, gin.H{
		"message": "user is successfully created",
	})
}

// @Summary SignIn
// @Tags auth
// @Description handler for SignIn request, allows user to authenticate
// @ID signIn
// @Param input body entities.UserSignInData true "signInRequest info"
// @Accept json
// @Produce json
// @Success 200 {object} docs.SuccessSignIn
// @Failure 400 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /auth/sign-in [post]
func (handlers AuthHandlers) signIn(c *gin.Context) {
	signInReq := entities.NewUserSignInData()
	if err := c.ShouldBindJSON(&signInReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logrus.Info("PASSWORD : ", signInReq.Password)
	err := handlers.ServiceAuth.CheckUser(signInReq)
	if err != nil {
		httpStatusCode, err := validateError(err)
		logrus.Error(err)
		c.JSON(httpStatusCode, gin.H{
			"error": err.Error(),
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone": signInReq.PhoneNumber,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func validateError(err error) (int, error) {
	if err == sql.ErrNoRows {
		return http.StatusBadRequest, errors.New("user with this data isn't found")
	}
	if err == services.ErrInvalidPhone || err == services.ErrPasswordEmpty || err == services.ErrPasswordIncorrect {
		return http.StatusBadRequest, err
	}
	return http.StatusInternalServerError, err
}
