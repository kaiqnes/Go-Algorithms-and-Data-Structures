package model

type stack struct {
	values []interface{}
}

type Stack interface {
	Push(val interface{})
	Pop() interface{}
	Size() int
}

func NewStack() Stack {
	return &stack{}
}

func (s *stack) Size() int {
	return len(s.values)
}

func (s *stack) Push(val interface{}) {
	s.values = append(s.values, val) // Simply append the new value to the end of the stack
}

func (s *stack) Pop() interface{} {
	var val interface{}
	if s.Size() > 0 {
		index := len(s.values) - 1    // Get the index of the top most element.
		val = (s.values)[index]       // Index into the slice and obtain the element.
		s.values = (s.values)[:index] // Remove it from the stack by slicing it off.
	}
	return val
}
