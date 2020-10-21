package main 

import (

   "fmt"
   "long"
   "net/http"
)

  func homePage(w http.ResponseWriter, r *http.Request) {
   fmt.Printf(w, "Homepage Endpoint Hit")
}

 func handlerRequests() {
 http.Handlerfunc("/", homepage)
llog.fatal(http.ListerAndServer(":8081", nil))
}
func main() {
  handle Requests()
}

