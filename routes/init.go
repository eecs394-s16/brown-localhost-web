package routes

import (
  "net/http"
  // "html/template"
  // "path"

  "github.com/gorilla/mux"

  "github.com/codegangsta/negroni"
)

// type Book struct {
//     Title  string
//     Author string
// }

func GetRouter() *negroni.Negroni {
  n  := negroni.New()
  n.Use(negroni.HandlerFunc(configureResponseMiddleware))

  r  := mux.NewRouter()
  addSongRoutes(r)

  n.UseHandler(r)
  return n
}

func configureResponseMiddleware(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
  // book := Book{"test", "Michael"}
  // fp := path.Join("web", "index.html")
  //
  // tmpl, err := template.ParseFiles(fp)
  // if err != nil {
  //     http.Error(w, err.Error(), http.StatusInternalServerError)
  //     return
  // }
  //
  // if err := tmpl.Execute(w, book); err != nil {
  //     http.Error(w, err.Error(), http.StatusInternalServerError)
  // }

  w.Header().Set("Content-Type", "application/json")
  next(w, req)
}
