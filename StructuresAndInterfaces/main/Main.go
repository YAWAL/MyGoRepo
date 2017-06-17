package main

import (
	"fmt"
	"math"
)

//локальная переменная типа Circle, чьи поля по умолчанию будут равны нулю
// (0 для int, 0.0 для float, "" для string, nil для указателей, …)
type Circle struct {
	x float64
	y float64
	r float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r*c.r
}

func main()  {



	/*
	Это выделит память для всех полей, присвоит каждому из них нулевое значение и вернет указатель (*Circle).
	 Часто, при создании структуры мы хотим присвоить полям структуры какие-нибудь значения.
	  Существует два способа сделать это. Первый способ:
	 */
	newCircle := new(Circle)
	newCircle.x = 1.54
	newCircle.y = 2.456
	newCircle.r = 5.46


	c := Circle{0, 0, 5}
	fmt.Println(circleArea(c))

	circle1 := Circle{x: 0, y: 0, r: 5}
	fmt.Println(circle1.r)
	fmt.Println(newCircle)


}
