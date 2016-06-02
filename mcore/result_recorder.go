package mcore

import (
	"fmt"
	"io"
	"net/http"

	"github.com/revel/revel"
)

type Results struct {
	Errs []error
	Msgs []string
}

func NewResults() *Results {
	return &Results{Errs: []error{}, Msgs: []string{}}
}

// Error merge errors
func (rr *Results) Error() error {
	if rr.IsHasError() {
		sb := NewStringBuffer()
		for _, err := range rr.Errs {
			sb.Appendf("%v\n", err)
		}
		return fmt.Errorf(sb.String())
	}
	return nil
}

// RecordError .
func (rr *Results) RecordError(err error) *Results {
	if err != nil {
		rr.Errs = append(rr.Errs, err)
	}
	return rr
}

// RecordErr .
func (rr *Results) RecordErr(errMsg string, args ...interface{}) *Results {
	if errMsg == "" {
		return rr
	}
	err := fmt.Errorf(errMsg, args...)
	rr.RecordError(err)
	return rr
}

// RecordMsg .
func (rr *Results) RecordMsg(msg string, args ...interface{}) *Results {
	if msg == "" {
		return rr
	}
	rr.Msgs = append(rr.Msgs, fmt.Sprintf(msg, args...))
	return rr
}

// Record .
func (rr *Results) Record(err error, okMsg string, args ...interface{}) *Results {
	if err != nil {
		rr.RecordError(err)
	} else {
		rr.RecordMsg(okMsg, args...)
	}
	return rr
}

// Done record done
func (rr *Results) Done() *Results {
	rr.RecordMsg("Done")
	return rr
}

func (rr *Results) Html() string {
	w := NewStringBuffer()
	rr.HtmlWrite(w)
	return w.String()
}

func (rr *Results) HtmlWrite(out io.Writer) {
	if len(rr.Errs) == 0 && len(rr.Msgs) == 0 {
		rr.Done()
	}
	// has error
	if len(rr.Errs) > 0 {
		out.Write([]byte(`<div class="error"><ul>`))
		for _, err := range rr.Errs {
			line := fmt.Sprintf(`<li>%v</li>`, err)
			out.Write([]byte(line))
		}
		out.Write([]byte(`</ul></div>`))
	}

	// has message
	if len(rr.Msgs) > 0 {
		out.Write([]byte(`<div class="info"><ul>`))
		for _, msg := range rr.Msgs {
			line := fmt.Sprintf(`<li>%v</li>`, msg)
			out.Write([]byte(line))
		}
		out.Write([]byte(`</ul></div>`))
	}
}

func (rr *Results) IsHasError() bool {
	return len(rr.Errs) > 0
}

func (rr *Results) IsHasMsg() bool {
	return len(rr.Msgs) > 0
}

// PrintErrors prints errors
func (rr *Results) PrintErrors() {
	if !rr.IsHasError() {
		fmt.Println("No errors")
		return
	}
	fmt.Printf("Errors:\n")
	for i, err := range rr.Errs {
		fmt.Printf("\t %d: %v\n", i+1, err)
	}
}

// PrintMsgs prints messages.
func (rr *Results) PrintMsgs() {
	if !rr.IsHasMsg() {
		fmt.Printf("No messages.\n")
		return
	}
	fmt.Printf("Messages:\n")
	for i, v := range rr.Msgs {
		fmt.Printf("\t %d: %v\n", i+1, v)
	}
}

// Print prints errors and messages.
func (rr *Results) Print() {
	rr.PrintErrors()
	rr.PrintMsgs()
}

// Apply implements revel.Result
func (rr *Results) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusOK, "text/html")
	rr.HtmlWrite(resp.Out)
}
