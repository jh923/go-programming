package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("This function is assigned to a variable. ")
	}
	f()

	gorwell := orwell()
	fmt.Printf("%T\n", gorwell())
	fmt.Println(gorwell())
	fmt.Println(bar()())

	ls := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	mul := even(mul, ls...)
	print(mul)
}

func orwell() func() int {
	return func() int {
		return 1984
	}
}

func bar() func() int {
	return func() int {
		return 451
	}
}

func mul(is ...int) int {
	total := 1
	for _, v := range is {
		total = total * v
	}
	return total
}

func even(f func(is ...int) int, xs ...int) int {
	var ys []int
	for _, v := range xs {
		if v%2 == 0 {
			ys = append(ys, v)
		}
	}
	return f(ys...)
}
