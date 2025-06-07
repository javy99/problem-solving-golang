package validpalindrome

// Time complexity: O(n)
// Space complexity: O(n)
func IsPalindromeReverseString(s string) bool {
	clean := make([]rune, 0, len(s))
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			clean = append(clean, r)
		} else if r >= 'A' && r <= 'Z' {
			clean = append(clean, r+32)
		}
	}
	n := len(clean)
	for i := 0; i < n/2; i++ {
		if clean[i] != clean[n-1-i] {
			return false
		}
	}
	return true
}

// Time complexity: O(n)
// Space complexity: O(1)
func IsPalindromeTwoPointers(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !((s[left] >= 'a' && s[left] <= 'z') || (s[left] >= 'A' && s[left] <= 'Z') || (s[left] >= '0' && s[left] <= '9')) {
			left++
		}
		for left < right && !((s[right] >= 'a' && s[right] <= 'z') || (s[right] >= 'A' && s[right] <= 'Z') || (s[right] >= '0' && s[right] <= '9')) {
			right--
		}
		l := s[left]
		r := s[right]
		if l >= 'A' && l <= 'Z' {
			l += 32
		}
		if r >= 'A' && r <= 'Z' {
			r += 32
		}
		if l != r {
			return false
		}
		left++
		right--
	}
	return true
}
