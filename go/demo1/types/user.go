package types

import "fmt"

func NewUser() {
	u := User{}
	fmt.Printf("user: %v \n", u)
	fmt.Printf("user: %+v \n", u)


	// 取地址
	up := &User{}
	fmt.Printf("user: %+v \n", up)

	up2 := new(User)
	fmt.Printf("user: %+v \n", up2)

	// 接口赋值
	u3 := User{Name: "李白", Age: 18}

	u3.Name = "张飞"
	u3.Age = 2000
	println(u3.Name)

	// 指针
	var up4 *User
	println(up4)
}



type User struct {
	Name string
	Age  int
}

// 结构体接收器
func (u User) ChangeName(name string) {
	fmt.Printf("change name 中 u 的地址 %p \n", &u)
	u.Name = name
}

// 指针接收器
func (u *User) ChangeAge(age int) {
	fmt.Printf("change age 中 u 的地址 %p \n", &u)
	u.Age = age
}

func ChangeUser() {
	u1 := User{Name: "miver", Age: 18}
	println(&u1)
	u1.ChangeName("miaut")
	u1.ChangeAge(19)
	fmt.Printf("%+v", u1)

	u2 := &User{Name: "miver", Age: 18}
	println(&u1)
	u2.ChangeName("miaut")
	u2.ChangeAge(19)
	fmt.Printf("%+v \n", u2)
}

type Integer int

// 衍生类型不能互相访问方法,但能访问类型,仅仅访问
func UseInteger() {
	var a1 int = 10
	a2 := Integer(a1)
	println(a1, a2)
}

type Fish struct {
	Name string
}

func (f Fish) Swim() {
	println("飞起来")
}

type FakeFish Fish

func UseFish () {
	a1 := Fish{}
	a1.Swim()
	
	// a2 = FakeFish(a1)
	
}


// 向后兼容
type Yu = Fish