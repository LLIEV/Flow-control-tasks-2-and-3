/* Golang BEINNER.ТЕОРИЯ 
Тема: управление потоком
Задача № 2 (программa кторая выводит числа от 1 до 100, которые делятся на 3)
Задача № 3 (программа которая выводит числа от 1 до 100. Но для кратных 
трём нужно вывети «Fizz» ВМЕСТО числа, для кратных пяти - «Buzz», а
для кратых как трём, так и пяти — «FizzBuzz»).*/

package main

import "fmt"

/*unc main() {
		i:= 1
		i++
		for i <= 100 {
		if i%3 == 0 {
		fmt.Println(i)
		 
}

func main() {
	for i := 1; i <= 100; i++ { 
		if i % 3 == 0 && i % 5 == 0 { 
			fmt.Println("FizzBuzz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz") 
		} else if i % 3 == 0 {
			fmt.Println("Fizz") 
		} else {
			fmt.Println(i) 
		}
	}
}
*/
func main() {
	for i := 1; i <= 100; i++ { 
		switch {
			case i % 3 == 0 && i % 5 == 0: fmt.Println(i, "FizzBuzz")
			case i % 3 == 0: fmt.Println(i, "Fizz")
			case i % 5 == 0: fmt.Println(i, "Buzz")
			default: fmt.Println(i) 
		}
	}
}