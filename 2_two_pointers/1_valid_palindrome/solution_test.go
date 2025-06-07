package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		name     string
		fn       func(string) bool
		s        string
		expected bool
	}{
		// IsPalindromeReverseString
		{"Reverse String - Palindrome with spaces/punctuations", IsPalindromeReverseString, "Was it a car or a cat I saw?", true},
		{"Reverse String - Not a palindrome", IsPalindromeReverseString, "tab a cat", false},
		{"Reverse String - Single character", IsPalindromeReverseString, "a", true},
		{"Reverse String - Empty string", IsPalindromeReverseString, "", true},
		{"Reverse String - Mixed case palindrome", IsPalindromeReverseString, "Able was I ere I saw Elba", true},
		{"Reverse String - Only non-alphanumeric", IsPalindromeReverseString, "!!!", true},
		{"Reverse String - Palindrome with numbers", IsPalindromeReverseString, "12321", true},
		{"Reverse String - Not palindrome with numbers", IsPalindromeReverseString, "12345", false},

		// IsPalindromeTwoPointers
		{"Two	Pointers - Palindrome with spaces/punctuations", IsPalindromeTwoPointers, "Was it a car or a cat I saw?", true},
		{"Two	Pointers - Not a palindrome", IsPalindromeTwoPointers, "tab a cat", false},
		{"Two	Pointers - Single character", IsPalindromeTwoPointers, "a", true},
		{"Two	Pointers - Empty string", IsPalindromeTwoPointers, "", true},
		{"Two	Pointers - Mixed case palindrome", IsPalindromeTwoPointers, "Able was I ere I saw Elba", true},
		{"Two	Pointers - Only non-alphanumeric", IsPalindromeTwoPointers, "!!!", true},
		{"Two	Pointers - Palindrome with numbers", IsPalindromeTwoPointers, "12321", true},
		{"Two	Pointers - Not palindrome with numbers", IsPalindromeTwoPointers, "12345", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.fn(c.s)
			if got != c.expected {
				t.Errorf("%s failed: expected %v, got %v", c.name, c.expected, got)
			}
		})
	}
}
