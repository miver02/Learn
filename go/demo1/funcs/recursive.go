package funcs

func Recursive(n int) {
	// 递归一定要有退出机制
	if n > 10 {
		return 
	}
	Recursive(n + 1)
}