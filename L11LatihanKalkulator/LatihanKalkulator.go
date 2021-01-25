package main

import (
	"fmt"
	"strconv"
	"strings"
)

// func historiKalkulasi(number float64, total float64, jenis string) (historiSementara string) {
// 	historiSementara = fmt.Sprintf("%f", number) + jenis + fmt.Sprintf("%f", total)
// 	total = tambah(number, total) //if else
// 	historiSementara = historiSementara + " = " + fmt.Sprintf("%f", total)
// 	return
// }

func historiKalkulasiNew(angka1, angka2 float64, total *float64, jenis string) (historiSementara string) {
	historiSementara = fmt.Sprintf("%f", angka1) + jenis + fmt.Sprintf("%f", angka2)
	historiSementara = historiSementara + " = " + fmt.Sprintf("%f", *total)
	return
}
func tambah(angka1, angka2 float64, total *float64) (historiSementara string) {
	*total = angka1 + angka2
	historiSementara = ""
	historiSementara = historiKalkulasiNew(angka1, angka2, total, "+")
	return
}
func kurang(angka1, angka2 float64, total *float64) (historiSementara string) {

	*total = angka1 - angka2
	historiSementara = historiKalkulasiNew(angka1, angka2, total, "-")
	return
}

func kali(angka1, angka2 float64, total *float64) (historiSementara string) {
	*total = angka1 * angka2
	historiSementara = historiKalkulasiNew(angka1, angka2, total, "*")
	return
}

func bagi(angka1, angka2 float64, total *float64) (historiSementara string) {
	*total = angka1 / angka2
	historiSementara = historiKalkulasiNew(angka1, angka2, total, "/")
	return
}

func menu() {
	fmt.Println("Pilih Menu")
	fmt.Println("1. Kalkulator")
	fmt.Println("2. Histori")
	fmt.Println("3. Keluar")
	fmt.Println("masukkan pilihan menu")
}

func main() {
	var inputNama, inputMenu, inputDalamMenu, jenisAritmatik string
	var total float64
	var angka1, angka2 float64
	var err error
	var histori []string
	// var historiSementara string

	fmt.Print("Masukkan nama ")
	fmt.Scanln(&inputNama)
	for {
		menu()
		fmt.Scanln(&inputMenu)

		if inputMenu == "3" {
			fmt.Print("Anda keluar dari aplikasi")
			break
		} else if inputMenu == "1" {
			fmt.Print("Anda masuk menu calculator")
			state := 1
			for {
				if state%2 == 1 {

					if state == 1 {
						fmt.Println("masukkan angka, untuk keluar masukkan simbol #")
						fmt.Println("hasil", total)
						fmt.Scanln(&inputDalamMenu)
						angka2, err = strconv.ParseFloat(inputDalamMenu, 64)
						if err != nil {
						}
						angka1 = total
						histori = append(histori, tambah(angka1, angka2, &total))
					} else {
						jenisAritmatik = inputDalamMenu
						fmt.Println("masukkan angka, untuk keluar masukkan simbol #")
						fmt.Println("hasil", total)
						fmt.Scanln(&inputDalamMenu)

						angka2, err = strconv.ParseFloat(inputDalamMenu, 64)
						angka1 = total
						if err != nil {
						}
						if jenisAritmatik == "+" {
							histori = append(histori, tambah(angka1, angka2, &total))
						} else if jenisAritmatik == "-" {
							histori = append(histori, kurang(angka1, angka2, &total))
						} else if jenisAritmatik == "*" {
							histori = append(histori, kali(angka1, angka2, &total))
						} else if jenisAritmatik == "/" {
							histori = append(histori, bagi(angka1, angka2, &total))
						}

					}
				} else {
					for {

						fmt.Println("masukkan aritmatika (+ - / *), kembali ke menu masukkan simbol #")
						fmt.Println("hasil", total)
						fmt.Scanln(&inputDalamMenu)
						index := strings.Index("+-*/", inputDalamMenu)
						fmt.Println(index)
						if index >= 0 || inputDalamMenu == "#" {
							break
						}
					}

				}

				if inputDalamMenu == "#" {
					fmt.Print("Anda keluar dari kalkulator")
					break
				}
				state++
			}

		} else if inputMenu == "2" {
			for i, v := range histori {
				fmt.Println((i + 1), " >> ", v)
			}
		}

	}

}
