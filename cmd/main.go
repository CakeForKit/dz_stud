package main

import (
	task1 "dz1/internal/task_1"
	task2 "dz1/internal/task_2"
	task3 "dz1/internal/task_3"
	"fmt"
	"os"
)

func example_1() {
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
	_, _, err = task1.FilterCommonDigits(-123, 456)
	if err != nil {
		println("Ошибка:", err.Error()) // negative numbers are not allowed
	}

	// Пример 4: пустой результат
	_, _, err = task1.FilterCommonDigits(111, 111)
	if err != nil {
		println("Ошибка:", err.Error()) // resulting number is empty
	}
}

func example_2() {
	// Задание 2
	f1 := "./task_2/files/file1.txt"
	f2 := "./task_2/files/file2.txt"
	f3 := "./task_2/files/file3.txt"
	// Пример 1: нормальная работа с общими словами
	fmt.Println("Пример 1: нормальная работа")
	err := task2.FindCommonWords(f1, f2, f3)
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

func example_3() {
	// Пример 1: Нормальное масштабирование
	slice1 := []int{1, 2, 3}
	err := task3.ScaleSlice(&slice1, 3)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Результат: %v\n", slice1) // [1 2 3 1 2 3 1 2 3]
	}

	// Пример 2: Переполнение
	slice2 := make([]int, 1000000)        // 1,000,000 элементов
	err = task3.ScaleSlice(&slice2, 5000) // 1,000,000 * 5,000 = 5,000,000,000 > 4,294,967,295
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err) // ErrOverflow
	} else {
		fmt.Printf("Результат: %v\n", slice2)
	}

	// Пример 3: Коэффициент 0
	slice3 := []int{1, 2, 3}
	err = task3.ScaleSlice(&slice3, 0)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Результат: %v\n", slice3) // []
	}

	// Пример 4: Пустой срез
	slice4 := []int{}
	err = task3.ScaleSlice(&slice4, 5)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Результат: %v\n", slice4) // []
	}

	// Пример 5: Коэффициент 1
	slice5 := []int{1, 2, 3}
	err = task3.ScaleSlice(&slice5, 1)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Результат: %v\n", slice5) // [1 2 3]
	}
}

// Пример использования
func main() {
	example_1()
	example_2()
	example_3()
}
