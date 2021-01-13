package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	//読み取れる入力がある限り受け取り続ける
	if stdin.Scan() {
		for stdin.Scan() {
			text := stdin.Text()
			fmt.Printf("%s", text)
		}
	}
}
