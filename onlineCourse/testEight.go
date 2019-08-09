/**
The tests can be found at https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type user struct {
	First string
	Age   int
}

type user83 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user83

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user83

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	//test81()
	//test82()
	//test83()
	//test84()
	test85()
}

func test81() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	userCollection := []user{u1, u2, u3}
	b, err := json.Marshal(userCollection)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func test82() {
	str := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(str)

	var people []person

	err := json.Unmarshal([]byte(str), &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for i, person := range people {
		fmt.Println("Person number", i)
		fmt.Printf("\t %s %s aged %d, likes to say:\n", person.First, person.Last, person.Age)
		for _, saying := range person.Sayings {
			fmt.Println("\t\t", saying)
		}
	}
}

func test83() {
	u1 := user83{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user83{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user83{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user83{u1, u2, u3}

	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did something wrong and here's the error:", err)
	}
}

func test84() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry",
		"fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

func test85() {
	u1 := user83{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user83{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user83{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user83{u1, u2, u3}

	//fmt.Println(users)

	fmt.Println("----------------------------------------")
	fmt.Println("Only sort sayings")
	fmt.Println("----------------------------------------")
	for _, user := range users {
		sort.Strings(user.Sayings)
		fmt.Printf("%s %s, aged: %d, likes to say:\n", user.First, user.Last, user.Age)
		for _, saying := range user.Sayings {
			fmt.Printf("\t%s\n", saying)
		}
	}

	fmt.Println("----------------------------------------")
	fmt.Println("Sort saying and by age")
	fmt.Println("----------------------------------------")
	sort.Sort(ByAge(users))
	for _, user := range users {
		sort.Strings(user.Sayings)
		fmt.Printf("%s %s, aged: %d, likes to say:\n", user.First, user.Last, user.Age)
		for _, saying := range user.Sayings {
			fmt.Printf("\t%s\n", saying)
		}
	}

	fmt.Println("----------------------------------------")
	fmt.Println("Sort by sayings and by last name")
	fmt.Println("----------------------------------------")
	sort.Sort(ByLast(users))
	for _, user := range users {
		sort.Strings(user.Sayings)
		fmt.Printf("%s %s, aged: %d, likes to say:\n", user.First, user.Last, user.Age)
		for _, saying := range user.Sayings {
			fmt.Printf("\t%s\n", saying)
		}
	}

}
