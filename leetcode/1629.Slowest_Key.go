package main

import "fmt"

func slowestKey(releaseTimes []int, keysPressed string) byte {
	maxDelay := releaseTimes[0]
	key := keysPressed[0]
	for i := 1; i < len(releaseTimes); i++ {
		delay := releaseTimes[i] - releaseTimes[i-1]
		if delay == maxDelay {
			if key < keysPressed[i] {
				key = keysPressed[i]
			}
		} else if delay > maxDelay {
			maxDelay = delay
			key = keysPressed[i]
		}
	}
	return key
}

func main() {
	fmt.Println(string(slowestKey([]int{9, 29, 49, 50}, "cbcd")))
}
