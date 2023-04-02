package config

type Conf struct {
	UserJson string 
	ProductJson string
	ShopCards string
	CategoryJson string
}


func Load() Conf {
	PathConf := Conf{
		UserJson: "../data/user.json",
		ProductJson: "../data/product.json",
		ShopCards: "../data/shop_cart.json",
		CategoryJson: "../data/category.json",
		
	}

	return PathConf
}