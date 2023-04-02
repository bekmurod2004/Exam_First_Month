package storage

import (
	"Versus/models"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Userman models.UserMan
type ProductMan models.ProductMan
type ShopCardMan models.ShopCardMan
type Task2 models.Task2
type Task2Manager models.Task2Manager
type Tas7 models.Tas7
type CategoryMan models.CategoryMan

func UsReader(jsonPath string) []Userman {
	var a []Userman

	date, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	err = json.Unmarshal(date, &a)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	return a

}

func ProdReader(jsonPath string) []ProductMan {
	var a []ProductMan

	date, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	err = json.Unmarshal(date, &a)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	return a

}

func SuperSort(jsonPath string) []ShopCardMan {
	// var answer []ShopCardMan

	date, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	var vw []ShopCardMan
	err = json.Unmarshal(date, &vw)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	for i := 0; i < len(vw); i++ {
		swaper := ""
		for j := i + 1; j < len(vw); j++ {
			if vw[i].Time > vw[j].Time {
				swaper = vw[j].Time
				vw[j].Time = vw[i].Time
				vw[i].Time = swaper
			}

		}

	}

	date, err = json.Marshal(vw)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = os.WriteFile(jsonPath, date, 0644)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil

}

func WriteNewShopCatd(jsonPath string, a ShopCardMan) error {

	date, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	var wasCard []ShopCardMan
	err = json.Unmarshal(date, &wasCard)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	wasCard = append(wasCard, a)

	date, err = json.Marshal(wasCard)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = os.WriteFile(jsonPath, date, 0644)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil

}

func SuperHistory(prode, shopcards, UserIdi string) []Task2 {
	// Read prode
	var Rprode []ProductMan
	date, err := os.ReadFile(prode)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	err = json.Unmarshal(date, &Rprode)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	// Read shopcards
	date, err = os.ReadFile(shopcards)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	var RshoCards []ShopCardMan
	err = json.Unmarshal(date, &RshoCards)
	if err != nil {
		fmt.Println(nil)
		return nil
	}
	var test []Task2

	var martId string

	for _, v := range RshoCards {
		var gok Task2
		if v.UserId == UserIdi {
			gok.Time = v.Time
			gok.ProductCount = v.Count
			martId = v.ProductId

			for _, k := range Rprode {
				if martId == k.Id {
					gok.ProductName = k.Name
					gok.ProductPrice = k.Price

				}

			}

			summ := 0

			for i := 0; i < gok.ProductCount; i++ {
				summ += gok.ProductPrice
			}
			gok.TotalSumm = summ

			test = append(test, gok)

		}
	}

	// fmt.Println(test)

	return test
}

func TotalSummer(prode, shopcards, idi string) int {
	ans := SuperHistory(prode, shopcards, idi)
	totalis := 0
	for _, v := range ans {
		totalis += v.TotalSumm
	}

	return totalis
}

func FinderCount(prode, shopcards, name string) int {
	// Read prode
	var Rprode []ProductMan
	date, err := os.ReadFile(prode)
	if err != nil {
		fmt.Println(nil)
		return 0
	}

	err = json.Unmarshal(date, &Rprode)
	if err != nil {
		fmt.Println(nil)
		return 0
	}

	var idiProd string
	for _, v := range Rprode {
		if v.Name == name {
			idiProd = v.Id
		}

	}

	// Read shopcards
	date, err = os.ReadFile(shopcards)
	if err != nil {
		fmt.Println(nil)
		return 0
	}

	var RshoCards []ShopCardMan
	err = json.Unmarshal(date, &RshoCards)
	if err != nil {
		fmt.Println(nil)
		return 0
	}

	count := 0
	for _, v := range RshoCards {
		if v.ProductId == idiProd {
			count += v.Count

		}

	}

	return count
}

func TopTen(prode, shopcards string) []string {
	// Read prode
	var Rprode []ProductMan
	date, err := os.ReadFile(prode)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	err = json.Unmarshal(date, &Rprode)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	// Read shopcards
	date, err = os.ReadFile(shopcards)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	var RshoCards []ShopCardMan
	err = json.Unmarshal(date, &RshoCards)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	mapa := make(map[string]string)
	for _, v := range Rprode {
		mapa[v.Name] = v.Id

	}

	mapas := make(map[string]int)
	for i, v := range mapa {
		count := 0
		for _, k := range RshoCards {
			if v == k.ProductId {
				count += k.Count
				mapas[i] = count

			} else {
				mapas[i] = count
			}

		}

	}

	/// Sort mapas/
	////////////////////////////////////////////
	keys := make([]string, 0, len(mapas))

	for key := range mapas {
		keys = append(keys, key)

	}

	sort.SliceStable(keys,
		func(i, j int) bool {
			return mapas[keys[i]] < mapas[keys[j]]
		})

	

	answer := []string{}
	///////////////////////////////////////////////////////////////////////////////

	for i := 0; i < len(keys); i++ {
		k := keys[i]
		v := strconv.Itoa(mapas[keys[i]])
		lister := " Name: " + k + " Count: " + v

		answer = append(answer, lister)

	}
	///////////////////////////////////////////////////////////////////////////////

	// fmt.Println(answer)

	return answer
}

func WhenBetter(prode, shopcards string) []Tas7 {
	var Rprode []ProductMan
	date, err := os.ReadFile(prode)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	err = json.Unmarshal(date, &Rprode)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	date, err = os.ReadFile(shopcards)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	var RshoCards []ShopCardMan
	err = json.Unmarshal(date, &RshoCards)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	var arr []Tas7
	for _, v := range RshoCards {
		a := Tas7{}
		a.Date = v.Time
		a.Count = v.Count

		for _, k := range Rprode {
			if v.ProductId == k.Id {

				a.Name = k.Name

			}

		}

		arr = append(arr, a)

	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Count > arr[j].Count
	})

	for _, v := range arr {
		fmt.Println(v)

	}

	return arr
}

