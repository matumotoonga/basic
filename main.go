package main

import (
	"bufio"
	"github.com/matumotoonga/basic/com"
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

		} else {
			command := strings.ToUpper(head)

			if command == "QUIT" {
				break
			} else if command == "CLS" {
				com.Cls()
			} else {

			}
			
		}

	}
}
func main() {
	do(os.Stdin, os.Stdout)
}
