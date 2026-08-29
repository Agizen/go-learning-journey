/*----------------------------------*/
/*    Agid Khezri Go  Developer     */
/*----------------------------------*/

package main

import "fmt"

func main() {
    var num int
    sum := 0
    //------------
    fmt.Print("Enter Your Number:")
    fmt.Scanln(&num)
    //------------
    for i := 1;i <= num;i++ {
        sum += i
    }
    //------------
    fmt.Println(sum)
}
