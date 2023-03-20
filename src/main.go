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

	// loads the HTML templates from the files in the templates directory
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	cardDetailsTmpl := template.Must(template.ParseFiles("templates/card-details.html"))
	nameSearchTmpl := template.Must(template.ParseFiles("templates/name-search.html"))
	setSearchTmpl := template.Must(template.ParseFiles("templates/set-search.html"))
	typeSearchTmpl := template.Must(template.ParseFiles("templates/type-search.html"))
	setCardsTmpl := template.Must(template.ParseFiles("templates/set-cards.html"))

	// API key, left there on purpose for the person correcting me, will regenerate a new one once it is done
	c := tcg.NewClient("82a5348e-9d27-4af9-9d49-3886047cae6d")

	// handles the static files in the assets directory
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	// handles routing and errors
	// 🠗 🠗 🠗 🠗 🠗 🠗 🠗 🠗 🠗 🠗 🠗 🠗 🠗 🠗 🠗

	// index
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing template: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	})

	// handles requests to get card details from fetched id
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

	// handles search queries by name
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

	// handles search queries by expansion pack
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

	// loads cards from queried expansion pack
	http.HandleFunc("/set-cards", func(w http.ResponseWriter, r *http.Request) {
		setCard := r.URL.Query().Get("set-name")

		query := fmt.Sprintf("set.name:%s", setCard)

		cards, err := c.GetCards(
			request.Query(query),
			request.PageSize(250),
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting cards: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		err = setCardsTmpl.Execute(w, cards)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing template: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	})

	// handles search queries by types
	http.HandleFunc("/type-search", func(w http.ResponseWriter, r *http.Request) {
		types := r.URL.Query().Get("types")
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

	// launches the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting HTTP server: %s", err.Error())
	}
}
