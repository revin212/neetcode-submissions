func isPalindrome(s string) bool {
	leftPointer := 0
	rightPointer := len(s) - 1

	for leftPointer < rightPointer {
		for leftPointer < rightPointer && !isAlphaNum(s[leftPointer]) {
			leftPointer++
		}
		for leftPointer < rightPointer && !isAlphaNum(s[rightPointer]) {
			rightPointer--
		}

		if toLowerByte(s[leftPointer]) != toLowerByte(s[rightPointer]) {
			return false
		}

		leftPointer++
		rightPointer--
	}
	return true
}

func isAlphaNum(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

func toLowerByte(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}