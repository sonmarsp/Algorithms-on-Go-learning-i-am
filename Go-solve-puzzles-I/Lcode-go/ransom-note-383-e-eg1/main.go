package main

import "fmt"

func main() {

	//Input: ransomNote = "a", magazine = "b"
	//Output: false
	//Input: ransomNote = "aa", magazine = "aab"
	//Output: true

	fmt.Println(canConstruct("a", "b"))    // Output: false
	fmt.Println(canConstruct("aa", "ab"))  // Output: false
	fmt.Println(canConstruct("aa", "aab")) // Output: true

}

func canConstruct(ransomNote string, magazine string) bool {
	// Create a map to count the occurrences of each character in the magazine
	charCount := make(map[rune]int)

	// Count each character in the magazine
	for _, char := range magazine {
		charCount[char]++
	}

	// Check if the ransomNote can be constructed
	for _, char := range ransomNote {
		if charCount[char] > 0 {
			charCount[char]--
		} else {
			return false
		}
	}

	return true
}
