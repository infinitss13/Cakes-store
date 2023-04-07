package services

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/infinitss13/Cakes-store-catalog-service/config"
	"github.com/infinitss13/Cakes-store-catalog-service/entities"
	"github.com/sirupsen/logrus"
)

type ServiceCatalog struct {
	Database DatabaseCatalog
}

func NewServiceCatalog(database DatabaseCatalog) *ServiceCatalog {
	return &ServiceCatalog{Database: database}
}

type DatabaseCatalog interface {
	GetCatalog() ([]entities.Cake, error)
	GetCatalogWithLimit(limit int) ([]entities.Cake, error)
}

func (s ServiceCatalog) GetCatalog(limit string) (cakes []entities.Cake, err error) {
	var limitNum int

	if limit != "" {
		limitNum, err = strconv.Atoi(limit)
		if err != nil {
			return nil, err
		}
		cakes, err = s.Database.GetCatalogWithLimit(limitNum)
	} else {
		cakes, err = s.Database.GetCatalog()
	}
	if err != nil {
		return nil, err
	}
	if cakes == nil {
		return nil, errors.New("error : Catalog is empty, pinganite Stasa")
	}
	return cakes, nil
}

func (s ServiceCatalog) Auth() gin.HandlerFunc {
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

func (s ServiceCatalog) VerifyToken(tokenSigned string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenSigned, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv("ACCESS_KEY", "")), nil
	})
	if err != nil {
		return "", errors.New("Error parsing claims")
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

func (s ServiceCatalog) GetToken(context *gin.Context) (string, error) {
	tokenString := context.GetHeader("Authorization")
	if tokenString == "" {
		return "", errors.New("no access token")
	}
	splitedToken := strings.Split(tokenString, " ")
	return splitedToken[1], nil
}
