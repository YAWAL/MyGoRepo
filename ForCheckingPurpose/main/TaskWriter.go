package main

import "fmt"

func WriteTask()  {
	fmt.Println("Дано натуральне число n. \n" +
		"Чи можна представити його у вигляді суми двох квадратів натуральних чисел? \n" +
		"Якщо можна, то : \n" +
		"a) вказати пару х, у таких натуральних чисел, що n = х^2 + y^2 \n" +
		"б) вказати всі пари х, у таких натуральних чисел, що n = х^2 + y^2, x >= y \n")
}