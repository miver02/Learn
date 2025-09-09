package funcs


func Functional1() string {
	return "hello go!"
}

func UseFunctional() {
	// 方法赋值给变量
	functional := Functional1
	value := functional()
	println(value)
}

func Functional2(){ 
	//新定义了一个方法，赋值给了 fn
	// 只能在Functional2里使用  
	fn := func() string {
		return "hello"
	}
	fn()
}

func Functional4(){ 
	//新定义了一个方法，赋值给了 fn
	// 只能在Functional2里使用  
	fn := func() string {
		return "hello"
	}() // 匿名方法立即调用
	println(fn)
}

// 返回一个 返回sting类型的方法 
func Functional3() func() string {
	return func() string {
		return "hello go!"
	}
}