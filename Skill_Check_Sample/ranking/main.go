package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	ans := 0

	for i := 0; i < 5; i++ {
		sc.Scan()
		var input, _ = strconv.Atoi(sc.Text())
		if i == 0 {
			ans = input
		} else if ans > input {
			ans = input
		}
	}
	fmt.Println(ans)
}
