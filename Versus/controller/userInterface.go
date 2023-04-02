package controller

import (
	"Versus/config"
	"Versus/models"
	"Versus/storage"
	"fmt"
)

type Userman models.UserMan
type ProductMan models.ProductMan
type CardManager models.CardManager
type StorageI storage.StorageI

var GlobalUserId string

type Task2 models.Task2
type Task2Manager models.Task2Manager

type Tas7 models.Tas7

var IdUser string

func WhoAreYou() {

	var namer string
	fmt.Print("Write your name: ")
	fmt.Scanln(&namer)
	configurator := config.Load()

	users := storage.UsReader(configurator.UserJson)

	nextStep := false
	username := ""

	for _, v := range users {
		if v.Name == namer {
			nextStep = true
			username = v.Name + " " + v.Surname
			IdUser = v.Id
			GlobalUserId = v.Id

		}

	}

	if nextStep {
		fmt.Println("|| Assalomu alekum ", username, " ||")
		ViewProduct()
	} else {
		fmt.Println(namer, " isimli odam yoq")
	}

}

func ViewProduct() {

	prods := storage.ProdReader(config.Load().ProductJson)
	for _, v := range prods {
		fmt.Println(v.Name, v.Price)
	}

	WBuy()

	var intrf StorageI
	intrf = &CardManager{}
	fmt.Print("Choise Tasks 1,2,3,4,5,6,7,8,9,10 -> ")
	var choise int
	fmt.Scanln(&choise)
	if choise == 1 {
		fmt.Print("Task 1 ->")

		intrf.SortingShopCarts()

		fmt.Println("Succesfully Sorted")

	} else if choise == 2 {
		fmt.Println("Task 2 ->")

		for _, v := range intrf.ViewHistoryUser() {
			fmt.Println("Name:", v.ProductName, "Price:", v.ProductPrice, "Count:", v.ProductCount, "Total:", v.TotalSumm, "Time:", v.Time)

		}

	} else if choise == 3 {
		fmt.Println("Task 3 ->")

		fmt.Println(" Obshiy ", intrf.FullCash(), "sum ishlatilgan")
	} else if choise == 4 {
		fmt.Print("Task 4 ->")
		var proNam string
		fmt.Println("Tell name product :")
		fmt.Scanln(&proNam)
		fmt.Println(intrf.HowMuchSaledProduct(proNam))

	} else if choise == 5 {
		fmt.Println("Task 5 ->")
		for _, v := range intrf.Top10BestProducts() {
			fmt.Println(v)

		}

	} else if choise == 6 {
		fmt.Println("Task 6 ->")
		for _, v := range intrf.Top10WorthProducts() {
			fmt.Println(v)

		}
	} else if choise == 7 {
		fmt.Println("Task 7 ->")
		for _, v := range intrf.TableBestProducts() {
			fmt.Println("Name: ", v.Name, "Sana: ", v.Date, "Count: ", v.Count)
		}
	} else if choise == 8 {
		fmt.Println("Task 8 ->")
		for _, v := range intrf.CategoryCountsProducts() {
			fmt.Println(v)

		}
	} else if choise == 9 {
		fmt.Println("Task 9 ->")
		fmt.Println(intrf.VeryActiveUser(), "dokonda bogan 'activlik qilgan'")
	} else if choise == 10 {
		var caser int
		fmt.Println("which case you want 1 or 2")
		fmt.Scanln(&caser)
		if caser == 1 {
			fmt.Println("Task 10 case 1 ->")
			intrf.Task10Case1()

		} else if caser == 2 {

			fmt.Println("Task 10 case 2 ->")
			intrf.Task10Case2()
		}
	}

}
