package mcore

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
)

// StringBuffer ideas from java StringBuffer.
type StringBuffer struct {
	*bytes.Buffer
}

// NewStringBuffer create no args.
func NewStringBuffer(v ...interface{}) (sb StringBuffer) {
	buf := bytes.NewBufferString(fmt.Sprint(v...))
	sb.Buffer = buf
	return
}

// Append append
func (sb StringBuffer) Append(args ...interface{}) {
	fmt.Fprint(sb.Buffer, args...)
}

// AppendByteArray append
func (sb StringBuffer) AppendByteArray(bs []byte) {
	sb.Append(string(bs))
}

// AppendLines append lines
func (sb StringBuffer) AppendLines(lines []string) {
	for _, line := range lines {
		sb.AppendLine(line)
	}
}

// AppendArrayCSVLine append "[a,b,c] as a,b,c\n"
func (sb StringBuffer) AppendArrayCSVLine(values []string) {
	for index, v := range values {
		sb.Append(v)
		if index < len(values)-1 {
			sb.Append(",")
		}
	}
	sb.Append("\n")
}

// AppendLine append line
func (sb StringBuffer) AppendLine(args ...interface{}) {
	sb.Append(args...)
	sb.Append("\n")
}

// AppendEachLine append line by line
func (sb StringBuffer) AppendEachLine(lines ...interface{}) {
	for _, line := range lines {
		sb.AppendLine(line)
	}
}

// AppendLinef append line with format
func (sb StringBuffer) AppendLinef(format string, args ...interface{}) {
	sb.Appendf(format, args...)
	sb.Append("\n")
}

// Appendf appends with format
func (sb StringBuffer) Appendf(format string, args ...interface{}) {
	fmt.Fprintf(sb.Buffer, format, args...)
}

// String return string
func (sb StringBuffer) String() string {
	return fmt.Sprint(sb.Buffer)
}

// WriteTo write
func (sb StringBuffer) WriteTo(dest io.Writer) (int64, error) {
	return sb.Buffer.WriteTo(dest)
}

// Print StringBuffer content to stdout.
func (sb StringBuffer) Print() {
	sb.WriteTo(os.Stdout)
}

// WriteToFile write to file
func (sb StringBuffer) WriteToFile(file string) (int64, error) {
	dest, err := NewFileWriter(file)
	if err != nil {
		return 0, err
	}
	return sb.WriteTo(dest)
}

// Clear clear value
func (sb StringBuffer) Clear() {
	sb.Buffer = bytes.NewBufferString("")
}

// HTML returns html template.HTML
func (sb StringBuffer) HTML() template.HTML {
	return template.HTML(sb.String())
}
