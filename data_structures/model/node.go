package model

type Node interface {
	SetNext(Node)
	HasNext() bool
	GetNext() Node
	GetValue() interface{}
}

type singleLinkNode struct {
	value interface{}
	next  Node
}

type doubleLinkNode struct {
	value    interface{}
	previous Node
	next     Node
}

func NewSingleLinkNode(val interface{}) Node {
	return &singleLinkNode{value: val}
}

func (sln *singleLinkNode) GetNext() Node {
	return sln.next
}

func (sln *singleLinkNode) HasNext() bool {
	return sln.next != nil
}

func (sln *singleLinkNode) SetNext(node Node) {
	sln.next = node
}

func (sln *singleLinkNode) GetValue() interface{} {
	return sln.value
}

func NewDoubleLinkNode(val interface{}) Node {
	return &doubleLinkNode{value: val}
}

func (dln *doubleLinkNode) GetNext() Node {
	return dln.next
}

func (dln *doubleLinkNode) HasNext() bool {
	return dln.next != nil
}

func (dln *doubleLinkNode) SetNext(node Node) {
	dln.next = node
}

func (dln *doubleLinkNode) GetValue() interface{} {
	return dln.value
}
