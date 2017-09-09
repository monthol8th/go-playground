package main

import "fmt"

func main() {
    var a [10]int

    for i := range a {
        a[i] = i
        fmt.Println(a[i])
    }

}
