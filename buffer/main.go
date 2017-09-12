package main

import ("fmt";"time")

func main() {
    buff := make(chan int,3)
    buff <- 1
    buff <- 2
    buff <- 3
    go func(b chan int){
        for {
            fmt.Println(<-buff)
            time.Sleep(time.Second)
        }
    }(buff)
    time.Sleep(time.Second * 5)
    buff <- 4
    var input string
    fmt.Scanln(&input)
}
