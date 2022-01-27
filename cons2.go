package main

import "fmt"

type fn func(interface{}, interface{}) interface{}
type ffn func(fn) interface{}

func cons(a, b interface{}) ffn {
	return func(f fn) interface{} {
		return f(a, b)
	}
}

func cdr(pair ffn) interface{} {
	return pair(first)
}

func car(pair ffn) interface{} {
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
	fmt.Printf("pair type %T\n", pair)

	y := cdr(pair)
	fmt.Printf("cdr %T %v\n", y, y)
	z := car(pair)
	fmt.Printf("car %T %v\n", z, z)

	var list interface{}
	list = cons(0, cons(1, cons(2, cons(3, nil))))
	fmt.Printf("list type %T\n", list)

	var hd interface{}
	for list != nil {
		hd, list = cdr(list.(ffn)), car(list.(ffn))
		fmt.Printf("head %v, rest %v\n", hd, list)
	}
}
