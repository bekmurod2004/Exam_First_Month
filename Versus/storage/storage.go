package storage 


type StorageI interface{
	SortingShopCarts() []ShopCardMan   // 1
	ViewHistoryUser() []Task2  // 2
	FullCash() int // 3
	HowMuchSaledProduct(productName string) int// 4
	Top10BestProducts() []string // 5
	Top10WorthProducts() []string // 6
	TableBestProducts() []Tas7// 7
	CategoryCountsProducts() []string // 8
	VeryActiveUser() []string // 9
	Task10Case1() // 10
	Task10Case2() // 10

}