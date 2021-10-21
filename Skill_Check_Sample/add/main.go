package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanLineInt() int32 {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return int32(i)
}

func main() {
	sc.Split(bufio.ScanWords)
	var a, b = scanLineInt(), scanLineInt()

	var ans int32 = a + b
	fmt.Println(ans)
}
