package massert

import "fmt"

// Equal assert Equal
func Equal(input, expect interface{}, msgs ...interface{}) {
	sInput := fmt.Sprintf("%v", input)
	sExpect := fmt.Sprintf("%v", expect)
	if sInput == sExpect {
		return
	}
	if len(msgs) != 0 {
		msg := fmt.Sprintf(msgs...)
		fmt.Printf("%s\n", msg)
	}
	fmt.Printf("Error: input: %v , expect: %v\n", input, expect)
}

// True assert True
func True(input interface{}, msgs ...interface{}) {
	Equal(input, true, msgs...)
}

// False assert False
func False(input interface{}, msgs ...interface{}) {
	Equal(input, false, msgs...)
}
