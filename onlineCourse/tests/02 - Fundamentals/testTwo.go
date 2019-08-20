/**
The tests can be found at https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#
*/

package two

import "fmt"

func main() {
	test2()
}

func test2() {
	//test 2.1 - print a number in base 10, base 2 and base 16
	num := 100
	fmt.Printf("%d\t%b\t%X\n", num, num, num)

	//test 2.2 - comparing values
	a := 10
	b := 100
	g := a == b
	h := a <= b
	j := a != b
	k := a < b
	l := a > b
	fmt.Println(g, h, j, k, l)

	//test 2.3 - constants
	const t23 = 23
	const t231 int = 23
	fmt.Printf("%v\t%v\n", t23, t231)

	//test 2.4 - bit shift
	a = 7
	fmt.Printf("%d\t%b\t%X\n", a, a, a)
	a = a << 1
	fmt.Printf("%d\t%b\t%X\n", a, a, a)

	//test 2.5 - raw string literal
	c := `Literal string`
	fmt.Println(c)

	//test 2.6 - iota
	const (
		_  = 2019 + iota
		n1 = 2019 + iota
		n2 = 2019 + iota
		n3 = 2019 + iota
		n4 = 2019 + iota
	)
	fmt.Println(n1, n2, n3, n4)
}
