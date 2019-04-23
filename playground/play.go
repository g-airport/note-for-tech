package main



func main() {
	//s1 := []string{"A", "B", "C"}
	//fmt.Printf("before foo function, s1 is \t%v\n", s1)
	//s2 := foo(s1)
	//fmt.Printf("after foo function, s1 is \t%v\n", s1)
	//fmt.Printf("after foo function, s2 is \t%v\n", s2)
	//fmt.Println(f())
	//
	//fmt.Println(s([]int{1,2}))
}
func foo(s []string) []string {
	//s[2] = "new"
	s = append(s, "New")
	return s
}

func f() int {
	var res int
	defer func() {
		res++
	}() //0

	return res
}
