package main

import "fmt"

func isEven(num int) bool{
    return num % 2 == 0
}

func main() {
    var number int
    //-------------------------
    fmt.Print("Enter A number:")
    fmt.Scanln(&number)
    //-------------------------
    fmt.Println(isEven(number))
}
