package mtml

import (
	"bytes"
	"html/template"
)

// MergeText merge text and data
func MergeText(text string, data interface{}) (string, error) {
	t := template.New("merge")
	t.Parse(text)
	out := bytes.NewBufferString("")
	if err := t.Execute(out, data); err != nil {
		return "", err
	}
	return out.String(), nil
}

// MergeToHtml merge text and data, returns template.HTML
func MergeToHtml(text string, data interface{}) template.HTML {
	txt, err := MergeText(text, data)
	logger.CheckError(err)
	return template.HTML(txt)
}

// MergeFile merge files and data to string
func MergeFile(data interface{}, file ...string) (string, error) {
	t, err := template.ParseFiles(file...)
	if err != nil {
		return "", err
	}
	out := bytes.NewBufferString("")
	if err := t.Execute(out, data); err != nil {
		return "", err
	}
	return out.String(), nil
}
