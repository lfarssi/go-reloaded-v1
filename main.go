package main

import (
	"fmt"
	functions "go-reloaded/func"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	length := len(args)
	if length < 2 {
		fmt.Println("less arguments you need to enter the input and the output files")
		os.Exit(1)
	} else if length > 2 {
		fmt.Println("too much arguments")
		os.Exit(1)
	} else {
		in := args[0]
		out := args[1]
		inExt := ""
		index := -1
		outExt := ""
		for i := len(in) - 1; i > 0; i-- {
			if in[i] == '.' {
				index = i
			}
		}
		inExt = in[index+1:]
		for j := len(out) - 1; j > 0; j-- {
			if out[j] == '.' {
				index = j
			}
		}
		outExt = out[index+1:]
		if inExt != "txt" || outExt != "txt" {
			fmt.Println("the extension must be .txt")
			os.Exit(1)
		} else {
			res, err := os.ReadFile(in)
			if err != nil {
				fmt.Println("Err msg: ", err)
				return
			}
			t := ""
			insideParenthese := false
			for _, v := range res {
				if v == '(' {
					t += " " + string(v)
					insideParenthese = true
				} else if v == ')' {
					t += string(v) + " "
					insideParenthese = false

				} else {
					if insideParenthese {
						if v == ',' {
							t += string(v)
						} else if v != ' ' {
							t += string(v)
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

			arr1 := strings.Fields(string(t))
			res2 := ""
			for _, item := range arr1 {
				if strings.HasPrefix(item, "(") && strings.HasSuffix(item, ")") {
					content := item[1 : len(item)-1]
					if strings.Contains(content, ",") {
						res2 += "(" + content + ") "
					} else {
						switch content {
						case "cap":
							res2 += "(cap,1) "
						case "low":
							res2 += "(low,1)"
						case "up":
							res2 += "(up,1)"
						case "hex", "bin":
							res2 += "(" + content + ") "
						default:
							res2 += "(" + content + ") "
						}
					}
				} else {
					res2 += item + " "
				}

			}
			arr := strings.Fields(res2)
			for i := 0; i < len(arr); i++ {
				insideParenthese2 := false
				if strings.HasPrefix(arr[i], "(") && strings.HasSuffix(arr[i], ")") {
					insideParenthese2 = true
				} else {
					insideParenthese2 = false
				}
				var Akwas []string
				var action string
				var nb int
				if insideParenthese2 {
					arr[i] = strings.Trim(arr[i], "()")
					Akwas = strings.Split(arr[i], ",")
					arr[i] = ""
					i--
					action = Akwas[0]
					if len(Akwas) == 2 {
						nb, err = strconv.Atoi(Akwas[1])
						if err != nil {
							fmt.Println("msg err : not a number ", err)
							continue
						}
					}
				}
				if action == "cap" || action == "low" || action == "up" {
					for j := 1; j <= nb; j++ {
						if i-j < 0 {
							break
						}
						if !functions.IsWord(arr[i-j]) {
							continue
						}
						if action == "cap" {
							arr[i-j] = functions.Capitalize(arr[i-j])

						} else if action == "low" {
							arr[i-j] = functions.ToLower(arr[i-j])

						} else if action == "up" {
							arr[i-j] = functions.ToUpper(arr[i-j])

						}
					}

				} else if action == "bin" {
					integer, err := strconv.ParseInt(arr[i-1], 2, 64)
					if err != nil {
						i -= 2
						continue
					}
					arr[i-1] = strconv.FormatInt(integer, 10)
					arr[i] = ""
					i--

				} else if action == "hex" {
					integer, err := strconv.ParseInt(arr[i-1], 16, 64)
					if err != nil {
						fmt.Println("you can't convert")
						continue
					}
					arr[i-1] = strconv.FormatInt(integer, 10)
					arr[i] = ""
					i--

				}
			}
			str2 := functions.TextFormated(arr)
			resulta := functions.Quote(str2)
			finalres := ""
			lines := strings.Split(string(resulta), "\n")
			for _, line := range lines {
				finalres += line + "\n"
			}
			err = os.WriteFile(out, []byte(finalres), 0o644)
			if err != nil {
				fmt.Println("Error writing file:", err)
				return
			}
			fmt.Println("bravoo!!!")
		}
	}
}
