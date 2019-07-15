package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	var a [10]int // LA CLAVE ES QUE ESTE ARRAY SE MANTIENE INALTERADO CADA VEZ QUE SE LLAMA A LA FUNC
	return func(x int) int {
		switch {
		case x == 0:
			a[x] = 0
		case x == 1:
			a[x] = 1
		default:
			a[x] = a[x-1] + a[x-2]
		}
		return a[x]
	}

}

func main() {

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
