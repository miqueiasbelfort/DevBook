package router

import "github.com/gorilla/mux"

//Retornar um router com todas as rotas configuradas
func GerarRouters() *mux.Router {
	return mux.NewRouter()
}
