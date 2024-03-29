package main

import "fmt"

/*
原型模式是创建型模式中的最后一个，它主要用于对象创建成本比较大的情况。
*/

// Resume 简历类，里面包含简历的基本信息
type Resume struct {
	name       string
	age        int64
	sex        string
	company    string
	experience string
}

// 设置简历个人信息
func (r *Resume) setPersonInfo(name string, age int64, sex string) {
	r.name = name
	r.age = age
	r.sex = sex
}

// 设置工作经验
func (r *Resume) setWorkExperience(company string, experience string) {
	r.company = company
	r.experience = experience
}

// 显示简历内容
func (r *Resume) display() {
	fmt.Printf("我的名字是%s，性别%s，今年%d岁，在%s工作，工作经验为:%s \n", r.name, r.sex, r.age, r.company, r.experience)
}

// 深拷贝，原型模式的核心
func (r *Resume) clone() *Resume {
	return &Resume{
		name:       r.name,
		sex:        r.sex,
		age:        r.age,
		company:    r.company,
		experience: r.experience,
	}
}

func main() {
	fmt.Println("---------------------------原简历")
	resume := &Resume{
		name:       "王工作",
		sex:        "男",
		age:        10,
		company:    "光明顶无限责任公司",
		experience: "学武功和划水、摸鱼",
	}
	resume.display()
	fmt.Println("---------------------------简历特别好，抄")
	copyResume := resume.clone()
	copyResume.setPersonInfo("李工作", 21, "男")
	copyResume.display()
}
