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

type stateFn func() (state, bool)

func main() {
	t := newTurnstyle()

	testEvents := []event{push, coin, push, push, coin, push, coin, coin, coin, push, push}
	for _, e := range testEvents {
		t.dispatch(e)
	}
	fmt.Println("Day's takings:", t.takings)
}

type turnstyle struct {
	takings int

	state   state
	machine [][]stateFn
}

func newTurnstyle() *turnstyle {
	t := &turnstyle{takings: 0, state: locked}
	t.machine = [][]stateFn{
		{t.lockedPush, t.lockedCoin},
		{t.unlockedPush, t.unlockedCoin},
	}
	return t
}

func (t *turnstyle) dispatch(e event) {
	fn := t.machine[t.state][e]
	next, transition := fn()
	if transition {
		t.state = next
	}
}

func (t *turnstyle) lockedPush() (state, bool) {
	sideEffect("ouch")
	return 0, false
}
func (t *turnstyle) lockedCoin() (state, bool) {
	sideEffect("green light")
	t.takings = t.takings + 1
	return unlocked, true
}

func (t *turnstyle) unlockedPush() (state, bool) {
	sideEffect("click")
	return locked, true
}
func (t *turnstyle) unlockedCoin() (state, bool) {
	sideEffect("cha-ching")
	t.takings = t.takings + 1
	return 0, false
}

func sideEffect(s string) {
	fmt.Println(s)
}
