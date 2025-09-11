package types


type ArrayList struct{

}

// 实现 List 接口
func (a *ArrayList) Add(idx int, val any) error {
	// 这里只是示例，具体实现需要根据内部结构补充
	return nil
}

func (a *ArrayList) Append(aa any) {
	// 这里只是示例，具体实现需要根据内部结构补充
}

func (a *ArrayList) Delete(idx int) (any, error) {
	// 这里只是示例，具体实现需要根据内部结构补充
	return nil, nil
}

