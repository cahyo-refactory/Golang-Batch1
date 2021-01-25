package main

import "fmt"

type Mahasiswa struct {
	nama, asal string
	age        int
	ipk        float32
}

func (mahasiswa Mahasiswa) halloMahasiswa(name string) {
	fmt.Println("hello ", name, "nama saya", mahasiswa.nama)
}

func main() {

	var andi Mahasiswa
	andi.nama = "Andi"
	andi.asal = "Brebes"
	andi.age = 17
	andi.ipk = 3.0
	fmt.Println(andi)
	fmt.Println(andi.nama)

	afni := Mahasiswa{
		nama: "Nur Afni",
		asal: "sorong",
		age:  16,
		ipk:  3.8,
	}

	fmt.Println(afni)
	fmt.Println(afni.nama)

	siti := Mahasiswa{"siti nurbaya", "aceh", 20, 4}

	fmt.Println(siti)
	fmt.Println(siti.nama)

	siti.halloMahasiswa("siti")

}
