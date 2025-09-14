package main

import task1 "dz1/internal/task_1"

// Пример использования
func main() {
	// Пример 1: нормальная работа
	result1, result2, err := task1.FilterCommonDigits(123, 456)
	if err != nil {
		println("Ошибка:", err.Error())
	} else {
		println("Результат:", result1, result2) // 123, 456
	}

	// Пример 2: общие цифры
	result1, result2, err = task1.FilterCommonDigits(123, 345)
	if err != nil {
		println("Ошибка:", err.Error())
	} else {
		println("Результат:", result1, result2) // 12, 45
	}

	// Пример 3: отрицательные числа
	result1, result2, err = task1.FilterCommonDigits(-123, 456)
	if err != nil {
		println("Ошибка:", err.Error()) // negative numbers are not allowed
	}

	// Пример 4: пустой результат
	result1, result2, err = task1.FilterCommonDigits(111, 111)
	if err != nil {
		println("Ошибка:", err.Error()) // resulting number is empty
	}
}
