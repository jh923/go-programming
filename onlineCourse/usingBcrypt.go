package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main()	{
	//Encrypt a password into a hash
	pass := `myEpicPassWord:)`
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost) //min cost is 4. Used to save runtime, max is 31
	if err != nil	{
		fmt.Println(err)
	}
	fmt.Println(bs)
	fmt.Println(pass)


	err = nil

	//Check if a password matches a hash
	err = bcrypt.CompareHashAndPassword(bs, []byte(pass))
	if err != nil	{
		fmt.Println("You can't login")
		return
	}
	fmt.Println("You logged in")
}