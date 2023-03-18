package main

import (
	"fmt"
	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
	"html/template"
	"log"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	cardDetailsTmpl := template.Must(template.ParseFiles("templates/card-details.html"))
	nameSearchTmpl := template.Must(template.ParseFiles("templates/name-search.html"))
	setSearchTmpl := template.Must(template.ParseFiles("templates/set-search.html"))
	typeSearchTmpl := template.Must(template.ParseFiles("templates/type-search.html"))

	c := tcg.NewClient("82a5348e-9d27-4af9-9d49-3886047cae6d")

	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing template: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/card-details", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		if id == "" {
			http.Error(w, "Missing card ID", http.StatusBadRequest)
			return
		}

		card, err := c.GetCardByID(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching card details from API: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		err = cardDetailsTmpl.Execute(w, card)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing template: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/name-search", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		var query string
		if name != "" {
			query = fmt.Sprintf("name:%s", name)
		}

		cards, err := c.GetCards(
			request.Query(query),
			request.PageSize(250),
		)
		if err != nil {
			log.Fatalf("Error getting cards: %s", err.Error())

		}
		err = nameSearchTmpl.Execute(w, cards)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing template: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/set-search", func(w http.ResponseWriter, r *http.Request) {
		setName := r.URL.Query().Get("set-name")

		var query string
		if setName != "" {
			query = fmt.Sprintf("set.name:%s", setName)
		}

		sets, err := c.GetSets(
			request.Query(query),
			request.PageSize(100),
		)
		if err != nil {
			log.Fatalf("Error getting sets: %s", err.Error())
		}

		err = setSearchTmpl.Execute(w, sets)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing template: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/type-search", func(w http.ResponseWriter, r *http.Request) {
		types := r.URL.Query().Get("type")
		var query string
		if types != "" {
			query = fmt.Sprintf("types:%s", types)
		}

		cards, err := c.GetCards(
			request.Query(query),
			request.PageSize(250),
		)
		if err != nil {
			log.Fatalf("Error getting cards: %s", err.Error())

		}
		err = typeSearchTmpl.Execute(w, cards)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing template: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting HTTP server: %s", err.Error())
	}
}
