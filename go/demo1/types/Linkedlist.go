package types

import "time"

type LinkedList struct {
	head *node
	tail *node

	Len int

	CreatTime time.Time
}

func (l *LinkedList) Add(idx int, val any) {

}

func (l *LinkedList) AddV1() {

}

type node struct {
	prod *node
}

