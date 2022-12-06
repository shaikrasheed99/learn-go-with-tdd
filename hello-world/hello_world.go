package helloworld

const helloPrefix = "Hello"

func HelloWorld() string {
	return "Hello World"
}

func Greet(name string) string {
	if name == "" {
		return helloPrefix
	}
	return helloPrefix + ", " + name
}
