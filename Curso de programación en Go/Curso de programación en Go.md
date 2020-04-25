# Go

[Info](https://es.wikipedia.org/wiki/Go_(lenguaje_de_programaci%C3%B3n))

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



## Goroutunes y channels 

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkServer(server string, channel chan string) {
	_, error := http.Get(server)
	if error != nil {
		fmt.Printf(server, "is not working :(\n")
		channel <- server + "is not working :(\n"
	} else {
		fmt.Printf(server, "is working :)\n")
		channel <- server + "is working :)\n" // forma de enviar al canal
	}
}

func main() {
	startTime := time.Now()
	channel := make(chan string)
	servers := []string{
		"http://google.com",
		"http://twitter.com",
		"http://github.com",
		"http://youtube.com",
	}
	for _, server := range servers {
		go checkServer(server, channel) // Este código se ejecutará concurrentemente 
	}
	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel) // Necesario para mantener un canal de comunicación entre las goroutines y el hilo principal (main)
	}
	diffTime := time.Since(startTime)
	fmt.Println(diffTime)
}

```



## Ciclos while en go

```go
	counter := 0
	for {
		if counter > 5 {
			break
		}
		for _, server := range servers {
			go checkServer(server, channel)
		}
		time.Sleep(5 * time.Second) // duerme 5 segundos
		counter++
	}
```



## Creando el servidor y manejando rutas en backend

> file: main.go

```go
package main

func main() {
	myServer := newServer(":3000")
	myServer.listen()
}
```

> file: server.go

```go
package main

import "net/http"

/**
	Crea una conexión y escucha
 */
type Server struct {
	port   string
	router *Router
}

func newServer(port string) *Server {
	return &Server{
		port:   port,
		router: newRouter(),
	}
}

func (s *Server) listen() error {
	http.Handle("/", s.router) // [3]
	error := http.ListenAndServe(s.port, nil)
	if error != nil {
		return error
	}
	return nil
}
```

> file: router.go
>

```go
package main

import (
	"fmt"
	"net/http"
)

/**
	Maneja las URLs que nos llegan
 */
type Router struct {
	rules map[string]http.HandlerFunc // cada string (url) le asociamos un Handler [1]
}

func newRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, responde *http.Request) {
	fmt.Fprintln(w, "Hello World!") // el primer parámetro es dónde se va escribir y el segundo qué
}
```



[1]

## type [HandlerFunc](https://golang.org/src/net/http/server.go?s=61706:61753#L1998)

The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is a function with the appropriate signature, HandlerFunc(f) is a Handler[2] that calls f.

```go
type HandlerFunc func(ResponseWriter, *Request)
```



[2]

## type [Handler](https://golang.org/src/net/http/server.go?s=2756:2819#L76)

A Handler responds to an HTTP request.

ServeHTTP should write reply headers and data to the ResponseWriter and then return. Returning signals that the request is finished; it is not valid to use the ResponseWriter or read from the Request.Body after or concurrently with the completion of the ServeHTTP call.

Depending on the HTTP client software, HTTP protocol version, and any intermediaries between the client and the Go server, it may not be possible to read from the Request.Body after writing to the ResponseWriter. Cautious handlers should read the Request.Body first, and then reply.

Except for reading the body, handlers should not modify the provided Request.

If ServeHTTP panics, the server (the caller of ServeHTTP) assumes that the effect of the panic was isolated to the active request. It recovers the panic, logs a stack trace to the server error log, and either closes the network connection or sends an HTTP/2 RST_STREAM, depending on the HTTP protocol. To abort a handler so the client sees an interrupted response but the server doesn't log an error, panic with the value ErrAbortHandler.

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

 

[3]

## func [Handle](https://golang.org/pkg/net/http/#Handle)

```
func Handle(pattern string, handler Handler)
```

Handle registers the handler for the given pattern in the DefaultServeMux. The documentation for ServeMux explains how patterns are matched.



## Middleware

Middleware es una abstraccion de las rutas con las solicitudes del usuario.

![](Imágenes/4.png)

## Manejando request HTTP

```go
// Buscamos que el path existe, es decir, tiene un Handler asociado. Lo cual nos permitirá hacer operaciones con éste
func (r *Router) FindHandler(path string) (http.HandlerFunc, bool){
	handler, error := r.rules[path]
	return handler, error
}
```



## Asignando las rutas HTTP 

> file: router.go

```go
func (r *Router) ServeHTTP(w http.ResponseWriter, responde *http.Request) {
	handler, exist := r.FindHandler(responde.URL.Path)
	if !exist{
		w.WriteHeader(http.StatusNotFound)
	}else{
		handler(w, responde)
	}
}
```

> file: server.go

```go
func (s *Server) setHandler(path string, handle http.HandlerFunc){
   s.router.rules[path] = handle
}
```

> file: handlers.go

```go
func HandlerRoot(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintln(w, "Hello World!")
}

