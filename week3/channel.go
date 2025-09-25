package main

/*
import (

	"fmt"

)
*/
func squareIt2(inputChan, outputChan chan int) { // channel chan을 원하고, 해당 채널 내의 데이터는 int형임을 알려줌
	for x := range inputChan {
		outputChan <- x * x
	}
}

func squarerCuber(sqInChan, sqOutChan, cuInChan, cuOutChan, exitChan chan int) {
	var squareX int
	var cubeX int
	for {
		select {
		case squareX = <-sqInChan:
			sqOutChan <- squareX * squareX
		case cubeX = <-cuInChan:
			cuOutChan <- cubeX * cubeX * cubeX
		case <-exitChan:
			return
		}
	}
}

/*
func main() {
	inputChannel := make(chan int)
	outputChannel := make(chan int, 10)
	go squareIt2(inputChannel, outputChannel) // 고루틴으로 실행, 함수 백그라운드에서 실행
	for i := 0; i < 10; i++ {
		inputChannel <- i
	}
	for i := 0; i < 10; i++ { //출력 채널이 비어있으면 반복문이 언제 멈추어야할지 모르기때문에, 일단 입력을 10번 받기에, 출력도 10번으로 수기 조정
		fmt.Println(<-outputChannel)
	}

	close(inputChannel)

	semaphore := make(chan struct{})
	go func() {
		semaphore <- struct{}{}
		fmt.Println("sinalling")

	}()
	<-semaphore
	fmt.Println("exiting")

	// select 예제
	sqInChan := make(chan int, 10)
	cuInChan := make(chan int, 10)
	sqOutChan := make(chan int, 10)
	cuOutChan := make(chan int, 10)
	exitChan := make(chan int)

	go squarerCuber(sqInChan, sqOutChan, cuInChan, cuOutChan, exitChan)
	for i := 0; i < 10; i++ {
		sqInChan <- i
		cuInChan <- i
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("squarer says %d\n", <-sqOutChan)
		fmt.Printf("cuber says %d\n", <-cuOutChan)
	}
	exitChan <- 0

	//5.3 별도의 고루틴에서 함수의 생성과 호출을 동시에 진행하기

	inputChan2 := make(chan int, 10)
	finishChan := make(chan int)
	outputChan2 := make(chan int, 10)
	go func(inputChan, finishChan chan int) {
		for {
			select {
			case x := <-inputChan:
				outputChan2 <- x * x
			case _ = <-finishChan:
				return
			}
		}
	}(inputChan2, finishChan) //생성과 동시에 실행하고 go루틴으로 실행
	for i := 0; i < 10; i++ {
		inputChan2 <- i
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-outputChan2)
	}
	finishChan <- 1

}
*/
