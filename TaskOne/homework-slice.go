package main

import "fmt"

// doubleSlice 接收一个整数切片的指针，将每个元素乘以 2
func doubleSlice(slicePtr *[]int) {
	slice := *slicePtr // 先解引用得到切片
	for i := range slice {
		slice[i] *= 2
	}
	// 注意：这里修改的是底层数组，所以即使不重新赋值，原切片也会变
	// 但如果用了 append 并导致扩容，则必须通过指针赋值回去
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", nums)

	doubleSlice(&nums)

	fmt.Println("修改后切片:", nums)
}
