package main

import (
	"bufio"
	"github.com/matumotoonga/basic/com"
	"github.com/matumotoonga/basic/text"
	"io"
	"os"
	"strconv"
	"strings"
)

const PROMPT = "* "

func do(in io.Reader, out io.Writer) {
	for {
		io.WriteString(out, PROMPT)

		scanner := bufio.NewScanner(in)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		slin := strings.Split(line," ")

		var head = ""
		// var secd = ""

		for i,v := range slin {
			if i == 0 { head = v }
			// if i == 1 { secd = v }
		}
		n, _ := strconv.Atoi(head)

		if n > 0 {
			text.Insert(n,line)

		} else {
			command := strings.ToUpper(head)

			if command == "QUIT" || command == "BYE" {
				break
			} else if command == "BUILD" {
				com.Build()
				break

			} else if command == "CLS" {
				com.Cls()

			} else if command == "CODE" {
				com.Code()
				
			} else if command == "DIR" {
				com.Dir()

			} else if command == "LIST" {
				text.List(out)
			
			} else if command == "LOAD" {

			
			} else if command == "SAVE" {

			
			} else {

			}
			
		}

	}
}
func main() {
	do(os.Stdin, os.Stdout)
}
