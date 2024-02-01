package main

import (
	"errors"
	"fmt"
)

// ResourcePoolConfig 配置资源池的结构体
type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

// Option 定义了配置选项的函数类型
type Option func(*ResourcePoolConfig) error

// NewResourcePoolConfig 创建一个新的 ResourcePoolConfig 实例，并应用所有的配置选项
func NewResourcePoolConfig(name string, opts ...Option) (*ResourcePoolConfig, error) {
	if name == "" {
		return nil, errors.New("name cannot be blank")
	}

	config := &ResourcePoolConfig{
		name:     name,
		maxTotal: 8, // 默认值
		maxIdle:  8, // 默认值
		minIdle:  0, // 默认值
	}

	for _, opt := range opts {
		err := opt(config)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}

// WithMaxTotal 设置 maxTotal 配置选项
func WithMaxTotal(maxTotal int) Option {
	return func(c *ResourcePoolConfig) error {
		if maxTotal <= 0 {
			return errors.New("maxTotal must be positive")
		}
		c.maxTotal = maxTotal
		return nil
	}
}

// WithMaxIdle 设置 maxIdle 配置选项
func WithMaxIdle(maxIdle int) Option {
	return func(c *ResourcePoolConfig) error {
		if maxIdle < 0 {
			return errors.New("maxIdle cannot be negative")
		}
		if maxIdle > c.maxTotal {
			return errors.New("maxIdle cannot be greater than maxTotal")
		}
		c.maxIdle = maxIdle
		return nil
	}
}

// WithMinIdle 设置 minIdle 配置选项
func WithMinIdle(minIdle int) Option {
	return func(c *ResourcePoolConfig) error {
		if minIdle < 0 {
			return errors.New("minIdle cannot be negative")
		}
		if minIdle > c.maxTotal || minIdle > c.maxIdle {
			return errors.New("minIdle cannot be greater than maxTotal or maxIdle")
		}
		c.minIdle = minIdle
		return nil
	}
}

func main() {
	// option模式将配置项之间的约束条件放到了WithXxx中，在NewResourcePoolConfig中通过遍历opts来检查约束条件
	// 所以需要注意调用WithXxx的顺序，否则可能会导致配置不正确。
	// 在build模式中，约束条件放到了Build方法中，通过Build方法来检查约束条件，所以不需要考虑调用顺序。
	config, err := NewResourcePoolConfig("dbconnectionpool",
		WithMaxTotal(16),
		WithMaxIdle(13),
		WithMinIdle(12),
	)

	if err != nil {
		fmt.Printf("Error configuring pool: %s\n", err)
		return
	}

	fmt.Printf("ResourcePoolConfig: %+v\n", config)
}
