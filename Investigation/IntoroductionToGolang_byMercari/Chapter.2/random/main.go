package main

import(
	"math/rand"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(10)
	println(s)

	now := time.Now()
	println(now)
}

/* 1~10の間でランダムな数値が出力 */
