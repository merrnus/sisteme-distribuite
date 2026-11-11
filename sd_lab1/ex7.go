// ex7: Citeste cuvinte dintr-un fisier si afiseaza de cate ori apare fiecare.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
		return
	}	
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	words := []string{}
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}

	for w, c := range counts {
		fmt.Println(w, c)
	}
}
