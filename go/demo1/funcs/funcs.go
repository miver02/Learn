package funcs


// 常见函数
func Func1() {

}

// 带参数的函数
func Func2(a int) {

}

// 携带不同类型的参数
func Func3(a int, b string) {

}

// 携带同类型参数
func Func4(a, b int) {

}

// 需要返回值
func Func5(a, b int, c string) string {
	return "hello go!"
}

// 多个返回值
func Func6(a, b int, c string) (string, string) {
	return "hello go!", "你好,go!"
}

// 返回值命名: 要么都有名字,要么都没有名字
func Func7(a, b int, c string) (name string, age string) {
	return "hello go!", "18"
}

// 默认返回命名参数
func Func8(a, b int, c string) (name string, age string) {
	name = "hello go!"
	age  = "18"
	return 
}

// 默认返回对应类型的零值
func Func9() (name string, age int) {
	// 等价于"" 0
	// 对应类型的零值
	return 
} 