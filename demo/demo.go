package main

import (
	"log"
	"time"

	storestruct "github.com/mbarbita/golib-storestruct"
)

type obj1 struct {
	Name string
	Age  int
	When time.Time
}

type obj2 struct {
	Name    string
	ID      int
	Planets map[string]string
}

type obj3 struct {
	Name    string
	ID      int
	Planets map[string]string
	Visited bool
}

func main() {
	obj11 := &obj1{
		Name: "BMM",
		Age:  45,
		When: time.Now(),
	}

	//Save obj11
	if err := storestruct.Save("./obj1.txt", obj11); err != nil {
		log.Fatalln(err)
	}
	// load obj12
	var obj12 = new(obj1)
	if err := storestruct.Load("./obj1.txt", obj12); err != nil {
		log.Fatalln(err)
	}
	log.Print(obj12, "\n", "\n")

	var obj21 = new(obj2)
	obj21.Name = "Ra"
	obj21.ID = 1
	obj21.Planets = make(map[string]string)
	obj21.Planets["Earth"] = "Humans"
	obj21.Planets["Mars"] = "None"

	//Save obj21
	if err := storestruct.Save("./obj2.txt", obj21); err != nil {
		log.Fatalln(err)
	}
	// load obj22
	var obj22 = new(obj2)
	if err := storestruct.Load("./obj2.txt", obj22); err != nil {
		log.Fatalln(err)
	}
	log.Print(obj22, "\n", "\n")

	obj31 := &obj3{
		Name: "Fer",
		ID:   2,
		Planets: map[string]string{
			"asd": "asd",
			"sss": "sss",
		},
		Visited: false,
	}

	//Save obj31
	if err := storestruct.Save("./obj3.txt", obj31); err != nil {
		log.Fatalln(err)
	}
	// load obj32
	var obj32 = new(obj3)
	if err := storestruct.Load("./obj3.txt", obj32); err != nil {
		log.Fatalln(err)
	}
	log.Print(obj32, "\n", "\n")

}
