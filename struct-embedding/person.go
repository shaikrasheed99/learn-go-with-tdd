package structembedding

type Address struct {
	state   string
	country string
}

type Person struct {
	name    string
	age     int
	address Address
}

func (p Person) modifyState(state string) Person {
	p.address.state = state
	return p
}

func (p Person) modifyAge(age int) Person {
	p.age = age
	return p
}
