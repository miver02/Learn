// package main
package variable

// 首字母大写:包外可访问
// 首字母小写:包外不可访问
var Global = "全局变量"
var internal = "内部变量"


var (
	// 包变量
	First string = "1"
	second int = 2
	abb string = "sdff"
)



func VariableDemo() {
	var a int = 10
	// var a int = "10" // 同作用域下,不能重复声明
	println(a)

	var abb int = 10 // 可以覆盖包变量
	println(abb)

	var Global int = 12313
	println(Global)

	var b int = 20
	println(b)
	println(a + b)

	var c uint = 30
	println(c)

	var (
		d string = "aafff"
		e int = 123
	)
	println(d, e)

}

func main() {
	VariableDemo()
}