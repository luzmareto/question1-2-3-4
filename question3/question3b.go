package question

import "fmt"

func more5words(strings []string) []string {
	filteredStrings := []string{}

	// Mengiterasi setiap string dalam array
	for _, str := range strings {
		if len(str) > 5 {
			filteredStrings = append(filteredStrings, str)
		}
	}

	return filteredStrings
}

func Question3b() {
	strings := []string{"apple", "banana", "cherry", "orange", "kiwi"}

	// Memmengembalikan array dengan string panjang > 5
	filteredStrings := more5words(strings)

	fmt.Println("strings with a length greater than 5 characters:")
	fmt.Println(filteredStrings)
}
