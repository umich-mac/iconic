package main

import "fmt"

func ExampleHashFile() {
	fmt.Println(HashFile("test.png"))
	// Output: 2f5bbd5c5147380eecd94daa0cf0681c30ca43803f10868a2b9a2aedd26b0f21
}
