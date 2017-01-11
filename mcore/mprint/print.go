package mprint

import (
	"fmt"
	"io"
	"os"
)

type Result interface {
	Apply(out io.Writer)
}

type StringArrayResult struct {
	head string
	rows []string
}

func NewStringArrayResult(head string, rows []string) Result {
	return &StringArrayResult{head: head, rows: rows}
}

func (c StringArrayResult) Apply(out io.Writer) {
	// process head
	if c.head != "" {
		fmt.Fprintf(out, "=== %s ===", c.head)
	}
	// no rows
	if len(c.rows) == 0 {
		fmt.Fprintln(out, "No result.")
		return
	}
	// do print rows
	for i, row := range c.rows {
		fmt.Fprintf(out, "%d.%s\n", i+1, row)
	}
}

func Render(out io.Writer, r Result) {
	r.Apply(out)
}

func StdRender(r Result) {
	r.Apply(os.Stdout)
}
