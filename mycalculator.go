package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	input    string
	isarabic bool
}

func (num *Number) RomeToInt(romenumericstoint map[string]int) {
	var buffer []int
	var buffer_elem string
	var sum int
	if strings.Count(num.input, "V") > 1 || strings.Count(num.input, "L") > 1 {
		panic("Ошибка! В римской системе нет такого числа!")
	}
	for i := range num.input {
		if strings.Count(num.input, string(num.input[i])) < 4 {
			buffer_elem += string(num.input[i])
			if _, ok := romenumericstoint[buffer_elem]; ok {
				if i == len(num.input)-1 {
					buffer = append(buffer, romenumericstoint[buffer_elem])
				}
			} else if _, ok := romenumericstoint[string(num.input[i])]; ok {
				buffer = append(buffer, romenumericstoint[buffer_elem[:i]])
				buffer_elem = string(num.input[i])
				if i == len(num.input)-1 {
					buffer = append(buffer, romenumericstoint[buffer_elem])
				}
			} else {
				panic("Ошибка! В римской системе нет такого числа!")
			}
		} else {
			panic("Ошибка! В римской системе нет такого числа!")
		}
	}
	for i := range buffer {
		if i != len(buffer)-1 {
			if buffer[i+1] > buffer[i] {
				panic("Ошибка! В римской системе нет такого числа!")
			}
		}
		sum += buffer[i]
	}
	if sum > 10 {
		panic("Ошибка! Введены некорректные данные!")
	}
	num.input = strconv.Itoa(sum)
	num.isarabic = false
}
func (num *Number) IsAllowable(romenumerics map[string]int) {
	number, err := strconv.Atoi(num.input)
	if err != nil {
		num.RomeToInt(romenumerics)
	}
	if number > 10 {
		panic("Ошибка! Введены некорректные данные!")
	}

}
func (num *Number) Calculation() int {
	number, _ := strconv.Atoi(num.input)
	return number
}
func IntToRome(responseInt int, inttoromenumerics map[int]string) string {
	var responseRome string
	for i := 100; i > 0; i /= 10 {
		responseRome += inttoromenumerics[(responseInt/i)*i]
		responseInt -= (responseInt / i) * i
	}
	return responseRome
}

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text := strings.Split(input, " ")
	if len(text) != 4 {
		panic("Ошибка! Введены некорректные данные!")
	}
	number1 := Number{text[0], true}
	number2 := Number{text[2], true}
	operation := text[1]
	romenumericstoint := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
		"XX":   20,
		"XXX":  30,
		"XL":   40,
		"L":    50,
		"LX":   60,
		"LXX":  70,
		"LXXX": 80,
		"XC":   90,
		"C":    100,
	}
	inttoromenumerics := map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		20:  "XX",
		30:  "XXX",
		40:  "XL",
		50:  "L",
		60:  "LX",
		70:  "LXX",
		80:  "LXXX",
		90:  "XC",
		100: "C",
	}
	operations := map[string]func() int{
		"+": func() int { return number1.Calculation() + number2.Calculation() },
		"-": func() int { return number1.Calculation() - number2.Calculation() },
		"*": func() int { return number1.Calculation() * number2.Calculation() },
		"/": func() int { return number1.Calculation() / number2.Calculation() },
	}
	if _, ok := operations[operation]; ok {
		number1.IsAllowable(romenumericstoint)
		number2.IsAllowable(romenumericstoint)
		if number1.isarabic != number2.isarabic {
			panic("Ошибка! Используются разные системы счисления!")
		} else if number1.isarabic == true {
			fmt.Println(operations[operation]())
		} else {
			if number1.input <= number2.input && (operation == "-" || operation == "/") {
				panic("Ошибка! Введены некорректные данные!")
			}
			fmt.Println(IntToRome(operations[operation](), inttoromenumerics))
		}
	} else {
		panic("Ошибка! Невозможно совершить операцию!")
	}
}
