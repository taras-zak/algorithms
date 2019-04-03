package searching

func BinarySearch(arr []int, low, high, target int) int {
	if high < low { return -1}
	mid := low + (high - low) / 2
	switch {
	case arr[mid] == target:
		return mid
	case target < arr[mid]:
		return BinarySearch(arr, low, mid - 1, target)
	case  arr[mid] < target:
		return BinarySearch(arr, mid + 1, high, target)
	}
	return -1
}

func BinarySearchIter(arr []int, target int) int {
	low, high := 0, len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
