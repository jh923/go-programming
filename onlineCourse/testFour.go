/**
The tests can be found at https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#
*/

package main

import "fmt"

func main() {
	test4()
}

func test4() {
	t41 := [5]int{1, 2, 3, 4, 5}
	for i, v := range t41 {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", t41)

	fmt.Println()

	t42 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range t42 {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", t42)

	fmt.Println()

	t43 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(t43[0:5])
	fmt.Println(t43[5:])
	fmt.Println(t43[2:7])
	fmt.Println(t43[1:6])

	fmt.Println()

	t44 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	t44 = append(t44, 52)
	fmt.Println(t44)
	t44 = append(t44, 53, 54, 55)
	fmt.Println(t44)
	t44_2 := []int{56, 57, 58, 59, 60}
	t44 = append(t44, t44_2...) //... used to just get the values from the slice t44_2 making it of tye many ints rather than type slice
	fmt.Println(t44)

	fmt.Println()

	t45 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	t45 = append(t45[:3], t45[6:]...)
	fmt.Println(t45)

	fmt.Println()

	m46 := make([]string, 50, 50)
	m46 = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`,
		` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`,
		` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`,
		` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`,
		` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`,
		` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`,
		` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	for i := 0; i < len(m46); i++ {
		fmt.Println(i, m46[i])
	}
	fmt.Println("Length: ", len(m46), "Cap: ", cap(m46))

	fmt.Println()

	s47_1 := []string{"James", "Bond", "Shaken, not stirred"}
	s47_2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	mds47 := [][]string{s47_1, s47_2}
	fmt.Println(mds47)

	for i, xs := range mds47 {
		fmt.Println("Record ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}

	fmt.Println()

	m48 := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range m48 {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	fmt.Println()

	m48[`c_jack`] = []string{`Blond hair`, `Money`, `Anime`}
	for k, v := range m48 {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	fmt.Println()

	delete(m48, `c_jack`)
	for k, v := range m48 {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
