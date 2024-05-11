// control statement using for loop and some conditions

package main

import (
	"fmt"
)

func main()  {
	var count = 10 
	for i := 0; i <= count; i++ {
		if i%2==0 {
			fmt.Println(i)
		}
		
	}
}