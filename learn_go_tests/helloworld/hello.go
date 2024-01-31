package helloworld

const (
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name string, language string) string {
	prefixMessage := greetingPrefix(language)

	if len(name) == 0 {
		name = "World"
	}

	return prefixMessage + name

}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = spanishHelloPrefix
	case "french":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
