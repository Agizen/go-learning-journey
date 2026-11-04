package main

import "fmt"

func main() {
	var (
		x    int
		sign string
		y    int
	)
	for {

		fmt.Print("Enter The X:")
		fmt.Scanln(&x)
		fmt.Print("Enter The Sign(+, -, *, /):")
		fmt.Scanln(&sign)
		fmt.Print("Enter The Y:")
		fmt.Scanln(&y)

		switch sign {
		case "+":
			fmt.Println("===================")
			fmt.Printf("%v + %v = %v\n", x, y, x+y)
			fmt.Println("===================")
		case "-":
			fmt.Println("===================")
			fmt.Printf("%v - %v = %v\n", x, y, x-y)
			fmt.Println("===================")
		case "*":
			fmt.Println("===================")
			fmt.Printf("%v * %v = %v\n", x, y, x*y)
			fmt.Println("===================")
		case "/":
			fmt.Println("===================")
			fmt.Printf("%v / %v = %v\n", x, y, x/y)
			fmt.Println("===================")
		}
	}
}
