package funcs

import "fmt"

func Defer() {
	defer func() {
		println("第一个defer")
	}()

	defer func() {
		println("第二个defer")
	}()
}


func DeferV1() {
	i := 0
	// 闭包不传参,获取上下文中最终变量被赋予的值
	defer func() {
		println(i)
	}()
	i = 1
}

func DeferV2() {
	i := 0
	//闭包传参,获取闭包上文变量的值
	defer func(i int) {
		println(i)
	}(i)
	i = 1
}

func DeferReturnV1() int {
	i := 0
	defer func() {
		i = 1
	}()
	return i
}

func DeferReturnV2() (i int) {
	i = 0
	defer func() {
		i = 1
	}()
	return i
}

func DeferReturnV3() int {
	i := 0
	defer func(i int) {
		i = 1
	}(i)
	return i
}

func DeferTestV1() {
	println("=======test1=======")
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("out: %p ", &i)
			println(i)
		}()
	}
}

func DeferTestV2() {
	println("=======test2=======")
	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Printf("out: %p ", &val)
			println(val)
		}(i)
	}
}

func DeferTestV3() {
	println("=======test3=======")
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			
			fmt.Printf("out: %p ", &j)
			println(j)
		}()
	}
}
