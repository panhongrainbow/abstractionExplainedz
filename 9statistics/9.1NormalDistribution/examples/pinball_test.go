package examples

import (
	"fmt"
	"testing"
)

func Test_Check_pinball(t *testing.T) {
	//
	node := pinball{}
	node.pinballTree(0, 6)

	// node.ListPinballTree(10)
	node.ListPinballTree(6)
	test := node.algorithmComparatorCountLayer(8)
	fmt.Println(test)
	var layer = 8
	node.OneBallGoThrough(&layer)
	fmt.Println(layer)
	var position = 0
	node.GoToTheRightEdge(&position)
	fmt.Println(">>", position)
	position = 0
	node.GoToTheLeftEdge(&position)
	fmt.Println("<<", position)

	distribution := node.PinballTable(1000)
	fmt.Println(distribution)
	position = 0
	node.OneBallGoThrough(&position)
	fmt.Println(position)

}
