package main

import (
	"module/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil) // ListenAndServe - Ouça e sirva a requisição

}
