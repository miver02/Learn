package main

import "github.com/miver02/Learn/go/demo1/hello_world"
import "github.com/miver02/Learn/go/demo1/basic_type"
import "github.com/miver02/Learn/go/demo1/variable"



func main() {
	hello_world.HelloWorldDemo()
	basic_type.BasicTypeDemo()
	println(variable.Global)
	// println(variable.internal) // 内部变量不能被外部访问

}