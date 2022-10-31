package main

import "fmt"

func forLoop(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("타잔은 %d원짜리 팬티를 입고, %d원짜리 칼을 차고 노래를 한다.\n", i*10, (i+1)*10)
	}
}

func main() {
	forLoop(5)
}
