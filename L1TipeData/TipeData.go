package main

import "fmt"

func main() {
	fmt.Println("Tipe Data")
	fmt.Println("Bilangan Bulat ", 6)      //bilangan bulat
	fmt.Println("Bilangan Desimal ", 6.5)  //float
	fmt.Println("Bilangan Boolean ", true) //bool
	fmt.Println("Ini String")              //string

	var nama string
	var ipk float32

	nama = "Andini"
	ipk = 3.99

	tinggi := 181

	total := ipk * 3

	fmt.Println("Nama ", nama)
	fmt.Println("Total IPK ", total)
	fmt.Println("tinggi ", tinggi, "cm")

	text := "ini string"
	fmt.Println(text[1]) //bukan n melainkan byte dari ascii
	fmt.Println(string(text[1]))
	// int8 = -128 sampai 127

	var angka int16
	var angkaConvert int8 //-128 sampai 127
	angka = 129
	angkaConvert = int8(angka)
	fmt.Println(angkaConvert)
	//uint dimulai dari positif 0 sampai 256 | int dimulai dari negatif -128 sampai 127

	const gravitasi float32 = 9.8

	// var (
	// 	kota1     string
	// 	propinsi1 string
	// )

	// var (
	// 	kota2     = "balikpapan"
	// 	propinsi2 = "kalimantan timur"
	// )
	{
		fmt.Println(angkaConvert)
		angkaConvert = 100
	}
	fmt.Println("berat badan", angkaConvert)
}
