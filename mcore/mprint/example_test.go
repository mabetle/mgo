package mprint

var vs = []string{"a", "b"}

func Example() {
	r := NewStringArrayResult("", vs)
	StdRender(r)
	// Output:
	// 1.a
	// 2.b
}

func ExampleSlice() {
	PrintSlice(vs)
	// Output:
	// 1.a
	// 2.b
}

func ExamplePrint() {
	Print(vs)
	// Output:
	// 1.a
	// 2.b
}
