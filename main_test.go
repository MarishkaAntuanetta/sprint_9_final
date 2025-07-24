package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	result1 := generateRandomElements(-5)
	if result1 != nil {
		t.Error("Для отрицательного размера должна возвращаться nil")
	}

	result2 := generateRandomElements(0)
	if result2 == nil {
		t.Error("Для размера 0 должен возвращаться пустой слайс, а не nil")
	}
	if len(result2) != 0 {
		t.Error("Длина слайса при размере 0 должна быть 0")
	}

	size := 10
	result3 := generateRandomElements(size)
	if result3 == nil {
		t.Error("Для положительного размера не должен возвращаться nil")
	}
	if len(result3) != size {
		t.Errorf("Ожидалась длина %d, получена: %d", size, len(result3))
	}

	for i, value := range result3 {
		if value < 0 || value >= 100 {
			t.Errorf("Элемент [%d] = %d — вне диапазона [0, 99]", i, value)
		}
	}
}

func TestMaximum(t *testing.T) {
	data1 := []int{}
	result1 := maximum(data1)
	if result1 != 0 {
		t.Errorf("Для пустого слайса должно возвращаться 0, получено: %d", result1)
	}

	data2 := []int{42}
	result2 := maximum(data2)
	if result2 != 42 {
		t.Errorf("Максимум из [42] должен быть 42, получен: %d", result2)
	}

	data3 := []int{1, 5, 3, 9, 2}
	result3 := maximum(data3)
	if result3 != 9 {
		t.Errorf("Максимум из [1,5,3,9,2] должен быть 9, получен: %d", result3)
	}

	data4 := []int{-10, -5, -8, -1}
	result4 := maximum(data4)
	if result4 != -1 {
		t.Errorf("Максимум из отрицательных должен быть -1, получен: %d", result4)
	}

	data5 := []int{7, 7, 7, 7}
	result5 := maximum(data5)
	if result5 != 7 {
		t.Errorf("Максимум из [7,7,7,7] должен быть 7, получен: %d", result5)
	}
}

func TestMaxChunks(t *testing.T) {
	data1 := []int{}
	result1 := maxChunks(data1)
	if result1 != 0 {
		t.Errorf("Для пустого слайса должно возвращаться 0, получено: %d", result1)
	}

	data2 := []int{100}
	result2 := maxChunks(data2)
	if result2 != 100 {
		t.Errorf("Максимум из [100] должен быть 100, получен: %d", result2)
	}

	data3 := []int{3, 7, 2, 9, 1, 8, 5, 4}
	result3 := maxChunks(data3)
	expected := 9
	if result3 != expected {
		t.Errorf("Ожидалось %d, получено: %d", expected, result3)
	}

	data4 := []int{5, 5, 5, 5, 5}
	result4 := maxChunks(data4)
	if result4 != 5 {
		t.Errorf("Максимум из пятерок должен быть 5, получен: %d", result4)
	}

	data5 := []int{1, 99, 3, 45, 88, 7, 91, 4, 50}
	max1 := maximum(data5)
	max2 := maxChunks(data5)
	if max1 != max2 {
		t.Errorf("Результаты должны совпадать: maximum=%d, maxChunks=%d", max1, max2)
	}
}
