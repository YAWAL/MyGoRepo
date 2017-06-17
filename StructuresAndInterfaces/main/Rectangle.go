package main

import "fmt"

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func dst(x1, y1, x2, y2  float64)  float64{

	return x1+x2+y1+y2
}

func (r *Rectangle) area() float64 {
	l := dst(r.x1, r.y1, r.x2, r.y2)
	w := dst(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main()  {
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}
