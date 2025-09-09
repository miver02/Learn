package funcs

// 不定参数
func YouName(name string, aliases ...string) {
	// aliases是一个切片
}


func UseYouName() {
	YouName("大狗")
	YouName("大狗", "小狗", "二狗")
	aliases := []string{"大黄", "小白"}
	YouName("你好", aliases...)
}