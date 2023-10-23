package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Users struct {
	Name   string
	Skills string
	Age    int
}

func Index(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")

	user := Users{"Antonio Bravo", "Desarrollador FullStack", 39}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, user)
	}
}
func main() {
	http.HandleFunc("/", Index)
	//creamos el servidor
	fmt.Println("El servidor esta corriendo en el puerto 3000")
	fmt.Println("Ir al sitio web: http://localhost:3000")
	http.ListenAndServe("localhost:3000", nil)
}
