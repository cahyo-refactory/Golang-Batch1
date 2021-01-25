package main

import "fmt"

func main() {
	//Array
	var cities [3]string
	cities[0] = "balikpapan"
	cities[1] = "sleman"
	cities[2] = "surabaya"

	fmt.Println(cities)
	fmt.Println(cities[0])

	var ages = [3]int{
		30, 20, 10,
	}

	fmt.Println(ages)
	fmt.Println(ages[0])
	fmt.Println(len(ages))

	var angkaArray = [...]string{
		"nol",
		"satu",
		"dua",
	}

	fmt.Println(angkaArray)
	fmt.Println(len(angkaArray))
	fmt.Println(cap(angkaArray))
	//slice
	fmt.Println("===== Ini Slice ======")
	var angkaSlice = []string{
		"nol",
		"satu",
		"dua",
	}

	fmt.Println(angkaSlice)
	fmt.Println(len(angkaSlice))
	fmt.Println(cap(angkaSlice))

	angkaSlice = append(angkaSlice, "tiga")

	fmt.Println(angkaSlice)
	fmt.Println(len(angkaSlice))
	fmt.Println(cap(angkaSlice))
	fmt.Println(angkaSlice[0])
	// fmt.Println(angkaSlice[4]) ini error

	//slice dari array
	fmt.Println("Membuat slice dari array")
	var angkaBiasa = [...]string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
	}

	var slice1 = angkaBiasa[2:4]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	newSlice := make([]string, 3, 6)

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice[0] = "0"
	newSlice[1] = "1"
	newSlice[2] = "2"

	fmt.Println(newSlice[0])
	fmt.Println(newSlice[1])
	fmt.Println(newSlice[2])

	newSlice = append(newSlice, "3")
	//cap hanya informasi
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice = append(newSlice, "4")
	//cap hanya informasi
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice = append(newSlice, "5")
	//cap hanya informasi
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice = append(newSlice, "6")
	//cap hanya informasi
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fmt.Println("Slice refrens to array")
	var alphabet = [...]string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
	}
	var alphaSlice = alphabet[0:5]
	fmt.Println(alphabet)
	fmt.Println(alphaSlice)

	alphaSlice[1] = "b2"

	fmt.Println(alphabet)
	fmt.Println(alphaSlice)

	var alphaSlice2 = alphabet[0:5]
	fmt.Println(alphabet)
	fmt.Println(alphaSlice)
	fmt.Println(alphaSlice2)
	alphaSlice[1] = "b3"
	fmt.Println(alphabet)
	fmt.Println(alphaSlice)
	fmt.Println(alphaSlice2)

	fmt.Println(len(alphaSlice), cap(alphaSlice))
	alphaSlice = append(alphaSlice, "f2")
	fmt.Println(alphabet)
	fmt.Println(alphaSlice)
	fmt.Println(alphaSlice2)
	alphaSlice2 = append(alphaSlice2, "f3")
	fmt.Println(alphabet)
	fmt.Println(alphaSlice)
	fmt.Println(alphaSlice2)

	//map
	user := map[string]string{
		"nama": "cahyo",
		"kota": "balikpapan",
	}
	fmt.Println(user)
	fmt.Println(len(user))

	user["umur"] = "30"

	fmt.Println(user)
	fmt.Println(len(user))

	delete(user, "kota")

	fmt.Println(user)
	fmt.Println(len(user))

	user["nama"] = "arman"

	fmt.Println(user)
	fmt.Println(len(user))

	var karyawan = make(map[string]string)
	karyawan["NIP"] = "01234"

	fmt.Println("Map dengan slice")
	var mahasiswa = []map[string]string{
		map[string]string{"nama": "arman", "nim": "1234"},
		map[string]string{"nama": "agus", "nim": "2345"},
	}

	fmt.Println(mahasiswa[0]["nama"])
	mahasiswa = append(mahasiswa, map[string]string{"nama": "cahyo", "nim": "3456"})
	fmt.Println(mahasiswa)

	var namaPesertaXapiens = []string{
		"budi",
		"siti",
		"rahman",
	}
	fmt.Println(namaPesertaXapiens)
	fmt.Println(len(namaPesertaXapiens))
	fmt.Println(cap(namaPesertaXapiens))
	fmt.Println(namaPesertaXapiens[1])

	namaPesertaXapiens = append(namaPesertaXapiens, "pebri")

	fmt.Println(namaPesertaXapiens)
	fmt.Println(len(namaPesertaXapiens))
	fmt.Println(cap(namaPesertaXapiens))
	fmt.Println(namaPesertaXapiens[1])

	namaPesertaXapiens = append(namaPesertaXapiens, "nurul")

	fmt.Println(namaPesertaXapiens)
	fmt.Println(len(namaPesertaXapiens))
	fmt.Println(cap(namaPesertaXapiens))
	fmt.Println(namaPesertaXapiens[1])

}
