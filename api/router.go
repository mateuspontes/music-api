package api

import (
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/dimfeld/httptreemux"
	"github.com/mateuspontes/music-api/db"
)

func NewRouter() http.Handler {
	router := httptreemux.NewContextMux()
	session, err := mgo.Dial("localhost:27017/musics")
	if err != nil {
		log.Fatal(err)
	}

	repository := db.NewMusicRepository(session)

	router.Handler(http.MethodGet, "/api/musics", &ListMusicHandler{repository})
	router.Handler(http.MethodGet, "/api/musics/:id", &GetMusicHandler{repository})

	return router
}
