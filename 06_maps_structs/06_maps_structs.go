package main

import (
	"fmt"
	"reflect"
)

func mapOperations() {

	// The zero value of a map is nil. A nil map has no keys, nor can keys be added. eg. var m[string]string, so m value is nil, we gotta use make function or map[string][string]{}

	fmt.Println("mapOperations:")

	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      2750293,
		"Florida":    2932398,
		"New York":   29230929,
		"Ohio":       2039283092,
	}
	fmt.Println(statePopulations)

	// m := map[[]int]string{} not possible with slice

	m := map[[3]int]string{}
	n := make(map[string]int)
	o := make(map[string]int, 10)
	fmt.Println(m, n, o)

	fmt.Println(statePopulations["California"])
	fmt.Println(statePopulations)
	statePopulations["California"] = 600
	fmt.Println(statePopulations)

	//* THE ORDER OF THE MAP IS NOT FIXED, IT KEEPS CHANGING

	// delete
	delete(statePopulations, "California")
	fmt.Println(statePopulations)

	// keys not present return a 'Zero Value'
	value, ok := statePopulations["California"]
	_, ok2 := statePopulations["California"] // if u just want to check if the key is present in the map, u can use _ write only variable
	fmt.Println(value, ok, ok2)

	// length
	fmt.Println(len(statePopulations))

	//* map are reference-type,
	// manipulating one will affect the other references
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)

	// If the top-level type is just a type name, you can omit it from the elements of the literal.

	type Vertex struct {
		Lat, Long float64
	}

	var z = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(z)

}

type Doctor struct {
	number     int
	actorName  string
	companions []string
	episodes   []string
}

// same as other for nameing, PascalCase - export, camelCase - package scope
//* this applies both to the struct name - Doctor and the fields, if the fields start with lower case if cannot be accessed outside the package (like private), if they start with upper case then they can be export and accessed outside this package too eg.

type Rectangle struct {
	Length  float32
	Breadth float32
}

func structOperations() {

	// * A struct is a collection of fields.

	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.number, aDoctor.actorName, aDoctor.companions)
	fmt.Println(aDoctor.companions[1])

	// recomended to NOT use this syntax, as if we add new variables or change the order of the Doctor type then we will have to change the order here too and all the values / at least empty values has to be specified
	b := Doctor{
		3,
		"Jon Pertwee",
		[]string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
		[]string{},
	}
	fmt.Println(b)

	// using this syntax we can mention the values for the fields we want to and in any order
	c := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
	}
	d := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		episodes: []string{
			"pilot",
			"parent",
		},
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}
	fmt.Println(c)
	fmt.Println(d)

	// anonymous struct
	// used when we have to create a temp struct type (as a json) as a reply to an api call request
	e := struct{ name string }{name: "abc"}
	fmt.Println(e)

	//* STRUCT ARE VALUE-TYPE
	f := e
	f.name = "def"
	fmt.Println("e:", e)
	fmt.Println("f:", f)

	//* we can use pointers to point to the same struct object

	g := &e
	g.name = "def"
	fmt.Println("e:", e)
	fmt.Println("g:", g)

	//* struct doesn't have inheritance, It has composition through embedding

	embedding()

	tags()

}

func embedding() {

	type Animal struct {
		name   string
		origin string
	}

	type Bird struct {
		Animal
		speedKPH float32
		canFly   bool
	}

	//* Bird2 is different than Bird
	type Bird2 struct {
		animal   Animal
		speedKPH float32
		canFly   bool
	}

	b := Bird{}
	b.name = "Emu"
	b.origin = "Australia"
	b.speedKPH = 48
	b.canFly = false
	fmt.Println(b)

	c := Bird{
		Animal:   Animal{name: "Emu", origin: "Australia"},
		speedKPH: 48,
		canFly:   false,
	}
	fmt.Println(c)
	fmt.Println(c.name, c.canFly)

	d := Bird2{
		animal:   Animal{name: "Emu", origin: "Australia"},
		speedKPH: 48,
		canFly:   false,
	}
	fmt.Println(d)
	fmt.Println(d.animal.name, d.canFly)
	// Bird2 cannot directly access the fields of Animal

	// embedding should only used in such a scenario where eg. we have different complex controllers, then we can have a base controller and then embed that base controller in all the different complex controllers so that we can get base behaviour into a custom type (and there no actual use of polymorphism or inheritance)
	// for any other use case look into interface

}

func tags() {

	// tags are used to specify some specific information about that field

	// all tags do is provide a string of text and something else has to figure out what to do with it, some other validation library has to be used to parse it and then apply logic based on what is specified in that tag

	// also used for json representations

	type Animal struct {
		name   string `required,max:"100"`
		origin string
	}

	// here the tag can specify that the name field should be required and max length should be 100

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)

	// https://stackoverflow.com/questions/59959839/not-compatible-with-reflect-structtag-get

}

func main() {

	mapOperations()

	structOperations()

}
