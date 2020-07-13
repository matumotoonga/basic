package text

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

const NEWLINE = "\n"

var filename = ""
var TEXT map[string]string

func Insert(n int, s string) {
	if TEXT == nil {
		TEXT = make(map[string]string)
	}
	m := fmt.Sprintf("L%05d", n)
	TEXT[m] = s
}

func List(out io.Writer) {
	ks := []string{}
	for k, _ := range TEXT {
		ks = append(ks, k)
	}
	sort.Strings(ks)

	for _, k := range ks {
		io.WriteString(out, TEXT[k]+NEWLINE)
	}
}

func Load(n string) {
	for i, v := range os.Args {
		if i == 1 {
			filename = v
		}
	}
	if filename == "" {
		if n == "" {
			return
		}
		filename = n
	}
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _ , err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		sbuf := string(line)
		head := strings.Split(sbuf," ")[0]
		n, _ := strconv.Atoi(head)

		if n > 0 {
			Insert(n,sbuf)
		}
	}

}

func Save(n string) {

}
