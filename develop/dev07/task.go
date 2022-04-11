package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func main() {
	// определяем функции которая будет принимать каналы
	or := func(channels ...<-chan interface{}) <-chan interface{} {
		// создаем single канал, который будет отправлен по завершении одного из каналов. 
		outCh := make(chan interface{})
		// создаем wg
		var wg sync.WaitGroup
		wg.Add(1)
		// Цикл который будет запускать горутины
		for _, channel := range channels {
			go func(channel <-chan interface{}) {
				// запуск цикла по каналу, при закрытии канала, произойдет выход из цикла, и отнимем wg
				for range channel {}
				wg.Done()
			}(channel)
		}
		
		wg.Wait()
		// закрыаем сингл канал и возвращаем его. 
		close(outCh)
		return outCh
	}
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v \n", time.Since(start))
}
