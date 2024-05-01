package main

// Import format and strings for functions
import (
	"fmt"
	"strings"
)

// Showcasing different string functionalities
func main() {
	// Concatenation
	str1 := "gopher"
	str2 := "world"
	concatenated := str1 + ", " + str2
	fmt.Println("Concatenation:", concatenated)

	// Length of a string
	fmt.Println("Length of str1:", len(str1))

	// Accessing individual characters
	fmt.Println("First character of str2:", string(str2[0]))

	// Substring
	fmt.Println("Substring of concatenated:", concatenated[3:8])

	// Check if a string contains a substring
	fmt.Println("Does concatenated contain 'world'?", strings.Contains(concatenated, "world"))

	// Count occurrences of a substring
	fmt.Println("Occurrences of 'l' in concatenated:", strings.Count(concatenated, "l"))

	// Replace part of a string
	replaced := strings.Replace(concatenated, "world", "yay!", -1)
	fmt.Println("Replaced string':", replaced)

	// Convert to lowercase
	fmt.Println("Lowercase of concatenated:", strings.ToLower(concatenated))

	// Convert to uppercase
	fmt.Println("Uppercase of concatenated:", strings.ToUpper(concatenated))

	// Trim whitespace from the beginning and end of a string
	whitespaceString := "  trimmed   "
	fmt.Println("Trimmed whitespace:", strings.TrimSpace(whitespaceString))

	// Split a string into substrings
	sentence := "I am a gopher."
	words := strings.Split(sentence, " ")
	fmt.Println("Words in sentence:", words)

	// Join substrings into a single string
	joined := strings.Join(words, "-")
	fmt.Println("Joined words:", joined)
}
