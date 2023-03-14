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

	c := tcg.NewClient("82a5348e-9d27-4af9-9d49-3886047cae6d")

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing template: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		setName := r.URL.Query().Get("set-name")
		cardType := r.URL.Query().Get("types")

		var query string
		if name != "" && setName != "" && cardType != "" {
			query = fmt.Sprintf("name:\"%s\" set.name:\"%s\" types:\"%s\"", name, setName, cardType)
		} else if name != "" && cardType != "" {
			query = fmt.Sprintf("name:\"%s\" types:\"%s\"", name, cardType)
		} else if setName != "" && cardType != "" {
			query = fmt.Sprintf("set.name:\"%s\" types:\"%s\"", setName, cardType)
		} else if name != "" {
			query = fmt.Sprintf("name:\"%s\"", name)
		} else if setName != "" {
			query = fmt.Sprintf("set.name:\"%s\"", setName)
		} else if cardType != "" {
			query = fmt.Sprintf("types:\"%s\"", cardType)
		}

		cards, err := c.GetCards(
			request.Query(query),
			request.PageSize(7),
		)

		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching data from API: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, cards)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing template: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		fmt.Println(query)
		for _, card := range cards {
			fmt.Println(card.Name)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting HTTP server: %s", err.Error())
	}
}
