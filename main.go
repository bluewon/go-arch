package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	First string
}

func main() {
	// p1 := Person{
	// 	First: "Jenny",
	// }
	// p2 := Person{
	// 	First: "Yuuminn",
	// }
	// xp := []Person{p1, p2}
	// bs, err := json.Marshal(xp)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(string(bs))
	// xp2 := []Person{}

	// err2 := json.Unmarshal(bs, &xp2)
	// if err2 != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println("back into data structure ", xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := Person{
		First: "Jenny",
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encode bad data", err)
	}
}
func bar(w http.ResponseWriter, r *http.Request) {
	var p1 Person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("decode bad data",err)
	}

	fmt.Println("Person:",p1)
}
