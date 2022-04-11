package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	a := `qwe\45`
	b, err := Unpack(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}
// Unpack - распоковать строку
func Unpack(s string) (str string, err error) {
	// приводим к нижнему регистру
	s = strings.ToLower(s)
	// проверка на пустую строку
	if s == "" {
		return str, nil
	}
	// проверка на первый элемент, если он не в пределе алфавита
	if s[0] < 97 || s[0] > 122{
		return str, fmt.Errorf("incorrent first element")
	}
	// цикл по символам строки
	for i := 0; i <= len(s)-1; i++ {
		// Если символ в пределах алфавита, то добавляем букву итоговую строку
		if s[i] >= 97 && s[i] <= 122 {
			str += string(s[i])
		// Если у нас есть escape последовательность, добавляем в итоговую строку следующий символ и увеличиваем номер итерации на 1
		} else if s[i] == '\\' {
			str += string(s[i+1])
			i++
		// здесь получаем число, повторения символов. С учетом того, что один символ добавили
		} else {
			l := make([]byte, 0)
			for n := i; n <= len(s) - 1; n++{
				if s[n] >= 48 && s[n] <= 57 {
					l = append(l, s[n])
				} else {
					break
				}
			}
			fmt.Println(l)
			a, err := strconv.Atoi(string(l))
				if err != nil {
					panic(err)
				}
				y := strings.Repeat(string(s[i-1]), a-1)
				str += y
		}

	}
	return
}
