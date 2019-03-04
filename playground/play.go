package main

import "fmt"

func main() {
	s1 := []string{"A", "B", "C"}
	fmt.Printf("before foo function, s1 is \t%v\n", s1)
	s2 := foo(s1)
	fmt.Printf("after foo function, s1 is \t%v\n", s1)
	fmt.Printf("after foo function, s2 is \t%v", s2)
}
func foo(s []string) []string {
	//s[2] = "new"
	s = append(s, "New")
	return  s
}
