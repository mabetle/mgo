package mcore

import (
)

func ExampleNewRandom() {
	for i := 0; i < 100; i++ {
		println(i, ":", NewRandom())
	}

	println("======================================")

	for i := 0; i < 100; i++ {
		println(i, ":", NewRangeRandom(1000000))
	}

}
