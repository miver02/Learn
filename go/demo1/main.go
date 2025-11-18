package main

import (
	hello_world "github.com/miver02/learn-program/go/demo1/Hello_world"
	"github.com/miver02/learn-program/go/demo1/basic_type"
	"github.com/miver02/learn-program/go/demo1/buildin_type"
	"github.com/miver02/learn-program/go/demo1/component"
	"github.com/miver02/learn-program/go/demo1/consts"
	"github.com/miver02/learn-program/go/demo1/control"
	"github.com/miver02/learn-program/go/demo1/funcs"
	"github.com/miver02/learn-program/go/demo1/generics"
	"github.com/miver02/learn-program/go/demo1/goroutine"
	"github.com/miver02/learn-program/go/demo1/types"
	"github.com/miver02/learn-program/go/demo1/variable"
)

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
	// 结构体组合
	outerV1 := component.OuterV1{}
	outerV1.OuterV1Im()

	// 泛型
	list1 := [9]int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	generics.UseList()
	generics.Sum(1.2, 2.3, 3.3, 4.1)
	generics.Sum(1, 2, 3, 4)
	generics.Max(1, 2, 3, 4)
	generics.Min(1, 2, 3, 4)
	// 修复类型不匹配，将数组转换为切片
	generics.AddSlice(list1[:], 5, 6)

	// 并行
	goroutine.BaseTest()
	goroutine.ChanInt()
	goroutine.ChanSem()
	goroutine.SyncMutex()
}

func main() {
	// BaseGrammer1()
	// BaseGrammer2()
	BaseGrammer3()

}
