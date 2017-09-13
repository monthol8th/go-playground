package main

import "fmt"

func main(){
    mom := make(map[string]map[string]int)
    
    fmt.Println(mom)

    mom["a"] = make(map[string]int)
    mom["b"] = make(map[string]int)
    
    mom["a"]["a"] = 1
    mom["a"]["b"] = 2
    mom["b"]["a"] = 3
    mom["b"]["b"] = 4
    fmt.Println(mom)
}
