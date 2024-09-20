package functions

import (
	"fmt"
	"strconv"
	"strings"
)

// formater le text like this are , rare ... reae => are, rare... reae
func TextFormated(s []string) string {
	var res string
	ponc := ".?:!;,l"
	for i, ch := range s {
		if ch == "," || ch == "!" || ch == "?" || ch == ":" || ch == ";" || ch == "." {
			if len(res) > 0 && res[len(res)-1] == ' ' {
				res = res[:len(res)-1]
			}
			if len(res) > 0 && !strings.ContainsAny(res, ponc) {
				res += ch
			} else {
				res += ch
			}
			if i < len(s)-1 && s[i+1] != " " {
				res += " "
			}
		} else {
			if strings.ContainsAny(ch, ponc){
				res += Freeze(ch) + " "
			} else {
				res += ch + " "
			}
			
		}
	}
	return strings.TrimSpace(res)
}
func Freeze(s string) string {
	var res string
	for _, ch := range s {
		if ch == ',' || ch == '!' || ch == '?' || ch == ':' || ch == ';' || ch == '.' {
			res += string(ch) + string(' ')
		} else {
			res += string(ch)
		}
	}
	return res
}

// handle if there is a word between ' make it a quote like ' and ' => 'and'
func HandleQuote(s string) string {
	var result string
	wordInside := ""
	quoteOpen := false

	for i := 0; i < len(s); i++ {
		word := s[i]
		if word == '\'' {
			if i > 0 && i < len(s)-1 && ((s[i-1] >= 'a' && s[i-1] <= 'z') || (s[i-1] >= 'A' && s[i-1] <= 'Z')) && ((s[i+1] >= 'a' && s[i+1] <= 'z') || (s[i+1] >= 'A' && s[i+1] <= 'Z')) {
				result += string(word)
				continue
			}
			if quoteOpen {
				result += strings.TrimSpace(wordInside) + "'"
				if i < len(s)-1 && s[i+1] != ' ' {
					result += " "
				}
				quoteOpen = false
				wordInside = ""
			} else {
				quoteOpen = true
				wordInside = ""
				if i > 0 && s[i-1] != ' ' {
					result += " "
				}
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
	if s[0] == 'a' || s[0] == 'o' || s[0] == 'i' || s[0] == 'e' || s[0] == 'u' || s[0] == 'h' {
		return true
	}
	return false
}

// when the string after a is vowel i replace it with an , if the string after an is not vowel i replace it with a
func HandleVowel(s string) string {
	arr := strings.Fields(s)
	for i := 0; i < len(arr); i++ {
		if strings.ToLower(arr[i]) == "a" && i+1 < len(arr) && IsVowel(strings.ToLower(arr[i+1])) {
			arr[i] += "n"
		} else if strings.ToLower(arr[i]) == "an" && i+1 < len(arr) && !IsVowel(strings.ToLower(arr[i+1])) && IsWord(strings.ToLower(arr[i+1])) && i < len(arr)-1 {
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
// func HandleParenthese(s string) string {
// 	t := ""
// 	insideParenthese := false
// 	beforeVergule := true
// 	for _, v := range s {
// 		if v == '(' {
// 			t += string(v)
// 			insideParenthese = true
// 		} else if v == ')' {
// 			t += string(v)
// 			insideParenthese = false

// 		} else {
// 			if insideParenthese {
// 				if v == ',' {
// 					beforeVergule = false
// 				}
// 				if beforeVergule {
// 					t += string(v)
// 				} else {
// 					if v != ' ' {
// 						t += string(v)
// 					}
// 				}
// 			} else {
// 				if v == ',' || v == '.' || v == ':' || v == '!' || v == '?' || v == ';' {
// 					t += " " + string(v) + " "
// 				} else {
// 					tok += string(v)
// 				}
// 			}
// 		}

// 	}
// 	return t
// }

// /*
// in this function i put the item between the () and modify it with my rules
// first rule if the flag is without number i add 1 default number in it
// second if the flag has a number i split with comma and add with  space after comma
// */
// func HandleParentheseParam(s []string) string {
// 	res2 := ""
// 	for _, item := range s {
// 		if strings.HasPrefix(item, "(") && strings.HasSuffix(item, ")") {
// 			content := item[1 : len(item)-1]
// 			if strings.Contains(content, ",") {
// 				arr := strings.Split(content, ",")
// 				content = arr[0] + ", " + arr[1]
// 				res2 += "(" + content + ") "
// 			} else {
// 				switch content {
// 				case "cap":
// 					res2 += "(cap, 1) "
// 				case "low":
// 					res2 += "(low, 1) "
// 				case "up":
// 					res2 += "(up, 1) "
// 				case "hex", "bin":
// 					res2 += " (" + content + ") "
// 				default:
// 					res2 += " (" + content + ") "
// 				}
// 			}
// 		} else {
// 			res2 += item + " "
// 		}

// 	}

// 	return res2
// }

/*
this is the most important function in my program
i handle the parenthese i after that i convert it to a slice to handle the parameter inside the parenthese
after that i convert it to a slice again to handle the flag based on the instruction
*/
func HandleFlag(s string) string {
	//str := HandleParenthese(s)
	//arr1 := strings.Fields(string(str))
	//res2 := HandleParentheseParam(arr1)
	arr := strings.Fields(s)
	for i := 0; i < len(arr); i++ {
		if arr[i] == "(bin)" {
			arr = append(arr[:i], arr[i+1:]...)
			// convert the string to 64 bit integer base 2
			if i-1 < 0 {
				break
			}
			integer, err := strconv.ParseInt(arr[i-1], 2, 32)
			if err != nil {
				fmt.Println("you can't convert")
				continue
			}
			arr[i-1] = strconv.Itoa(int(integer))
			i--

		} else if arr[i] == "(hex)" {
			arr = append(arr[:i], arr[i+1:]...)
			// convert the string to 64 bit integer base 16
			if i-1 < 0 {
				break
			}
			integer, err := strconv.ParseInt(arr[i-1], 16, 64)
			if err != nil {
				fmt.Println("you can't convert")
				continue
			}
			arr[i-1] = strconv.Itoa(int(integer))
			i--

		} else if arr[i] == "(cap)" {
			if i > 0 {
			arr[i-1] = Capitalize(arr[i-1])
		}
			arr = append(arr[:i], arr[i+1:]...)
			i--
		} else if arr[i] == "(low)" {
			if i > 0 {
			arr[i-1] = strings.ToLower(arr[i-1])
		}
			arr = append(arr[:i], arr[i+1:]...)
			i--
		} else if arr[i] == "(up)" {
			if i > 0 {

				arr[i-1] = strings.ToUpper(arr[i-1])
			}
			arr = append(arr[:i], arr[i+1:]...)
			i--
		} else if (arr[i] == "(cap," || arr[i] == "(low," || arr[i] == "(up,") && strings.ContainsAny(arr[i+1],")") {
			if i+1 == len(arr) {
				continue
			}
			var err error
			var nb int
			//remove )  from the number
			// convert the number to int value
			nb, err = strconv.Atoi(strings.TrimSuffix(arr[i+1], ")"))

			if err != nil {
				fmt.Println("msg err : The params is not a number ", err)
				continue
			}
			// store temporairement the index
			temp := i
			if nb < 0 {
				continue
			}
			// loop for the number in the flag
			for j := 1; j <= nb; j++ {
				// break the loop before err out of range
				if i-j < 0 {
					break
				}
				// apply the changes based on the number
				if arr[i] == "(cap," {
					arr[i-j] = Capitalize(arr[i-j])
				} else if arr[i] == "(low," {
					arr[i-j] = strings.ToLower(arr[i-j])
				} else if arr[i] == "(up," {
					arr[i-j] = strings.ToUpper(arr[i-j])

				}
				i = temp
			}
			// remove the flag and the number
			arr = append(arr[:i], arr[i+2:]...)
			// return one step to continue from the correct position
			i = temp - 1
		}
	}
	return strings.Join(arr, " ")
}

// this function is used to skip the non words runes
// func Word(s string) string {
// 	for _, r := range s {
// 		if !(r >= 'a' && r <= 'z') && !(r >= 'A' && r <= 'Z') {
// 			continue
// 		}
// 	}
// 	return s
// }

// check if the string is a letter or not
// func IsWord(s string) bool {
// 	ponc := ",;:!?."
// 	return !strings.ContainsAny(s, ponc)
// }

func Capitalize(word string) string {
	word = strings.ToLower(word)
	for i := 0; i < len(word); i++ {
		word = strings.ToUpper(string(word[0])) + word[1:]
	}
	return word
}

func IsWord(s string) bool {
	if s == "'" || s == "!" || s== "," || s== "." || s=="\n" || s=="?" || s==";" {
		return false
	}
	return true
}