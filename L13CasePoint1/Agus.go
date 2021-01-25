package main

import (
	"fmt"
	"regexp"
)

// Global var
var contact Phone
var contactBook = []Phone{}
var inputMenu, inputSubMenu, inputNameSementara, inputNumberSementara string

// Phone example
type Phone struct {
	name, number, email string
}

func checkName(name string) bool {
	matchString, _ := regexp.MatchString(`^[a-zA-Z]+$`, name)
	if !matchString || len(name) < 5 {
		return false
	}
	return true
}

func checkNumber(number string) bool {
	matchString, _ := regexp.MatchString(`^[\d]+$`, number)
	if !matchString || len(number) < 10 {
		return false
	}
	return true
}

func checkEmail(email string) bool {
	matchString, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	return matchString
}

func main() {
	menu()
}

func menu() {
	for {
		fmt.Println()
		fmt.Println("Pilih Menu")
		fmt.Println("1. Input nomor telepon")
		fmt.Println("2. Update nomor telepon")
		fmt.Println("3. List nomor telepon")
		fmt.Println()
		fmt.Println("Masukkan karakter '!' untuk keluar")
		fmt.Print("Masukkan pilihan menu : ? ")

		fmt.Scanln(&inputMenu)
		if inputMenu == "!" {
			fmt.Println()
			fmt.Println("Anda keluar dari aplikasi")
			fmt.Println()
			break
		} else if inputMenu == "1" {
			fmt.Println()
			fmt.Println("Anda Masuk Pilihan INPUT NOMOR TELEPON")
			fmt.Println()
			addContactName()
			addContactNumber()
			addContactEmail()
			contactBook = append(contactBook, contact)
		} else if inputMenu == "2" {
			fmt.Println()
			fmt.Println("Anda Masuk Pilihan UPDATE NOMOR TELEPON")
			fmt.Println()
			updateContact()
		} else if inputMenu == "3" {
			contactList()
		} else {
			fmt.Println()
			fmt.Println("Masukkan pilihan yang benar")
			fmt.Println()
			continue
		}
	}
}

func addContactName() {
	for {
		fmt.Println("Masukkan Nama Anda : ")
		fmt.Scanln(&contact.name)

		if !checkName(contact.name) {
			fmt.Println()
			fmt.Println("FORMAT NAMA HANYA ALPHABET dan TIDAK BOLEH KURANG dari 5")
			fmt.Println()
			continue
		}
		break
	}
}

func addContactNumber() {
	for {
		fmt.Println("Masukkan Nomor Anda : ")
		fmt.Scanln(&contact.number)

		if !checkNumber(contact.number) {
			fmt.Println()
			fmt.Println("FORMAT NOMOR HANYA ANGKA DENGAN MINIMAL 10 KARAKTER")
			fmt.Println()
			continue
		}
		break
	}
}

func addContactEmail() {
	for {
		fmt.Println("Masukkan Email Anda : ")
		fmt.Scanln(&contact.email)

		if contact.email == "!" {
			menu()
		}

		if !checkEmail(contact.email) {
			fmt.Println()
			fmt.Println("EMAIL YANG DIMASUKKAN TIDAK VALID")
			fmt.Println()
			continue
		}
		break
	}
}

func updateContact() {
	for {
		fmt.Print("Masukkan nama : ")
		fmt.Scanln(&inputNameSementara)
		noData := false

		if inputNameSementara == "!" {
			menu()
		}

		for i := 0; i < len(contactBook); i++ {
			if contactBook[i].name == inputNameSementara {
				fmt.Print("Masukkan nomor : ")
				fmt.Println()
				fmt.Scanln(&inputNumberSementara)
				fmt.Println("KONTAK BERHASIL DI UPDATE")
				fmt.Println("Nama\t\t : ", contactBook[i].name)
				fmt.Println("Nomor sebelumnya : ", contactBook[i].number)
				contactBook[i].number = inputNumberSementara
				fmt.Println("Nomor Baru\t : ", contactBook[i].number)
				noData = false
				break
			} else {
				noData = true
			}
		}

		if noData == true {
			fmt.Println("NAMA TIDAK ADA")
			continue
		}
		break
	}
}

func contactList() {
	for i := 0; i < len(contactBook); i++ {
		fmt.Println()
		fmt.Println("Kontak ", i+1)
		fmt.Println("Nama\t: ", contactBook[i].name)
		fmt.Println("Nomor\t: ", contactBook[i].number)
		fmt.Println("Email\t: ", contactBook[i].email)
		fmt.Println()
	}
}
