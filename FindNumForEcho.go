package main

import "fmt"

/*
Задание:
Напишите программу, которая считывает целые числа с консоли по одному числу в строке.
Для каждого введённого числа проверить:
-если число меньше 10, то пропускаем это число;
-если число больше 100, то прекращаем считывать числа;
-в остальных случаях вывести это число обратно на консоль в отдельной строке.
*/

func main() {

	var num int
	var answer string
	fmt.Println("Каким способом решить задачу? Если switch, то введите YES, иначе NO")
	fmt.Scan(&answer)

	switch answer {

	case "NO":
		for fmt.Scan(&num); ; fmt.Scan(&num) {
			if num < 10 {
				continue
			} else if num > 100 {
				break
			} else {
				fmt.Println(num)
			}
		}

	case "YES":
		for fmt.Scan(&num); num <= 100; fmt.Scan(&num) {
			switch {
			case num < 10:
				continue
			default:
				fmt.Println(num)
			}
		}

	default:
		fmt.Println("Пожалуйста, введите либо YES, либо NO")
	}
}
