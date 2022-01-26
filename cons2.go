package main

import "fmt"

type fn func(interface{}, interface{}) interface{}

func cons(a, b interface{}) func(func(interface{}, interface{}) interface{}) interface{} {
	return func(f func(interface{}, interface{}) interface{}) interface{} {
		return f(a, b)
	}
}

func cdr(pair func(func(interface{}, interface{}) interface{}) interface{}) interface{} {
	return pair(first)
}

func car(pair func(func(interface{}, interface{}) interface{}) interface{}) interface{} {
	return pair(second)
}

func first(a interface{}, b interface{}) interface{} {
	return a
}

func second(a interface{}, b interface{}) interface{} {
	return b
}

func main() {
	pair := cons(1, 2)
	fmt.Printf("%T\n", pair)

	y := cdr(pair)
	fmt.Printf("%T %v\n", y, y)
	z := car(pair)
	fmt.Printf("%T %v\n", z, z)
}
