package utils

func Min(elements ...int) int {
	minimum := elements[0]
	for _, el := range elements {
		if minimum > el {
			minimum = el
		}
	}
	return minimum
}
