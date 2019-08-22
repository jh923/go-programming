package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	tpl.ParseFiles("exampleP2.gohtml")

	type Duelist struct {
		Name       string
		AceMonster string
	}

	type Trainer struct {
		Name        string
		MainPokemon string
	}

	type Childhood struct {
		Trainers []Trainer
		Duelists []Duelist
	}

	yugi := Duelist{
		Name:       "Yugi",
		AceMonster: "Dark Magician",
	}
	kaiba := Duelist{
		Name:       "Kaiba",
		AceMonster: "Blue eyes white dragon",
	}
	joe := Duelist{
		Name:       "Joe",
		AceMonster: "Red eyes black dragon",
	}

	ds := []Duelist{yugi, kaiba, joe}

	ash := Trainer{
		Name:        "Ash",
		MainPokemon: "Pikachu",
	}

	misty := Trainer{
		Name:        "Misty",
		MainPokemon: "Psyduck",
	}

	brock := Trainer{
		Name:        "Brock",
		MainPokemon: "Onix",
	}

	ts := []Trainer{ash, misty, brock}

	ch := Childhood{ts, ds}

	err := tpl.ExecuteTemplate(os.Stdout, "example.gohtml", ch)
	if err != nil {
		log.Fatal(err)
	}
}
