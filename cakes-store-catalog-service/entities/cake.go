package entities

type Cake struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
	Description string  `json:"description"`
	BiscuitType string  `json:"biscuitType"`
	CreamType   string  `json:"creamType"`
	ToppingType string  `json:"toppingType"`
	FillingType string  `json:"fillingType"`
	Berries     string  `json:"berries"`
	Weight      string  `json:"weight"`
	IsCustom    bool    `json:"isCustom"`
	CustomText  string  `json:"customText"`
}
