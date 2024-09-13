package functions

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// formater le text like this are , rare ... reae => are, rare... reae
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
// handle if there is a word between ' make it a quote like ' and ' => 'and'
func HandleQuote(s string) string {
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

// check if the string if vowel
func IsVowel(s string) bool {
	if s[0] == 'a' || s[0] == 'o' || s[0] == 'i' || s[0] == 'e' || s[0] == 'u' {
		return true
	}
	return false
}

// when the string after a is vowel i replace it with an , if the string after an is not vowel i replace it with a
func HandleVowel(s string) string {
	arr := strings.Fields(s)
	for i := 0; i < len(arr); i++ {
		if strings.ToLower(arr[i]) == "a" && i+1 < len(arr) && IsVowel(arr[i+1]) {
			arr[i] += "n"
		} else if strings.ToLower(arr[i]) == "an" && i+1 < len(arr) && !IsVowel(arr[i+1]) {
			arr[i] = arr[i][:len(arr[i])-1]
		}
	}
	str := ""
	for _, ch := range arr {
		str += ch + " "
	}
	return str
}
/*
the function is also for separate punctuation with spaces but the most important rule 
is to handle the between parentheses flag 
explination: adding the flag like with the space to check the format is it correct
and the number without spaces 
*/
func HandleParenthese(s string) string {
	t := ""
	insideParenthese := false
	beforeVergule := true
	for _, v := range s {
		if v == '(' {
			t += " " + string(v)
			insideParenthese = true
		} else if v == ')' {
			t += string(v) + " "
			insideParenthese = false

		} else {
			if insideParenthese {
				if v == ',' {
					beforeVergule = false
				}
				if beforeVergule{
					t += string(v)
					} else {
						if v != ' '{
							t+=string(v)
						}
					}
			} else {
				if v == ',' || v == '.' || v == ':' || v == '!' || v == '?' || v == ';' {
					t += " " + string(v) + " "
				} else {
					t += string(v)
				}
			}
		}

	}
	return t
}

/*
in this function i put the item between the () and modify it with my rules
first rule if the flag is without number i add 1 default number in it
second if the flag has a number i split with comma and add with  space after comma 
*/
func HandleParentheseParam(s []string) string {
	res2 := ""
	for _, item := range s {
		if strings.HasPrefix(item, "(") && strings.HasSuffix(item, ")") {
			content := item[1 : len(item)-1]
			if strings.Contains(content, ",") {
				arr:= strings.Split(content, ",")
				content = ""
				for i, c:= range arr{
					if !IsWord(c){
						arr[i] = ", " + c
					}
					content += arr[i]
				}
				
				res2 += "(" + content + ") "
			} else {
				switch content {
				case "cap":
					res2 += "(cap, 1) "
				case "low":
					res2 += "(low, 1) "
				case "up":
					res2 += "(up, 1) "
				case "hex", "bin":
					res2 += " (" + content + ") "
				default:
					res2 += " (" + content + ") "
				}
			}
		} else {
			res2 += item + " "
		}

	}

	return res2
}
/*
this is the most important function in my program
i handle the parenthese i after that i convert it to a slice to handle the parameter inside the parenthese
after that i convert it to a slice again to handle the flag based on the instruction 
*/
func HandleFlag(s string) string {
	str := HandleParenthese(s)
	arr1 := strings.Fields(string(str))
	res2 := HandleParentheseParam(arr1)
	arr := strings.Fields(res2)
	for i := 0; i < len(arr); i++ {
		if arr[i] == "(cap," || arr[i] == "(low," || arr[i] == "(up," {
			if i+1 == len(arr) {
				continue
			}
			//remove ( and , from the flag
			arr[i] = strings.TrimPrefix(arr[i], "(")
			arr[i] = strings.TrimSuffix(arr[i], ",")
			//remove )  from the number
			arr[i+1] = strings.TrimSuffix(arr[i+1], ")")
			var err error
			var nb int
			// convert the number to int value
			nb, err = strconv.Atoi(arr[i+1])
			if err != nil {
				fmt.Println("msg err : The params is not a number ", err)
				continue
			}
			// store temporairement the index
			temp := i

			// loop for the number in the flag
			for j := 1; j <= nb; j++ {
				// break the loop before err out of range
				if i-j < 0 {
					break
				}
				// skip if it's not a word or it's a flag
				if !IsWord(arr[i-j]) || arr[i-j]== "cap" || arr[i-j] == "low" || arr[i-j]== "up" || arr[i-j]== "(bin)" || arr[i-j]== "(hex)" {
					continue
				}
				// apply the changes based on the number
				if arr[i] == "cap" {
					arr[i-j] = Capitalize(arr[i-j])
				} else if arr[i] == "low" {
					arr[i-j] = ToLower(arr[i-j])
				} else if arr[i] == "up" {
					arr[i-j] = ToUpper(arr[i-j])
				}
			}
			// remove the flag and the number 
			arr = append(arr[:i],arr[i+2:]... )
			// return one step to continue from the correct position
			i = temp-1
		} else if arr[i] == "(bin)" {
			// convert the string to 64 bit integer base 2
			integer, err := strconv.ParseInt(arr[i-1], 2, 64)
			if err != nil {
				fmt.Println("you can't convert")
				continue
			}
			arr[i-1] = strconv.Itoa(int(integer))
			arr = append(arr[:i], arr[i+1:]...)
			i--

		} else if arr[i] == "(hex)" {
			// convert the string to 64 bit integer base 16
			integer, err := strconv.ParseInt(arr[i-1], 16, 64)
			if err != nil {
				fmt.Println("you can't convert")
				continue
			}
			arr[i-1] = strconv.Itoa(int(integer))
			arr = append(arr[:i], arr[i+1:]...)
			i--

		}
	}
	return strings.Join(arr, " ")
}

// check if the string is a letter or not
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
	for _, ch := range s {
		if unicode.IsLetter(ch) {
			res = append(res, unicode.ToUpper(ch))
		} else {
			res = append(res, ch)
		}
	}
	return string(res)
}

func ToLower(s string) string {
	var res []rune
	for _, i := range s {
		if unicode.IsLetter(i) {
			res = append(res,unicode.ToLower(i))
		} else {
			res = append(res, i)
		}
	}
	return string(res)
}
