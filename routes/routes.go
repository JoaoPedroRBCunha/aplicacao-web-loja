package routes

import (
	"module/routes/rotas"

	"github.com/gorilla/mux"
)

func GerarRotas() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
