package main

import "fmt"

type A struct{
    q int
    w int
    e int
}

func create() A {
    return A{1,2,3}
}

func (a A) sum() int {
    return a.q+a.w+a.e
}

type B interface {
    sum() int
}

func main() {
    var b B
    b = create()
    c := b
    fmt.Println(b.sum())
    fmt.Println(c.sum())
    fmt.Println(b.sum())
}
