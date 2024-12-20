package main

import "fmt"

func main() {
	l := NewLocale("en-US")
	fmt.Println("Language:", l.Language())
}
