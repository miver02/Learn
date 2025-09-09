package basic_type

import "unicode/utf8"

func StringDemo() {
	println("hello" + "go")
	println(len("abc"))
	println(len("你好"))
    println(utf8.RuneCountInString("你好"))

}

