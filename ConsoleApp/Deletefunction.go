package main

import (
	"fmt"
	"log"
)

func deleteFromMap(y map[int]string) {
	fmt.Println("print num beebox:")
	var keyForMap int
	//	errors for key input
	_, err := fmt.Scan(&keyForMap)
	if err != nil {
		log.Fatal("the input num is not number:")
	}
	delete(y, keyForMap)
}
