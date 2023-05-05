package examples

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/*
axle is not a binary tree.
It is only simulating the ratios of pinballs falling to the left or to the right when the pinballs hit an axle.
*/
type axle struct {
	depth         int     // depth is the depth of the axle
	ratioToLeft   float64 // Ratio for pinballs to go left side
	ratioToRight  float64 // Ratio for pinballs to go right side, RatioToRight = 1 - RatioToLeft
	nextLeftAxle  *axle   // When pinballs hit the next left axle, use it
	nextRightAxle *axle   // When pinballs hit the next right axle, use it
}

// axleTree constructs axle tree by calling makeNextBothNode and itself recursively.
func (node *axle) axleTree(depthNo, depthLimit int, isRealWorld bool) {
	// decide whether to use a real-world machine or an idealized one
	if isRealWorld {
		node.ratioToLeft = rand.Float64()
	} else {
		node.ratioToLeft = 0.5
	}
	node.ratioToRight = 1 - node.ratioToLeft

	// Set ratioToLeft and ratioToRight for next 2 axles
	node.makeNextBothNode(depthNo, depthLimit, isRealWorld)
}

// makeNextBothNode makes the next left axle and the next right axle.
func (node *axle) makeNextBothNode(depth, maxDepth int, isRealWorld bool) {
	// Return if exceeding depth maxDepth
	if depth > maxDepth {
		return
	}
	node.depth = depth

	// Set ratioToLeft and ratioToRight for the left axle
	node.nextLeftAxle = &axle{}
	// decide whether to use a real-world machine or an idealized one
	if isRealWorld {
		node.nextLeftAxle.ratioToLeft = rand.Float64()
	} else {
		node.nextLeftAxle.ratioToLeft = 0.5
	}
	node.nextLeftAxle.ratioToRight = 1 - node.nextLeftAxle.ratioToLeft

	// Set ratioToLeft and ratioToRight for the right axle
	node.nextRightAxle = &axle{}
	// decide whether to use a real-world machine or an idealized one
	if isRealWorld {
		node.nextRightAxle.ratioToLeft = rand.Float64()
	} else {
		node.nextRightAxle.ratioToLeft = 0.5
	}
	node.nextRightAxle.ratioToRight = 1 - node.nextRightAxle.ratioToLeft

	// Go to the next two axles
	node.nextLeftAxle.axleTree(depth+1, maxDepth, isRealWorld)
	node.nextRightAxle.axleTree(depth+1, maxDepth, isRealWorld)
	return
}

// ListAxleTree recursively traverses an axle tree and prints node info.
func (node *axle) ListAxleTree(tab int) {
	// Checks for a null node and returns early if null
	if node == nil {
		return
	}

	// Prints the node info with a level of indentation based on the tab value
	fmt.Printf("%s", strings.Repeat("|_____", tab))
	fmt.Printf("Layer: %d, Left: %f, Right: %f\n", node.depth, node.ratioToLeft, node.ratioToRight)

	// Recursively calls the next left axle
	node.nextLeftAxle.ListAxleTree(tab + 1)

	//Recursively calls the next right axle
	node.nextRightAxle.ListAxleTree(tab + 1)
}

// CountSpecificDepthAxle counts the number of axles at the specific depth.
func (node *axle) CountSpecificDepthAxle(specificDepth int) (count int) {
	count = 1 << specificDepth
	return
}

/*
algorithmComparatorForCountSpecificDepthAxle also counts the number of axles at the specific depth.
However, it goes through the axle tree recursively and counts the number of axles at the specific depth.
对数器，用来检验CountLastDepthAxle的正确性
找遍所有的节点，看看有没有深度为specificDepth的节点
*/
func (node *axle) algorithmComparatorForCountSpecificDepthAxle(depth int) (count int) {
	if node == nil {
		return
	}

	// If the depth of the node is equal to the specific depth, count it.
	if node.depth == depth {
		count++
	}

	// Recursively calls the next left axle
	count = count + node.nextLeftAxle.algorithmComparatorForCountSpecificDepthAxle(depth)

	// Recursively calls the next right axle
	count = count + node.nextRightAxle.algorithmComparatorForCountSpecificDepthAxle(depth)

	return
}

// PlayPinballTable is to simulate a person to play pinball.
func (node *axle) PlayPinballTable(balls int) (distribution []int) {
	// Check the position of the ball when it goes to the right edge side
	var rightEdge int
	node.GoRightEdgeSide(&rightEdge)

	// capacity = 1 + rightEdge*2 because the distribution is from -rightEdge to rightEdge
	distribution = make([]int, 1+rightEdge*2)

	for i := 0; i < balls; i++ {
		var position int
		node.OneBallGoThrough(&position)
		distribution[position+rightEdge]++ // add rightEdge to make the position positive
	}

	return
}

// GoRightEdgeSide directs the pinball to the farthest right edge until reaching a null axle. (球会到最整个弹珠台的最右边)
func (node *axle) GoRightEdgeSide(position *int) {
	// When the node is on the right edge side, it will be nil. Then return
	if node == nil {
		return
	}

	// When the ball goes to the right side, the position will be added 1
	*position = *position + 1
	next := node.nextRightAxle
	next.GoRightEdgeSide(position)
	return
}

// GoToTheLeftEdge directs the pinball to the farthest left edge until reaching a null axle. (球会到最整个弹珠台的最左边)
func (node *axle) GoToTheLeftEdge(position *int) {
	// When the node is on the right edge side, it will be nil. Then return
	if node == nil {
		return
	}

	// When the ball goes to the left side, the position will be subtracted 1
	*position = *position - 1
	next := node.nextLeftAxle
	next.GoToTheLeftEdge(position)
	return
}

// OneBallGoThrough simulates the ball falling layer by layer until reaching the bottom.
func (node *axle) OneBallGoThrough(position *int) {
	if node == nil {
		return
	}
	var next *axle
	if GoRightSide(node.ratioToRight) {
		// When the ball goes to the right side, the position will be added 1
		*position = *position + 1
		next = node.nextRightAxle
	} else {
		// When the ball goes to the left side, the position will be subtracted 1
		*position = *position - 1
		next = node.nextLeftAxle
	}

	// Call it recursively
	next.OneBallGoThrough(position)
	return
}

// GoRightSide is to simulates a ball hits the axle, it falls to the left or to the right
func GoRightSide(probability float64) (decision bool) {
	if probability == 0 {
		fmt.Println("error !")
	}

	// randomly generates a value to determine whether to go right or not
	seed := time.Now().UnixNano()
	src := rand.NewSource(seed)
	random := rand.New(src)
	if random.Float64() > probability {
		decision = true
	}
	return
}
