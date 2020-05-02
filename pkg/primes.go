package primeutils

import (
	"math"
)

/*Num represents a number type that may be Prime.*/
type Num struct {
	number  int64
	isPrime bool
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
