package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаём канал
	dataChannel := make(chan int)

	// Переменная для времени работы программы в секундах
	N := 5

	// Запускаем горутин для отправки данных
	go func() {
		counter := 0 // Счётчик отправляемых значений
		for {
			counter++
			dataChannel <- counter              // Отправляем значение в канал
			fmt.Println("Отправлено:", counter) // Вывод сообщения об отправленной переменной
			time.Sleep(1 * time.Second)         // Ждём 1 секунду перед следующей отправкой
		}
	}()

	// Запускаем горутину для чтения данных
	go func() {
		for {
			val := <-dataChannel          // Получаем значение из канала
			fmt.Println("Получено:", val) // Вывод сообщения о полученной переменной
		}
	}()

	// Ждём N секунд
	time.Sleep(time.Duration(N) * time.Second)

	// Завершаем программу
	fmt.Println("Программа завершена")
}
