package main
import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
)

type NameInput struct {
    Name string `json:"name"`
  }
type NameResponse struct {
    Greeting string `json:"greeting"`
}
func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "name: , %s!\n", p.ByName("name"))
}

func sayhello(rw http.ResponseWriter, req *http.Request, p httprouter.Params){
    var n NameInput
    var r NameResponse
    decoder := json.NewDecoder(req.Body)
    fmt.Println(req.Body)

    err1 := decoder.Decode(&n)
    if err1 != nil {
    panic(err1)
    }

    r.Greeting="Hello,"+n.Name

    output,_ := json.Marshal(r)
    fmt.Fprintf(rw,string(output))
}

func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", hello)
    mux.POST("/sayhello", sayhello)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}