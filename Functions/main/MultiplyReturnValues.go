package main

import "fmt"

/*
Возврат нескольких значений часто используется для возврата ошибки вместе с результатом (x, err := f())
или логического значения, говорящего об успешном выполнении (x, ok := f()).
 */
func f() (int, int) {
	return 5, 6
}

func withError() (int, string) {
	return 404, "page not found"
}

func withBool()(int, bool)  {
	return 101, true
}

func main() {
	x, y := f()
	fmt.Println("callinf f() -> ", x, y)


	x, err := withError()
	fmt.Println("calling withError() -> ", x, err)

	x, bool := withBool()
	fmt.Println("calling withBool() -> ", x, bool)

}
