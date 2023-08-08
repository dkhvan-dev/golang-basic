package controllers

import (
	"net/http"
	"testBasicHttp/internal/service"
)

func New() {
	http.HandleFunc("/albums", service.GetAlbums)
	http.HandleFunc("/albums/{id}", service.GetAlbumById)
	http.HandleFunc("/albums/add", service.CreateAlbum)
}
