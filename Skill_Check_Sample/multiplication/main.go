package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Lineinput() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	var a, b int

	a, _ = strconv.Atoi(Lineinput())
	b, _ = strconv.Atoi(Lineinput())
	// fmt.Println(a)
	ans := a * b
	fmt.Println(ans)
}
