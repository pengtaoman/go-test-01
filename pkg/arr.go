package apis

func StrArr() {
	var str [4]string
	str[0] = "ARRRRRRR"

	for i := 0; i < len(str); i++ {
		println(">>>>>>>>>>>>>>>>>>>>>>" + str[i])
	}

	for _, value := range str {
		println("=================" + value)
	}
}
