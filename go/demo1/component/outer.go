package component



type Inner struct {
	Name string
}

func (i Inner) InnerIm(name string) {
	println("这是" + name + "的InnerIm")
}

func (i Inner) IName() string {
	return "inner"
}


func (i Inner) InnerName() {
	println("这是" + i.IName() + "的InnerName")
}


type Outer struct {
	Inner
}

func (o Outer) OuterName() {
	println("这是" + o.OName() + "的OuterName")
}

func (o Outer) OName() string {
	return "outer"
}

// 结构体嵌套:组合
type OuterV1 struct {
	*Inner
	Outer
}





func (o1 OuterV1) OuterV1Im() {
	var inner Inner
	inner.Name = "张三"
	inner.InnerIm(inner.Name)

	outer := Outer{}
	outer.Name = "李四"
	outer.InnerIm(outer.Name)

	// 没有多态 
	outer.InnerName()
	outer.OuterName()


}