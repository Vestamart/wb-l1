package l1_1

import "fmt"

type human struct {
	Name string
	Age  int
}

type action struct {
	human
	ActionDesc string
}

func (h *human) Greeting() {
	fmt.Printf("%s ему %d лет.\n", h.Name, h.Age)
}

func (a *action) ActionLog() {
	fmt.Printf("%s выполнил действие: %s \n", a.Name, a.ActionDesc)
}

func Start() {
	action := action{human: human{
		Name: "Nikita",
		Age:  19,
	},
		ActionDesc: "Сделал задание l1.1",
	}
	action.Greeting()
	action.ActionLog()
}
