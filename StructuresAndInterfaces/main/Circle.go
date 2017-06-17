package main

import (
	"math"
	"fmt"
)

type Circle1 struct {
	x float64
	y float64
	r float64
}

func circleArea1(c *Circle1) float64 {
	return math.Pi * c.r*c.r
}

/*
Между ключевым словом func и именем функции мы добавили «получателя».
 Получатель похож на параметр — у него есть имя и тип, но объявление функции таким способом позволяет нам вызывать
  функцию с помощью оператора .
 */
func (c *Circle1) area() float64 {
	return math.Pi * c.r*c.r
}

func main()  {

	c := Circle1{1, 2, 3}
	fmt.Println(c.area())

}