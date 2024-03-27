package puppy

import "github.com/badalsura/dog"
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