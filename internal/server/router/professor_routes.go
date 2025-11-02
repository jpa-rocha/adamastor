package router

import (
	"log"
	"net/http"

	"github.com/jpa-rocha/adamastor/internal/server/templates"
	"github.com/jpa-rocha/adamastor/internal/server/utilities"
)

func HandleLittleProfessor(w http.ResponseWriter, r *http.Request) {
	err := templates.Professor().Render(r.Context(), w)
	if err != nil {
		log.Println("error rendering content")
	}
}

func startProfessor(w http.ResponseWriter, r *http.Request) {
	game := utilities.Game{}
	game.PrepareRound(r)
	game.MakeRound()
	err := templates.StartProfessor(game).Render(r.Context(), w)
	if err != nil {
		log.Println("error rendering content")
	}
}

//nolint:gocritic //dont want switch statement for result creation
func playRound(w http.ResponseWriter, r *http.Request) {
	game := utilities.Game{}
	game.PrepareRound(r)
	if game.Answer == game.Result {
		game.Win++
	} else {
		game.Loss++
	}
	game.NumRounds--
	if game.NumRounds == 0 {
		if game.Win > game.Loss {
			game.Final = 1
		} else if game.Win < game.Loss {
			game.Final = -1
		} else {
			game.Final = 0
		}
		err := templates.ShowResults(game).Render(r.Context(), w)
		if err != nil {
			log.Println("error rendering content")
		}
		return
	}
	game.MakeRound()
	err := templates.PlayRound(game).Render(r.Context(), w)
	if err != nil {
		log.Println("error rendering content")
	}
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	err := templates.Layout(templates.Index()).Render(r.Context(), w)
	if err != nil {
		log.Println("error rendering content")
	}
}
