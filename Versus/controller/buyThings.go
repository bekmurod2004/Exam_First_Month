package controller

import (
	"Versus/config"
	"Versus/models"
	"Versus/storage"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ShopCardMan models.ShopCardMan

func WBuy() {
	unq := uuid.New()
	var nameP string
	var countProd int
	fmt.Print("nma sotvomoqchisiz -> ")
	fmt.Scanln(&nameP)
	fmt.Print("qancha omoqchisiz -> ")
	fmt.Scanln(&countProd)

	var newShopCatd ShopCardMan

	prods := storage.ProdReader(config.Load().ProductJson)

	// product Id
	for _, v := range prods {
		if v.Name == nameP {
			newShopCatd.ProductId = v.Id
		}
	}

	// user Id
	newShopCatd.UserId = IdUser

	// id
	newShopCatd.Id = unq.String()

	// count products
	newShopCatd.Count = countProd

	// status
	newShopCatd.Status = true

	// time
	t := time.Now()
	newShopCatd.Time = t.Format("2006.01.02 15:04:05")

	storage.WriteNewShopCatd(config.Load().ShopCards, storage.ShopCardMan(newShopCatd))

}
