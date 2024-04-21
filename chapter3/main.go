//string

package main

import (
	"fmt"
)

func main()  {
	fmt.Println(len("Hello,World")) //returns the number  bytes in a string, So, it will output 11, which is the number of characters in "Hello,World".
	fmt.Println("Hello,world"[1]) // This line prints the byte (ASCII value)
	fmt.Println("Hello, "+"world") //concatenation
}