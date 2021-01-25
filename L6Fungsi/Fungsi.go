package main

import "fmt"

// basic fungsi
func sayHello() {
	fmt.Println("Hallo Xapiens")
}

//fungsi dengan parameter
func sayName(name string) {
	fmt.Println("Hallo ", name)
}

// fungsi dengan return
func sayReturn(name string) string {
	return "Hello " + name
}

// fungsi dengan multiple return
func sayMultipleReturn() (string, string, string) {
	return "ngalik", "sleman", "DIY"
}

//alias variabel return
func sayMultipleReturnAlias() (kec, kab, prov string) {
	kec = "ngalik"
	kab = "sleman"
	prov = "DIY"

	return
}

//fungsi variadic
//jika ada parameter tambahan tipe int juga maka caranya gmn
//jawaban ada di bawah
func sumAll(pengali int, pembagi int, numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total * pengali
}

//fungsi sebagai parameter
type Filter func(string) string

func textFilter(name string, filter Filter) {
	filtered := filter(name)
	fmt.Println("Hello ", filtered)
}

func filter(text string) string {
	if text == "jahat" {
		return "....."
	} else {
		return text
	}
}

func filter2(text string) string {
	if text == "cahyo" {
		return "....."
	} else {
		return text
	}
}

//anonymus function
type BlackList func(string) bool

func registerUser(name string, blakList BlackList) {
	if blakList(name) {
		fmt.Println("are you blocked ", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

//recrusive
func factorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorial(value-1)
	}
}

func main() {
	fmt.Println(5 * 4 * 3 * 2 * 1)
	fmt.Println(factorial(5))

	var getMinMax = func(numbers ...int) (int, int) {
		var min, max int
		min = numbers[0]
		max = numbers[0]
		for _, v := range numbers {
			if min > v {
				min = v
			}
			if max < v {
				max = v
			}
		}
		return min, max
	}

	min, max := getMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(min, max)

	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blacklist)
	registerUser("arman", func(name string) bool {
		return name == "admin"
	})

	textFilter("cahyo", filter)
	textFilter("jahat", filter)

	textFilter("cahyo", filter2)
	textFilter("jahat", filter2)

	sayHello()

	sayName("xapiens")

	fmt.Println(sayReturn("xapiens"))

	fmt.Println(sayMultipleReturn())
	kec, kab, prov := sayMultipleReturn()
	fmt.Println("2 ", kec, kab, prov)

	fmt.Println(sayMultipleReturnAlias())

	fmt.Println(sumAll(100, 1, 2, 3, 4, 5))

}
