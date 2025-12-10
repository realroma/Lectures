package main

import (
	"fmt"
	"sync"
	"time"
)

func even(arr [5]int, wg *sync.WaitGroup, ch chan<- int, oddCh chan<- int, evenCh <-chan int) {
	fmt.Println("Gorutine even is srtarting.")
	for i := range arr {
		select {
		case <-evenCh: //Слушаем канал.
			ch <- arr[i]
			oddCh <- 0 //Запускаем другой поток, передав в канал, который он слушает, значение.
		default:
			fmt.Println("Wait.")
		}
	}
	wg.Done()
}

func odd(arr [5]int, wg *sync.WaitGroup, ch chan<- int, oddCh <-chan int, evenCh chan<- int) {
	fmt.Println("Gorutine odd is starting.")
	for i := range arr {
		select {
		case <-oddCh:
			ch <- arr[i]
			evenCh <- 0
		default:
			fmt.Println("Wait.")
		}
	}
	wg.Done()
}

func add(ch chan int, arr *[10]int) {
	fmt.Println("Add in queue.")
	a := 0
	for i := range ch {
		arr[a] = i
		a++
	}
}

func main() {
	//Создание массивов с значениями.
	var evenArr [5]int = [5]int{1, 3, 5, 7, 9}
	var oddArr [5]int = [5]int{2, 4, 6, 8, 10}
	var arr [10]int
	fmt.Println("Creating chanels.")
	//Создание канала с буфером на десять значений и двух каналов для взаимодействия между потоками(Горутинами).
	ch := make(chan int, 10)
	evenCh := make(chan int, 1)
	oddCh := make(chan int, 1)
	//Вносим значение для запуска потокока в определённом порядке.
	evenCh <- 0
	//Создаём мьютекс, который ожидает выполнения потоков и заполнения буфера канала для последующей записи в массив.
	wg := &sync.WaitGroup{}
	fmt.Println("Starting gorutines.")
	wg.Add(2)
	go even(evenArr, wg, ch, oddCh, evenCh)
	go odd(oddArr, wg, ch, oddCh, evenCh)
	wg.Wait()
	//Закрываем канал, закрываем все циклы слушающие данный канал и избегаем блокировки канала.
	close(ch)
	//Запись в массив из канала.
	go add(ch, &arr)
	/*Ожидание записи из канала.[Можно реализовать через print(<-ch), пока канал не получит значение main не закончится.
	 */
	time.Sleep(1 * time.Second)
	fmt.Println(arr)
}
