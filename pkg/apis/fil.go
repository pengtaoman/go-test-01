package apis

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(fileName string) {

	file, _ := os.OpenFile(fileName, os.O_RDWR, 0)
	scan := bufio.NewScanner(file)

	for scan.Scan() {
		fmt.Println(">>>> :", scan.Text())
	}
}
