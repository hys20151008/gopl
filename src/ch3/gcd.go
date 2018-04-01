package main

import "fmt"

func gcd(x,y int) int {
    if y != 0 {
	x, y = y, x%y
    }
    return x
}

func main() {
    fmt.Println(gcd(2,3))
}
