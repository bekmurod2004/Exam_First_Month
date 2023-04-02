package controller

import (
	"Versus/config"
	"Versus/storage"
	"fmt"
	"sort"
)

func (c *CardManager) SortingShopCarts() []storage.ShopCardMan {
	ans := storage.SuperSort(config.Load().ShopCards)

	return ans
}

func (c *CardManager) ViewHistoryUser() []storage.Task2 {
	a := storage.SuperHistory(config.Load().ProductJson, config.Load().ShopCards, GlobalUserId)

	return a
}

func (c *CardManager) FullCash() int {
	a := storage.TotalSummer(config.Load().ProductJson, config.Load().ShopCards, GlobalUserId)

	return a
}

func (c *CardManager) HowMuchSaledProduct(proNam string) int {
	ans := storage.FinderCount(config.Load().ProductJson, config.Load().ShopCards, proNam)

	return ans

}

func (c *CardManager) Top10BestProducts() []string {

	ans := storage.TopTen(config.Load().ProductJson, config.Load().ShopCards)

	top := []string{}
	for i := len(ans) - 1; i > 2; i-- {
		top = append(top, ans[i])

	}

	return top
}

func (c *CardManager) Top10WorthProducts() []string {

	ans := storage.TopTen(config.Load().ProductJson, config.Load().ShopCards)

	worth := []string{}
	for i := 0; i < len(ans)-3; i++ {
		worth = append(worth, ans[i])
	}

	return worth
}

func (c *CardManager) TableBestProducts() []storage.Tas7 {

	ans := storage.WhenBetter(config.Load().ProductJson, config.Load().ShopCards)

	return ans
}

func (c *CardManager) CategoryCountsProducts() []string {
	ans := storage.CatMad(config.Load().CategoryJson, config.Load().ShopCards, config.Load().ProductJson)

	return ans
}

func (c *CardManager) VeryActiveUser() []string {
	ans := storage.FindSuperUser(config.Load().ProductJson, config.Load().ShopCards, config.Load().UserJson)

	return ans
}

func (c *CardManager) Task10Case1() {
	fmt.Println("-------------------------------------------------------------------------------------")
	prods := storage.ProdReader(config.Load().ProductJson)
	for _, v := range prods {
		fmt.Println(v.Name, v.Price)
	}

	act := true
	var productSh string
	var hoM int
	choosen := []string{}
	for act {
		fmt.Print("product name: ")
		fmt.Scanln(&productSh)
		fmt.Print("how much product: ")
		fmt.Scanln(&hoM)

		for i := 0; i < hoM; i++ {
			if productSh != "stop" {
				choosen = append(choosen, productSh)
			}

		}
		if productSh == "stop" {
			act = false
		}

	}

	pricerro := []int{}

	for _, v := range choosen {

		for _, k := range prods {
			if v == k.Name {
				pricerro = append(pricerro, k.Price)

			}

		}

	}

	if len(pricerro) >= 9 {
		summ := 0
		for i := 0; i < 9; i++ {
			summ += pricerro[i]

		}

		fmt.Println(summ, "tekinga:", choosen[0])

	} else {
		summ := 0
		for i := 0; i < len(pricerro); i++ {
			summ += pricerro[i]

		}
		fmt.Println("9 tadan kam boginuchun", summ)

	}

}

func (c *CardManager) Task10Case2() {
	fmt.Println("-------------------------------------------------------------------------------------")
	prods := storage.ProdReader(config.Load().ProductJson)
	for _, v := range prods {
		fmt.Println(v.Name, v.Price)
	}

	act := true
	var productSh string
	var hoM int
	choosen := []string{}
	for act {
		fmt.Print("product name: ")
		fmt.Scanln(&productSh)
		fmt.Print("how much product: ")
		fmt.Scanln(&hoM)

		for i := 0; i < hoM; i++ {
			if productSh != "stop" {
				choosen = append(choosen, productSh)
			}

		}
		if productSh == "stop" {
			act = false
		}

	}

	pricerro := []int{}

	for _, v := range choosen {

		for _, k := range prods {
			if v == k.Name {
				pricerro = append(pricerro, k.Price)

			}

		}

	}

	if len(pricerro) >= 9 {
		sort.Sort(sort.IntSlice(pricerro))
		summ := 0
		for i := 1; i < len(pricerro); i++ {
			summ += pricerro[i]
		}
		free := pricerro[0]

		fmt.Println(summ, "eng arzon narsani puli optashaldi:", free)

	} else {
		summ := 0
		for i := 0; i < len(pricerro); i++ {
			summ += pricerro[i]
		}

		fmt.Println("9 tadan kam product narhi: ", summ)

	}

}
