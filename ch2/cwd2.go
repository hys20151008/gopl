package main

import (
    "os"
    "fmt"
)

//var cwd string

func main() {
    cwd, err := os.Getwd()
    if err != nil {
	fmt.Println(err)
    }
    fmt.Println(cwd)
}
