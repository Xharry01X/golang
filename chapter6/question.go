package main

import (
    "fmt"
)

// 1m=3.28ft
func main() {
    var m int
    fmt.Println("Enter a number in meters:")
    fmt.Scanf("%d", &m)

    f := float64(m) / 3.28 // Convert m to float64 to perform floating-point division

    fmt.Println("Value in feet:", f)
}
