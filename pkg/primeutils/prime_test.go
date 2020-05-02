package primeutils

import (
	"fmt"
	"testing"
)

func TestPrime(t *testing.T) {
	c := make(chan *Num)
	go CheckIfPrime(3, c)
	num := <-c
	if num.IsPrime != true {
		t.Error("Was expecting true")
	}
	fmt.Println(num.Number)
}
