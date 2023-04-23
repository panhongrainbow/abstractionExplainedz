package examples

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type pinball struct {
	layer     int
	left      float64
	right     float64 // can ignore
	leftNode  *pinball
	rightNode *pinball
}

func (node *pinball) pinballTree(layer, limit int) {
	// node.left = rand.Float64()
	node.left = 0.5
	node.right = 1 - node.left
	node.makeNextNode(layer, limit)
}

func (node *pinball) makeNextNode(layer, limit int) {
	if layer > limit {
		return
	}
	node.layer = layer
	//
	if node.layer < limit {
		node.leftNode = &pinball{
			// left: rand.Float64(),
			left: 0.5,
		}
		node.leftNode.right = 1 - node.leftNode.left
		node.leftNode.pinballTree(layer+1, limit)
		//
		node.rightNode = &pinball{
			// left: rand.Float64(),
			left: 0.5,
		}
		node.rightNode.right = 1 - node.rightNode.left
		node.rightNode.pinballTree(layer+1, limit)
	}

	return
}

func (node *pinball) ListPinballTree(tab int) {
	if node == nil {
		return
	}
	if node.left == 0 {
		fmt.Println("left is 0")
	}
	if node.right == 0 {
		fmt.Println("right is 0")
	}
	fmt.Printf("%s", strings.Repeat("|_____", tab))
	fmt.Printf("Layer: %d, Left: %f, Right: %f\n", node.layer, node.left, node.right)
	node.leftNode.ListPinballTree(tab + 1)
	node.rightNode.ListPinballTree(tab + 1)
}

func (node *pinball) CountLayer(layer int) (count int) {
	count = 1 << layer
	return
}

func (node *pinball) algorithmComparatorCountLayer(layer int) (count int) {
	if node == nil {
		return
	}
	if node.layer == layer {
		count++
	}
	count = count + node.leftNode.algorithmComparatorCountLayer(layer)
	count = count + node.rightNode.algorithmComparatorCountLayer(layer)

	return
}

func (node *pinball) PinballTable(balls int) (distribution []int) {
	var leftEdge int
	node.GoToTheLeftEdge(&leftEdge)
	distribution = make([]int, -1*leftEdge*2+1)

	for i := 0; i < balls; i++ {
		var position int
		node.OneBallGoThrough(&position)
		distribution[position-leftEdge]++
	}

	return
}

func (node *pinball) OneBallGoThrough(position *int) {
	if node == nil {
		return
	}
	var next *pinball
	if node.right == 0 {
		fmt.Println("right is 0")
		fmt.Println(node.layer)
	}
	if node.left == 0 {
		fmt.Println("left is 0")
		fmt.Println(node.layer)
	}
	direction := GoRightSide(node.right)
	if direction {
		*position = *position + 1
		next = node.rightNode
	} else {
		*position = *position - 1
		next = node.leftNode
	}
	fmt.Println(">>>", *position, node.left, node.right, direction, node.layer)
	next.OneBallGoThrough(position)
	return
}

func (node *pinball) GoToTheRightEdge(position *int) {
	if node == nil {
		return
	}
	*position = *position + 1
	next := node.rightNode
	next.GoToTheRightEdge(position)
	return
}

func (node *pinball) GoToTheLeftEdge(position *int) {
	if node == nil {
		return
	}
	*position = *position - 1
	next := node.leftNode
	next.GoToTheLeftEdge(position)
	return
}

func GoRightSide(probability float64) (decision bool) {
	seed := time.Now().UnixNano()
	src := rand.NewSource(seed)
	r := rand.New(src)
	if probability == 0 {
		fmt.Println("error !")
	}
	if r.Float64() > probability {
		decision = true
	}
	return
}
