package uti

var StrConst string = "23asdfasdfasdfasdfadf"

type met struct {
	name   string
	age    int
	gender bool
}

func GetMet() met {

	return met{"asdf", 23, false}
}

func change(mm *met) {

}

func TestChange() {
	mm1 := GetMet()
	change(mm1)
}
