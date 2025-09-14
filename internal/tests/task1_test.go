package tests

import (
	task1 "dz1/internal/task_1"
	"errors"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

var (
	ErrNegNums  = errors.New("negative numbers are not allowed")
	ErrEmptyNum = errors.New("resulting number is empty")
)

func TestFilterCommonDigits(t *testing.T) {
	suite.RunSuite(t, new(FilterCommonDigitsSuite))
}

type FilterCommonDigitsSuite struct {
	suite.Suite
}

func (s *FilterCommonDigitsSuite) TestFilterCommonDigits(t provider.T) {
	t.Parallel()

	t.Run("Обычный случай: нет общих цифр", func(t provider.T) {
		t.Parallel()
		// Given
		a, b := 123, 456

		// When
		result1, result2, err := task1.FilterCommonDigits(a, b)

		// Then
		// t.Require().NoError(err)
		t.Require().Error(err)
		t.Require().Equal(123, result1)
		t.Require().Equal(456, result2)
	})

	t.Run("Есть общие цифры", func(t provider.T) {
		t.Parallel()
		// Given
		a, b := 12345, 56789

		// When
		result1, result2, err := task1.FilterCommonDigits(a, b)

		// Then
		t.Require().NoError(err)
		t.Require().Equal(1234, result1)
		t.Require().Equal(6789, result2)
	})

	t.Run("Все цифры общие - ошибка EmptyNum", func(t provider.T) {
		t.Parallel()
		// Given
		a, b := 111, 111

		// When
		result1, result2, err := task1.FilterCommonDigits(a, b)

		// Then
		t.Require().Error(err)
		t.Require().Equal(ErrEmptyNum, err)
		t.Require().Equal(0, result1)
		t.Require().Equal(0, result2)
	})

	t.Run("Отрицательные числа - ошибка NegNums", func(t provider.T) {
		t.Parallel()
		// Given
		a, b := -123, 456

		// When
		result1, result2, err := task1.FilterCommonDigits(a, b)

		// Then
		t.Require().Error(err)
		t.Require().Equal(ErrNegNums, err)
		t.Require().Equal(0, result1)
		t.Require().Equal(0, result2)
	})

	t.Run("Нулевые значения", func(t provider.T) {
		t.Parallel()
		// Given
		a, b := 0, 123

		// When
		result1, result2, err := task1.FilterCommonDigits(a, b)

		// Then
		t.Require().NoError(err)
		t.Require().Equal(0, result1)
		t.Require().Equal(123, result2)
	})

	t.Run("Большие числа с общими цифрами - ошибка EmptyNum", func(t provider.T) {
		t.Parallel()
		// Given
		a, b := 987654321, 123456789

		// When
		result1, result2, err := task1.FilterCommonDigits(a, b)

		// Then
		t.Require().Error(err)
		t.Require().Equal(ErrEmptyNum, err)
		t.Require().Equal(0, result1)
		t.Require().Equal(0, result2)
	})

	t.Run("Оба числа отрицательные", func(t provider.T) {
		t.Parallel()
		// Given
		a, b := -123, -456

		// When
		result1, result2, err := task1.FilterCommonDigits(a, b)

		// Then
		t.Require().Error(err)
		t.Require().Equal(ErrNegNums, err)
		t.Require().Equal(0, result1)
		t.Require().Equal(0, result2)
	})

	t.Run("Частично общие цифры", func(t provider.T) {
		t.Parallel()
		// Given
		a, b := 123456, 456789

		// When
		result1, result2, err := task1.FilterCommonDigits(a, b)

		// Then
		t.Require().NoError(err)
		t.Require().Equal(123, result1)
		t.Require().Equal(789, result2)
	})
}
