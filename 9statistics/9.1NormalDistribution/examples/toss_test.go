package examples

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
Test_Check_Toss_And_Count contains two subtests:
TossTenTimes and OneThousandCount. It checks the values of heads, average, and distribution array elements.
*/
func Test_Check_Toss_And_Count(t *testing.T) {
	// The TossTenTimes subtest focuses on testing the TossTenTimes function
	t.Run("TossTenTimes", func(t *testing.T) {
		// Assign the return value of the TossTenTimes function to the heads variable
		heads := TossTenTimes()
		for i := 0; i < 50; i++ {
			// Check whether the value of heads is between 0 and 10
			require.GreaterOrEqual(t, 10, heads, "heads should be less or equal than 10")
			require.LessOrEqual(t, 0, heads, "heads should be greater or equal than between 0")
		}
	})
	// The OneThousandCount subtest focuses on testing the OneThousandCount function
	t.Run("OneThousandCount", func(t *testing.T) {
		// Assign the return values of the OneThousandCount function to the distribution and average variables
		distribution, average := OneThousandCount()
		// Check whether the value of average is between 4 and 6
		require.Greater(t, float32(6), average, "average should be less or equal than 10")
		require.Less(t, float32(4), average, "average should be greater or equal than between 0")
		// Check whether the elements in the distribution array are arranged in the expected order
		require.Greater(t, distribution[5], distribution[4], "distribution[5] should be greater than distribution[4]")
		require.Greater(t, distribution[5], distribution[6], "distribution[5] should be greater than distribution[6]")
		require.Greater(t, distribution[4], distribution[3], "distribution[4] should be greater than distribution[3]")
		require.Greater(t, distribution[4], distribution[7], "distribution[4] should be greater than distribution[7]")
		require.Greater(t, distribution[6], distribution[3], "distribution[4] should be greater than distribution[3]")
		require.Greater(t, distribution[6], distribution[7], "distribution[4] should be greater than distribution[7]")
		require.Greater(t, distribution[3], distribution[2], "distribution[3] should be greater than distribution[2]")
		require.Greater(t, distribution[3], distribution[8], "distribution[3] should be greater than distribution[8]")
		require.Greater(t, distribution[7], distribution[2], "distribution[7] should be greater than distribution[2]")
		require.Greater(t, distribution[7], distribution[8], "distribution[7] should be greater than distribution[8]")
		require.Greater(t, distribution[2], distribution[1], "distribution[2] should be greater than distribution[1]")
		require.Greater(t, distribution[2], distribution[9], "distribution[2] should be greater than distribution[9]")
		require.Greater(t, distribution[8], distribution[1], "distribution[8] should be greater than distribution[1]")
		require.Greater(t, distribution[8], distribution[9], "distribution[8] should be greater than distribution[9]")
		require.Greater(t, distribution[1], distribution[0], "distribution[1] should be greater than distribution[0]")
		require.Greater(t, distribution[1], distribution[10], "distribution[1] should be greater than distribution[10]")
		require.Greater(t, distribution[9], distribution[0], "distribution[9] should be greater than distribution[0]")
		require.Greater(t, distribution[9], distribution[10], "distribution[9] should be greater than distribution[10]")
	})
}
