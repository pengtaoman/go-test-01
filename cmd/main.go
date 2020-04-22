package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func createPointer() *int {
	var rr = 2344
	return &rr
}

func printPointer(po *int) {
	fmt.Printf("VVVVVVVVV : %d", *po)
}

func main() {
	a := "ali#dsd##"
	fmt.Println("###########", a)
	rep := strings.NewReplacer("#", "--")
	ss := rep.Replace(a)
	fmt.Println(ss)

	time1 := time.Now()
	println(time1.Year())
	println(time1.Month())

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 3; i++ {
		input, _ := reader.ReadString('\n')
		fmt.Println("Input is ::::", input)
	}

	s001 := "        a a        sd  dferef  svdaeev  -- dfder        "
	s001 = strings.TrimSpace(s001)
	println(s001)

	a11 := 23
	b22 := 3.2
	println(float64(a11) * b22)
	log.Printf("------------------------------------------------------")
	fi, err := os.Stat("/Users/pengtao/my-work/tmp/go/cmd/main.go")
	if err != nil {
		println("EEEEEEEEEEEEEEEE")
		println(err)
	}
	fmt.Println("################", fi.Size())
	println("EEEEEEEEEEEEEEEE===========================")
	a33 := "544"
	a44, _ := strconv.ParseFloat(a33, 64)

	// if err != nil {
	// 	println("EEEEEEEEEEEEEEEE")
	// 	println(err)
	// }

	fmt.Println("################", a44)

	time2 := time.Now().Unix()
	fmt.Println("################", time2)

	for iii := 0; iii < 10; iii++ {

		if iii > 4 {

			continue
		}
		println("!!!!!!!!!!!!!" + strconv.Itoa(iii))
	}

	fmt.Printf("KKK :%#v", "\n")
	fmt.Printf("KKK :%v", "\n")
	fmt.Printf("KKK ===============")
	fmt.Printf("KKK :%v", "\n")
	fmt.Printf("KKK :%v", "\n")
	fmt.Printf("KKK :%v", "\n")

	amount := 7
	println(&amount)
	var myp int
	myp = 121
	var mypp *int
	mypp = &myp
	myp = 1212323
	fmt.Println(reflect.TypeOf(mypp))
	fmt.Println(*mypp)

	pppoo := createPointer()
	fmt.Printf("HHHHHHH: %d", *pppoo)
	println()
	printPointer(&myp)

	diffFmt()

	println()
	println()
	println()
}

func diffFmt() {
	testStr := []byte("test str")
	println()
	fmt.Println(testStr, "hello world") //[116 101 115 116 32 115 116 114] hello world
	fmt.Printf("111111%s", testStr)     //test str
	fmt.Println()
	fmt.Sprintf("22222%s", testStr) //空，无IO输出
	fmt.Println()
	printStr := fmt.Sprintf("33333%s", testStr)
	fmt.Println(printStr) //test str
}
