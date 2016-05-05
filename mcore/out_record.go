package mcore

import (
	"fmt"
	"io"
	"os"
)

type OutRecorder struct {
	Writer io.Writer
	Type   string
}

func NewOutRecorder(w io.Writer, typ string) *OutRecorder {
	return &OutRecorder{Writer: w, Type: typ}
}

func NewHtmlOutRecorder(w io.Writer) *OutRecorder {
	return NewOutRecorder(w, "html")
}

func NewStdOutRecorder() *OutRecorder {
	return NewOutRecorder(os.Stdout, "")
}

// RecordError
func (r *OutRecorder) RecordError(err error) {
	if err == nil {
		return
	}
	r.Writef("error", fmt.Sprintf("Error: %v", err))
}

// RecordMsg
func (r *OutRecorder) RecordMsg(msg string, args ...interface{}) {
	r.Writef("info", msg, args...)
}

func (r *OutRecorder) Record(err error, okMsg string, args ...interface{}) {
	if err != nil {
		r.RecordError(err)
	} else {
		r.RecordMsg(okMsg, args...)
	}
}

// Writef
func (r *OutRecorder) Writef(cssClass string, msg string, args ...interface{}) {
	outMsg := fmt.Sprintf(msg, args...)
	if r.Type == "html" {
		// html
		outMsg = fmt.Sprintf(`<p class="%s">%s</p>`, cssClass, outMsg)
	} else {
		// text, add new line
		outMsg = fmt.Sprintf("[%s] %s\n", cssClass, outMsg)
	}
	r.Writer.Write([]byte(outMsg))
}
