package generics

import (
	"errors"
	"fmt"
)

type List[T any] interface {
	Add(idx int, t T)
	Append(t T)
}

func UseList() {
	// 错误的方式：接口不能直接实例化
	// var l List[int]  // 这是 nil，调用方法会panic

	// 正确的方式：需要实现接口的具体类型
	var l List[int] = &ArrayList[int]{}
	l.Append(12)
}

// 实现 List 接口的具体类型
type ArrayList[T any] struct {
	items []T
}

func (a *ArrayList[T]) Add(idx int, t T) {
	if idx < len(a.items) {
		a.items[idx] = t
	} else {
		a.items = append(a.items, t)
	}
}

func (a *ArrayList[T]) Append(t T) {
	a.items = append(a.items, t)
	println("添加元素:", t)
}

func AddSlice[T any](slice []T, idx int, val T) ([]T, error) {
	println("======数组分割=====")
	if idx < 0 || idx > len(slice) {
		return nil, errors.New("idx传入错误或者数据为空")
	}

	res := make([]T, 0, len(slice)+1)
	for i := 0; i < idx; i++ {
		res = append(res, slice[i])
	}

	res = append(res, val)

	for i := idx; i < len(slice); i++ {
		res = append(res, slice[i])
	}

	println(res)
	fmt.Println("结果切片:", res)
	return res, nil
}
