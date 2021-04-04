package stringPractise

import (
	"fmt"
	"strings"
)

func StringPractise() {
	str := "I love work and go go go language!"
	ret := strings.Split(str, "g")
	for _, s := range ret {
		fmt.Println(s)
	}
	flg_end := strings.HasSuffix("test.mp3", "mp3")
	fmt.Println(flg_end)

	flg_start := strings.HasPrefix("li@11", "@")
	fmt.Println(flg_start)

}
