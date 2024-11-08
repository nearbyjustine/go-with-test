package slice

import (
	"fmt"
)

type Human struct {
	name   string
	isCute bool
	age    int
}

func PrintSlice() {
	var people = []Human{
		{
			name:   "Ae",
			isCute: true,
			age:    0,
		},
	}

	fmt.Println(people)
}
