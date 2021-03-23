package main

import "fmt"

type event int

const (
	push event = iota
	coin event = iota
)

type state func(event) state

func main() {
	s := stateLocked

	testEvents := []event{push, coin, push, push, coin, push, coin, coin, coin, push, push}
	for _, e := range testEvents {
		s = dispatch(s, e)
	}
}

func dispatch(s state, e event) state {
	next := s(e)
	if next == nil {
		return s
	}
	return next
}

func stateLocked(e event) state {
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
func stateUnlocked(e event) state {
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
