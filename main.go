package main

import (
  "fmt"
  "net/http"
)

func main(){
  fmt.Println("yes")
  router := http.NewServeMux()
  router.HandleFunc("/",hd1)
  router.HandleFunc("/sumon",hd2)

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
