package services

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/infinitss13/Cakes-store-card-service/config"
	"github.com/infinitss13/Cakes-store-card-service/entities"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	CreateCard(userId int, cake entities.Cake) error
	AddToCart(userID int, Cake entities.Cake) error
	IsCartEmpty(userID int) (bool, error)
	GetCartItems(userId int) (map[string]interface{}, error)
}
type Service struct {
	Database Repository
}

func NewService(Database Repository) Service {
	return Service{Database: Database}
}

func (s Service) AddItemToCart(userId int, cake entities.Cake) error {
	//var cakeCart json.RawMessage
	//cakeCart, err := json.Marshal(cart.Cake)
	//if err != nil {
	//	return err
	//}

	isCardEmpty, err := s.Database.IsCartEmpty(userId)
	if err != nil {
		return err
	}
	if isCardEmpty {
		err = s.Database.CreateCard(userId, cake)
		if err != nil {
			return err
		}
		return nil
	}
	err = s.Database.AddToCart(userId, cake)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) GetCart(userId int) (map[string]interface{}, error) {
	cakes, err := s.Database.GetCartItems(userId)
	if err != nil {
		return nil, err
	}
	return cakes, nil
}

func (s Service) Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		splitToken, err := s.GetToken(context)
		if err != nil {
			logrus.Error(err)
			context.JSON(http.StatusInternalServerError, err)
		}
		fmt.Println(splitToken)
		id, err := s.VerifyToken(splitToken)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		context.Set("userID", id)
		context.Next()
	}
}

type JWTClaim struct {
	id string `json:"userID"`
	jwt.StandardClaims
}

func (s Service) VerifyToken(tokenSigned string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenSigned, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv("ACCESS_KEY", "")), nil
	})
	if err != nil {
		return "", errors.New("error parsing claims")
	}
	claims := token.Claims.(*JWTClaim)
	fmt.Println("CLAIMS")
	fmt.Println(claims.id)
	fmt.Println(claims.ExpiresAt)
	fmt.Println(claims.StandardClaims.ExpiresAt)
	if claims.StandardClaims.ExpiresAt < time.Now().Local().Unix() {
		return "", errors.New("token expired")
	}

	return claims.id, nil
}

func (s Service) GetToken(context *gin.Context) (string, error) {
	tokenString := context.GetHeader("Authorization")
	if tokenString == "" {
		return "", errors.New("no access token")
	}
	splitedToken := strings.Split(tokenString, " ")
	return splitedToken[1], nil
}
