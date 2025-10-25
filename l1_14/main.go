package main

import "fmt"

func DescribeType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("Переданное значение %v имеет тип: int\n", v)
	case string:
		fmt.Printf("Переданное значение \"%v\" имеет тип: string\n", v)
	case bool:
		fmt.Printf("Переданное значение %v имеет тип: bool\n", v)
	case chan int:
		fmt.Printf("Переданное значение %v имеет тип: chan int (канал)\n", v)
	case nil:
		fmt.Println("Переданное значение имеет тип: nil (нулевое значение интерфейса)")
	default:
		// Сюда попадут все остальные типы, включая chan string, float64, struct и т.д.
		fmt.Printf("Переданное значение %v имеет нераспознанный тип: %T\n", v, v)
	}
}

func main() {
	DescribeType(42)

	DescribeType("Hello, WB!")

	DescribeType(true)

	intChan := make(chan int)
	DescribeType(intChan)

	DescribeType(3.14)

	var empty interface{}
	DescribeType(empty)

	strChan := make(chan string)
	DescribeType(strChan)
}
