package main

import "fmt"

// 公司结构体
type Company struct { //结构体声明
	name       string
	address    string
	website    string
	createDate int
}

func main() { //主函数
	p := NewPerson("知链")
	p.SetCompany(2018)
	fmt.Println(p)
	fmt.Println(p.name, " createDate = ", p.GetCompany())
}
func NewPerson(name string) *Company { //工程函数构建
	return &Company{
		name: name,
	}
}

// 设置get方法，有返回值
func (p *Company) GetCompany() int { //函数方法封装
	return p.createDate
}

// SetCompany 为了访问age 和 sal 我们编写一对SetXxx的方法和GetXxx的方法
func (p *Company) SetCompany(createDate int) { //函数方法封装
	if createDate < 2022 && createDate > 1949 {
		p.createDate = createDate
	} else {
		fmt.Println("创建时间范围不正确..")
		//给程序员给一个默认值
	}
}
