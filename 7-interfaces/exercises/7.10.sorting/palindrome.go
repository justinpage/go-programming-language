package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/tabwriter"
)

type sequence []rune

func (s sequence) Len() int           { return len(s) }
func (s sequence) Less(i, j int) bool { return s[i] < s[j] }
func (s sequence) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	examples := []string{
		"taco cat", "madam", "racecar",
		"A man, a plan, a canal, Panama!",
		"Palindrome",
		"Was it a car or a cat I saw?",
		"No 'x' in Nixon",
		"Justin Page",
	}

	const format = "%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)

	fmt.Fprintf(tw, format, "Sequence", "Palindrome")
	fmt.Fprintf(tw, format, "--------", "----------")

	for _, ex := range examples {
		fmt.Fprintf(tw, format, ex, isPalindrome(sequence(sanitize(ex))))
	}

	tw.Flush()
}

func sanitize(s string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(strings.ToLower(s), "")
}

func isPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		if j := s.Len() - i - 1; !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
	}
	return true
}
