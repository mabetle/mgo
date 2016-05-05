package mxml

import (
	"encoding/xml"
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
