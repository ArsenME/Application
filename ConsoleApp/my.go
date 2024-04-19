package main

import (
	"fmt"
	"log"
)

func main() {

	for {

		bees := data()
		var want string
		fmt.Println("how do you want?:")
		//errors
		_, err := fmt.Scanf("%s", &want)
		if err != nil {
			log.Fatal("error reading input")
		}
		switch want {
		case "add":
			fmt.Println("print num for search beebox:")
			var keyForMap int
			//	errors for key input
			_, err := fmt.Scan(&keyForMap)
			if err != nil {
				log.Fatal("the input num is not number:")
			}
			addInMap(keyForMap, bees)
		case "delete":
			deleteFromMap(bees)

		case "watch":
			fmt.Println(bees)

		case "stop":
			break
		default:
			break
		}
		fmt.Println(bees)

	}
}
