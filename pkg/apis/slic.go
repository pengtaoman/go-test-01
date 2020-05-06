package apis

import (
	"fmt"
	"reflect"
	"strconv"
)

func Slic() {

	sl := make([]string, 8)
	println(len(sl))

	fmt.Println(reflect.TypeOf(sl))

	// for i := 0; i < 20; i++ {
	// 	sl[i] = "asdfasdf" + string(i)
	// }

}

func Dyslic() {

	var sl []string
	for i := 0; i < 20; i++ {
		// sl[i] = "asdfasdf" + string(i)
		sl = append(sl, "asdfasdf----"+strconv.Itoa(i))
	}

	for _, value := range sl {
		println("!!!!!!!!!!!!! ::: " + value)
	}
}

func SliOver() {
	ss := [5]string{"aa", "bb", "cc", "dd", "ee"}

	sc1 := ss[1:3]

	ss[2] = "asdfae3e"

	fmt.Println(sc1)
}

func SliZero() {
	var ss []int
	fmt.Println("####", ss)
}
