package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jailtonjunior94/go-tags/enums"
	"github.com/jailtonjunior94/go-tags/validations"
)

type Person struct {
	ID         int              `json:"id" validate:"required"`
	Name       string           `json:"name" validate:"required"`
	Email      string           `json:"email" validate:"required,email"`
	PersonType enums.PersonType `json:"type" validate:"required,person_type"`
}

func main() {
	person := Person{ID: 1, Name: "John Doe", Email: "john.doe@example.com", PersonType: "Unknown"}
	err := validations.Validations(person)
	if err != nil {
		log.Fatal(err)
	}

	uJson, _ := json.Marshal(person)
	fmt.Println(string(uJson))
}
