package main

import "fmt"

func Counter(out chan<- int){
    for i :=0; i < 100; i++ {
        out <- i
    }
    close(out)
}

func Squarer(in <-chan int, out chan<- int) {
    for x:= range in {
	out <- x * x
    }
    close(out)
}

func Printer(in <-chan int){
    for x := range in {
        fmt.Println(x)
    }
}

func main() {
    naturals := make(chan int)
    squares := make(chan int)

    go Counter(naturals)
    go Squarer(naturals, squares)
    Printer(squares)
}
