package structDemo

type Student struct {
	Name      string
	Age       int
	Flg       bool
	Intereset []string
}

//通过函数参数初始化结构体
func InitFunc(p *Student) {
	p.Name = "Nami"
	p.Age = 18
	p.Flg = true
	p.Intereset = append(p.Intereset, "shopping")
	p.Intereset = append(p.Intereset, "sleeping")
}

func InitStruct() *Student {
	p := new(Student)
	p.Name = "Siwang"
	p.Age = 19
	p.Flg = true
	p.Intereset = append(p.Intereset, "Sleeping")
	p.Intereset = append(p.Intereset, "making")
	return p //不能返回局部变量的地址值，只能返回值
}
