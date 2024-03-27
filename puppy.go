package puppy

import (
	"fmt"

	"github.com/badalsura/dog"
)
func Bark() string {
	return "woof!"
}

func Barks() string {
	return "Woof! woof! woorf!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From10() {
	fmt.Println("I'm from version 1.0.0")
}