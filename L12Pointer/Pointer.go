package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {

	address1 := Address{"klaten", "jateng", "indonesia"}
	address2 := address1

	address1.city = "bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	var address3 *Address = &address1
	//addrres3 pasti akan menerima perubahan apapun dari address1 dan sebaliknya

	address3.city = "solo"

	fmt.Println("1 1", address1)
	fmt.Println("1 2", address2)
	fmt.Println("1 3", address3)

	var address4 = &Address{"balikpapan", "kaltim", "indonesia"}
	fmt.Println("2 1", address1)
	fmt.Println("2 2", address2)
	fmt.Println("2 3", address3)
	fmt.Println("2 4", address4)

	// address3 = &Address{"banjarmasin", "kalsel", "indonesia"}
	// fmt.Println("3 1", address1)
	// fmt.Println("3 2", address2)
	// fmt.Println("3 3", address3)
	// fmt.Println("3 4", address4)

	var address5 *Address = &address1
	// address5.city = "bandung"
	*address5 = Address{"surabaya", "jatim", "indonesia"}
	fmt.Println("4 1", address1)
	fmt.Println("4 2", address2)
	fmt.Println("4 3", address3)
	fmt.Println("4 4", address4)
	fmt.Println("4 5", address5)

	address1.city = "tamrin"
	address3.province = "jakarta"
	address5.country = "Ina"

	fmt.Println("5 1", address1)
	fmt.Println("5 2", address2)
	fmt.Println("5 3", address3)
	fmt.Println("5 4", address4)
	fmt.Println("5 5", address5)

	var address6 *Address = new(Address)
	fmt.Println("6 6", address6)

	var ankgas float64
	ankgas = 4.4

	changeCity(&address2, &ankgas)
	changeCity(&address1, &ankgas)

	fmt.Println("7 1", address1)
	fmt.Println("7 2", address2)
	fmt.Println("7 3", address3)
	fmt.Println("7 5", address5)

	address2.city = "kudus"
	fmt.Println("8 1", address1)
	fmt.Println("8 2", address2)
	fmt.Println("8 3", address3)
	fmt.Println("8 5", address5)

	cahyo := Person{"cahyo"}
	cahyo.school()
	fmt.Println(cahyo.name)
}

func changeCity(addres *Address, angka *float64) {
	addres.city = "semarang"
	fmt.Println(*angka)
}

type Person struct {
	name string
}

func (person *Person) school() {
	person.name = "Hitler"
}
