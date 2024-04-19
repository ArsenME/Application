package main

import (
	"fmt"
	"log"
)

func addInMap(x int, y map[int]string) {
	fmt.Println("input subjects: ")
	s := ""
	_, err := fmt.Scanln(&s)
	if err != nil {
		log.Fatal("Error reading input:", err)
	}
	y[x] = s

}
