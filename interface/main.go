package main

import "fmt"

type Something interface {
    callName() string
}

type A struct {
    name string
}

type B struct {
    name,surname string
}

type C struct {
    name,surname string
}

type D struct{
    name string
}


func (a *A) callName() string {
    return a.name
}

func (b *B) callName() string {
    return b.name+" "+b.surname
}

func (c *C) callName() string {
    c.surname = c.surname+"+"
    return c.surname+" "+c.name
}

func (d D) callName() string {
    return d.name
}

func call(s Something) {
    fmt.Println(s.callName())
}

func main() {
    a := A{"qwertyuio"}
    b := B{"qwe","asd"}
    c := C{"qwe","asd"}
    d := D{"kuy"}

    call(&a)
    call(&b)
    call(&c)
   
    call(d)

    fmt.Println("Oh,",a.callName())

    s := []Something{&a,&b,&c,d}
    for _,something := range s {
        call(something)
    }
}
