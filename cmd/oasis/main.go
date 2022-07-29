package main

import (
	"fmt"

	"github.com/tkdailey11/oasis/pkg/db"
)

func main() {
	if result := db.Insert("Hello World"); result {
		fmt.Println("SUCCESS")
	} else {
		fmt.Println("FAILURE")
	}
}