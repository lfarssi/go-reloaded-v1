package main

import (
	"fmt"
	"go-reloaded/func"
	"os"
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
			finalres := ""
			lines := strings.Split(string(res), "\n")
			for _, line := range lines {
				line = functions.HandleVowel(line)
				line = functions.HandleQuote(line)
				line = functions.HandleFlag(line)
				line = functions.TextFormated(strings.Fields(line))
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
