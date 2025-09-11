package types

import "time"

type LinkedList struct {
	head *node
	tail *node

	Len int

	CreatTime time.Time
}


type node struct {
	prod *node
}

// 使用 node 和 LinkedList，实现 Append 和 Delete 方法

// Append 向链表尾部添加一个新节点
func (l *LinkedList) Append(val any) {
	newNode := &node{}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.prod = newNode
		l.tail = newNode
	}
	l.Len++
}

// Delete 删除指定索引的节点，并返回其值（这里只返回nil作为示例）
func (l *LinkedList) Delete(idx int) (any, error) {
	if idx < 0 || idx >= l.Len || l.head == nil {
		return nil, nil
	}
	var prev *node
	cur := l.head
	for i := 0; i < idx; i++ {
		prev = cur
		cur = cur.prod
	}
	if prev == nil {
		// 删除头节点
		l.head = cur.prod
		if l.head == nil {
			l.tail = nil
		}
	} else {
		prev.prod = cur.prod
		if cur == l.tail {
			l.tail = prev
		}
	}
	l.Len--
	return nil, nil
}
