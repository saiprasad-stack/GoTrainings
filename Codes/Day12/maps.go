package main

import "fmt"

func main() {
	fmt.Println("Maps in GO")

	// Contact Application -- Arvin - 12345567
	// var Map1 map[string]int
	// _ = Map1
	// fmt.Println(Map1["123"])
	// Map1["123"] = 123445

	// make()
	// contactBook := make(map[string]int)
	// contactBook["Arvin"] = 123456677
	// fmt.Println(contactBook)

	// delete a key
	// delete(contactBook, "Arvin")
	// fmt.Println(contactBook)

	contactBook := map[string]int{
		"Arvin":  1234,
		"Golang": 1997,
	}
	// _ = contactBook

	for key, value := range contactBook {
		fmt.Println(key, "-", value)
	}

	for key := range contactBook {
		fmt.Println("Value: ", key)
	}
	for _, value := range contactBook {
		fmt.Println("Value: ", value)
	}

	// length of map
	fmt.Println(len(contactBook))

	// comma , ok syntax
	key1, ok := contactBook["Arvin"]
	if ok {
		fmt.Println("it exists", key1)
	} else {
		fmt.Println("No key found", key1)
	}

	// type Student struct {
	// 	id  int
	// 	age int
	// }

	// contactBook2 = map[Student]string{
	// 	{id: 12, age: 34}: "India",
	// 	{id: 24, age: 45}: "Shri Lanka",
	// }
	// _ = contactBook2

	//

	hobbyMap := make(map[int][]string)
	hobbyMap[1] = []string{"Cricket", "Hockey", "Tennis"}
	hobbyMap[2] = []string{"Singing", "Dance"}

	// map2 := make(map[string]map[int]string)
	// map2 {"City":{1:"Gandhi Nagar"}}
	// map2["City"][1]

}
