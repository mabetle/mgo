package mjson

import (
	"encoding/json"
)

// Marshal returns string of obj json format.
func Marshal(obj interface{}) string {
	//b, err := json.Marshal(obj)
	b, err := json.MarshalIndent(obj, "", "  ")
	// Marshal error
	if logger.CheckError(err) {

	}
	return string(b)
}
