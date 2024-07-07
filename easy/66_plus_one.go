package easy

func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			if i-1 == -1 {
				newArray := make([]int, len(digits)+1)
				copy(newArray[1:], digits)
				newArray[0] = 1
				return newArray
			}
			continue
		}

		digits[i] += 1
		break
	}
	return digits
}
