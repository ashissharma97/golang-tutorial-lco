package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(6) + 1
	fmt.Printf("Dice Number is %v \n", random)

	switch random {
	case 1:
		fmt.Println("We got 1 open")
	case 2:
		fmt.Println("We got 2")
	case 3:
		fmt.Println("We got 3")
	case 4:
		fmt.Println("We got 4")
	case 5:
		fmt.Println("We got 5")
		// fallthrough
	case 6:
		fmt.Println("We got 6 roll it again")
	default:
		fmt.Println("Not expected")
	}

}
