package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
)

type phonebook struct {
	name   string
	number string
	email  string
}

var currentData phonebook
var listData []phonebook
var selectedMenu string
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var clear map[string]func()
var scanner = bufio.NewScanner(os.Stdin)

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	menu()
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

func changeMenu() {
	clearScreen()
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

	listData = append(listData, currentData)

	menu()
}

func validate(tp string) {
	switch tp {
	case "name":
		if currentData.name == "!" {
			exit()
		}

		if len(currentData.name) < 5 {
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
		if len(currentData.number) < 10 {
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

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

func exit() {
	fmt.Println("Goodbye")
	os.Exit(0)
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

		for i := 0; i < len(listData); i++ {
			s := &listData[i]
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

func checkExist() bool {
	for _, s := range listData {
		if s.name == currentData.name {
			return true
		}
	}

	return false
}

func list() {
	for i, s := range listData {
		fmt.Print(i+1, " Name : ", s.name)
		fmt.Print(", Number : ", s.number)
		fmt.Println(", Email : ", s.email)
	}
	fmt.Println("- - - - - - - - - - - -")
	fmt.Println("back to menu? (y/n) ")
	fmt.Scan(&selectedMenu)
	changeMenu()
}

func clearScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
