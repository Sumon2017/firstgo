package main

import (
  "fmt"
  "net/http"
  "os"
  "os/signal"
  "syscall"
)

func main(){
  fmt.Println("program started")
  ch := make(chan os.Signal, 1)
  signal.Notify(ch , syscall.SIGINT , syscall.SIGTERM)
  router := http.NewServeMux()
  router.HandleFunc("/",hd1)
  router.HandleFunc("/sumon",hd2)

  go func(){
    <-ch
    fmt.Println("")
    fmt.Println("program closed gracefully")
    os.Exit(0)
  }()

  err := http.ListenAndServe(":3000",router)
  if err != nil {
    fmt.Println("Error: could not open server")
  }
}

func hd1(w http.ResponseWriter , r *http.Request){
  w.Write([]byte("hello from golang"))
}
func hd2(w http.ResponseWriter , r *http.Request){
  w.Write([]byte("hello from sumon"))
}
