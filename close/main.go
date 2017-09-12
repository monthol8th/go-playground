package main

import ("fmt";"time")

func sender(c chan int){
    for i := 0; i < 10; i++ {
        c <- i
        time.Sleep(time.Second)
    }
    close(c)
}

func reciever(c chan int){
    for v := range c{
            fmt.Println(v)
    }
}

func main() {
    c := make(chan int)
    go sender(c)
    go reciever(c)
    var input string
    fmt.Scanln(&input)
}
