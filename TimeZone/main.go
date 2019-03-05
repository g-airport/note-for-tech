package main

import (
	"fmt"
	"time"
)

func main() {
	test1()
	test2()
	test3()
}

func timeIntro() {
	l := "Asia/Hong_Kong"
	time.LoadLocation(l)

	offset := 0
	time.FixedZone("UTC-8", offset)

	location := time.UTC
	time.Now().In(location)

	var time1 time.Time
	var time2 time.Time
	var time1Zone *time.Location

	// via time.Date -> time struct instance
	time1 = time.Date(time2.Year(), time2.Month(),
		0, 0, 0, 0, 0, time1Zone)
	_ = time1
}

func Daily(t time.Time) time.Time {
	// zero time in golang not 0 year 0 day
	// use 1
	return time.Date(1, 1, t.Day(), t.Hour(),
		0, 0, 0, t.Location())
}

func test1() {
	location, _ := time.LoadLocation("Asia/Hong_Kong")

	t := time.Date(1904, 10, 29, 11, 59, 59, 59, location)

	fmt.Println(t) //1904-10-29 11:59:59.000000059 +0736 LMT
}

func test2() {
	location, _ := time.LoadLocation("Asia/Hong_Kong")

	t := time.Date(1904, 10, 30, 0, 0, 0, 0, location)

	fmt.Println(t) //1904-10-30 00:00:00 +0736 LMT
}

func test3() {
	location, _ := time.LoadLocation("Asia/Hong_Kong")

	t := time.Date(1942, 10, 30, 0, 0, 0, 0, location)

	fmt.Println(t) //日本标准时间 1942-10-30 00:00:00 +0900 JST

}
