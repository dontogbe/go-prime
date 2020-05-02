package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/dontogbe/go-prime/pkg/primeutils"
)

func main() {
	num1, _ := strconv.ParseInt(os.Args[1], 10, 64)
	num2, _ := strconv.ParseInt(os.Args[2], 10, 64)
	var numCPU = runtime.GOMAXPROCS(0)
	fmt.Println("Using", numCPU, "cores")
	var chanMax int64
	var c = make(chan *primeutils.Num, numCPU)
	timeout := time.After(5 * time.Second)

	fmt.Println("Finding prime numbers between", num1, "and ", num2)
	if num2 > num1 {
		chanMax = num2 - num1
		for num := num1; num <= num2; num++ {
			go primeutils.CheckIfPrime(num, c)
		}
		fmt.Println("Goroutines fired!!!")
	} else {
		chanMax = num1 - num2
		for num := num2; num <= num1; num++ {
			go primeutils.CheckIfPrime(num, c)
		}
		fmt.Println("Goroutines fired!!!")
	}
	for i := 0; i < int(chanMax); i++ {
		select {
		case n := <-c:
			if n.IsPrime {
				fmt.Println(n.Number, "is prime")
			}
		case <-timeout:
			continue
		}
	}
}
