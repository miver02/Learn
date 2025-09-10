package control


func ForLoop() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	for i := 0; i < 10; {
		println(i)
		i++
	}

}

func Loop2() {
	i := 0
	for i < 10 {
		println(i)
		i++
	}
}

func ForArr() {
	arr := [3]int{1, 2, 3}
	for index, val := range arr {
		println("下标 ", index, "值", val)
	}

}

func ForMap() {
	m := map[string]int {
		"key1": 100,
		"key2": 102,
	}

	for k, v := range m {
		println("key ", k, "value ", v)
	}
}