package main

import (
	"fmt"
)

func main()	{
	res := fact(10)
	fmt.Println(res)
}

func fact(n int) int	{
	if n == 1 {
		return n
	}
	return n * fact(n-1)
}

