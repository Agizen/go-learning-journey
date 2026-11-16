package main

import (
	"fmt"
	"math"
)

func CalculateHypotenuse(SideA, SideB float64) {
	result := math.Sqrt((SideA * SideA) + (SideB * SideB))
	fmt.Println("===== Input Results =====")
	fmt.Printf("Side A: %.2f\n", SideA)
	fmt.Printf("Side B: %.2f\n", SideB)
	fmt.Println("\n===== Calculation Results =====")
	fmt.Printf("Hypotenuse: %.2f\n", result)
}

func main() {
	var sideA, sideB float64
	fmt.Print("Enter Side A: ")
	_, err := fmt.Scanln(&sideA)
	if err != nil {
		fmt.Println("Invalid Input!")
		return
	}
	if sideA <= 0 {
		fmt.Println("Side A Should Be A Positive And Bigger than 0!")
		return
	}
	fmt.Print("Enter Side B: ")
	_, err1 := fmt.Scanln(&sideB)
	if err1 != nil {
		fmt.Println("Invalid Input!")
		return
	}
	if sideB <= 0 {
		fmt.Println("Side B Should Be A Positive And Bigger than 0!")
		return
	}
	CalculateHypotenuse(sideA, sideB)
}
