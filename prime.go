package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"time"
)

/*Num represents a number type that may be Prime.*/
type Num struct {
	number  int64
	isPrime bool
}

func main() {
	num1, _ := strconv.ParseInt(os.Args[1], 10, 64)
	num2, _ := strconv.ParseInt(os.Args[2], 10, 64)
	var numCPU = runtime.GOMAXPROCS(0)
	fmt.Println("Using", numCPU, "cores")
	var chanMax int64
	var c = make(chan *Num, numCPU)
	timeout := time.After(5 * time.Second)

	fmt.Println("Finding prime numbers between", num1, "and ", num2)
	if num2 > num1 {
		chanMax = num2 - num1
		for num := num1; num <= num2; num++ {
			go checkIfPrime(num, c)
		}
		fmt.Println("Goroutines fired!!!")
	} else {
		chanMax = num1 - num2
		for num := num2; num <= num1; num++ {
			go checkIfPrime(num, c)
		}
		fmt.Println("Goroutines fired!!!")
	}
	for i := 0; i < int(chanMax); i++ {
		select {
		case n := <-c:
			if n.isPrime {
				fmt.Println(n.number, "is prime")
			}
		case <-timeout:
			continue
		}
	}
}

func checkIfPrime(num int64, c chan *Num) {
	n := &Num{number: num, isPrime: false}
	if num <= 1 {
		c <- n
		return
	}
	if num == 2 || num == 3 {
		n.isPrime = true
		c <- n
	} else if num%2 == 0 {
		c <- n
	} else {
		limit := math.Sqrt(float64(num))
		for i := int64(3); i <= int64(limit); i += 2 {
			if num%i == 0 {
				c <- n
				return
			}
		}
		n.isPrime = true
		c <- n
	}

}
