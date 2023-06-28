package question

import "fmt"

func getCount(numbers []int) int {
	total := 0

	// Mengiterasi setiap bilangan dalam array
	for _, num := range numbers {
		total += num
	}

	return total
}

func Question3a() {
	bilangan := []int{2, 5, 10, 3, 8}

	// Menghitung jumlah total angka
	total := getCount(bilangan)

	fmt.Println("Total:", total)
}
