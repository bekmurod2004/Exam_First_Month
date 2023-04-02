package models

type UserMan struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Balance int    `json:"balance"`
}

type ProductMan struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category_id"`
}

type ShopCardMan struct {
	Id        string `json:"id"`
	ProductId string `json:"productId"`
	UserId    string `json:"userId"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Time      string `json:"time"`
}

type CategoryMan struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	ParenId string `json:"paren_id"`
}

type CardManager struct {
}

////////////////////////////////////////////////////////////////////////////

type Task2 struct {
	ProductName  string
	ProductPrice int
	ProductCount int
	TotalSumm    int
	Time         string
}

type Task2Manager struct {
}

type Tas7 struct {
	Name  string
	Date  string
	Count int
}
