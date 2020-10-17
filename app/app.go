package app

import(
  "fmt"
  "net/http"
  "os"

  "github.com/kokoqalb/noor-core/router"
)

type App interface {
  Addr() string
  Mux() http.Handler
}

func New(r router.Router) App {
  app := new(app)
  addr := os.Getenv("PORT")
  if addr = "" {
    log.Println("$PORT not set -- using 8080 by default")
    addr = "8080"
}
  app.addr = fmt.Sprintf(":%s", addr)
  router.Route()
}

