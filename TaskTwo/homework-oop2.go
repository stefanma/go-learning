package main

import "fmt"

// Person 结构体
type Person struct {
	Name string
	Age  int
}

// Employee 组合 Person，并添加自己的字段
type Employee struct {
	Person     // 匿名字段，实现组合
	EmployeeID string
}

// 为 Employee 实现 PrintInfo 方法
func (e Employee) PrintInfo() {
	fmt.Printf("姓名: %s, 年龄: %d, 工号: %s\n", e.Name, e.Age, e.EmployeeID)
}

// func main() {
// 	emp := Employee{
// 		Person:     Person{Name: "张三tyu", Age: 30},
// 		EmployeeID: "E12345",
// 	}

// 	emp.PrintInfo()

// 	// 也可以直接访问嵌入字段
// 	fmt.Println("直接访问:", emp.Name, emp.Age, emp.EmployeeID)
// }
