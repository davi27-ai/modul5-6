package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif n: ")
	fmt.Scan(&n)
	if n < 1 {
		fmt.Println("Input harus bilangan bulat positif.")
		return
	}
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("Hasil penjumlahan dari 1 sampai %d adalah %d\n", n, sum)
}
