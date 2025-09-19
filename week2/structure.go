package week2

import "fmt"

type User struct {
	Name string
	Age  int
} //구조체는 참조타입이 아닌 값 타입, 구조체의 정의

func nameAndAge(uid int) User {
	switch uid {
	case 0:
		return User{"Baheer Kamal", 24}
	case 1:
		return User{"Tanmay Bakshi", 16}
	default:
		return User{"", -1}
	}
}

func incrementAge(user User) {
	user.Age++ //값을 복사하기에 원본에는 영향이 x
	fmt.Println("2.20 인자전달: ", user.Age)
}

func incrementAge2(user *User) {
	user.Age++ //참조로 받았기에 원본에 영향 o
	fmt.Println("2.21 포인터 전달: ", user.Age)
}

/*
c++에서는 포인터지만 고에서는 포인터가 아님, 단지 참조
c에서는 포인터를 역참조하는 경우가 아니라면 -> 연산자를 사용해야하지만,
go는 포인터가 아니므로 그러한 제한이 없음
*/

// 구조체에 함수 할당
func (user *User) incrementAge3() {
	user.Age++
	fmt.Println("2.23 포인터 리시버를 이용한 구조체 변화: ", user.Age)
} //인자명과 타입을 함수명 앞으로 빼기, 구조체 위에 인스턴스 메서드를 만듬

/*
func main() {
	//2.19
	user := nameAndAge(1)
	fmt.Println("2.19. User age:")
	fmt.Println(user.Age)
	//2.20
	kathy := User{"Kathy", 19}
	incrementAge(kathy)
	fmt.Println(kathy.Age)
	//2.21
	lathy := User{"lathy", 19}
	incrementAge2(&lathy)
	fmt.Println(lathy.Age)
	//2.23
	jathy := User{"jathy", 19}
	jathy.incrementAge3() //구조체에 함수를 할당한 경우
	fmt.Println(jathy.Age)

}
*/
