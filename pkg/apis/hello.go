package apis

import (
	"fmt"
	"strconv"
)

func SayHello() {

	println("!!!!!!!!!!!!!!!!!!!!! SAY HELLO !!!!!!!!!!!!!!!!!!!!!!!!!!")

	q := strconv.Quote("Hello, 世界")

	fmt.Println("!!!!!!!!!!!!!!!!", q)

}
