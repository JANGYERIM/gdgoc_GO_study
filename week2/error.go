package week2

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
} // error 인터페이스 정의

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
} // 인자가 없는 함수 error는 사람이 읽을 수 있는 문자열로 표현된 오류를 반환

var DivisionByZero = errors.New("division by zero") // 오류 변수 정의

func Divide(number, d float32) (float32, error) {
	if d == 0 {
		return 0, DivisionByZero
	}
	return number / d, nil
}

// 2.29 패닉의 동작: 단일 함수 호출

var SampleError = errors.New("This is a tese error")

func TestRecover() {
	defer func() {
		if recover() != nil {
			fmt.Println("got an error!")
		} else {
			fmt.Println("no error")
		}
	}()
	panic(SampleError)
	fmt.Println("Hello!")
}

//2.30 패닉의 동작: 중첩 함수 호출

func TestPanic() {
	panic(SampleError)
	fmt.Println("Hello from testPanic")

}

func TestRecover2() {
	defer func() {
		if recover() != nil {
			fmt.Println("got an error!")
		} else {
			fmt.Println("no error")
		}
	}()
	TestPanic()
	fmt.Println("Hello from testRecover!")
}
