package main

import "fmt"

func avg(input []int) float64 {
    sum := 0.0

    for _, x := range input {
        sum += float64(x)
    }
    return sum/float64(len(input))
}

func spreadit(input ...int) {
    for _,x := range input {
        fmt.Println(x)
    }
}

func generator(multipler int) (func (int) int) {
    return func (x int) int {
        return multipler*x
    }
}

func incrementgen() (func () int) {
    i:=0
    return func () int {
        ret := i
        i++
        return ret
    }
}

func end(text string) {
    fmt.Println(text)
}

func main() {

    defer end("eiei")
    x := [10]int{1,2,3,4,5,6,7,8,9,10}
    fmt.Println(avg(x[:]))
    spreadit(x[:]...)
    mul1:=generator(2)
    mul2:=generator(4)

    fmt.Println(mul1(10),mul2(10),mul1(10))
    inc := incrementgen()
    inc2 := incrementgen()
    fmt.Println(inc2())
    fmt.Println(inc())
    fmt.Println(inc())
}
