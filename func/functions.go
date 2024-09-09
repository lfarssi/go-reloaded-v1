package functions

import "fmt"

func TextFormated(s  []string) string{
	var res string
	for i, ch := range s {
		if ch == "," || ch == "!" || ch == "?" || ch == ":" || ch ==";" {
			fmt.Print("yes")
			if s[i-1] == " " {
				fmt.Print("zeb")
				res += res[:len(res)-1]
			    res += ch
			}
			res += ch
			res += " " 
			} else {
				res += ch + " "
		}
	}	
	return res
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