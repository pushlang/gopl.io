package main

import "fmt"

type stateAPI1 interface {
	prettyName() string
}
type stateAPI2 interface {
	dumbName() string
}
type stateAPI12 interface {
	stateAPI1
	stateAPI2
}

func (s *stateP) prettyName() string {
	return s.pName()
}
func (s *stateD) dumbName() string {
	return s.dName()
}
// func (s *statePD) prettyName() string {
// 	return s.pName()
// }
// func (s *statePD) dumbName() string {
// 	return s.dName()
// }

type stateP struct {
	pn    int
	pName func() string
}

type stateD struct {
	pd    int
	dName func() string
}

type statePD struct {
	*stateP
	*stateD
}

type state struct {
	stateAPI12
}

func main() {
	var st1 stateAPI1
	var st2 stateAPI2
	var st12 stateAPI12

	a := &stateP{0, func() string { return "State prettyName!" }}
	b := &stateD{0, func() string { return "State dumbName!" }}
	c := &state{&statePD{a, b}}

	st1 = a
	fmt.Println(st1.prettyName())
	st2 = b
	fmt.Println(st2.dumbName())
	st12 = c
	fmt.Println(st12.dumbName(), st12.prettyName())

}
