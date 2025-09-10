package control


func IfOnly(age int){
	if age >= 18 {
		println( "成年了")
	}
}

func IfElse(age int){
	if age >= 18 {
		println( "成年了")
	} else {
		println( "没有成年")
	}
}

func IfElseIf(age int){
	if age >= 18 {
		println( "成年了")
	} else if age >= 12 {
		println( "少年")
	} else {
		println( "小孩")
	}
}

func IfNewVariable(start int,end int) string {
	if distance := end-start; distance>100 {
		return "远了"
	} else if distance >60 {
		return "有点远"
	} else {
		return "还挺好"
	}
}
	