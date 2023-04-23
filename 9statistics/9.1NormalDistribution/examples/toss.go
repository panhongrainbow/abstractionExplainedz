package examples

import (
	"math/rand"
)

/*
TossTenTimes tosses 10 coins and return the number of heads.
*/
func TossTenTimes() (heads int) {
	for i := 0; i < 10; i++ {
		if rand.Float64() < 0.5 {
			heads++
		}
	}
	return
}

// OneThousandCount counts the average number and the distribution of executions a thousand times.
func OneThousandCount() (distribution [11]int, average float32) {
	var total float32
	for i := 0; i < 1000; i++ {
		heads := TossTenTimes()
		distribution[heads]++
		total = total + float32(heads)
	}
	average = total / 1000.0
	return
}
