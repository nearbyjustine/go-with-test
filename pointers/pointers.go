package pointers

func ChangeString(name *string) {
	*name = "Douglas"
}

func ChangePointerString(name *string) {
	*name = "Luffy"
}
