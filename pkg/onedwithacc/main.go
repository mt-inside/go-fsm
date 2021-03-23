package main

import "fmt"

type event int

const (
	push event = iota
	coin event = iota
)

type stateFn func(event) stateFn

func main() {
	t := newTurnstyle()

	testEvents := []event{push, coin, push, push, coin, push, coin, coin, coin, push, push}
	for _, e := range testEvents {
		t.dispatch(e)
	}
	fmt.Println("Day's takings:", t.takings)
}

type turnstyle struct {
	state   stateFn
	takings int
}

func newTurnstyle() *turnstyle {
	t := &turnstyle{takings: 0}
	t.state = t.stateLocked
	return t
}

func (t *turnstyle) dispatch(e event) {
	next := t.state(e)
	if next != nil {
		t.state = next
	}
}

func (t *turnstyle) stateLocked(e event) stateFn {
	switch e {
	case push:
		sideEffect("ouch")
		return nil
	case coin:
		sideEffect("green light")
		t.takings = t.takings + 1
		return t.stateUnlocked
	}
	panic("unknown event")
}
func (t *turnstyle) stateUnlocked(e event) stateFn {
	switch e {
	case push:
		sideEffect("click")
		return t.stateLocked
	case coin:
		sideEffect("cha-ching")
		t.takings = t.takings + 1
		return nil
	}
	panic("unknown event")
}

func sideEffect(s string) {
	fmt.Println(s)
}
