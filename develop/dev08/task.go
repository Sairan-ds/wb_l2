package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// принимаем данные из stdin в цикле
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	}
}

func execInput(input string) error {
	// чистим строку от переноса строки и пустых мест
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSpace(input)
	// делим строку, где разделитель пробел
	args := strings.Split(input, " ")
	// запуск команд на смену директории или выход
	switch args[0] {
	case "cd" :
		if len(args) < 2 {
			return errors.New("path required")
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	// запуск CMD
	return cmd.Run()
}
