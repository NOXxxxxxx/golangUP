package fileOperation

import (
	"fmt"
	"os"
)

func FilePractice() {
	f, err := os.Create("./fileOperation/createFromGO")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	fmt.Println("Create Success")

	_, err1 := f.WriteString("goooooooooooooooooooooooooooO")
	if err1 != nil {
		fmt.Println("Writing Failed", err1)
		return
	}

	content, err2 := os.ReadFile("./fileOperation/createFromGO")
	if err2 != nil {
		fmt.Println("Writing Failed!", err2)
		return
	}

	fmt.Printf("%s\n", content)

}
