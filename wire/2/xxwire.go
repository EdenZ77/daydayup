//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitEndingA(name string) EndingA {
	wire.Build(NewMonster, NewPlayer, NewEndingA)
	return EndingA{}
}

func InitEndingB(name string) EndingB {
	wire.Build(NewMonster, NewPlayer, NewEndingB)
	return EndingB{}
}
