package main

import "fmt"

func main() {
	var i int

loop:
	for i = 0; i < 10; i++ {
		fmt.Printf("for i = %d\n", i)
		switch i {
		case 2, 3:
			fmt.Printf("Quebrando switch, i == %d.\n", i)
			break
		case 5:
			fmt.Printf("Quebrando loop, i == %d.\n", i)
			break loop
		}
	}
	fmt.Println("Fim.")
}
