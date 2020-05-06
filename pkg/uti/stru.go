package uti

import (
	"fmt"
)

type str struct {
	sli []string
	mp  map[string]int
}

func TestStru() {

	var sli []string
	sli = append(sli, "asdfas", "2434", "3q4f3q34")
	str := str{sli, nil}

	fmt.Println("##############%#v", str)
}
