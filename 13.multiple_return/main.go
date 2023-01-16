package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("분모로 0을 사용할 수 없습니다.")
	}
	return a / b, nil
}
func main() {
	res, err := divide(4, 0)
	if err != nil {
		fmt.Println("ERROR :", err.Error())
	} else {
		fmt.Println("4 / 2 =", res)
	}
}
