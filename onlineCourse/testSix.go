/**
The tests can be found at https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#
*/

package main

import (
	"fmt"
	"math"
)

func main()	{
	defer test63()
	x := test611()
	y, z := test612()
	fmt.Println(x, y, z)
	ls := []int{1,2,34,5,6,78,9}
	fmt.Println(test621(ls...))
	fmt.Println(test622(ls))
	fmt.Println("Exiting main")
	t64 := test64Person{
		first: "James",
		last: "Bond",
		age: 32,
	}
	t64.test64Speak()
	t65sq := t65Square{
		length: 10,
	}
	fmt.Println(t65sq.t65Area())
	t65c := t65Circle{
		diameter: 15,
	}
	fmt.Println(t65c.t65Area())
	fmt.Println(t65Area(t65sq))

	func(){
		fmt.Println("This is for test 6.6")
	}()

	t67 := t67()
	fmt.Println(t67)

	t68v:= t68()
	fmt.Println(t68v())

	t59xs := []int{2,3,4,5,6,7,8,9,10,11,1,21,1234,23345,32,1}
	fmt.Println(t59Even(t59Mul, t59xs...))
}

func test611() int {
	return 1
}

func test612() (int, string)	{
	return 1, "s"
}

func test621(xs ...int)	int {
	total := 0
	for _, v := range xs {
		total += v
	}
	return total
}

func test622(xs []int)	int {
	total := 0
	for _, v := range xs {
		total += v
	}
	return total
}

func test63()	{
	fmt.Println("Executing deferred function")
}


func (p test64Person) test64Speak()	{
	fmt.Printf("I am %s %s and I am %d years old.\n", p.first, p.last, p.age)
}

type test64Person struct {
	first string
	last string
	age int
}

type t65Circle struct {
	diameter float64
}

type t65Square struct {
	length float64
}

type t65Shape interface {
	t65Area() float64
}

func (s t65Square) t65Area() float64{
	return s.length * s.length
}

func (c t65Circle) t65Area() float64{
	return  math.Pi * (c.diameter/2) * (c.diameter/2)
}

func t65Area(s t65Shape) float64	{
	return s.t65Area()
}

func t67() string	{
	return "This is for test 6.7"
}


func t68() func()	int	{
	return func() int {
		return 42
	}
}


func t59Mul(is ...int) int {
	total := 1
	for _, v := range is	{
		total = total * v
	}
	return total
}

func t59Even(f func(is ...int) int, xs ...int)	int {
	var ys []int
	for _, v := range xs	{
		if v%2 == 0 {
			ys = append(ys, v)
		}
	}
	return f(ys...)
}

