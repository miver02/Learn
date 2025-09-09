package funcs



func Closure(name string) func() string {
	// 返回的函数会绑定上下文
	return func() string {
		return "hello " + name
	}
}