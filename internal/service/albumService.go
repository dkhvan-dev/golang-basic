package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testBasicHttp/internal/model"
	"testBasicHttp/internal/utils"
)

var (
	id     = 4
	albums = []model.Album{
		{Id: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{Id: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{Id: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
)

type AlbumCreateDto struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		handlerMsg := []byte(`{
			"success": false,
			"message": "Check your HTTP method: Invalid HTTP method executed",
		}`)

		utils.ReturnJsonResponse(w, 400, handlerMsg)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(albums)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAlbumById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		handlerMsg := []byte(`{
			"message": "Check your HTTP method: Invalid HTTP method executed",
		}`)

		utils.ReturnJsonResponse(w, 400, handlerMsg)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	for _, album := range albums {
		if album.Id == id {
			json.NewEncoder(w).Encode(album)
			break
		}
	}

}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		handlerMsg := []byte(`{
			"success": false,
			"message": "Check your HTTP method: Invalid HTTP method executed",
		}`)

		utils.ReturnJsonResponse(w, 400, handlerMsg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var dto AlbumCreateDto

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &dto)
	if err != nil {
		log.Fatal(err)
	}

	entity := model.Album{
		Id:     id,
		Title:  dto.Title,
		Artist: dto.Artist,
		Price:  dto.Price,
	}

	id += 1

	albums = append(albums, entity)
	json.NewEncoder(w).Encode(entity)
}
