package main

import (
	"fmt"
	"math"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
func main() {
	// Проинициализовали массив длиной 5 и заполнили данными
	array := [5]int{2, 4, 6, 8, 10}
	// далее создали объект типа Wait Group из пакета Sync. Package sync provides basic synchronization primitives
	// Wait Group используется для ожидания завершения всех запущенных горутин.
	// Примечание. Если WaitGroup явно передается в функции, это следует делать с помощью Pointer.

	var wg = sync.WaitGroup{}
	// Перебираем с помощью for range перебираем каждый элемент пока не останется элементов
	// _ обозначает неиспользуемую, неименованную переменную
	for _, arg := range array {
		//Мы хотим запустить 1 concurrent task
		wg.Add(1)
		//Запустили горутину
		go squareAndPrint(float64(arg), &wg)
	}
	wg.Wait()
}

func squareAndPrint(num float64, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%f: %f | ", num, math.Pow(num, 2))
}

/*
	Add используется для установки или настройки счетчика в группе ожидания.
	Этот счетчик отслеживает количество ожидающих горутин. Положительные дельты увеличивают счетчик,
	а отрицательные дельты уменьшают его. Вам следует избегать отрицательного значения счетчика и гарантировать,
	что новое использование WaitGroup начнется только после полного завершения предыдущего использования.
*/
