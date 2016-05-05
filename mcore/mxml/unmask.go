package mxml

import (
	"encoding/xml"
)

// Unmarshal try to unmarshal xml string.
func Unmarshal(str string, v interface{}) error {
	err := xml.Unmarshal([]byte(str), v)
	logger.CheckError(err)
	return err
}
