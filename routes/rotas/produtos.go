package rotas

import "module/controllers"

var RotasProdutos = []Rota{
	{
		URI:    "/",
		Funcao: controllers.Index,
	},
	{
		URI:    "/new",
		Funcao: controllers.New,
	},
	{
		URI:    "/insert",
		Funcao: controllers.Insert,
	},
	{
		URI:    "/delete",
		Funcao: controllers.Delete,
	},
	{
		URI:    "/edit",
		Funcao: controllers.Edit,
	},
	{
		URI:    "/update",
		Funcao: controllers.Update,
	},
}
