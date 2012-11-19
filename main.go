package mvc

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var Env struct {
	RootPath string
}

type RouteHandler struct {
	Version int
}

//dummy types
type Result string
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

type Action func(c Context) Result

var routes = make(map[string]Action)

func ContextBuilder(w http.ResponseWriter, r *http.Request) Context {
	return Context{w, r}
}
func (handler *RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if a, ok := routes[path]; ok {
		log.Println("found action for", path)
		fmt.Fprintln(w, a(ContextBuilder(w, r)))
		return
	}
	log.Println("no action found for", path)
	http.NotFound(w, r)
}

func MapRoute(pattern string, handler Action) {
	routes[pattern] = handler
}

func StartServer(port string) {
	log.Printf("Starting server on http://localhost%s/", port)
	handler := RouteHandler{Version: 1}
	if err := http.ListenAndServe(port, &handler); err != nil {
		panic(err)
	}
}

//action builder
func FileResultAction(filename string) Action {
	return func(ctx Context) Result {
		bytes, _ := ioutil.ReadFile("views/" + filename)
		return Result(bytes)
	}
}

//initialization code
func init() {
	Env.RootPath, _ = os.Getwd()
}
