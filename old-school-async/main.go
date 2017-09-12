package main

import "fmt"

func reasonOfLife(c chan int) {
    c <- 42
}

func main() {
    var c chan int
    var arr [10]int
    var sl []int
    c = make(chan int)
    go reasonOfLife(c)
    
    sl = arr[:]
    sl[0] = 1
    x := <- c

    fmt.Println(x)

}
