package main

import (
	"reflect"
	"testing"
)

type FTest struct {
	s     any
	name  string
	index []int
	value int
}

type D1 struct {
	d int
}
type D2 struct {
	d int
}

type S0 struct {
	A, B, C int
	D1
	D2
}

type S1 struct {
	B int
	S0
}

type S2 struct {
	A int
	*S1
}

type S1x struct {
	S1
}

type S1y struct {
	S1
}

type S3 struct {
	S1x
	S2
	D, E int
	*S1y
}

type S4 struct {
	*S4
	A int
}

// The X in S6 and S7 annihilate, but they also block the X in S8.S9.
type S5 struct {
	S6
	S7
	S8
}

type S6 struct {
	X int
}

type S7 S6

type S8 struct {
	S9
}

type S9 struct {
	X int
	Y int
}

// The X in S11.S6 and S12.S6 annihilate, but they also block the X in S13.S8.S9.
type S10 struct {
	S11
	S12
	S13
}

type S11 struct {
	S6
}

type S12 struct {
	S6
}

type S13 struct {
	S8
}

// The X in S15.S11.S1 and S16.S11.S1 annihilate.
type S14 struct {
	S15
	S16
}

type S15 struct {
	S11
}

type S16 struct {
	S11
}

var fieldTests = []FTest{
	//{struct{}{}, "", nil, 0},
	//{struct{}{}, "Foo", nil, 0},
	//{S0{A: 'a'}, "A", []int{0}, 'a'},
	//{S0{}, "D", nil, 0},
	//{S1{S0: S0{A: 'a'}}, "A", []int{1, 0}, 'a'},
	//{S1{B: 'b'}, "B", []int{0}, 'b'},
	//{S1{}, "S0", []int{1}, 0},
	//{S1{S0: S0{C: 'c'}}, "C", []int{1, 2}, 'c'},
	//{S2{A: 'a'}, "A", []int{0}, 'a'},
	//{S2{}, "S1", []int{1}, 0},
	//{S2{S1: &S1{B: 'b'}}, "B", []int{1, 0}, 'b'},
	//{S2{S1: &S1{S0: S0{C: 'c'}}}, "C", []int{1, 1, 2}, 'c'},
	//{S2{}, "D", nil, 0},
	//{S3{}, "S1", nil, 0},
	//{S3{S2: S2{A: 'a'}}, "A", []int{1, 0}, 'a'},
	//{S3{}, "B", nil, 0},
	//{S3{D: 'd'}, "D", []int{2}, 0},
	//{S3{E: 'e'}, "E", []int{3}, 'e'},
	//{S4{A: 'a'}, "A", []int{1}, 'a'},
	//{S4{}, "B", nil, 0},
	//{S5{}, "X", nil, 0},
	//{S5{}, "Y", []int{2, 0, 1}, 0},
	//{S10{}, "X", nil, 0},
	{S10{}, "Y", []int{2, 0, 0, 1}, 0},
	//{S14{}, "X", nil, 0},
}

func TestFieldByIndex(t *testing.T) {
	for _, test := range fieldTests {
		s := reflect.TypeOf(test.s)
		f := s.FieldByIndex(test.index)
		if f.Name != "" {
			if test.index != nil {
				if f.Name != test.name {
					t.Errorf("%s.%s found; want %s", s.Name(), f.Name, test.name)
				}
			} else {
				t.Errorf("%s.%s found", s.Name(), f.Name)
			}
		} else if len(test.index) > 0 {
			t.Errorf("%s.%s not found", s.Name(), test.name)
		}

		if test.value != 0 {
			v := reflect.ValueOf(test.s).FieldByIndex(test.index)
			if v.IsValid() {
				if x, ok := v.Interface().(int); ok {
					if x != test.value {
						t.Errorf("%s%v is %d; want %d", s.Name(), test.index, x, test.value)
					}
				} else {
					t.Errorf("%s%v value not an int", s.Name(), test.index)
				}
			} else {
				t.Errorf("%s%v value not found", s.Name(), test.index)
			}
		}
	}
}

func TestFieldByName(t *testing.T) {
	for _, test := range fieldTests {
		s := reflect.TypeOf(test.s)
		f, found := s.FieldByName(test.name)
		if found {
			if test.index != nil {
				// Verify field depth and index.
				if len(f.Index) != len(test.index) {
					t.Errorf("%s.%s depth %d; want %d: %v vs %v", s.Name(), test.name, len(f.Index), len(test.index), f.Index, test.index)
				} else {
					for i, x := range f.Index {
						if x != test.index[i] {
							t.Errorf("%s.%s.Index[%d] is %d; want %d", s.Name(), test.name, i, x, test.index[i])
						}
					}
				}
			} else {
				t.Errorf("%s.%s found", s.Name(), f.Name)
			}
		} else if len(test.index) > 0 {
			t.Errorf("%s.%s not found", s.Name(), test.name)
		}

		if test.value != 0 {
			v := reflect.ValueOf(test.s).FieldByName(test.name)
			if v.IsValid() {
				if x, ok := v.Interface().(int); ok {
					if x != test.value {
						t.Errorf("%s.%s is %d; want %d", s.Name(), test.name, x, test.value)
					}
				} else {
					t.Errorf("%s.%s value not an int", s.Name(), test.name)
				}
			} else {
				t.Errorf("%s.%s value not found", s.Name(), test.name)
			}
		}
	}
}
