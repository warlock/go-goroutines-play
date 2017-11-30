package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    c1 := make(chan float64, 1)
    c2 := make(chan int, 1)

    go func() {
        // simulate spending time to do work to get answer
        time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
        c1 <- 1.2
    }()

    go func() {
        time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
        c2 <- 34
    }()

    x := <-c1
    y := <-c2
    fmt.Println(x, y)
}