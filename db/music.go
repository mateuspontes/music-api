package db

import (
	"github.com/mateuspontes/music-api/music"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const MusicCollection = "musics"

type MusicRepository struct {
	session *mgo.Session
}

func NewMusicRepository(session *mgo.Session) *MusicRepository {
	return &MusicRepository{session}
}

func (r *MusicRepository) FindAll() ([]*music.Music, error) {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(MusicCollection)
	query := bson.M{}

	documents := make([]*music.Music, 0)
	err := collection.Find(query).All(&documents)
	return documents, err
}

func (r *MusicRepository) FindById(id string) (*music.Music, error) {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(MusicCollection)
	query := bson.M{"_id": id}

	music := &music.Music{}

	err := collection.Find(query).One(music)

	return music, err
}
