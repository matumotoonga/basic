package text

import (
	"fmt"
	"io"
	"sort"
)
const NEWLINE = "\n"

var TEXT map[string]string

func Insert(n int, s string) {
	if TEXT == nil {
		TEXT = make(map[string]string)
	}
	m := fmt.Sprintf("L%05d", n)
	TEXT[m] =  s 
}

func List(out io.Writer) {
	ks := []string{}
	for k, _ := range TEXT {
		ks = append(ks, k)
	}
	sort.Strings(ks)

	for _, k := range ks {
		io.WriteString(out,TEXT[k] + NEWLINE)
	}


}