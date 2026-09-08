package main

import "fmt"

func main() {
    var num, result int
    /*------------*/
    fmt.Print("Enter A number:")
    fmt.Scanln(&num)    
    /*------------*/
    for i := 1;i <= 10;i++ {
        result = num * i
        fmt.Printf("%d x %d = %d\n", i, num, result)    
    }
}
