package main

import (
	"module/routes"
	"net/http"
)

func main() {
	r := routes.GerarRotas()
	http.ListenAndServe(":8000", r) // ListenAndServe - Ouça e sirva a requisição

}
