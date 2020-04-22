# Go

## Primer programa en Go

```go
package main 

import "fmt" // Importaciones en Go

func main(){
	fmt.Println("Hola Mundo") // La función fmt sirve para hacer impresiones en pantalla
}
```

[Descarga de Go](https://golang.org/dl/)



## Comandos básicos 

```bash
go build [filename]
# Este comando sirve para compilar el archivo que le indiquemos. Eso genera un ejecutable, es decir, un binario.

go run [filename]
# Ejecuta el binario generado
```



### Variables y funciones 

```go
package main

import "fmt"

func main(){
	var nombre string = "pepito"
	apellido := "romero"
	fmt.Println(nombre)
	fmt.Println(apellido)

	a := 10.
	var b float64 = 3
	fmt.Println(a/b)

	x := true
	y := false
	fmt.Println(x || y)
	fmt.Println(x && y)
	fmt.Println(!x)
	privada()
	Publica()

}

// Declaración de una función privada, su uso queda restringido al paquete main
func privada(){
	fmt.Println("soy privada")
}

// Declaración de una función pública, para todos los paquetes 
func Publica(){
	fmt.Println("soy pública")
}


func decirHola(){
	defer fmt.Println("mundo") // defer se ejecutará al final
	fmt.Print("hola ")
}

```



## Leer de teclado 

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Introduzca operación: ")
	scanner.Scan()
	operacion := scanner.Text()
	valores := strings.Split(operacion, "+") // Troceamos 
	fmt.Println(valores)
	fmt.Println(valores[0] + valores[1])
	operador1, _ := strconv.Atoi(valores[0]) // Convertimos a entero
	operador2, _ := strconv.Atoi(valores[1])

	fmt.Println(operador1 + operador2)
}

```



## Manejo de errores y condicionales

```go
	operador1, error1 := strconv.Atoi(valores[0])
	operador2, _ := strconv.Atoi(valores[1])

	if error1 != nil {
		fmt.Println(error1)
	} else {
		fmt.Println("Transformación correcta")
	}
```



## Switch 

```go
switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
	case "-":
		fmt.Println(operador1 - operador2)
	case "*":
		fmt.Println(operador1 * operador2)
	case "/":
		fmt.Println(operador1 / operador2)
	default:
```



## Struct y receivers

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calculator struct{}

// calculate podrá ser usada como parte de calculator, y hacer llamadas desde él
func (calculator) calculate(operators string, operation string) int {
	numbers := strings.Split(operators, operation)
	firstOperator := parseInt(numbers[0])
	secondOperator := parseInt(numbers[1])

	switch operation {
	case "+":
		return firstOperator + secondOperator
	case "-":
		return firstOperator - secondOperator
	case "*":
		return firstOperator * secondOperator
	case "/":
		return firstOperator / secondOperator
	default:
		return 0
	}
}

func read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func parseInt(str string) int {
	result, _ := strconv.Atoi(str)
	return result
}

func main() {
	fmt.Println("Introduce operation (+,-,* ó /): ")
	operation := read()
	fmt.Println("Introduce operators (ex: 2+2): ")
	operators := read()
	c := calculator{}
	fmt.Println("Result: ")
	fmt.Println(c.calculate(operators, operation))
}

```



## Struct con variables 

```go
type task struct {
	name        string
	description string
	done        bool
}

func (t task) setName(newName string) {
	t.name = newName
}
```



## Punteros 

```go
func main() {
	x := 25
	fmt.Println(x)
	cambiarValor(&x)
	fmt.Println(x)
}

func cambiarValor(a *int) {
	*a = 30
}
```

 ```go
func (t *task) setName(newName string) {
	t.name = newName
}

t := &task{
		name:        "Buy bread",
		description: "Find it and buy cheap",
	}
 ```



## Slides 

```go
package main

import "fmt"

type app struct {
	tasks []*task
}

func (a *app) addItem(t *task) {
	a.tasks = append(a.tasks, t)
}

func (a *app) removeItem(index int) {
	a.tasks = append(a.tasks[:index], a.tasks[index+1:]...)
}

type task struct {
	name        string
	description string
	done        bool
}

func (t *task) setName(newName string) {
	t.name = newName
}

func (t *task) setDescription(newDescription string) {
	t.description = newDescription
}

func (t *task) check() {
	t.done = true
}

func main() {
	t1 := &task{
		name:        "Buy bread",
		description: "Find it and buy cheap",
	}
	t2 := &task{
		name:        "Buy milk",
		description: "Find it",
	}
	t3 := &task{
		name:        "Buy sugar",
		description: "Find it and buy one",
	}
	myAppList := &app{
		tasks: []*task{
			t1, t2,
		},
	}

	fmt.Println(myAppList.tasks[0])

	myAppList.addItem(t3)
	fmt.Println(len(myAppList.tasks))
	myAppList.removeItem(0)
	fmt.Println(len(myAppList.tasks))
}

```



## Ciclo for y range

```go
func (a *app) print() {
	for i := 0; i < len(a.tasks); i++ {
		fmt.Println("Index ", i, "name: ", a.tasks[i].name)
	}
}

func (a *app) printNewWay() {
	for index, task := range a.tasks {
		fmt.Println("Index ", index, "name: ", task.name)
	}
}
```



## Maps

```go
	myMap := make(map[string]*app)
	myMap["Manuel"] = firstList
	myMap["Pedro"] = secondList

	fmt.Println("Lista de Manuel")
	myMap["Manuel"].print()

	fmt.Println("Lista de Pedro")
	myMap["Pedro"].printNewWay()
```



## Interfaces

```go
package main

import "fmt"

type animal interface {
	move() string
}

type dog struct{}
type fish struct{}
type bird struct{}

func (dog) move() string {
	return "I am a dog and I am walking"
}

func (fish) move() string {
	return "I am a fish and I am swimming"
}

func (bird) move() string {
	return "I am a bird and I am flying"
}

func moveAnimal(a animal) {
	fmt.Println(a.move())
}

func main() {
	dPeter := dog{}
	moveAnimal(dPeter)
	fWilly := fish{}
	moveAnimal(fWilly)
	bCovi := bird{}
	moveAnimal(bCovi)
}
```



## Imprimiendo el contenido de una Página Web usando Interfaces

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

type webWriter struct{}

func (webWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	response, error := http.Get("http://google.com")
	webW := webWriter{}
	if error != nil {
		fmt.Println(error)
	}
	io.Copy(webW, response.Body)
}
```



## Introducción al problema de la concurrencia  

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkServer(server string) {
	_, error := http.Get(server)
	if error != nil {
		fmt.Printf(server, "is not working :(\n")
	} else {
		fmt.Printf(server, "is working :)\n")
	}
}

func main() {
	startTime := time.Now()
	servers := []string{
		"http://google.com",
		"http://twitter.com",
		"http://github.com",
		"http://youtube.com",
	}
	for _, server := range servers {
		checkServer(server)
	}
	diffTime := time.Since(startTime)
	fmt.Println(diffTime)
}

```

