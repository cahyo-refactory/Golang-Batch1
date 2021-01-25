package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Phone struct {
	name, number, email string
}

var contactBook = []Phone{}
var selectedMenu string
var currentData Phone
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var scanner = bufio.NewScanner(os.Stdin)

func checkHuruf(name string) bool {
	isStringAlphabetic, _ := regexp.MatchString(`^[a-z A-Z]+$`, name)
	return isStringAlphabetic
}

func checkNomor(nomor string) bool {
	isStringNumeric, _ := regexp.MatchString(`^[\d]+$`, nomor)
	return isStringNumeric
}

func checkEmail(email string) bool {
	isStringValid, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	return isStringValid
}

func validate(tp string) {
	switch tp {
	case "name":
		if currentData.name == "!" {
			exit()
		}
		if checkHuruf(currentData.name) && len(currentData.name) < 5 {
			fmt.Println("Minimal Name length is 5")
			fmt.Print("Name : ")
			scanner.Scan()
			currentData.name = scanner.Text()
			validate("name")
		}
	case "number":
		if currentData.number == "!" {
			exit()
		}
		if checkNomor(currentData.number) && len(currentData.number) < 10 {
			fmt.Println("Minimal Number length is 10")
			fmt.Print("Number : ")
			fmt.Scanln(&currentData.number)
			validate("number")
		}
	case "email":
		if currentData.email == "!" {
			exit()
		}
		if !isEmailValid(currentData.email) {
			fmt.Println("Email Not Valid")
			fmt.Print("Email : ")
			fmt.Scanln(&currentData.email)
			validate("email")
		}
	case "!":
		exit()
	}
}
func exit() {
	fmt.Println("Goodbye")
	os.Exit(0)
}
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

func list() {
	for i, s := range contactBook {
		fmt.Print(i+1, " Name : ", s.name)
		fmt.Print(", Number : ", s.number)
		fmt.Println(", Email : ", s.email)
	}
	fmt.Println("- - - - - - - - - - - -")
	fmt.Println("back to menu? (y/n) ")
	fmt.Scan(&selectedMenu)
	changeMenu()
}
func menu() {
	fmt.Println("Menu Aplikasi")
	fmt.Println("1. Add Phone Number")
	fmt.Println("2. Update Phone Number")
	fmt.Println("3. List Phone Number")
	fmt.Println("!. Exit App")
	fmt.Println("- - - - - - - - - - - - -")
	fmt.Print("Pilih Menu : ")
	fmt.Scanln(&selectedMenu)
	changeMenu()
}
func add() {

	fmt.Println("Input Data")
	fmt.Println("-------------------")
	fmt.Print("Name : ")
	scanner.Scan()
	currentData.name = scanner.Text()
	validate("name")
	fmt.Print("Number : ")
	fmt.Scanln(&currentData.number)
	validate("number")
	fmt.Print("Email : ")
	fmt.Scanln(&currentData.email)
	validate("email")

	contactBook = append(contactBook, currentData)

	menu()
}
func changeMenu() {
	switch selectedMenu {
	case "1":
		add()
	case "2":
		update()
	case "3":
		list()
	case "!":
		exit()
	case "y":
		menu()
	case "n":
		exit()
	default:
		fmt.Println("Choice not available")
		menu()
	}
}
func checkExist() bool {
	for _, s := range contactBook {
		if s.name == currentData.name {
			return true
		}
	}

	return false
}
func update() {
	fmt.Print("Input Name : ")
	scanner.Scan()
	currentData.name = scanner.Text()
	if checkExist() {
		fmt.Print("Number : ")
		fmt.Scanln(&currentData.number)
		validate("number")
		fmt.Print("Email : ")
		fmt.Scanln(&currentData.email)
		validate("email")

		for i := 0; i < len(contactBook); i++ {
			s := &contactBook[i]
			if s.name == currentData.name {
				s.number = currentData.number
				s.email = currentData.email
				break
			}
		}
	} else {
		fmt.Println("Data not exist")
	}

	menu()
}

func main() {
	menu()
}
