package solute

import (
	"strconv"
	"strings"
)

var days = []int{
	31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
}

// [统计共同度过的日子数](https://leetcode.cn/problems/count-days-spent-together/)
func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	if arriveAlice > leaveBob || arriveBob > leaveAlice {
		return 0
	}

	left, right := arriveAlice, leaveAlice
	if arriveAlice < arriveBob {
		left = arriveBob
	}

	if leaveAlice > leaveBob {
		right = leaveBob
	}

	fieldLeft := strings.Split(left, "-")
	fieldRight := strings.Split(right, "-")

	monthleft, _ := strconv.Atoi(fieldLeft[0])
	dayleft, _ := strconv.Atoi(fieldLeft[1])
	monthright, _ := strconv.Atoi(fieldRight[0])
	dayright, _ := strconv.Atoi(fieldRight[1])

	if monthleft == monthright {
		return dayright - dayleft + 1
	}

	result := days[monthleft-1] - dayleft + 1 + dayright

	for j := monthleft + 1; j < monthright; j++ {
		result += days[j-1]
	}
	return result
}
