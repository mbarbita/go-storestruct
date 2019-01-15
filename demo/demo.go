package main

import (
	"log"
	"time"

	storestruct "github.com/mbarbita/golib-storestruct"
)

type obj struct {
	Name string
	Age  int
	When time.Time
}

func main() {
	o1 := &obj{
		Name: "BMM",
		Age:  45,
		When: time.Now(),
	}

	//Save
	if err := storestruct.Save("./demo.txt", o1); err != nil {
		log.Fatalln(err)
	}
	// load
	var o2 = new(obj)
	if err := storestruct.Load("./demo.txt", o2); err != nil {
		log.Fatalln(err)
	}
	log.Println(o2)

}
