package main

import "fmt"

/*
首先来看，如果没有工厂模式，在开发者创建一个类的对象时，如果有很多不同种类的对象将会如何实现，代码如下：
*/

// Fruit 水果类
type Fruit struct {
}

func (f *Fruit) Show(name string) {
	if name == "apple" {
		fmt.Println("我是苹果")
	} else if name == "banana" {
		fmt.Println("我是香蕉")
	} else if name == "pear" {
		fmt.Println("我是梨")
	}
}

// NewFruit 创建一个Fruit对象
func NewFruit(name string) *Fruit {
	fruit := new(Fruit)
	if name == "apple" {
		//创建apple逻辑
	} else if name == "banana" {
		//创建banana逻辑
	} else if name == "pear" {
		//创建pear逻辑
	}
	return fruit
}

/*
(1) 在Fruit类中包含很多“if…else…”代码块，整个类的代码相当冗长，代码越长，阅读难度、维护难度和测试难度也越大；而且大量条件语句的存在还将影响系统的性能，程序在执行过程中需要做大量的条件判断。
(2) Fruit类的职责过重，它负责初始化和显示所有的水果对象，将各种水果对象的初始化代码和显示代码集中在一个类中实现，违反了“单一职责原则”，不利于类的重用和维护；
(3) 当需要增加新类型的水果时，必须修改Fruit类的构造函数NewFruit()和其他相关方法源代码，违反了“开闭原则”。

关键是来观察main()函数，main()函数与Fruit类是两个模块。当业务层希望创建一个对象的时候，将直接依赖Fruit类型的构造方法NewFruit()，
这样随着Fruit的越来越复杂，那么业务层的开发逻辑也需要依赖Fruit模块的更新，且随之改变，这样将导致业务层开发需要观察Fruit模块做改动，影响业务层的开发效率和稳定性。
整体的依赖关系为：业务逻辑层 ---> 基础类模块
那么如何将业务层创建对象与基础类模块做解耦呢，这里即可以在中间加一层工厂模块层，来降低业务逻辑层对基础模块层的直接依赖和耦合关联。
整体的依赖关系为：业务逻辑层 ---> 工厂模块 ---> 基础类模块
这样就引出了需要对工厂模块的一些设计和加工生成基础模块对象的模式。

简单工厂模式角色和职责

	简单工厂模式并不属于GoF的23种设计模式。他是开发者自发认为的一种非常简易的设计模式，其角色和职责如下：
	工厂（Factory）角色：简单工厂模式的核心，它负责实现创建所有实例的内部逻辑。工厂类可以被外界直接调用，创建所需的产品对象。
	抽象产品（AbstractProduct）角色：简单工厂模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。
	具体产品（Concrete Product）角色：简单工厂模式所创建的具体实例对象。
*/
func main() {
	apple := NewFruit("apple")
	apple.Show("apple")

	banana := NewFruit("banana")
	banana.Show("banana")

	pear := NewFruit("pear")
	pear.Show("pear")
}