func HandlerAPI(w http.ResponseWriter, r *http.Request){
   fmt.Fprintln(w, "I am an API")
}
```

> file: main.go

```go
myServer := newServer(":3000")
myServer.setHandler("/", HandlerRoot)
myServer.setHandler("/api", HandlerAPI)
myServer.listen()
```

En el main nos encargamos de asignar qué ruta, va con qué Handler. Para ello introducimos en el `rules` de `Router` cada item (`string=path, http.HandlerFunc`). 

El método `setHandler` se encarga de vincular cada path con su respectivo Handler. Los Handlers están en el fichero `handlers.go`. 

El método `ServeHTTP` de `Router` es el engargado de llamar al respectivo Handle, en función del `Path` del Request.



## Middleware

```go
File: main.go
package main

import (
  "log"
  "net/http"
)

func middlewareOne(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println("Executing middlewareOne")
    next.ServeHTTP(w, r)
    log.Println("Executing middlewareOne again")
  })
}

func middlewareTwo(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println("Executing middlewareTwo")
    if r.URL.Path != "/" {
      return
    }
    next.ServeHTTP(w, r)
    log.Println("Executing middlewareTwo again")
  })
}

func final(w http.ResponseWriter, r *http.Request) {
  log.Println("Executing finalHandler")
  w.Write([]byte("OK"))
}

func main() {
  finalHandler := http.HandlerFunc(final)

  http.Handle("/", middlewareOne(middlewareTwo(finalHandler)))
  http.ListenAndServe(":3000", nil)
}
```

```bash
$ go run main.go
2014/10/13 20:27:36 Executing middlewareOne
2014/10/13 20:27:36 Executing middlewareTwo
2014/10/13 20:27:36 Executing finalHandler
2014/10/13 20:27:36 Executing middlewareTwo again
2014/10/13 20:27:36 Executing middlewareOne again
```

[Fuente](https://www.alexedwards.net/blog/making-and-using-middleware) 



### Agregando múltiples middlewares

Un Middleware recibe un Handler, el Handler que tiene que llamar después, es decir, el que se encuentra después de que el Middleware realice sus operaciones, y devuelve otro Handler: un Handler que ya sí llama al Handler final (en nuestro ejemplo a `HandlerAPI`)

```go
File: types.go
package main

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc
```

```go
File: middlewares.go
package main

import (
   "fmt"
   "net/http"
)

func checkAuth() Middleware{
   return func(f http.HandlerFunc) http.HandlerFunc{
      return func (w http.ResponseWriter, r *http.Request){
         flag := true
         fmt.Println("Checking authentication")
         if flag{
            f(w,r)
         }else{
            return
         }
      }
   }
}
```

```go
File: server.go
func (s *Server) addMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc{
   for _, middleware := range middlewares{
      f = middleware(f)
   }
   return f
}
```

```go
File: main.go
package main

