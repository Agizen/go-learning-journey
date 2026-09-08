package main

import "fmt"

func isPrime(number int) bool{
    if number <= 1 {
        return false    
    } else {
        for i := 2;i <= number;i++ {
            if number % i == 0 {
                return false           
            } else {
                return true            
            }
        }
    }
}

func main() {
    var num int
    //-----
    fmt.Print("Enter Your Number:")
    fmt.Scanln(&num)
    //-----
}
