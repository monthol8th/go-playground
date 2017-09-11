package main

import "fmt"

func main() {
    arr := [5]int{1,2,3,4,5}
    fmt.Println(arr)
    var x *int
    x = &arr[0]
    fmt.Println(*x)
    *x -= 1
    fmt.Println(*x)
}
