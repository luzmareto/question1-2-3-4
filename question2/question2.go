package question_2

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type CharacterCount struct {
	Char  rune
	Count int
}

// Menghitung kemunculan setiap karakter dalam string
func Count(input string) []CharacterCount {
	counts := make(map[rune]int)

	// Mengiterasi setiap karakter dalam string
	for _, char := range input {
		counts[char]++
	}

	// Mengonversi map ke slice of struct
	result := []CharacterCount{}
	for char, count := range counts {
		result = append(result, CharacterCount{Char: char, Count: count})
	}

	// Mengurutkan slice berdasarkan karakter
	sort.Slice(result, func(i, j int) bool {
		return result[i].Char < result[j].Char
	})

	return result
}

func Question2() {
	fmt.Print("Enter words, spaces, numbers: ")

	// Membaca input string dari terminal
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Menghilangkan leading dan trailing whitespace pada input
	input = strings.TrimSpace(input)

	// Menghitung kemunculan setiap karakter dalam string
	result := Count(input)

	// Menampilkan hasil kemunculan karakter
	for _, charCount := range result {
		fmt.Printf("Karakter '%c' muncul sebanyak %d kali\n", charCount.Char, charCount.Count)
	}
}
