package demo01

import (
	"fmt"
	"strconv"
)

func Test03(){
	str := " "
	for i:=0;i<10000;i++{
		str += "hello " + strconv.Itoa(i)
		//fmt.Println(str)
	}
	arraylist:=[5]int{1,2,3,3,4}
	//var arraylist [5]int
	//arraylist = [5]int{1,2,3,4,5}
	for i:=0;i<len(arraylist);i++{
		fmt.Println(i)
	}
}


