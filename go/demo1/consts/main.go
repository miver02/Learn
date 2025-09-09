package main
package consts

// 常量
const External = " 包外"
const internal = "包内"
const (
	a = 123
)

const (
	StatesA = iota + 18
	StatesB = iota
	StatesC = iota
	StatesD = iota << 2
)

func ConstsDemo() {
	const a = 123
	println(a)

	println(StatesA)
	println(StatesB)
	println(StatesC)
	println(StatesD)

}

func main() {
	ConstsDemo()
}