package main

import (
    "fmt"
)

func main() {
    var f, c int

    fmt.Println("Enter the value of Celsius:")
    fmt.Scanf("%d", &c)

    f = (c * 9 / 5) + 32

    fmt.Println("Value in Fahrenheit:", f)
}
