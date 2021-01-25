package main

import (
	"errors"
	"fmt"
)

func logRecover() {
	fmt.Println("selesai memanggil fungsi")
	pesanError := recover()
	if pesanError != nil {
		fmt.Println("Terjadi Error ", pesanError)
	}
}

func pembagian(value int) {
	defer logRecover()
	fmt.Println("ini pembagian")
	total := 10 / value
	fmt.Println(total)
}

func pembagianDenganError(value int) (int, error) {
	if value == 0 {
		return 0, errors.New("Pembagi tidak boleh nol")
	} else {
		return (10 / value), nil
	}
}

func main() {
	//pembagian 0
	//tipe data tidak sesuai
	//range tipe data tidak sesuai
	//defer --> try
	//panic --> error
	//recover --> catch
	pembagian(0)
	hasil, err := pembagianDenganError(0)
	fmt.Println("hasil ", hasil, " pesan error", err)
}
