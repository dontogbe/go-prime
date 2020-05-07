package primeutils

import (
	"testing"
)

func TestPrime(t *testing.T) {
	tests := map[string]struct {
		num  int64
		ch   chan *Num
		want bool
	}{
		"negative": {num: -9, ch: make(chan *Num), want: false},
		"odd":      {num: 999, ch: make(chan *Num), want: false},
		"even":     {num: 1000, ch: make(chan *Num), want: false},
		"square":   {num: 49, ch: make(chan *Num), want: false},
		"prime":    {num: 99929, ch: make(chan *Num), want: true},
		"two":      {num: 2, ch: make(chan *Num), want: true},
		"three":    {num: 3, ch: make(chan *Num), want: true},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			go CheckIfPrime(tc.num, tc.ch)
			num := <-tc.ch
			if num.IsPrime != tc.want {
				t.Fatalf("Checked %#v for prime. Was expecting %#v but got %#v", tc.num, tc.want, num.IsPrime)
			}

		})
	}

}
