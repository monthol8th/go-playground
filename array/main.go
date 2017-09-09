package main

import "fmt"

func main() {
    var a [10]int

    for i := range a {
        a[i] = i
        fmt.Println(a[i])
    }

    var slicea []int
    slicea = a[:]

    for _,x := range slicea {
        fmt.Println(x)
    }

    var x map[string]int

    x = make(map[string]int)
    x["plan"]=1
    fmt.Println(x["plan"])
    _,ok := x["plan2"]
    fmt.Println(ok)
}
