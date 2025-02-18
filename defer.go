package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

