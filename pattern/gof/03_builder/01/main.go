package main

import (
	"errors"
	"fmt"
)

/*
假设我们要创建一个资源池，需要设置资源名称(name)、最大总资源数量(maxTotal)、最大空闲资源数量(maxIdle)、最小空闲资源数量(minIdle)等。
其中name必填，maxTotal、maxIdle、minIdle非必填，但是填了一个其它两个也需要填，而且数据值有限制，如不能等于0，数据间有限制，如maxIdle不能大于maxTotal。

碰到这种问题如何处理呢？

我们可以使用构造函数，所有的判断都在构造函数里做。但是一旦输入参数很多，会导致调用的时候容易写乱，而且构造函数里判断太多，后面需求有变化，构造函数也需要更改，不满足开放封闭原则。

如果使用set，因为数据间有限制，很容易漏掉部分配置。而且有时资源对象为不可变对象，就不能暴露set方法。

这个时候，建造者模式就能发挥作用了。建造者将输入数据整理好，将数据以对象的方式传递给资源类的构造函数，资源类拿到数据直接获取数值即可，是不是就达到了分离的效果。今后有规则上的变动，只需要修改Builder即可。
*/

// ResourceParams Product内的参数
type ResourceParams struct {
	name     string
	maxTotal int64
	maxIdle  int64
	minIdle  int64
}

// ResourceProduct is the interface that Product接口
type ResourceProduct interface {
	show()
}

// RedisResourceProduct 实际Product， 有show函数
// 对于实际的产品，我们没有提供任何的构造函数和set方法，它的构造工作全部交给了这个产品的Builder，
// 这个Builder提供一系列的setXxx方法，并在setXxx方法和build方法中进行规则校验，在build方法最后将生产的属性赋值给实际Product
// 这样就实现了按需
type RedisResourceProduct struct {
	resourceParams ResourceParams
}

func (p *RedisResourceProduct) show() {
	fmt.Printf("Product的数据为 %+v ", p.resourceParams)
}

// ResourceBuilder 资源类创建接口
type ResourceBuilder interface {
	setName(name string) ResourceBuilder
	setMaxTotal(maxTotal int64) ResourceBuilder
	setMaxIdle(maxIdle int64) ResourceBuilder
	setMinIdle(minIdle int64) ResourceBuilder
	getError() error
	build() (p ResourceProduct)
}

// RedisResourceBuilder 实际建造者
type RedisResourceBuilder struct {
	resourceParams ResourceParams
	err            error
}

// 获取错误信息
func (r *RedisResourceBuilder) getError() error {
	return r.err
}

/**
 * @Description: 设置名称
 * @receiver r
 * @param name
 * @return ResourceBuilder
 */
func (r *RedisResourceBuilder) setName(name string) ResourceBuilder {
	if name == "" {
		r.err = errors.New("name为空")
		return r
	}
	r.resourceParams.name = name
	fmt.Println("RedisResourceBuilder setName ", name)
	return r
}

/**
 * @Description: 设置maxTotal值，值不能小于0
 * @receiver r
 * @param maxTotal
 * @return ResourceBuilder
 */
func (r *RedisResourceBuilder) setMaxTotal(maxTotal int64) ResourceBuilder {
	if maxTotal <= 0 {
		r.err = errors.New("maxTotal小于0")
		return r
	}
	r.resourceParams.maxTotal = maxTotal
	fmt.Println("RedisResourceBuilder setMaxTotal ", maxTotal)
	return r
}

/**
 * @Description: 设置maxIdle值，值不能小于0
 * @receiver r
 * @param maxIdle
 * @return ResourceBuilder
 */
func (r *RedisResourceBuilder) setMaxIdle(maxIdle int64) ResourceBuilder {
	if maxIdle <= 0 {
		r.err = errors.New("maxIdle小于0")
		return r
	}
	r.resourceParams.maxIdle = maxIdle
	fmt.Println("RedisResourceBuilder setMaxIdle ", maxIdle)
	return r
}

/**
 * @Description: 设置minIdle值，值不能小于0
 * @receiver r
 * @param minIdle
 * @return ResourceBuilder
 */
func (r *RedisResourceBuilder) setMinIdle(minIdle int64) ResourceBuilder {
	if minIdle <= 0 {
		r.err = errors.New("minIdle小于0")
		return r
	}
	r.resourceParams.minIdle = minIdle
	fmt.Println("RedisResourceBuilder setMinIdle ", minIdle)
	return r
}

/*
*
  - @Description: 构建product
    1. 做参数校验
    2. 根据参数生成product
  - @receiver r
  - @return p
*/
func (r *RedisResourceBuilder) build() (p ResourceProduct) {
	// 校验逻辑放到这里来做，包括必填项校验、依赖关系校验、约束条件校验等
	if r.resourceParams.name == "" {
		r.err = errors.New("name为空")
		return
	}
	// 都为0表示是默认值；这三个属性要不都设置，要不都不设置
	if !((r.resourceParams.maxIdle == 0 && r.resourceParams.minIdle == 0 && r.resourceParams.maxTotal == 0) ||
		(r.resourceParams.maxIdle != 0 && r.resourceParams.minIdle != 0 && r.resourceParams.maxTotal != 0)) {
		r.err = errors.New("数据需要保持一致")
		return
	}

	if r.resourceParams.maxIdle > r.resourceParams.maxTotal {
		r.err = errors.New("maxIdle > maxTotal")
		return
	}
	if r.resourceParams.minIdle > r.resourceParams.maxTotal || r.resourceParams.minIdle > r.resourceParams.maxIdle {
		r.err = errors.New("minIdle > maxTotal|maxIdle")
		return
	}
	fmt.Println("RedisResourceBuilder build")
	product := &RedisResourceProduct{
		resourceParams: r.resourceParams,
	}
	return product
}

// Director 指挥者
type Director struct {
}

/**
 * @Description: 指挥者控制建造过程
 * @receiver d
 * @param builder
 * @return *ResourceProduct
 */
func (d *Director) construct(builder ResourceBuilder) ResourceProduct {
	resourceProduct := builder.setName("redis").
		setMinIdle(10).
		setMaxIdle(10).
		setMaxTotal(20).
		build()

	err := builder.getError()
	if err != nil {
		fmt.Println("构建失败，原因为" + err.Error())
		return nil
	}
	return resourceProduct
}

func main() {
	builder := &RedisResourceBuilder{}

	var director Director
	product := director.construct(builder)

	if product == nil {
		return
	}

	product.show()
}
