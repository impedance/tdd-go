package main
// import "fmt"

const engilshHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "


func Hello(name string, language string) string {
  if name == "" {
    name = "World"
  }

  return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {

  switch language {
  case "French":
    prefix = frenchHelloPrefix
  case "Spanish":
    prefix = spanishHelloPrefix
  default:
    prefix = engilshHelloPrefix
  }
  return 
}

func main() {
}
