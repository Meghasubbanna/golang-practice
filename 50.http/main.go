package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Server started listening ")
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})
	http.HandleFunc("/greet", Sayhi)
	http.HandleFunc("/greetby", GreetBy)
	err := http.ListenAndServe("0.0.0.0:9091", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(err, "Hello World")
}
func GreetBy(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("Hit here?")
		g := Greet{}
		err := json.NewDecoder(r.Body).Decode(&g)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
		} else {
			fmt.Fprintln(w, "data has been read", g)
		}
	} else {
		w.Write([]byte("Unimpolemented method"))
		w.WriteHeader(http.StatusNotImplemented)

	}
}
func Sayhi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Heloo VSCO")

}

//	func AddEmployee(w http.ResponseWriter,r*http.Request){
//		if r.Method=="POST"{
//			if bytes,err:=io.ReadAll(r.Body);err!=nil{
//				w.WriteHeader(http.StatusBadRequest)
//				w.Write([]byte())
//			}
//		}
//	}
type Greet struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

// type Employee struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"name"`
// }
