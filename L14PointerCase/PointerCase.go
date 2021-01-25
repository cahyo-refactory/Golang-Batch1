package main

import (
	"fmt"
	"strconv"
)

//total belanjaan
//item | quantitiy | harga satuan

//perbedaan array dan slice
//array itu length fix | slice length growth
//slice itu bisa referense ke array | slice bisa dibuat tanpa ada array terlebih dahulu

//input belanjaan

type belanjaan struct {
	item        string
	Quantity    int
	hargaSatuan float64
}

type belanjaanKumpulam struct {
	listBelanjaan []belanjaan
}

func main() {
	var listBelanjaan []belanjaan
	var listBelanjaanMethod belanjaanKumpulam
	var itemBelanjaan belanjaan
	var item string
	var quantity string
	var hargaSatuan string
	var totalPotongan float64

	fmt.Println("Masukkan nama Item")
	fmt.Scanln(&item)
	fmt.Println("Masukkan jumlah atau quantity")
	fmt.Scanln(&quantity)
	fmt.Println("Masukkan harga satuan")
	fmt.Scanln(&hargaSatuan)

	itemBelanjaan.item = item
	itemBelanjaan.Quantity, _ = strconv.Atoi(quantity)
	itemBelanjaan.hargaSatuan, _ = strconv.ParseFloat(hargaSatuan, 64)
	listBelanjaan = append(listBelanjaan, itemBelanjaan)
	ListBelanjaan2(listBelanjaan)

	vInt, _ := strconv.Atoi(quantity)
	vFloat, _ := strconv.ParseFloat(hargaSatuan, 64)
	listBelanjaanMethod.listBelanjaan = append(listBelanjaanMethod.listBelanjaan, belanjaan{item, vInt, vFloat})

	fmt.Println("Lakukan perubahan pada item")

	fmt.Println("Masukkan nama Item")
	fmt.Scanln(&item)
	fmt.Println("Masukkan jumlah atau quantity")
	fmt.Scanln(&quantity)
	fmt.Println("Masukkan harga satuan")
	fmt.Scanln(&hargaSatuan)

	vInt, _ = strconv.Atoi(quantity)
	vFloat, _ = strconv.ParseFloat(hargaSatuan, 64)

	listBelanjaanMethod.UpdateBelanjaan3(item, vInt, vFloat)
	listBelanjaanMethod.ListBelanjaan3()

	listBelanjaanMethod.DiskonBelanjaan3(35, &totalPotongan)

	fmt.Println("Total potongan ", totalPotongan)

}

//ListBelanjaan example tidak bisa kita gunakan atau ERROR
// func (listBelanjaan []belanjaan) ListBelanjaan() {
// 	for i, v := range listBelanjaan {
// 		fmt.Print((i + 1), "Item ", v.item)
// 		fmt.Print("| Quantity ", v.Quantity)
// 		fmt.Print("| Harga Satuan ", v.hargaSatuan)
// 		fmt.Println("| Total ", (v.hargaSatuan * float64(v.Quantity)))
// 	}
// }

//ListBelanjaan2 example
func ListBelanjaan2(listBelanjaan []belanjaan) {
	fmt.Println("ini menggunakan parameter struct")
	for i, v := range listBelanjaan {
		fmt.Print((i + 1), "Item ", v.item)
		fmt.Print("| Quantity ", v.Quantity)
		fmt.Print("| Harga Satuan ", v.hargaSatuan)
		fmt.Println("| Total ", (v.hargaSatuan * float64(v.Quantity)))
	}
}

func (listBelanjaanMethod belanjaanKumpulam) ListBelanjaan3() {
	fmt.Println("ini menggunakan struct type method")
	for i, v := range listBelanjaanMethod.listBelanjaan {
		fmt.Print((i + 1), "Item ", v.item)
		fmt.Print("| Quantity ", v.Quantity)
		fmt.Print("| Harga Satuan ", v.hargaSatuan)
		fmt.Println("| Total ", (v.hargaSatuan * float64(v.Quantity)))
	}
}

func (listBelanjaanMethod *belanjaanKumpulam) UpdateBelanjaan3(
	item string,
	Quantity int,
	hargaSatuan float64) {
	for i, v := range listBelanjaanMethod.listBelanjaan {
		if v.item == item {
			listBelanjaanMethod.listBelanjaan[i].item = item
			listBelanjaanMethod.listBelanjaan[i].Quantity = Quantity
			listBelanjaanMethod.listBelanjaan[i].hargaSatuan = hargaSatuan
		}

	}
}

func (listBelanjaanMethod belanjaanKumpulam) DiskonBelanjaan3(
	diskon float64,
	totalPotongan *float64) {
	var total float64

	for _, v := range listBelanjaanMethod.listBelanjaan {
		total = total + (float64(v.Quantity) * v.hargaSatuan)
	}

	*totalPotongan = total * diskon / 100

	fmt.Println(totalPotongan)
	fmt.Println(*totalPotongan)
}
