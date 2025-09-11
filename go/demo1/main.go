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
import "github.com/miver02/Learn/go/demo1/control"
import "github.com/miver02/Learn/go/demo1/buildin_type"

//BaseGrammer3
import "github.com/miver02/Learn/go/demo1/types"





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

	// 不定参数
	funcs.UseYouName()

	// defer是后进先出
	funcs.Defer()

	// 闭包与参数
	funcs.DeferV1()
	funcs.DeferV2()

	//defer 语句会将其后面的函数调用推迟到包含该 defer 语句的函数即将返回时执行
	println(funcs.DeferReturnV1())
	println(funcs.DeferReturnV2())
	println(funcs.DeferReturnV3())

	funcs.DeferTestV1()
	funcs.DeferTestV2()
	funcs.DeferTestV3()

	// 循环
	control.ForArr()
	control.ForMap()
	control.Swith(0)
	control.SwithBool(0)

	// 数组
	buildin_type.Array()
	// 切片
	buildin_type.Slice()
	// 子切片
	buildin_type.SubSlice()
	// 切片共享
	buildin_type.ShareSlice()
	// Map
	buildin_type.Map()
}

func BaseGrammer3() {
	// 接口
	types.NewUser()
	// 结构体和指针接收器区别
	types.ChangeUser()
	// 衍生类型
	types.UseInteger()
}


func main() {
	// BaseGrammer1()
	// BaseGrammer2()
	BaseGrammer3()

}