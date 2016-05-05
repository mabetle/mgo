package mjson

import (
	"encoding/json"
)

// Unmarshal try to unmarshal json string.
func Unmarshal(str string, v interface{}) error {
	err := json.Unmarshal([]byte(str), v)
	logger.CheckError(err)
	return err
}
