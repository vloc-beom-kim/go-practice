package main

import (
	"fmt"
)

type job struct {
	Name   string
	Result string
}

func worker(j chan job) {
	for {
		myjob := <-j
		// if myjob.Name == "close" {
		// 	time.Sleep(time.Second * 1)
		// 	myjob.Result = "종료"
		// } else {
		// }
		fmt.Println("작업 수행:", myjob.Name)
		myjob.Result = "완료"
		j <- myjob
	}
}

func main() {

	done := make(chan job)
	go worker(done)
	// time.Sleep(time.Second * 1)
	for _, i := range []string{"job 1", "job 2"} {
		done <- job{Name: i}
		result := <-done
		fmt.Println("Job: ", result.Name, result.Result)
	}

	// done <- job{Name: "close"}
	// result := <-done
	// fmt.Println("결과: ", result)
	// <-done

}
