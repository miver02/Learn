package generics

import "errors"

func Sum[T Numble](vals ...T) T {
	println("=====泛型求和=====")
	var res T
	for _, val := range vals {
		res += val
	}

	println(res)
	return res
}

// 泛型约束：只约束泛型
type Numble interface {
	int | int64 | float64
}

func Max[T Numble](vals ...T) (T, error) {
	println("======泛型求最大值======")
	if len(vals) == 0 {
		var t T
		return t, errors.New("数组为空")
	}
	max := vals[0]
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	println(max)
	return max, nil
}

func Min[T Numble](vals ...T) (T, error) {
	println("======泛型求最小值======")
	if len(vals) == 0 {
		var t T
		return t, errors.New("数组为空")
	}
	min := vals[0]
	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	println(min)
	return min, nil
}