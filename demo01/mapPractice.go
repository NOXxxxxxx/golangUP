package demo01

import "strings"

func WordCountFunc(str string) map[string]int {
	s := strings.Fields(str) //将字符串拆分成字符串切片
	m := make(map[string]int)

	//遍历拆分后的字符串切片
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok { //当ok==true，说明s[i] 作为key存在
			m[s[i]] = m[s[i]] + 1 //m[s[i]] +1

		} else {
			m[s[i]] = 1
		}
	}

	return m
}