func CatMad(category, shopcards, prode string) []string {
	var cats []CategoryMan
	date, err := os.ReadFile(category)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	err = json.Unmarshal(date, &cats)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	/////////////////////////////////////////////////

	var Rprode []ProductMan
	date, err = os.ReadFile(prode)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	err = json.Unmarshal(date, &Rprode)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	///////////////////////////////////////////////////////////
	date, err = os.ReadFile(shopcards)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	var RshoCards []ShopCardMan
	err = json.Unmarshal(date, &RshoCards)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	mapa := make(map[string][]string)

	for _, v := range cats {
		subar := []string{}
		for _, j := range Rprode {
			if v.Id == j.Category {
				subar = append(subar, j.Id)
			}

		}

		mapa[v.Name] = subar

	}

	mass := make(map[string]int)
	for i, v := range mapa {
		countCounters := 0
		for _, l := range v {
			for _, j := range RshoCards {
				if l == j.ProductId {
					countCounters += j.Count
				}
				
			}
			
		}

		mass[i] =countCounters
		
	}

	answer :=[]string{}
	for i, v := range mass {
		f := strconv.Itoa(v)
		liter := i +": "+ f
		answer = append(answer, liter)
	}

	return answer

}

func FindSuperUser(prode,shopcards, suser string) []string {

	var SupUs []Userman
	date, err := os.ReadFile(suser)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	err = json.Unmarshal(date, &SupUs)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	// Read shopcards
	date, err = os.ReadFile(shopcards)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	var RshoCards []ShopCardMan
	err = json.Unmarshal(date, &RshoCards)
	if err != nil {
		fmt.Println(nil)
		return nil
	}

	mapa := make(map[string]int)

	for _, v := range SupUs {
		cuter :=0
		for _, g := range RshoCards {
			if v.Id == g.UserId {
				cuter++
				
			}
			
		}
		mapa[v.Name] = cuter
		
	}

/////////////////////////////////Sorting/////////
	keys := make([]string, 0, len(mapa))

	for key := range mapa {
		keys = append(keys, key)

	}

	sort.SliceStable(keys,
		func(i, j int) bool {
			return mapa[keys[i]] > mapa[keys[j]]
		})
/////////////////////////////////////////////////////////////
topUser := []string{}
for _, v := range keys {
	word := strconv.Itoa(mapa[v])
	topUser = append(topUser,v, word  )
	break
	
}
	

	


	return topUser
}
