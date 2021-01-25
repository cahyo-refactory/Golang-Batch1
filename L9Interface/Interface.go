package main

import "fmt"

//HasName example
type HasName interface {
	GetName() string
	GetAge() int
}

//Person example
type Person struct {
	name string
	age  int
}

//SayHello example
func SayHello(person HasName) {
	fmt.Println("Hello ", person.GetName(), "anda berumur", person.GetAge())
}

func (person Person) GetName() string {
	return person.name
}

func (person Person) GetAge() int {
	return person.age
}

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "string"
	}
}

func assertions() interface{} {
	return 1
}

func main() {
	//interface
	var person Person
	person.name = "tika"
	person.age = 45

	SayHello(person)

	//interface kosong
	var data interface{} = Ups(3)
	fmt.Println(data)

	//interface assertions
	//tanpa switch
	var result interface{} = assertions()
	// var resultString string = result.(string)
	// fmt.Println(resultString)
	// var resultInt int = result.(int)
	// fmt.Println(resultInt)
	
	//dengan switch
	switch value := result.(type) {
	case string:
		fmt.Println("Program sudah benar bertipe string nilai ", value)
	default:
		fmt.Println("Error karena tipe data bukan string nilai", value)
	}

}
