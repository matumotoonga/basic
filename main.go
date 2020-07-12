package main

import (
	"bufio"
	"io"
	"os"
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

	}
}
func main() {
	do(os.Stdin, os.Stdout)
}
