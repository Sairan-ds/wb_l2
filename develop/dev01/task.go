package main

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

import (
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	t := CurrentTime()
	log.Printf("Current time is : %s", t)
}
// CurrentTime Получить точное время с использованием  NTP библиотеки
func CurrentTime() time.Time {
	l := log.New(os.Stderr, "", 0)
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		l.Fatal(err.Error())
	}
	return t
}
