package controllers

import (
	"log"
	"module/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "index.html", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	//novoProduto := models.CriarNovoProduto()
	temp.ExecuteTemplate(w, "new.html", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}
		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}
		models.CriaNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}