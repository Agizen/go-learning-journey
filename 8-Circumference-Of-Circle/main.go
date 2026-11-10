package main

import "fmt"

func QxPI(qatar float64) {
	fmt.Println("=============")
	fmt.Printf("%v x 3.14 = %v\n", qatar, qatar*3.14)
	fmt.Println("=============")
}

func RxPIx2(radius float64) {
	fmt.Println("=============")
	fmt.Printf("%v x 3.14 x 2 = %v\n", radius, radius*3.14*2)
	fmt.Println("=============")
}

func main() {
	for {
		var choise int
		fmt.Println("1-Calculating by (Qatar x PI).")
		fmt.Println("2-Calculating by (Radius x PI x 2).")
		fmt.Print("Enter Your Choise: ")
		fmt.Scanln(&choise)

		switch choise {
		case 1:
			var qatar float64
			fmt.Print("Enter Your Qatar(5.78, etc.): ")
			fmt.Scanln(&qatar)
			QxPI(qatar)
		case 2:
			var radius float64
			fmt.Print("Enter Your Radius(5.8, etc.): ")
			fmt.Scanln(&radius)
			RxPIx2(radius)
		}
	}
}
