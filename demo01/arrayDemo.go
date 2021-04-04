package demo01

import "fmt"

func SliceTest() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	s := arr[1:3:5]
	fmt.Println("s=", s)
	fmt.Println("len(s)=", len(s))
	fmt.Println("cap(s)=", cap(s))
}
