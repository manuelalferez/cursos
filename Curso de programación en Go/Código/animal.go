package main

import "fmt"

type animal interface {
	move() string
}

type dog struct{}
type fish struct{}
type bird struct{}

func (dog) move() string {
	return "I am a dog and I am walking"
}

func (fish) move() string {
	return "I am a fish and I am swimming"
}

func (bird) move() string {
	return "I am a bird and I am flying"
}

func moveAnimal(a animal) {
	fmt.Println(a.move())
}

func main() {
	dPeter := dog{}
	moveAnimal(dPeter)
	fWilly := fish{}
	moveAnimal(fWilly)
	bCovi := bird{}
	moveAnimal(bCovi)
}
