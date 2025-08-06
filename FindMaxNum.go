package main

import "fmt"
/*
Задание:
Последовательность состоит из натуральных чисел и завершается числом 0. 
Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.
*/
func main() {

	var n int
	x, counts := 0, 0, 0

	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if n < x {
			continue
		} else if n > x {
			x = n
			counts = 1
			fmt.Println("Новое максимальноe -", x, ", количество -", counts)
		} else {
			counts++
			fmt.Println("Количество максимальных увеличено: ", counts)
		}
	}
	fmt.Println(counts)

}
