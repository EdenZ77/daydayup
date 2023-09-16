package main

import "fmt"

// Aircraft 飞行器接口，有fly函数
type Aircraft interface {
	fly()
	landing()
}

// Helicopter 直升机类，拥有正常飞行、降落功能
type Helicopter struct {
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

/**
 * @Description: 给直升机增加武装功能
 * @receiver a
 */
func (a *WeaponAircraft) fly() {
	a.Aircraft.fly()
	fmt.Println("增加武装功能")
}

// RescueAircraft 救援直升机
type RescueAircraft struct {
	Aircraft
}

/**
 * @Description: 给直升机增加救援功能
 * @receiver r
 */
func (r *RescueAircraft) fly() {
	r.Aircraft.fly()
	fmt.Println("增加救援功能")
}

/*
装饰器模式相对于简单的组合关系，还有两个比较特殊的地方。
第一个比较特殊的地方是：装饰器类和原始类继承同样的父类，这样我们可以对原始类“嵌套”多个装饰器类。
	比如飞行器这个接口，装饰器类(武装、救援)和原始类(直升机)都实现了这个接口，那么我们就可以对原始类进行多次装饰
第二个比较特殊的地方是：装饰器类是对功能的增强，这也是装饰器模式应用场景的一个重要特点。
	实际上，符合“组合关系”这种代码结构的设计模式有很多，比如之前讲过的代理模式、桥接模式，还有现在的装饰器模式。
	尽管它们的代码结构很相似，但是每种设计模式的意图是不同的。
	就拿比较相似的代理模式和装饰器模式来说吧，代理模式中，代理类附加的是跟原始类无关的功能，而在装饰器模式中，装饰器类附加的是跟原始类相关的增强功能。
*/

func main() {
	//普通直升机
	fmt.Println("------------普通直升机")
	helicopter := &Helicopter{}
	helicopter.fly()
	helicopter.landing()

	//武装直升机
	fmt.Println("------------武装直升机")
	weaponAircraft := &WeaponAircraft{
		Aircraft: helicopter,
	}
	weaponAircraft.fly()

	//救援直升机
	fmt.Println("------------救援直升机")
	rescueAircraft := &RescueAircraft{
		Aircraft: helicopter,
	}
	rescueAircraft.fly()

	//武装救援直升机
	fmt.Println("------------武装救援直升机")
	weaponRescueAircraft := &RescueAircraft{
		Aircraft: weaponAircraft,
	}
	weaponRescueAircraft.fly()
}
