package model

import (
	"errors"
	"fmt"
)

const (
	head = iota
)

type singleLinkedList struct {
	head Node
	tail Node
}

func NewLinkedList() *singleLinkedList {
	return &singleLinkedList{}
}

func (ll *singleLinkedList) PrintHeadTail() {
	fmt.Println("Head:", ll.head.GetValue(), "\tTail:", ll.tail.GetValue())
}

func (ll *singleLinkedList) PrintFull() {
	currentNode := ll.head
	if !ll.IsEmptyList() {
		count := 1
		fmt.Printf("%v -> ", ll.head.GetValue())
		for currentNode.HasNext() {
			currentNode = currentNode.GetNext()
			if currentNode.HasNext() {
				fmt.Printf("%v -> ", currentNode.GetValue())
				count++
			} else {
				fmt.Printf("%v\n", currentNode.GetValue())
			}
		}
	} else {
		fmt.Println("LinkedList is nil")
	}
}

func (ll *singleLinkedList) Size() int {
	var (
		current = ll.head
		count   = 1
	)

	for current.HasNext() {
		count++
		current = current.GetNext()
	}

	return count
}

func (ll *singleLinkedList) InsertFirst(val interface{}) {
	node := NewSingleLinkNode(val)
	if ll.IsEmptyList() {
		ll.head = node
	}

	if ll.tail != nil {
		ll.tail.SetNext(node)
	}

	ll.tail = node
}

func (ll *singleLinkedList) InsertLast(val interface{}) {
	node := NewSingleLinkNode(val)
	if ll.tail != nil {
		ll.tail.SetNext(node)
	}

	ll.tail = node
}

func (ll *singleLinkedList) InsertMiddle(val interface{}, position int) error {
	if position < head {
		return errors.New("position must be bigger than head (0)")
	}

	node := NewSingleLinkNode(val)
	if ll.IsEmptyList() {
		ll.head = node
	}

	if ll.tail == nil {
		ll.tail = node
	}

	var (
		current = ll.head
		count   = 1
	)

	var hasChange bool
	for current.HasNext() {
		if count == position {
			temp := current.GetNext()
			current.SetNext(node)
			current.GetNext().SetNext(temp)
			hasChange = true
			break
		}
		count++
		current = current.GetNext()
	}

	if !hasChange {
		return errors.New("middle position not found")
	}

	return nil
}

func (ll *singleLinkedList) DeleteFirst() error {
	if ll.IsEmptyList() {
		return errors.New("list is empty")
	}

	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
		return nil
	}

	ll.head = ll.head.GetNext()
	return nil
}

func (ll *singleLinkedList) DeleteLast() error {
	if ll.IsEmptyList() {
		return errors.New("list is empty")
	}

	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
		return nil
	}

	current := ll.head
	for current.HasNext() {
		if !current.GetNext().HasNext() {
			current.SetNext(nil)
			ll.tail = current
			break
		}
		current = current.GetNext()
	}

	return nil
}

func (ll *singleLinkedList) DeleteMiddle(position int) error {
	if ll.isValidPosition(position) {
		return errors.New("position must be bigger than head (0)")
	}

	var (
		current = ll.head
		count   = 0
	)

	var hasChange bool
	for current.HasNext() {
		if count == position-1 {
			toBeDeleted := current.GetNext()
			if toBeDeleted == nil || toBeDeleted == ll.tail {
				return errors.New("middle position not found")
			}
			afterPosition := toBeDeleted.GetNext()
			current.SetNext(afterPosition)
			hasChange = true
			break
		}
		count++
		current = current.GetNext()
	}

	if !hasChange {
		return errors.New("middle position not found")
	}

	return nil
}

func (ll *singleLinkedList) GetNode(position int) (Node, error) {
	if ll.isValidPosition(position) {
		return nil, errors.New("position must be bigger than head (0)")
	}

	if ll.IsEmptyList() {
		return nil, errors.New("list is empty")
	}

	var (
		current = ll.head
		count   = 0
	)

	for current.HasNext() && count <= position {
		if count == position {
			return current, nil
		}
		count++
		current = current.GetNext()
	}
	return nil, errors.New("position not found")
}

func (ll *singleLinkedList) IsEmptyList() bool {
	return ll.head == nil
}

func (ll *singleLinkedList) isValidPosition(position int) bool {
	return position < head
}
