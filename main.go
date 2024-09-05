package main

import (
	"fmt"
	"io"
	"os"
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
			lAkwas := false
			keyword := ""
			for i := 0; i < len(res); i++ {
				if res[i] == '(' {
					lAkwas = true
					continue
				}
				if res[i] == ')' {
					lAkwas = false
					continue
				}
				if lAkwas {
					keyword += string(res[i])
				}
			}
			fmt.Println(string(keyword))
		}
	}
}
