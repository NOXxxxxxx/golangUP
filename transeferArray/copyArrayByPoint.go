package transeferArray

import "fmt"

func Modify2(a *[5]int) {
	a[1]=66
	fmt.Println(*a)
}
