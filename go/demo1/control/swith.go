package control

func Swith(status int) {
	switch status {
	case 0:
		println("初始化")
	case 1:
		println("运行中")
	default:
		println("未知状态")
	}

}

func SwithBool(age int) {
	switch {
	case age >= 18:
		println("成年")
	case 0 < age && age < 18:
		println("未成年")
	default:
		println("未知状态")
	}

}