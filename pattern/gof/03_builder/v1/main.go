package main

import (
	"errors"
	"fmt"
)

/*
参考资料：GPT，将极客时间的java建造者模式改为go语言实现
*/

// ResourcePoolConfig 配置资源池的结构体
type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

// Builder 建造者模式的结构体
type Builder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

const (
	DEFAULT_MAX_TOTAL = 8
	DEFAULT_MAX_IDLE  = 8
	DEFAULT_MIN_IDLE  = 0
)

// NewBuilder 创建一个新的 Builder 实例
func NewBuilder() *Builder {
	return &Builder{
		maxTotal: DEFAULT_MAX_TOTAL,
		maxIdle:  DEFAULT_MAX_IDLE,
		minIdle:  DEFAULT_MIN_IDLE,
	}
}

// Build 构建一个新的 ResourcePoolConfig 实例
func (b *Builder) Build() (*ResourcePoolConfig, error) {
	if b.name == "" {
		return nil, errors.New("name cannot be blank")
	}
	if b.maxIdle > b.maxTotal {
		return nil, errors.New("maxIdle cannot be greater than maxTotal")
	}
	if b.minIdle > b.maxTotal || b.minIdle > b.maxIdle {
		return nil, errors.New("minIdle cannot be greater than maxTotal or maxIdle")
	}

	return &ResourcePoolConfig{
		name:     b.name,
		maxTotal: b.maxTotal,
		maxIdle:  b.maxIdle,
		minIdle:  b.minIdle,
	}, nil
}

// SetName 设置 Builder 的 name 属性
func (b *Builder) SetName(name string) *Builder {
	if name == "" {
		panic("name cannot be blank")
	}
	b.name = name
	return b
}

// SetMaxTotal 设置 Builder 的 maxTotal 属性
func (b *Builder) SetMaxTotal(maxTotal int) *Builder {
	if maxTotal <= 0 {
		panic("maxTotal must be positive")
	}
	b.maxTotal = maxTotal
	return b
}

// SetMaxIdle 设置 Builder 的 maxIdle 属性
func (b *Builder) SetMaxIdle(maxIdle int) *Builder {
	if maxIdle < 0 {
		panic("maxIdle cannot be negative")
	}
	b.maxIdle = maxIdle
	return b
}

// SetMinIdle 设置 Builder 的 minIdle 属性
func (b *Builder) SetMinIdle(minIdle int) *Builder {
	if minIdle < 0 {
		panic("minIdle cannot be negative")
	}
	b.minIdle = minIdle
	return b
}

func main() {
	builder := NewBuilder()
	config, err := builder.SetName("dbconnectionpool").
		SetMaxTotal(16).
		SetMaxIdle(10).
		SetMinIdle(12).
		Build()

	if err != nil {
		fmt.Printf("Error building config: %s\n", err)
		return
	}

	fmt.Printf("ResourcePoolConfig: %+v\n", config)
}
