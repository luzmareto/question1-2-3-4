package question_1

import (
	"fmt"
	"strconv"
)

// operasi penjumlahan
func Input(bilangan1, bilangan2 int) int {
	return bilangan1 + bilangan2
}

// membuat input dari terminal
func ReadInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func Question1() {
	//Membaca input data pertama dari terminal
	fmt.Print("Enter the first value: ")
	input1, _ := strconv.Atoi(ReadInput())

	//Membaca input data kedua dari terminal
	fmt.Print("Enter the secound value: ")
	input2, _ := strconv.Atoi(ReadInput())

	// Menghitung jumlah dua angka
	result := Input(input1, input2)

	// Menampilkan hasil penjumlahan
	fmt.Println("Result:", result)
}
