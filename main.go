package mvc

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

var Env struct {
	RootPath            string
	PublicDirectoryPath string
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

func MapRoute(pattern string, handler Action) {
	routes[pattern] = handler
}

func StartServer(port string) {
	log.Printf("Starting server on http://localhost%s/", port)
	handler := MvcHandler{Version: 1}
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
	Env.PublicDirectoryPath = path.Join(Env.RootPath, "public")
}
