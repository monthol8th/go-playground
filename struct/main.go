package main

import ("fmt" ; "math")

type Circle struct {
    x float64
    y float64
    r float64
}

type NamedCircle struct {
    Circle
    name string
}

func circleArea(c Circle) float64 {
    return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r
}

func area(c Circle) float64 {
    return math.Pi * c.r *c.r
}

func main() {
    c := Circle{x: 0, y: 0, r: 5}
    fmt.Println(circleArea(c))
    fmt.Println(c.area())
    fmt.Println(area(c))

    nc:= new(NamedCircle)
    nc.r = 1
    fmt.Println(nc.Circle.area())
    fmt.Println(nc.area())
}
