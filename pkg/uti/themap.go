package uti

import (
	"fmt"
)

func TestMap() {

	mm := map[string]int{"aa": -1, "bb": 34354}
	key, ok := mm["cc"]
	fmt.Println(key, ok)

	//delete(mm, "bb")
	key, ok = mm["bb"]
	fmt.Println(key, ok)

}
