package main

import "fmt"

func main() {
	count := 1
	for count <= 5 {
		fmt.Println("Perulangan ke ", count)
		count++
	}

	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan ke ", i)
	}

	slice := []string{"solo", "klaten", "sleman"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for _, value := range slice {
		// fmt.Println("index ", i, "value", value)
		fmt.Println("value", value)
	}

	nilai := 1
	for {
		if nilai == 3 {
			nilai++
			continue
		}
		fmt.Println("perulangan ke ", nilai)
		nilai++

		if nilai == 10 { //jika sudah teroptimasi maka stop
			break
		}
	}

	for k := 1; k < 10; k++ {
		if k == 3 {
			continue
		}
		fmt.Println("perulangan ke ", k)

	}
}
