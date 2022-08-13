package interfacee

import "fmt"

func MainInterface() {
	var i interface{}

	i = 10
	fmt.Printf("type is %T, value is %v\n", i, i)

	i = "ten"
	fmt.Printf("type is %T, value is %v\n", i, i)

	if s, ok := i.(string); ok {
		fmt.Printf("type is %T, value is %v\n", s, s)
	}
	//.(string) -> Type Assertion

	i = struct {
		number int
		text   string
	}{
		number: 10,
		text:   "ten",
	}
	fmt.Printf("type is %T, value is %v\n", i, i)

	i = func() string {
		return "10"
	}
	fmt.Printf("type is %T, value is %v\n", i, i)

}

type color interface {
	Color() string
}

type cat struct{}

func (c cat) Color() string {
	return "Blact"
}

type dog struct{}

func (d dog) Color() string {
	return "White"
}

func printAnimalColor(c color) {
	fmt.Println(c.Color())
}

func MainAnimalColor() {
	c := cat{}
	printAnimalColor(c)
	d := dog{}
	printAnimalColor(d)
}
