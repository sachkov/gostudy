package main

import (
	"fmt"
	"os"
	"flag"

	"gostudy/muu"
)

func main() {
	var memo string
	var try string
	var k int
	var b int
	var fn = flag.String("s", "", "First number")

	flag.Parse()

	mu := new(muu.Muu)

	fmt.Println("Игра быки и коровы!")
	fmt.Println("Можно вводить цифры от 0 до 9")
	fmt.Println("Для выхода введите '0000' или 'exit'")

	if *fn != "" {
		memo = *fn
	} else {
		fmt.Println("Загадайте число: ")
		fmt.Scan(&memo)
	}

	er := mu.Check(memo)
	if er != "" {
		fmt.Println(er)
		os.Exit(3)
	}

	for {
		fmt.Println("Введите число на проверку: ")
		fmt.Scan(&try)

		if try == "0000" || try == "exit" {
			break
		}

		er := mu.Check(try)
		if er != "" {
			fmt.Println(er)
			continue
		}

		if try == memo {
			fmt.Println("Вы угадали")
			break
		}

		if len(memo) != len(try) {
			fmt.Println("Длины чисел не совпадают!")
			continue
		}

		for i := 0; i < len(try); i++ {
			for y := 0; y < len(memo); y++ {
				if try[i] == memo[y] {
					if i == y {
						b++
					} else {
						k++
					}
					continue
				}
			}
		}

		fmt.Println("Обнаружено коров:", k, " и быков:", b)
		k = 0
		b = 0
	}
}

// func check(x string) error {
// 	matched, _ := regexp.Match(`\d`, []byte(x))
// 	if !matched {
// 		return errors.New("Надо вводить только цифры")
// 	}
// 	for i := 0; i < len(x); i++ {
// 		for y := 0; y < len(x); y++ {
// 			if i == y {
// 				continue
// 			}
// 			if x[i] == x[y] {
// 				return errors.New("Цифры в числе не должны посторятся")
// 			}
// 		}
// 	}
// 	return nil
// }
