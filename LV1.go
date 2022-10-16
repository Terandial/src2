package main

import "fmt"

func main() {

	var a string

	fmt.Scanf("%s", &a)

	c := []rune(a)
	for b := 0; b < len(c)/2; b++ {
		if c[b] == c[len(c)-b-1] {
			fmt.Printf("%s", string(c[b]))
			continue
		} else {
			fmt.Println("false")
			break

		}
	}

}
