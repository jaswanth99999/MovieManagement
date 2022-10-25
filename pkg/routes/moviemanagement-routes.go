package routes

import (
	"github.com/gorilla/mux"
	"github.com/jaswanth99999/MovieManagement/pkg/controllers"
)

var RegisterMoviesRoutes = func(router *mux.Router) {
	router.HandleFunc("/movie/", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/", controllers.GetMovie).Methods("GET")
	router.HandleFunc("/movie/{movieId}", controllers.GetMovieById).Methods("GET")
	router.HandleFunc("/movie/{movieId}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{movieId}", controllers.DeleteMovie).Methods("DELETE")
}
