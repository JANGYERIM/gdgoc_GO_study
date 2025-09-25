package main

import (
	"fmt"
	//"time"
)

func squareIt(x int) {
	fmt.Println("5.1 스레드 프로그램")
	fmt.Println(x * x)
}

/*
func main() {
	squareIt(2)
	go squareIt(3)
	time.Sleep(1 * time.Millisecond) //main 함수가 종료되지 않도록 기다림(고 함수가 반환되기를 기다려야하기 때문에)
	// 실제로 작업 완료를 기다리기 위해서는 '채널'이라는 기능을 이용해야 함
}
*/
