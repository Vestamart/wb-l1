package l1_1

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human
	ActionDesc string
}

func (h *Human) Greeting() {
	fmt.Printf("%s ему %d лет.\n", h.Name, h.Age)
}

func (a *Action) ActionLog() {
	fmt.Printf("%s выполнил действие: %s \n", a.Name, a.ActionDesc)
}

func Start() {
	action := Action{Human: Human{
		Name: "Nikita",
		Age:  19,
	},
		ActionDesc: "Сделал задание l1.1",
	}
	action.Greeting()
	action.ActionLog()
}
