package mvc

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
)

type MvcHandler struct {
	Version int
}

//Static file handler
func staticFileHandler(w http.ResponseWriter, r *http.Request) bool {
	urlpath := r.URL.Path[1:]
	if len(urlpath) == 0 {
		return false
	}
	filepath := path.Join(Env.PublicDirectoryPath, urlpath)
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	w.Header().Add("Content-Type", mime.TypeByExtension(path.Ext(filepath)))
	fd, err := os.Open(filepath)
	io.Copy(w, fd)
	return true
}

//Core request handler and dispatcher
func (handler *MvcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if staticFileHandler(w, r) {
		return
	}
	if a, ok := routes[path]; ok {
		log.Println("found action for", path)
		fmt.Fprintln(w, a(ContextBuilder(w, r)))
		return
	}
	log.Println("no action found for", path)
	http.NotFound(w, r)
}
