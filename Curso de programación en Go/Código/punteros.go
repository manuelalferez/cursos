package main

import "fmt"

func main() {
	x := 25
	fmt.Println(x)
	cambiarValor(&x)
	fmt.Println(x)
}

func cambiarValor(a *int) {
	*a = 30
}
