package main

import "fmt"

/*

那么如何将业务层创建对象与基础类模块做解耦呢，这里即可以在中间加一层工厂模块层，来降低业务逻辑层对基础模块层的直接依赖和耦合关联。
整体的依赖关系为：业务逻辑层 ---> 工厂模块 ---> 基础类模块
这样就引出了需要对工厂模块的一些设计和加工生成基础模块对象的模式。

简单工厂模式角色和职责

	简单工厂模式并不属于GoF的23种设计模式。他是开发者自发认为的一种非常简易的设计模式，其角色和职责如下：
	工厂（Factory）角色：简单工厂模式的核心，它负责实现创建所有实例的内部逻辑。工厂类可以被外界直接调用，创建所需的产品对象。
	抽象产品（AbstractProduct）角色：简单工厂模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。
	具体产品（Concrete Product）角色：简单工厂模式所创建的具体实例对象。
*/

// ======= 抽象层 =========

// Fruit 水果类(抽象接口)
type Fruit interface {
	Show() //接口的某方法
}

// ======= 基础类模块 =========

type Apple struct {
}

func (apple *Apple) Show() {
	fmt.Println("我是苹果")
}

type Banana struct {
}

func (banana *Banana) Show() {
	fmt.Println("我是香蕉")
}

type Pear struct {
}

func (pear *Pear) Show() {
	fmt.Println("我是梨")
}

// ========= 工厂模块  =========

// Factory 一个工厂， 有一个生产水果的机器，返回一个抽象水果的指针
type Factory struct{}

func (fac *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit

	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}

	return fruit
}

// ==========业务逻辑层==============
/*
上述代码可以看出，业务逻辑层只会和工厂模块进行依赖，这样业务逻辑层将不再关心Fruit类是具体怎么创建基础对象的。

缺点：
1. 对工厂类职责过重，一旦不能工作，系统受到影响。
2. 增加系统中类的个数，复杂度和理解度增加。
3. 违反“开闭原则”，添加新产品需要修改工厂逻辑，工厂越来越复杂。

适用场景：
1. 工厂类负责创建的对象比较少，由于创建的对象较少，不会造成工厂方法中的业务逻辑太过复杂。
2. 客户端只知道传入工厂类的参数，对于如何创建对象并不关心。
*/
func main() {
	factory := new(Factory)

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
