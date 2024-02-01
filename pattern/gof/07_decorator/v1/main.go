package main

import "fmt"

/*
参考资料：https://mp.weixin.qq.com/s?__biz=MzUzNzAzMTc3MA==&mid=2247484431&idx=1&sn=2a7cd975bd703b478efe62f971297253&chksm=faec613acd9be82c0839181df010d0fb1bb8ef42ffd7d6d42040c7a8b57a96cde6caa70d34fa&scene=178&cur_album_id=1908992469812199431#rd


*/

// Aircraft 飞行器接口，有fly和landing方法
type Aircraft interface {
	fly()
	landing()
}

// Helicopter 直升机结构体
type Helicopter struct{}

// NewHelicopter 创建一个新的Helicopter实例
func NewHelicopter() *Helicopter {
	return &Helicopter{}
}

func (h *Helicopter) fly() {
	fmt.Println("我是普通直升机")
}

func (h *Helicopter) landing() {
	fmt.Println("我有降落功能")
}

// WeaponAircraft 武装直升机
type WeaponAircraft struct {
	Aircraft
}

// NewWeaponAircraft 创建一个新的WeaponAircraft实例
func NewWeaponAircraft(aircraft Aircraft) *WeaponAircraft {
	return &WeaponAircraft{
		Aircraft: aircraft,
	}
}

func (a *WeaponAircraft) fly() {
	a.Aircraft.fly()
	fmt.Println("增加武装功能")
}

// RescueAircraft 救援直升机
type RescueAircraft struct {
	Aircraft
}

// NewRescueAircraft 创建一个新的RescueAircraft实例
func NewRescueAircraft(aircraft Aircraft) *RescueAircraft {
	return &RescueAircraft{
		Aircraft: aircraft,
	}
}

func (r *RescueAircraft) fly() {
	r.Aircraft.fly()
	fmt.Println("增加救援功能")
}

func main() {
	// 普通直升机
	fmt.Println("------------普通直升机")
	helicopter := NewHelicopter()
	helicopter.fly()
	helicopter.landing()

	// 武装直升机
	fmt.Println("------------武装直升机")
	weaponAircraft := NewWeaponAircraft(helicopter)
	weaponAircraft.fly()

	// 救援直升机
	fmt.Println("------------救援直升机")
	rescueAircraft := NewRescueAircraft(helicopter)
	rescueAircraft.fly()

	// 武装救援直升机
	fmt.Println("------------武装救援直升机")
	weaponRescueAircraft := NewRescueAircraft(weaponAircraft)
	weaponRescueAircraft.fly()
}
