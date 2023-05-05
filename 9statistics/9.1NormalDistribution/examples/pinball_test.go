package examples

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Check_pinball(t *testing.T) {
	// Creates a root axle node
	root := axle{}
	root.axleTree(0, 10, true)

	// List the axle tree
	root.ListAxleTree(0)

	// Calls the comparator algorithm and count the number of axles at the specific depth
	require.Equal(t,
		root.algorithmComparatorForCountSpecificDepthAxle(10),
		root.CountSpecificDepthAxle(10),
	)

	// One ball moves right from the edge from the position 0
	var position int
	root.GoRightEdgeSide(&position)
	require.Equal(t, 12, position)

	// One ball moves left from the edge from the position 0
	position = 0
	root.GoToTheLeftEdge(&position)
	require.Equal(t, -12, position)

	// One ball falling layer by layer.
	position = 0
	root.OneBallGoThrough(&position)
	require.LessOrEqual(t, -12, position)
	require.GreaterOrEqual(t, 12, position)

	// 1000 balls falling
	distribution := root.PlayPinballTable(1000)

	// Prints the distribution
	fmt.Println("distribution:")
	fmt.Println(distribution)
}
