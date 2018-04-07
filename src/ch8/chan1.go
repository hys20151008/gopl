package main

import "fmt"

func main() {
    ch := make(chan string, 3)
    ch <- "A"
    ch <- "B"
    ch <- "C"

    fmt.Println(<-ch)
    fmt.Println(len(ch))
}



