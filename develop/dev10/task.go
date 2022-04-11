package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	// парсим таймаут, либо задаем стандарные 10 сек
	timeout := flag.Int("timeout", 10, "timeout for connection")
	flag.Parse()
	// проверка на количество аргументов
	if flag.NArg() != 2 {
		panic("incorrect args")
	}
	// складываем адресс из ip и port
	args := flag.Args()
	address := args[0] + ":" + args[1]
	// подключаемся по tcp к адрессу, и таймаутом на подключение
	conn, err := net.DialTimeout("tcp", address, time.Duration(*timeout)*time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	// по окончании разрыв функции
	defer conn.Close()
	// настраиваем контекст для завершения программы
	ctx, cancel := context.WithCancel(context.Background())
	// канал на прием сигнала из ОС
	signalChan := make(chan os.Signal)
	// запись в канал сигнала из ОС
	signal.Notify(signalChan, syscall.SIGINT)
	// горутина на ожидание сигнала из ОС 
	go func() {
		<-signalChan
		cancel()
	}()
	// запуск в горутине функции общения с сервером
	go Connect(conn, cancel)
		// ожидание завершения контекста
	<-ctx.Done()
	fmt.Println("\nConnection closed")
}

func Connect(conn net.Conn, cancelFunc context.CancelFunc) {
	// читаем из conn, записываем в stdout
	for {

		// Чтение входных данных от stdin
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		// Отправляем в socket
		fmt.Fprintf(conn, text+"\n")
		// Прослушиваем ответ
		message, _ := bufio.NewReader(conn).ReadString('\n')
		

		fmt.Print("Message from server: " + message)

	}

}
