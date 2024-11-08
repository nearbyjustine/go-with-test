package slice

import "fmt"

type Human struct {
	Name   string
	IsCute bool
	Age    int
} 

func (h Human) PrintDescription() string {
	if h.IsCute {
		return fmt.Sprintf("%v is cute!! 🥹", h.Name)
	}
	return fmt.Sprintf("%v ain't that cute 😾", h.Name)
}

func GetPeople() []Human {
	var people = []Human{
		{
			Name:   "Ae",
			IsCute: true,
			Age:    0,
		},
	}

	return people
}
