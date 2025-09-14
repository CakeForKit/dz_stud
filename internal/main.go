package main

import (
	task1 "dz1/internal/task_1"
	task2 "dz1/internal/task_2"
	"fmt"
	"os"
)

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
	_ = result1
	_ = result2

	// Задание 2
	f1 := "./task_2/files/file1.txt"
	f2 := "./task_2/files/file2.txt"
	f3 := "./task_2/files/file3.txt"
	// Пример 1: нормальная работа с общими словами
	fmt.Println("Пример 1: нормальная работа")
	err = task2.FindCommonWords(f1, f2, f3)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Println("Результат записан в res.txt")
		// Читаем и выводим результат
		content, _ := os.ReadFile("res.txt")
		fmt.Printf("Общие слова: %s\n", string(content))
	}
	fmt.Println()

	// Пример 2: несуществующий файл
	fmt.Println("Пример 2: несуществующий файл")
	err = task2.FindCommonWords(f1, "nonexistent.txt")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err) // ErrOpenFile
	} else {
		fmt.Println("Успешно обработано")
	}
	fmt.Println()

	// Пример 3: нет общих слов
	fmt.Println("Пример 3: нет общих слов")
	err = task2.FindCommonWords(f1, "file4.txt")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Println("Результат записан в res.txt")
		content, _ := os.ReadFile("res.txt")
		if len(content) == 0 {
			fmt.Println("Нет общих слов")
		} else {
			fmt.Printf("Общие слова: %s\n", string(content))
		}
	}
	fmt.Println()

	// Пример 4: пустые файлы
	fmt.Println("Пример 4: пустые файлы")
	err = task2.FindCommonWords("empty1.txt", "empty2.txt")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Println("Результат записан в res.txt")
		content, _ := os.ReadFile("res.txt")
		if len(content) == 0 {
			fmt.Println("Нет общих слов")
		}
	}
	fmt.Println()

	// Пример 5: только один файл
	fmt.Println("Пример 5: только один файл")
	err = task2.FindCommonWords(f1)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Println("Результат записан в res.txt")
		content, _ := os.ReadFile("res.txt")
		fmt.Printf("Слова из файла: %s\n", string(content))
	}
	fmt.Println()

	// Пример 6: нет файлов
	fmt.Println("Пример 6: нет файлов")
	err = task2.FindCommonWords()
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Println("Нет файлов для обработки")
	}
}
