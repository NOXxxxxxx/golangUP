package main

import (
	"fmt"
	"golangUP/demo01"
	"golangUP/fibo"
	"golangUP/fileOperation"
	"golangUP/stringPractise"
	"golangUP/structDemo"
	"golangUP/transeferArray"
	"time"
)

func main() {
	var start int64 = time.Now().UnixNano()
	fibo.Fbn(10)
	var end int64 = time.Now().UnixNano()
	fmt.Println(end - start)

	demo01.Test03()
	newa := [5]int{2, 3, 4, 5, 6}
	transeferArray.Modify(newa)

	newa2 := [5]int{5, 3, 2, 2, 1}
	transeferArray.Modify2(&newa2)
	fmt.Println(newa2)
	a, b := 10, 20
	demo01.Swap(a, b)
	fmt.Println("a:", a, "b:", b)
	demo01.SwapPoint(&a, &b)
	fmt.Println("a:", a, "b:", b)

	demo01.SliceTest()

	str := "i love go go language! wow wow wow"
	mRet := demo01.WordCountFunc(str)

	for k, v := range mRet {
		fmt.Printf("%q:%d\n", k, v)
	}

	var man structDemo.Person = structDemo.Person{Name: "andy", Sex: 'm', Age: 18}
	fmt.Println(man)
	var man2 *structDemo.Person = &structDemo.Person{Name: "timo", Sex: 'f', Age: 199}
	fmt.Println(*man2)
	fmt.Printf("%p\n", man2)
	fmt.Print("man2.name:", &man2.Name)

	fmt.Println("********************************")

	var stu structDemo.Student
	structDemo.InitFunc(&stu)
	fmt.Println(stu)

	fmt.Println("**********返回值初始化******************")

	fmt.Println("********************string操作****************")

	stringPractise.StringPractise()

	fmt.Println("****************文件操作*********************")

	fileOperation.FilePractice()
}