func main() {
   myServer := newServer(":3000")
   myServer.setHandler("/", HandlerRoot)
   myServer.setHandler("/api", myServer.addMiddleware(HandlerAPI, checkAuth()))
   myServer.listen()
}
```

![](Imágenes/6.jpg)

El middleware `checkAuth` se situará entre el Request y el `HandlerAPI`, con el propósito de checkear (en este caso hace una simulación) si el usuario está logeado. 

En la entrada de nuestro programa, en el `main.go`, añadimos un handler a nuestros Server, `myServer`, con `setHandler`. Ahora, como Handler vamos a llamar antes a `addMiddleware`, que hará el recorrido que se muestra en la imagen de arriba. 



### Agregando otro middleware

```go
File: middlewares.go
func Loggin() Middleware{
   return func(f http.HandlerFunc) http.HandlerFunc{
      return func(w http.ResponseWriter, r *http.Request){
         start := time.Now()
         defer func(){
            log.Println(r.URL.Path, time.Since(start)) // añade la fecha en la impresión
         }()
         f(w,r)
      }
   }
}
```



## Manejando distintas peticiones, POST

```go
File: router.go
package main

import (
   "net/http"
)

/**
   Maneja las URLs que nos llegan
 */
type Router struct {
   rules map[string]map[string]http.HandlerFunc
}

func newRouter() *Router {
   return &Router{
      rules: make(map[string]map[string]http.HandlerFunc),
   }
}

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool){
   _, pathExist := r.rules[path]
   handler, methodExist := r.rules[path][method]
   return handler, methodExist, pathExist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, response *http.Request) {
   handler, methodExist, pathExist := r.FindHandler(response.URL.Path, response.Method)
   if !pathExist{
      w.WriteHeader(http.StatusNotFound)
   }else if !methodExist{
      w.WriteHeader(http.StatusMethodNotAllowed)
   }else{
      handler(w, response)
   }
}
```

```go
File: server.go
func (s *Server) setHandler(method string, path string, handle http.HandlerFunc){
   _, pathExist:= s.router.rules[path]
   if !pathExist{
      s.router.rules[path] = make(map[string]http.HandlerFunc)
   }
   s.router.rules[path][method] = handle
}
```

```go
File: main.go
func main() {
   myServer := newServer(":3000")
   myServer.setHandler("GET", "/", HandlerRoot)
   myServer.setHandler("POST", "/api", myServer.addMiddleware(HandlerAPI, checkAuth(), Loggin()))
   myServer.listen()
}
```

Apoyándonos en [Postman](https://chrome.google.com/webstore/detail/postman/fhbjgbiflinjbdggehcddcbncdddomop?hl=es-419) para hacer las distintas peticiones



## Post Request con JSON

```go
File: handlers.go
func PostRequest(w http.ResponseWriter, r *http.Request){
   decoder := json.NewDecoder(r.Body)
   var metaData MetaData
   error := decoder.Decode(&metaData)
   if error != nil {
      fmt.Fprintf(w, "error: %v", error)
      return
   }else{
      fmt.Fprintf(w, "Payload: %v", metaData)
   }
}
```

![](Imágenes/5.png)



## Respondiendo al Resquest JSON

```go
File: handlers.go
func UserPostRequest(w http.ResponseWriter, r *http.Request){
   decoder := json.NewDecoder(r.Body)
   var user User
   error := decoder.Decode(&user)
   if error != nil {
      fmt.Fprintf(w, "error: %v", error)
      return
   }else{
      response, error := user.toJSON()
      if error != nil{
         w.WriteHeader(http.StatusInternalServerError)
         return
      }
      w.Header().Set("Content-Type", "application/json")
      w.Write(response)
   }
}
```

```go
File: types.go
type User struct {
   Name  string `json:"name"`
   Email string `json:"email"`
   Phone string `json:"phone"`
}

func (u *User) toJSON() ([]byte, error){
   return json.Marshal(u) // nos convierte el struct en JSON (atributos)
}
```

![](Imágenes/7.png)



## Crear un módulo en go 

```bash
git mod init [respotiry-url]
```



## Incluyendo los mods 

En la terminal:

````bash
go mod init [name-mod]
````

En el archivo `go.mod`:

```go
requiere [repository-url] [version]
```

En nuestro `main.go`: 

```go
import "[repository-url]"
```



