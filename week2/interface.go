package week2

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Dog struct {
	Name  string
	Owner *Person
	Age   int
}

func (p *Person) incrementAge() {
	p.Age++
}

func (p *Person) getAge() int {
	return p.Age
} // 인터페이스의 함수가 하나라도 포인터 리시버가 필요하다면 함수 모두 포인터 리시버로 정의해야함

func (d *Dog) incrementAge() {
	d.Age++
}

func (d *Dog) getAge() int {
	return d.Age
}

type Living interface {
	incrementAge()
	getAge() int
} // 인터페이스(여러 함수를 클래스처럼 표현, 오버라이딩 하듯이 정의하고 사용)
// 단지 규약일뿐 각각의 함수는 정의해주어야함(이 두가지 함수를 구현하는 모든 타입을 living 타입으로 간주)

func IncrementAndPrintAge(being Living) {
	being.incrementAge()
	fmt.Println(being.getAge())
}
