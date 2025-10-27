package main

import (
	"fmt"
	"math/big"
)

func performBigIntOperations(aStr, bStr string) {
	a := new(big.Int)
	b := new(big.Int)

	_, okA := a.SetString(aStr, 10)
	_, okB := b.SetString(bStr, 10)

	if !okA || !okB {
		fmt.Println("Ошибка: Не удалось преобразовать строки в big.Int.")
		return
	}

	fmt.Printf("Число A: %s\n", a.String())
	fmt.Printf("Число B: %s\n", b.String())
	fmt.Println("--------------------------------------------------")

	sum := new(big.Int).Add(a, b)
	fmt.Printf("Сложение (A + B) = %s\n", sum.String())

	diff := new(big.Int).Sub(a, b)
	fmt.Printf("Вычитание (A - B) = %s\n", diff.String())

	prod := new(big.Int).Mul(a, b)
	fmt.Printf("Умножение (A * B) = %s\n", prod.String())

	zero := big.NewInt(0)
	if b.Cmp(zero) != 0 {
		quotient := new(big.Int).Div(a, b)
		fmt.Printf("Деление (A / B) = %s\n", quotient.String())

		// Также можно найти остаток (модуль)
		rem := new(big.Int).Mod(a, b)
		fmt.Printf("Остаток (A %% B) = %s\n", rem.String())
	} else {
		fmt.Println("Деление (A / B): ОШИБКА, деление на ноль.")
	}
}

func main() {
	fmt.Println("--- Пример 1: Числа > 2^20 ---")
	performBigIntOperations("1234567", "2345678")

	fmt.Println("\n--- Пример 2: Очень большие числа (> 2^64) ---")

	aLarge := "1"
	for i := 0; i < 100; i++ {
		aLarge += "0"
	}
	bLarge := "1"
	for i := 0; i < 99; i++ {
		bLarge += "0"
	}

	performBigIntOperations(aLarge, bLarge)
}
