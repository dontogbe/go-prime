package primeutils

import (
	"math"
)

/*Num represents a number type that may be Prime.*/
type Num struct {
	Number  int64
	IsPrime bool
}

/*CheckIfPrime performs the actual prime validation.*/
func CheckIfPrime(num int64, c chan *Num) {
	n := &Num{Number: num, IsPrime: false}
	if num <= 1 {
		c <- n
		return
	}
	if num == 2 || num == 3 {
		n.IsPrime = true
		c <- n
		return
	}
	if num%2 == 0 {
		c <- n
		return
	}
	limit := math.Sqrt(float64(num))
	for i := int64(3); i <= int64(limit); i += 2 {
		if num%i == 0 {
			c <- n
			return
		}
	}
	n.IsPrime = true
	c <- n
}
