package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)
	milis := nanos / 100000
	fmt.Println("Sec:", secs)
	fmt.Println("Milis:", milis)
	fmt.Println("Nanos:", nanos)
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
