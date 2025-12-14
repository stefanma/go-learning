package main

import (
	"math"
)

// Shape 接口定义
type Shape9 interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 结构体
type Rectangle9 struct {
	Width, Height float64
}

// Circle 结构体
type Circle9 struct {
	Radius float64
}

// Rectangle 实现 Shape 接口
func (r Rectangle9) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle9) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 实现 Shape 接口
func (c Circle9) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle9) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// func main() {
// 	rect := Rectangle9{Width: 3, Height: 4}
// 	circle := Circle9{Radius: 5}

// 	fmt.Printf("矩形面积: %.2f, 周长: %.2f\n", rect.Area(), rect.Perimeter())
// 	fmt.Printf("圆形面积: %.2f, 周长: %.2f\n", circle.Area(), circle.Perimeter())

// 	// 也可以通过接口变量调用（体现多态）
// 	var s Shape9

// 	s = rect
// 	fmt.Printf("通过接口调用矩形: 面积=%.2f\n", s.Area())

// 	s = circle
// 	fmt.Printf("通过接口调用圆形: 面积=%.2f\n", s.Area())
// }
