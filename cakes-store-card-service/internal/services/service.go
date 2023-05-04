package services

import (
	"github.com/infinitss13/Cakes-store-card-service/entities"
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

func (s Service) AddItemToCart(cart entities.UserCart) error {
	//var cakeCart json.RawMessage
	//cakeCart, err := json.Marshal(cart.Cake)
	//if err != nil {
	//	return err
	//}

	isCardEmpty, err := s.Database.IsCartEmpty(cart.UserID)
	if err != nil {
		return err
	}
	if isCardEmpty {
		err = s.Database.CreateCard(cart.UserID, cart.Cake)
		if err != nil {
			return err
		}
		return nil
	}
	err = s.Database.AddToCart(cart.UserID, cart.Cake)
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
