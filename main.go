package main

import "fmt"

func main () {
	atoz := "the quick brown fox jumps over the lazy dog\n"

	fmt.Printf("%s\n", atoz[0:9])
	fmt.Printf("%s\n", atoz[:9])  // Default to the beginning of the string
	fmt.Printf("%s\n", atoz[15:19])  // Default to the beginning of the string
	fmt.Printf("%s\n", atoz[15:])  // Take me to the end

	//for i, r := range atoz {
	//	fmt.Printf("index: %d \tCharacter: %c \n", i, r)
	//}

	for _, r := range atoz {
		fmt.Printf("\tCharacter: %c \n", r)
	}
	fmt.Printf("length of ayoz: %d\n", len(atoz))

	// BAckquote for double quotes
	atoz2 := `the quick brown "fox""" jumps over the lazy dog\n`
	fmt.Printf("%s", atoz2)

}
