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
		os.Exit(21)
	} else {
		in := args[0]
		out := args[1]

		if !strings.HasSuffix(in, ".txt") || !strings.HasSuffix(out, ".txt") {
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
			for i, line := range lines {
				line = functions.HandleFlag(line)
				line = functions.TextFormated(strings.Fields(line))
				line = functions.TextFormated(strings.Fields(line))
				line = functions.HandleQuote(line)
				line = functions.HandleVowel(line)
				finalres += line 
				if i != len(lines)-1{
					finalres+="\n"
				}
			}

			err = os.WriteFile(out, []byte(finalres), 0o644)
			if err != nil {
				fmt.Println("Error writing file:", err)
				return
			}
			//fmt.Println("bravoo!!!")
		}
	}
}
