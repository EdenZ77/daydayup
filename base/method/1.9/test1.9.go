package main

type Interface1 interface {
	M1(int, string) error
}

type Interface2 interface {
	M1(int, string) error
	M2(string2 string)
}

type Interface3 interface {
	Interface1
	Interface2 // Error: duplicate method M1
}

type Interface4 interface {
	Interface2
	M2(string) // Error: duplicate method M2
}

func main() {
}
