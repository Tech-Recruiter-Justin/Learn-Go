package hello

const (
	englishGreeting = "Hello, "
	spanishGreeting = "Hola, "
	frenchGreeting  = "Bonjour, "
)

func SayHello(name string, lang string) string {

	if name == "" {
		name = "world"
	}

	switch lang {
	case "Spanish":
		return spanishGreeting + name
	case "French":
		return frenchGreeting + name
	}
	return englishGreeting + name
}
