package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

/*Num represents a number type that may be Prime.*/
type Num struct {
	number  float64
	isPrime bool
}

func main() {
	num1, _ := strconv.ParseFloat(os.Args[1], 64)
	num2, _ := strconv.ParseFloat(os.Args[2], 64)
	var chanMax int64
	var c chan bool

	fmt.Println("Finding prime numbers between", num1, "and ", num2)
	if num2 > num1 {
		chanMax = int64(num2 - num1)
		c = make(chan bool, chanMax)
		for num := num1; num <= num2; num++ {
			go checkIfPrime(num, c)

		}
	} else {
		chanMax = int64(num1 - num2)
		c = make(chan bool, chanMax)
		for num := num2; num <= num1; num++ {
			go checkIfPrime(num, c)
		}
	}
	for i := 0; i < int(chanMax); i++ {
		select {
		case <-c:
		}
	}
}

func checkIfPrime(num float64, c chan bool) {
	if num <= 1 {
		c <- false
	}
	if num == 2 || num == 3 {
		fmt.Println(num, "is prime")
		c <- true
	} else if int64(num)%2 == 0 {
		c <- false
	} else {
		limit := math.Sqrt(float64(num))
		for n := float64(3); n <= limit; n += 2 {
			if int64(num)%int64(n) == 0 {
				c <- false
				return
			}
		}
		fmt.Println(num, "is prime")
		c <- true
	}

}
