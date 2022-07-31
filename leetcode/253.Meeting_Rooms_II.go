package main

import (
	"fmt"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	start := make([]int, 0, len(intervals))
	for _, interval := range intervals {
		start = append(start, interval[0])
	}
	finish := make([]int, 0, len(intervals))
	for _, interval := range intervals {
		finish = append(finish, interval[1])
	}
	sort.Ints(start)
	sort.Ints(finish)

	startPointer := 0
	finishPointer := 0
	roomCount := 0
	maxRoomCount := 0
	for startPointer < len(intervals) && finishPointer < len(intervals) {
		if start[startPointer] < finish[finishPointer] {
			roomCount++
			if roomCount > maxRoomCount {
				maxRoomCount = roomCount
			}
			startPointer++
		} else if start[startPointer] >= finish[finishPointer] {
			roomCount--
			finishPointer++
		}
	}
	return maxRoomCount
}

func main() {
	fmt.Println(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
}
