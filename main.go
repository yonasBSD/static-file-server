package main

/*
import (
  "net/http"
  "fmt"
  "log"
)

func main() {
  port := 9999

  fileServer := http.FileServer(http.Dir("."))
  http.Handle("/", fileServer)

  fmt.Printf("Server started at :%d\n", port)

  if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
    log.Fatal(err)
  }
}
*/

import (
  "net"
  "net/http"
  "flag"
  "fmt"
  "os"
)

var address string
var port int
var directoryPath string

func init() {
  flag.StringVar(&directoryPath, "dir", ".", "directory to serve")
  flag.StringVar(&address, "address", "0.0.0.0", "port number")
  flag.IntVar(&port, "port", 0, "port number")
}

func main() {
  flag.Parse()

  // Check if the directory exists
  _, err := os.Stat(directoryPath)
  if os.IsNotExist(err) {
    fmt.Printf("Directory '%s' not found.\n", directoryPath)
    return
  }

  // Create a file server handler to serve the directory's contents
  fileServer := http.FileServer(http.Dir(directoryPath))

  // Create a new HTTP server and handle requests
  http.Handle("/", fileServer)

  // Start the server
  listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
  if err != nil {
    panic(err)
  }

  port = listener.Addr().(*net.TCPAddr).Port

  fmt.Printf("Server started at http://%s:%d\n", address, port)
  err = http.Serve(listener, nil)

  if err != nil {
    fmt.Printf("Error starting server: %s\n", err)
  }
}
