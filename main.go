package main

import (
	"fmt"
	"goweek2/week2"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	Harley := week2.Person{"Harley", 21}
	Snowy := week2.Dog{"Snowy", &Harley, 6}
	fmt.Println("2.27 두개의 구조체 인스턴스로 호출하기")
	week2.IncrementAndPrintAge(&Harley)
	week2.IncrementAndPrintAge(&Snowy) //참조가 아닌 값으로 넘기면 오류, 해당 함수를 구현하지 않았으니까

	// 2.28 간단한 DivisionByZero 오류
	fmt.Println("2.28 간단한 DivisionByZero 오류")
	n1, e1 := week2.Divide(1, 1)
	fmt.Println(n1)
	if e1 != nil {
		fmt.Println(e1.Error())

	}
	n2, e2 := week2.Divide(1, 0)
	fmt.Println(n2)
	if e2 != nil {
		fmt.Println(e2.Error())
	}

	//2.29 패닉의 동작: 단일 함수 호출
	fmt.Println("2.29 패닉의 동작: 단일 함수 호출")
	week2.TestRecover()

	//2.30 패닉의 동작: 중첩 함수 호출
	week2.TestRecover2()

	//3.1 외부 API 호출
	fmt.Println("3.1 외부 API 호출")
	body, _ := week2.SearchByID("tt3896198")
	fmt.Println(body.Title)
	body, _ = week2.SearchByName("Game of")
	fmt.Println(body.Title)

	/*
		//3.2 Otiai10 primes 패키지 사용하기
		fmt.Println("3.2 Otiai10 primes 패키지 사용하기")
		args := os.Args[1:]
		if len(args) != 1 {
			fmt.Println("Usage:", os.Args[0], "<number>")
			os.Exit(1)

		}
		number, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		f := primes.Factorize(int64(number))
		fmt.Println("primes:", len(f.Powers()) == 1)
	*/

	//3.4 소수 확인 패키지 사용해서 서버 연결하기
	e := echo.New()
	e.GET("/:number", func(c echo.Context) error {
		nstr := c.Param("number")
		n, err := strconv.Atoi(nstr)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusOK, strconv.FormatBool(week2.IsPrime(n)))

	})
	e.Logger.Fatal(e.Start(":1323"))
}
