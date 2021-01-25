package main

import (
	"fmt"
	"strconv"
)

// History struct
type History struct {
	log string
}

func tambah(angka1, angka2 float64) float64 {
	return (angka1 + angka2)
}
func kurang(angka1, angka2 float64) float64 {
	return (angka1 - angka2)
}
func kali(angka1, angka2 float64) float64 {
	return (angka1 * angka2)
}
func bagi(angka1, angka2 float64) float64 {
	return (angka1 / angka2)
}
func menu() {
	fmt.Println()
	fmt.Println("Pilih Menu")
	fmt.Println("1. Kalkulator")
	fmt.Println("2. Histori")
	fmt.Println("3. Keluar")
	fmt.Println("masukkan pilihan menu")
}
func kalkulator() {
	var inputMenu, inputDalamMenu, jenisAritmatik string
	var total float64
	var number float64
	var err error
	var historiSementara string
	var histori []string
	for {

		menu()
		fmt.Scanln(&inputMenu)
		if inputMenu == "3" {
			fmt.Println("Anda keluar dari aplikasi")
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
						number, err = strconv.ParseFloat(inputDalamMenu, 64)
						if err != nil {
						}

						historiSementara = fmt.Sprintf("%f", number) + "+" + fmt.Sprintf("%f", total)
						total = tambah(number, total)
						histori = append(histori, historiSementara)
					} else {
						jenisAritmatik = inputDalamMenu
						fmt.Println("masukkan angka, untuk keluar masukkan simbol #")
						fmt.Println("hasil", total)
						fmt.Scanln(&inputDalamMenu)
						number, err = strconv.ParseFloat(inputDalamMenu, 64)
						if err != nil {
						}
						if jenisAritmatik == "+" {
							total = tambah(total, number)
						} else if jenisAritmatik == "-" {
							total = kurang(total, number)
						} else if jenisAritmatik == "*" {
							total = kali(total, number)
						} else if jenisAritmatik == "/" {
							total = bagi(total, number)
						}
					}
				} else {
					fmt.Println("masukkan aritmatika (+ - / *), kembali ke menu masukkan simbol #")
					fmt.Println("hasil", total)
					fmt.Scanln(&inputDalamMenu)
					number, err = strconv.ParseFloat(inputDalamMenu, 64)
					if err != nil {
					}
					total = tambah(number, total)
				}
				if inputDalamMenu == "#" {
					fmt.Print("Anda keluar dari kalkulator")
					break
				}
				state++
			}
		}
	}
}
func main() {
	// var histori History
	var inputNama string
	fmt.Print("Masukkan nama ")
	fmt.Scanln(&inputNama)
	kalkulator()
}
