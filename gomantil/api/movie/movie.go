package movie

import (
	"github.com/mantil-io/mantil.go"
)

type Movie struct {
	kv *mantil.KV
}

type MovieItem struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Rating string `json:"rating"`
}

func New() *Movie {
	kv, _ := mantil.NewKV("movies")
	return &Movie{
		kv: kv,
	}
}
