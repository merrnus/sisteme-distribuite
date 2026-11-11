// ex5: Inverseaza ordinea elementelor unui slice de string-uri, fara alt slice.

package main

import "fmt"

func main() {
	s := []string{"Laboratorul", "de", "sisteme", "distribuite"}

	for i, j := 0, len(s)-1; i < j; i, j=i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	fmt.Println(s)
}
