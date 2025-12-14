package main

// addTen 接收一个整数指针，并将其指向的值加 10
func addTen(p *int) {
	*p += 10
}

// func main() {
// 	num := 5
// 	fmt.Println("原始值:", num)

// 	addTen(&num) // 传入地址

// 	fmt.Println("修改后值:", num)
// }
