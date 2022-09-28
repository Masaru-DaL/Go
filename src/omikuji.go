package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	Omikuji_num := rand.Intn(6)
	switch Omikuji_num {
	case 6:
		fmt.Println("出た数字は", Omikuji_num, "! 大吉")
	case 5, 4:
		fmt.Println("出た数字は", Omikuji_num, "! 中吉")
	case 3, 2:
		fmt.Println("出た数字は", Omikuji_num, "! 吉")
	case 1:
		fmt.Println("出た数字は", Omikuji_num, "! 凶")
	}
}
