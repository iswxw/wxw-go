/*
@Time: 2022/3/18 16:19
@Author: wxw
@File: demo01_composit
*/
package main

import "fmt"

// 公司的人员组织就是一个典型的树状的结构，现在假设我们现在有部门，和员工，两种角色，一个部门下面可以存在子部门和员工，员工下面不能再包含其他节点。
// 我们现在要实现一个统计一个部门下员工数量的功能

func main() {
	got := NewOrganization().Count()
	fmt.Println("got = ", got)
}

// IOrganization 组织接口，实现统计人数的功能
type IOrganization interface {
	Count() int
}

// Employee 员工
type Employee struct {
	Name string
}

// Count 人数统计
func (Employee) Count() int {
	return 1
}

// Department 部门
type Department struct {
	Name string

	SubOrganizations []IOrganization
}

// Count 人数统计
func (d Department) Count() int {
	c := 0
	for _, org := range d.SubOrganizations {
		c += org.Count()
	}
	return c
}

// AddSub 添加子节点
func (d *Department) AddSub(org IOrganization) {
	d.SubOrganizations = append(d.SubOrganizations, org)
}

// NewOrganization 构建组织架构 demo
func NewOrganization() IOrganization {
	root := &Department{Name: "root"}
	for i := 0; i < 10; i++ {
		root.AddSub(&Employee{})
		root.AddSub(&Department{Name: "sub", SubOrganizations: []IOrganization{&Employee{}}})
	}
	return root
}
