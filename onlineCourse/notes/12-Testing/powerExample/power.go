package main

import	(
	"fmt"
)

func main()	{
	fmt.Println("1^3 =", power(1,3))
	fmt.Println("2^3 =", power(2,3))
	fmt.Println("3^3 =", power(3,3))
}


func power(a, b int) int	{
	result := 1
	for i := 0; i < b; i++	{
		result = result * a
	}
	return  result
}