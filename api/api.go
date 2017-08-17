package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guilhermebr/minesweeper/minesweeper"
	"github.com/guilhermebr/minesweeper/storage/memory"
	"github.com/guilhermebr/minesweeper/types"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

type Services struct {
	logger      *logrus.Logger
	GameService types.GameService
}

func Start(log *logrus.Logger) error {

	db := memory.New()

	services := Services{
		logger: log,
		GameService: &minesweeper.GameService{
			Store: memory.NewGameStore(db),
		},
	}
	// API Routes
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", services.healthcheck).Methods("GET")
	r.HandleFunc("/game", services.createGame).Methods("POST")

	// Middleware
	n := negroni.Classic()
	n.UseHandler(r)

	//Run Server
	log.Infoln("Server running on port :3000")
	http.ListenAndServe(":3000", n)
	return nil
}
