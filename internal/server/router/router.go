package router

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/jpa-rocha/adamastor/internal/server/templates"
	"github.com/jpa-rocha/adamastor/public"
)

type Router struct {
	Mux *http.ServeMux
	Err error
}

func NewRouter() *Router {
	router := Router{
		Mux: http.NewServeMux(),
		Err: nil,
	}
	return &router
}

func (r *Router) SetStaticPaths() {
	static := []string{
		"/static/",
	}

	const base = "assets"
	staticFS, err := fs.Sub(public.Assets, base)
	if err != nil {
		r.Err = err
		return
	}
	for _, path := range static {
		r.Mux.Handle(path, http.FileServer(http.FS(staticFS)))
	}
}

func (r *Router) HandleRoutes() {
	r.SetStaticPaths()

	// Main Routes
	r.Mux.HandleFunc("/cv", HandleCV)
	r.Mux.HandleFunc("/projects", HandleProjects)
	r.Mux.HandleFunc("/about", HandleAbout)

	// Little Professor Routes
	r.Mux.HandleFunc("/littleprofessor", HandleLittleProfessor)
	r.Mux.HandleFunc("/start_professor", startProfessor)
	r.Mux.HandleFunc("/play_round", playRound)
	r.Mux.HandleFunc("/", HandleIndex)
}

// TODO: Check templ Render error and display 500 in case of error
func HandleCV(w http.ResponseWriter, r *http.Request) {
	err := templates.Layout(templates.CV()).Render(r.Context(), w)
	if err != nil {
		log.Println("error rendering content")
	}
}

func HandleProjects(w http.ResponseWriter, r *http.Request) {
	err := templates.Layout(templates.Project()).Render(r.Context(), w)
	if err != nil {
		log.Println("error rendering content")
	}
}

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	err := templates.Layout(templates.About()).Render(r.Context(), w)
	if err != nil {
		log.Println("error rendering content")
	}
}
