package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jaswanth99999/MovieManagement/pkg/models"
	"github.com/jaswanth99999/MovieManagement/pkg/utils"
)

var NewMovie models.Movie

func GetMovie(w http.ResponseWriter, r *http.Request) {
	NewMovies := models.GetAllMovies()
	res, _ := json.Marshal(NewMovies)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	Id, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error While parsing")
	}
	movieDetails, _ := models.GetMovieById(Id)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	CreateMovie := &models.Movie{}
	utils.ParseBody(r, CreateMovie)
	m := CreateMovie.CreateMovie()
	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	movie := models.DeleteMovie(ID)
	res, _ := json.Marshal(movie)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var updateMovie = &models.Movie{}
	utils.ParseBody(r, updateMovie)
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	movieDetails, db := models.GetMovieById(ID)
	if updateMovie.Name != "" {
		movieDetails.Name = updateMovie.Name
	}
	if updateMovie.Hero != "" {
		movieDetails.Hero = updateMovie.Hero
	}
	if updateMovie.Director != "" {
		movieDetails.Director = updateMovie.Director
	}
	db.Save(&movieDetails)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
