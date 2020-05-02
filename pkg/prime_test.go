package primeutils

import (
	"fmt"
	"testing"
)

func TestPrime(t *testing.T) {
	c := make(chan *Num)
	go checkIfPrime(3, c)
	num := <-c
	if num.isPrime != true {
		t.Error("Was expecting true")
	}
	fmt.Println(num.number)
}
