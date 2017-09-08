package main

import "fmt";

var publicX int16;

func test() int16 {
    return 1 + publicX;
}

func main() {
    var x string;
    x = "plankung";
    fmt.Println(x);
    y:= 5;
    fmt.Println(y);
    y = 6;
    fmt.Println(y);

    publicX = 10;
    fmt.Println(test());

    var a,b,c=1,2,3
    fmt.Println(a,b,c)
    var (
        d=4
        e=5
        f=6
    )
    fmt.Println(d,e,f)
}
