package functions

import (
	"strings"
	"unicode"
)

func TextFormated(s []string) string {
	var res string
	ponc := ".?:!;,"
	for i, ch := range s {
		if ch == "," || ch == "!" || ch == "?" || ch == ":" || ch == ";" || ch == "." {
			if len(res) > 0 && res[len(res)-1] == ' ' {
				res = res[:len(res)-1]
			}
			if len(res) > 0 && !strings.ContainsAny(res, ponc) {
				res += ch
			} else {
				res += string(ch)
			}
			if i+1 < len(s) && s[i+1] != " " {
				res += " "
			}
		} else {
			res += ch + " "
		}
	}
	return strings.TrimSpace(res)
}

func Quote(s string) string {
	var result string
	wordInside := ""
	quoteOpen := false

	for i := 0; i < len(s); i++ {
		word := s[i]
		if word == '\'' {
			if quoteOpen {
				result += strings.TrimSpace(wordInside) + "'"
				quoteOpen = false
				wordInside = ""
			} else {
				quoteOpen = true
				wordInside = ""
				result += "'"
			}
			continue
		}
		if quoteOpen {
			wordInside += string(word)
		} else {
			result += string(word)
		}
	}
	if quoteOpen {
		result += wordInside
	}
	return result
}

func IsVowel(s string) bool {
	if s[0] == 'a' || s[0] == 'o' || s[0] == 'i' || s[0] == 'e' || s[0] == 'u' {
		return true
	}
	return false
}
func HandleVowel(arr []string, i int) []string {
	if strings.ToLower(arr[i]) == "a" && i+1 < len(arr) && IsVowel(arr[i+1]) {
		arr[i] += "n"
	} else if strings.ToLower(arr[i]) == "an" && i+1 < len(arr) && !IsVowel(arr[i+1]) {
		arr[i] = arr[i][:len(arr[i])-1]
	}
	return arr
}

func IsAlphabet(s byte) bool {
	return (s < 'a' || s > 'z') && (s < 'A' || s > 'Z') && (s < '0' || s > '9')
}

func IsWord(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
func Capitalize(word string) string {
	word = ToLower(word)
	for i := 0; i < len(word); i++ {
		word = ToUpper(string(word[0])) + word[1:]
	}
	return word
}

// func Split(s, sep string) []string {
// 	var result []string

// 	if len(sep) == 0 {
// 		return []string{s}
// 	}

// 	start := 0
// 	for i := 0; i < len(s); i++ {
// 		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
// 			result = append(result, s[start:i])
// 			// Move the start position past the separator
// 			start = i + len(sep)
// 			// Skip over the separator
// 			i += len(sep) - 1
// 		}
// 	}

// Add the last segment after the last separator
// 	result = append(result, s[start:])

// 	return result
// }

func ToUpper(s string) string {
	var res []rune
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			res = append(res, i-32)
		} else if i == 'é' {
			res = append(res, unicode.ToUpper(i))
		} else {
			res = append(res, i)
		}
	}
	return string(res)
}

func ToLower(s string) string {
	var res []rune
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			res = append(res, i+32)
		} else {
			res = append(res, i)
		}
	}
	return string(res)
}
