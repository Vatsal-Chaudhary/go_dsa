package array

import (
	"math"
)

func earliestFinishTime(landStartTime, landDuration, waterStartTime, waterDuration []int) int {
	if len(landStartTime) != len(landDuration) || len(waterDuration) != len(waterStartTime) {
		return -1
	}

	ans := math.MaxInt

	for i, landStart := range landStartTime {
		for j, waterStart := range waterStartTime {

			landFinish := landStart + landDuration[i]

			waterRideStart := max(waterStart, landFinish)

			plan1Finish := waterRideStart + waterDuration[j]

			waterFinish := waterStart + waterDuration[j]

			landRideStart := max(landStartTime[i], waterFinish)

			plan2Finish := landRideStart + landDuration[i]

			ans = min(ans, plan1Finish)
			ans = min(ans, plan2Finish)

		}
	}

	return ans
}
