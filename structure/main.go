package main

import "fmt"

func main() {
    for i:=0; i<=10; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }

    j:=1;
    for j<=1024 {
        fmt.Println(j);
        j*=2;
    }
}
