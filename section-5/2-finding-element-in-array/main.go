package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"Sandy","Provo","St. George","Salt Lake City","Draper","South Jordan","Murray"}

	for i,v := range str{
		if v == "Sandy" {
			fmt.Println(i)
		}
	}

	sortedList := sort.StringSlice(str)
	sort.Sort(sort.Reverse())
	fmt.Println(sortedList)

	sort.in
	index := sortedList.Search("Sandy")
	fmt.Println(index)

}
