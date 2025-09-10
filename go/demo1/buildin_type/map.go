package buildin_type

import "fmt"

func Map() {
	println("======Map=======")
	m1 := map[string]int{
		"唐山": 123,
		"李白": 234,
		"小名": 2424,
	}

	m1["李白"] = 1000

	// 容量
	m2 := make(map[string]int, 12)
	m2["key2"] = 12

	val, ok := m1["小白"]
	if ok {
		println(val)
	}

	val = m1["李白"]
	fmt.Printf("李白: %d \n", val)
	val1 := m1["大黄"]
	fmt.Printf("大黄: %d \n", val1)

	delete(m1, "李白")
	for key, val := range m1 {
		fmt.Printf("new: %s: %d\n", key, val)
	}
}