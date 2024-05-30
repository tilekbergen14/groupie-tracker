package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/tilekbergen14/groupie-tracker/types"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}

func (s *Server) Start() error {
	// handler := http.StripPrefix("/static/", http.FileServer(http.Dir("web")))
	// http.Handle("/static/", handler)
	http.HandleFunc("/", s.handleHome)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Cannot get info!")
		return
	}

	var artists []types.Artist
	err1 := json.NewDecoder(res.Body).Decode(&artists)
	if err1 != nil {
		fmt.Println("Cannot get info")
		return
	}
	t, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t.Execute(w, artists)
}
