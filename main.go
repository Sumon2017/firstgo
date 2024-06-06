package main

import (
  "fmt"
  "net/http"
)

func main(){
  fmt.Println("yes")
  router := http.NewServeMux()
  router.HandleFunc("/",hd1)

  err := http.ListenAndServe(":3000",router)
  if err != nil {
    fmt.Println("could not open server")
  }else{
    fmt.Println("server started on 3000")
  }
}

func hd1(w http.ResponseWriter , r *http.Request){
  w.Write([]byte("hello from golang"))
}
