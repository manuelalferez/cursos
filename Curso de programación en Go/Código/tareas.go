package main

import "fmt"

type app struct {
	tasks []*task
}

func (a *app) addItem(t *task) {
	a.tasks = append(a.tasks, t)
}

func (a *app) removeItem(index int) {
	a.tasks = append(a.tasks[:index], a.tasks[index+1:]...)
}

func (a *app) print() {
	for i := 0; i < len(a.tasks); i++ {
		fmt.Println("Index ", i, "name: ", a.tasks[i].name)
	}
}

func (a *app) printNewWay() {
	for index, task := range a.tasks {
		fmt.Println("Index ", index, "name: ", task.name)
	}
}

type task struct {
	name        string
	description string
	done        bool
}

func (t *task) setName(newName string) {
	t.name = newName
}

func (t *task) setDescription(newDescription string) {
	t.description = newDescription
}

func (t *task) check() {
	t.done = true
}

func main() {
	t1 := &task{
		name:        "Buy bread",
		description: "Find it and buy cheap",
	}
	t2 := &task{
		name:        "Buy milk",
		description: "Find it",
	}
	t3 := &task{
		name:        "Buy sugar",
		description: "Find it and buy one",
	}
	t4 := &task{
		name:        "Buy pencils",
		description: "Find it and buy cheap",
	}
	t5 := &task{
		name:        "Buy wheels",
		description: "Find it",
	}

	firstList := &app{
		tasks: []*task{
			t1, t2,
		},
	}
	secondList := &app{
		tasks: []*task{
			t4, t5,
		},
	}

	firstList.print()
	fmt.Println("added new element")
	firstList.addItem(t3)
	firstList.print()
	fmt.Println("remove first element")
	firstList.removeItem(0)
	firstList.printNewWay()

	myMap := make(map[string]*app)
	myMap["Manuel"] = firstList
	myMap["Pedro"] = secondList

	fmt.Println("Lista de Manuel")
	myMap["Manuel"].print()

	fmt.Println("Lista de Pedro")
	myMap["Pedro"].printNewWay()
}
