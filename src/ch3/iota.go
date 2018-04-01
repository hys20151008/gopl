package main

import "fmt"


const (
    FlagUp uint = 1 << iota
    FlagBroadcast
    FlagLoopback
)

func main() {
    fmt.Printf("%u\t%u\t%u", FlagUp, FlagBroadcast, FlagLoopback)
}
