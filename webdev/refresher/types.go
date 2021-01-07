package main
import "fmt"

var y int // package level scope
type person struct {
	fname string
	lname string
}
type secretAgent struct {
	p person
	haveLicenseToKill bool
}
func(p person) Speak() {
	fmt.Println("Good morning from ",p.fname)
}
func(sa secretAgent) Speak() {
	var x string
	if (sa.haveLicenseToKill) {
		x = "can"
	} else {
		x = "can't"
	}
	fmt.Println("I ",sa.p.fname," ",x, "kill")
}

type human interface {
	Speak()
}

func saySomething(h human) {
	h.Speak()
}


func main() {
	x := 10.0
	fmt.Printf("%T\n",x)
	xs := []int{1,2,3}
	fmt.Println(xs)
	m := map[string]int {
		"delcin":10,
		"beuls":4,
	}
	fmt.Println(m)

	p1 := person {
		"Delcin",
		"Raj",
	}
	fmt.Println(p1)

	p1.Speak()
	sa1 := secretAgent{
		person {
			"James",
			"Bond",
		},
		true,
	}
	saySomething(p1)
	saySomething(sa1)
}
