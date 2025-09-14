package task2

import (
	"bufio"
	"os"
	"sort"
)

func FindCommonWords(filenames ...string) error {
	if len(filenames) == 0 {
		return nil
	}

	// Читаем слова из первого файла
	firstFileWords, err := readWordsFromFile(filenames[0])
	if err != nil {
		return ErrOpenFile
	}

	// Создаем слайс для общих слов
	commonWords := firstFileWords

	// Пересекаем с остальными файлами
	for i := 1; i < len(filenames); i++ {
		fileWords, err := readWordsFromFile(filenames[i])
		if err != nil {
			return ErrOpenFile
		}

		commonWords = intersect(commonWords, fileWords)
		if len(commonWords) == 0 {
			break // нет общих слов
		}
	}

	// Сортируем для удобства (опционально)
	sort.Strings(commonWords)

	// Записываем результат
	return writeToFile(commonWords)
}

// readWordsFromFile читает уникальные слова из файла
func readWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrOpenFile
	}
	defer file.Close()

	wordSet := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		if word != "" {
			wordSet[word] = true
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, ErrOpenFile
	}

	// Преобразуем мапу в слайс
	result := make([]string, 0, len(wordSet))
	for word := range wordSet {
		result = append(result, word)
	}

	return result, nil
}

// intersect находит пересечение двух слайсов слов
func intersect(a, b []string) []string {
	// Создаем мапу для второго слайса
	bSet := make(map[string]bool)
	for _, word := range b {
		bSet[word] = true
	}

	// Ищем пересечение
	result := make([]string, 0)
	for _, word := range a {
		if bSet[word] {
			result = append(result, word)
		}
	}

	return result
}

// writeToFile записывает слова в файл
func writeToFile(words []string) error {
	file, err := os.Create("res.txt")
	if err != nil {
		return ErrOpenFile
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, word := range words {
		_, err := writer.WriteString(word + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
