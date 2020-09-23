package main

import(
  "fmt"
  "log"
  "os"
  "net/http"
)

func indexHandler(writer http.ResponseWriter, req *http.Request) {
  fmt.Fprintln(writer, "hello, world")
}

func main() {
  http.HandleFunc("/", indexHandler)

  port := os.Getenv("PORT")

  // Set port to 8080 if not defined
  if defaultPort := "8080"; port == "" {
    log.Println("$PORT not set - using", defaultPort, "by default")
    port = defaultPort
  }
  log.Println("Listening on port", port)

  // http.ListenAndServe "always returns a non-nil error"??
  if err := http.ListenAndServe(":" + port, nil); err != nil {
    log.Fatal(err)
  }
}
