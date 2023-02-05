package airportrobot

import "fmt"

func main() {
	germanGreeter := German{}
	italianGreeter := Italian{}
	portuguesGreeter := Portuguese{}

	fmt.Println(SayHello("Dietrich", germanGreeter)) // => "I can speak German: Hallo Dietrich!"
	fmt.Println(SayHello("Kevin", italianGreeter))   // => "I can speak Italian: Ciao Kevin!"
	fmt.Println(SayHello("Kevin", portuguesGreeter)) // => "I can speak Portuguese: Olá Kevin!"
}

/*
Como primer paso, defina una interfaz Greeter con dos métodos.

LanguageName que devuelve el nombre del idioma (a string) en el que se supone que el robot saluda al visitante.
Greet que acepta el nombre de un visitante (a string) y devuelve un string con el mensaje de saludo en un idioma específico.

A continuación, implemente una función SayHello que acepte el nombre del visitante y todo lo que implemente la interfaz Greeter como argumentos y devuelva la cadena de saludo deseada.
Por ejemplo, imagine una implementación Greeter con alemana para la cual; LanguageName regresa "German" y Greet regresa "Hallo {name}!":
*/
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s!", g.LanguageName(), g.Greet(name))
}

// Declarate type German language
type German struct{}

// Call interfaz "methods" Greeter as German language
func (g German) LanguageName() string {
	return "German"
}

func (g German) Greet(name string) string {
	return fmt.Sprintf("Hallo %s", name)
}

// Call interfaz Greeter as Italian language
type Italian struct{}

// Call interfaz "methods" Greeter as Italian language
func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s", name)
}

// Call interfaz Greeter as Portuguese language
type Portuguese struct{}

// Call interfaz "methods" Greeter as Portuguese language
func (p Portuguese) LanguageName() string {
	return "Portuguese"
}
func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s", name)
}
