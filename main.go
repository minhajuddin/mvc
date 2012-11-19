package mvc

import (
  "log"
	"fmt"
	"net/http"
)

type RouteHandler struct {
	Version int
}

//dummy types
type Result string
type Context string

type Action func(c Context) Result

var routes = make(map[string]Action)

func (handler *RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path
  if a, ok := routes[path]; ok {
    log.Println("found action for", path)
		fmt.Fprintln(w, a(Context("Some stuff")))
		return
	}
  log.Println("no action found for", path)
	http.NotFound(w, r)
}

func MapRoute(pattern string, handler Action){
  routes[pattern] = handler
}

func StartServer(port string){
  log.Printf("Starting server on http://localhost%s/", port)
	handler := RouteHandler{Version: 1}
	if err := http.ListenAndServe(port, &handler); err != nil {
		panic(err)
	}
}

//initialization code
func init(){
}
