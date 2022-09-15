package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	v := 6

	fmt.Println("Hello, World!", v)
	fmt.Println(add(6, 8))

}
