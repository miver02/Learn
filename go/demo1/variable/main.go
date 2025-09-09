package variable

var Global = "全局变量"
var internal = "内部变量"

func VariableDemo() {
	var a int = 10
	println(a)

	var b int = 20
	println(b)

	var c uint = 30
	println(c)

	println(a + b)
}
