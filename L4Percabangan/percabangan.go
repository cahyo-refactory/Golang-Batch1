package main

import "fmt"

func main() {
	suhu := 20

	if suhu <= 0 {
		fmt.Println("beku")
	} else if suhu > 0 && suhu <= 100 {
		fmt.Println("cair")
	} else {
		fmt.Println("beku")
	}

	nama := "dwi"

	if length := len(nama); length <= 5 {
		fmt.Println(length, "nama terlalu pendek")
	} else {
		fmt.Println(length, "nama sudah benar")
	}

	if len(nama) <= 5 {
	} else {
		fmt.Println("nama sudah benar")
	}

	switch nama {
	case "budi":
		fmt.Println("Hallo Budi")
	case "dwi":
		fmt.Println("Hallo dwi")
	default:
		fmt.Println("nama belum terdaftar")
	}

	switch length := len(nama); length <= 5 {
	case true:
		fmt.Println("nama terlalu pendek")
	default:
		fmt.Println("nama sudah benar")
	}
	nama = "cahyo purnomo"
	suhu = -3
	length := len(nama)

	switch {
	case length <= 5:
		fmt.Println("nama terlalu pendek")
	case length > 5:
		fmt.Println("nama sudah benar")
	case suhu <= 0:
		fmt.Println("Beku")
	}

}
