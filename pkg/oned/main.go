package main

import "fmt"

type event int

const (
	push event = iota
	coin event = iota
)

type stateFn func(event) stateFn

func main() {
	/* This example is called "1-D".
	 * We actually store the "state" as a pointer straight to the function that implements it.
	 * But we could have a state enum and a slice of the funcions (see 2-D) */
	s := stateLocked

	testEvents := []event{push, coin, push, push, coin, push, coin, coin, coin, push, push}
	for _, e := range testEvents {
		s = dispatch(s, e)
	}
}

func dispatch(s stateFn, e event) stateFn {
	next := s(e)
	if next == nil {
		return s
	}
	return next
}

func stateLocked(e event) stateFn {
	switch e {
	case push:
		sideEffect("ouch")
		return nil
	case coin:
		sideEffect("green light")
		return stateUnlocked
	}
	panic("unknown event")
}
func stateUnlocked(e event) stateFn {
	switch e {
	case push:
		sideEffect("click")
		return stateLocked
	case coin:
		sideEffect("cha-ching")
		return nil
	}
	panic("unknown event")
}

func sideEffect(s string) {
	fmt.Println(s)
}
