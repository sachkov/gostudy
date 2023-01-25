package muu

import (
	"regexp"
)

type Muu struct {
	bulls  int16
	cows   int16
	result string
}

func (*Muu) Check(x string) string {
	matched, _ := regexp.Match(`\d`, []byte(x))
	if !matched {
		return "Надо вводить только цифры"
	}
	for i := 0; i < len(x); i++ {
		for y := 0; y < len(x); y++ {
			if i == y {
				continue
			}
			if x[i] == x[y] {
				return "Цифры в числе не должны повторятся"
			}
		}
	}
	return ""
}
