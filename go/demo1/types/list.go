package types

// type 名字 interface
// 里面只能有方法,并且不需要func关键字
type List interface {
	Add(idx int, val any) error
	Append(aa any)
	Delete(idx int) (any, error)
	// toSlice(idx int) (any error)
}

// 实现接口
