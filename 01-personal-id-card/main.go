/*-------------------------------*/
/*    Agid Khezri Go Develper    */
/*-------------------------------*/

package main

import "fmt"

func main() {
    var name string
    var age int
    var height float64
    var result string
    //------------------
    fmt.Print("What is Your Name:")
    fmt.Scanln(&name)
    fmt.Print("How Old Are You:")
    fmt.Scanln(&age)
    fmt.Print("How much are you tall:")
    fmt.Scanln(&height)
    //------------------
    if age % 2 == 0 {
        result = "Even"
    } else {
        result = "Odd"    
    }
    //------------------
    fmt.Println("==========================")
    fmt.Printf("Name: %s\nAge: %d\nHeight: %.1f cm\nIs Age Odd Or Even: %s\n", name, age, height, result)
    fmt.Println("==========================")
}
