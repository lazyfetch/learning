/*
On LeetCode it belongs to the easy class, when I solved it the initial execution
time was 40ms, but after some string formatting I came to a result of 9ms.

	What I remembered for myself

1. Always look at code in the direction of simplification
2. You don't always need hundreds of temporary variables (refers to the first point).
3. A regular expression requires simplification, and understanding why simplification gives you more speed is important.

The most important thing is not that I could see a solution, but that you understand
how the logic can be arranged and your code works, even if initially slow
*/

package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = re.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[len(s)-1-i] != s[i] {
			return false
		}
	}
	return true
}

// for test

func main() {

	msg := "A man, a plan, a canal: Panama" // type something

	println(isPalindrome(msg))

}
