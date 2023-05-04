package entities

type UserCart struct {
	UserID int  `json:"userId"`
	Cake   Cake `json:"cake"`
}

func NewUserCart() UserCart {
	return UserCart{Cake: Cake{}}
}

type Cake struct {
	CakeID int     `json:"cake_id"`
	Title  string  `json:"title"`
	Price  float32 `json:"price"`
	ImgUrl string  `json:"imgUrl"`
	Value  int     `json:"value"`
}
