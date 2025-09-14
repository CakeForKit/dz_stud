package tests

import (
	task2 "dz1/internal/task_2"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindCommonWords(t *testing.T) {

	// Создаем временную директорию для тестовых файлов
	testDir := t.TempDir()

	t.Run("Нормальный случай: есть общие слова", func(t *testing.T) {
		defer emptyResultFile(t)
		// Given
		file1 := createTestFile(t, testDir, "file1.txt", "hello world go programming test example")
		file2 := createTestFile(t, testDir, "file2.txt", "world python go test sample")
		file3 := createTestFile(t, testDir, "file3.txt", "java script go world test case")

		// When
		err := task2.FindCommonWords(file1, file2, file3)

		// Then
		require.NoError(t, err)
		content := readResultFile(t)
		expected := "go\ntest\nworld\n"
		assert.Equal(t, expected, content)
	})

	t.Run("Нет общих слов", func(t *testing.T) {
		defer emptyResultFile(t)
		// Given
		file1 := createTestFile(t, testDir, "file1.txt", "hello world go")
		file2 := createTestFile(t, testDir, "file2.txt", "python java cpp")

		// When
		err := task2.FindCommonWords(file1, file2)

		// Then
		require.NoError(t, err)
		content := readResultFile(t)
		assert.Empty(t, content)
	})

	t.Run("Один файл", func(t *testing.T) {
		defer emptyResultFile(t)
		// Given
		file1 := createTestFile(t, testDir, "single.txt", "hello world test")

		// When
		err := task2.FindCommonWords(file1)

		// Then
		require.NoError(t, err)
		content := readResultFile(t)
		expected := "hello\ntest\nworld\n"
		assert.Equal(t, expected, content)
	})

	t.Run("Пустые файлы", func(t *testing.T) {
		defer emptyResultFile(t)
		// Given
		file1 := createTestFile(t, testDir, "empty1.txt", "")
		file2 := createTestFile(t, testDir, "empty2.txt", "")

		// When
		err := task2.FindCommonWords(file1, file2)

		// Then
		require.NoError(t, err)
		content := readResultFile(t)
		assert.Empty(t, content)
	})

	t.Run("Файлы с дубликатами слов", func(t *testing.T) {
		defer emptyResultFile(t)
		// Given
		file1 := createTestFile(t, testDir, "dup1.txt", "word word word test test")
		file2 := createTestFile(t, testDir, "dup2.txt", "test sample word example")

		// When
		err := task2.FindCommonWords(file1, file2)

		// Then
		require.NoError(t, err)
		content := readResultFile(t)
		expected := "test\nword\n"
		assert.Equal(t, expected, content)
	})

	t.Run("Учет регистра", func(t *testing.T) {
		defer emptyResultFile(t)
		// Given
		file1 := createTestFile(t, testDir, "case1.txt", "Hello World TEST")
		file2 := createTestFile(t, testDir, "case2.txt", "hello world test")

		// When
		err := task2.FindCommonWords(file1, file2)

		// Then
		require.NoError(t, err)
		content := readResultFile(t)
		assert.Empty(t, content)
	})

	t.Run("Разные разделители", func(t *testing.T) {
		defer emptyResultFile(t)
		// Given
		file1 := createTestFile(t, testDir, "delim1.txt", "hello\nworld\ttest  example")
		file2 := createTestFile(t, testDir, "delim2.txt", "world test sample")

		// When
		err := task2.FindCommonWords(file1, file2)

		// Then
		require.NoError(t, err)
		content := readResultFile(t)
		expected := "test\nworld\n"
		assert.Equal(t, expected, content)
	})

	t.Run("Нет файлов", func(t *testing.T) {
		defer emptyResultFile(t)
		// When
		err := task2.FindCommonWords()

		// Then
		require.NoError(t, err)
		content := readResultFile(t)
		assert.Empty(t, content)
	})

}

func TestFindCommonWords_Errors(t *testing.T) {

	testDir := t.TempDir()

	t.Run("Несуществующий файл", func(t *testing.T) {
		defer emptyResultFile(t)
		// Given
		existingFile := createTestFile(t, testDir, "existing.txt", "hello world")
		nonExistingFile := filepath.Join(testDir, "nonexistent.txt")

		// When
		err := task2.FindCommonWords(existingFile, nonExistingFile)

		// Then
		require.Error(t, err)
		assert.Equal(t, task2.ErrOpenFile, err)
	})

	t.Run("Все файлы не существуют", func(t *testing.T) {
		defer emptyResultFile(t)
		// When
		err := task2.FindCommonWords(
			filepath.Join(testDir, "nonexistent1.txt"),
			filepath.Join(testDir, "nonexistent2.txt"),
		)

		// Then
		require.Error(t, err)
		assert.Equal(t, task2.ErrOpenFile, err)
	})

	t.Run("Нет прав на чтение файла", func(t *testing.T) {
		defer emptyResultFile(t)
		// Skip on Windows
		if runtime.GOOS == "windows" {
			t.Skip("Skipping permission test on Windows")
		}

		// Skip if running as root
		if os.Geteuid() == 0 {
			t.Skip("Skipping permission test when running as root")
		}

		// Given
		noReadFile := createTestFile(t, testDir, "no_read.txt", "test content")
		err := os.Chmod(noReadFile, 0222) // Только запись, нет чтения
		require.NoError(t, err)

		// When
		err = task2.FindCommonWords(noReadFile)

		// Then
		require.Error(t, err)
		assert.Equal(t, task2.ErrOpenFile, err)
	})

	t.Run("Директория вместо файла", func(t *testing.T) {
		defer emptyResultFile(t)
		// When
		err := task2.FindCommonWords(testDir)

		// Then
		require.Error(t, err)
		assert.Equal(t, task2.ErrOpenFile, err)
	})

	t.Run("Пустой путь к файлу", func(t *testing.T) {
		defer emptyResultFile(t)
		// When
		err := task2.FindCommonWords("")

		// Then
		require.Error(t, err)
		assert.Equal(t, task2.ErrOpenFile, err)
	})
}

func TestFindCommonWords_EdgeCases(t *testing.T) {

	testDir := t.TempDir()

	t.Run("Файл только с пробелами", func(t *testing.T) {
		defer emptyResultFile(t)
		// Given
		file1 := createTestFile(t, testDir, "spaces1.txt", "   \n\t  \n ")
		file2 := createTestFile(t, testDir, "spaces2.txt", "hello world")

		// When
		err := task2.FindCommonWords(file1, file2)

		// Then
		require.NoError(t, err)
		content := readResultFile(t)
		assert.Empty(t, content)
	})

	t.Run("Специальные символы в словах", func(t *testing.T) {
		defer emptyResultFile(t)
		// Given
		file1 := createTestFile(t, testDir, "special1.txt", "hello-world test@mail go1.18")
		file2 := createTestFile(t, testDir, "special2.txt", "hello-world go1.18")

		// When
		err := task2.FindCommonWords(file1, file2)

		// Then
		require.NoError(t, err)
		content := readResultFile(t)
		expected := "go1.18\nhello-world\n"
		assert.Equal(t, expected, content)
	})

	t.Run("Много файлов", func(t *testing.T) {
		defer emptyResultFile(t)
		// Given
		files := []string{
			"common word1 word2",
			"word1 common word3",
			"word2 word3 common",
			"word4 common word5",
			"word5 word6 common",
		}

		var filePaths []string
		for i, content := range files {
			filePaths = append(filePaths, createTestFile(t, testDir,
				fmt.Sprintf("multi%d.txt", i+1), content))
		}

		// When
		err := task2.FindCommonWords(filePaths...)

		// Then
		require.NoError(t, err)
		content := readResultFile(t)
		assert.Equal(t, "common\n", content)
	})

	t.Run("Большие файлы", func(t *testing.T) {
		defer emptyResultFile(t)
		// Given
		content := ""
		for i := 0; i < 100; i++ { // Уменьшим размер для скорости
			content += "word "
		}
		content += "common"

		file1 := createTestFile(t, testDir, "large1.txt", content)
		file2 := createTestFile(t, testDir, "large2.txt", "common test")

		// When
		err := task2.FindCommonWords(file1, file2)

		// Then
		require.NoError(t, err)
		content = readResultFile(t)
		assert.Equal(t, "common\n", content)
	})
}

// Вспомогательные функции

func createTestFile(t *testing.T, dir, filename, content string) string {
	t.Helper()

	filePath := filepath.Join(dir, filename)
	err := os.WriteFile(filePath, []byte(content), 0644)
	require.NoError(t, err)

	return filePath
}

func readResultFile(t *testing.T) string {
	t.Helper()

	content, err := os.ReadFile("res.txt")
	if os.IsNotExist(err) {
		return ""
	}
	require.NoError(t, err)

	return string(content)
}

func emptyResultFile(t *testing.T) {
	t.Helper()

	err := os.WriteFile("res.txt", []byte{}, 0644)
	require.NoError(t, err)
}
