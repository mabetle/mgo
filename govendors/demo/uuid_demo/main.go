package main

import (
	"github.com/mabetle/govendors/uuid"
)

func main() {
	for i := 0; i < 10; i++ {
		println(uuid.New())
	}
}
