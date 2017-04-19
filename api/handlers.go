package api

import (
	"encoding/json"
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/mateuspontes/music-api/db"
)

type ListMusicHandler struct {
	repository *db.MusicRepository
}

func (h *ListMusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	musics, err := h.repository.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(musics)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type GetMusicHandler struct {
	repository *db.MusicRepository
}

func (h *GetMusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
	music, err := h.repository.FindById(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(music)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
