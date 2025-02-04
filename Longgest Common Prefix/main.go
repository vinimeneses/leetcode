package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, s := range strs[1:] {
		for !strings.HasPrefix(s, prefix) {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func main() {
	strs1 := []string{"flower", "flow", "flight"}
	fmt.Printf("Entrada: %v\nSaída: %q\n\n", strs1, longestCommonPrefix(strs1))

	strs2 := []string{"dog", "racecar", "car"}
	fmt.Printf("Entrada: %v\nSaída: %q\n", strs2, longestCommonPrefix(strs2))
}
