package model

import (
	"fmt"
	"testing"
)

func TestSingleLinkedList(t *testing.T) {
	ll := NewLinkedList()
	//ll.PrintFull()
	//ll.PrintHeadTail()

	ll.InsertFirst("firstNode")
	//ll.PrintFull()
	//ll.PrintHeadTail()

	ll.InsertLast("secondNode")
	//ll.PrintFull()
	//ll.PrintHeadTail()
	ll.InsertLast("thirdNode")
	//ll.PrintFull()
	//ll.PrintHeadTail()
	ll.InsertLast("fourthNode")
	ll.PrintFull()
	ll.PrintHeadTail()
	fmt.Println()

	//if err := ll.InsertMiddle("middle1", 4); err != nil {
	//	fmt.Println(err)
	//} else {
	//	ll.PrintFull()
	//	ll.PrintHeadTail()
	//}
	//if err := ll.InsertMiddle("middle2", 3); err != nil {
	//	fmt.Print(err, "\n\n")
	//} else {
	//	ll.PrintFull()
	//	ll.PrintHeadTail()
	//}
	//if err := ll.InsertMiddle("middle3", 4); err != nil {
	//	fmt.Println(err)
	//} else {
	//	ll.PrintFull()
	//	ll.PrintHeadTail()
	//}

	//ll.DeleteFirst()
	//ll.PrintFull()
	//ll.PrintHeadTail()
	//fmt.Println()
	//ll.DeleteFirst()
	//ll.PrintFull()
	//ll.PrintHeadTail()
	//fmt.Println()

	//ll.DeleteLast()
	//ll.PrintFull()
	//ll.PrintHeadTail()
	//fmt.Println()
	//ll.DeleteLast()
	//ll.PrintFull()
	//ll.PrintHeadTail()
	//fmt.Println()

	//if err := ll.DeleteMiddle(1); err != nil {
	//	fmt.Println(err)
	//} else {
	//	ll.PrintFull()
	//	ll.PrintHeadTail()
	//}

	n, err := ll.GetNode(2)
	fmt.Println(n, err)
}
