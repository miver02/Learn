package buildin_type

import "fmt"

// 切片
func Slice() {
	println("======切片======")
	s1 := []int{9, 8, 7, 6, 7, 5, 5}
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 4)
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s3 := make([]int, 4)
	fmt.Printf("s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))
	
	// 添加元素
	s4 := make([]int, 0, 4)
	s4 = append(s4, 1)
	fmt.Printf("s4: %v, len: %d, cap: %d \n", s4, len(s4), cap(s4))

	
}

func SubSlice() {
	println("=====子切片=====")
	// 子切片
	// 容量是按照start开始计算,直到原切片s1结尾的长度
	s1 := []int{9, 8, 7, 6, 7, 5, 5}

	s5 := s1[1:3]
	fmt.Printf("s5: %v, len: %d, cap: %d \n", s5, len(s5), cap(s5))

	s6 := s1[3:]
	fmt.Printf("s6: %v, len: %d, cap: %d \n", s6, len(s6), cap(s6))

	s7 := s1[:3]
	fmt.Printf("s7: %v, len: %d, cap: %d \n", s7, len(s7), cap(s7))
}

func ShareSlice() {
	println("=====切片共享=====")
	s1 := []int{9, 8, 7, 6, 7, 5, 5}

	s5 := s1[1:3]
	fmt.Printf("s5: %v, len: %d, cap: %d \n", s5, len(s5), cap(s5))

	s5[0] = 99
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s5: %v, len: %d, cap: %d \n", s5, len(s5), cap(s5))

	s5 = append(s5, 199)
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s5: %v, len: %d, cap: %d \n", s5, len(s5), cap(s5))

	s1 = append(s1, 199)
	s5 = append(s5, 199)
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s5: %v, len: %d, cap: %d \n", s5, len(s5), cap(s5))

}