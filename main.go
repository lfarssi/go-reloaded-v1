package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"go-reloaded/func"
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
			file, err := os.Open(in)
			if err != nil {
				fmt.Println("Err msg: ", err)
				return
			}
			defer file.Close()
			res, err := io.ReadAll(file)
			if err != nil {
				fmt.Println("err msg :", err)
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
						if v > 32 && v < 48 {
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
						// Apply rules to update the content
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
			//fmt.Println(arr)
			for i := 0; i < len(arr); i++ {
				insideParenthese2 := false
				if strings.HasPrefix(arr[i], "(") && strings.HasSuffix(arr[i], ")") {
					insideParenthese2 = true
				} else{
					insideParenthese2 = false
				}
				var Akwas []string
				var action string
				var nb int
				if insideParenthese2 {
					arr[i]= strings.Trim(arr[i],"()")
					Akwas = strings.Split(arr[i],",")
					arr[i]=""
					action = Akwas[0]
					if len(Akwas)==2{
						nb, err = strconv.Atoi(Akwas[1])
						if err != nil {
							fmt.Println("msg err : not a number ", err)
							continue
						}
					}
				}
				if action == "cap" || action == "low" || action == "up" {
					for j:=1 ; j <= nb ; j++ {
						if i-j-1 < 0 {
							break
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
						fmt.Println("you can't convert")
						return
					}
					arr[i-1] = strconv.FormatInt(integer, 10)

				} else if action == "hex" {
					integer, err := strconv.ParseInt(arr[i-1], 16, 64)
					if err != nil {
						fmt.Println("you can't convert")
						return
					}
					arr[i-1] = strconv.FormatInt(integer, 10)

				}
			}

			/*lAkwas := false
			keyword := ""
			for i := 0; i < len(res); i++ {
				index := 0
				if res[i] == '(' {
					lAkwas = true
					index = res[i]
					continue
				}
				if res[i] == ')' {
					lAkwas = false
					continue
				}
				if lAkwas {
					keyword += string(res[i])
				}
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			res = append(res, i+32)
		} else {
			res = append(res, i)
		}
	}
	return string(res)
}

				if keyword == "cap" {
					for i := index-1 ; i > 0 ; i--{
						if res[i] == ' ' {
							break
						}

					}category:formatters go
										arr[i+1] = ""
						arr[i-1] = ""	keyword  = ""
				} else if keyword == "low" {
					for i := index-1 ; i > 0 ; i--{
						if res[i] == ' ' {
							break
						}

					}
					keyword = ""
				}
			}*/
			
			str2 := functions.TextFormated(arr)
			//arr2 := strings.Fields(string(str2))
			fmt.Printf("%v\n", str2)
		}
	}
}
