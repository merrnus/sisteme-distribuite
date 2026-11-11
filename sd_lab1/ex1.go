// ex1: Citeste un numar natural de la tastatura si verifica daca este prim.

package main 

import "fmt"

func main() {
	var n int 
	fmt.Scan(&n)

	prime := n >= 2
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			prime = false
			break

		}
	}
	
	if prime {
	fmt.Println ("Prim")
	} else {
		fmt.Println("Nu prim")
	}

}

