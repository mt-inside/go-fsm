package main

import "fmt"

type event int

const (
	push event = iota
	coin event = iota
)

type state int

const (
	locked   state = iota
	unlocked state = iota
)

func main() {
	s := locked

	testEvents := []event{push, coin, push, push, coin, push, coin, coin, coin, push, push}
	for _, e := range testEvents {
		if s == locked {
			if e == push {
				sideEffect("ouch")
			} else if e == coin {
				sideEffect("green light")
				s = unlocked
			}
		} else if s == unlocked {
			if e == push {
				sideEffect("click")
				s = locked
			} else if e == coin {
				sideEffect("cha-ching")
			}
		}
	}
}

func sideEffect(s string) {
	fmt.Println(s)
}
