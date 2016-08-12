package mxml

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/mabetle/mgo/mcore"
)

// Marshal returns string of obj xml format.
func Marshal(obj interface{}) string {
	//b, err := xml.Marshal(obj)
	b, err := xml.MarshalIndent(obj, "", "  ")
	// Marshal error
	if logger.CheckError(err) {

	}
	return string(b)
}

// Write to io.Writer
func Write(obj interface{}, w io.Writer) error {
	s := Marshal(obj)
	if s == "" {
		return fmt.Errorf("Parse object to xml error")
	}
	_, err := w.Write([]byte(s))
	return err
}

// WriteFile write marshal xml to file
func WriteFile(obj interface{}, file string) error {
	s := Marshal(obj)
	if s == "" {
		return fmt.Errorf("Parse object to xml error")
	}
	_, err := mcore.WriteFile(file, s)
	return err
}
