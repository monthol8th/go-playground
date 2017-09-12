package main

import ("fmt";"time")

func wait(data,quit chan int) {
    fmt.Println("wait")
    for{
        select {
            case x:= <- data:
                fmt.Println(x)
            case <- quit:
                fmt.Println("finish")
                return
        }
    }
}

func sender(data,quit chan int) {
    fmt.Println("start")
    for i:=0; i<10; i++ {
        fmt.Println("send",i)
        data<- i
        time.Sleep(time.Second * 1)
    }
    quit <- 0
}

func main() {
    var data,quit chan int
    data = make(chan int)
    quit = make(chan int)
    go sender(data, quit)
    wait(data, quit)

    var input string
    fmt.Scanln(&input)
}
