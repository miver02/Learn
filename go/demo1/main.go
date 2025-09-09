package main
// package demo1

// BaseGrammer1
import "github.com/miver02/Learn/go/demo1/hello_world"
import "github.com/miver02/Learn/go/demo1/basic_type"
import "github.com/miver02/Learn/go/demo1/variable"
import "github.com/miver02/Learn/go/demo1/consts"

// BaseGrammer2
// 方法签名:名字 + 参数列表 + 返回值 
import "github.com/miver02/Learn/go/demo1/funcs"



func BaseGrammer1() {
	hello_world.HelloWorldDemo()
	basic_type.BasicTypeDemo()
	println(variable.Global)
	// println(variable.internal) // 内部变量不能被外部访问
	variable.VariableDemo()
	consts.ConstsDemo()
}

func BaseGrammer2() {
	// 使用 := 的前提是左边至少有一个新变量
	name, age := funcs.Func9()
	println(name, age)

	// 递归调用
	funcs.Recursive(0)

	// 方法赋值给变量
	funcs.UseFunctional()

	// 匿名方法立即调用
	funcs.Functional4()

	// 闭包
	str := funcs.Closure("world!")
	println(str())
}


func main() {
	// BaseGrammer1()
	BaseGrammer2()

}