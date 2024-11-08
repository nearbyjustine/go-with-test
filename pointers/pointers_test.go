package pointers

import (
	"testing"
)

func TestPointers(t *testing.T) {
	t.Run("mutate the original variable", func(t *testing.T) {
		var name = "Leo"
		ChangeString(&name)
		var expected = "Douglas"

		if name != expected {
			t.Errorf("expected '%v' but got '%v'", expected, name)
		}
	})

	t.Run("set the value of a pointer", func(t *testing.T) {
		var name *string = new(string)
		ChangePointerString(name)
		var expected = "Luffy"
		if *name != expected {
			t.Errorf("name address: %v", &name)
		}
	})

	t.Run("use another variable to change variable", func(t *testing.T) {
		var testAge = 17

		var age int = 16
		var refAge *int = new(int)

		refAge = &age
		*refAge = testAge

		if testAge != age {
			t.Errorf("expected '%v' got '%v'", testAge, age)
		}
	})
}
