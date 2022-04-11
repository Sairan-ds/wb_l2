package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	testStr := []string{"Пятак", "пятак", "пятка", "тяпка", "листок", "слиток", "столик"}

	fmt.Println(anagramFinder(testStr))
}

//  Функция
func anagramFinder(str []string) *map[string][]string {
	// Удалямем дубликаты и строки приводим к нижнему регистру
	s := deleteDublicate(str)
	// Временная мапа куда бы добавляем все повторения а ключом является отсортированный набор букв слова.
	tempMap := make(map[string][]string) // map[акптя:[пятак пятка тяпка] иклост:[листок слиток столик]]  получится подобная мапа
	for _, val := range s {
		tempMap[strSort(val)] = append(tempMap[strSort(val)], val)
	}

	// мапа множест анаграм, у которой ключ первое встретившейся в словаре слово из множества. Значение ссылка на массив, каждый элемент которого, слово из
	//множества
	endMap := make(map[string][]string)
	for _, v := range tempMap {
		endMap[v[0]] = v[1:]
	}
	return &endMap
}

// сортировка строк
func strSort(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// реализация структуры и интерфейса для сортировки
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}
// приводим к нижнему регистру и удаляем дубликаты
func deleteDublicate(s []string) []string {
	temp := map[string]struct{}{}
	for _, r := range s {
		temp[strings.ToLower(r)] = struct{}{}
	}
	str := make([]string, 0, len(temp))
	for k := range temp {
		str = append(str, k)
	}
	return str
}
