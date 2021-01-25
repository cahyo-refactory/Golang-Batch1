package main

import (
	"fmt"
)

func menu() {
	fmt.Println("Pilih Menu")
	fmt.Println("1. Input nomor telepon")
	fmt.Println("2. Update nomor telepon")
	fmt.Println("3. List nomor telepon")
	fmt.Println("4. keluar")
	fmt.Println("Input pilihan menu")
}
func main() {
	var inputNama, inputNotelp, inputEmail, inputMenu string
	fmt.Print("Masukkan nama = ")
	fmt.Scanln(&inputNama)
	fmt.Print("Masukkan No.Telepon = ")
	fmt.Scanln(&inputNotelp)
	fmt.Print("Masukkan Email = ")
	fmt.Scanln(&inputEmail)
	for {
		menu()
		fmt.Scanln(&inputMenu)
		if inputMenu == "4" {
			fmt.Print("Anda keluar dari aplikasi")
			break
		} else if inputMenu == "1" {
			fmt.Print("Input nomor telepon")
			// state := 1
		} else if inputMenu == "2" {
			fmt.Print("Update nomor telepon")
		}
		// } else if inputMenu == "3" {
		// 	fmt.Print("List nomor telepon")
		// }
	}
}
