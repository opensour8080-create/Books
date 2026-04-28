package main

import "fmt"

func main() {
    x := 8           // или var x int = 8
    var p *int       // указатель на int
    p = &x           // p хранит адрес x
    fmt.Println(p)   // выведет адрес (например 0xc0000140a0)
    fmt.Println(*p)  // выведет значение по адресу (8)
}