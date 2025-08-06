package main

import "fmt"

/*
Задание:
Найдите первое число от 1 до n включительно,
 кратное c, но НЕ кратное d.
*/
func main() {

	var n, c, d int
	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)

	for i := 0; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Print("Первое число, кратное ", c, ", но не ", d, " : ", i)
			break
		} else {
			continue
		}
	}

}
