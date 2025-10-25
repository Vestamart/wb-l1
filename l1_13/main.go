package main

import "fmt"

func SwapArithmetic(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func SwapXOR(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func main() {
	a1 := 10
	b1 := 25
	fmt.Println("Арифметический метод")
	fmt.Printf("До обмена: a = %d, b = %d\n", a1, b1)

	a1, b1 = SwapArithmetic(a1, b1)

	fmt.Printf("После обмена: a = %d, b = %d\n", a1, b1)

	a2 := 10
	b2 := 25
	fmt.Println("Метод XOR")
	fmt.Printf("До обмена: a = %d, b = %d\n", a2, b2)

	a2, b2 = SwapXOR(a2, b2)

	fmt.Printf("После обмена: a = %d, b = %d\n", a2, b2)

	a3 := 10
	b3 := 25
	fmt.Println("Параллельное присваивание")
	fmt.Printf("До обмена: a = %d, b = %d\n", a3, b3)

	a3, b3 = b3, a3
	fmt.Printf("После обмена: a = %d, b = %d\n", a3, b3)

}
