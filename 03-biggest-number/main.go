package main

import "fmt"

func main() {
    var num1, num2, num3 int
    /*--------------*/
    fmt.Print("Enter Number 1:")
    fmt.Scanln(&num1)
    fmt.Print("Enter Number 2:")
    fmt.Scanln(&num2)
    fmt.Print("Enter Number 3:")
    fmt.Scanln(&num3)
    /*--------------*/
    biggest := num1
    
    if num2 > biggest {
        biggest = num2    
    }
    if num3 > biggest {
        biggest = num3
    }
    fmt.Print("\n=======================")
    fmt.Printf("\nThe Biggest Number Is %d\n", biggest)
    fmt.Print("=======================\n\n")
}
